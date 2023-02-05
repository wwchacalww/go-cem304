package db_test

import (
	"database/sql"
	"log"
	"testing"
	"time"
	"wwchacalww/go-cem304/adapters/db"
	"wwchacalww/go-cem304/domain/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var dB *sql.DB

func setUpClassTest() {
	dB, _ = sql.Open("sqlite3", ":memory:")
	createClassroomTable(dB)
	createClass(dB)
}

func createClassroomTable(db *sql.DB) {
	table := `CREATE TABLE classrooms (
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
}

func createClass(db *sql.DB) {
	insert := `insert into classrooms values
	("id-1", "1º ano A - Vespertino", "1º ano", "Ensino Médio", "vespertino", "Descrição da turma", "0 alunos", "2023", 1, $1, $2),
	("id-2", "1º ano B - Vespertino", "1º ano", "Ensino Médio", "vespertino", "Descrição da turma", "0 alunos", "2023", 1, $1, $2);`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec(time.Now(), time.Now().Add(1*time.Hour))
}

func TestClassroom_Get(t *testing.T) {
	setUpClassTest()
	defer dB.Close()
	classDB := db.NewClassroomDB(dB)
	class, err := classDB.FindById("id-1")
	require.Nil(t, err)
	require.Equal(t, "1º ano A - Vespertino", class.GetName())
}

func TestClassroom_FindName(t *testing.T) {
	setUpClassTest()
	defer dB.Close()
	classDB := db.NewClassroomDB(dB)
	class, err := classDB.FindByName("Vespertino")
	require.Nil(t, err)
	require.Equal(t, len(class), 2)
}

func TestClassroom_List(t *testing.T) {
	setUpClassTest()
	defer dB.Close()
	classDB := db.NewClassroomDB(dB)
	class, err := classDB.List("2023")
	require.Nil(t, err)
	require.Equal(t, len(class), 2)
}

func TestClassroom_Create(t *testing.T) {
	setUpClassTest()
	defer dB.Close()
	classDB := db.NewClassroomDB(dB)

	class := model.NewClassroom()
	class.Name = "1º ano C - Vespertino"
	class.Level = "1º ano"
	class.Grade = "Ensino Médio"
	class.Shift = "Vespertino"
	class.Year = "2023"

	err := classDB.Create(class)
	require.Nil(t, err)

	classFind, err := classDB.FindById(class.GetID())
	require.Nil(t, err)
	require.Equal(t, classFind.GetName(), class.GetName())
	require.True(t, classFind.GetStatus())

	classDisable, err := classDB.Disable(class.GetID())
	require.Nil(t, err)
	require.False(t, classDisable.GetStatus())

	classEnable, err := classDB.Enable(class.GetID())
	require.Nil(t, err)
	log.Println(classEnable.GetANNE())
	require.True(t, classEnable.GetStatus())

	changeANNE, err := classDB.ANNE(class.GetID(), "3 alunos ANNE (1 TGD, 2 TDHA)")
	require.Nil(t, err)
	log.Println(changeANNE.GetANNE())
	require.Equal(t, changeANNE.GetANNE(), "3 alunos ANNE (1 TGD, 2 TDHA)")
}
