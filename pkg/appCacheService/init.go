package appCacheService

var _ AppCacheService = &appCacheService{}

type appCacheService struct {
	authPrefix *string
}

type AppCacheService interface {
	GetAuthPrefix() string
}

func NewCacheService(authPrefix string) (*appCacheService, error) {
	return &appCacheService{
		&authPrefix,
	}, nil
}
