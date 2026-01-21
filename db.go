package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	dsn := "root:Harsha@1627@tcp(127.0.0.1:3306)/medical_ai_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to MySQL successfully")
	return db, nil
}
func insertPatient(db *sql.DB, name string, age int, gender string) error {
	query := "INSERT INTO patients (name, age, gender) VALUES (?, ?, ?)"
	_, err := db.Exec(query, name, age, gender)
	return err
}

func getAllPatients(db *sql.DB) ([]Patient, error) {
	rows, err := db.Query("SELECT patient_id, name, age, gender FROM patients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var p Patient
		err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Gender)
		if err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, nil
}
func insertReport(db *sql.DB, patientID int, text string) error {
	query := "INSERT INTO reports (patient_id, report_text) VALUES (?, ?)"
	_, err := db.Exec(query, patientID, text)
	return err
}

func getReportsByPatient(db *sql.DB, patientID int) ([]Report, error) {
	rows, err := db.Query(
		"SELECT report_id, patient_id, report_text, created_at FROM reports WHERE patient_id = ?",
		patientID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []Report
	for rows.Next() {
		var r Report
		err := rows.Scan(&r.ID, &r.PatientID, &r.Text, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		reports = append(reports, r)
	}

	return reports, nil
}
func insertExtraction(db *sql.DB, reportID int, e *AIExtraction) error {
	query := `
	INSERT INTO extracted_data (report_id, diagnosis, key_findings, severity)
	VALUES (?, ?, ?, ?)
	`
	_, err := db.Exec(query, reportID, e.Diagnosis, e.KeyFindings, e.Severity)
	return err
}
