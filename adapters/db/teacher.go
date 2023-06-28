package db

import (
	"database/sql"
	"log"
	"time"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type TeacherDB struct {
	db *sql.DB
}

func NewTeacherDB(db *sql.DB) *TeacherDB {
	return &TeacherDB{db: db}
}

func (t *TeacherDB) Create(input repository.TeacherInput) (model.TeacherInterface, error) {
	teacher := model.NewTeacher(input.Name, input.Nick, input.CPF, input.License)
	teacher.Save(input.Email, input.BirthDay, input.Gender, input.Fones, input.License, input.Note)

	_, err := teacher.IsValid()
	if err != nil {
		return nil, err
	}

	stmt, err := t.db.Prepare(`INSERT INTO teachers (
		id,
		name,
		nick,
		birth_day,
		gender,
		cpf,
		fones,
		email,
		license,
		note,
		status,
		created_at,
		updated_at
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`)
	if err != nil {
		log.Println("Error SQL")
		return nil, err
	}

	_, err = stmt.Exec(
		teacher.GetID(),
		teacher.GetName(),
		teacher.GetNick(),
		teacher.GetBirthDay(),
		teacher.GetGender(),
		teacher.GetCPF(),
		teacher.GetFones(),
		teacher.GetEmail(),
		teacher.GetLicense(),
		teacher.GetNote(),
		teacher.GetStatus(),
		teacher.GetCreatedAt(),
		teacher.GetUpdatedAt(),
	)
	if err != nil {
		log.Println("Error INSERT")
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		log.Println("Error CLOSE")
		return nil, err
	}
	return teacher, nil
}

func (t *TeacherDB) FindById(id string) (model.TeacherInterface, error) {
	var classrooms []model.ClassroomInterface
	var subjects []model.SubjectInterface
	// Verificar se o professor está enturmado e com disciplina
	count := 0
	rows, err := t.db.Query(`
	SELECT 
		classroom_id, c.name, c.level, c.grade, c.shift, c.description, c.ANNE, c.year, c.status, c.created_at, c.updated_at,
		subject_id, slug,	s.license,	s.level, s.grade,	s.note,	wch,	s.year,	s.semester, s.created_at, s.updated_at 
	from classrooms_subjects_teachers cst 
		LEFT JOIN teachers t ON t.id = cst.teacher_id
		LEFT JOIN subjects s ON s.id = cst.subject_id
		LEFT JOIN classrooms c ON c.id = cst.classroom_id
	WHERE cst.teacher_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var classroom model.Classroom
		var subject model.Subject
		var ipSubject struct {
			Note     sql.NullString
			CH       sql.NullInt32
			Semester sql.NullString
		}
		err := rows.Scan(
			&classroom.ID,
			&classroom.Name,
			&classroom.Level,
			&classroom.Grade,
			&classroom.Shift,
			&classroom.Description,
			&classroom.ANNE,
			&classroom.Year,
			&classroom.Status,
			&classroom.CreatedAt,
			&classroom.UpdatedAt,
			&subject.ID,
			&subject.Name,
			&subject.License,
			&subject.Level,
			&subject.Grade,
			&ipSubject.Note,
			&ipSubject.CH,
			&subject.Year,
			&ipSubject.Semester,
			&subject.CreatedAt,
			&subject.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		classrooms = append(classrooms, &classroom)
		subjects = append(subjects, &subject)
		count = count + 1
	}
	rows.Close()
	var teacher model.Teacher
	var ip struct {
		CPF      sql.NullString
		Email    sql.NullString
		Fones    sql.NullString
		Gender   sql.NullString
		BirthDay sql.NullTime
		Note     sql.NullString
	}
	fields := "id, name, nick,	birth_day,	gender,	cpf,	fones,	email,	license,	note,	status, created_at, updated_at"

	stmt, err := t.db.Prepare(`SELECT ` + fields + `
		FROM teachers WHERE id = $1
	`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Nick,
		&ip.BirthDay,
		&ip.Gender,
		&ip.CPF,
		&ip.Fones,
		&ip.Email,
		&teacher.License,
		&ip.Note,
		&teacher.Status,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	teacher.BirthDay = ip.BirthDay.Time
	teacher.Gender = ip.Gender.String
	teacher.CPF = ip.CPF.String
	teacher.Fones = ip.Fones.String
	teacher.Email = ip.Email.String
	teacher.Note = ip.Note.String

	if count > 0 {
		teacher.Classrooms = classrooms
		teacher.Subjects = subjects
	}
	stmt.Close()

	return &teacher, nil
}

func (t *TeacherDB) FindByCPF(cpf string) (model.TeacherInterface, error) {
	var teacher model.Teacher
	var ip struct {
		CPF      sql.NullString
		Email    sql.NullString
		Fones    sql.NullString
		Gender   sql.NullString
		BirthDay sql.NullTime
		Note     sql.NullString
	}
	fields := "id, name, nick,	birth_day,	gender,	cpf,	fones,	email,	license,	note,	status, created_at, updated_at"

	stmt, err := t.db.Prepare(`SELECT ` + fields + `
		FROM teachers WHERE cpf = $1
	`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(cpf).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Nick,
		&ip.BirthDay,
		&ip.Gender,
		&ip.CPF,
		&ip.Fones,
		&ip.Email,
		&teacher.License,
		&ip.Note,
		&teacher.Status,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	teacher.BirthDay = ip.BirthDay.Time
	teacher.Gender = ip.Gender.String
	teacher.CPF = ip.CPF.String
	teacher.Fones = ip.Fones.String
	teacher.Email = ip.Email.String
	teacher.Note = ip.Note.String

	stmt.Close()

	return &teacher, nil
}

func (t *TeacherDB) FindByName(name string) ([]model.TeacherInterface, error) {
	var teachers []model.TeacherInterface
	fields := "id, name, nick,	birth_day,	gender,	cpf,	fones,	email,	license,	note,	status, created_at, updated_at"
	rows, err := t.db.Query("SELECT "+fields+" FROM teachers WHERE name like $1 ORDER BY name ASC", "%"+name+"%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var teacher model.Teacher
		var ip struct {
			CPF      sql.NullString
			Email    sql.NullString
			Fones    sql.NullString
			Gender   sql.NullString
			BirthDay sql.NullTime
			Note     sql.NullString
		}
		err = rows.Scan(
			&teacher.ID,
			&teacher.Name,
			&teacher.Nick,
			&ip.BirthDay,
			&ip.Gender,
			&ip.CPF,
			&ip.Fones,
			&ip.Email,
			&teacher.License,
			&ip.Note,
			&teacher.Status,
			&teacher.CreatedAt,
			&teacher.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		teacher.BirthDay = ip.BirthDay.Time
		teacher.Gender = ip.Gender.String
		teacher.CPF = ip.CPF.String
		teacher.Fones = ip.Fones.String
		teacher.Email = ip.Email.String
		teacher.Note = ip.Note.String

		teachers = append(teachers, &teacher)
	}

	return teachers, nil
}

func (t *TeacherDB) AttachClassroomSubject(id, classroom_id, subject_id, slug string, wch int32, start_course, end_course time.Time) error {
	_, err := t.db.Exec("INSERT INTO classrooms_subjects_teachers ( classroom_id, subject_id, teacher_id, wch, slug, start_course, end_course) values ($1, $2, $3, $4, $5, $6, $7)",
		classroom_id, subject_id, id, wch, slug, start_course, end_course,
	)
	if err != nil {
		return err
	}

	return nil
}

func (t *TeacherDB) Save(teacher model.TeacherInterface) error {
	_, err := t.db.Exec("UPDATE teachers SET email= $1, fones=$2, license=$3, gender=$4, birth_day=$5, note=$6 WHERE id=$6",
		teacher.GetEmail(), teacher.GetFones(), teacher.GetLicense(), teacher.GetGender(), teacher.GetBirthDay(), teacher.GetNote(),
	)
	if err != nil {
		return err
	}

	return nil
}
