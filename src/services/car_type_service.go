package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: &BaseService[models.CarType, dto.CreateCarTypeRequest, dto.UpdateCarTypeRequest, dto.CarTypeResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *CarTypeService) Create(ctx context.Context, req *dto.CreateCarTypeRequest) (*dto.CarTypeResponse, error) {

	return s.base.Create(ctx, req)

}

//update

func (s *CarTypeService) Update(ctx context.Context, id int, req *dto.UpdateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Update(ctx, id, req)

}

//Delete

func (s *CarTypeService) Delete(ctx context.Context, id int) error {

	return s.base.Delete(ctx, id)
}

// GetById
func (s *CarTypeService) GetById(ctx context.Context, id int) (*dto.CarTypeResponse, error) {
	return s.base.GetById(ctx, id)

}
func (s *CarTypeService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarTypeResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
