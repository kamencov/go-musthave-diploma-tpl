package db

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"testing"
)

func TestDateBase_SaveTableUserAndUpdateToken(t *testing.T) {
	tests := []struct {
		name        string
		login       string
		accessToken string
		expentedErr error
	}{
		{
			name:        "Success",
			login:       "test",
			accessToken: "test",
		},
		{
			name:        "Error",
			login:       "test",
			accessToken: "test",
			expentedErr: customerrors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			repo := NewMockUser(ctrl)

			repo.EXPECT().SaveTableUserAndUpdateToken(gomock.Any(), gomock.Any()).Return(tt.expentedErr).AnyTimes()

			repo.SaveTableUserAndUpdateToken(tt.login, tt.accessToken)

			ctrl.Finish()

		})
	}
}

func TestDateBase_GetLoginID(t *testing.T) {

	tests := []struct {
		name        string
		login       string
		loginID     int
		expentedErr error
	}{
		{
			name:    "Success",
			login:   "test",
			loginID: 1,
		},
		{
			name:        "Error",
			login:       "test",
			loginID:     1,
			expentedErr: customerrors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := NewMockUser(ctrl)
			repo.EXPECT().GetLoginID(gomock.Any()).Return(tt.loginID, tt.expentedErr).AnyTimes()
			_, err := repo.GetLoginID(tt.login)
			if !errors.Is(err, tt.expentedErr) {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestDateBase_SearchLoginByToken(t *testing.T) {

	tests := []struct {
		name        string
		accessToken string
		searchTokin string
		expentedErr error
	}{
		{
			name:        "Success",
			accessToken: "test",
			searchTokin: "test",
		},
		{
			name:        "Error",
			accessToken: "test",
			searchTokin: "test",
			expentedErr: customerrors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := NewMockUser(ctrl)
			repo.EXPECT().SearchLoginByToken(gomock.Any()).Return(tt.searchTokin, tt.expentedErr).AnyTimes()
			_, err := repo.SearchLoginByToken(tt.accessToken)
			if !errors.Is(err, tt.expentedErr) {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestDateBase_CheckTableUserLogin(t *testing.T) {

	tests := []struct {
		name        string
		login       string
		expentedErr error
	}{
		{
			name:  "Success",
			login: "test",
		},
		{
			name:        "Error",
			login:       "test",
			expentedErr: customerrors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := NewMockUser(ctrl)
			repo.EXPECT().CheckTableUserLogin(gomock.Any()).Return(tt.expentedErr).AnyTimes()
			err := repo.CheckTableUserLogin(tt.login)

			if !errors.Is(err, tt.expentedErr) {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestDateBase_CheckTableUserPassword(t *testing.T) {
	tests := []struct {
		name             string
		existingPassword string
		expentedErr      bool
	}{
		{
			name:             "Success",
			existingPassword: "test",
		},
		{
			name:             "Error",
			existingPassword: "test",
			expentedErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := NewMockUser(ctrl)
			repo.EXPECT().CheckTableUserPassword(gomock.Any()).Return(tt.existingPassword, tt.expentedErr).AnyTimes()
			_, err := repo.CheckTableUserPassword(tt.existingPassword)

			if err != tt.expentedErr {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
