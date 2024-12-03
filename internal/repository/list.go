package repository

import (
	"context"
	"fmt"

	"github.com/gmcriptobox/otus-go-final-project/internal/entity"
	"github.com/gmcriptobox/otus-go-final-project/internal/repository/client"
)

var (
	isExistsSQL  = "SELECT EXISTS(SELECT 1 FROM %s WHERE ip = $1 and mask = $2)"
	insertSQL    = "INSERT INTO %s (ip, mask, binary_prefix) VALUES ($1, $2, $3)"
	selectAllSQL = "SELECT ip, mask, binary_prefix FROM %s"
	deleteSQL    = "DELETE FROM %s WHERE ip = $1 and mask = $2"
)

type ListRepo struct {
	client    *client.PostgresSQL
	tableName string
}

func NewListRepo(client *client.PostgresSQL, tableName string) *ListRepo {
	return &ListRepo{
		client:    client,
		tableName: tableName,
	}
}

func (l *ListRepo) IsExists(ctx context.Context, ip, mask string) (bool, error) {
	var exists bool
	err := l.client.DB.QueryRowContext(ctx, fmt.Sprintf(isExistsSQL, l.tableName), ip, mask).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (l *ListRepo) Add(ctx context.Context, network entity.Network) error {
	_, err := l.client.DB.ExecContext(ctx, fmt.Sprintf(insertSQL, l.tableName), network.IP,
		network.Mask, network.BinaryPrefix)
	if err != nil {
		return err
	}
	return nil
}

func (l *ListRepo) GetAll(ctx context.Context) ([]entity.Network, error) {
	networkList := make([]entity.Network, 0)
	err := l.client.DB.SelectContext(ctx, &networkList, fmt.Sprintf(selectAllSQL, l.tableName))
	if err != nil {
		return nil, err
	}
	return networkList, nil
}

func (l *ListRepo) Remove(ctx context.Context, ip, mask string) error {
	_, err := l.client.DB.ExecContext(ctx, fmt.Sprintf(deleteSQL, l.tableName), ip, mask)
	if err != nil {
		return err
	}
	return nil
}
