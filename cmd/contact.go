package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Jordy-6/CRM-Go/internal/storage"
	"github.com/spf13/cobra"
)

var store storage.Storer

func SetStore(s storage.Storer) {
	store = s
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact to the CRM",
	Long:  `add is a command that allows you to add a new contact to the CRM system.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a contact from the CRM",
	Long:  `delete is a command that allows you to delete a contact from the CRM system.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a contact in the CRM",
	Long:  `update is a command that allows you to update a contact in the CRM system.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

var getCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contacts in the CRM",
	Long:  `list is a command that retrieves and displays all contacts stored in the CRM system.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := store.GetAll()
		if err != nil {
			fmt.Println("⚠️", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(getCmd)
}
