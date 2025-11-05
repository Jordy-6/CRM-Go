package main

// Display home page
// Leave home page

import (
	"github.com/Jordy-6/CRM-Go/internal/app"
	"github.com/Jordy-6/CRM-Go/internal/storage"
)

func main() {
	var store = storage.NewMemoryStore()
	app.Run(store)
}
