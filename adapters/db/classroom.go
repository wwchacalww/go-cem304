package db

import (
	"database/sql"
	"log"
	"time"
	"wwchacalww/go-cem304/domain/model"
)

type ClassroomDB struct {
	db *sql.DB
}

func NewClassroomDB(db *sql.DB) *ClassroomDB {
	return &ClassroomDB{db: db}
}

func (c *ClassroomDB) Create(class model.ClassroomInterface) error {
	stmt, err := c.db.Prepare(`INSERT INTO classrooms (
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
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		class.GetID(),
		class.GetName(),
		class.GetLevel(),
		class.GetGrade(),
		class.GetShift(),
		class.GetDescription(),
		class.GetANNE(),
		class.GetYear(),
		class.GetStatus(),
		class.GetCreatedAt(),
		class.GetUpdatedAt(),
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *ClassroomDB) FindById(id string) (model.ClassroomInterface, error) {
	var class model.Classroom
	stmt, err := c.db.Prepare("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
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

	var students []model.StudentInterface
	std_fields := "id, name, birth_day, gender, anne, note, ieducar, educa_df, status, cpf, created_at, students.updated_at"
	rows, err := c.db.Query("SELECT "+std_fields+" FROM students WHERE classroom_id = $1 ORDER BY name", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var student model.Student
		var cpf sql.NullString
		err = rows.Scan(
			&student.ID,
			&student.Name,
			&student.BirthDay,
			&student.Gender,
			&student.ANNE,
			&student.Note,
			&student.Educar,
			&student.EducaDF,
			&student.Status,
			&cpf,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		student.CPF = cpf.String
		if err != nil {
			return nil, err
		}

		students = append(students, &student)
	}
	class.Students = students
	return &class, nil
}

func (c *ClassroomDB) FindByName(name string) ([]model.ClassroomInterface, error) {
	var classrooms []model.ClassroomInterface
	rows, err := c.db.Query("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where name like $1 ORDER BY name ASC", "%"+name+"%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var class model.Classroom
		err = rows.Scan(
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
		classrooms = append(classrooms, &class)
	}

	return classrooms, nil
}

func (c *ClassroomDB) List(year string) ([]model.ClassroomInterface, error) {
	var classrooms []model.ClassroomInterface
	rows, err := c.db.Query("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where year like $1 ORDER BY name ASC", year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var class model.Classroom
		err = rows.Scan(
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
		classrooms = append(classrooms, &class)
	}
	return classrooms, nil
}

func (c *ClassroomDB) Enable(id string) (model.ClassroomInterface, error) {
	_, err := c.db.Exec("update classrooms set status=true where id=$1", id)
	if err != nil {
		return nil, err
	}

	var class model.Classroom
	stmt, err := c.db.Prepare("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
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
	return &class, nil
}

func (c *ClassroomDB) Disable(id string) (model.ClassroomInterface, error) {
	_, err := c.db.Exec("update classrooms set status=false where id=$1", id)
	if err != nil {
		return nil, err
	}
	log.Println("teste")
	var class model.Classroom
	stmt, err := c.db.Prepare("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at FROM classrooms where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
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

	return &class, nil
}

func (c *ClassroomDB) ANNE(id, anne string) (model.ClassroomInterface, error) {
	_, err := c.db.Exec("UPDATE classrooms SET ANNE=$1 WHERE id=$2", anne, id)
	if err != nil {
		return nil, err
	}

	var class model.Classroom
	stmt, err := c.db.Prepare("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
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
	return &class, nil
}

func (c *ClassroomDB) AddMass(mass []model.ClassroomInterface) ([]model.ClassroomInterface, error) {
	var query string
	query = `INSERT INTO classrooms (
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
			query = query + "('" + c.GetID() + "',"
		} else {
			query = query + ", ('" + c.GetID() + "',"
		}
		query = query + "'" + c.GetName() + "',"
		query = query + "'" + c.GetLevel() + "',"
		query = query + "'" + c.GetGrade() + "',"
		query = query + "'" + c.GetShift() + "',"
		query = query + "'" + c.GetDescription() + "',"
		query = query + "'" + c.GetANNE() + "',"
		query = query + "'" + c.GetYear() + "',"
		query = query + "true,"
		query = query + "'" + string(c.GetCreatedAt().Format(time.RFC3339)) + "',"
		query = query + "'" + string(c.GetUpdatedAt().Format(time.RFC3339)) + "') "
	}

	_, err := c.db.Exec(query)
	if err != nil {
		return nil, err
	}

	return mass, nil
}
