package model

type ScientificTestMLRequest struct {
	Professions map[string]int `json:"test_result"`
}

type PersonalityTestMLRequest struct {
	TestResult string `json:"test_result"`
}

type AITestMLRequest struct {
	AQ map[string]string `json:"questions"`
}
