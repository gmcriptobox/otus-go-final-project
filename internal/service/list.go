package service

import (
	"context"
	"strings"

	"github.com/gmcriptobox/otus-go-final-project/internal/repository"
)

type ListService struct {
	repo repository.IListRepo
}

func NewListService(repo repository.IListRepo) *ListService {
	return &ListService{
		repo: repo,
	}
}

func (l *ListService) IsContains(ctx context.Context, ip string) (bool, error) {
	networks, err := l.repo.GetAll(ctx)
	if err != nil {
		return false, err
	}

	binaryIP, err := IPAddressToBinary(ip)
	if err != nil {
		return false, err
	}

	for _, network := range networks {
		if strings.HasPrefix(binaryIP, network.BinaryPrefix) {
			return true, nil
		}
	}
	return false, err
}

func (l *ListService) Add(ctx context.Context, network string) error {
	networkEntity, err := GetNetwork(network)
	if err != nil {
		return err
	}

	isExists, err := l.repo.IsExists(ctx, networkEntity.IP, networkEntity.Mask)
	if err != nil {
		return err
	}
	if isExists {
		return ErrNetworkAlreadyExists
	}

	err = l.repo.Add(ctx, networkEntity)
	if err != nil {
		return err
	}
	return nil
}

func (l *ListService) Remove(ctx context.Context, network string) error {
	networkEntity, err := GetNetwork(network)
	if err != nil {
		return err
	}

	err = l.repo.Remove(ctx, networkEntity.IP, networkEntity.Mask)
	if err != nil {
		return err
	}
	return nil
}
