package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cockroachdb/cockroach-go/crdb"
	"github.com/pkg/errors"

	"github.com/shijuvar/gokit/examples/nats-streaming/pb"
)

// QueryStore syncs data model to be used for query operations
// Because it's store for read model, denormalized data would be inserted

type QueryStore struct{}

func (store QueryStore) SyncOrderQueryModel(order pb.Order) error {

	// Run a transaction to sync the query model.
	// Node: There is an issue with this at this moment
	err := crdb.ExecuteTx(context.Background(), db, nil, func(tx *sql.Tx) error {
		return createOrderQueryModel(tx, order)
	})
	if err != nil {
		return errors.Wrap(err, "Error on syncing query store")
	}
	return nil
}

func createOrderQueryModel(tx *sql.Tx, order pb.Order) error {

	// Insert order into the "orders" table.
	sql := fmt.Sprintf("INSERT INTO orders (id, customerid, status, createdon, restaurantid) VALUES ('%s','%s','%s',%s,'%s')", order.OrderId, order.CustomerId, order.Status, order.CreatedOn, order.RestaurantId)
	_, err := tx.Exec(sql)
	if err != nil {
		return errors.Wrap(err, "Error on insert into orders")
	}
	// Insert order items into the "orderitems" table.
	// Because it's store for read model, we can insert denormalized data
	for _, v := range order.OrderItems {
		sql = fmt.Sprintf("INSERT INTO orderitems (orderid, customerid, code, name, unitprice, quantity) VALUES ('%s','%s','%s','%s',%s, %s)", order.OrderId, order.CustomerId, v.Code, v.Name, v.UnitPrice, v.Quantity)
		_, err := tx.Exec(sql)
		if err != nil {
			return errors.Wrap(err, "Error on insert into order items")
		}
	}
	return nil
}
