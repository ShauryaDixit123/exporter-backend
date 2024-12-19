package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsImagesRepository interface {
	Insert(d rdbms.CreateImage) (string, error)
}
type ImagesService interface {
	GetPresignedURLForDownload(fileName string) (string, error)
	GetSignedURLAndSave(
		f []rdbms.CreateImage,
	) ([]rdbms.UploadImageRes, error)
}
