package repository

import (
	"time"
	"wwchacalww/go-cem304/domain/model"
)

type TeacherInput struct {
	Name     string `json:"name" valid:"required,stringlength(5|20)"`
	Nick     string `json:"nick" valid:"alphanum,required"`
	CPF      string `json:"cpf" valid:"optional"`
	Email    string `json:"email" valid:"optional,email"`
	Fones    string `json:"fones" valid:"optional"`
	License  string `json:"license" valid:"optional"`
	Gender   string `json:"gender" valid:"optional"`
	BirthDay string `json:"birth_day" valid:"optional"`
	Note     string `json:"note" valid:"optional"`
}

type TeacherRepositoryInterface interface {
	Create(input TeacherInput) (model.TeacherInterface, error)
	FindById(id string) (model.TeacherInterface, error)
	FindByCPF(cpf string) (model.TeacherInterface, error)
	FindByName(name string) ([]model.TeacherInterface, error)
	AttachClassroomSubject(id, classroom_id, subject_id, slug, start_course, end_course string, wch int32) (model.TeacherInterface, error)
	Save(id string, input TeacherInput) (model.TeacherInterface, error)
}

type TeacherPersistence interface {
	Create(input TeacherInput) (model.TeacherInterface, error)
	FindById(id string) (model.TeacherInterface, error)
	FindByName(name string) ([]model.TeacherInterface, error)
	FindByCPF(cpf string) (model.TeacherInterface, error)
	AttachClassroomSubject(id, classroom_id, subject_id, slug string, wch int32, start_course, end_course time.Time) error
	Save(teacher model.TeacherInterface) error
}
