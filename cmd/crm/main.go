package main

import (
	"github.com/Jordy-6/CRM-Go/cmd"
	"github.com/Jordy-6/CRM-Go/internal/config"
	"github.com/Jordy-6/CRM-Go/internal/database"
	"github.com/Jordy-6/CRM-Go/internal/storage"
)

func main() {

	config.InitConfig()

	store := StoreSetup()

	cmd.SetStore(store)

	// Exécuter les commandes Cobra
	cmd.Execute()
}

func StoreSetup() storage.Storer {
	var store storage.Storer
	switch config.AppConfig.Storer.Type {
	case "gorm":
		database.ConnectDB()
		// Initialiser le store
		store = storage.NewGormStore(database.DB)
	case "json":
		store = storage.NewJsonStore()
	case "memory":
		store = storage.NewMemoryStore()
	default:
		panic("Type de stockage inconnu")

	}
	return store
}
