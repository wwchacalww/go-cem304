package repository

import (
	"log"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"
)

type TeacherRepository struct {
	Persistence repository.TeacherPersistence
}

func NewTeacherRepository(persistence repository.TeacherPersistence) *TeacherRepository {
	return &TeacherRepository{Persistence: persistence}
}

func (repo *TeacherRepository) Create(input repository.TeacherInput) (model.TeacherInterface, error) {
	teacher, err := repo.Persistence.Create(input)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (repo *TeacherRepository) FindById(id string) (model.TeacherInterface, error) {
	Teacher, err := repo.Persistence.FindById(id)
	if err != nil {
		return &model.Teacher{}, err
	}

	return Teacher, nil
}

func (repo *TeacherRepository) FindByCPF(cpf string) (model.TeacherInterface, error) {
	Teacher, err := repo.Persistence.FindByCPF(cpf)
	if err != nil {
		return &model.Teacher{}, err
	}

	return Teacher, nil
}

func (repo *TeacherRepository) FindByName(name string) ([]model.TeacherInterface, error) {
	var result []model.TeacherInterface
	teachers, err := repo.Persistence.FindByName(name)
	if err != nil {
		return nil, err
	}

	if len(teachers) > 0 {
		for _, t := range teachers {
			teacher, err := repo.Persistence.FindById(t.GetID())
			if err != nil {
				return nil, err
			}
			result = append(result, teacher)
		}
	}

	return result, nil
}

func (repo *TeacherRepository) AttachClassroomSubject(id, classroom_id, subject_id, slug, start_course, end_course string, wch int32) (model.TeacherInterface, error) {
	st, err := utils.String2Time(start_course)
	if err != nil {
		log.Println("Deu ruim")
		return &model.Teacher{}, err
	}
	ec, err := utils.String2Time(end_course)
	if err != nil {
		return &model.Teacher{}, err
	}
	err = repo.Persistence.AttachClassroomSubject(id, classroom_id, subject_id, slug, wch, st, ec)
	if err != nil {
		return &model.Teacher{}, err
	}

	teacher, _ := repo.Persistence.FindById(id)
	return teacher, nil
}

func (repo *TeacherRepository) Save(id string, input repository.TeacherInput) (model.TeacherInterface, error) {
	teacher, err := repo.Persistence.FindById(id)
	if err != nil {
		return &model.Teacher{}, err
	}

	teacher.Save(input.Email, input.BirthDay, input.Gender, input.Fones, input.License, input.Note)
	err = repo.Persistence.Save(teacher)
	if err != nil {
		return &model.Teacher{}, err
	}
	return teacher, nil
}
