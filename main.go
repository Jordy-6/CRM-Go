package main

// Display home page
// Leave home page

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Jordy-6/CRM-Go/contact"
)

func main() {

	addFlag := flag.Bool("add", false, "Ajouter un contact via flags")
	nameFlag := flag.String("name", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *addFlag {
		contact.AddContactWithFlag(*nameFlag, *emailFlag)
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

		contact.AddContact()
	case 2:
		contact.GetContact()
	case 3:
		contact.DeleteContact()
	case 4:
		contact.UpdateContact()
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}
