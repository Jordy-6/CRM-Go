package contact

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Jordy-6/CRM-Go/internal/storage"
)

// Api call for contact

// Add
func AddContact(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	if name == "" || email == "" {
		fmt.Println("❌ Name and email are required")
		return
	}

	newContact := &storage.Contact{
		Name:  name,
		Email: email,
	}

	error := store.Add(newContact)
	if error != nil {
		fmt.Println("❌ Error adding contact:", error)
		return
	}

	fmt.Println("✅ Contact added successfully.")
}

func AddContactWithFlag(store storage.Storer, name string, email string) {
	if name == "" || email == "" {
		fmt.Println("❌ Name and email are required")
		return
	}

	newContact := &storage.Contact{
		Name:  name,
		Email: email,
	}
	store.Add(newContact)
	fmt.Println("✅ Contact added successfully.")
}

// Delete
func DeleteContact(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Id contact to delete : ")

	input, _ := reader.ReadString('\n')
	idContact, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("❌ ID invalide.")
		return
	}

	err = store.Delete(idContact)
	if err != nil {
		fmt.Println("⚠️", err)
	} else {
		fmt.Println("🗑️ Contact deleted.")
	}
}

// Update
func UpdateContact(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Id to update : ")

	input, _ := reader.ReadString('\n')
	idContact, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("❌ ID invalide.")
		return
	}

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter new email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	newContact := &storage.Contact{
		ID:    idContact,
		Name:  name,
		Email: email,
	}

	err = store.Update(idContact, newContact)
	if err != nil {
		fmt.Println("⚠️", err)
	} else {
		fmt.Println("✅ Contact updated.")
	}
}

// List
func GetContact(store storage.Storer) {
	err := store.GetAll()
	if err != nil {
		fmt.Println("⚠️", err)
	}
}
