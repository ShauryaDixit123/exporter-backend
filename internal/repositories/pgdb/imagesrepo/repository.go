package imagesrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofrs/uuid"
)

type Repository struct {
	logger   logging.Logger
	dbClient *goqu.Database
}

func New(logger logging.Logger,
	dbClient *goqu.Database) *Repository {
	return &Repository{
		logger:   logger,
		dbClient: dbClient,
	}
}

func (r *Repository) Insert(d rdbms.CreateImage) (string, error) {
	var id uuid.UUID
	if _, er := r.dbClient.Insert(TABLE).Rows(
		goqu.Record{
			S3PATH:      d.S3Path,
			MIMETYPE:    d.MimeType,
			UPLOADED_BY: d.UploadedBy,
		},
	).Returning("id").Executor().ScanVal(&id); er != nil {
		return "", er
	}
	return id.String(), nil
}
