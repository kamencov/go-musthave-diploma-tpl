package orders

import (
	"github.com/golang/mock/gomock"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/logger"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/models"
	"testing"
	"time"
)

func TestMockStorageGetLoginByAccessToken(t *testing.T) {
	tests := []struct {
		name        string
		login       string
		loginID     int
		loginErr    error
		order       string
		now         time.Time
		loyalty     models.Loyalty
		loyaltyErr  error
		expentedErr error
	}{
		{
			name:    "Successful_test",
			login:   "test",
			loginID: 1,
			order:   "2222",
			now:     time.Now(),
			loyalty: models.Loyalty{
				UserID: 1,
			},
			loyaltyErr: customerrors.ErrNoOrderInLoyalty,
		},
		{
			name:        "Bad_login",
			login:       "test",
			loginID:     1,
			loginErr:    customerrors.ErrNotFound,
			order:       "2222",
			now:         time.Now(),
			expentedErr: customerrors.ErrNotFound,
		},
		{
			name:    "Order_by_another_user",
			login:   "test",
			loginID: 1,
			order:   "2222",
			now:     time.Now(),
			loyalty: models.Loyalty{
				UserID: 2,
			},
			expentedErr: customerrors.ErrAnotherUsersOrder,
		},
		{
			name:    "Order_already_registered",
			login:   "test",
			loginID: 1,
			order:   "2222",
			now:     time.Now(),
			loyalty: models.Loyalty{
				UserID: 1,
			},
			expentedErr: customerrors.ErrOrderRegistered,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			loger := logger.NewLogger()
			repo := NewMockStorage(ctrl)
			repo.EXPECT().GetLoginID(tt.login).Return(tt.loginID, tt.loginErr).AnyTimes()
			repo.EXPECT().GetLoyalty(tt.order).Return(tt.loyalty, tt.loyaltyErr).AnyTimes()
			repo.EXPECT().SaveOrder(tt.loginID, tt.order, gomock.Any(), tt.now).Return(nil).AnyTimes()

			service := NewService(repo, loger)

			err := service.GetLoginByAccessToken(tt.order, tt.login, tt.now)
			if tt.expentedErr != err {
				t.Errorf("expected behavior %t, got %t", tt.expentedErr, err)
			}
		})
	}
}

func TestServiceCheckWriteOffOfFunds(t *testing.T) {
	tests := []struct {
		name        string
		login       string
		loginID     int
		loginErr    error
		order       string
		sumBody     float32
		sumBonus    float32
		sumErr      error
		now         time.Time
		expentedErr error
	}{
		{
			name:     "Successful_check_bonus",
			login:    "test",
			loginID:  1,
			order:    "2222",
			sumBody:  10,
			sumBonus: 15,
			now:      time.Now(),
		},
		{
			name:        "Bad_login",
			login:       "test",
			loginID:     1,
			loginErr:    customerrors.ErrNotFound,
			order:       "2222",
			sumBody:     10,
			sumBonus:    15,
			now:         time.Now(),
			expentedErr: customerrors.ErrNotFound,
		},
		{
			name:        "No_bonus_in_data",
			login:       "test",
			loginID:     1,
			order:       "2222",
			sumBody:     10,
			sumBonus:    15,
			sumErr:      customerrors.ErrNotData,
			now:         time.Now(),
			expentedErr: customerrors.ErrNotData,
		},
		{
			name:        "Not_enough_bonuses",
			login:       "test",
			loginID:     1,
			order:       "2222",
			sumBody:     20,
			sumBonus:    15,
			now:         time.Now(),
			expentedErr: customerrors.ErrNotEnoughBonuses,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			loger := logger.NewLogger()
			repo := NewMockStorage(ctrl)
			repo.EXPECT().GetLoginID(tt.login).Return(tt.loginID, tt.loginErr).AnyTimes()
			repo.EXPECT().GetSumBonus(tt.loginID).Return(tt.sumBonus, tt.sumErr).AnyTimes()
			repo.EXPECT().SaveOrder(tt.loginID, tt.order, gomock.Any(), tt.now).Return(nil).AnyTimes()
			repo.EXPECT().UpdateOrder(tt.order, tt.sumBody, tt.now).Return(nil).AnyTimes()

			service := NewService(repo, loger)

			err := service.CheckWriteOffOfFunds(tt.login, tt.order, tt.sumBody, tt.now)

			if tt.expentedErr != err {
				t.Errorf("expented behavior %t, got %t", tt.expentedErr, err)
			}
		})
	}
}
