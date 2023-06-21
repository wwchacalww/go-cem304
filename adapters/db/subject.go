package db

import (
	"database/sql"
	"log"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type SubjectDB struct {
	db *sql.DB
}

func NewSubjectDB(db *sql.DB) *SubjectDB {
	return &SubjectDB{db: db}
}

func (t *SubjectDB) Create(input repository.SubjectInput) (model.SubjectInterface, error) {
	subject := model.NewSubject(input.Name, input.License, input.Level)
	subject.Semester = input.Semester
	subject.Note = input.Note
	subject.CH = input.CH
	_, err := subject.IsValid()
	if err != nil {
		return nil, err
	}

	stmt, err := t.db.Prepare(`INSERT INTO subjects (
		id,
		name,
		license,
		level,
		grade,
		note,
		ch,
		year,
		semester,
		created_at,
		updated_at
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`)
	if err != nil {
		log.Println("Error SQL")
		return nil, err
	}

	_, err = stmt.Exec(
		subject.GetID(),
		subject.GetName(),
		subject.GetLicense(),
		subject.GetLevel(),
		subject.GetGrade(),
		subject.GetNote(),
		subject.GetCH(),
		subject.GetYear(),
		subject.GetSemester(),
		subject.GetCreatedAt(),
		subject.GetUpdatedAt(),
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
	return subject, nil
}

func (t *SubjectDB) FindById(id string) (model.SubjectInterface, error) {
	var subject model.Subject
	var ip struct {
		Note     sql.NullString
		CH       sql.NullInt32
		Semester sql.NullString
	}
	fields := "id, name,	license,	level,	grade,	note,	ch,	year,	semester, created_at, updated_at"

	stmt, err := t.db.Prepare(`SELECT ` + fields + `
		FROM subjects WHERE id = $1
	`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&subject.ID,
		&subject.Name,
		&subject.License,
		&subject.Level,
		&subject.Grade,
		&ip.Note,
		&ip.CH,
		&subject.Year,
		&ip.Semester,
		&subject.CreatedAt,
		&subject.UpdatedAt,
	)

	subject.Note = ip.Note.String
	subject.CH = int(ip.CH.Int32)
	subject.Semester = ip.Semester.String

	stmt.Close()

	return &subject, nil
}

func (t *SubjectDB) FindByLicense(license string) ([]model.SubjectInterface, error) {
	var subjects []model.SubjectInterface
	fields := "id, name,	license,	level,	grade,	note,	ch,	year,	semester, created_at, updated_at"

	rows, err := t.db.Query("SELECT "+fields+" FROM subjects WHERE license like $1 ORDER BY name ASC", "%"+license+"%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var subject model.Subject
		var ip struct {
			Note     sql.NullString
			CH       sql.NullInt32
			Semester sql.NullString
		}
		err = rows.Scan(
			&subject.ID,
			&subject.Name,
			&subject.License,
			&subject.Level,
			&subject.Grade,
			&ip.Note,
			&ip.CH,
			&subject.Year,
			&ip.Semester,
			&subject.CreatedAt,
			&subject.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		subject.Note = ip.Note.String
		subject.CH = int(ip.CH.Int32)
		subject.Semester = ip.Semester.String

		subjects = append(subjects, &subject)
	}

	return subjects, nil
}

func (t *SubjectDB) FindByName(name string) ([]model.SubjectInterface, error) {
	var subjects []model.SubjectInterface
	fields := "id, name,	license,	level,	grade,	note,	ch,	year,	semester, created_at, updated_at"

	rows, err := t.db.Query("SELECT "+fields+" FROM subjects WHERE name like $1 ORDER BY name ASC", "%"+name+"%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var subject model.Subject
		var ip struct {
			Note     sql.NullString
			CH       sql.NullInt32
			Semester sql.NullString
		}
		err = rows.Scan(
			&subject.ID,
			&subject.Name,
			&subject.License,
			&subject.Level,
			&subject.Grade,
			&ip.Note,
			&ip.CH,
			&subject.Year,
			&ip.Semester,
			&subject.CreatedAt,
			&subject.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		subject.Note = ip.Note.String
		subject.CH = int(ip.CH.Int32)
		subject.Semester = ip.Semester.String

		subjects = append(subjects, &subject)
	}

	return subjects, nil
}
