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
	txt = "DECLARAÇÃO DE ESCOLARIDADE - GERAL"
	pdf.CellFormat(190, 8, txt, "0", 1, "C", false, 0, "")
	pdf.Ln(6)

	txt = "DADOS DA INSTITUIÇÃO EDUCACIONAL"
	pdf.SetFont("Roboto-Bold", "", 10)
	pdf.CellFormat(190, 6, txt, "1", 1, "L", false, 0, "")
	txt = "Instituição Educacional: CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA - INEP: 53009029"
	pdf.SetFont("Roboto", "", 10)
	pdf.CellFormat(190, 5, txt, "RL", 1, "L", false, 0, "")
	txt = "Portaria: PORTARIA Nº 03 DE 12/01/2004-SEDF"
	pdf.CellFormat(190, 5, txt, "RL", 1, "L", false, 0, "")
	txt = "Endereço: QN 54 Conjunto 4 LOTE 01"
	pdf.CellFormat(190, 5, txt, "RL", 1, "L", false, 0, "")
	txt = "Bairro: Samambaia Sul (Samambaia) - Samambaia-DF - CEP: 7256-004"
	pdf.CellFormat(190, 5, txt, "RL", 1, "L", false, 0, "")
	txt = "Telefone: (61) 39017718"
	pdf.CellFormat(190, 5, txt, "RL", 1, "L", false, 0, "")

	err := pdf.OutputFileAndClose("pdf/statement-" + input.Student.GetID() + ".pdf")
	return err
}
