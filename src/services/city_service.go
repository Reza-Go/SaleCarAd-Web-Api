package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateUpdateCityRequest, dto.CreateUpdateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: &BaseService[models.City, dto.CreateUpdateCityRequest, dto.CreateUpdateCityRequest, dto.CityResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Country"},
			},
		},
	}
}

func (s *CityService) Create(ctx context.Context, req *dto.CreateUpdateCityRequest) (*dto.CityResponse, error) {

	return s.base.Create(ctx, req)

}

//update

func (s *CityService) Update(ctx context.Context, id int, req *dto.CreateUpdateCityRequest) (*dto.CityResponse, error) {
	return s.base.Update(ctx, id, req)

}

//Delete

func (s *CityService) Delete(ctx context.Context, id int) error {

	return s.base.Delete(ctx, id)
}

// GetById
func (s *CityService) GetById(ctx context.Context, id int) (*dto.CityResponse, error) {
	return s.base.GetById(ctx, id)

}
func (s *CityService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CityResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
