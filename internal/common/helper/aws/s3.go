package aws

import (
	"exporterbackend/internal/configs"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service interface {
	GetPresignedURLForUpload(s3Path string, contentType string) (string, error)
	GetPresignedURLForDownload(s3Path string) (string, error)
}

type S3 struct {
	s3Client *s3.S3
	s3Config configs.S3Config
}

func NewS3(s3Config configs.S3Config, s3Client *s3.S3) *S3 {
	return &S3{
		s3Client: s3Client,
		s3Config: s3Config,
	}
}

func (s *S3) GetPresignedURLForUpload(
	s3Path string,
	contentType string,
) (string, error) {
	req, _ := s.s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      &s.s3Config.Bucket,
		Key:         &s3Path,
		ContentType: &contentType,
	})
	return req.Presign(15 * time.Minute)
}

func (s *S3) GetPresignedURLForDownload(
	s3Path string,
) (string, error) {
	req, _ := s.s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &s.s3Config.Bucket,
		Key:    &s3Path,
	})
	return req.Presign(15 * time.Minute)
}
