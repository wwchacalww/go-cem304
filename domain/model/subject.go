package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type SubjectInterface interface {
	IsValid() (bool, error)
	GetID() string
	GetName() string
	GetLicense() string
	GetLevel() string
	GetGrade() string
	GetNote() string
	GetCH() int
	GetYear() int
	GetSemester() string
	GetClassrooms() []ClassroomInterface
	GetTeacher() TeacherInterface
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	AttachClassroom(classroom ClassroomInterface) error
}

type Subject struct {
	ID         string               `valid:"uuidv4" json:"id"`
	Name       string               `valid:"required,stringlength(5|150)" json:"name"`
	License    string               `valid:"required,stringlength(5|50)" json:"license"`
	Level      string               `valid:"required,stringlength(5|50)" json:"level"`
	Grade      string               `valid:"required,stringlength(5|50)" json:"grade"`
	Note       string               `valid:"optional" json:"note"`
	CH         int                  `valid:"optional" json:"ch"`
	Year       int                  `valid:"required" json:"year"`
	Semester   string               `valid:"optional" json:"semester"`
	Classrooms []ClassroomInterface `valid:"optional" json:"classrooms"`
	Teacher    TeacherInterface     `valid:"optional" json:"teacher"`
	CreatedAt  time.Time            `valid:"optional" json:"created_at"`
	UpdatedAt  time.Time            `valid:"optional" json:"updated_at"`
}

func NewSubject(name, license, level string) *Subject {
	Subject := Subject{
		ID:        uuid.NewV4().String(),
		Name:      name,
		License:   license,
		Level:     level,
		Grade:     "Ensino Médio",
		Year:      time.Now().Year(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &Subject
}

func (s *Subject) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(s)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Subject) GetID() string {
	return s.ID
}

func (s *Subject) GetName() string {
	return s.Name
}

func (s *Subject) GetLicense() string {
	return s.License
}

func (s *Subject) GetLevel() string {
	return s.Level
}

func (s *Subject) GetGrade() string {
	return s.Grade
}

func (s *Subject) GetNote() string {
	return s.Note
}

func (s *Subject) GetCH() int {
	return s.CH
}

func (s *Subject) GetYear() int {
	return s.Year
}

func (s *Subject) GetSemester() string {
	return s.Semester
}

func (s *Subject) GetClassrooms() []ClassroomInterface {
	return s.Classrooms
}

func (s *Subject) GetTeacher() TeacherInterface {
	return s.Teacher
}

func (s *Subject) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *Subject) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *Subject) AttachClassroom(classroom ClassroomInterface) error {
	if s.Classrooms == nil {
		s.Classrooms = []ClassroomInterface{classroom}
		return nil
	}
	check := false
	for _, c := range s.Classrooms {
		if c.GetID() == classroom.GetID() {
			check = true
			break
		}
	}
	if !check {
		s.Classrooms = append(s.Classrooms, classroom)
	}
	return nil
}
