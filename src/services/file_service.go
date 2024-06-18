package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
)

type FileService struct {
	base *BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]
}

func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		base: &BaseService[models.File, dto.CreateFileRequest, dto.UpdateFileRequest, dto.FileResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *FileService) Create(ctx context.Context, req *dto.CreateFileRequest) (*dto.FileResponse, error) {

	return s.base.Create(ctx, req)

}

//update

func (s *FileService) Update(ctx context.Context, id int, req *dto.UpdateFileRequest) (*dto.FileResponse, error) {
	return s.base.Update(ctx, id, req)

}

//Delete

func (s *FileService) Delete(ctx context.Context, id int) error {

	return s.base.Delete(ctx, id)
}

// GetById
func (s *FileService) GetById(ctx context.Context, id int) (*dto.FileResponse, error) {
	return s.base.GetById(ctx, id)

}
func (s *FileService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.FileResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
