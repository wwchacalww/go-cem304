package utils

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"wwchacalww/go-cem304/domain/dtos"
)

func CsvToCards(f multipart.File, term, classroom_id string) (dtos.InputImportCard, error) {
	var cards dtos.InputImportCard
	cards.Term = term
	cards.ClassroomID = classroom_id
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		return dtos.InputImportCard{}, err
	}
	subs, err := registerSubjects(data[0])
	if err != nil {
		return dtos.InputImportCard{}, err
	}

	for i, line := range data {
		if i != 0 {
			card, err := registerCard(line, subs)
			if err != nil {
				return dtos.InputImportCard{}, err
			}
			cards.Cards = append(cards.Cards, card)
		}

	}
	return cards, nil
}

func registerSubjects(line []string) ([]dtos.SubjectCrtl, error) {
	var subs []dtos.SubjectCrtl
	for j, field := range line {
		if j%2 == 1 {
			var sbctrl = dtos.SubjectCrtl{
				Ind:     j,
				Subject: field,
			}
			subs = append(subs, sbctrl)
		}
	}
	// if len(subs) != 7 {
	// 	return nil, fmt.Errorf("Invalid list of subjects")
	// }
	return subs, nil
}

func findSubject(i int, subs []dtos.SubjectCrtl) (string, error) {
	for _, field := range subs {
		if field.Ind == i {
			return field.Subject, nil
		}

	}
	return "", fmt.Errorf("Fail on list subjects")
}

func registerCard(line []string, subs []dtos.SubjectCrtl) (dtos.CardStudent, error) {
	var ip dtos.CardStudent
	var subIp dtos.SubjectIp
	subLenght := len(subs)*2 + 1
	if len(line) == subLenght {
		for j, field := range line {
			if j == 0 {
				ip.EducaDF = field
			} else {
				if j%2 == 1 {
					sub, err := findSubject(j, subs)
					if err != nil {
						return dtos.CardStudent{}, err
					}
					subIp.Name = sub
					subIp.Note = field
				} else {
					n, err := strconv.Atoi(field)
					if err != nil {
						return dtos.CardStudent{}, err
					}
					subIp.Absences = n
					ip.Subjects = append(ip.Subjects, subIp)
				}
			}
		}
	}

	return ip, nil
}
