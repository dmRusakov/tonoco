package userCacheService

var _ UserCacheService = &userCacheService{}

type userCacheService struct {
	authPrefix *string
}

type UserCacheService interface {
	GetAuthPrefix() string
}

func NewCacheService(authPrefix string) (*userCacheService, error) {
	return &userCacheService{
		&authPrefix,
	}, nil
}
