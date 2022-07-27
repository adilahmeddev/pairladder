package models

import "github.com/google/uuid"

type Person struct {
	ID uuid.UUID
}

func NewPerson(name string) Person {
	return Person{ID: uuid.New()}
}
