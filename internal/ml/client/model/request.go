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

type AITestMLProfessionsRequest struct {
	AQ map[string]string `json:"user_answers"`
}

type AICommonProfessionsRequest struct {
	FirstTest  []string `json:"test_1"`
	SecondTest []string `json:"test_2"`
	ThirdTest  []string `json:"test_3"`
}

type ImageGenerateRequest struct {
	Profession string `json:"profession"`
}
