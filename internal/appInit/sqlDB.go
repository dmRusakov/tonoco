package appInit

import (
	"github.com/dmRusakov/tonoco/pkg/common/core/closer"
	"github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

// productDBInit - product database initialization
func (a *App) productDBInit() (err error) {
	a.sqlDB, err = postgresql.NewClient(a.Ctx, 5, 3*time.Second, a.Cfg.DataStorage.ToPostgreSQLConfig(), false)
	if err != nil {
		return err
	}

	closer.AddN(a.sqlDB)

	return nil
}
