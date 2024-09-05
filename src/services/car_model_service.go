package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Company.Country"},
				{string: "CarType"},
				{string: "Gearbox"},
				{string: "CarModelColors.Color"},
				{string: "CarModelYears.PersianYear"},
				{string: "CarModelYears.CarModelPriceHistories"},
				{string: "CarModelProperties.Property.Category"},
				{string: "CarModelImages.Image"},
				{string: "CarModelComments.User"},
			},
		},
	}
}

func (s *CarModelService) Create(ctx context.Context, req *dto.CreateCarModelRequest) (*dto.CarModelResponse, error) {

	return s.base.Create(ctx, req)

}

//update

func (s *CarModelService) Update(ctx context.Context, id int, req *dto.UpdateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Update(ctx, id, req)

}

//Delete

func (s *CarModelService) Delete(ctx context.Context, id int) error {

	return s.base.Delete(ctx, id)
}

// GetById
func (s *CarModelService) GetById(ctx context.Context, id int) (*dto.CarModelResponse, error) {
	return s.base.GetById(ctx, id)

}
func (s *CarModelService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
