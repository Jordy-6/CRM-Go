package main

import (
	"github.com/Jordy-6/CRM-Go/cmd"
	"github.com/Jordy-6/CRM-Go/internal/storage"
)

func main() {
	// Initialiser le store
	store := storage.NewJsonStore()

	// Configurer le store pour Cobra
	cmd.SetStore(store)

	// Exécuter les commandes Cobra
	cmd.Execute()
}
