package repository

import (
	"context"
	"time"

	"github.com/his-vita/patients-service/pkg/database/postgres"
)

type (
	SqlStore interface {
		GetQuery(queryName string) (string, error)
	}

	DbHandler interface {
		TxOrDb(ctx context.Context) postgres.DbExecutor
		WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc)
		DefaultTimeoutCtx() (context.Context, context.CancelFunc)
	}
)
