package appInit

import (
	"context"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/redisdb"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
)

// redisClientInit - redis client initialization
func (a *App) redisClientInit() (err error) {
	// check if redis client already initialized
	if a.cacheDB == nil {
		return nil
	}

	// connect to redis
	a.cacheDB, err = redisdb.Connect(context.Background(), a.Cfg.CacheStorage.ToRedisConfig())
	if err != nil {
		return err
	}

	return nil
}

// AppCacheServiceInit - appCacheService initialization
func (a *App) AppCacheServiceInit() (err error) {
	// check if appCacheService already initialized
	if a.AppCacheService != nil {
		return nil
	}

	// init redis client
	err = a.redisClientInit()
	if err != nil {
		return err
	}

	a.AppCacheService, err = appCacheService.NewCacheService(a.cacheDB, "app")
	if err != nil {
		return err
	}

	return nil
}

// UserCacheServiceInit - userCacheService initialization
func (a *App) UserCacheServiceInit() (err error) {
	// check if userCacheService already initialized
	if a.UserCacheService != nil {
		return nil
	}

	// check redis client init
	err = a.redisClientInit()
	if err != nil {
		return err
	}

	a.UserCacheService, err = userCacheService.NewCacheService(a.cacheDB, "user")
	if err != nil {
		return err
	}

	return nil
}
