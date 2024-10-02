package order

import (
	"bytes"
	"context"
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
		body           string
		login          string
		responseError  error
		expectedStatus int
	}{
		{
			name:           "Successful_post",
			body:           "22664155",
			login:          "test",
			expectedStatus: http.StatusAccepted,
		},
		{
			name:           "User_not_authenticated",
			login:          "",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalid_order_numbers",
			body:           "5",
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "Order_another_user",
			body:           "22664155",
			login:          "test",
			responseError:  customerrors.ErrAnotherUsersOrder,
			expectedStatus: http.StatusConflict,
		},
		{
			name:           "Order_another_user",
			body:           "22664155",
			login:          "test",
			responseError:  customerrors.ErrOrderRegistered,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Order_another_user",
			body:           "22664155",
			login:          "test",
			responseError:  errors.New("cannot loading order"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loger := logger.NewLogger()
			ctrl := gomock.NewController(t)

			repo := orders.NewMockStorage(ctrl)
			repo.EXPECT().GetUserByAccessToken(tt.body, tt.login, gomock.Any()).Return(tt.responseError).AnyTimes()

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
