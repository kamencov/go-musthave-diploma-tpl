package db

import (
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/models"
	"log/slog"
)

type Order struct {
	OrderID string
	Status  string
}

func (d *DateBase) GetAllOrders() ([]Order, error) {
	var orders []Order

	query := "SELECT order_id, order_status FROM loyalty WHERE order_status IN ('REGISTERED', 'PROCESSING', 'NEW')" //берем только в статусе REGISTERED и PROCESSING и NEW
	rows, err := d.Gets(query)
	if err != nil {
		slog.Error("Error :", "gets rows in worker - ", customerrors.ErrNotData)
		return nil, err
	}
	if rows.Err() != nil {
		slog.Error("Error :", "gets rows in worker - ", customerrors.ErrNotData)
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			slog.Error("Error ", "closing row set:", err)
		}
	}()

	for rows.Next() {
		var (
			order       string
			orderStatus string
		)

		if err = rows.Scan(&order, &orderStatus); err != nil {
			slog.Error("Error worker", "Error scanning rows in worker :", err)
			continue
		}

		orders = append(orders, Order{
			OrderID: order,
			Status:  orderStatus,
		})
	}

	return orders, nil
}

func (d *DateBase) SaveNewStatusAndBonus(accrual models.ResponseAccrual) error {
	querySave := "UPDATE loyalty SET order_status = $1, bonus = $2 WHERE order_id = $3"
	err := d.Save(querySave, accrual.Status, accrual.Accrual, accrual.Order)
	if err != nil {
		slog.Error("Error: ", "saving data in worker: ", err)
		return err
	}
	return nil
}
