package appInit

import (
	"github.com/dmRusakov/tonoco/pkg/common/core/clock"
	"github.com/dmRusakov/tonoco/pkg/common/core/identity"
)

// clockInit - clock initialization
func (a *App) clockInit() (err error) {
	a.clock = clock.New()
	return nil

}

// generatorInit - generator initialization
func (a *App) generatorInit() (err error) {
	a.generator = identity.NewGenerator()
	return nil
}
