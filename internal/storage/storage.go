package storage

type Storer interface {
	Add(contact *Contact) error
	Update(id int, contact *Contact) error
	Delete(id int) error
	GetAll() error
}

type Contact struct {
	ID    int
	Name  string
	Email string
}
