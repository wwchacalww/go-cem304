package reportpdf

import (
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

type Transcript struct {
	level     int
	subject   string
	grade     float64
	absences  int
	workload  int
	formation string
}

func StudentTranscript() error {
	std_tct := []Transcript{}
	lp := Transcript{
		level:     1,
		subject:   "Língua Portuguesa",
		grade:     5.5,
		absences:  12,
		workload:  140,
		formation: "FGB",
	}
	std_tct = append(std_tct, lp)
	lp = Transcript{
		level:     2,
		subject:   "Língua Portuguesa",
		grade:     7.5,
		absences:  30,
		workload:  140,
		formation: "FGB",
	}
	std_tct = append(std_tct, lp)
	edf := Transcript{
		level:     1,
		subject:   "Educação Física",
		grade:     9.5,
		absences:  3,
		workload:  40,
		formation: "FGB",
	}
	std_tct = append(std_tct, edf)
	edf = Transcript{
		level:     2,
		subject:   "Educação Física",
		grade:     7.5,
		absences:  2,
		workload:  40,
		formation: "FGB",
	}
	std_tct = append(std_tct, edf)
	art := Transcript{
		level:     1,
		subject:   "Arte",
		grade:     9.5,
		absences:  3,
		workload:  40,
		formation: "FGB",
	}
	std_tct = append(std_tct, art)
	art = Transcript{
		level:     2,
		subject:   "Arte",
		grade:     7.5,
		absences:  2,
		workload:  40,
		formation: "FGB",
	}
	std_tct = append(std_tct, art)
	lem := Transcript{
		level:     1,
		subject:   "Língua Inglesa",
		grade:     9.5,
		absences:  3,
		workload:  40,
		formation: "FGB",
	}
	std_tct = append(std_tct, lem)
	lem = Transcript{
		level:     2,
		subject:   "Língua Inglesa",
		grade:     7.5,
		absences:  2,
		workload:  40,
		formation: "FGB",
	}
	std_tct = append(std_tct, lem)
	mat := Transcript{
		level:     1,
		subject:   "Matemática",
		grade:     6.5,
		absences:  12,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, mat)
	mat = Transcript{
		level:     2,
		subject:   "Matemática",
		grade:     8.5,
		absences:  30,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, mat)
	bio := Transcript{
		level:     1,
		subject:   "Biologia",
		grade:     6.5,
		absences:  12,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, bio)
	bio = Transcript{
		level:     2,
		subject:   "Biologia",
		grade:     8.5,
		absences:  30,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, bio)
	fis := Transcript{
		level:     1,
		subject:   "Física",
		grade:     6.5,
		absences:  12,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, fis)
	fis = Transcript{
		level:     2,
		subject:   "Física",
		grade:     8.5,
		absences:  30,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, fis)
	qui := Transcript{
		level:     1,
		subject:   "Quimica",
		grade:     6.5,
		absences:  12,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, qui)
	qui = Transcript{
		level:     2,
		subject:   "Quimica",
		grade:     8.5,
		absences:  30,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, qui)
	fil := Transcript{
		level:     1,
		subject:   "Filosofia",
		grade:     6.5,
		absences:  12,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, fil)
	fil = Transcript{
		level:     2,
		subject:   "Filosofia",
		grade:     8.5,
		absences:  30,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, fil)
	geo := Transcript{
		level:     1,
		subject:   "Geografia",
		grade:     6.5,
		absences:  12,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, geo)
	geo = Transcript{
		level:     2,
		subject:   "Geografia",
		grade:     8.5,
		absences:  30,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, geo)
	his := Transcript{
		level:     1,
		subject:   "História",
		grade:     6.5,
		absences:  12,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, his)
	his = Transcript{
		level:     2,
		subject:   "História",
		grade:     8.5,
		absences:  30,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, his)
	soc := Transcript{
		level:     1,
		subject:   "Sociologia",
		grade:     6.5,
		absences:  12,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, soc)
	soc = Transcript{
		level:     2,
		subject:   "Sociologia",
		grade:     8.5,
		absences:  30,
		workload:  104,
		formation: "FGB",
	}
	std_tct = append(std_tct, soc)
	esp := Transcript{
		level:     1,
		subject:   "Língua Espanhola",
		grade:     6.5,
		absences:  12,
		workload:  40,
		formation: "IF",
	}
	std_tct = append(std_tct, esp)
	esp = Transcript{
		level:     2,
		subject:   "Língua Espanhola",
		grade:     8.5,
		absences:  30,
		workload:  31,
		formation: "IF",
	}
	std_tct = append(std_tct, esp)
	pv := Transcript{
		level:     1,
		subject:   "Projeto de Vida",
		grade:     6.5,
		absences:  12,
		workload:  40,
		formation: "IF",
	}
	std_tct = append(std_tct, pv)
	pv = Transcript{
		level:     2,
		subject:   "Projeto de Vida",
		grade:     8.5,
		absences:  30,
		workload:  40,
		formation: "IF",
	}
	std_tct = append(std_tct, pv)
	pvI := Transcript{
		level:     1,
		subject:   "Projeto de Vida",
		grade:     6.5,
		absences:  12,
		workload:  40,
		formation: "IF",
	}
	std_tct = append(std_tct, pvI)
	pvI = Transcript{
		level:     2,
		subject:   "Projeto de Vida",
		grade:     8.5,
		absences:  30,
		workload:  40,
		formation: "IF",
	}
	std_tct = append(std_tct, pvI)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(5, 5, 5)
	pdf.AddPage()

	pdf.AddUTF8Font("Roboto", "", "pdf/Roboto-Regular.ttf")
	pdf.AddUTF8Font("Roboto-Bold", "", "pdf/Roboto-Bold.ttf")
	pdf.SetFillColor(220, 220, 220)

	pdf.SetFont("Roboto-Bold", "", 9)

	// Cabeçalho
	pdf.Image("pdf/logoGDF.png", 5, 4, 25, 29, false, "", 0, "")
	txt := "GDF – SECRETARIA DE ESTADO DE EDUCAÇÃO DO DISTRITO FEDERAL"
	pdf.CellFormat(200, 4, txt, "0", 1, "C", false, 0, "")
	pdf.Ln(2)
	txt = "CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA"
	pdf.CellFormat(200, 4, txt, "0", 1, "C", false, 0, "")
	txt = "CRE - SAMAMBAIA"
	pdf.SetFont("Roboto", "", 9)
	pdf.CellFormat(200, 4, txt, "0", 1, "C", false, 0, "")
	txt = "PORTARIANº 03 DE 12/01/2004-SEDF"
	pdf.CellFormat(200, 4, txt, "0", 1, "C", false, 0, "")
	txt = "QN 304 Conjunto 4 - Samambaia Sul(Samambaia) - DF"
	pdf.CellFormat(200, 4, txt, "0", 1, "C", false, 0, "")
	pdf.Ln(2)
	txt = "HISTÓRICO ESCOLAR - ENSINO MÉDIO"
	pdf.SetFont("Roboto-Bold", "", 11)
	pdf.CellFormat(200, 4, txt, "0", 1, "C", false, 0, "")

	// Fundamentação legal
	txt = ""
	pdf.CellFormat(200, 2, txt, "LTR", 1, "C", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 9)
	txt = "FUNDAMENTAÇÃO LEGAL DO CURSO"
	pdf.CellFormat(200, 2, txt, "LR", 1, "C", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 4, txt, "LR", 1, "C", false, 0, "")
	pdf.SetFont("Roboto", "", 9)
	txt = "Lei n.º 9394/96; Lei nº 13.415 de 2017 Parecer CNE/CEB n.º 3/2018; Resolução CNE/CEB n.º 3/2018; Portaria MEC nº 649/2018 Resolução n.º"
	pdf.CellFormat(200, 4, txt, "LR", 1, "J", false, 0, "")
	txt = "2/2020 – CEDF, com alterações dadas pela Resolução nº 1/2021-CEDF, publicada no DODF nº 30, de 12 de fevereiro de 2021, pela Resolução nº"
	pdf.CellFormat(200, 4, txt, "LR", 1, "J", false, 0, "")
	txt = "2/2021-CEDF, publicada no DODF n.º 126, de 7 de julho de 2021, e pela Resolução n.º 3/2021-CEDF, publicada no DODF nº 158, de 20 de"
	pdf.CellFormat(200, 4, txt, "LR", 1, "J", false, 0, "")
	txt = "agosto de 2021. Portaria nº 1.094/SEEDF, de 16 de novembro de 2022, publicada no DODF n.º 214, de 17 de novembro de 2022 e Parecer n.º"
	pdf.CellFormat(200, 4, txt, "LR", 1, "J", false, 0, "")
	txt = "210/2022-CEDF, de 8 de novembro de 2022, que valida o Plano de Implementação do Novo Ensino Médio da Rede Pública de Ensino do Distrito"
	pdf.CellFormat(200, 4, txt, "LR", 1, "J", false, 0, "")
	txt = "Federal, incluindo o Quadro-Resumo da Matriz Curricular."
	pdf.CellFormat(200, 4, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 2, txt, "LBR", 1, "L", false, 0, "")

	// Dados do aluno ou aluna
	pdf.SetFont("Roboto-Bold", "", 9)
	txt = "Estudante: LAURA SIQUEIRA SILVA"
	pdf.CellFormat(140, 5, txt, "1", 0, "L", false, 0, "")
	txt = "RA: 000120049515-9/DF"
	pdf.CellFormat(60, 5, txt, "1", 1, "L", false, 0, "")
	pdf.SetFont("Roboto", "", 9)
	txt = "Data Nasc.: 27/09/2005"
	pdf.CellFormat(45, 5, txt, "1", 0, "L", false, 0, "")
	txt = "Sexo: F"
	pdf.CellFormat(20, 5, txt, "1", 0, "L", false, 0, "")
	txt = "Nacionalidade: Brasileira"
	pdf.CellFormat(50, 5, txt, "1", 0, "L", false, 0, "")
	txt = "Naturalidade: SÃO MATEUS"
	pdf.CellFormat(85, 5, txt, "1", 1, "L", false, 0, "")
	txt = "Carteira deIdentidade: 3.556.946 SSP-DF"
	pdf.CellFormat(100, 5, txt, "1", 0, "L", false, 0, "")
	txt = "CPF: 06599751148"
	pdf.CellFormat(100, 5, txt, "1", 1, "L", false, 0, "")
	txt = "Filiação: CELESTE SHUINA SIQUEIRA; FRANCISCO CARLOS LINDEMBERGUE SILVA"
	pdf.CellFormat(200, 5, txt, "1", 1, "L", false, 0, "")

	// Componentes Curriculares
	pdf.SetFont("Roboto-Bold", "", 9)
	txt = "COMPONENTES/UNIDADES"
	pdf.CellFormat(65, 5, txt, "LTR", 0, "C", false, 0, "")
	txt = "1ª SÉRIE"
	pdf.CellFormat(45, 5, txt, "1", 0, "C", true, 0, "")
	txt = "2ª SÉRIE"
	pdf.CellFormat(45, 5, txt, "1", 0, "C", true, 0, "")
	txt = "3ª SÉRIE"
	pdf.CellFormat(45, 5, txt, "1", 1, "C", true, 0, "")
	txt = "CURRICULARES"
	pdf.CellFormat(65, 5, txt, "LBR", 0, "C", false, 0, "")
	for i := range [3]int{} {
		if i == 2 {
			pdf.CellFormat(18, 5, "M/C/N", "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 5, "CH", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 5, "FALTAS", "1", 1, "C", false, 0, "")
		} else {
			pdf.CellFormat(18, 5, "M/C/N", "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 5, "CH", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 5, "FALTAS", "1", 0, "C", false, 0, "")
		}
	}
	txt = "Formação Geral Básica - FGB"
	pdf.CellFormat(65, 5, txt, "1", 0, "C", true, 0, "")
	txt = ""
	pdf.CellFormat(135, 5, txt, "1", 1, "C", true, 0, "")

	// Notas do Histórico
	pdf.SetFont("Roboto", "", 8)
	ctrl := ""
	lvl := 0
	abs := 0
	for _, sbj := range std_tct {
		if lvl < sbj.level {
			lvl = sbj.level
		}
		if ctrl == "" && sbj.formation == "FGB" {
			ctrl = sbj.subject
			lvl = sbj.level
			abs = sbj.absences
			pdf.CellFormat(65, 4, ctrl, "1", 0, "L", false, 0, "")
			pdf.CellFormat(18, 4, strconv.FormatFloat(sbj.grade, 'f', 2, 64), "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 4, strconv.Itoa(sbj.workload), "1", 0, "C", false, 0, "")
		} else {
			if sbj.level == 1 && sbj.formation == "FGB" {
				if lvl == 1 {
					pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
				} else if lvl == 2 {
					pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
				} else {
					pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 1, "C", false, 0, "")
				}
				abs = sbj.absences
				ctrl = sbj.subject
				pdf.CellFormat(65, 4, ctrl, "1", 0, "L", false, 0, "")
				pdf.CellFormat(18, 4, strconv.FormatFloat(sbj.grade, 'f', 2, 64), "1", 0, "C", false, 0, "")
				pdf.CellFormat(9, 4, strconv.Itoa(sbj.workload), "1", 0, "C", false, 0, "")
			}
			if sbj.level > 1 && sbj.formation == "FGB" {
				pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
				abs = sbj.absences
				ctrl = sbj.subject
				pdf.CellFormat(18, 4, strconv.FormatFloat(sbj.grade, 'f', 2, 64), "1", 0, "C", false, 0, "")
				pdf.CellFormat(9, 4, strconv.Itoa(sbj.workload), "1", 0, "C", false, 0, "")
			}
		}
	}
	if lvl == 1 {
		pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
	} else if lvl == 2 {
		pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
	} else {
		pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 1, "C", false, 0, "")
	}

	// Itinerário Formativo
	pdf.SetFont("Roboto-Bold", "", 9)
	txt = "Itinerários Formativos - IF"
	pdf.CellFormat(65, 5, txt, "1", 0, "C", true, 0, "")
	txt = ""
	pdf.CellFormat(135, 5, txt, "1", 1, "C", true, 0, "")
	pdf.SetFont("Roboto", "", 8)
	for _, sbj := range std_tct {
		if sbj.subject == "Língua Espanhola" && sbj.level == 1 {
			ctrl = sbj.subject
			abs = sbj.absences
			pdf.CellFormat(65, 4, ctrl, "1", 0, "L", false, 0, "")
			pdf.CellFormat(18, 4, strconv.FormatFloat(sbj.grade, 'f', 2, 64), "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 4, strconv.Itoa(sbj.workload), "1", 0, "C", false, 0, "")
		} else {
			if sbj.level == 1 && sbj.formation == "IF" {
				if lvl == 1 {
					pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
				} else if lvl == 2 {
					pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
					pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
				} else {
					pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 1, "C", false, 0, "")
				}
				abs = sbj.absences
				ctrl = sbj.subject
				pdf.CellFormat(65, 4, ctrl, "1", 0, "L", false, 0, "")
				pdf.CellFormat(18, 4, strconv.FormatFloat(sbj.grade, 'f', 2, 64), "1", 0, "C", false, 0, "")
				pdf.CellFormat(9, 4, strconv.Itoa(sbj.workload), "1", 0, "C", false, 0, "")
			}
			if sbj.level > 1 && sbj.formation == "IF" {
				pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
				abs = sbj.absences
				ctrl = sbj.subject
				pdf.CellFormat(18, 4, strconv.FormatFloat(sbj.grade, 'f', 2, 64), "1", 0, "C", false, 0, "")
				pdf.CellFormat(9, 4, strconv.Itoa(sbj.workload), "1", 0, "C", false, 0, "")
			}
		}
	}
	if lvl == 1 {
		pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
	} else if lvl == 2 {
		pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
		pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
	} else {
		pdf.CellFormat(18, 4, strconv.Itoa(abs), "1", 1, "C", false, 0, "")
	}
	for i := range [10]int{} {
		pdf.CellFormat(65, 4, "Itinerário Formativo", "1", 0*i, "L", false, 0, "")
		pdf.CellFormat(18, 4, "EP", "1", 0, "C", false, 0, "")
		pdf.CellFormat(9, 4, "40", "1", 0, "C", false, 0, "")
		if lvl == 1 {
			pdf.CellFormat(18, 4, "0", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
		} else if lvl == 2 {
			pdf.CellFormat(18, 4, "0", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "EP", "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 4, "40", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "0", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "---", "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 4, "---", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "---", "1", 1, "C", false, 0, "")
		} else {
			pdf.CellFormat(18, 4, "0", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "EP", "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 4, "40", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "0", "1", 0, "C", false, 0, "")
			pdf.CellFormat(9, 4, "EP", "1", 0, "C", false, 0, "")
			pdf.CellFormat(18, 4, "0", "1", 1, "C", false, 0, "")
		}
	}

	// Media
	pdf.SetFont("Roboto-Bold", "", 9)
	txt = "Média por Área do Conhecimento"
	pdf.CellFormat(65, 5, txt, "1", 0, "C", true, 0, "")
	txt = "1ª SÉRIE"
	pdf.CellFormat(45, 5, txt, "1", 0, "C", true, 0, "")
	txt = "2ª SÉRIE"
	pdf.CellFormat(45, 5, txt, "1", 0, "C", true, 0, "")
	txt = "3ª SÉRIE"
	pdf.CellFormat(45, 5, txt, "1", 1, "C", true, 0, "")

	pdf.SetFont("Roboto", "", 9)
	for i := range [5]int{} {
		if i == 0 {
			txt = "Língua Portuguesa"
		}
		if i == 1 {
			txt = "Linguagens e suas Tecnologias"
		}
		if i == 2 {
			txt = "Matemática e suas Tecnologias"
		}
		if i == 3 {
			txt = "Ciências da Natureza e suas Tecnologias"
		}
		if i == 4 {
			txt = "Ciências Humanas e Sociais Aplicadas "
		}

		pdf.CellFormat(65, 5, txt, "1", 0, "C", false, 0, "")
		txt = "9.0"
		pdf.CellFormat(45, 5, txt, "1", 0, "C", false, 0, "")
		txt = "6,7"
		pdf.CellFormat(45, 5, txt, "1", 0, "C", false, 0, "")
		txt = "---"
		pdf.CellFormat(45, 5, txt, "1", 1, "C", false, 0, "")
	}
	pdf.SetFont("Roboto-Bold", "", 9)
	txt = "Média Final/Global"
	pdf.CellFormat(65, 5, txt, "1", 0, "C", true, 0, "")
	txt = "9.0"
	pdf.CellFormat(45, 5, txt, "1", 0, "C", true, 0, "")
	txt = "6,7"
	pdf.CellFormat(45, 5, txt, "1", 0, "C", true, 0, "")
	txt = "---"
	pdf.CellFormat(45, 5, txt, "1", 1, "C", true, 0, "")

	// Dias Letivos
	pdf.SetFont("Roboto", "", 8)
	txt = "Dias Letivos"
	pdf.CellFormat(65, 4, txt, "1", 0, "C", false, 0, "")
	txt = "200"
	pdf.CellFormat(45, 4, txt, "1", 0, "C", false, 0, "")
	txt = "200"
	pdf.CellFormat(45, 4, txt, "1", 0, "C", false, 0, "")
	txt = "---"
	pdf.CellFormat(45, 4, txt, "1", 1, "C", false, 0, "")
	txt = "Total de Faltas"
	pdf.CellFormat(65, 4, txt, "1", 0, "C", false, 0, "")
	txt = "54:00"
	pdf.CellFormat(45, 4, txt, "1", 0, "C", false, 0, "")
	txt = "32:30"
	pdf.CellFormat(45, 4, txt, "1", 0, "C", false, 0, "")
	txt = "---"
	pdf.CellFormat(45, 4, txt, "1", 1, "C", false, 0, "")
	txt = "Carga Horária Anual"
	pdf.CellFormat(65, 4, txt, "1", 0, "C", false, 0, "")
	txt = "1.120"
	pdf.CellFormat(45, 4, txt, "1", 0, "C", false, 0, "")
	txt = "1.110"
	pdf.CellFormat(45, 4, txt, "1", 0, "C", false, 0, "")
	txt = "---"
	pdf.CellFormat(45, 4, txt, "1", 1, "C", false, 0, "")

	// Serie Ano Escola Cidade Resultado
	pdf.SetFont("Roboto-Bold", "", 8)
	txt = "Série"
	pdf.CellFormat(15, 5, txt, "1", 0, "C", false, 0, "")
	txt = "Ano"
	pdf.CellFormat(15, 5, txt, "1", 0, "C", false, 0, "")
	txt = "INSTITUIÇÃO EDUCACIONAL/UNIDADE ESCOLAR"
	pdf.CellFormat(100, 5, txt, "1", 0, "C", false, 0, "")
	txt = "CIDADE-UF"
	pdf.CellFormat(40, 5, txt, "1", 0, "C", false, 0, "")
	txt = "RESUL. FINAL"
	pdf.CellFormat(30, 5, txt, "1", 1, "C", false, 0, "")

	pdf.SetFont("Roboto", "", 8)
	txt = "1ª"
	pdf.CellFormat(15, 5, txt, "1", 0, "C", false, 0, "")
	txt = "2021"
	pdf.CellFormat(15, 5, txt, "1", 0, "C", false, 0, "")
	txt = "Centro de Ensino Médio 304 de Samambaia"
	pdf.CellFormat(100, 5, txt, "1", 0, "C", false, 0, "")
	txt = "Samambaia-DF"
	pdf.CellFormat(40, 5, txt, "1", 0, "C", false, 0, "")
	txt = "APROVADO"
	pdf.CellFormat(30, 5, txt, "1", 1, "C", false, 0, "")

	txt = "2ª"
	pdf.CellFormat(15, 5, txt, "1", 0, "C", false, 0, "")
	txt = "2021"
	pdf.CellFormat(15, 5, txt, "1", 0, "C", false, 0, "")
	txt = "Centro de Ensino Médio 304 de Samambaia"
	pdf.CellFormat(100, 5, txt, "1", 0, "C", false, 0, "")
	txt = "Samambaia-DF"
	pdf.CellFormat(40, 5, txt, "1", 0, "C", false, 0, "")
	txt = "APROVADO"
	pdf.CellFormat(30, 5, txt, "1", 1, "C", false, 0, "")

	txt = "3ª"
	pdf.CellFormat(15, 5, txt, "1", 0, "C", false, 0, "")
	txt = "2021"
	pdf.CellFormat(15, 5, txt, "1", 0, "C", false, 0, "")
	txt = "Centro de Ensino Médio 304 de Samambaia"
	pdf.CellFormat(100, 5, txt, "1", 0, "C", false, 0, "")
	txt = "Samambaia-DF"
	pdf.CellFormat(40, 5, txt, "1", 0, "C", false, 0, "")
	txt = "APROVADO"
	pdf.CellFormat(30, 5, txt, "1", 1, "C", false, 0, "")

	pdf.AddPage()

	// Observações Gerais
	pdf.SetFont("Roboto-Bold", "", 8)
	txt = ""
	pdf.CellFormat(200, 5, txt, "LTR", 1, "C", false, 0, "")
	txt = "Observações Gerais:"
	pdf.CellFormat(200, 5, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 2, txt, "LR", 1, "L", false, 0, "")
	pdf.SetFont("Roboto", "", 8)
	txt = "1. O Curso Novo Ensino Médio está organizado em duas fases: a Fase 1 compreende a 1ª e a 2ª séries, composta por 4 (quatro) semestres letivos, e a"
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = "Fase 2 compreende a 3ª serie,composta por 2 (dois) semestres letivos."
	pdf.CellFormat(200, 5, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 2, txt, "LR", 1, "L", false, 0, "")
	txt = "2. Todos os componentes curriculares das Áreas de Conhecimento da FGB são obrigatórios para os estudantes."
	pdf.CellFormat(200, 5, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LR", 1, "L", false, 0, "")
	txt = "3. A frequência inferiora 75%, ao final de cada série, resulta na retenção do estudante de acordo como inciso VI do Art. 24 da LDB."
	pdf.CellFormat(200, 2, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LR", 1, "L", false, 0, "")
	txt = "4. As unidades curriculares dos itinerários formativos estão arranjadas de quatro formas: I – Língua Espanhola: unidadecurricular obrigatória; II –"
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = "Projeto de Vida: unidade curricular obrigatória para orientação do percurso formativo do estudante; III- (1) Eletivas: unidades curriculares de escolha"
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = "do estudante para ampliação das aprendizagens; III (2) Projetos Interventivos: unidades curriculares para atendimento das necessidades pedagógicas"
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = "dos estudantes; IV– Trilhas de Aprendizagem: unidades curriculares planejadas de forma a caracterizar uma área de aprofundamento do estudante."
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LR", 1, "L", false, 0, "")
	txt = "5. Nos IF, Língua Espanhola e Projeto de Vida são obrigatórios para todos os estudantes."
	pdf.CellFormat(200, 5, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LR", 1, "L", false, 0, "")
	txt = "6. As unidades curriculares eletivas  eas Trilhas de Aprendizagem propostas devem ser baseadas nos eixos estruturantes da Portaria no 1.432/2018 e"
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = "no Currículo em Movimento do Novo Ensino Médio (Parecer n.º 112 /2020 – CEDF - Portaria n.º 507, de 30/12/2020, DODF n.º 1, de"
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = "04/01/2021)."
	pdf.CellFormat(200, 5, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LR", 1, "L", false, 0, "")
	txt = "7. Em caso de “Adequação curricular na temporalidade”, para estudantes com a flexibilização do tempo para realizar as atividades e o desenvolvimento"
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = "de conteúdo, a situação final na série será: CURS.TEMPORALIDADE."
	pdf.CellFormat(200, 5, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LR", 1, "L", false, 0, "")
	txt = "8. Dados transcritos na íntegra do documento de origem referente aos anos - ,considerando o ingresso do estudante nesta Rede Pública de Ensino."
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LR", 1, "L", false, 0, "")
	txt = "9. Os campos do histórico escolar marcadoscom *** (três asteriscos), representam a aprovação obtida no(s) componente(s) curricular(es), cursado(s)"
	pdf.CellFormat(200, 5, txt, "LR", 1, "J", false, 0, "")
	txt = "em regime de dependência – progressão parcial."
	pdf.CellFormat(200, 5, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LR", 1, "L", false, 0, "")
	txt = ""
	pdf.CellFormat(200, 3, txt, "LBR", 1, "L", false, 0, "")
	pdf.Ln(3)

	// Legendas
	pdf.SetFont("Roboto-Bold", "", 10)
	txt = "LEGENDAS:"
	pdf.CellFormat(60, 10, txt, "1", 0, "C", false, 0, "")
	txt = "Menções/Conceitos/Notas que aprovam"
	pdf.CellFormat(80, 10, txt, "1", 0, "C", false, 0, "")
	txt = "INFORMAÇÕES"
	pdf.CellFormat(60, 5, txt, "LTR", 1, "C", false, 0, "")
	txt = ""
	pdf.CellFormat(140, 5, txt, "0", 0, "C", false, 0, "")
	txt = ""
	pdf.CellFormat(60, 5, txt, "LBR", 1, "C", false, 0, "")

	pdf.CellFormat(60, 85, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 85, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(60, 85, "", "1", 1, "C", false, 0, "")

	pdf.SetFont("Roboto", "", 7)
	txt = `CH= carga horáriacumprida pelo estudante
	AE= aproveitamento deestudos
	ABA= abandono
	RF = reprovado por faltas
	CUR= cursando
	TRANSF = transferido
	M/C/N= Menções/Conceitos/Notas
	EP = Envolvimento Pleno
	ER= Envolvimento Regular
	ES= Envolvimento Satisfatório`
	pdf.SetY(155)
	pdf.MultiCell(60, 5, txt, "0", "L", false)
	pdf.SetY(143)
	pdf.SetX(65)
	txt = `1. Os registros dos resultdos da avaliação na FGB ocorrem em escala numérica de notas de 0 (zero) a 10 (dez) por componentecurricular, com	carga horáriaanualeregistros de notas bimestraiseao final do ano letivo.
2. A média simples, por componente curricular na FGB é de 5,0 (cinco)	pontos,emescala numérica de 0 (zero)a 10 (dez).
3. As menções dos Itinerários Formativos (IF), dos Itinerários Integradores (II), das Trilhas de Aprendizagem, dos Projetos Interventivos e do Projeto de Vida são computadas em média modal. A	médiafinaléamenção (EP, ESou ER)commaior frequência nos registros.
4. Médias por áreas do conhecimento: é a média aritmética simples dos componentescurriculares decada uma dasáreas.
5. Média FINAL/GLOBAL na FGB éfeitaa partir das médias das notas das Áreas do Conhecimento, ao final de cada série. Sendo a única média	considerada paraaprovação ou reprovação.
6. Na 1ª sérieindependente das médias, o estudanteterá PROGRESSÃO	CONTINUADA.
7. Média FINAL/GLOBAL para aprovação na 2.ª e 3.ª séries: igual ou	maior que 5,0`
	pdf.MultiCell(80, 4, txt, "0", "J", false)
	txt = `Nos turnos matutino e vespertino,a hora-aulaequivale a 50 minutos. Documento escolar válido, somente, mediante assinaturas e carimbos do(a) Diretor(a) e do(a) Chefe de Secretaria Escolar e a inexistência de rasuras.`
	pdf.SetY(171)
	pdf.SetX(145)
	pdf.MultiCell(60, 4, txt, "0", "J", false)

	// Carimbo
	pdf.SetY(230)
	pdf.Image("pdf/logoGDF.png", 28, 231, 19, 22, false, "", 0, "")
	pdf.CellFormat(20, 5, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(160, 5, "", "LTR", 1, "C", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 16)
	txt = "GOVERNO DO DISTRITO FEDERAL"
	pdf.CellFormat(20, 5, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(160, 5, txt, "LR", 1, "C", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 10)
	txt = "SECRETARIA DE ESTADO DE EDUCAÇÃO"
	pdf.CellFormat(20, 4, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(160, 4, txt, "LR", 1, "C", false, 0, "")
	pdf.SetFont("Roboto", "", 8)
	txt = "CENTRO DE ENSINO MÉDIO 304 DE SAMAMBAIA"
	pdf.CellFormat(20, 4, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(160, 4, txt, "LR", 1, "C", false, 0, "")
	pdf.SetFont("Roboto", "", 7)
	txt = "Conferido o presente documento, declaramos sua autenticidade e regularidade, de acordo com os"
	pdf.CellFormat(20, 2, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(160, 2, "", "LR", 1, "C", false, 0, "")
	pdf.CellFormat(20, 4, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(20, 4, "", "L", 0, "J", false, 0, "")
	pdf.CellFormat(140, 4, txt, "R", 1, "C", false, 0, "")
	txt = "registros escolares e com a legislação vigente"
	pdf.CellFormat(20, 4, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(160, 4, txt, "LR", 1, "L", false, 0, "")
	txt = "Brasília, 7 de março de 2024."
	pdf.CellFormat(20, 7, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(160, 7, txt, "LR", 1, "R", false, 0, "")
	pdf.CellFormat(20, 4, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(70, 4, "Diretor", "L", 0, "C", false, 0, "")
	pdf.CellFormat(70, 4, "Secretário Escolar", "0", 0, "C", false, 0, "")
	pdf.CellFormat(20, 4, "", "R", 1, "J", false, 0, "")

	txt = "(assinatura e carimbo)"
	pdf.CellFormat(20, 4, "", "0", 0, "J", false, 0, "")
	pdf.CellFormat(70, 4, txt, "LB", 0, "C", false, 0, "")
	pdf.CellFormat(70, 4, txt, "B", 0, "C", false, 0, "")
	pdf.CellFormat(20, 4, "", "RB", 1, "J", false, 0, "")

	err := pdf.OutputFileAndClose("pdf/student-transcript.pdf")
	return err
}
