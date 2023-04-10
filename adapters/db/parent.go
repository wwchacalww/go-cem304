package db

import (
	"database/sql"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type ParentDB struct {
	db *sql.DB
}

func NewParentDB(db *sql.DB) *ParentDB {
	return &ParentDB{db: db}
}

func (p *ParentDB) Create(input repository.ParentInput) (model.ParentInterface, error) {
	parent := model.NewParent(input.Name, input.Relationship)
	parent.Email = input.Email
	parent.CPF = input.CPF
	parent.Fones = input.Fones
	parent.Responsible = input.Responsible
	parent.CPF = input.CPF

	_, err := parent.IsValid()
	if err != nil {
		return nil, err
	}

	stmt, err := p.db.Prepare(`INSERT INTO p (
		id,
		name,
		cpf,
		email,
		fones
	) values ($1, $2, $3, $4, $5)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		parent.GetID(),
		parent.GetName(),
		parent.GetCPF(),
		parent.GetEmail(),
		parent.GetFones(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	pvStmt, err := p.db.Prepare(`INSERT INTO parents_students (
		parent_id,
		student_id,
		relationship,
		responsible
	) values ($1, $2, $3, $4)`)
	if err != nil {
		return nil, err
	}

	_, err = pvStmt.Exec(
		parent.ID,
		input.Student_id,
		parent.Relationship,
		parent.Responsible,
	)
	if err != nil {
		return nil, err
	}
	err = pvStmt.Close()
	if err != nil {
		return nil, err
	}
	return parent, nil
}

func (p *ParentDB) FindById(id string) (model.ParentInterface, error) {
	var parent model.Parent
	var students []model.StudentInterface
	fields := "parents_id, p.name, p.cpf, p.email, p.fones, relationship, student_id, s.name, birth_day, gender, anne, note, ieducar, educa_df, status, address, city, cep, s.fones, s.cpf, created_at, updated_at"

	rows, err := p.db.Query(`SELECT `+fields+` 
		FROM parents_students ps
		LEFT JOIN parents p ON ps.parent_id = p.id
		LEFT JOIN students s ON ps.student_id = s.id
		WHERE ps.parent_id=$1
	`, id)
	defer rows.Close()
	for rows.Next() {
		var std model.Student
		err = rows.Scan(
			&parent.ID,
			&parent.Name,
			&parent.CPF,
			&parent.Email,
			&parent.Fones,
			&parent.Relationship,
			&std.ID,
			&std.Name,
			&std.BirthDay,
			&std.Gender,
			&std.ANNE,
			&std.Note,
			&std.Educar,
			&std.EducaDF,
			&std.Status,
			&std.Address,
			&std.City,
			&std.CEP,
			&std.Fones,
			&std.CPF,
			&std.CreatedAt,
			&std.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &std)
	}
	parent.Students = append(parent.Students, students...)

	return &parent, nil
}

func (p *ParentDB) FindByCPF(cpf string) (model.ParentInterface, error) {
	var parent model.Parent
	var students []model.StudentInterface
	fields := "parents_id, p.name, p.cpf, p.email, p.fones, relationship, student_id, s.name, birth_day, gender, anne, note, ieducar, educa_df, status, address, city, cep, s.fones, s.cpf, created_at, updated_at"

	rows, err := p.db.Query(`SELECT `+fields+` 
		FROM parents_students ps
		LEFT JOIN parents p ON ps.parent_id = p.id
		LEFT JOIN students s ON ps.student_id = s.id
		WHERE p.cpf=$1
	`, cpf)
	defer rows.Close()
	for rows.Next() {
		var std model.Student
		err = rows.Scan(
			&parent.ID,
			&parent.Name,
			&parent.CPF,
			&parent.Email,
			&parent.Fones,
			&parent.Relationship,
			&std.ID,
			&std.Name,
			&std.BirthDay,
			&std.Gender,
			&std.ANNE,
			&std.Note,
			&std.Educar,
			&std.EducaDF,
			&std.Status,
			&std.Address,
			&std.City,
			&std.CEP,
			&std.Fones,
			&std.CPF,
			&std.CreatedAt,
			&std.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &std)
	}
	parent.Students = append(parent.Students, students...)

	return &parent, nil
}

func (p *ParentDB) FindByName(name string) ([]model.ParentInterface, error) {
	var parents []model.ParentInterface
	var students []model.StudentInterface
	fields := "parents_id, p.name, p.cpf, p.email, p.fones, relationship, student_id, s.name, birth_day, gender, anne, note, ieducar, educa_df, status, address, city, cep, s.fones, s.cpf, created_at, updated_at"

	rows, err := p.db.Query(`SELECT `+fields+` 
		FROM parents_students ps
		LEFT JOIN parents p ON ps.parent_id = p.id
		LEFT JOIN students s ON ps.student_id = s.id
		WHERE p.name like $1
	`, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var std model.Student
		var parent model.Parent
		err = rows.Scan(
			&parent.ID,
			&parent.Name,
			&parent.CPF,
			&parent.Email,
			&parent.Fones,
			&parent.Relationship,
			&std.ID,
			&std.Name,
			&std.BirthDay,
			&std.Gender,
			&std.ANNE,
			&std.Note,
			&std.Educar,
			&std.EducaDF,
			&std.Status,
			&std.Address,
			&std.City,
			&std.CEP,
			&std.Fones,
			&std.CPF,
			&std.CreatedAt,
			&std.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &std)
		parent.Students = append(parent.Students, students...)
		parents = append(parents, &parent)
	}

	return parents, nil
}

func (p *ParentDB) ChangeFone(id, fones string) error {
	_, err := p.db.Exec("UPDATE parents SET fones=$1 WHERE id=$2", fones, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ParentDB) ChangeEmail(id, email string) error {
	_, err := p.db.Exec("UPDATE parents SET email=$1 WHERE id=$2", email, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ParentDB) ChangeCPF(id, cpf string) error {
	_, err := p.db.Exec("UPDATE parents SET cpf=$1 WHERE id=$2", cpf, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ParentDB) ChangeResponsible(id, student_id string, status bool) error {
	_, err := p.db.Exec("UPDATE parents_students SET responsible=$1 WHERE parent_id=$2 AND student_id=$3", status, id, student_id)
	if err != nil {
		return err
	}

	return nil
}
