package salesorderrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type Repository struct {
	logger logging.Logger
	goquDB *goqu.Database
}

func New(
	logger logging.Logger,
	goquDB *goqu.Database,
) *Repository {
	return &Repository{
		logger: logger,
		goquDB: goquDB,
	}
}

func (r *Repository) Insert(
	data rdbms.SalesOrder,
) error {
	query := r.goquDB.Insert(
		TABLE,
	).Prepared(true).Rows(
		goqu.Record{
			ID:                      data.Id,
			ACCOUNT_ID:              data.AccountId,
			PO_ID:                   data.POId,
			FLOW_INSTANCE_ID:        data.FlowInstanceId,
			FLOW_INSTANCE_PARAMS_ID: data.FlowInstanceParamsId,
			SO_NUMBER:               data.SONumber,
			SUPPLIER_ID:             data.SupplierId,
			DUE_DATE:                data.DueDate,
			CREATED_BY:              data.CreatedBy,
			STATUS:                  data.Status,
		},
	)
	_, err := query.Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectByPOId(
	id string,
) ([]rdbms.SalesOrder, error) {
	var salesOrder []rdbms.SalesOrder
	if err := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ID,
		PO_ID,
		SO_NUMBER,
		FLOW_INSTANCE_ID,
		FLOW_INSTANCE_PARAMS_ID,
		SUPPLIER_ID,
		DUE_DATE,
		CREATED_BY,
		CREATED_ON,
		STATUS,
	).Where(
		goqu.C(PO_ID).Eq(id),
	).ScanStructs(&salesOrder); err != nil {
		return []rdbms.SalesOrder{}, err
	}
	return salesOrder, nil
}

func (r *Repository) SelectBySOId(
	id string,
) (rdbms.SalesOrder, error) {
	var salesOrder rdbms.SalesOrder
	if _, err := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ID,
		PO_ID,
		SO_NUMBER,
		SUPPLIER_ID,
		FLOW_INSTANCE_ID,
		FLOW_INSTANCE_PARAMS_ID,
		DUE_DATE,
		STATUS,
		CREATED_BY,
	).Where(
		goqu.C(ID).Eq(id),
	).ScanStruct(&salesOrder); err != nil {
		return rdbms.SalesOrder{}, err
	}
	return salesOrder, nil
}

