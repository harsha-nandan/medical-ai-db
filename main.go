package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Medical AI DB ---")
		fmt.Println("1. Add Patient")
		fmt.Println("2. View Patients")
		fmt.Println("3. Add Medical Report")
		fmt.Println("4. View Patient Reports")
		fmt.Println("5. Process Report with AI")
		fmt.Println("6. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name, gender string
			var age int

			fmt.Print("Enter name: ")
			name, _ = reader.ReadString('\n')

			fmt.Print("Enter age: ")
			fmt.Scanln(&age)

			fmt.Print("Enter gender: ")
			fmt.Scanln(&gender)

			err := insertPatient(db, name, age, gender)
			if err != nil {
				fmt.Println("❌ Error inserting patient:", err)
			} else {
				fmt.Println("✅ Patient added successfully")
			}

		case 2:
			patients, err := getAllPatients(db)
			if err != nil {
				fmt.Println("❌ Error fetching patients:", err)
				continue
			}

			fmt.Println("\nPatients:")
			for _, p := range patients {
				fmt.Printf("ID: %d | Name: %s | Age: %d | Gender: %s\n",
					p.ID, p.Name, p.Age, p.Gender)
			}

		case 3:
			var patientID int
			fmt.Print("Enter patient ID: ")
			fmt.Scanln(&patientID)

			fmt.Println("Enter medical report text (single line):")
			reportText, _ := reader.ReadString('\n')

			err := insertReport(db, patientID, reportText)
			if err != nil {
				fmt.Println("❌ Error adding report:", err)
			} else {
				fmt.Println("✅ Medical report added")
			}
		case 4:
			var patientID int
			fmt.Print("Enter patient ID: ")
			fmt.Scanln(&patientID)

			reports, err := getReportsByPatient(db, patientID)
			if err != nil {
				fmt.Println("❌ Error fetching reports:", err)
				continue
			}

			if len(reports) == 0 {
				fmt.Println("No reports found")
				continue
			}

			for _, r := range reports {
				fmt.Println("--------------------")
				fmt.Println("Report ID:", r.ID)
				fmt.Println("Date:", r.CreatedAt)
				fmt.Println("Text:", r.Text)
			}
		case 5:
			var reportID int
			fmt.Print("Enter report ID: ")
			fmt.Scanln(&reportID)

			var reportText string
			err := db.QueryRow(
				"SELECT report_text FROM reports WHERE report_id = ?",
				reportID,
			).Scan(&reportText)

			if err != nil {
				fmt.Println("❌ Report not found")
				continue
			}

			extraction, err := extractMedicalInfo(reportText)
			if err != nil {
				fmt.Println("❌ AI error:", err)
				continue
			}

			err = insertExtraction(db, reportID, extraction)
			if err != nil {
				fmt.Println("❌ DB error:", err)
			} else {
				fmt.Println("✅ AI extraction stored successfully")
				fmt.Println("Diagnosis:", extraction.Diagnosis)
				fmt.Println("Severity:", extraction.Severity)
			}

		case 6:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
