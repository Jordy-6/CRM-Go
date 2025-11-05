package app

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Jordy-6/CRM-Go/cmd"
	"github.com/Jordy-6/CRM-Go/contact"
	"github.com/Jordy-6/CRM-Go/internal/storage"
)

func Run(store storage.Storer) {

	cmd.SetStore(store)

	addFlag := flag.Bool("add", false, "Ajouter un contact via flags")
	nameFlag := flag.String("name", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *addFlag {
		contact.AddContactWithFlag(store, *nameFlag, *emailFlag)
	}

	for {
		homePage()
	}
}

func homePage() {
	fmt.Println("Welcome to the CRM System")
	fmt.Println("1. Add Contacts")
	fmt.Println("2. List Contacts")
	fmt.Println("3. Delete Contact")
	fmt.Println("4. Update Contact")
	fmt.Println("5. Exit")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	choiceInt, _ := strconv.Atoi(choice)

	switch choiceInt {
	case 1:
		os.Args = []string{"crm-go", "add"}
		cmd.Execute()
	case 2:
		os.Args = []string{"crm-go", "list"}
		cmd.Execute()
	case 3:
		os.Args = []string{"crm-go", "delete"}
		cmd.Execute()
	case 4:
		os.Args = []string{"crm-go", "update"}
		cmd.Execute()
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}
