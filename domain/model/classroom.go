package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ClassroomInterface interface {
	IsValid() (bool, error)
	GetID() string
	GetName() string
	GetLevel() string
	GetGrade() string
	GetShift() string
	GetDescription() string
	GetANNE() string
	GetYear() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetStatus() bool
	GetStudents() []StudentInterface
	Enable() error
	Disable() error
}

type Classroom struct {
	ID          string             `valid:"uuidv4" json:"id"`
	Name        string             `valid:"required,stringlength(5|50)" json:"Name"`
	Level       string             `valid:"required" json:"level"`
	Grade       string             `valid:"required" json:"grade"`
	Shift       string             `valid:"required" json:"shift"`
	Description string             `valid:"optional" json:"description"`
	ANNE        string             `valid:"optional" json:"ANNE"`
	Year        string             `valid:"required" json:"year"`
	Status      bool               `valid:"optional" json:"status"`
	Students    []StudentInterface `valid:"optional" json:"students"`
	CreatedAt   time.Time          `valid:"optional" json:"created_at"`
	UpdatedAt   time.Time          `valid:"optional" json:"updated_at"`
}

func NewClassroom() *Classroom {
	classroom := Classroom{
		ID:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    true,
	}

	return &classroom
}

func (c *Classroom) IsValid() (bool, error) {
	if c.Status != false && c.Status != true {
		c.Status = false
	}

	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Classroom) GetID() string {
	return c.ID
}

func (c *Classroom) GetName() string {
	return c.Name
}

func (c *Classroom) GetLevel() string {
	return c.Level
}

func (c *Classroom) GetGrade() string {
	return c.Grade
}

func (c *Classroom) GetShift() string {
	return c.Shift
}

func (c *Classroom) GetDescription() string {
	return c.Description
}

func (c *Classroom) GetANNE() string {
	return c.ANNE
}

func (c *Classroom) GetYear() string {
	return c.Year
}

func (c *Classroom) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *Classroom) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

func (c *Classroom) GetStatus() bool {
	return c.Status
}

func (c *Classroom) GetStudents() []StudentInterface {
	return c.Students
}

func (c *Classroom) Enable() error {
	c.Status = true
	return nil
}

func (c *Classroom) Disable() error {
	c.Status = false
	return nil
}
