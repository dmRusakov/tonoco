package appInit

import (
	"github.com/dmRusakov/tonoco/pkg/common/core/clock"
	"github.com/dmRusakov/tonoco/pkg/common/core/identity"
)

// clockInit - clock initialization
func (a *App) clockInit() (err error) {
	// if clock already initialized
	if a.clock != nil {
		return nil
	}

	// new clock
	a.clock = clock.New()
	return nil

}

// generatorInit - generator initialization
func (a *App) generatorInit() (err error) {
	// if generator already initialized
	if a.generator != nil {
		return nil
	}

	// new generator
	a.generator = identity.NewGenerator()
	return nil
}
