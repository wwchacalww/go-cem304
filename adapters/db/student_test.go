package db_test

import (
	"database/sql"
	"log"
	"testing"
	"time"
	"wwchacalww/go-cem304/adapters/db"
	"wwchacalww/go-cem304/domain/repository"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var dB_student_test *sql.DB

func setUpStudentTest() {
	dB_student_test, _ = sql.Open("sqlite3", ":memory:")
	createStudentTable(dB_student_test)
	createstudent(dB_student_test)
	createNewClass(dB_student_test)
}

func createStudentTable(db *sql.DB) {
	table := `CREATE TABLE students (
		"id" string,
		"name" string,
		"birth_day" Date,
		"gender" string,
		"anne" string,
		"note" string,
		"ieducar" string,
		"educa_df" string,
		"classroom_id" string,
		"status" boolean,
		"created_at" Date,
		"updated_at" Date
	);
	`
	tableClass := `CREATE TABLE classrooms (
		"id" string,
		"name" string,
		"level" string,
		"grade" string,
		"shift" string,
		"description" string,
		"ANNE" string,
		"year" string,
		"status" boolean,
		"created_at" Date,
		"updated_at" Date
	);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
	db.Exec(tableClass)
}

func createstudent(db *sql.DB) {
	insert := `insert into students values
	("id-1", "Fulana de Tal Lima", $1, "Feminino", "", "Notações", 1101, "cod-educa-df-2023", "class-1", 1, $2, $3),
	("id-2", "Beltrano da Silva Lima", $1, "Masculino", "", "Notações", 1102, "cod-educa-df-2023", "class-2", 1, $2, $3);`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec(time.Now(), time.Now().Add(1*time.Hour), time.Now().Add(2*time.Hour))
}

func createNewClass(db *sql.DB) {
	insert := `insert into classrooms values
	("class-1", "1º ano A - Vespertino", "1º ano", "Ensino Médio", "vespertino", "Descrição da turma", "0 alunos", "2023", 1, $1, $2),
	("class-2", "1º ano B - Vespertino", "1º ano", "Ensino Médio", "vespertino", "Descrição da turma", "0 alunos", "2023", 1, $1, $2);`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec(time.Now(), time.Now().Add(1*time.Hour))
}

func TestStudent_Get(t *testing.T) {
	setUpStudentTest()
	defer dB_student_test.Close()
	studentDB := db.NewStudentDB(dB_student_test)
	student, err := studentDB.FindById("id-1")
	require.Nil(t, err)
	require.Equal(t, "Fulana de Tal Lima", student.GetName())
	require.Equal(t, "1º ano A - Vespertino", student.GetClassroom().GetName())
}

func TestStudent_FindName(t *testing.T) {
	setUpStudentTest()
	defer dB_student_test.Close()
	studentDB := db.NewStudentDB(dB_student_test)
	students, err := studentDB.FindByName("Lima")
	require.Nil(t, err)
	require.Equal(t, students[0].GetName(), "Beltrano da Silva Lima")
	require.Equal(t, students[1].GetClassroom().GetName(), "1º ano A - Vespertino")
}

func TestStudent_List(t *testing.T) {
	setUpStudentTest()
	defer dB_student_test.Close()
	studentDB := db.NewStudentDB(dB_student_test)
	students, err := studentDB.List("class-1")
	require.Nil(t, err)
	require.Equal(t, len(students), 1)
	require.Equal(t, students[0].GetClassroom().GetName(), "1º ano A - Vespertino")
}

func TestStudent_Create(t *testing.T) {
	setUpStudentTest()
	defer dB_student_test.Close()
	studentDB := db.NewStudentDB(dB_student_test)

	input := repository.StudentInput{
		Name:        "Siclano da Silva Lima",
		BirthDay:    "29/11/2016",
		Gender:      "Masculino",
		Educar:      234234,
		ANNE:        "",
		Note:        "",
		EducaDF:     "",
		ClassroomID: "class-1",
	}

	student, err := studentDB.Create(input)
	require.Nil(t, err)
	// log.Println(student)

	studentFind, err := studentDB.FindById(student.GetID())
	require.Nil(t, err)
	require.Equal(t, studentFind.GetName(), student.GetName())
	require.True(t, studentFind.GetStatus())

	// studentDisable, err := studentDB.Disable(student.GetID())
	// require.Nil(t, err)
	// require.False(t, studentDisable.GetStatus())

	// studentEnable, err := studentDB.Enable(student.GetID())
	// require.Nil(t, err)
	// log.Println(studentEnable.GetANNE())
	// require.True(t, studentEnable.GetStatus())

	// changeANNE, err := studentDB.ANNE(student.GetID(), "3 alunos ANNE (1 TGD, 2 TDHA)")
	// require.Nil(t, err)
	// log.Println(changeANNE.GetANNE())
	// require.Equal(t, changeANNE.GetANNE(), "3 alunos ANNE (1 TGD, 2 TDHA)")
}
