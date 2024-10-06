package withdraw

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/logger"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/middleware"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/service/orders"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerPost(t *testing.T) {
	tests := []struct {
		name           string
		login          string
		requestBody    RequestBody
		whenBodyBad    bool
		sumBonus       float32
		sumBonusError  error
		responseError  error
		expectedStatus int
	}{
		{
			name:  "Successful_post",
			login: "test",
			requestBody: RequestBody{
				Order: "22664155",
				Sum:   55.5,
			},
			sumBonus:       55.5,
			expectedStatus: http.StatusOK,
		},
		{
			name:  "Invalid_body",
			login: "test",
			requestBody: RequestBody{
				Order: "22664155",
				Sum:   55.5,
			},
			whenBodyBad:    true,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:  "Invalid_order",
			login: "test",
			requestBody: RequestBody{
				Order: "226641551",
				Sum:   55.5,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:  "Invalid_login",
			login: "",
			requestBody: RequestBody{
				Order: "22664155",
				Sum:   55.5,
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:  "Order_is_not_in_the_database",
			login: "test",
			requestBody: RequestBody{
				Order: "22664155",
				Sum:   55.5,
			},
			sumBonusError:  customerrors.ErrNotData,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:  "Don't_have_enough_bonuses",
			login: "test",
			requestBody: RequestBody{
				Order: "22664155",
				Sum:   55.5,
			},
			sumBonus:       53.5,
			expectedStatus: http.StatusPaymentRequired,
		},
		{
			name:  "Check_status_500",
			login: "test",
			requestBody: RequestBody{
				Order: "22664155",
				Sum:   55.5,
			},
			responseError:  errors.New("check status 500"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			loger := logger.NewLogger()

			repo := orders.NewMockStorage(ctrl)
			repo.EXPECT().GetLoginID(gomock.Any()).Return(1, tt.responseError).AnyTimes()
			repo.EXPECT().GetSumBonus(gomock.Any()).Return(tt.sumBonus, tt.sumBonusError).AnyTimes()
			repo.EXPECT().SaveOrder(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			repo.EXPECT().UpdateOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			//repo.EXPECT().CheckWriteOffOfFunds(tt.login, tt.requestBody.Order, tt.requestBody.Sum, gomock.Any()).Return(tt.responseError).AnyTimes()

			serv := orders.NewService(repo, loger)
			handler := NewHandler(serv, loger)

			req, err := http.NewRequest("POST", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			req = req.WithContext(context.WithValue(req.Context(), middleware.LoginContentKey, tt.login))

			if tt.whenBodyBad {
				req.Body = io.NopCloser(bytes.NewBufferString("incorrect body"))
			} else {
				req.Body = jsonReader(tt.requestBody)
			}

			w := httptest.NewRecorder()

			handler.Post(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status code %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func jsonReader(v interface{}) io.ReadCloser {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return io.NopCloser(bytes.NewBuffer(b))
}
