package reportpdf

import (
	"log"
	"strconv"
	"wwchacalww/go-cem304/domain/model"

	"github.com/jung-kurt/gofpdf"
)

type ReportClassroom struct {
	Classroom string `json:"classroom"`
	Group     string `json:"group"`
	Quantity  int    `json:"quantity"`
	Female    int    `json:"female"`
	Male      int    `json:"male"`
	ANEE      int    `json:"anee"`
}

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

func ReportAllClass(classrooms []model.ClassroomInterface) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	for _, class := range classrooms {
		if len(class.GetStudents()) > 0 {
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
		}
	}

	err := pdf.OutputFileAndClose("pdf/all-classrooms.pdf")
	return err
}

func DiaryClass(class model.ClassroomInterface) error {
	pdf := gofpdf.New("L", "mm", "A4", "")
	ts := strconv.Itoa(len(class.GetStudents()))
	numberPage := "1"
	pdf.SetFooterFunc(func() {
		// Footer
		pdf.SetY(-15)
		pdf.CellFormat(130, 8, "Total de "+ts+" alunos", "0", 0, "L", false, 0, "")
		pdf.CellFormat(130, 8, "Página "+numberPage+" de 2", "0", 1, "R", false, 0, "")
	})
	pdf.AddPage()
	pdf.AddUTF8Font("Roboto", "", "pdf/Roboto-Regular.ttf")
	pdf.AddUTF8Font("Roboto-Bold", "", "pdf/Roboto-Bold.ttf")

	pdf.SetFont("Roboto-Bold", "", 12)
	// Cell(float w [, float h [, string txt [, mixed border [, int ln [, string align [, boolean fill [, mixed link]]]]]]])
	pdf.Image("pdf/logoGDF.png", 10, 10, 20, 23, false, "", 0, "")
	txt := "CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA"
	pdf.CellFormat(260, 8, txt, "0", 1, "C", false, 0, "")
	txt = "DIÁRIO PROVISÓRIO DE CLASSE - " + class.GetName()
	pdf.CellFormat(260, 8, txt, "0", 1, "C", false, 0, "")

	pdf.CellFormat(20, 8, "", "0", 0, "C", false, 0, "")
	pdf.CellFormat(60, 8, "ABRIL DE 2023", "0", 0, "C", false, 0, "")
	pdf.CellFormat(35, 8, "DISCIPLINA:", "0", 0, "R", false, 0, "")
	txt = "____________________________"
	pdf.CellFormat(55, 8, txt, "B", 0, "L", false, 0, "")
	pdf.CellFormat(35, 8, "PROFESSOR(A):", "0", 0, "R", false, 0, "")
	pdf.CellFormat(55, 8, txt, "B", 1, "L", false, 0, "")

	pdf.SetFont("Roboto-Bold", "", 8)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(200, 200, 200)
	pdf.SetFillColor(100, 100, 100)
	pdf.CellFormat(20, 6, "N.º", "1", 0, "C", true, 0, "")
	days := [9]string{"13", "14", "15", "16", "17", "23", "24", "27", "28"}

	pdf.CellFormat(114, 6, "    Nome do Aluno", "1", 0, "L", true, 0, "")
	for _, d := range days {
		pdf.CellFormat(6, 6, d, "1", 0, "C", true, 0, "")
	}
	pdf.CellFormat(90, 6, "Anotações", "1", 1, "C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFillColor(225, 225, 225)
	numberPage = "1"
	for i, student := range class.GetStudents() {
		if i == 30 {
			pdf.AddPage()

			pdf.SetFont("Roboto-Bold", "", 12)
			// Cell(float w [, float h [, string txt [, mixed border [, int ln [, string align [, boolean fill [, mixed link]]]]]]])
			pdf.Image("pdf/logoGDF.png", 10, 10, 20, 23, false, "", 0, "")
			txt := "CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA"
			pdf.CellFormat(260, 8, txt, "0", 1, "C", false, 0, "")
			txt = "DIÁRIO PROVISÓRIO DE CLASSE - " + class.GetName()
			pdf.CellFormat(260, 8, txt, "0", 1, "C", false, 0, "")

			pdf.CellFormat(20, 8, "", "0", 0, "C", false, 0, "")
			pdf.CellFormat(60, 8, "ABRIL DE 2023", "0", 0, "C", false, 0, "")
			pdf.CellFormat(35, 8, "DISCIPLINA:", "0", 0, "R", false, 0, "")
			txt = "____________________________"
			pdf.CellFormat(55, 8, txt, "B", 0, "L", false, 0, "")
			pdf.CellFormat(35, 8, "PROFESSOR(A):", "0", 0, "R", false, 0, "")
			pdf.CellFormat(55, 8, txt, "B", 1, "L", false, 0, "")

			pdf.SetFont("Roboto-Bold", "", 8)
			pdf.SetTextColor(255, 255, 255)
			pdf.SetDrawColor(200, 200, 200)
			pdf.SetFillColor(100, 100, 100)
			pdf.CellFormat(20, 6, "N.º", "1", 0, "C", true, 0, "")

			pdf.CellFormat(114, 6, "    Nome do Aluno", "1", 0, "L", true, 0, "")
			for _, d := range days {
				pdf.CellFormat(6, 6, d, "1", 0, "C", true, 0, "")
			}
			pdf.CellFormat(90, 6, "Anotações", "1", 1, "C", true, 0, "")
			pdf.SetTextColor(0, 0, 0)

			pdf.SetFillColor(225, 225, 225)
			numberPage = "2"
		}
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
		pdf.CellFormat(114, 4.6, "  "+student.GetName(), "1", 0, "L", fill, 0, "")
		for _, d := range days {
			if d == "asflkj" {
				d = ""
			}
			pdf.CellFormat(6, 4.6, "", "1", 0, "C", fill, 0, "")
		}
		pdf.CellFormat(90, 4.6, "", "1", 1, "C", fill, 0, "")
	}

	err := pdf.OutputFileAndClose("pdf/diary_class.pdf")
	return err
}

func DiaryAllClass(classrooms []model.ClassroomInterface) error {
	pdf := gofpdf.New("L", "mm", "A4", "")
	numberPage := "1"
	nPages := "1"
	pdf.SetFooterFunc(func() {
		// Footer
		pdf.SetY(-15)
		pdf.CellFormat(260, 8, "Página "+numberPage+" de "+nPages, "0", 1, "R", false, 0, "")
	})
	for _, class := range classrooms {
		if len(class.GetStudents()) > 0 {

			ts := strconv.Itoa(len(class.GetStudents()))

			pdf.AddPage()
			if len(class.GetStudents()) >= 31 {
				nPages = "2"
				log.Println(nPages)
			} else {
				nPages = "1"
				log.Println(len(class.GetStudents()))
			}
			pdf.AddUTF8Font("Roboto", "", "pdf/Roboto-Regular.ttf")
			pdf.AddUTF8Font("Roboto-Bold", "", "pdf/Roboto-Bold.ttf")

			pdf.SetFont("Roboto-Bold", "", 12)
			// Cell(float w [, float h [, string txt [, mixed border [, int ln [, string align [, boolean fill [, mixed link]]]]]]])
			pdf.Image("pdf/logoGDF.png", 10, 10, 20, 23, false, "", 0, "")
			txt := "CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA"
			pdf.CellFormat(260, 8, txt, "0", 1, "C", false, 0, "")
			txt = "DIÁRIO PROVISÓRIO DE CLASSE - " + class.GetName()
			pdf.CellFormat(260, 8, txt, "0", 1, "C", false, 0, "")

			pdf.CellFormat(20, 8, "", "0", 0, "C", false, 0, "")
			pdf.CellFormat(60, 8, "JUNHO DE 2023", "0", 0, "C", false, 0, "")
			pdf.CellFormat(35, 8, "DISCIPLINA:", "0", 0, "R", false, 0, "")
			txt = "____________________________"
			pdf.CellFormat(55, 8, txt, "B", 0, "L", false, 0, "")
			pdf.CellFormat(35, 8, "PROFESSOR(A):", "0", 0, "R", false, 0, "")
			pdf.CellFormat(55, 8, txt, "B", 1, "L", false, 0, "")

			pdf.SetFont("Roboto-Bold", "", 8)
			pdf.SetTextColor(255, 255, 255)
			pdf.SetDrawColor(200, 200, 200)
			pdf.SetFillColor(100, 100, 100)
			pdf.CellFormat(20, 6, "N.º", "1", 0, "C", true, 0, "")
			days := [23]string{
				"01", "02", "03",
				"05", "06", "07", "08", "09",
				"12", "13", "14", "15", "16",
				"19", "20", "21", "22", "23",
				"26", "27", "28", "29", "30",
			}

			pdf.CellFormat(86, 6, "    Nome do Aluno", "1", 0, "L", true, 0, "")
			for _, d := range days {
				pdf.CellFormat(6, 6, d, "1", 0, "C", true, 0, "")
			}
			pdf.CellFormat(36, 6, "Anotações", "1", 1, "C", true, 0, "")
			pdf.SetTextColor(0, 0, 0)

			pdf.SetFillColor(225, 225, 225)
			numberPage = "1"

			for i, student := range class.GetStudents() {
				if i == 30 {
					pdf.AddPage()

					pdf.SetFont("Roboto-Bold", "", 12)
					// Cell(float w [, float h [, string txt [, mixed border [, int ln [, string align [, boolean fill [, mixed link]]]]]]])
					pdf.Image("pdf/logoGDF.png", 10, 10, 20, 23, false, "", 0, "")
					txt := "CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA"
					pdf.CellFormat(260, 8, txt, "0", 1, "C", false, 0, "")
					txt = "DIÁRIO PROVISÓRIO DE CLASSE - " + class.GetName()
					pdf.CellFormat(260, 8, txt, "0", 1, "C", false, 0, "")

					pdf.CellFormat(20, 8, "", "0", 0, "C", false, 0, "")
					pdf.CellFormat(60, 8, "JUNHO DE 2023", "0", 0, "C", false, 0, "")
					pdf.CellFormat(35, 8, "DISCIPLINA:", "0", 0, "R", false, 0, "")
					txt = "____________________________"
					pdf.CellFormat(55, 8, txt, "B", 0, "L", false, 0, "")
					pdf.CellFormat(35, 8, "PROFESSOR(A):", "0", 0, "R", false, 0, "")
					pdf.CellFormat(55, 8, txt, "B", 1, "L", false, 0, "")

					pdf.SetFont("Roboto-Bold", "", 8)
					pdf.SetTextColor(255, 255, 255)
					pdf.SetDrawColor(200, 200, 200)
					pdf.SetFillColor(100, 100, 100)
					pdf.CellFormat(20, 6, "N.º", "1", 0, "C", true, 0, "")

					pdf.CellFormat(86, 6, "    Nome do Aluno", "1", 0, "L", true, 0, "")
					for _, d := range days {
						pdf.CellFormat(6, 6, d, "1", 0, "C", true, 0, "")
					}
					pdf.CellFormat(36, 6, "Anotações", "1", 1, "C", true, 0, "")
					pdf.SetTextColor(0, 0, 0)

					pdf.SetFillColor(225, 225, 225)
					numberPage = "2"
				}
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
				pdf.CellFormat(86, 4.6, "  "+student.GetName(), "1", 0, "L", fill, 0, "")
				for _, d := range days {
					if d == "asflkj" {
						d = ""
					}
					pdf.CellFormat(6, 4.6, "", "1", 0, "C", fill, 0, "")
				}
				pdf.CellFormat(36, 4.6, "", "1", 1, "C", fill, 0, "")
			}
			ts = strconv.Itoa(len(class.GetStudents()))
			pdf.CellFormat(130, 8, "Total de "+ts+" alunos", "0", 0, "L", false, 0, "")

		}
	}
	err := pdf.OutputFileAndClose("pdf/diary_all_classrooms.pdf")
	return err
}

func FolderCoverClass(class model.ClassroomInterface) error {
	pdf := gofpdf.New("L", "mm", "A4", "")

	pdf.AddPage()
	pdf.AddUTF8Font("Roboto", "", "pdf/Roboto-Regular.ttf")
	pdf.AddUTF8Font("Roboto-Bold", "", "pdf/Roboto-Bold.ttf")

	for i, std := range class.GetStudents() {
		if i != 0 && i%2 == 0 {
			pdf.AddPage()
		}
		pdf.SetFont("Roboto-Bold", "", 12)
		pdf.CellFormat(260, 4, "", "LTR", 1, "C", false, 0, "")

		pdf.CellFormat(25, 6, "Código", "L", 0, "C", false, 0, "")
		pdf.SetFont("Roboto", "", 12)
		txt := strconv.FormatInt(std.GetEducar(), 10)
		pdf.CellFormat(160, 6, txt, "0", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Bold", "", 12)
		txt = "Data de Nascimento"
		pdf.CellFormat(45, 6, txt, "0", 0, "C", false, 0, "")
		pdf.SetFont("Roboto", "", 12)
		txt = std.GetBirthDay().Format("02/01/2006")
		pdf.CellFormat(30, 6, txt, "R", 1, "C", false, 0, "")

		pdf.SetFont("Roboto-Bold", "", 12)
		pdf.CellFormat(25, 8, "Nome", "L", 0, "C", false, 0, "")
		pdf.SetFont("Roboto", "", 16)
		txt = std.GetName()
		pdf.CellFormat(235, 8, txt, "R", 1, "L", false, 0, "")

		pdf.CellFormat(260, 4, "", "LR", 1, "L", false, 0, "")

		pdf.SetFont("Roboto-Bold", "", 12)
		pdf.CellFormat(25, 6, "Ano", "LB", 0, "C", false, 0, "")
		pdf.CellFormat(75, 6, "Turma", "B", 0, "C", false, 0, "")
		pdf.CellFormat(70, 6, "Situação", "B", 0, "C", false, 0, "")
		pdf.CellFormat(90, 6, "Observação", "BR", 1, "L", false, 0, "")

		pdf.SetFont("Roboto", "", 16)
		pdf.CellFormat(25, 8, "2023", "1", 0, "C", false, 0, "")
		txt = class.GetName()[0:9]
		pdf.CellFormat(75, 8, txt, "1", 0, "L", false, 0, "")
		pdf.CellFormat(70, 8, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(90, 8, "", "LR", 1, "L", false, 0, "")

		pdf.CellFormat(25, 8, "2024", "1", 0, "C", false, 0, "")
		pdf.CellFormat(75, 8, "", "1", 0, "L", false, 0, "")
		pdf.CellFormat(70, 8, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(90, 8, "", "LR", 1, "L", false, 0, "")

		pdf.CellFormat(25, 8, "2025", "1", 0, "C", false, 0, "")
		pdf.CellFormat(75, 8, "", "1", 0, "L", false, 0, "")
		pdf.CellFormat(70, 8, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(90, 8, "", "LR", 1, "L", false, 0, "")

		pdf.CellFormat(170, 4, "", "L", 0, "L", false, 0, "")
		pdf.CellFormat(90, 4, "", "LR", 1, "L", false, 0, "")

		pdf.SetFont("Roboto-Bold", "", 12)
		pdf.CellFormat(25, 8, "ANEE", "L", 0, "C", false, 0, "")
		pdf.CellFormat(145, 8, "", "0", 0, "L", false, 0, "")
		pdf.CellFormat(90, 8, "", "LR", 1, "L", false, 0, "")

		pdf.SetFont("Roboto", "", 12)
		txtAnne := ""
		if std.GetANNE() == "" {
			txt = "NÃO"
		} else {
			txt = "SIM"
			txtAnne = std.GetANNE()
		}
		pdf.CellFormat(25, 8, txt, "1", 0, "C", false, 0, "")
		pdf.CellFormat(145, 8, txtAnne, "1", 0, "L", false, 0, "")
		pdf.CellFormat(90, 8, "", "LBR", 1, "L", false, 0, "")
		pdf.Ln(20)

	}

	err := pdf.OutputFileAndClose("pdf/folderCover.pdf")
	return err
}
