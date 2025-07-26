package database

import (
	"database/sql"
	"strconv"

	"github.com/juliocesarboaroli/goexpert-clean-arch-challenge/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindAll() []entity.Order {
	var orders []entity.Order
	rows, err := r.Db.Query(`SELECT * FROM orders`)
	if err != nil {
		panic(err)
	}

	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			panic(err)
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		price, _ := strconv.ParseFloat(result[1], 64)
		tax, _ := strconv.ParseFloat(result[2], 64)
		finalPrice, _ := strconv.ParseFloat(result[3], 64)
		orders = append(orders, entity.Order{
			ID:         result[0],
			Price:      price,
			Tax:        tax,
			FinalPrice: finalPrice,
		})
	}

	return orders
}
