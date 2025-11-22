package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonStore struct {
	contacts map[int]*Contact
	nextID   int
}

func NewJsonStore() *JsonStore {
	store := &JsonStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		data, _ := json.MarshalIndent(store.contacts, "", "  ")
		os.WriteFile(path, data, 0644)
	}
	return store
}

var path = "contacts.json"

func (m *JsonStore) Add(contact *Contact) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("impossible de lire le fichier %s: %w", path, err)
	}
	var contacts map[int]*Contact
	if len(data) > 0 {
		if err := json.Unmarshal(data, &contacts); err != nil {
			contacts = make(map[int]*Contact)
		}
	} else {
		contacts = make(map[int]*Contact)
	}

	contact.ID = m.nextID
	m.nextID++

	contacts[contact.ID] = contact
	m.contacts = contacts

	data, err = json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de sérialiser les données: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s: %w", path, err)
	}

	return nil
}

func (m *JsonStore) Update(id int, contact *Contact) error {

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("impossible de lire le fichier %s: %w", path, err)
	}

	var contacts map[int]*Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return fmt.Errorf("impossible de parser le fichier %s: %w", path, err)
	}

	if _, exists := contacts[id]; !exists {
		return fmt.Errorf("utilisateur introuvable")
	}

	contacts[id] = contact

	data, err = json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de sérialiser les données: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s: %w", path, err)
	}

	return nil
}

func (m *JsonStore) Delete(id int) error {

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("impossible de lire le fichier %s: %w", path, err)
	}

	var contacts map[int]*Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return fmt.Errorf("impossible de parser le fichier %s: %w", path, err)
	}

	if _, exists := contacts[id]; !exists {
		return fmt.Errorf("utilisateur introuvable")
	}

	delete(contacts, id)

	data, err = json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de sérialiser les données: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s: %w", path, err)
	}

	return nil
}

func (m *JsonStore) GetAll() error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("impossible de lire le fichier %s: %w", path, err)
	}

	var contacts map[int]*Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return fmt.Errorf("impossible de parser le fichier %s: %w", path, err)
	}

	if len(contacts) == 0 {
		return fmt.Errorf("no contacts")
	}

	for _, c := range contacts {
		fmt.Printf("[%d] %s - %s\n", c.ID, c.Name, c.Email)
	}
	return nil
}
