package check

import (
	"log"
	"strconv"
	"wwchacalww/go-cem304/domain/model"
	"wwchacalww/go-cem304/domain/utils"
)

type OutputCheckStudentsInClass struct {
	Classroom string `json:"classroom_name"`
	Result    string `json:"result"`
}

func CheckStudentsInClass(list []utils.InputCheckStudentInClass, classroom_id string, class []model.StudentInterface) (OutputCheckStudentsInClass, error) {
	var output OutputCheckStudentsInClass
	ctxt := ""
	rtxt := class[0].GetClassroom().GetName()
	ctrl := len((list)) - len(class)
	log.Println(ctrl)
	if ctrl == 0 {
		ctxt = ctxt + "As turmas tem o mesmo número de alunos, agora vou verificar se conferer os alunos."
		for _, study := range class {
			res := study.GetName() + " não encontrado na lista... \n\r"
			for _, listStudy := range list {
				if listStudy.Educar == study.GetEducar() {
					res = study.GetName() + " OK... \n\r"
				}
			}
			ctxt = ctxt + res
		}
	}
	if ctrl > 0 {
		ctxt = ctxt + "A lista tem mais alunos que a turma, conferindo alunos...\n\r "
		for _, listStudy := range list {
			res := strconv.FormatInt(listStudy.Educar, 10) + " não encontrado na turma... \n\r"
			for _, study := range class {
				if listStudy.Educar == study.GetEducar() {
					res = study.GetName() + " OK... \n\r"
				}
			}
			ctxt = ctxt + res
		}
	}

	if ctrl < 0 {
		ctxt = ctxt + "A lista esta com menos alunos que a turma, conferindo alunos...\n\r "
		for _, study := range class {
			res := study.GetName() + " não encontrado na lista... \n\r"
			for _, listStudy := range list {
				if listStudy.Educar == study.GetEducar() {
					res = study.GetName() + " OK... \n\r"
				}
			}
			ctxt = ctxt + res
		}
	}

	output.Classroom = rtxt
	output.Result = ctxt
	return output, nil
}
