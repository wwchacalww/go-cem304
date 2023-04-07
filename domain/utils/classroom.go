package utils

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"time"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type CheckResult struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

type Parent struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
	Resposible   bool   `json:"responsible"`
}

type Input struct {
	Ieducar  int64     `json:"ieducar"`
	Name     string    `json:"name"`
	BirthDay time.Time `json:"birth_day"`
	Address  string    `json:"address"`
	City     string    `json:"city"`
	CEP      string    `json:"CEP"`
	Fones    string    `json:"fones"`
	CPF      string    `json:"CPF"`
	Parents  []Parent
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

func ReportToStudents(f multipart.File) ([]Input, error) {
	var mother Parent
	var father Parent
	var result []Input

	reader := csv.NewReader(f)
	reader.Comma = '£'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var input Input
	for i, l := range data {
		switch {
		case i%10 == 0:
			ieducarStr := strings.Replace(l[0], " ", "", -1)
			n, err := strconv.ParseInt(ieducarStr, 10, 64)
			if err != nil {
				return nil, err
			}
			input.Ieducar = n
		case i%10 == 1:
			name, birth_date, err := ExtractNameAndBirthday(l[0])
			if err != nil {
				return nil, err
			}
			input.Name = name
			input.BirthDay = birth_date
		case i%10 == 2:
			if len(l[0]) > 8 {
				father.Name = l[0]
				father.Relationship = "Pai"
				father.Resposible = false
			}
		case i%10 == 3:
			if len(l[0]) > 8 {
				mother.Name = l[0]
				mother.Relationship = "Mãe"
				mother.Resposible = false
			}
		case i%10 == 4:
			if l[0] == father.Name {
				father.Resposible = true
			} else {
				mother.Resposible = true
			}
			input.Parents = []Parent{mother, father}
		case i%10 == 5:
			input.Address = l[0]
		case i%10 == 6:
			input.City = l[0]
		case i%10 == 7:
			input.CEP = l[0]
		case i%10 == 8:
			input.Fones = l[0]
		case i%10 == 9:
			input.CPF = l[0]
			result = append(result, input)
		}
	}
	return result, nil
}

func ExtractNameAndBirthday(str string) (string, time.Time, error) {
	name := ""
	birth_day := time.Now()
	strSplit := strings.Split(str, " ")
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	for _, txt := range strSplit {
		if !re.MatchString(txt) {
			name = name + " " + txt
		} else {
			dt := strings.Split(txt, "/")
			if len(dt) != 3 {
				return "", time.Now(), fmt.Errorf("Invalid date")
			}
			y, err := strconv.Atoi(dt[2])
			if err != nil {
				return "", time.Now(), fmt.Errorf("Invalid date")
			}
			m, err := strconv.Atoi(dt[1])
			if err != nil {
				return "", time.Now(), fmt.Errorf("Invalid date")
			}
			d, err := strconv.Atoi(dt[0])

			if err != nil {
				return "", time.Now(), fmt.Errorf("Invalid date")
			}
			birth_day = time.Date(y, time.Month(m), d, 12, 15, 5, 5, time.Local)
		}
	}
	return name, birth_day, nil
}
