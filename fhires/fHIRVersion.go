package fhires

// All published FHIR Versions.
//
// Status: active
//
// http://hl7.org/fhir/ValueSet/FHIR-version
type FHIRVersion string

const (
	// Oldest archived version of FHIR.
	FHIRVersion0_01 FHIRVersion = "0.01"
	// 1st Draft for Comment (Sept 2012 Ballot).
	FHIRVersion0_05 FHIRVersion = "0.05"
	// 2nd Draft for Comment (January 2013 Ballot).
	FHIRVersion0_06 FHIRVersion = "0.06"
	// DSTU 1 Ballot version.
	FHIRVersion0_11 FHIRVersion = "0.11"
	// DSTU 1 version.
	FHIRVersion0_0 FHIRVersion = "0.0"
	// DSTU 1 Official version.
	FHIRVersion0_0_80 FHIRVersion = "0.0.80"
	// DSTU 1 Official version Technical Errata #1.
	FHIRVersion0_0_81 FHIRVersion = "0.0.81"
	// DSTU 1 Official version Technical Errata #2.
	FHIRVersion0_0_82 FHIRVersion = "0.0.82"
	// January 2015 Ballot.
	FHIRVersion0_4 FHIRVersion = "0.4"
	// Draft For Comment (January 2015 Ballot).
	FHIRVersion0_4_0 FHIRVersion = "0.4.0"
	// May 2015 Ballot.
	FHIRVersion0_5 FHIRVersion = "0.5"
	// DSTU 2 Ballot version (May 2015 Ballot).
	FHIRVersion0_5_0 FHIRVersion = "0.5.0"
	// DSTU 2 version.
	FHIRVersion1_0 FHIRVersion = "1.0"
	// DSTU 2 QA Preview + CQIF Ballot (Sep 2015).
	FHIRVersion1_0_0 FHIRVersion = "1.0.0"
	// DSTU 2 (Official version).
	FHIRVersion1_0_1 FHIRVersion = "1.0.1"
	// DSTU 2 (Official version) with 1 technical errata.
	FHIRVersion1_0_2 FHIRVersion = "1.0.2"
	// GAO Ballot version.
	FHIRVersion1_1 FHIRVersion = "1.1"
	// GAO Ballot + draft changes to main FHIR standard.
	FHIRVersion1_1_0 FHIRVersion = "1.1.0"
	// Connectathon 12 (Montreal) version.
	FHIRVersion1_4 FHIRVersion = "1.4"
	// CQF on FHIR Ballot + Connectathon 12 (Montreal).
	FHIRVersion1_4_0 FHIRVersion = "1.4.0"
	// Connectathon 13 (Baltimore) version.
	FHIRVersion1_6 FHIRVersion = "1.6"
	// FHIR STU3 Ballot + Connectathon 13 (Baltimore).
	FHIRVersion1_6_0 FHIRVersion = "1.6.0"
	// Connectathon 14 (San Antonio) version.
	FHIRVersion1_8 FHIRVersion = "1.8"
	// FHIR STU3 Candidate + Connectathon 14 (San Antonio).
	FHIRVersion1_8_0 FHIRVersion = "1.8.0"
	// STU3 version.
	FHIRVersion3_0 FHIRVersion = "3.0"
	// FHIR Release 3 (STU).
	FHIRVersion3_0_0 FHIRVersion = "3.0.0"
	// FHIR Release 3 (STU) with 1 technical errata.
	FHIRVersion3_0_1 FHIRVersion = "3.0.1"
	// FHIR Release 3 (STU) with 2 technical errata.
	FHIRVersion3_0_2 FHIRVersion = "3.0.2"
	// R4 Ballot #1 version.
	FHIRVersion3_3 FHIRVersion = "3.3"
	// R4 Ballot #1 + Connectaton 18 (Cologne).
	FHIRVersion3_3_0 FHIRVersion = "3.3.0"
	// R4 Ballot #2 version.
	FHIRVersion3_5 FHIRVersion = "3.5"
	// R4 Ballot #2 + Connectathon 19 (Baltimore).
	FHIRVersion3_5_0 FHIRVersion = "3.5.0"
	// R4 version.
	FHIRVersion4_0 FHIRVersion = "4.0"
	// FHIR Release 4 (Normative + STU).
	FHIRVersion4_0_0 FHIRVersion = "4.0.0"
	// FHIR Release 4 (Normative + STU) with 1 technical errata.
	FHIRVersion4_0_1 FHIRVersion = "4.0.1"
	// R4B Ballot #1 version.
	FHIRVersion4_1 FHIRVersion = "4.1"
	// R4B Ballot #1 + Connectathon 27 (Virtual).
	FHIRVersion4_1_0 FHIRVersion = "4.1.0"
	// R5 Preview #1 version.
	FHIRVersion4_2 FHIRVersion = "4.2"
	// R5 Preview #1 + Connectathon 23 (Sydney).
	FHIRVersion4_2_0 FHIRVersion = "4.2.0"
	// R4B version.
	FHIRVersion4_3 FHIRVersion = "4.3"
	// FHIR Release 4B (Normative + STU).
	FHIRVersion4_3_0 FHIRVersion = "4.3.0"
	// FHIR Release 4B CI-Builld.
	FHIRVersion4_3_0Cibuild FHIRVersion = "4.3.0-cibuild"
	// FHIR Release 4B Snapshot #1.
	FHIRVersion4_3_0Snapshot1 FHIRVersion = "4.3.0-snapshot1"
	// R5 Preview #2 version.
	FHIRVersion4_4 FHIRVersion = "4.4"
	// R5 Preview #2 + Connectathon 24 (Virtual).
	FHIRVersion4_4_0 FHIRVersion = "4.4.0"
	// R5 Preview #3 version.
	FHIRVersion4_5 FHIRVersion = "4.5"
	// R5 Preview #3 + Connectathon 25 (Virtual).
	FHIRVersion4_5_0 FHIRVersion = "4.5.0"
	// R5 Draft Ballot version.
	FHIRVersion4_6 FHIRVersion = "4.6"
	// R5 Draft Ballot + Connectathon 27 (Virtual).
	FHIRVersion4_6_0 FHIRVersion = "4.6.0"
	// R5 Versions.
	FHIRVersion5_0 FHIRVersion = "5.0"
	// R5 Final Version.
	FHIRVersion5_0_0 FHIRVersion = "5.0.0"
	// R5 Rolling ci-build.
	FHIRVersion5_0_0Cibuild FHIRVersion = "5.0.0-cibuild"
	// R5 Preview #2.
	FHIRVersion5_0_0Snapshot1 FHIRVersion = "5.0.0-snapshot1"
	// R5 Interim tooling stage.
	FHIRVersion5_0_0Snapshot2 FHIRVersion = "5.0.0-snapshot2"
	// R5 Ballot.
	FHIRVersion5_0_0Ballot FHIRVersion = "5.0.0-ballot"
	// R5 January 2023 Staging Release + Connectathon 32.
	FHIRVersion5_0_0Snapshot3 FHIRVersion = "5.0.0-snapshot3"
	// R5 Final QA.
	FHIRVersion5_0_0DraftFinal FHIRVersion = "5.0.0-draft-final"
)

