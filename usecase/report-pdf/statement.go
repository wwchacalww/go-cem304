package reportpdf

import (
	"wwchacalww/go-cem304/domain/model"

	"github.com/jung-kurt/gofpdf"
)

type InputStatementSchooling struct {
	Student      model.StudentInterface
	Matricula    string
	FathersName  string
	MothersName  string
	Address      string
	Phones       string
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
	Phones       string `json:"fones"`
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

	pdf.SetFont("Roboto-Bold", "", 12)
	// Cell(float w [, float h [, string txt [, mixed border [, int ln [, string align [, boolean fill [, mixed link]]]]]]])
	pdf.Image("pdf/logoGDF.png", 10, 10, 20, 23, false, "", 0, "")
	txt := "GDF - SECRETARIA DE ESTADO DE EDUCAÇÃO DO DISTRITO FEDERAL"
	pdf.CellFormat(190, 8, txt, "0", 1, "C", false, 0, "")

	err := pdf.OutputFileAndClose("pdf/statement-" + input.Student.GetID() + ".pdf")
	return err
}
