package app

import (
	"github.com/FayxChance/tha-backend-tezos/internal/infrastructure/persistence"
	"github.com/FayxChance/tha-backend-tezos/internal/router"
)

type App struct {
	Router   router.Router
	Database persistence.SQLite3Database
}

func (a *App) SetupApp() error {
	var db persistence.SQLite3Database
	err := db.SetupDatabase()
	if err != nil {
		return err
	}
	a.Database = db

	var r router.Router
	err = r.SetupRouter(db)
	if err != nil {
		return err
	}
	a.Router = r
	return nil
}
