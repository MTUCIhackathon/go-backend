package model

type ScientificTestMLRequest struct {
	Professions map[string]int `json:"professions"`
}

type PersonalityTestMLRequest struct {
	TestResult string `json:"test_result"`
}
