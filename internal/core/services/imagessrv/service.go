package imagessrv

import (
	"exporterbackend/internal/common/helper/aws"
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"fmt"
)

type Service struct {
	logger     logging.Logger
	imagesRepo ports.RdbmsImagesRepository
	s3Service  aws.S3Service
}

func New(logger logging.Logger,
	imagesRepo ports.RdbmsImagesRepository,
	s3Service aws.S3Service,
) *Service {
	return &Service{
		logger:     logger,
		imagesRepo: imagesRepo,
		s3Service:  s3Service,
	}
}

func (s *Service) getPresignedURLForUpload(fileName string, contentType string) (string, error) {
	return s.s3Service.GetPresignedURLForUpload("/imgs/"+fileName, contentType)
}

func (s *Service) GetPresignedURLForDownload(fileName string) (string, error) {
	return s.s3Service.GetPresignedURLForDownload(fileName)
}
func (s *Service) GetSignedURLAndSave(
	f []rdbms.CreateImage,
) ([]rdbms.UploadImageRes, error) {
	var res []rdbms.UploadImageRes
	for _, v := range f {
		if id, er := s.imagesRepo.Insert(v); er != nil {
			fmt.Println(er, "Errrrr")
			return nil, er
		} else {
			fmt.Println(id, "id")
			if url, er := s.getPresignedURLForUpload(v.FileName, v.MimeType); er != nil {
				fmt.Println(er, "errrrr")
				return nil, er
			} else {
				fmt.Println(url, "idzxczxcz")
				res = append(res, rdbms.UploadImageRes{
					Id:        id,
					SignedURL: url,
				})
			}
		}

	}
	return res, nil
}
