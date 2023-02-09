package reportpdf

import (
	"strconv"
	"wwchacalww/go-cem304/domain/model"

	"github.com/jung-kurt/gofpdf"
)

func StudentsInClass(class model.ClassroomInterface) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Roboto", "", "pdf/Roboto-Regular.ttf")
	pdf.AddUTF8Font("Roboto-Bold", "", "pdf/Roboto-Bold.ttf")

	pdf.SetFont("Roboto-Bold", "", 12)
	// Cell(float w [, float h [, string txt [, mixed border [, int ln [, string align [, boolean fill [, mixed link]]]]]]])
	pdf.Image("pdf/logoGDF.png", 10, 10, 20, 23, false, "", 0, "")
	pdf.Ln(5)
	txt := "CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA"
	pdf.CellFormat(190, 8, txt, "0", 1, "C", false, 0, "")
	txt = "Lista de Alunos - " + class.GetName()
	pdf.CellFormat(190, 8, txt, "0", 1, "C", false, 0, "")
	pdf.Ln(5)
	pdf.SetFont("Roboto-Bold", "", 8)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(200, 200, 200)
	pdf.SetFillColor(100, 100, 100)
	pdf.CellFormat(20, 6, "N.º", "1", 0, "C", true, 0, "")
	pdf.CellFormat(170, 6, "    Nome do Aluno", "1", 1, "L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFillColor(225, 225, 225)

	for i, student := range class.GetStudents() {
		var n string
		fill := true
		if i%2 == 0 {
			fill = false
		}
		n = strconv.Itoa(i + 1)
		if i < 9 {
			n = "0" + strconv.Itoa(i+1)
		}
		pdf.CellFormat(20, 4.6, n, "1", 0, "C", fill, 0, "")
		pdf.CellFormat(170, 4.6, "  "+student.GetName(), "1", 1, "L", fill, 0, "")
	}

	ts := strconv.Itoa(len(class.GetStudents()))
	pdf.Ln(1)
	pdf.CellFormat(190, 4.6, "Total de "+ts+" alunos", "0", 1, "R", false, 0, "")

	err := pdf.OutputFileAndClose("pdf/hello.pdf")
	return err
}
