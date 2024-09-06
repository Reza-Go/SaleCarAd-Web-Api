package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *PersianYearService) Create(ctx context.Context, req *dto.CreatePersianYearRequest) (*dto.PersianYearResponse, error) {

	return s.base.Create(ctx, req)

}

//update

func (s *PersianYearService) Update(ctx context.Context, id int, req *dto.UpdatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Update(ctx, id, req)

}

//Delete

func (s *PersianYearService) Delete(ctx context.Context, id int) error {

	return s.base.Delete(ctx, id)
}

// GetById
func (s *PersianYearService) GetById(ctx context.Context, id int) (*dto.PersianYearResponse, error) {
	return s.base.GetById(ctx, id)

}
func (s *PersianYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PersianYearResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
