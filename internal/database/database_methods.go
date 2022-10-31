package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"l0/internal/models"
	"log"
)

// const values for connection to database
const (
	username = "postgres"
	password = ""
	host     = "localhost"
	port     = "5432"
	database = "l0"
)

// closeConnection - close connection to database
func closeConnection(link *pgx.Conn) {
	err := link.Close(context.Background())
	if err != nil {
		log.Printf("Connection closing fail: %v\n", err)
	}
}

// connection - connect to database
func connection() *pgx.Conn {
	connectionUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
	conn, err := pgx.Connect(context.Background(), connectionUrl)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil
	}
	return conn
}

// AddMessageToDatabase - addition message to four tables of database
func AddMessageToDatabase(order models.Order) error {
	conn := connection()
	defer closeConnection(conn)

	// add data to orders
	_, err := conn.Exec(context.Background(),
		"INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shred) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);",
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService,
		order.Shardkey, order.SmId, order.DateCreated, order.OofShred)
	if err != nil {
		return errors.New(fmt.Sprintf("Orders insertion failed (%v)\n", err))
	}

	// add data to delivery
	_, err = conn.Exec(context.Background(),
		"INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);",
		order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
		order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return errors.New(fmt.Sprintf("Delivery insertion failed (%v)\n", err))
	}

	// add data to payment
	_, err = conn.Exec(context.Background(),
		"INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);",
		order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider,
		order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal,
		order.Payment.CustomFee)
	if err != nil {
		return errors.New(fmt.Sprintf("Payment insertion failed (%v)\n", err))
	}

	// add all items to database
	for i := range order.Items {
		_, err = conn.Exec(context.Background(),
			"INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);",
			order.OrderUID, order.Items[i].ChrtID, order.Items[i].TrackNumber, order.Items[i].Price, order.Items[i].Rid,
			order.Items[i].Name, order.Items[i].Sale, order.Items[i].Size, order.Items[i].TotalPrice, order.Items[i].NmID,
			order.Items[i].Brand, order.Items[i].Status)
		if err != nil {
			return errors.New(fmt.Sprintf("Item %v insertion failed (%v)\n", i, err))
		}
	}

	log.Printf("Got new message!")
	return nil
}

// SyncCacheAndDatabase - synchronize cache and database values
// copy values from database to cache
func SyncCacheAndDatabase() error {
	conn := connection()
	defer closeConnection(conn)

	var countOfRowsInTable int

	err := conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM orders;").Scan(&countOfRowsInTable)
	if err != nil {
		return errors.New(fmt.Sprintf("QueryRow failed (%v)\n", err))
	}

	if len(models.Cache) != countOfRowsInTable {

		// copy orders to cache
		rows, err := conn.Query(context.Background(), "select * from orders;")
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow (orders) failed (%v)\n", err))
		}
		for rows.Next() {
			var order models.Order
			err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry,
				&order.Locale, &order.InternalSignature, &order.CustomerId,
				&order.DeliveryService, &order.Shardkey, &order.SmId,
				&order.DateCreated, &order.OofShred)
			if err != nil {
				return errors.New(fmt.Sprintf("Error in scanning order row (%v)\n", err))
			}
			if _, found := models.Cache[order.OrderUID]; !found {
				models.Cache[order.OrderUID] = order
			}
		}

		// copy delivery to cache
		rows, err = conn.Query(context.Background(), "SELECT * FROM delivery;")
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow (delivery) failed (%v)\n", err))
		}
		for rows.Next() {
			var delivery models.Delivery
			var uid string
			err := rows.Scan(&uid, &delivery.Name, &delivery.Phone,
				&delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region,
				&delivery.Email)
			if err != nil {
				return errors.New(fmt.Sprintf("Error in scanning delivery row (%v)\n", err))
			}
			if value, found := models.Cache[uid]; found {
				value.Delivery = delivery
				models.Cache[value.OrderUID] = value
			}
		}

		// copy payment to cache
		rows, err = conn.Query(context.Background(), "SELECT * FROM payment;")
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow (payment) failed (%v)\n", err))
		}
		for rows.Next() {
			var payment models.Payment
			var uid string
			err := rows.Scan(&uid, &payment.Transaction, &payment.RequestID,
				&payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDT,
				&payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
			if err != nil {
				return errors.New(fmt.Sprintf("Error in scanning payment row (%v)\n", err))
			}
			if value, found := models.Cache[uid]; found {
				value.Payment = payment
				models.Cache[value.OrderUID] = value
			}
		}

		// copy items to cache
		rows, err = conn.Query(context.Background(), "SELECT * FROM items;")
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow (items) failed (%v)\n", err))
		}
		for rows.Next() {
			var item models.Items
			var uid string
			err := rows.Scan(&uid, &item.ChrtID, &item.TrackNumber, &item.Price,
				&item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID,
				&item.Brand, &item.Status)
			if err != nil {
				return errors.New(fmt.Sprintf("Error in scanning item row (%v)\n", err))
			}
			if value, found := models.Cache[uid]; found {
				value.Items = append(value.Items, item)
				models.Cache[value.OrderUID] = value
			}
		}

	}

	log.Printf("Cache and database synchronized.\n")
	return nil
}
