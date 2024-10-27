package lineitemsrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
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
	data []rdbms.OrderLineItems,
) error {
	query := []goqu.Record{}
	for _, v := range data {
		query = append(query, goqu.Record{
			ID:            v.Id,
			PO_ID:         v.PoId,
			SO_ID:         v.SoId,
			LI_REF_ID:     v.LiRefId,
			ITEM_CODE:     v.ItemCode,
			DESCRIPTION:   v.Description,
			BATCH_COUNT:   v.BatchCount,
			QUANTITY:      v.Quantity,
			OG:            v.OG,
			DELVIERY_DATE: v.DeliveryDate,
		})
	}
	result := r.goquDB.Insert(
		TABLE,
	).Prepared(true).Rows(
		query,
	)
	if _, er := result.Executor().Exec(); er != nil {
		return er
	}
	fmt.Println("inserted")
	return nil
}

func (r *Repository) GetPOLineItems(
	id string,
	soID *string,
	og *bool,
) ([]rdbms.OrderLineItems, error) {
	var lineItems []rdbms.OrderLineItems
	exp := []exp.Expression{}
	notNullStr := "NOT_NULL"
	if soID != nil && soID == &notNullStr {
		fmt.Println("here1")
		exp = append(exp, goqu.C(SO_ID).IsNotNull())
	} else {
		if soID != nil {
			fmt.Println("here2")
			exp = append(exp, goqu.C(SO_ID).Eq(*soID))
		} else if soID == nil && !*og {
			fmt.Println("here3")
			exp = append(exp, goqu.C(SO_ID).IsNull())
		}
	}
	if og != nil && *og {
		fmt.Println("here4")
		exp = append(exp, goqu.C(OG).Eq(true))
	}
	exp = append(exp, goqu.C(PO_ID).Eq(id))
	fmt.Println("ogsssss", exp)
	query := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ID,
		PO_ID,
		SO_ID,
		LI_REF_ID,
		ITEM_CODE,
		DESCRIPTION,
		BATCH_COUNT,
		QUANTITY,
		OG,
		DELVIERY_DATE,
		CREATED_ON,
	).Where(exp...)
	err := query.ScanStructs(&lineItems)
	if err != nil {
		return []rdbms.OrderLineItems{}, err
	}
	return lineItems, nil
}

func (r *Repository) DeletePOLineItems(id string, by string) error {
	exp := []exp.Expression{}
	switch by {
	case "IS_OG":
		exp = append(exp, goqu.C(OG).Eq(true))
	}
	exp = append(exp, goqu.C(PO_ID).Eq(id))
	query := r.goquDB.From(
		TABLE,
	).Prepared(true).Delete().Where(exp...)
	_, err := query.Executor().Exec()
	if err != nil {
		return err
	}
	fmt.Println("deleted")
	return nil
}

func (r *Repository) FetchLastInsertedLIForSOByRefId(
	id string,
) (rdbms.OrderLineItems, error) {
	var lineItem rdbms.OrderLineItems
	query := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ID,
		PO_ID,
		SO_ID,
		LI_REF_ID,
		ITEM_CODE,
		DESCRIPTION,
		BATCH_COUNT,
		QUANTITY,
		OG,
		DELVIERY_DATE,
		CREATED_ON,
	).Where(goqu.C(LI_REF_ID).Eq(id)).Order(goqu.C(CREATED_ON).Desc()).Limit(1)
	if _, err := query.Executor().ScanStruct(&lineItem); err != nil {
		return rdbms.OrderLineItems{}, err
	}
	return lineItem, nil
}

func (r *Repository) GetLIBySOid(
	id string,
) ([]rdbms.OrderLineItems, error) {
	var lineItems []rdbms.OrderLineItems
	fmt.Println("sosfsdfss")

	query := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ID,
		PO_ID,
		SO_ID,
		LI_REF_ID,
		ITEM_CODE,
		DESCRIPTION,
		BATCH_COUNT,
		QUANTITY,
		OG,
		DELVIERY_DATE,
		CREATED_ON,
	).Where(goqu.C(SO_ID).Eq(id))
	fmt.Println("sosfsdfsafdss")

	if err := query.Executor().ScanStructs(&lineItems); err != nil {
		return []rdbms.OrderLineItems{}, err
	}
	return lineItems, nil
}

// id is line item id
func (r *Repository) DeleteLineItemById(
	id string,
) error {
	query := r.goquDB.From(TABLE).Delete().Where(goqu.C(ID).Eq(id))
	if _, err := query.Executor().Exec(); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("deleted")
	return nil
}

func (r *Repository) UpdateLineItemById(
	id string,
	bq *int, // BATCH COUNT
	quan *int, // QUANTITY
	itemCode *string,
	description *string,
	deliveryDate *string,
) error {
	record := goqu.Record{}
	if itemCode != nil {
		record[ITEM_CODE] = *itemCode
	}
	if description != nil {
		record[DESCRIPTION] = *description
	}
	if deliveryDate != nil {
		record[DELVIERY_DATE] = *deliveryDate
	}
	if quan != nil {
		record[QUANTITY] = *quan
	}
	if bq != nil {
		record[BATCH_COUNT] = *bq
	}
	query := r.goquDB.Update(
		TABLE,
	).Prepared(true).Set(
		record,
	).Where(goqu.C(ID).Eq(id))
	if _, err := query.Executor().Exec(); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("updated")
	return nil
}

func (r *Repository) GetLineItemsBySOId(
	id string,
) ([]rdbms.OrderLineItems, error) {
	var lineItems []rdbms.OrderLineItems
	if er := r.goquDB.From(
		TABLE,
	).Select(
		ID,
		PO_ID,
		SO_ID,
		LI_REF_ID,
		ITEM_CODE,
		DESCRIPTION,
		BATCH_COUNT,
		QUANTITY,
		OG,
		DELVIERY_DATE,
		CREATED_ON,
	).Where(goqu.C(SO_ID).Eq(id)).ScanStructs(&lineItems); er != nil {
		return []rdbms.OrderLineItems{}, er
	}
	return lineItems, nil
}

func (r *Repository) ExpiredLineItems() ([]rdbms.OrderLineItems, error) {
	var li []rdbms.OrderLineItems
	today := time.Now().Format("2006-01-02")
	if er := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ID,
		PO_ID,
		SO_ID,
		LI_REF_ID,
		DELVIERY_DATE,
	).Where(goqu.C(DELVIERY_DATE).Lt(today)).ScanStructs(&li); er != nil {
		return []rdbms.OrderLineItems{}, er
	}
	fmt.Println("expired line items", li)
	return li, nil
}
