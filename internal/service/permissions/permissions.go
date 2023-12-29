package permissions

import (
	"SSO/internal/storage"
	"context"
	"fmt"
	"log/slog"
)

type Permissions struct {
	l           *slog.Logger
	permStorage storage.PermissionsStorage
}

func New(l *slog.Logger, permStorage storage.PermissionsStorage) *Permissions {
	return &Permissions{
		l:           l,
		permStorage: permStorage,
	}
}

func (p *Permissions) SetUserPermission(ctx context.Context, userId int, permission int32) (err error) {
	const op = "service.permissions.SetUserPermission"
	if err := p.permStorage.Save(ctx, userId, permission); err != nil {
		p.l.Error(fmt.Errorf("%s: %w", op, err).Error())
		return err
	}
	return nil
}

func (p *Permissions) GetUserPermission(ctx context.Context, userId int) (permission int32, err error) {
	const op = "service.permissions.GetUserPermission"
	perm, err := p.permStorage.Get(ctx, userId)
	if err != nil {
		p.l.Error(fmt.Errorf("%s: %w", op, err).Error())
		return perm, err
	}
	return perm, nil
}
