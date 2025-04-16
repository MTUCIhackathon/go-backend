package model

type FirstTestMLResponse struct {
	Professions []string `json:"professions"`
}

type SecondTestMLResponse struct {
	PersonalityType string   `json:"personality_type"`
	Description     string   `json:"description"`
	Professions     []string `json:"professions"`
}
