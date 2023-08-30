package db

import (
	"database/sql"
	"fmt"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type EvaluationDB struct {
	db *sql.DB
}

func NewEvaluationDB(db *sql.DB) *EvaluationDB {
	return &EvaluationDB{db: db}
}

func (e *EvaluationDB) Create(input repository.EvaluationInput) error {
	stmt, err := e.db.Prepare(`INSERT INTO evaluations (
		student_id,
		classroom_subject_teacher_id,
		term,
		note,
		absences
	) values ($1, $2, $3, $4,$5)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		input.Student,
		input.Subject,
		input.Term,
		input.Note,
		input.Absences,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	return err
}

func (e *EvaluationDB) Update(note string, id, absences int) error {
	_, err := e.db.Exec("UPDATE evaluations SET note=$1, absences=$2 WHERE id=$3", note, absences, id)
	if err != nil {
		return err
	}

	return nil
}

func (e *EvaluationDB) FindById(id int) (model.EvaluationInterface, error) {
	var student_id string
	var subject_id int
	var term string
	stmt, err := e.db.Prepare(`
		SELECT 
			id, student_id, classroom_subject_teacher_id, term, note, absences 
		FROM
			evaluations
		WHERE
			id = $1
	`)
	if err != nil {
		return nil, err
	}
	var eval model.Evaluation
	err = stmt.QueryRow(id).Scan(
		&eval.ID,
		&student_id,
		&subject_id,
		&term,
		&eval.Note,
		&eval.Absences,
	)
	if err != nil {
		return nil, err
	}

	std, err := e.findStudent(student_id)
	if err != nil {
		return nil, err
	}
	subject, err := e.findSubject(subject_id)
	if err != nil {
		return nil, err
	}

	eval.Student = std
	eval.Subject = subject
	eval.Term = term

	return &eval, nil
}

func (e *EvaluationDB) FindEvaluation(student_id, term string, subject_id int) (model.EvaluationInterface, error) {
	var eval model.Evaluation
	std, err := e.findStudent(student_id)
	if err != nil {
		return nil, err
	}
	subject, err := e.findSubject(subject_id)
	if err != nil {
		return nil, err
	}

	eval.Student = std
	eval.Subject = subject
	eval.Term = term

	stmt, err := e.db.Prepare(`
		SELECT 
			id, note, absences 
		FROM
			evaluations
		WHERE
			student_id = $1 AND
			term = $2 AND
			classroom_subject_teacher_id = $3
		LIMIT 1
	`)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(student_id, term, subject_id).Scan(
		&eval.ID,
		&eval.Note,
		&eval.Absences,
	)
	if err != nil {
		return nil, err
	}

	return &eval, nil
}

func (e *EvaluationDB) CheckEvaluationExists(student_id, term string, subject_id int) (int, error) {
	var ID int
	e.db.QueryRow(`
	SELECT 
			id
		FROM
			evaluations
		WHERE
			student_id = $1 AND
			term = $2 AND
			classroom_subject_teacher_id = $3
		LIMIT 1
		`, student_id, term, subject_id).Scan(&ID)
	if ID == 0 {
		return ID, nil
	}
	return ID, nil
}

func (e *EvaluationDB) GetAndCheckStudent(classroom_id, educadf string) (string, error) {
	var student_id string
	var classroom sql.NullString
	stmt, err := e.db.Prepare("SELECT id, classroom_id FROM students WHERE educa_df=$1")
	if err != nil {
		return "", err
	}

	err = stmt.QueryRow(educadf).Scan(&student_id, &classroom)
	if err != nil {
		return "", err
	}
	if classroom.String != classroom_id {
		return student_id, fmt.Errorf("Check classroom of student_id " + student_id)
	}

	return student_id, nil
}

func (e *EvaluationDB) GetSubjectId(classroom_id, name string) (int, error) {
	var subject_id int
	stmt, err := e.db.Prepare(`
	SELECT 
		cst.id
	FROM classrooms_subjects_teachers cst 
	LEFT JOIN subjects s ON s.id = cst.subject_id
	WHERE cst.classroom_id=$1 AND s.name = $2
	LIMIT 1
	`)
	if err != nil {
		return 0, err
	}
	stmt.QueryRow(classroom_id, name).Scan(&subject_id)
	return subject_id, nil
}

func (e *EvaluationDB) findStudent(id string) (model.StudentInterface, error) {
	var student model.Student
	var classroom_id string
	var ip struct {
		Address sql.NullString
		City    sql.NullString
		CEP     sql.NullString
		Fones   sql.NullString
		CPF     sql.NullString
	}

	std_fields := "id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, address, city, cep, fones, cpf, created_at, students.updated_at"
	stmt, err := e.db.Prepare("SELECT " + std_fields + " from students WHERE id = $1")
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
		&ip.Address,
		&ip.City,
		&ip.CEP,
		&ip.Fones,
		&ip.CPF,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	student.Address = ip.Address.String
	student.City = ip.City.String
	student.CEP = ip.CEP.String
	student.Fones = ip.Fones.String
	student.CPF = ip.CPF.String

	stmt.Close()

	if classroom_id != "" {
		var class model.Classroom
		class_fields := "id, name, level, grade, shift, description, ANNE, year, status, created_at, updated_at"

		classStmt, err := e.db.Prepare("SELECT " + class_fields + " from classrooms where id=$1")
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

	parentRows, err := e.db.Query(`SELECT p.id, p.name, cpf, email, fones, relationship, responsible 
		FROM parents_students ps 
		LEFT JOIN parents p ON p.id = ps.parent_id
		WHERE ps.student_id = $1`, student.GetID())
	if err != nil {
		return nil, err
	}
	defer parentRows.Close()
	var parents []model.ParentInterface
	for parentRows.Next() {
		var par model.Parent
		var cpf, email, fones, relationship sql.NullString
		err = parentRows.Scan(
			&par.ID,
			&par.Name,
			&cpf,
			&email,
			&fones,
			&relationship,
			&par.Responsible,
		)
		if err != nil {
			return nil, err
		}
		par.CPF = cpf.String
		par.Email = email.String
		par.Fones = fones.String
		par.Relationship = relationship.String
		parents = append(parents, &par)
	}
	if len(parents) > 0 {
		student.Parents = parents
	}

	return &student, nil
}

func (e *EvaluationDB) findSubject(id int) (model.SubjectInterface, error) {
	var subject model.Subject
	var teacher model.Teacher
	var ip struct {
		Classroom_id sql.NullString
		CH           sql.NullInt32
		Semester     sql.NullString
		CPF          sql.NullString
		Email        sql.NullString
		Fones        sql.NullString
		Gender       sql.NullString
		BirthDay     sql.NullTime
		Note         sql.NullString
	}
	fields := `cst.id, s.name, s.license, s.level, s.grade, slug, wch, s.year, s.semester, classroom_id, s.created_at, s.updated_at,
	t.id, t.name, t.nick,	t.birth_day,	t.gender,	t.cpf,	t.fones,	t.email,	t.license,	t.note,	t.status, t.created_at, t.updated_at`

	stmt, err := e.db.Prepare(`SELECT ` + fields + `
		FROM classrooms_subjects_teachers cst
		LEFT JOIN subjects s ON cst.subject_id = s.id
		LEFT JOIN teachers t ON cst.teacher_id = t.id
		WHERE cst.id = $1
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
		&subject.Note,
		&subject.CH,
		&subject.Year,
		&ip.Semester,
		&ip.Classroom_id,
		&subject.CreatedAt,
		&subject.UpdatedAt,
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

	subject.CH = int(ip.CH.Int32)
	subject.Semester = ip.Semester.String
	subject.Teacher = &teacher

	stmt.Close()

	return &subject, nil
}
