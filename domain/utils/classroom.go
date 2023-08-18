package utils

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"

	uuid "github.com/satori/go.uuid"
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

func FindClassBySlug(classrooms []model.ClassroomInterface, slug string) (model.ClassroomInterface, error) {
	if len(slug) != 2 {
		return nil, fmt.Errorf("Classroom name invalid")
	}

	afternoon := string(slug[0]) + "º ano " + string(slug[1]) + " - Vespertino"
	morning := string(slug[0]) + "º ano " + string(slug[1]) + " - Matutino"
	for _, class := range classrooms {
		if class.GetName() == morning || class.GetName() == afternoon {
			return class, nil
		}
	}
	return nil, fmt.Errorf("Classroom not found. " + afternoon + " and " + morning)
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

func ReportToStudents(f multipart.File, classroom_id string) ([]repository.StudentInput, error) {
	var result []repository.StudentInput

	reader := csv.NewReader(f)
	reader.Comma = '£'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var input repository.StudentInput
	for i, l := range data {
		switch {
		case i%10 == 0:
			ieducarStr := strings.Replace(l[0], " ", "", -1)
			n, err := strconv.ParseInt(ieducarStr, 10, 64)
			if err != nil {
				return nil, err
			}
			input.Educar = n
			input.Status = true
			input.ClassroomID = classroom_id
		case i%10 == 1:
			name, birth_date, err := ExtractNameAndBirthday(l[0])
			if err != nil {
				return nil, err
			}
			input.Name = name
			input.BirthDay = birth_date
		case i%10 == 2:
			if len(l[0]) > 8 {
				input.Father = l[0]
			}
		case i%10 == 3:
			if len(l[0]) > 8 {
				input.Mother = l[0]
			}
		case i%10 == 4:
			if len(l[0]) > 8 {
				input.Responsible = l[0]
			} else {
				input.Responsible = input.Mother
			}
		case i%10 == 5:
			input.Address = l[0]
		case i%10 == 6:
			input.City = l[0]
		case i%10 == 7:
			input.CEP = l[0]
		case i%10 == 8:
			input.Fones = l[0]
		case i%10 == 9:
			input.CPF = uuid.NewV4().String()
			if l[0] != "NULL" {
				input.CPF = l[0]
			}
			result = append(result, input)
		}
	}
	return result, nil
}

func ExtractNameAndBirthday(str string) (string, string, error) {
	name := ""
	birth_day := ""
	strSplit := strings.Split(str, " ")
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	for _, txt := range strSplit {
		if !re.MatchString(txt) {
			name = name + " " + txt
		} else {
			dt := strings.Split(txt, "/")
			if len(dt) != 3 {
				return "", "", fmt.Errorf("Invalid date")
			}
			birth_day = txt
		}
	}
	return name, birth_day, nil
}
