package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/logger"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/models"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/service"
	"net/http"
	"strconv"
	"time"
)

type WorkerAccrual struct {
	storage *service.Service
	log     *logger.Logger
}

func NewWorkerAccrual(storage *service.Service, log *logger.Logger) *WorkerAccrual {
	return &WorkerAccrual{
		storage: storage,
		log:     log,
	}
}

func (w *WorkerAccrual) StartWorkerAccrual(ctx context.Context, addressAccrual string) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			go w.getAccrual(ctx, addressAccrual)
		case <-ctx.Done():
			return
		}
	}
}

func (w *WorkerAccrual) getAccrual(ctx context.Context, addressAccrual string) {
	query := "SELECT order_id, order_status FROM loyalty WHERE order_status IN ('REGISTERED', 'PROCESSING', 'NEW')" //берем только в статусе REGISTERED и PROCESSING и NEW
	rows, err := w.storage.Gets(query)
	if err != nil {
		w.log.Error("Error:", customerrors.ErrNotData)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			w.log.Error("Error closing rowset:", err)
		}
	}()

	for rows.Next() {
		var (
			order       string
			orderStatus string
		)

		var accrual models.ResponseAccrual

		if err = rows.Scan(&order, &orderStatus); err != nil {
			w.log.Error("Error scanning rows:", err)
			continue
		}

		req, err := http.Get(fmt.Sprintf("%s/api/orders/%s", addressAccrual, order))
		if err != nil {
			w.log.Error("Error making request:", err)
			continue
		}
		defer req.Body.Close()

		if err = json.NewDecoder(req.Body).Decode(&accrual); err != nil {
			w.log.Error("Error decoding response:", err)
			continue
		}

		if req.StatusCode == http.StatusTooManyRequests {
			timeSleep, err := strconv.Atoi(req.Header.Get("Retry-After"))
			if err != nil {
				time.Sleep(60 * time.Second)
			} else {
				time.Sleep(time.Duration(timeSleep) * time.Second)
			}
		}

		if orderStatus != accrual.Status {
			querySave := "UPDATE loyalty SET order_status = $1, bonus = $2 WHERE order_id = $3"
			err = w.storage.Save(querySave, accrual.Status, accrual.Accrual, accrual.Order)
			if err != nil {
				w.log.Error("Error saving data:", err)
			}
		}
	}
}