package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
)

type PropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	return &PropertyCategoryService{
		base: &BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{string: "Properties"}},
		},
	}
}

// Create
func (s *PropertyCategoryService) Create(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {

	return s.base.Create(ctx, req)

}

//update

func (s *PropertyCategoryService) Update(ctx context.Context, id int, req *dto.UpdatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return s.base.Update(ctx, id, req)

}

//Delete

func (s *PropertyCategoryService) Delete(ctx context.Context, id int) error {

	return s.base.Delete(ctx, id)

}

// GetById
func (s *PropertyCategoryService) GetById(ctx context.Context, id int) (*dto.PropertyCategoryResponse, error) {
	return s.base.GetById(ctx, id)

}
func (s *PropertyCategoryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyCategoryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
