package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *GearboxService) Create(ctx context.Context, req *dto.CreateGearboxRequest) (*dto.GearboxResponse, error) {

	return s.base.Create(ctx, req)

}

//update

func (s *GearboxService) Update(ctx context.Context, id int, req *dto.UpdateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Update(ctx, id, req)

}

//Delete

func (s *GearboxService) Delete(ctx context.Context, id int) error {

	return s.base.Delete(ctx, id)
}

// GetById
func (s *GearboxService) GetById(ctx context.Context, id int) (*dto.GearboxResponse, error) {
	return s.base.GetById(ctx, id)

}
func (s *GearboxService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
