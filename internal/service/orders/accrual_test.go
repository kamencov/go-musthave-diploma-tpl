package orders

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/logger"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/storage/db"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceGetAccrual(t *testing.T) {

	tests := []struct {
		name               string
		orders             []db.Order
		dbErr              error
		httpResponseStatus int
		httpResponseBody   string
		httpError          error
		saveErr            error
		expectedStatus     int
	}{
		{
			name: "Success",
			orders: []db.Order{{
				OrderID: "1",
				Status:  "NEW",
			}},
			httpResponseStatus: http.StatusOK,
			httpResponseBody: `{
                "order": "1",
                "status": "PROCESSED",
                "accrual": 500
            }`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := NewMockStorage(ctrl)
			repo.EXPECT().GetAllOrders().Return(tt.orders, tt.dbErr).AnyTimes()
			repo.EXPECT().SaveNewStatusAndBonus(gomock.Any()).Return(tt.saveErr).AnyTimes()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.httpResponseStatus)
				fmt.Fprint(w, tt.httpResponseBody)
			}))
			defer server.Close()

			addressAccrual := server.URL

			loger := logger.NewLogger()

			service := NewService(repo, loger)

			service.GetAccrual(addressAccrual)
		})
	}
}
