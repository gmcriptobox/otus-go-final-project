package service

import (
	"context"

	"github.com/gmcriptobox/otus-go-final-project/internal/config"
	"github.com/gmcriptobox/otus-go-final-project/internal/entity/request"
)

type Authorization struct {
	ipRateLimit       *RateLimit
	loginRateLimit    *RateLimit
	passwordRateLimit *RateLimit
	blackList         *ListService
	whiteList         *ListService
}

func NewAuthorization(config config.Config, blackList *ListService, whiteList *ListService) *Authorization {
	return &Authorization{
		ipRateLimit:       NewRateLimit(config.Bucket.IPLimit, config.Bucket.BucketTTL),
		loginRateLimit:    NewRateLimit(config.Bucket.LoginLimit, config.Bucket.BucketTTL),
		passwordRateLimit: NewRateLimit(config.Bucket.PasswordLimit, config.Bucket.BucketTTL),
		blackList:         blackList,
		whiteList:         whiteList,
	}
}

func (a *Authorization) Authorize(ctx context.Context, request request.AuthRequest) (bool, error) {
	isContains, err := a.blackList.IsContains(ctx, request.IP)
	if err != nil {
		return false, err
	}
	if isContains {
		return false, nil
	}

	isContains, err = a.whiteList.IsContains(ctx, request.IP)
	if err != nil {
		return false, err
	}
	if isContains {
		return true, nil
	}

	if !a.ipRateLimit.Allow(request.IP) {
		return false, nil
	}
	if !a.loginRateLimit.Allow(request.Login) {
		return false, nil
	}
	if !a.passwordRateLimit.Allow(request.Password) {
		return false, nil
	}
	return true, nil
}

func (a *Authorization) ResetBuckets(request request.BucketResetRequest) {
	a.ipRateLimit.ResetBucket(request.IP)
	a.loginRateLimit.ResetBucket(request.Login)
}
