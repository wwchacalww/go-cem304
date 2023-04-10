package db

import (
	"database/sql"
	"log"
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

func (std *StudentDB) Save(input repository.StudentInput) (model.StudentInterface, error) {
	student, err := std.FindByEducar(input.Educar)
	if err != nil { // new student
		student, err := model.NewStudent(input.Name, input.BirthDay, input.Educar)
		student.Gender = input.Gender
		student.ANNE = input.ANNE
		student.Note = input.Note
		student.EducaDF = input.EducaDF
		student.Address = input.Address
		student.City = input.City
		student.CEP = input.CEP
		student.Fones = input.Fones
		student.CPF = input.CPF

		if err != nil {
			return nil, err
		}

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
			classStmt.Close()
			student.Classroom = &class
		}

		if input.Mother != "" {
			mother := model.NewParent(input.Mother, "Mãe")
			if input.Mother == input.Responsible {
				mother.Responsible = true
			}
			student.Parents = append(student.Parents, mother)
		}

		if input.Father != "" {
			father := model.NewParent(input.Father, "Pai")
			if input.Father == input.Responsible {
				father.Responsible = true
			}
			student.Parents = append(student.Parents, father)
		}

		_, err = student.IsValid()
		if err != nil {
			return nil, err
		}
		newStd, err := std.create(student)
		if err != nil {
			return nil, err
		}
		return newStd, nil
	}

	NewStudent := model.Student{
		ID:        student.GetID(),
		Name:      student.GetName(),
		BirthDay:  student.GetBirthDay(),
		Gender:    input.Gender,
		ANNE:      input.ANNE,
		Note:      input.Note,
		Educar:    input.Educar,
		EducaDF:   input.EducaDF,
		Classroom: student.GetClassroom(),
		Status:    input.Status,
		Address:   input.Address,
		City:      input.City,
		CEP:       input.CEP,
		Fones:     input.Fones,
		CPF:       input.CPF,
	}

	if input.Mother != "" && student.GetParents() == nil {
		mother := model.NewParent(input.Mother, "Mãe")
		if input.Mother == input.Responsible {
			mother.Responsible = true
		}
		NewStudent.Parents = append(NewStudent.Parents, mother)
	}

	if input.Father != "" && student.GetParents() == nil {
		father := model.NewParent(input.Father, "Pai")
		if input.Father == input.Responsible {
			father.Responsible = true
		}
		NewStudent.Parents = append(NewStudent.Parents, father)
	}
	result, err := std.update(&NewStudent)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (std *StudentDB) FindById(id string) (model.StudentInterface, error) {
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

	parentRows, err := std.db.Query(`SELECT p.id, p.name, cpf, email, fones, relationship, responsible 
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

func (std *StudentDB) FindByEducar(educar int64) (model.StudentInterface, error) {
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

	parentRows, err := std.db.Query(`SELECT p.id, p.name, cpf, email, fones, relationship, responsible 
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
	var ip struct {
		Address sql.NullString
		City    sql.NullString
		CEP     sql.NullString
		Fones   sql.NullString
		CPF     sql.NullString
	}
	std_fields := "id, name, birth_day, gender, anne, note, ieducar, educa_df, classroom_id, status, address, city, cep, fones, cpf, created_at, students.updated_at"
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
			&ip.Address,
			&ip.City,
			&ip.CEP,
			&ip.Fones,
			&ip.CPF,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		student.Address = ip.Address.String
		student.City = ip.City.String
		student.CEP = ip.CEP.String
		student.Fones = ip.Fones.String
		student.CPF = ip.CPF.String
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

func (std *StudentDB) FindByParent(name string) ([]model.StudentInterface, error) {
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
	var ip struct {
		Address sql.NullString
		City    sql.NullString
		CEP     sql.NullString
		Fones   sql.NullString
		CPF     sql.NullString
	}
	std_fields := "id, name, birth_day, gender, anne, note, ieducar, educa_df, status, address, city, cep, fones, cpf, created_at, students.updated_at"
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
			&ip.Address,
			&ip.City,
			&ip.CEP,
			&ip.Fones,
			&ip.CPF,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		student.Classroom = &class
		student.Address = ip.Address.String
		student.City = ip.City.String
		student.CEP = ip.CEP.String
		student.Fones = ip.Fones.String
		student.CPF = ip.CPF.String

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

func (std *StudentDB) ChangeClassroom(id, classroom_id string) error {
	_, err := std.db.Exec("UPDATE students SET classroom_id=$1 WHERE id=$2", classroom_id, id)
	if err != nil {
		return err
	}

	return nil
}

func (std *StudentDB) create(student model.StudentInterface) (model.StudentInterface, error) {
	classIsSet := false
	if student.GetClassroom() != nil {
		classIsSet = true
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
		address, city, cep, fones, cpf,
		created_at,
		updated_at
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12, $13, $14, $15, $16, $17)`)
	if err != nil {
		return nil, err
	}
	if classIsSet {
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
			student.GetAddress(),
			student.GetCity(),
			student.GetCEP(),
			student.GetFones(),
			student.GetCPF(),
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
			student.GetAddress(),
			student.GetCity(),
			student.GetCEP(),
			student.GetFones(),
			student.GetCPF(),
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

	if student.GetParents() != nil {
		for i, p := range student.GetParents() {
			log.Println(p, i)
			_, err = std.db.Exec("INSERT INTO parents (id,name) values ( $1, $2)", p.GetID(), p.GetName())
			if err != nil {
				return nil, err
			}
			_, err = std.db.Exec(`INSERT INTO parents_students (
				parent_id,
				student_id,
				relationship,
				responsible
				) values ( 
					$1, $2, $3, $4
				)`, p.GetID(), student.GetID(), p.GetRelationship(), p.GetResponsible())
			if err != nil {
				return nil, err
			}
		}
	}
	return student, nil
}

func (std *StudentDB) update(student model.StudentInterface) (model.StudentInterface, error) {
	_, err := std.db.Exec(`UPDATE students SET 
		birth_day=$1, 
		gender=$2, 
		anne=$3, 
		note=$4, 
		ieducar=$5, 
		educa_df=$6, 
		classroom_id=$7, 
		status=$8, 
		address=$9, 
		city=$10, 
		cep=$11, 
		fones=$12, 
		cpf=$13, 
		updated_at=$14
		WHERE id=$15`,
		student.GetBirthDay(),
		student.GetGender(),
		student.GetANNE(),
		student.GetNote(),
		student.GetEducar(),
		student.GetEducaDF(),
		student.GetClassroom().GetID(),
		student.GetStatus(),
		student.GetAddress(),
		student.GetCity(),
		student.GetCEP(),
		student.GetFones(),
		student.GetCPF(),
		time.Now(),
		student.GetID(),
	)
	if err != nil {
		return nil, err
	}

	if student.GetParents() != nil {
		for i, p := range student.GetParents() {
			log.Println(p, i)
			_, err = std.db.Exec("INSERT INTO parents (id,name) values ( $1, $2)", p.GetID(), p.GetName())
			if err != nil {
				return nil, err
			}
			_, err = std.db.Exec(`INSERT INTO parents_students (
				parent_id,
				student_id,
				relationship,
				responsible
				) values ( 
					$1, $2, $3, $4
				)`, p.GetID(), student.GetID(), p.GetRelationship(), p.GetResponsible())
			if err != nil {
				return nil, err
			}
		}
	}

	return student, nil
}
