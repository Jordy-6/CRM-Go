package main

import (
	"github.com/Jordy-6/CRM-Go/cmd"
	"github.com/Jordy-6/CRM-Go/internal/database"
	"github.com/Jordy-6/CRM-Go/internal/storage"
)

func main() {
	database.ConnectDB()

	// Initialiser le store
	store := storage.NewGormStore(database.DB)

	// Configurer le store pour Cobra
	cmd.SetStore(store)

	// Exécuter les commandes Cobra
	cmd.Execute()
}
