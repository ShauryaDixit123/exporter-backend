package purchaseorderrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"fmt"

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

func (r *Repository) InsertOne(
	data rdbms.PurchaseOrder,
) error {
	result := r.goquDB.Insert(
		TABLE,
	).Prepared(true).Rows(
		goqu.Record{
			ID:                   data.Id,
			ACCOUNT_ID:           data.AccountId,
			PO_NUMBER:            data.PONumber,
			DUE_DATE:             data.DueDate,
			SHIPMENT_TERMS:       data.ShipmentMode,
			STATUS:               data.Status,
			TERMS_AND_CONDITIONS: data.TermsAndConditions,
			REMARKS:              data.Remarks,
			REJECTION_REASON:     data.RejectionReason,
			SUPPLIER_ID:          data.SupplierId,
			SHIPMENT_MODE:        data.ShipmentMode,
			CREATED_BY:           data.CreatedBy,
		},
	)
	_, err := result.Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectOne(id string) (rdbms.PurchaseOrder, error) {
	var data rdbms.PurchaseOrder
	_, err := r.goquDB.From(TABLE).Select(
		ID,
		USER_ID,
		ACCOUNT_ID,
		FLOW_INSTANCE_ID,
		FLOW_INSTANCE_PARAMS_ID,
		PO_NUMBER,
		PICKUP_LOCATION_ID,
		DROP_LOCATION_ID,
		SHIPMENT_TERMS,
		TERMS_AND_CONDITIONS,
		REMARKS,
		SUPPLIER_ID,
		STATUS,
		REJECTION_REASON,
		CREATED_BY,
	).Where(goqu.C(ID).Eq(id)).ScanStruct(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *Repository) UpdateOne(
	id string,
	data rdbms.PurchaseOrder,
) error {
	record := goqu.Record{
		USER_ID:                 data.UserId,
		PO_NUMBER:               data.PONumber,
		DUE_DATE:                data.DueDate,
		FLOW_INSTANCE_ID:        data.FlowInstanceId,
		FLOW_INSTANCE_PARAMS_ID: data.FlowInstanceParamsId,
		SHIPMENT_TERMS:          data.ShipmentMode,
		STATUS:                  data.Status,
		TERMS_AND_CONDITIONS:    data.TermsAndConditions,
		REMARKS:                 data.Remarks,
		REJECTION_REASON:        data.RejectionReason,
		SUPPLIER_ID:             data.SupplierId,
		SHIPMENT_MODE:           data.ShipmentMode,
	}
	_, err := r.goquDB.Update(TABLE).Prepared(true).Set(
		record,
	).Where(
		goqu.C(ID).Eq(id),
	).Executor().Exec()
	if err != nil {
		return err
	}
	fmt.Println("updated!")
	return nil
}

func (r *Repository) UpdateStatus(
	id string,
	status string,
	rejectionReason *string,
) error {
	record := goqu.Record{
		STATUS: status,
	}
	if rejectionReason != nil {
		record[REJECTION_REASON] = *rejectionReason
	}
	_, er := r.goquDB.Update(TABLE).Prepared(true).Set(record).Where(goqu.C(ID).Eq(id)).Executor().Exec()
	if er != nil {
		return er
	}
	return nil
}

func (r *Repository) DeleteOne(id string) error {
	_, er := r.goquDB.Delete(TABLE).Prepared(true).Where(goqu.C(ID).Eq(id)).Executor().Exec()
	if er != nil {
		return er
	}
	return nil
}
func (r *Repository) CheckPOExistsByCode(code string) (bool, error) {
	var po rdbms.PurchaseOrder
	a, err := r.goquDB.From(TABLE).Select(PO_NUMBER).Where(goqu.C(PO_NUMBER).Eq(code)).ScanStruct(&po)
	fmt.Println(a, po, err, "check po exists by code")
	if err != nil {
		return true, err
	}
	return a, nil
}

func (r *Repository) SelectMany(
	limit uint,
	offset uint,
	search ports.OrdersSearch,
	sort ports.OrdersSort,
	filters ports.OrdersFilters,
	account_id int,
) ([]rdbms.PurchaseOrder, error) {
	var orders []rdbms.PurchaseOrder

	w := []exp.Expression{}
	temp := false
	// fmt.Println(*filters.Status,"status")
	if filters.Status != nil {
		if *filters.Status == "To_be_Confirmed" {
			temp = true
		}
	}
	// fmt.Println(filters.IsConsignee, "filters.IsConsignee")

	if filters.User == "buyer" {
		w = append(w, goqu.C(ACCOUNT_ID).Eq(account_id))
		fmt.Println(filters.User, w, account_id, "yesssss")

	} else if filters.User == "Supplier" {
		// fmt.Println(filters.User, account_id, "ytuutuss")

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

	fmt.Println(search.Name, "search name")
	if search.Name != nil {
		searchClauses := r.buildSearchWhereClauses(*search.Name)
		fmt.Println(searchClauses, "searchclause")

		if len(searchClauses) > 1 {
			w = append(w, goqu.Or(searchClauses...))
		} else if len(searchClauses) == 1 {
			w = append(w, searchClauses[0])
		}
	}

	query := r.goquDB.From(TABLE).Prepared(true).Select(
		ID,
		USER_ID,
		ACCOUNT_ID,
		PO_NUMBER,
		DUE_DATE,
		SHIPMENT_TERMS,
		TERMS_AND_CONDITIONS,
		REMARKS,
		REJECTION_REASON,
		STATUS,
		SUPPLIER_ID,
		SHIPMENT_MODE,
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
			"SelectMany purchase order failed",
			"there was an issue while fetching purchase orders",
			err,
			map[string]any{},
			map[string]any{
				"query":  querySQL,
				"params": params,
				"error":  qerr,
			},
		)

		return []rdbms.PurchaseOrder{}, err
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
	default:
		columnName = CREATED_BY
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
	if filters.User == "buyer" {
		w = append(w, goqu.C(ACCOUNT_ID).Eq(account_id))
		fmt.Println(filters.User, w, account_id, "yesssss")

	} else if filters.User == "Supplier" {
		fmt.Println(filters.User, account_id, "ytuutuss")

		w = append(w, goqu.C(SUPPLIER_ID).Eq(account_id))
	}
	temp := false

	// fmt.Println("here")
	if filters.Status != nil {
		if *filters.Status == "To_be_Confirmed" {
			temp = true
		}
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
	fmt.Println(count, "counting in repo")
	if err != nil {
		// fmt.Println("here y5")
		return 0, err
	}

	return count, nil
}

func (r *Repository) buildSearchWhereClauses(search string) []exp.Expression {
	clauses := []exp.Expression{}

	clauses = append(clauses, goqu.I(PO_NUMBER).ILike(fmt.Sprintf("%%%s%%", search)))

	// Add more search conditions as needed

	return clauses
}
