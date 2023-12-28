package mysql

import (
	"SSO/internal/domain/models"
	"SSO/internal/storage/storageErrors"
	"context"
	"database/sql"
	"errors"
)

type AppStorage struct {
	db *sql.DB
}

func NewAppStorage(db *sql.DB) *AppStorage {
	return &AppStorage{
		db: db,
	}
}

func (a *AppStorage) Save(ctx context.Context, key []byte) error {
	if _, err := a.db.ExecContext(ctx, "INSERT INTO apps (key) VALUES (?)", key); err != nil {
		return err
	}
	return nil
}

func (a *AppStorage) GetByKey(ctx context.Context, key []byte) (models.App, error) {
	var app models.App
	if err := a.db.QueryRowContext(ctx, "SELECT * FROM apps WHERE secret_key=?", key).Scan(
		&app.Id, &app.Key,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return app, storageErrors.ErrAppNotFound
		}
		return app, err
	}
	return app, nil
}
