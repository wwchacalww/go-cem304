package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type TeacherInterface interface {
	IsValid() (bool, error)
	GetID() string
	GetName() string
	GetNick() string
	GetBirthDay() time.Time
	GetGender() string
	GetCPF() string
	GetFones() string
	GetEmail() string
	GetLicense() string
	GetNote() string
	GetStatus() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Save(email, bd, gender, fones, license, note string) error
}

type Teacher struct {
	ID        string    `valid:"uuidv4" json:"id"`
	Name      string    `valid:"required,stringlength(5|150)" json:"name"`
	Nick      string    `valid:"required,stringlength(5|50)" json:"nick"`
	BirthDay  time.Time `valid:"optional" json:"birth_day"`
	Gender    string    `valid:"optional" json:"gender"`
	CPF       string    `valid:"optional" json:"cpf"`
	Fones     string    `valid:"optional" json:"fones"`
	Email     string    `valid:"email,optional" json:"email"`
	License   string    `valid:"optional" json:"license"`
	Note      string    `valid:"optional" json:"note"`
	Status    bool      `valid:"optional" json:"status"`
	CreatedAt time.Time `valid:"optional" json:"created_at"`
	UpdatedAt time.Time `valid:"optional" json:"updated_at"`
}

func NewTeacher(name, nick, cpf, license string) *Teacher {
	Teacher := Teacher{
		ID:        uuid.NewV4().String(),
		Name:      name,
		Nick:      nick,
		CPF:       cpf,
		License:   license,
		Status:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &Teacher
}

func (t *Teacher) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *Teacher) GetID() string {
	return t.ID
}

func (t *Teacher) GetName() string {
	return t.Name
}

func (t *Teacher) GetNick() string {
	return t.Nick
}

func (t *Teacher) GetBirthDay() time.Time {
	return t.BirthDay
}

func (t *Teacher) GetGender() string {
	return t.Gender
}

func (t *Teacher) GetCPF() string {
	return t.CPF
}

func (t *Teacher) GetLicense() string {
	return t.License
}

func (t *Teacher) GetNote() string {
	return t.Note
}

func (t *Teacher) GetFones() string {
	return t.Fones
}

func (t *Teacher) GetEmail() string {
	return t.Email
}

func (t *Teacher) GetStatus() bool {
	return t.Status
}

func (t *Teacher) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *Teacher) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}

func (t *Teacher) Save(email, bd, gender, fones, license, note string) error {
	if email != "" {
		t.Email = email
	}
	if bd != "" {
		dt := strings.Split(bd, "/")
		if len(dt) != 3 {
			return fmt.Errorf("Birth day invalid")
		}
		y, err := strconv.Atoi(dt[2])
		if err != nil {
			return err
		}
		m, err := strconv.Atoi(dt[1])
		if err != nil {
			return err
		}
		d, err := strconv.Atoi(dt[0])

		if err != nil {
			return err
		}
		birth_day := time.Date(y, time.Month(m), d, 12, 15, 5, 5, time.Local)
		t.BirthDay = birth_day
	}
	if gender == "Masculino" || gender == "Feminino" {
		t.Gender = gender
	}
	if fones != "" {
		t.Fones = fones
	}
	if license != "" {
		t.License = license
	}
	if note != "" {
		t.Note = note
	}

	_, err := t.IsValid()
	if err != nil {
		return err
	}

	return nil
}
