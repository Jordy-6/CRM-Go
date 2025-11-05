package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var displayMenuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Display the CRM menu",
	Long:  `menu is a command that displays the CRM system menu for user interaction.`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			showMenu()
		}
	},
}

func showMenu() {
	fmt.Println("\n========================================")
	fmt.Println("Welcome to the CRM System")
	fmt.Println("========================================")
	fmt.Println("1. Add Contacts")
	fmt.Println("2. List Contacts")
	fmt.Println("3. Delete Contact")
	fmt.Println("4. Update Contact")
	fmt.Println("5. Exit")
	fmt.Println("========================================")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	choiceInt, _ := strconv.Atoi(choice)

	switch choiceInt {
	case 1:
		addCmd.Run(addCmd, []string{})
	case 2:
		getCmd.Run(getCmd, []string{})
	case 3:
		deleteCmd.Run(deleteCmd, []string{})
	case 4:
		updateCmd.Run(updateCmd, []string{})
	case 5:
		fmt.Println("👋 Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("❌ Invalid choice, please try again")
	}
}

func init() {
	rootCmd.AddCommand(displayMenuCmd)
}