func (code FHIRVersion) String() string {
	return code.Code()
}

func (code FHIRVersion) Code() string {
	return string(code)
}

func (code FHIRVersion) Display() string {
	switch code {
	case FHIRVersion0_01:
		return "0.01"
	case FHIRVersion0_05:
		return "0.05"
	case FHIRVersion0_06:
		return "0.06"
	case FHIRVersion0_11:
		return "0.11"
	case FHIRVersion0_0:
		return "0.0"
	case FHIRVersion0_0_80:
		return "0.0.80"
	case FHIRVersion0_0_81:
		return "0.0.81"
	case FHIRVersion0_0_82:
		return "0.0.82"
	case FHIRVersion0_4:
		return "0.4"
	case FHIRVersion0_4_0:
		return "0.4.0"
	case FHIRVersion0_5:
		return "0.5"
	case FHIRVersion0_5_0:
		return "0.5.0"
	case FHIRVersion1_0:
		return "1.0"
	case FHIRVersion1_0_0:
		return "1.0.0"
	case FHIRVersion1_0_1:
		return "1.0.1"
	case FHIRVersion1_0_2:
		return "1.0.2"
	case FHIRVersion1_1:
		return "1.1"
	case FHIRVersion1_1_0:
		return "1.1.0"
	case FHIRVersion1_4:
		return "1.4"
	case FHIRVersion1_4_0:
		return "1.4.0"
	case FHIRVersion1_6:
		return "1.6"
	case FHIRVersion1_6_0:
		return "1.6.0"
	case FHIRVersion1_8:
		return "1.8"
	case FHIRVersion1_8_0:
		return "1.8.0"
	case FHIRVersion3_0:
		return "3.0"
	case FHIRVersion3_0_0:
		return "3.0.0"
	case FHIRVersion3_0_1:
		return "3.0.1"
	case FHIRVersion3_0_2:
		return "3.0.2"
	case FHIRVersion3_3:
		return "3.3"
	case FHIRVersion3_3_0:
		return "3.3.0"
	case FHIRVersion3_5:
		return "3.5"
	case FHIRVersion3_5_0:
		return "3.5.0"
	case FHIRVersion4_0:
		return "4.0"
	case FHIRVersion4_0_0:
		return "4.0.0"
	case FHIRVersion4_0_1:
		return "4.0.1"
	case FHIRVersion4_1:
		return "4.1"
	case FHIRVersion4_1_0:
		return "4.1.0"
	case FHIRVersion4_2:
		return "4.2"
	case FHIRVersion4_2_0:
		return "4.2.0"
	case FHIRVersion4_3:
		return "4.3"
	case FHIRVersion4_3_0:
		return "4.3.0"
	case FHIRVersion4_3_0Cibuild:
		return "4.3.0-cibuild"
	case FHIRVersion4_3_0Snapshot1:
		return "4.3.0-snapshot1"
	case FHIRVersion4_4:
		return "4.4"
	case FHIRVersion4_4_0:
		return "4.4.0"
	case FHIRVersion4_5:
		return "4.5"
	case FHIRVersion4_5_0:
		return "4.5.0"
	case FHIRVersion4_6:
		return "4.6"
	case FHIRVersion4_6_0:
		return "4.6.0"
	case FHIRVersion5_0:
		return "5.0"
	case FHIRVersion5_0_0:
		return "5.0.0"
	case FHIRVersion5_0_0Cibuild:
		return "5.0.0-cibuild"
	case FHIRVersion5_0_0Snapshot1:
		return "5.0.0-snapshot1"
	case FHIRVersion5_0_0Snapshot2:
		return "5.0.0-snapshot2"
	case FHIRVersion5_0_0Ballot:
		return "5.0.0-ballot"
	case FHIRVersion5_0_0Snapshot3:
		return "5.0.0-snapshot3"
	case FHIRVersion5_0_0DraftFinal:
		return "5.0.0-draft-final"
	}
	return "<unknown>"
}

