package reportpdf

import (
	"strconv"
	"strings"
	"time"
	"wwchacalww/go-cem304/domain/model"

	"github.com/jung-kurt/gofpdf"
)

type InputStatementSchooling struct {
	Student      model.StudentInterface
	Matricula    string
	FathersName  string
	MothersName  string
	Address      string
	CEP          string
	PhoneOne     string `json:"fone_one"`
	PhoneTwo     string `json:"fone_two"`
	PhoneThree   string `json:"fone_three"`
	Nationality  string
	Birthplace   string
	CPF          string
	Neighborhood string
	City         string
	ParentCPF    string
}

type InputSS struct {
	Educar       int64  `json:"educar"`
	Matricula    string `json:"matricula"`
	FathersName  string `json:"nome_pai"`
	MothersName  string `json:"nome_mae"`
	Address      string `json:"endereco"`
	CEP          string `json:"cep"`
	PhoneOne     string `json:"fone_one"`
	PhoneTwo     string `json:"fone_two"`
	PhoneThree   string `json:"fone_three"`
	Nationality  string `json:"nacionalidade"`
	Birthplace   string `json:"naturalidade"`
	CPF          string `json:"cpf"`
	Neighborhood string `json:"bairro"`
	City         string `json:"cidade"`
	ParentCPF    string `json:"cpf_responsavel"`
}

