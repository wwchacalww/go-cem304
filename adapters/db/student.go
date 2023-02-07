package db

import (
	"database/sql"
	"time"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
	"wwchacalww/go-cem304/domain/utils"
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
		stmt.Close()
	}

	student, err := model.NewStudent(input.Name, input.BirthDay, input.Educar)
	student.Gender = input.Gender
	student.ANNE = input.ANNE
	student.Note = input.Note
	student.EducaDF = input.EducaDF

	if err != nil {
		return nil, err
	}

	_, err = student.IsValid()
	if err != nil {
		return nil, err
	}

	stmt, err := std.db.Prepare(`INSERT INTO Students (
		id,
		name,
		birth_day,
		gender,
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
			student.GetGender(),
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
			student.GetGender(),
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

	std_fields := "id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, created_at, students.updated_at"
	stmt, err := std.db.Prepare("SELECT " + std_fields + " from students WHERE id = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.Gender,
		&student.ANNE,
		&student.Note,
		&student.Educar,
		&student.EducaDF,
		&classroom_id,
		&student.Status,
		&student.CreatedAt,
		&student.UpdatedAt,
	)
	stmt.Close()

	if classroom_id != "" {
		var class model.Classroom
		class_fields := "id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at"

		classStmt, err := std.db.Prepare("SELECT " + class_fields + " from classrooms where id=$1")
		if err != nil {
			return nil, err
		}

		err = classStmt.QueryRow(classroom_id).Scan(
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
		student.Classroom = &class
		classStmt.Close()
	}

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (std *StudentDB) FindByName(name string) ([]model.StudentInterface, error) {
	startDate := time.Date(time.Now().Year(), time.January, 1, 12, 15, 5, 5, time.Local)
	var classrooms []model.Classroom
	class_fields := "id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at"
	classRows, err := std.db.Query("SELECT "+class_fields+" from classrooms WHERE created_at > $1", startDate)
	if err != nil {
		return nil, err
	}
	defer classRows.Close()
	for classRows.Next() {
		var class model.Classroom
		err = classRows.Scan(
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

		classrooms = append(classrooms, class)
	}

	var students []model.StudentInterface
	std_fields := "id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, created_at, students.updated_at"
	rows, err := std.db.Query("SELECT "+std_fields+" from students where name like $1 ORDER BY name ASC", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var student model.Student
		var classroom_id string
		err = rows.Scan(
			&student.ID,
			&student.Name,
			&student.BirthDay,
			&student.Gender,
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

		if classroom_id != "" {
			xpto, err := utils.FindClassById(classrooms, classroom_id)
			if err != nil {
				return nil, err
			}
			student.Classroom = &xpto
		}
		students = append(students, &student)
	}

	return students, nil
}

func (std *StudentDB) List(classroom_id string) (model.StudentInterface, error) {
	var student model.Student
	var class_id string
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where classroom_id=$1 ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(classroom_id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.Gender,
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
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.Gender,
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
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.Gender,
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
	stmt, err := std.db.Prepare("SELECT id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, created_at, updated_at from students where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&student.ID,
		&student.Name,
		&student.BirthDay,
		&student.Gender,
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

func (std *StudentDB) AddMass(mass []repository.StudentInput) ([]model.StudentInterface, error) {
	var class model.Classroom
	classroom_id := mass[0].ClassroomID
	stmt, err := std.db.Prepare("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where id=$1")
	if err != nil {
		return nil, err
	}

	var students []model.StudentInterface
	err = stmt.QueryRow(classroom_id).Scan(
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
	stmt.Close()

	var query string
	query = `INSERT INTO Students (
		id,
		name,
		birth_day,
		gender,
		anne,
		note,
		ieducar,
		educa_df,
		classroom_id,
		status,
		created_at,
		updated_at
	) values `
	for i, s := range mass {
		study, err := model.NewStudent(s.Name, s.BirthDay, s.Educar)
		if err != nil {
			return nil, err
		}
		study.ANNE = s.ANNE
		study.Note = s.Note
		study.Educar = s.Educar
		study.EducaDF = s.EducaDF
		study.Classroom = &class
		students = append(students, study)
		if i == 0 {
			query = query + "('" + study.GetID() + "',"
		} else {
			query = query + ", ('" + study.GetID() + "',"
		}
		query = query + "'" + study.GetName() + "',"
		query = query + "'" + string(study.GetBirthDay().Format(time.RFC3339)) + "',"
		query = query + "'" + study.GetGender() + "',"
		query = query + "'" + study.GetANNE() + "',"
		query = query + "'" + study.GetNote() + "',"
		query = query + string(study.GetEducar()) + ","
		query = query + "'" + study.GetEducaDF() + "',"
		query = query + "'" + classroom_id + "',"
		query = query + "true,"
		query = query + "'" + string(study.GetCreatedAt().Format(time.RFC3339)) + "',"
		query = query + "'" + string(study.GetUpdatedAt().Format(time.RFC3339)) + "') "
	}

	_, err = std.db.Exec(query)
	if err != nil {
		return nil, err
	}

	return students, nil
}
