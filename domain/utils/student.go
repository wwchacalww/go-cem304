package utils

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type InputCheckStudentInClass struct {
	Educar      int64  `json:"ieducar"`
	ClassroomID string `json:"classroom_id"`
}

type ResultCheckStudentsList struct {
	Classroom string `json:"classroom"`
	Message   string `json:"message"`
}

type NameEducaDF struct {
	Name    string `json:"name"`
	EducaDF string `json:"educa_df"`
}

func CsvToStudents(f multipart.File) ([]repository.StudentInput, error) {
	var list []repository.StudentInput
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		return []repository.StudentInput{}, err
	}
	for _, line := range data {
		if len(line) == 6 {
			var rec repository.StudentInput
			for j, field := range line {
				switch {
				case j == 0:
					rec.Educar, _ = strconv.ParseInt(field, 10, 64)
				case j == 1:
					rec.Name = field
				case j == 2:
					rec.BirthDay = field
				case j == 3:
					rec.Gender = field
				case j == 4:
					rec.ANNE = field
				case j == 5:
					rec.ClassroomID = field
				}
			}
			list = append(list, rec)
		}
	}

	return list, nil
}

func StudentsInClass(f multipart.File, classroom_id string) ([]InputCheckStudentInClass, error) {
	var list []InputCheckStudentInClass
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		return []InputCheckStudentInClass{}, err
	}
	for _, line := range data {
		if len(line) == 1 {
			var rec InputCheckStudentInClass
			for j, field := range line {
				switch {
				case j == 0:
					rec.Educar, _ = strconv.ParseInt(field, 10, 64)
				}
			}
			rec.ClassroomID = classroom_id
			list = append(list, rec)
		}
	}

	return list, nil
}

func StudentsInClassrooms(f multipart.File) ([]InputCheckStudentInClass, error) {
	var list []InputCheckStudentInClass
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		return []InputCheckStudentInClass{}, err
	}
	for _, line := range data {
		if len(line) == 2 {
			var rec InputCheckStudentInClass
			for j, field := range line {
				switch {
				case j == 0:
					rec.ClassroomID = field
				case j == 1:
					rec.Educar, _ = strconv.ParseInt(field, 10, 64)
				}
			}
			list = append(list, rec)
		}
	}

	return list, nil
}

func CheckStudentsList(f multipart.File) ([]ResultCheckStudentsList, error) {
	var result []ResultCheckStudentsList
	reader := csv.NewReader(f)
	reader.Comma = '£'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var input ResultCheckStudentsList
	input.Classroom = "1A"
	for _, l := range data {
		strSplit := strings.Split(l[0], " ")
		if len(strSplit) == 1 {
			result = append(result, input)
			input.Classroom = l[0]
			input.Message = "Turma registrada \n\r"
		}
		if len(strSplit) > 1 {
			ieducarStr := strings.Replace(strSplit[0], " ", "", -1)
			input.Message = input.Message + input.Classroom + ";" + ieducarStr + "\n\r"
		}
	}
	return result, nil
}

func CheckEducaDFList(f multipart.File) ([]NameEducaDF, error) {
	var result []NameEducaDF
	reader := csv.NewReader(f)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range data {
		if len(line) == 7 {
			var input NameEducaDF
			input.Name = line[1]
			input.EducaDF = line[2] + "-" + line[3] + "-" + line[4]
			result = append(result, input)
		}
	}

	return result, nil
}

func FindByName(name string, students []model.StudentInterface) (model.StudentInterface, error) {
	for _, std := range students {
		name = strings.ToUpper(name)
		name = strings.TrimPrefix(name, " ")
		name = strings.TrimSuffix(name, " ")
		std_name := strings.ToUpper(std.GetName())
		std_name = strings.TrimPrefix(std_name, " ")
		std_name = strings.TrimSuffix(std_name, " ")
		if name == std_name {
			return std, nil
		}
	}

	return nil, fmt.Errorf(name + " não encontrado")
}
