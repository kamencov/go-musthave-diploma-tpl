package customerrors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrNotFound          = errors.New("not found")
	ErrIsTruePassword    = errors.New("password is not true")
	ErrInCorrectMethod   = errors.New("incorrect method")
	ErrAnotherUsersOrder = errors.New("another user`s order")
	ErrOrderIsAlready    = errors.New("the older is already there")
	ErrNotData           = errors.New("not data to answer")
	ErrNotEnoughBonuses  = errors.New("not enough bonuses")
)

type APIError struct {
	Message string `json:"message"`
}
