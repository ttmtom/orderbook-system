package service

import (
	"orderbook/internal/adapter/database/postgres/repository"
)

type CommonService struct {
	repo *repository.CommonRepository
}

func NewCommonService(resp *repository.CommonRepository) *CommonService {
	return &CommonService{
		resp,
	}
}

func (cs *CommonService) GetAccessTokenTimeLimit() (uint, error) {
	accessTimeLimit, err := cs.repo.GetTimeLimit("access_token")

	if err != nil {
		return 0, err
	}
	return accessTimeLimit.Time, nil
}

func (cs *CommonService) GetRefreshTokenTimeLimit() (uint, error) {
	refreshTimeLimit, err := cs.repo.GetTimeLimit("refresh_token")

	if err != nil {
		return 0, err
	}
	return refreshTimeLimit.Time, nil
}

type JwtTokenTimeLimit struct {
	AccessTokenDuration  uint
	RefreshTokenDuration uint
}

func (cs *CommonService) GetJwtTokenTimeLimit() (*JwtTokenTimeLimit, error) {
	accessTokenTimeLimit, err := cs.GetAccessTokenTimeLimit()
	if err != nil {
		return nil, err
	}

	refreshTokenTimeLimit, err := cs.GetRefreshTokenTimeLimit()
	if err != nil {
		return nil, err
	}
	return &JwtTokenTimeLimit{
		accessTokenTimeLimit,
		refreshTokenTimeLimit,
	}, nil
}
