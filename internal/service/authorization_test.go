package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/gmcriptobox/otus-go-final-project/internal/config"
	"github.com/gmcriptobox/otus-go-final-project/internal/entity"
	"github.com/gmcriptobox/otus-go-final-project/internal/entity/request"
	mockrepository "github.com/gmcriptobox/otus-go-final-project/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type AuthCall struct {
	call      request.AuthRequest
	isAllowed bool
}

type TestCase struct {
	name        string
	authCall    []AuthCall
	whiteList   []entity.Network
	blackList   []entity.Network
	isBlackList bool
}

// buckets have this limit setting: ipLimit: 5, loginLimit: 3, passwordLimit: 3.
func TestAuthorize(t *testing.T) {
	testCases := []TestCase{
		{
			name: "Case with single auth call",
			authCall: []AuthCall{
				{request.AuthRequest{Login: "login", Password: "password", IP: "11.2.3.4"}, true},
			},
		},
		{
			name: "Case with password brute force attack",
			authCall: []AuthCall{
				{request.AuthRequest{Login: "login", Password: "password", IP: "128.128.10.33"}, true},
				{request.AuthRequest{Login: "login", Password: "password1", IP: "128.128.10.33"}, true},
				{request.AuthRequest{Login: "login", Password: "password2", IP: "128.128.10.33"}, true},
				{request.AuthRequest{Login: "login", Password: "password3", IP: "128.128.10.33"}, false},
			},
		},
		{
			name: "Case with login brute force attack",
			authCall: []AuthCall{
				{request.AuthRequest{Login: "login", Password: "qwerty123", IP: "128.128.10.33"}, true},
				{request.AuthRequest{Login: "randomlogin", Password: "qwerty123", IP: "128.128.10.33"}, true},
				{request.AuthRequest{Login: "uzver", Password: "qwerty123", IP: "128.128.10.33"}, true},
				{request.AuthRequest{Login: "wrrrrrr", Password: "qwerty123", IP: "128.128.10.33"}, false},
			},
		},
		{
			name: "Case when brute force attack from same ip",
			authCall: []AuthCall{
				{request.AuthRequest{Login: "trdfg", Password: "nmfawe52", IP: "77.42.11.32"}, true},
				{request.AuthRequest{Login: "235sfasf", Password: "nvvbnvqsa", IP: "77.42.11.32"}, true},
				{request.AuthRequest{Login: "yuyu66", Password: "7tittt", IP: "77.42.11.32"}, true},
				{request.AuthRequest{Login: "vbbbbbb", Password: "vivivivi", IP: "77.42.11.32"}, true},
				{request.AuthRequest{Login: "nncwww", Password: "qwerrtt", IP: "77.42.11.32"}, true},
				{request.AuthRequest{Login: "xccvbx", Password: "zxcqrewaa", IP: "77.42.11.32"}, false},
				{request.AuthRequest{Login: "oyqsss", Password: "vbcbwww", IP: "77.42.11.32"}, false},
			},
		},
		{
			name: "Case with login brute force attack with ip in white list",
			authCall: []AuthCall{
				{request.AuthRequest{Login: "login", Password: "qwerty123", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "randomlogin", Password: "qwerty123", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "uzver", Password: "qwerty123", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "wrrrrrr", Password: "qwerty123", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "olrqwe", Password: "qwerty123", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "misha228", Password: "qwerty123", IP: "192.128.12.66"}, true},
			},
			whiteList: []entity.Network{
				{IP: "192.128.12.0", Mask: "24", BinaryPrefix: "110000001000000000001100"},
			},
		},
		{
			name: "Case with password brute force attack with ip in white list",
			authCall: []AuthCall{
				{request.AuthRequest{Login: "user322", Password: "qwerty123", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "user322", Password: "322qwerty", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "user322", Password: "asdfghjkl", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "user322", Password: "zxccvbbvb", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "user322", Password: "53rasdfcsdg", IP: "192.128.12.66"}, true},
				{request.AuthRequest{Login: "user322", Password: "123512sdzzz", IP: "192.128.12.66"}, true},
			},
			whiteList: []entity.Network{
				{IP: "192.128.12.0", Mask: "24", BinaryPrefix: "110000001000000000001100"},
			},
		},
		{
			name: "Case when ip in black list",
			authCall: []AuthCall{
				{request.AuthRequest{Login: "oooooooo", Password: "passnotword", IP: "112.32.56.4"}, false},
			},
			blackList: []entity.Network{
				{IP: "112.32.56.0", Mask: "24", BinaryPrefix: "011100000010000000111000"},
			},
			isBlackList: true,
		},
		{
			name: "Case when ip in black list and in white list",
			authCall: []AuthCall{
				{request.AuthRequest{Login: "oooooooo", Password: "passnotword", IP: "112.32.56.151"}, false},
			},
			whiteList: []entity.Network{
				{IP: "112.32.56.0", Mask: "24", BinaryPrefix: "011100000010000000111000"},
			},
			blackList: []entity.Network{
				{IP: "112.32.56.0", Mask: "24", BinaryPrefix: "011100000010000000111000"},
			},
			isBlackList: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			blackList, whiteList, authorization := getAuthService(t, ctrl)

			for _, authCall := range tc.authCall {
				ctx := context.TODO()
				if !tc.isBlackList {
					whiteList.EXPECT().GetAll(ctx).Return(tc.whiteList, nil)
				}
				blackList.EXPECT().GetAll(ctx).Return(tc.blackList, nil)

				isAllowed, err := authorization.Authorize(ctx, authCall.call)

				require.NoError(t, err)
				require.Equal(t, authCall.isAllowed, isAllowed)
			}
		})
	}
}

func TestBucketReset(t *testing.T) {
	authCalls := []AuthCall{
		{request.AuthRequest{Login: "login", Password: "password", IP: "128.128.10.33"}, true},
		{request.AuthRequest{Login: "login", Password: "password1", IP: "128.128.10.33"}, true},
		{request.AuthRequest{Login: "login", Password: "password2", IP: "128.128.10.33"}, true},
		{request.AuthRequest{Login: "login", Password: "password3", IP: "128.128.10.33"}, false},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	blackList, whiteList, authorization := getAuthService(t, ctrl)

	for i := 0; i < 2; i++ {
		for _, authCall := range authCalls {
			ctx := context.TODO()
			whiteList.EXPECT().GetAll(ctx).Return([]entity.Network{}, nil)
			blackList.EXPECT().GetAll(ctx).Return([]entity.Network{}, nil)

			isAllowed, err := authorization.Authorize(ctx, authCall.call)

			require.NoError(t, err)
			require.Equal(t, authCall.isAllowed, isAllowed)
		}
		authorization.ResetBuckets(request.BucketResetRequest{Login: "login", IP: "128.128.10.33"})
	}
}

func getAuthService(t *testing.T, ctrl *gomock.Controller) (*mockrepository.MockIListRepo,
	*mockrepository.MockIListRepo, *Authorization,
) {
	t.Helper()

	blackList := mockrepository.NewMockIListRepo(ctrl)
	blackListService := NewListService(blackList)

	whiteList := mockrepository.NewMockIListRepo(ctrl)
	whiteListService := NewListService(whiteList)

	configuration := getConfig(t)
	authorization := NewAuthorization(configuration, blackListService, whiteListService)
	return blackList, whiteList, authorization
}

func getConfig(t *testing.T) config.Config {
	t.Helper()

	projectConfig, err := config.Read("../../configs/test_config.yaml")
	if err != nil {
		fmt.Println("error while reading config file: ", err)
	}
	require.NoError(t, err)

	return projectConfig
}
