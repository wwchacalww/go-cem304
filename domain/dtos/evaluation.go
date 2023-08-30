package dtos

type SubjectIp struct {
	Name     string `json:"name"`
	Note     string `json:"note"`
	Absences int    `json:"absences"`
}

type CardStudent struct {
	EducaDF  string      `json:"educadf"`
	Subjects []SubjectIp `json:"subjects"`
}

type SubjectCrtl struct {
	Ind     int
	Subject string
}

type InputImportCard struct {
	Term        string        `json:"term"`
	ClassroomID string        `json:"classroom_id"`
	Cards       []CardStudent `json:"cards"`
}
