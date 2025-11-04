package storage

import "fmt"

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (m *MemoryStore) Add(contact *Contact) error {
	contact.ID = m.nextID
	m.contacts[m.nextID] = contact
	m.nextID++
	return nil
}

func (m *MemoryStore) Update(id int, contact *Contact) error {

	if _, exists := m.contacts[id]; !exists {
		return fmt.Errorf("Invalid User")
	}

	m.contacts[id] = contact

	return nil
}

func (m *MemoryStore) Delete(id int) error {

	if _, exists := m.contacts[id]; !exists {
		return fmt.Errorf("Invalid User")
	}

	delete(m.contacts, id)
	return nil

}

func (m *MemoryStore) GetAll() error {
	if len(m.contacts) == 0 {
		return fmt.Errorf("No contacts")
	}

	for _, c := range m.contacts {
		fmt.Printf("[%d] %s - %s\n", c.ID, c.Name, c.Email)
	}

	return nil
}
