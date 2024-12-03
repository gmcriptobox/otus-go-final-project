//go:generate mockgen -source=$GOFILE
//-destination ./mocks/mock_list.go -package mockrepository MockListRepo

package repository

import (
	"context"

	"github.com/gmcriptobox/otus-go-final-project/internal/entity"
)

type IListRepo interface {
	IsExists(ctx context.Context, ip, mask string) (bool, error)
	Add(ctx context.Context, network entity.Network) error
	GetAll(ctx context.Context) ([]entity.Network, error)
	Remove(ctx context.Context, ip, mask string) error
}
