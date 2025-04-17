package model

type ScientificTestMLResponse struct {
	Professions []string `json:"professions"`
}

type PersonalityTestMLResponse struct {
	PersonalityType string   `json:"personality_type"`
	Description     string   `json:"description"`
	Professions     []string `json:"professions"`
}

type AITestMLResponse struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

type AITestMLProfessionsResponse struct {
	Professions []string `json:"top_professions"`
}

type AICommonProfessionsResponse struct {
	Professions []string `json:"top_professions"`
}
