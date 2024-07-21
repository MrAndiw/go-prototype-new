package model

import (
	"context"
	"fmt"
	"go-prototype-new/internal/entity"

	"github.com/jmoiron/sqlx"
)

type (
	HomeModel interface {
		GetHome(ctx context.Context, db *sqlx.DB) []entity.Home
	}

	homeModel struct {
	}
)

func NewHomeModel() HomeModel {
	return &homeModel{}
}

func (m *homeModel) GetHome(ctx context.Context, db *sqlx.DB) []entity.Home {

	// Query the database, storing results in a []Person (wrapped in []interface{})
	home := []entity.Home{}
	err := db.Select(&home, "SELECT id,name,description,logo,favicon,phone,email FROM public.home ORDER BY id ASC")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("home : ", home)
	return home
}
