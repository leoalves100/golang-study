// Ref: https://github.com/larapulse/migrator
package migrations

import (
	"api/src/database"
	"log"
	"os"

	"github.com/larapulse/migrator"
)

var migrations = []migrator.Migration{
	{
		Name: "2023_03_20_create_users_table",
		Up: func() migrator.Schema {
			var s migrator.Schema
			users := migrator.Table{
				Name:   "users",
				Engine: "INNODB",
			}

			users.ID("id")
			users.Varchar("name", 50)
			users.Varchar("nick", 50)
			users.Varchar("mail", 50)
			users.Varchar("password", 100)
			users.Unique("nick", "mail")
			users.Timestamps()

			s.CreateTable(users)

			return s
		},
		Down: func() migrator.Schema {
			var s migrator.Schema

			s.DropTableIfExists("usuarios")

			return s
		},
	},
}

func Migration() {
	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	m := migrator.Migrator{Pool: migrations}

	migrated, err := m.Migrate(db)
	if err != nil {
		log.Printf("Could not migrate: %v", err)
		os.Exit(1)
	}

	if len(migrated) == 0 {
		log.Print("Nothing were migrated.")
	}

	for _, m := range migrated {
		log.Printf("Migration: %s was migrated ✅", m)
	}
}

func Revert() {
	db, _ := database.Connection()

	m := migrator.Migrator{Pool: migrations}

	revert, err := m.Revert(db)
	if err != nil {
		log.Printf("Could not revert migrations: %v", err)
		os.Exit(1)
	}

	if len(revert) == 0 {
		log.Print("Nothing were revert.")
	}

	for _, m := range revert {
		log.Printf("Migration: %s was revert ✅", m)
	}
}
