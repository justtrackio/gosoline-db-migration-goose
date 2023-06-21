package goose

import (
	"context"

	"github.com/justtrackio/gosoline/pkg/application"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/db"
)

func Plugin(_ context.Context, _ cfg.GosoConf) ([]application.Option, error) {
	db.AddMigrationProvider("goose", runMigration)

	return nil, nil
}
