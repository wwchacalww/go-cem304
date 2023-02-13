package db

import (
	"database/sql"
	"strconv"
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
		classStmt, err := std.db.Prepare("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where id=$1")
		if err != nil {
			return nil, err
		}

		err = classStmt.QueryRow(input.ClassroomID).Scan(
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
		classStmt.Close()
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
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12)`)
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

func (std *StudentDB) FindByEducar(educar int64) (model.StudentInterface, error) {
	var student model.Student
	var classroom_id string

	std_fields := "id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, created_at, students.updated_at"
	stmt, err := std.db.Prepare("SELECT " + std_fields + " from students WHERE ieducar = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(educar).Scan(
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

func (std *StudentDB) List(classroom_id string) ([]model.StudentInterface, error) {
	var class model.Classroom
	classStmt, err := std.db.Prepare("SELECT id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at from classrooms where id=$1")
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
	classStmt.Close()

	var students []model.StudentInterface
	std_fields := "id, name, birth_day, gender, anne, note, ieducar, educa_df, status, created_at, students.updated_at"
	rows, err := std.db.Query("SELECT "+std_fields+" from students where classroom_id = $1 ORDER BY name ASC", classroom_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student model.Student
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
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		student.Classroom = &class
		students = append(students, &student)
	}

	return students, nil
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
	// Listar as turmas
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
	for _, s := range mass {
		class, err := utils.FindClassById(classrooms, s.ClassroomID)
		if err != nil {
			return nil, err
		}
		student, err := model.NewStudent(
			s.Name, s.BirthDay, s.Educar,
		)
		if err != nil {
			return nil, err
		}
		student.Gender = s.Gender
		student.ANNE = s.ANNE
		student.Classroom = &class
		students = append(students, student)
	}

	query := `INSERT INTO Students (
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

	for i, st := range students {
		if i == 0 {
			query = query + "('" + st.GetID() + "',"
		} else {
			query = query + ", ('" + st.GetID() + "',"
		}
		query = query + "'" + st.GetName() + "',"
		query = query + "'" + string(st.GetBirthDay().Format(time.RFC3339)) + "',"
		query = query + "'" + st.GetGender() + "',"
		query = query + "'" + st.GetANNE() + "',"
		query = query + "'" + st.GetNote() + "',"
		query = query + strconv.FormatInt(st.GetEducar(), 10) + ","
		query = query + "'" + st.GetEducaDF() + "',"
		query = query + "'" + st.GetClassroom().GetID() + "',"
		query = query + "true,"
		query = query + "'" + string(st.GetCreatedAt().Format(time.RFC3339)) + "',"
		query = query + "'" + string(st.GetUpdatedAt().Format(time.RFC3339)) + "') "
	}

	_, err = std.db.Exec(query)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (c *StudentDB) ChangeClassroom(id, classroom_id string) error {
	_, err := c.db.Exec("UPDATE students SET classroom_id=$1 WHERE id=$2", classroom_id, id)
	if err != nil {
		return err
	}

	return nil
}
