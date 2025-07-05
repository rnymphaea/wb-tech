package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"

	"wb-tech-l0/internal/database/models"
)

func (pgrepo *PostgresRepo) GetOrderByUID(ctx context.Context, uid string) (*models.Order, error) {
	orderQuery := `SELECT 
			o.uid, o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id, 
			o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard,
			d.name, d.phone, d.zip, d.city, d.address, d.region, d.email,
			p.transaction, p.request_id, p.currency, p.provider, p.amount, p.payment_dt, p.bank, 
			p.delivery_cost, p.goods_total, p.custom_fee
		FROM orders o
		JOIN deliveries d ON o.uid = d.order_uid
		JOIN payments p ON o.uid = p.order_uid
		WHERE o.uid = $1
	`

	order := models.Order{}
	err := pgrepo.pool.QueryRow(ctx, orderQuery, uid).Scan(
		&order.UID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, 
		&order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, 
		&order.OofShard,
		// Delivery
		&order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City, 
		&order.Delivery.Address, &order.Delivery.Region, &order.Delivery.Email,
		// Payment
		&order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency, 
		&order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDt, &order.Payment.Bank, 
		&order.Payment.DeliveryCost, &order.Payment.GoodsTotal, &order.Payment.CustomFee,
	)
	if err != nil {
		if err == pgx.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}

	itemsQuery := `SELECT 
			chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
		FROM items
		WHERE order_uid = $1
	`

	rows, err := pgrepo.pool.Query(ctx, itemsQuery, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	order.Items = make([]models.Item, 0)
	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, 
			&item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status,
		)
		if err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &order, nil
}

func (pgrepo *PostgresRepo) SaveOrder(ctx context.Context, order *models.Order) error {
	tx, err := pgrepo.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := pgrepo.saveOrder(ctx, tx, order); err != nil {
		return err
	}

	if err := pgrepo.saveDelivery(ctx, tx, order); err != nil {
		return err
	}

	if err := pgrepo.savePayment(ctx, tx, order); err != nil {
		return err
	}

	if err := pgrepo.saveItems(ctx, tx, order); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (pgrepo *PostgresRepo) saveOrder(ctx context.Context, tx pgx.Tx, order *models.Order) error {
	query := `
		INSERT INTO orders (
			uid, track_number, entry, locale, internal_signature, customer_id, 
			delivery_service, shardkey, sm_id, date_created, oof_shard
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := tx.Exec(ctx, query,
		order.UID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)
	return err
}

func (pgrepo *PostgresRepo) saveDelivery(ctx context.Context, tx pgx.Tx, order *models.Order) error {
	query := `
		INSERT INTO deliveries (
			order_uid, name, phone, zip, city, address, region, email
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := tx.Exec(ctx, query,
		order.UID,
		order.Delivery.Name,
		order.Delivery.Phone,
		order.Delivery.Zip,
		order.Delivery.City,
		order.Delivery.Address,
		order.Delivery.Region,
		order.Delivery.Email,
	)
	return err
}

func (pgrepo *PostgresRepo) savePayment(ctx context.Context, tx pgx.Tx, order *models.Order) error {
	query := `
		INSERT INTO payments (
			order_uid, transaction, request_id, currency, provider, amount, 
			payment_dt, bank, delivery_cost, goods_total, custom_fee
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := tx.Exec(ctx, query,
		order.UID,
		order.Payment.Transaction,
		order.Payment.RequestID,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.PaymentDt,
		order.Payment.Bank,
		order.Payment.DeliveryCost,
		order.Payment.GoodsTotal,
		order.Payment.CustomFee,
	)
	return err
}

func (pgrepo *PostgresRepo) saveItems(ctx context.Context, tx pgx.Tx, order *models.Order) error {
	query := `
		INSERT INTO items (
			order_uid, chrt_id, track_number, price, rid, name, sale, size, 
			total_price, nm_id, brand, status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	batch := &pgx.Batch{}
	for _, item := range order.Items {
		batch.Queue(query,
			order.UID,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status,
		)
	}

	br := tx.SendBatch(ctx, batch)
	defer br.Close()

	for i := 0; i < batch.Len(); i++ {
		if _, err := br.Exec(); err != nil {
			return err
		}
	}
	return nil
}