func (r *Repository) CheckSOExistsByCode(code string) (bool, error) {
	var so rdbms.SalesOrder
	exists, err := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ID,
	).Where(goqu.C(SO_NUMBER).Eq(code)).ScanStruct(&so)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *Repository) DeleteById(
	id string,
) error {
	if _, er := r.goquDB.Delete(
		TABLE,
	).Prepared(true).Where(goqu.C(ID).Eq(id)).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

func (r *Repository) UpdateBySOId(
	id string,
	data rdbms.SalesOrder,
) error {
	if _, er := r.goquDB.Update(
		TABLE,
	).Prepared(true).Set(
		goqu.Record{
			SO_NUMBER:   data.SONumber,
			MODIFIED_BY: data.ModifiedBy,
			DUE_DATE:    data.DueDate,
		},
	).Where(goqu.C(ID).Eq(id)).Executor().Exec(); er != nil {
		fmt.Println(er, "eror")
		return er
	}
	fmt.Println("updated")
	return nil
}

func (r *Repository) SelectMany(
	limit uint,
	offset uint,
	search ports.OrdersSearch,
	sort ports.OrdersSort,
	filters ports.OrdersFilters,
	account_id int,
) ([]rdbms.SalesOrder, error) {
	var orders []rdbms.SalesOrder

	w := []exp.Expression{}
	// fmt.Println(filters.IsConsignee, "filters.IsConsignee")
	temp := false

	fmt.Println(filters.User, "status")
	if filters.Status != nil {
		if *filters.Status == "To_be_Confirmed" {
			temp = true
		}
	}
	if filters.User == "buyer" {
		w = append(w, goqu.C(ACCOUNT_ID).Eq(account_id))
		fmt.Println(filters.User, w, account_id, "yesssss")

	} else if filters.User == "Supplier" {
		fmt.Println(filters.User, account_id, "ytuutuss")

		w = append(w, goqu.C(SUPPLIER_ID).Eq(account_id))
	}
	if temp || filters.Status == nil {
		fmt.Println("hereee1")
		w = append(w, goqu.C(STATUS).Neq("Confirmed"))
		w = append(w, goqu.C(STATUS).Neq("Completed"))
		w = append(w, goqu.C(STATUS).Neq("Trash"))
	} else {
		fmt.Println("hereee2")
		w = append(w, goqu.C(STATUS).Eq(*filters.Status))
	}

	if search.Name != nil {
		searchClauses := r.buildSearchWhereClauses(*search.Name)

		if len(searchClauses) > 1 {
			w = append(w, goqu.Or(searchClauses...))
		} else if len(searchClauses) == 1 {
			w = append(w, searchClauses[0])
		}
	}

	query := r.goquDB.From(TABLE).Prepared(true).Select(
		ID,
		ACCOUNT_ID,
		PO_ID,
		SO_NUMBER,
		SUPPLIER_ID,
		DUE_DATE,
		CREATED_BY,
		CREATED_ON,
		STATUS,
		FLOW_INSTANCE_ID,
		FLOW_INSTANCE_PARAMS_ID,
		// BOOKING_AGENT_ID,
		// BOOKING_STATUS,
		// TRACKING_ID,
	).Where(w...).Order(
		r.getOrderedExpression(sort),
	).Limit(
		limit,
	).Offset(
		offset,
	)

	err := query.ScanStructs(&orders)

	if err != nil {
		querySQL, params, qerr := query.ToSQL()

		r.logger.Error(
			"CONTACTS_LIST_SELECTION_FAILED",
			"there was an issue when updating the contact",
			err,
			map[string]any{},
			map[string]any{
				"query":  querySQL,
				"params": params,
				"error":  qerr,
			},
		)

		return []rdbms.SalesOrder{}, err
	}

	return orders, nil
}

func (r *Repository) getOrderedExpression(
	sort ports.OrdersSort,
) exp.OrderedExpression {
	var columnName string

	switch sort.Field {
	case ports.ORDER_CREATED_BY:
		columnName = CREATED_BY
	case ports.ORDER_MODIFIED_BY:
		columnName = MODIFIED_BY
		// case ports.ORDER_NAME:
		// 	columnName = ASSIGN_TO_NAME
		// default:
		// 	columnName = ASSIGN_TO_NAME
	}

	if sort.Order == ports.SORT_ASCENDING {
		return goqu.I(columnName).Asc()
	} else {
		return goqu.I(columnName).Desc()
	}

}

func (r *Repository) Count(
	search ports.OrdersSearch,
	filters ports.OrdersFilters,
	account_id int,
) (int64, error) {

	w := []exp.Expression{}
	temp := false

	fmt.Println(filters.User, "status")
	if filters.Status != nil {
		if *filters.Status == "To_be_Confirmed" {
			temp = true
		}
	}
	if filters.User == "buyer" {
		w = append(w, goqu.C(ACCOUNT_ID).Eq(account_id))
		fmt.Println(filters.User, w, account_id, "yesssss")

	} else if filters.User == "Supplier" {
		fmt.Println(filters.User, account_id, "ytuutuss")

		w = append(w, goqu.C(SUPPLIER_ID).Eq(account_id))
	} // fmt.Println(*filters.Status,"status");
	if temp || filters.Status == nil {
		fmt.Println("hereee1")
		w = append(w, goqu.C(STATUS).Neq("Confirmed"))
		w = append(w, goqu.C(STATUS).Neq("Completed"))
		w = append(w, goqu.C(STATUS).Neq("Trash"))
	} else {
		fmt.Println("hereee2")
		w = append(w, goqu.C(STATUS).Eq(*filters.Status))
	}
	// fmt.Println("fata", filters.IsConsignee)
	if search.Name != nil {
		searchClauses := r.buildSearchWhereClauses(*search.Name)

		if len(searchClauses) > 1 {
			w = append(w, goqu.Or(searchClauses...))
		} else if len(searchClauses) == 1 {
			w = append(w, searchClauses[0])
		}
	}
	// fmt.Println("fata3")
	count, err := r.goquDB.From(TABLE).Prepared(true).Where(w...).Count()
	// fmt.Println("fata4")
	fmt.Println(count, "count")
	if err != nil {
		// fmt.Println("here y5")
		return 0, err
	}

	return count, nil
}
func (r *Repository) UpdateStatus(
	id string,
	status string,
	reason *string,
) error {
	record := goqu.Record{
		STATUS: status,
	}
	if reason != nil {
		record[REJECTION_REASON] = reason
	}
	fmt.Println(record, status, "record")
	if _, er := r.goquDB.Update(
		TABLE,
	).Prepared(true).Set(
		record,
	).Where(goqu.C(ID).Eq(id)).Executor().Exec(); er != nil {
		fmt.Println(er, "eror")
		return er
	}
	fmt.Println("updated")
	return nil
}
func (r *Repository) FetchAllExpired() ([]rdbms.SalesOrder, error) {
	var pos []rdbms.SalesOrder
	exp := time.Now().AddDate(0, 0, 1).String()
	if er := r.goquDB.From(TABLE).Select(
		ID,
		PO_ID,
		SO_NUMBER,
		STATUS,
		DUE_DATE,
		FLOW_INSTANCE_ID,
		FLOW_INSTANCE_PARAMS_ID,
	).Where(goqu.C(DUE_DATE).Eq(exp)).Executor().ScanStructs(&pos); er != nil {
		return nil, er
	}
	return pos, nil
}

func (r *Repository) buildSearchWhereClauses(search string) []exp.Expression {
	clauses := []exp.Expression{}

	clauses = append(clauses, goqu.I(SO_NUMBER).ILike(fmt.Sprintf("%%%s%%", search)))

	return clauses
}
