package contact

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Api call for contact

// Contact type
type Contact struct {
	ID    int
	Name  string
	Email string
}

// var contacts []Contact
var contacts = make(map[int]Contact)
var contactIDCounter int = 1

// Add
func AddContact() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')

	contacts[contactIDCounter] = Contact{ID: contactIDCounter, Name: name, Email: email}

	contactIDCounter++

	return "✅ Contact added successfully!"
}

func AddContactWithFlag(name string, email string) {
	if name == "" || email == "" {
		fmt.Println("⚠️ Nom et email requis.")
		return
	}

	contacts[contactIDCounter] = Contact{
		ID:    contactIDCounter,
		Name:  name,
		Email: email,
	}
	fmt.Printf("✅ Contact ajouté : [%d] %s - %s\n", contactIDCounter, name, email)
	contactIDCounter++
}

// Delete
func DeleteContact() {

	if len(contacts) == 0 {
		fmt.Println("No contact.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Id contact to delete : ")

	input, _ := reader.ReadString('\n')
	idContact, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("❌ ID invalide.")
		return
	}

	if _, ok := contacts[idContact]; ok {
		delete(contacts, idContact)
		fmt.Println("🗑️ Contact deleted.")
	} else {
		fmt.Println("⚠️ Contact not found.")
	}
}

// Update
func UpdateContact() {

	if len(contacts) == 0 {
		fmt.Println("No contact.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Id to update : ")

	input, _ := reader.ReadString('\n')
	idContact, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("❌ ID invalide.")
		return
	}

	if _, ok := contacts[idContact]; ok {
		fmt.Print("Enter new name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("Enter new email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		contacts[idContact] = Contact{ID: idContact, Name: name, Email: email}
		fmt.Println("✅ Contact updated.")
	} else {
		fmt.Println("⚠️ Contact not found.")
	}

}

// List
func GetContact() {
	if len(contacts) == 0 {
		fmt.Println("No contact.")
		return
	}
	fmt.Println("📋 Liste des contacts :")
	for _, c := range contacts {
		fmt.Printf("[%d] %s - %s\n", c.ID, c.Name, c.Email)
	}
}
