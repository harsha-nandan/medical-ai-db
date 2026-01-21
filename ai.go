package main

import (
	"strings"
)

func extractMedicalInfo(reportText string) (*AIExtraction, error) {
	text := strings.ToLower(reportText)

	extraction := &AIExtraction{
		Diagnosis:   "Unknown",
		KeyFindings: "Not specified",
		Severity:    "Low",
	}

	// Simple AI-assisted pattern logic
	if strings.Contains(text, "hemoglobin") || strings.Contains(text, "hb") {
		extraction.KeyFindings = "Low hemoglobin levels detected"
		extraction.Diagnosis = "Possible Anemia"
		extraction.Severity = "Moderate"
	}

	if strings.Contains(text, "fatigue") || strings.Contains(text, "dizziness") {
		extraction.KeyFindings += "; fatigue and dizziness reported"
	}

	if strings.Contains(text, "severe") || strings.Contains(text, "critical") {
		extraction.Severity = "High"
	}

	return extraction, nil
}
