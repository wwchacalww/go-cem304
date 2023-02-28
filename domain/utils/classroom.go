package utils

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type CheckResult struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func CsvToClassrooms(f multipart.File) ([]repository.ClassroomInput, error) {
	var list []repository.ClassroomInput
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return []repository.ClassroomInput{}, err
	}

	for _, line := range data {
		if len(line) == 6 {
			var rec repository.ClassroomInput
			for j, field := range line {
				switch {
				case j == 0:
					rec.Name = field
				case j == 1:
					rec.Level = field
				case j == 2:
					rec.Grade = field
				case j == 3:
					rec.Shift = field
				case j == 4:
					rec.ANNE = field
				case j == 5:
					rec.Year = field
				}
			}
			list = append(list, rec)
		}
	}

	return list, nil
}

func FindClassById(classrooms []model.Classroom, classroom_id string) (model.Classroom, error) {
	for _, class := range classrooms {
		if class.GetID() == classroom_id {
			return class, nil
		}
	}
	return model.Classroom{}, fmt.Errorf("Classroom not found")
}

func CheckStudentInClassrooms(list []InputCheckStudentInClass, ieducar int64, classroom_id string) (CheckResult, error) {
	var result CheckResult
	if len(list) == 0 {
		return result, fmt.Errorf("list invalid")
	}

	result.Result = false
	result.Message = strconv.FormatInt(ieducar, 10) + "<->" + classroom_id + " NÃO ENCONTRADO"
	for _, std := range list {
		if std.Educar == ieducar {
			if std.ClassroomID == classroom_id {
				result.Result = true
				result.Message = "OK"
				return result, nil
			}
			result.Result = false
			result.Message = strconv.FormatInt(std.Educar, 10) + " <-> " + std.ClassroomID + " != " + classroom_id
			return result, nil
		}
	}

	return result, nil
}
