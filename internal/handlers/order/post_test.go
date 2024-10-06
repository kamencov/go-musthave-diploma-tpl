package order

import (
	"bytes"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/logger"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/middleware"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/models"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/service/orders"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerPost(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		login          string
		loginID        int
		loyalty        models.Loyalty
		loginErr       error
		loyaltyErr     error
		responseError  error
		expectedStatus int
	}{
		{
			name:    "Successful_post",
			body:    "22664155",
			login:   "test",
			loginID: 1,
			loyalty: models.Loyalty{
				UserID: 1,
			},
			loyaltyErr:     customerrors.ErrNoOrderInLoyalty,
			expectedStatus: http.StatusAccepted,
		},
		{
			name:           "User_not_authenticated",
			loginID:        0,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalid_order_numbers",
			body:           "5",
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:    "Order_another_user",
			body:    "22664155",
			login:   "test",
			loginID: 1,
			loyalty: models.Loyalty{
				UserID: 2,
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name:    "Order_this_user",
			body:    "22664155",
			login:   "test",
			loginID: 1,
			loyalty: models.Loyalty{
				UserID: 1,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Error_status_500",
			body:           "22664155",
			login:          "test",
			loginID:        1,
			loginErr:       errors.New("cannot loading order"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loger := logger.NewLogger()
			ctrl := gomock.NewController(t)

			repo := orders.NewMockStorage(ctrl)
			repo.EXPECT().GetLoginID(gomock.Any()).Return(tt.loginID, tt.loginErr).AnyTimes()
			repo.EXPECT().GetLoyalty(gomock.Any()).Return(tt.loyalty, tt.loyaltyErr).AnyTimes()
			repo.EXPECT().SaveOrder(tt.loginID, tt.body, gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			serv := orders.NewService(repo, loger)

			req, err := http.NewRequest("POST", "/", io.NopCloser(bytes.NewBufferString(tt.body)))
			if err != nil {
				t.Fatal(err)
			}

			req = req.WithContext(context.WithValue(req.Context(), middleware.LoginContentKey, tt.login))

			handler := NewHandler(serv, loger)

			w := httptest.NewRecorder()

			handler.Post(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected expectedStatus code %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
