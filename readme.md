# AI-Assisted Clinical Report Structuring System

## Overview
This project is a Go-based backend application designed to manage medical patient records and unstructured clinical reports.  
It extracts structured insights from clinical text using AI-assisted logic and stores the results in a relational SQL database.

The system focuses on **data integrity, reliability, and safe medical data handling**.

---

## Features
- Add and manage patient records
- Store unstructured clinical reports
- Extract structured information (diagnosis, key findings, severity)
- Enforce relational integrity using SQL foreign keys
- Console-based interface for simplicity and reliability

---

## Tech Stack
- **Language:** Go
- **Database:** MySQL
- **Concepts:** SQL, Foreign Keys, Backend Development, AI-assisted NLP
- **Tools:** Git, GitHub

---

## Database Design
- `patients` – stores patient details
- `reports` – stores raw clinical report text
- `extracted_data` – stores structured insights derived from reports

All reports are linked to patients using foreign key constraints.

---

## How AI Is Used
AI-assisted text processing is used to convert unstructured clinical reports into structured data fields.
The system is designed to avoid medical diagnosis or predictions and focuses purely on data structuring.

---

## How to Run
1. Set up MySQL and create the database
2. Update database credentials in the Go code
3. Run the application:
```bash
go run .
