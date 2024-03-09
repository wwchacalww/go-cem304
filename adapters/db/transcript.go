package db

import (
	"database/sql"
	"fmt"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/repository"
)

type TranscriptDB struct {
	db *sql.DB
}

func NewTranscriptDB(db *sql.DB) *TranscriptDB {
	return &TranscriptDB{db: db}
}

func (t *TranscriptDB) Save(input repository.TranscriptInput) (model.TranscriptInterface, error) {
	std, err := t.findByRA(input.EducaDF)
	if err != err {
		return nil, err
	}

	transcript_id, err := t.checkExistsTranscript(std.GetID(), input.Subject, input.Level)
	if err != err {
		return nil, err
	}

	if transcript_id <= 0 {
		transcript, err := t.create(input, std)
		if err != err {
			return nil, err
		}
		return transcript, err
	} else {
		err = t.update(transcript_id, input)
		if err != err {
			return nil, err
		}
		transcript := model.NewTranscript(
			input.Subject,
			input.Result,
			input.Note,
			input.Level,
			input.Formation,
			input.Absences,
			input.Workload,
			input.Year,
		)
		transcript.ID = transcript_id
		transcript.Student = std

		return transcript, nil
	}
}

func (t *TranscriptDB) findByRA(ra string) (model.StudentInterface, error) {
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
	stmt, err := t.db.Prepare("SELECT " + std_fields + " from students WHERE educa_df = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(ra).Scan(
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

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (t *TranscriptDB) checkExistsTranscript(student_id, subject, level string) (int, error) {
	transcript_id := 0
	stmt, err := t.db.Prepare("SELECT id FROM transcripts WHERE student_id = $1 AND subject = $2 AND level = $3")
	if err != nil {
		return 0, err
	}
	err = stmt.QueryRow(student_id, subject, level).Scan(
		&transcript_id,
	)
	return transcript_id, err
}

func (t *TranscriptDB) create(input repository.TranscriptInput, student model.StudentInterface) (model.TranscriptInterface, error) {
	transcript := model.NewTranscript(
		input.Subject,
		input.Result,
		input.Note,
		input.Level,
		input.Formation,
		input.Absences,
		input.Workload,
		input.Year,
	)
	transcript.Student = student
	valid, _ := transcript.IsValid()
	if valid == false {
		return nil, fmt.Errorf("Transcript invalid")
	}
	stmt, err := t.db.Prepare(`INSERT INTO transcripts (
		student_id,
		subject,
		result,
		note,
		absences,
		workload,
		level,
		formation,
		year
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	_, err = stmt.Exec(
		transcript.GetStudent().GetID(),
		transcript.GetSubject(),
		transcript.GetResult(),
		transcript.GetNote(),
		transcript.GetAbsences(),
		transcript.GetWorkload(),
		transcript.GetLevel(),
		transcript.GetFormation(),
		transcript.GetYear(),
	)

	if err != nil {
		return nil, err
	}

	return transcript, nil
}

func (t *TranscriptDB) update(transcript_id int, input repository.TranscriptInput) error {
	_, err := t.db.Exec(`UPDATE transcripts SET 
		subject=$1,
		result=$2,
		note=$3,
		level=$4,
		formation=$5,
		absences=$6,
		workload=$7,
		year=$8
		WHERE id=$9`,
		input.Subject,
		input.Result,
		input.Note,
		input.Level,
		input.Formation,
		input.Absences,
		input.Workload,
		input.Year,
		transcript_id,
	)
	return err
}
