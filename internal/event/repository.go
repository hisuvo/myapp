package event

import (
	"errors"

	"gorm.io/gorm"
)

var ErrEventNotFound = errors.New("event not found")

type Repository interface {
	Create(event *Event) error
	GetAll()(*[]Event, error)
	GetById(eventId uint) (*Event, error)
	Update(event *Event) error
}

type repository struct {
	db *gorm.DB
}

// Any function named NewXxx() is considered a constructor.
func NewRepository(db *gorm.DB) Repository{
	return &repository{db: db}
}

func (r *repository) Create(event *Event) error {
	// Todo: we can do this as details
	// result := r.db.Create(event)
	// if result.Error != nil {
	// 	return result.Error
	// }
	// return nil

	// Todo: short way
	return r.db.Create(event).Error
}

func (r *repository) GetAll()(*[]Event,error){
	var events *[]Event

	err := r.db.Find(&events).Error;

	if err != nil {
		return nil, ErrEventNotFound
	}

	return events, nil
}

func (r *repository) GetById(eventId uint) (*Event, error) {
	var event Event
	err := r.db.First(&event, eventId).Error

	if err != nil {
		if errors.Is(err, ErrEventNotFound){
			return nil, ErrEventNotFound
		}
		return nil, err
	}
	return &event, nil
}

func (r *repository) Update(event *Event) error {
	return r.db.Save(event).Error
}

// ---------------------------------Common Golang Design pattern-------------------------------------- //

// Todo: 
type Animal interface {
	Sound() error
}

type Dog struct {
	Name string
	Breed string
}

func NewDog(name string, bread string) Animal{
	return &Dog{
		Name: name,
		Breed: bread,
	}
}


func (d *Dog) Sound()error{
	return nil
}
