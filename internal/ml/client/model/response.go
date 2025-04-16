package model

type ScientificTestMLResponse struct {
	Professions []string `json:"professions"`
}

type PersonalityTestMLResponse struct {
	PersonalityType string   `json:"personality_type"`
	Description     string   `json:"description"`
	Professions     []string `json:"professions"`
}
