package orders

import (
	"errors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/models"
	"time"
)

func (s *Service) GetLoginByAccessToken(order string, login string, now time.Time) error {

	user, err := s.db.GetLoginID(login)
	if err != nil {
		return customerrors.ErrNotFound
	}
	loyalty, err := s.db.GetLoyalty(order)
	if errors.Is(err, customerrors.ErrNoOrderInLoyalty) {
		err = s.db.SaveOrder(user, order, models.NewOrder, now)
		if err != nil {
			return err
		}
		return nil
	}
	if loyalty.UserID != user {
		return customerrors.ErrAnotherUsersOrder
	}

	return customerrors.ErrOrderRegistered
}

func (s *Service) GetAllUserOrders(login string) ([]*models.OrdersUser, error) {
	return s.db.GetAllUserOrders(login)
}

func (s *Service) GetBalanceUser(login string) (*models.Balance, error) {
	return s.db.GetBalanceUser(login)
}

func (s *Service) GetWithdrawals(login string) ([]*models.Withdrawals, error) {
	return s.db.GetWithdrawals(login)
}

func (s *Service) CheckWriteOffOfFunds(login, order string, sum float32, now time.Time) error {
	return s.db.CheckWriteOffOfFunds(login, order, sum, now)
}
