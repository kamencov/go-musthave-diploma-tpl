package service

import (
	"context"
	"time"
)

func (s *Service) CheckTableUserLogin(ctx context.Context, login string) error {
	return s.db.CheckTableUserLogin(ctx, login)
}

func (s *Service) CheckTableUserPassword(ctx context.Context, password string) (string, bool) {
	return s.db.CheckTableUserPassword(ctx, password)
}

func (s *Service) CheckWriteOffOfFunds(ctx context.Context, order string, sum float64, now time.Time) error {
	return s.db.CheckWriteOffOfFunds(ctx, order, sum, now)
}