func (code FHIRVersion) Definition() string {
	switch code {
	case FHIRVersion0_01:
		return "Oldest archived version of FHIR."
	case FHIRVersion0_05:
		return "1st Draft for Comment (Sept 2012 Ballot)."
	case FHIRVersion0_06:
		return "2nd Draft for Comment (January 2013 Ballot)."
	case FHIRVersion0_11:
		return "DSTU 1 Ballot version."
	case FHIRVersion0_0:
		return "DSTU 1 version."
	case FHIRVersion0_0_80:
		return "DSTU 1 Official version."
	case FHIRVersion0_0_81:
		return "DSTU 1 Official version Technical Errata #1."
	case FHIRVersion0_0_82:
		return "DSTU 1 Official version Technical Errata #2."
	case FHIRVersion0_4:
		return "January 2015 Ballot."
	case FHIRVersion0_4_0:
		return "Draft For Comment (January 2015 Ballot)."
	case FHIRVersion0_5:
		return "May 2015 Ballot."
	case FHIRVersion0_5_0:
		return "DSTU 2 Ballot version (May 2015 Ballot)."
	case FHIRVersion1_0:
		return "DSTU 2 version."
	case FHIRVersion1_0_0:
		return "DSTU 2 QA Preview + CQIF Ballot (Sep 2015)."
	case FHIRVersion1_0_1:
		return "DSTU 2 (Official version)."
	case FHIRVersion1_0_2:
		return "DSTU 2 (Official version) with 1 technical errata."
	case FHIRVersion1_1:
		return "GAO Ballot version."
	case FHIRVersion1_1_0:
		return "GAO Ballot + draft changes to main FHIR standard."
	case FHIRVersion1_4:
		return "Connectathon 12 (Montreal) version."
	case FHIRVersion1_4_0:
		return "CQF on FHIR Ballot + Connectathon 12 (Montreal)."
	case FHIRVersion1_6:
		return "Connectathon 13 (Baltimore) version."
	case FHIRVersion1_6_0:
		return "FHIR STU3 Ballot + Connectathon 13 (Baltimore)."
	case FHIRVersion1_8:
		return "Connectathon 14 (San Antonio) version."
	case FHIRVersion1_8_0:
		return "FHIR STU3 Candidate + Connectathon 14 (San Antonio)."
	case FHIRVersion3_0:
		return "STU3 version."
	case FHIRVersion3_0_0:
		return "FHIR Release 3 (STU)."
	case FHIRVersion3_0_1:
		return "FHIR Release 3 (STU) with 1 technical errata."
	case FHIRVersion3_0_2:
		return "FHIR Release 3 (STU) with 2 technical errata."
	case FHIRVersion3_3:
		return "R4 Ballot #1 version."
	case FHIRVersion3_3_0:
		return "R4 Ballot #1 + Connectaton 18 (Cologne)."
	case FHIRVersion3_5:
		return "R4 Ballot #2 version."
	case FHIRVersion3_5_0:
		return "R4 Ballot #2 + Connectathon 19 (Baltimore)."
	case FHIRVersion4_0:
		return "R4 version."
	case FHIRVersion4_0_0:
		return "FHIR Release 4 (Normative + STU)."
	case FHIRVersion4_0_1:
		return "FHIR Release 4 (Normative + STU) with 1 technical errata."
	case FHIRVersion4_1:
		return "R4B Ballot #1 version."
	case FHIRVersion4_1_0:
		return "R4B Ballot #1 + Connectathon 27 (Virtual)."
	case FHIRVersion4_2:
		return "R5 Preview #1 version."
	case FHIRVersion4_2_0:
		return "R5 Preview #1 + Connectathon 23 (Sydney)."
	case FHIRVersion4_3:
		return "R4B version."
	case FHIRVersion4_3_0:
		return "FHIR Release 4B (Normative + STU)."
	case FHIRVersion4_3_0Cibuild:
		return "FHIR Release 4B CI-Builld."
	case FHIRVersion4_3_0Snapshot1:
		return "FHIR Release 4B Snapshot #1."
	case FHIRVersion4_4:
		return "R5 Preview #2 version."
	case FHIRVersion4_4_0:
		return "R5 Preview #2 + Connectathon 24 (Virtual)."
	case FHIRVersion4_5:
		return "R5 Preview #3 version."
	case FHIRVersion4_5_0:
		return "R5 Preview #3 + Connectathon 25 (Virtual)."
	case FHIRVersion4_6:
		return "R5 Draft Ballot version."
	case FHIRVersion4_6_0:
		return "R5 Draft Ballot + Connectathon 27 (Virtual)."
	case FHIRVersion5_0:
		return "R5 Versions."
	case FHIRVersion5_0_0:
		return "R5 Final Version."
	case FHIRVersion5_0_0Cibuild:
		return "R5 Rolling ci-build."
	case FHIRVersion5_0_0Snapshot1:
		return "R5 Preview #2."
	case FHIRVersion5_0_0Snapshot2:
		return "R5 Interim tooling stage."
	case FHIRVersion5_0_0Ballot:
		return "R5 Ballot."
	case FHIRVersion5_0_0Snapshot3:
		return "R5 January 2023 Staging Release + Connectathon 32."
	case FHIRVersion5_0_0DraftFinal:
		return "R5 Final QA."
	}
	return "<unknown>"
}
