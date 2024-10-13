package app

import (
	"github.com/FayxChance/tha-backend-tezos/internal/infrastructure/persistence"
	"github.com/FayxChance/tha-backend-tezos/internal/router"
)

// App struct represents the main application structure.
// It holds references to the router and the database.
type App struct {
	Router   router.Router               // Router for handling API routes
	Database persistence.SQLite3Database // SQLite database for storing delegations
}

// SetupApp initializes the app by setting up the database and the router.
func (a *App) SetupApp() error {
	var db persistence.SQLite3Database

	// Setup the database (SQLite) for storing delegations
	err := db.SetupDatabase()
	if err != nil {
		return err
	}
	a.Database = db

	var r router.Router

	// Setup the router and pass the database for data access
	err = r.SetupRouter(db)
	if err != nil {
		return err
	}
	a.Router = r

	return nil
}