func StatementSchooling(input InputStatementSchooling) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Roboto", "", "pdf/Roboto-Regular.ttf")
	pdf.AddUTF8Font("Roboto-Bold", "", "pdf/Roboto-Bold.ttf")

	pdf.SetFont("Roboto", "", 12)
	// Cell(float w [, float h [, string txt [, mixed border [, int ln [, string align [, boolean fill [, mixed link]]]]]]])
	pdf.Image("pdf/logoGDF.png", 10, 10, 20, 23, false, "", 0, "")
	txt := "GDF - SECRETARIA DE ESTADO DE EDUCAÇÃO DO DISTRITO FEDERAL"
	pdf.CellFormat(190, 6, txt, "0", 1, "C", false, 0, "")
	pdf.SetFont("Roboto", "", 8)
	txt = "CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA"
	pdf.CellFormat(190, 4.5, txt, "0", 1, "C", false, 0, "")
	pdf.SetFont("Roboto", "", 10)
	txt = "CRE - Samambaia"
	pdf.CellFormat(190, 4.5, txt, "0", 1, "C", false, 0, "")
	txt = "QN 304 Conjunto 4 LOTE 01"
	pdf.CellFormat(190, 6, txt, "0", 1, "C", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 14)
	txt = "DECLARAÇÃO DE ESCOLARIDADE - PASSE ESTUDANTIL"
	pdf.CellFormat(190, 8, txt, "0", 1, "C", false, 0, "")
	pdf.Ln(3)

	txt = "DADOS DA INSTITUIÇÃO EDUCACIONAL"
	pdf.SetFont("Roboto-Bold", "", 10)
	pdf.CellFormat(190, 6, txt, "1", 1, "L", false, 0, "")
	txt = "Instituição Educacional: CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA - INEP: 53009029"
	pdf.SetFont("Roboto", "", 10)
	pdf.CellFormat(190, 6, txt, "RL", 1, "L", false, 0, "")
	txt = "Portaria: PORTARIA Nº 03 DE 12/01/2004-SEDF"
	pdf.CellFormat(190, 6, txt, "RL", 1, "L", false, 0, "")
	txt = "Endereço: QN 54 Conjunto 4 LOTE 01"
	pdf.CellFormat(190, 6, txt, "RL", 1, "L", false, 0, "")
	txt = "Bairro: Samambaia Sul (Samambaia) - Samambaia-DF - CEP: 7256-004"
	pdf.CellFormat(190, 6, txt, "RL", 1, "L", false, 0, "")
	txt = "Telefone: (61) 39017718"
	pdf.CellFormat(190, 6, txt, "RBL", 1, "L", false, 0, "")
	pdf.Ln(3)

	txt = "DADOS DO ESTUDANTE"
	pdf.SetFont("Roboto-Bold", "", 10)
	pdf.CellFormat(135, 6, txt, "1", 0, "L", false, 0, "")
	txt = "Código do Estudante: " + strconv.FormatInt(input.Student.GetEducar(), 10)
	pdf.CellFormat(55, 6, txt, "1", 1, "L", false, 0, "")
	txt = "Nome do(a) Estudante: " + input.Student.GetName()
	pdf.SetFont("Roboto-Bold", "", 10)
	pdf.CellFormat(135, 6, txt, "1", 0, "L", false, 0, "")
	pdf.SetFont("Roboto", "", 10)
	txt = "Matrícula: " + input.Matricula
	pdf.CellFormat(55, 6, txt, "1", 1, "L", false, 0, "")
	txt = "Data de Nascimento: " + input.Student.GetBirthDay().Format("02/01/2006")
	pdf.CellFormat(55, 6, txt, "1", 0, "L", false, 0, "")
	txt = "Sexo: " + strings.Title(input.Student.GetGender())
	pdf.CellFormat(35, 6, txt, "1", 0, "L", false, 0, "")
	txt = "Nacionalidade: Brasileira"
	pdf.CellFormat(45, 6, txt, "1", 0, "L", false, 0, "")
	txt = "Naturalidade: " + input.Birthplace
	pdf.CellFormat(55, 6, txt, "1", 1, "L", false, 0, "")

	txt = "CPF: " + input.CPF
	pdf.CellFormat(55, 6, txt, "1", 0, "L", false, 0, "")
	txt = "RG: "
	pdf.CellFormat(35, 6, txt, "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 6, "", "1", 1, "L", false, 0, "")

	pdf.SetFont("Roboto-Bold", "", 10)
	pdf.CellFormat(190, 6, "FILIAÇÃO", "1", 1, "L", false, 0, "")
	pdf.SetFont("Roboto", "", 10)

	txt = "Filiação 1: " + input.FathersName
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Filiação 2: " + input.MothersName
	pdf.CellFormat(190, 6, txt, "LBR", 1, "L", false, 0, "")
	pdf.Ln(3)

	pdf.SetFont("Roboto-Bold", "", 10)
	pdf.CellFormat(190, 6, "DADOS DO RESPONSÁVEL", "1", 1, "L", false, 0, "")
	pdf.SetFont("Roboto", "", 10)

	txt = "Responsável: " + input.MothersName
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Endereço: " + input.Address
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Bairro: " + input.Neighborhood
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Cidade: " + input.City + " CEP: " + input.CEP
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Telefone1: " + input.PhoneOne
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Telefone2: " + input.PhoneTwo
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Celular: " + input.PhoneThree
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "CPF Responsável: " + input.ParentCPF
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	pdf.CellFormat(190, 3, "", "LBR", 1, "L", false, 0, "")
	pdf.Ln(3)

	pdf.SetFont("Roboto-Bold", "", 10)
	pdf.CellFormat(190, 6, "DADOS DO CURSO", "1", 1, "L", false, 0, "")
	pdf.SetFont("Roboto", "", 10)

	txt = "Modalidade de Ensino: Ensino Médio" + input.Student.GetClassroom().GetGrade()
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Série: " + input.Student.GetClassroom().GetLevel() + " Série de Referência: " + input.Student.GetClassroom().GetLevel()
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Turma: " + input.Student.GetClassroom().GetName()
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Período Letivo: Anual / 2023"
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	txt = "Data de Início: 13/02/2023 Data de Término: 22/12/2023"
	pdf.CellFormat(190, 6, txt, "LR", 1, "L", false, 0, "")
	if input.Student.GetClassroom().GetShift() == "matutino" {
		txt = "Dia das Aulas: Horário: 07:15:00 às 12:15:00"
	}
	if input.Student.GetClassroom().GetShift() == "vespertino" {
		txt = "Dia das Aulas: Horário: 13:15:00 às 18:15:00"
	}
	if input.Student.GetClassroom().GetShift() == "noturno" {
		txt = "Dia das Aulas: Horário: 19:00:00 às 23:00:00"
	}
	pdf.CellFormat(190, 6, txt, "LBR", 1, "L", false, 0, "")
	pdf.Ln(3)

	pdf.SetFont("Roboto-Bold", "", 10)
	pdf.CellFormat(190, 6, "INFORMAÇÕES COMPLEMENTARES", "1", 1, "L", false, 0, "")
	pdf.SetFont("Roboto", "", 10)
	pdf.CellFormat(190, 18, "", "LBR", 1, "L", false, 0, "")
	txt = "ESTE DOCUMENTO NÃO É VALIDO PARA MATRÍCULA"
	pdf.CellFormat(190, 4, txt, "0", 1, "L", false, 0, "")
	day := time.Now().Day()
	txt = "Local e Data: Samambaia-DF, " + strconv.Itoa(day) + " de fevereiro de 2023"
	pdf.CellFormat(190, 4, txt, "0", 1, "L", false, 0, "")

	pdf.Ln(5)
	txt = "DIRETOR"
	pdf.CellFormat(95, 4, txt, "0", 0, "C", false, 0, "")
	txt = "SECRETÁRIO ESCOLAR"
	pdf.CellFormat(95, 4, txt, "0", 1, "C", false, 0, "")
	pdf.SetFont("Roboto", "", 8)
	txt = "(ASSINATURA E CARIMBO)"
	pdf.CellFormat(95, 4, txt, "0", 0, "C", false, 0, "")
	pdf.CellFormat(95, 4, txt, "0", 1, "C", false, 0, "")

	err := pdf.OutputFileAndClose("pdf/statement-" + input.Student.GetID() + ".pdf")
	return err
}
