package converter

import (
	"context"
	"go-prototype-new/internal/entity"
	"go-prototype-new/internal/response"
)

func ConvertGetHome(ctx context.Context, home []entity.Home) []response.HomeResponse {
	var res []response.HomeResponse

	for _, d := range home {
		res = append(res, response.HomeResponse{
			Name:        d.Name,
			Description: d.Description,
			Logo:        d.Logo.String,
			Favicon:     d.Favicon.String,
			Phone:       d.Phone,
			Email:       d.Email,
		})
	}

	return res
}
