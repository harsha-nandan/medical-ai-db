package main

type Patient struct {
	ID     int
	Name   string
	Age    int
	Gender string
}

type Report struct {
	ID        int
	PatientID int
	Text      string
	CreatedAt string
}

type AIExtraction struct {
	Diagnosis   string `json:"diagnosis"`
	KeyFindings string `json:"key_findings"`
	Severity    string `json:"severity"`
}
