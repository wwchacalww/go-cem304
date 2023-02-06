package utils

import (
	"encoding/csv"
	"mime/multipart"
	"wwchacalww/go-cem304/domain/repository"
)

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