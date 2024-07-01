package model

import (
	"fmt"
	"go-prototype-new/internal/app/database"
)

type (
	HomeModel interface {
		GetHome(message string) string
	}

	homeModel struct {
		db database.Service
	}
)

func NewHomeModel(db database.Service) HomeModel {
	return &homeModel{
		db: db,
	}
}

func (m *homeModel) GetHome(message string) string {

	return fmt.Sprintf("Welcome to My Home %s", message)
}
