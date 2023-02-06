package db

import (
	"database/sql"
	"time"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type StudentDB struct {
	db *sql.DB
}

func NewStudentDB(db *sql.DB) *StudentDB {
	return &StudentDB{db: db}
}

func (std *StudentDB) Create(input repository.StudentInput) (model.StudentInterface, error) {
	classIsSet := false
	var class model.Classroom
	if input.ClassroomID != "" {
		stmt, err := std.db.Prepare("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where id=$1")
		if err != nil {
			return nil, err
		}

		err = stmt.QueryRow(input.ClassroomID).Scan(
			&class.ID,
			&class.Name,
			&class.Level,
			&class.Grade,
			&class.Shift,
			&class.Description,
			&class.ANNE,
			&class.Year,
			&class.Status,
			&class.CreatedAt,
			&class.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		classIsSet = true
	}

	student, err := model.NewStudent(input.Name, input.BirthDay, input.Educar)
	if err != nil {
		return nil, err
	}

	stmt, err := std.db.Prepare(`INSERT INTO Students (
		id,
		name,
		birth_day,
		anne,
		note,
		ieducar,
		educa_df,
		classroom_id,
		status,
		created_at,
		updated_at
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`)
	if err != nil {
		return nil, err
	}

	if classIsSet {
		student.Classroom = &class
		_, err = stmt.Exec(
			student.GetID(),
			student.GetName(),
			student.GetBirthDay(),
			student.GetANNE(),
			student.GetNote(),
			student.GetEducar(),
			student.GetEducaDF(),
			student.GetClassroom().GetID(),
			student.GetStatus(),
			student.GetCreatedAt(),
			student.GetUpdatedAt(),
		)
	} else {
		_, err = stmt.Exec(
			student.GetID(),
			student.GetName(),
			student.GetBirthDay(),
			student.GetANNE(),
			student.GetNote(),
			student.GetEducar(),
			student.GetEducaDF(),
			"",
			student.GetStatus(),
			student.GetCreatedAt(),
			student.GetUpdatedAt(),
		)
	}

	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (std *StudentDB) FindById(id string) (model.StudentInterface, error) {
	var student model.Student
	var classroom_id string
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.ANNE,
		&student.Note,
		&student.Educar,
		&student.EducaDF,
		&classroom_id,
		&student.Status,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (std *StudentDB) FindByName(name string) ([]model.StudentInterface, error) {
	var students []model.StudentInterface
	var classroom_id string
	rows, err := std.db.Query("SELECT id, name, birth_day, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where name like $1 ORDER BY name ASC", "%"+name+"%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var student model.Student
		err = rows.Scan(
			&student.ID,
			&student.ID,
			&student.Name,
			&student.BirthDay,
			&student.ANNE,
			&student.Note,
			&student.Educar,
			&student.EducaDF,
			&classroom_id,
			&student.Status,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

func (std *StudentDB) List(classroom_id string) (model.StudentInterface, error) {
	var student model.Student
	var class_id string
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where classroom_id=$1 ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(classroom_id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.ANNE,
		&student.Note,
		&student.Educar,
		&student.EducaDF,
		&class_id,
		&student.Status,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (std *StudentDB) Enable(id string) (model.StudentInterface, error) {
	_, err := std.db.Exec("update students set status=true where id=$1", id)
	if err != nil {
		return nil, err
	}

	var student model.Student
	var classroom_id string
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.ANNE,
		&student.Note,
		&student.Educar,
		&student.EducaDF,
		&classroom_id,
		&student.Status,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (std *StudentDB) Disable(id string) (model.StudentInterface, error) {
	_, err := std.db.Exec("update students set status=false where id=$1", id)
	if err != nil {
		return nil, err
	}

	var student model.Student
	var classroom_id string
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.ANNE,
		&student.Note,
		&student.Educar,
		&student.EducaDF,
		&classroom_id,
		&student.Status,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (std *StudentDB) ANNE(id, anne string) (model.StudentInterface, error) {
	_, err := std.db.Exec("UPDATE students SET ANNE=$1 WHERE id=$2", anne, id)
	if err != nil {
		return nil, err
	}

	var student model.Student
	var classroom_id string
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.ANNE,
		&student.Note,
		&student.Educar,
		&student.EducaDF,
		&classroom_id,
		&student.Status,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (std *StudentDB) AddMass(mass []model.StudentInterface) ([]model.StudentInterface, error) {
	var query string
	query = `INSERT INTO Students (
		id,
		name,
		level,
		grade,
		shift,
		description,
		ANNE,
		year,
		status,
		created_at,
		updated_at
	) values `
	for i, c := range mass {
		if i == 0 {
			query = query + "('" + std.GetID() + "',"
		} else {
			query = query + ", ('" + std.GetID() + "',"
		}
		query = query + "'" + std.GetName() + "',"
		query = query + "'" + std.GetLevel() + "',"
		query = query + "'" + std.GetGrade() + "',"
		query = query + "'" + std.GetShift() + "',"
		query = query + "'" + std.GetDescription() + "',"
		query = query + "'" + std.GetANNE() + "',"
		query = query + "'" + std.GetYear() + "',"
		query = query + "true,"
		query = query + "'" + string(std.GetCreatedAt().Format(time.RFC3339)) + "',"
		query = query + "'" + string(std.GetUpdatedAt().Format(time.RFC3339)) + "') "
	}

	_, err := std.db.Exec(query)
	if err != nil {
		return nil, err
	}

	return mass, nil
}
