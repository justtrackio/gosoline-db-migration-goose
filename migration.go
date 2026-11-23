package goose

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/db"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/pressly/goose/v3"
)

func runMigration(ctx context.Context, logger log.Logger, settings *db.Settings, db *sql.DB) error {
	var err error

	goose.SetLogger(newGooseLogger(logger))

	if err = goose.SetDialect(settings.Driver); err != nil {
		return fmt.Errorf("can not set db dialect: %w", err)
	}

	if err = goose.UpContext(ctx, db, settings.Migrations.Path, goose.WithAllowMissing()); err != nil {
		return fmt.Errorf("can not run up migrations from path %s: %w", settings.Migrations.Path, err)
	}

	return nil
}

type gooseLogger struct {
	logger log.Logger
}

func newGooseLogger(logger log.Logger) goose.Logger {
	return gooseLogger{
		logger: logger,
	}
}

func (g gooseLogger) Fatal(v ...interface{}) {
	g.logger.Error(context.Background(), fmt.Sprint(v...))
}

func (g gooseLogger) Fatalf(format string, v ...interface{}) {
	g.logger.Error(context.Background(), format, v...)
}

func (g gooseLogger) Print(v ...interface{}) {
	g.logger.Info(context.Background(), fmt.Sprint(v...))
}

func (g gooseLogger) Println(v ...interface{}) {
	g.logger.Info(context.Background(), fmt.Sprint(v...))
}

func (g gooseLogger) Printf(format string, v ...interface{}) {
	g.logger.Info(context.Background(), format, v...)
}
