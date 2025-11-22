# 🧩 Mini CRM – Command Line Project (Go)

## 🎯 Project Overview

This project is a **mini CRM (Customer Relationship Management)** built with **Go (Golang)**.
It allows you to **manage a list of contacts** directly from the command line.

Each contact includes:

* a unique **ID** (auto-incremented),
* a **name**,
* an **email address**.

---

## ⚙️ Main Features

### 🧭 Interactive Mode (Menu)

When you run:

```bash
go run .
```

The application launches an **interactive main menu** that lets you:

1. Add a contact
2. List all contacts
3. Delete a contact
4. Update a contact
5. Exit the application

### 🚀 Command-Line Flags Mode

You can also **add a contact directly** using flags — no menu needed:

```bash
go run . -add -name="Jordy" -email="jordy@mail.com"
```

➡️ Output:

```
✅ Contact added: [1] Ali - ali@mail.com
```

---
## 🏗️ Project Structure

```
CRM/
├── go.mod
├── main.go
└── contact/
    └── contact.go
```

### 📁 `contact/contact.go`

Contains the `Contact` struct and related CRUD functions.

### 📁 `main.go`

Entry point of the program — handles **flags** and the **interactive menu**.

---

## 🚀 How to Run

### Run in interactive mode

```bash
go run .
```

### Run with flags

Add a contact directly:

```bash
go run . -add -name="Jordy" -email="jordy@mail.com"
```

## 👨‍💻 Authors

Jordy PEREIRA-ELENGA MAKOUALA Developer

Romain MONMARCHE Developer




Pour ajouter viper dans cobra => command root + persistant prerun