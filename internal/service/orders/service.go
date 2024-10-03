package orders

import (
	"github.com/kamencov/go-musthave-diploma-tpl/internal/logger"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/models"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/storage/db"
	"time"
)

//go:generate mockgen -source=./service.go -destination=service_mock.go -package=orders
type Storage interface {
	SaveNewStatusAndBonus(accrual models.ResponseAccrual) error
	SaveOrder(userID int, orderID string, orderStatus string, now time.Time) error
	GetLoyalty(order string) (models.Loyalty, error)
	GetLoginID(login string) (int, error)
	GetAllUserOrders(login string) ([]*models.OrdersUser, error)
	GetAllOrders() ([]db.Order, error)
	GetSumBonus(userID int) (float32, error)
	GetBalanceUser(login string) (*models.Balance, error)
	GetWithdrawals(login string) ([]*models.Withdrawals, error)
	UpdateOrder(order string, sum float32, now time.Time) error
}

type Service struct {
	db   Storage
	logs *logger.Logger
}

func NewService(db Storage, logger *logger.Logger) *Service {
	return &Service{
		db,
		logger,
	}
}
