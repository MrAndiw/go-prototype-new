package repository

import (
	"context"
	"go-prototype-new/internal/app/database"
	"go-prototype-new/internal/converter"
	"go-prototype-new/internal/model"
	"go-prototype-new/internal/response"
)

type (
	HomeRepository interface {
		GetHome(ctx context.Context) []response.HomeResponse
	}

	homeRepository struct {
		DB   database.Service
		Home model.HomeModel
	}
)

func NewHomeRepository(db database.Service, home model.HomeModel) HomeRepository {
	return &homeRepository{
		DB:   db,
		Home: home,
	}
}

func (rep *homeRepository) GetHome(ctx context.Context) []response.HomeResponse {
	home := rep.Home.GetHome(ctx, rep.DB.Connect())

	return converter.ConvertGetHome(ctx, home)
}
