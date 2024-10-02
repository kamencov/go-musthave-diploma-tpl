package orders

import (
	"encoding/json"
	"fmt"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/models"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func (s *Service) GetAccrual(addressAccrual string) {
	orders, err := s.db.GetAllOrders()
	if err != nil {
		slog.Error("Error :", "gets rows in worker - ", customerrors.ErrNotData)
		return
	}
	for _, order := range orders {

		var accrual models.ResponseAccrual

		req, err := http.Get(fmt.Sprintf("%s/api/orders/%s", addressAccrual, order))
		if err != nil {
			slog.Error("Error worker", "Error making request in worker :", err)
			continue
		}
		defer req.Body.Close()

		if err = json.NewDecoder(req.Body).Decode(&accrual); err != nil {
			slog.Error("Error worker", "Error decoding response in worker:", err)
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

		if order.Status != accrual.Status {
			err = s.db.SaveNewStatusAndBonus(accrual)
			if err != nil {
				slog.Error("Error: ", "saving data in worker: ", err)
				continue
			}
		}
	}
}
