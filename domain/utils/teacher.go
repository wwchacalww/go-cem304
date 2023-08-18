package utils

import (
	"encoding/csv"
	"mime/multipart"
	"strconv"
	"strings"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

func CsvToTeachers(f multipart.File, classrooms []model.ClassroomInterface) ([]model.Subject, error) {
	var subjects []model.Subject
	var teachers []model.TeacherInterface
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range data {
		if len(line) == 8 {
			var ip_subject repository.SubjectInput
			var ip_teacher repository.TeacherInput
			var classroom model.ClassroomInterface
			var sub model.Subject
			for j, field := range line {
				switch {
				case j == 0:
					ip_teacher.Name = field
				case j == 1:
					ip_teacher.CPF = field
				case j == 2:
					ip_subject.Note = field
					classroom, _ = FindClassBySlug(classrooms, field)
				case j == 3:
					ip_subject.CH, _ = strconv.Atoi(field)
				case j == 4:
					ip_subject.Name = field
				}
			}
			tn := strings.Split(ip_teacher.Name, " ")
			if len(subjects) == 0 {
				teacher := model.NewTeacher(ip_teacher.Name, tn[0], ip_teacher.CPF, ip_subject.License)
				subject := model.NewSubject(ip_subject.Name, ip_subject.Name, classroom.GetLevel())
				teachers = append(teachers, teacher)
				subject.CH = ip_subject.CH
				subject.Note = ip_subject.Note + " - " + ip_subject.Name
				subject.AttachClassroom(classroom)
				subject.Teacher = teacher
				subjects = append(subjects, *subject)
			}
			for _, sbj := range subjects {
				if sbj.GetLevel() == classroom.GetLevel() && sbj.Teacher.GetCPF() == ip_teacher.CPF && sbj.GetName() == ip_subject.Name {
					sub = sbj
					sub.AttachClassroom(classroom)
				}
			}

			if sub.GetName() != "" {
				for i, sbj := range subjects {
					if sbj.GetID() == sub.GetID() {
						subjects = append(subjects[:i], subjects[i+1:]...)
					}
				}
				subjects = append(subjects, sub)
			} else {
				var checkTeacher = true
				var teacherExists model.TeacherInterface
				for _, t := range teachers {
					if t.GetCPF() == ip_teacher.CPF {
						checkTeacher = false
						teacherExists = t
						break
					}
				}
				teacher := model.NewTeacher(ip_teacher.Name, tn[0], ip_teacher.CPF, ip_subject.License)
				subject := model.NewSubject(ip_subject.Name, ip_subject.Name, classroom.GetLevel())
				subject.AttachClassroom(classroom)
				subject.CH = ip_subject.CH
				subject.Note = ip_subject.Note + " - " + ip_subject.Name
				if checkTeacher {
					teachers = append(teachers, teacher)
					subject.Teacher = teacher
				} else {
					subject.Teacher = teacherExists
				}
				subjects = append(subjects, *subject)
			}
			// for _, sbj := range subjects {
			// 	if sbj.GetLevel() == classroom.GetLevel() && sbj.Teacher.GetCPF() == ip_teacher.CPF && sbj.GetName() == ip_subject.Name {
			// 		var subject = sbj
			// 		subject.AttachClassroom(classroom)
			// 		subjects = append(subjects, subject)
			// 	} else {
			// 		teacher := model.NewTeacher(ip_teacher.Name, tn[0], ip_teacher.CPF, ip_subject.License)
			// 		subject := model.NewSubject(ip_subject.Name, ip_subject.Name, classroom.GetLevel())
			// 		subject.AttachClassroom(classroom)
			// 		subject.Teacher = teacher
			// 		subjects = append(subjects, *subject)
			// 	}
			// }
		}
	}

	// log.Println(subjects)
	return subjects, nil
}
