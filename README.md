# 🧩 Mini CRM – Command Line Project (Go)

## 🎯 Project Overview

This project is a **mini CRM (Customer Relationship Management)** built with **Go (Golang)** using **Cobra CLI** and **Viper** for configuration management.
It allows you to **manage contacts** from the command line with multiple storage backends.

Each contact includes:
* a unique **ID** (auto-incremented)
* a **name**
* an **email address**

---

## ⚙️ Features

### 🔧 Multiple Storage Backends
- **Memory**: In-memory storage (data lost on exit)
- **JSON**: File-based storage in `contacts.json`
- **GORM**: SQLite database storage

### 🚀 Cobra CLI Commands
```bash
# Add a contact (interactive)
go run .\cmd\crm add

# List all contacts
go run .\cmd\crm list

# Delete a contact (interactive)
go run .\cmd\crm delete

# Update a contact (interactive)
go run .\cmd\crm update

# Show help
go run .\cmd\crm --help
```

### ⚙️ Configuration via YAML
Storage type is configurable via `config/config.yaml`:

```yaml
database:
  name: CRM.db

storer:
  type: json  # Options: memory, json, gorm
```

---

## 🏗️ Project Structure

```
CRM/
├── go.mod
├── go.sum
├── README.md
├── config/
│   └── config.yaml          # Configuration file
├── cmd/
│   ├── root.go             # Cobra root command
│   ├── contact.go          # Contact-related commands
│   └── crm/
│       └── main.go         # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go       # Viper configuration loading
│   ├── database/
│   │   └── db.go          # GORM database connection
│   └── storage/
│       ├── storage.go      # Storage interface
│       ├── memory.go       # In-memory implementation
│       ├── json.go         # JSON file implementation
│       └── gorm.go         # GORM database implementation
└── contacts.json           # JSON storage file (auto-created)
```

---

## 📦 Installation & Setup

### Prerequisites
- **Go 1.25+** installed
- **Windows** (PowerShell or CMD)

### 1. Clone & Install Dependencies
```bash
git clone <repository-url>
cd CRM
go mod tidy
```

### 2. Configuration
Edit `config/config.yaml` to choose your storage backend:

```yaml
database:
  name: CRM.db

storer:
  type: json  # Change to: memory, json, or gorm
```

### 3. Build (Optional)
```bash
go build -o crm.exe .\cmd\crm
```

---

## 🚀 Usage Examples

### Command Line Interface

```bash
# Add a new contact
PS> go run .\cmd\crm add
Enter name: Alice
Enter email: alice@example.com
✅ Contact added successfully.

# List all contacts
PS> go run .\cmd\crm list
📋 Contacts:
[1] Alice - alice@example.com
[2] Bob - bob@test.com

# Delete a contact
PS> go run .\cmd\crm delete
Id contact to delete : 1
🗑️ Contact deleted.

# Update a contact
PS> go run .\cmd\crm update
Id to update : 2
Enter new name: Robert
Enter new email: robert@example.com
✅ Contact updated.
```

### Using Built Executable

```bash
# After building
.\crm.exe add
.\crm.exe list
.\crm.exe delete
.\crm.exe update
```

## 🔧 Storage Backends

### 🧠 Memory Storage
- **Type**: `memory`
- **Persistence**: None (data lost on exit)
- **Use case**: Testing, temporary data

### 📄 JSON Storage
- **Type**: `json`
- **File**: `contacts.json`
- **Persistence**: File-based
- **Use case**: Simple file storage

### 🗄️ Database Storage (GORM)
- **Type**: `gorm`
- **Database**: SQLite (`CRM.db`)
- **Persistence**: Database
- **Use case**: Production, advanced features

---

## 🛠️ Technology Stack

- **Language**: Go 1.25+
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Database ORM**: [GORM](https://gorm.io/)
- **SQLite Driver**: [glebarez/sqlite](https://github.com/glebarez/go-sqlite)

---

## 👨‍💻 Authors

- **Jordy PEREIRA-ELENGA MAKOUALA** - Developer
- **Romain MONMARCHE** - Developer
