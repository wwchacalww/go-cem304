package utils

import (
	"encoding/csv"
	"mime/multipart"
	"strconv"
	"wwchacalww/go-cem304/domain/repository"
)

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
