package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
)

type CarModelImageService struct {
	base *BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.UpdateCarModelImageRequest, dto.CarModelImageResponse]
}

func NewCarModelImageService(cfg *config.Config) *CarModelImageService {
	return &CarModelImageService{
		base: &BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.UpdateCarModelImageRequest, dto.CarModelImageResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Image"},
			},
		},
	}
}

func (s *CarModelImageService) Create(ctx context.Context, req *dto.CreateCarModelImageRequest) (*dto.CarModelImageResponse, error) {

	return s.base.Create(ctx, req)

}

//update

func (s *CarModelImageService) Update(ctx context.Context, id int, req *dto.UpdateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Update(ctx, id, req)

}

//Delete

func (s *CarModelImageService) Delete(ctx context.Context, id int) error {

	return s.base.Delete(ctx, id)
}

// GetById
func (s *CarModelImageService) GetById(ctx context.Context, id int) (*dto.CarModelImageResponse, error) {
	return s.base.GetById(ctx, id)

}
func (s *CarModelImageService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelImageResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
