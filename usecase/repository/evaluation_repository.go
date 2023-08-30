package repository

import (
	"wwchacalww/go-cem304/domain/dtos"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type EvaluationRepository struct {
	Persistence repository.EvaluationPersistence
}

func NewEvaluationRepository(persistence repository.EvaluationPersistence) *EvaluationRepository {
	return &EvaluationRepository{Persistence: persistence}
}

func (repo *EvaluationRepository) Create(input repository.EvaluationInput) (model.EvaluationInterface, error) {
	err := repo.Persistence.Create(input)
	if err != nil {
		return nil, err
	}

	eval, err := repo.Persistence.FindEvaluation(input.Student, input.Term, input.Subject)
	if err != nil {
		return nil, err
	}
	return eval, nil
}

func (repo *EvaluationRepository) Update(note string, id, absences int) (model.EvaluationInterface, error) {
	err := repo.Persistence.Update(note, id, absences)
	if err != nil {
		return nil, err
	}
	eval, err := repo.Persistence.FindById(id)
	if err != nil {
		return nil, err
	}
	return eval, nil
}

func (repo *EvaluationRepository) ImportCSVEvaluation(input dtos.InputImportCard) error {
	var inputs []repository.EvaluationInput

	for _, std := range input.Cards {
		var ip repository.EvaluationInput
		ip.Term = input.Term
		student_id, err := repo.Persistence.GetAndCheckStudent(input.ClassroomID, std.EducaDF)
		if err != nil {
			return err
		}
		ip.Student = student_id
		for _, sbj := range std.Subjects {
			id, err := repo.Persistence.GetSubjectId(input.ClassroomID, sbj.Name)
			if err != nil {
				return err
			}
			ip.Subject = id
			ip.Note = sbj.Note
			ip.Absences = sbj.Absences
			inputs = append(inputs, ip)
		}
	}

	for _, reg := range inputs {
		check, err := repo.Persistence.CheckEvaluationExists(reg.Student, reg.Term, reg.Subject)
		if err != nil {
			return err
		}
		if check != 0 {
			err = repo.Persistence.Update(reg.Note, check, reg.Absences)
			if err != nil {
				return err
			}
		} else {
			err = repo.Persistence.Create(reg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
