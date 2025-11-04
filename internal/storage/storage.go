package storage

type Storer interface {
	Add(message string) error
	Update(id int, newMessage string) error
	Delete(id int) error
	GetAll() ([]string, error)
}

type Contact struct {
	ID    int
	Name  string
	Email string
}
