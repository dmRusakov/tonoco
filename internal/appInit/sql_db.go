package appInit

import (
	"github.com/dmRusakov/tonoco/pkg/common/core/closer"
	"github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

// SqlDBInit - sql database initialization
func (a *App) SqlDBInit() (err error) {
	// if already initialized
	if a.sqlDB != nil {
		return nil
	}

	// new sqlDB
	a.sqlDB, err = postgresql.NewClient(a.Ctx, 5, 3*time.Second, a.Cfg.DataStorage.ToPostgreSQLConfig(), false)
	if err != nil {
		return err
	}

	closer.AddN(a.sqlDB)

	return nil
}
