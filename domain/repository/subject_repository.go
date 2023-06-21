package repository

import "wwchacalww/go-cem304/domain/model"

type SubjectInput struct {
	Name     string `valid:"required,stringlength(5|150)" json:"name"`
	License  string `valid:"required,stringlength(5|50)" json:"license"`
	Level    string `valid:"required,stringlength(5|50)" json:"level"`
	Grade    string `valid:"required,stringlength(5|50)" json:"grade"`
	Note     string `valid:"optional" json:"note"`
	CH       int    `valid:"optional" json:"ch"`
	Year     int    `valid:"required" json:"year"`
	Semester string `valid:"optional" json:"semester"`
}

type SubjectRepositoryInterface interface {
	Create(input SubjectInput) (model.SubjectInterface, error)
	FindById(id string) (model.SubjectInterface, error)
	FindByLicense(license string) ([]model.SubjectInterface, error)
	FindByName(name string) ([]model.SubjectInterface, error)
}

type SubjectPersistence interface {
	Create(input SubjectInput) (model.SubjectInterface, error)
	FindById(id string) (model.SubjectInterface, error)
	FindByLicense(license string) ([]model.SubjectInterface, error)
	FindByName(name string) ([]model.SubjectInterface, error)
}
