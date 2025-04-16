package model

type FirstTestMLRequest struct {
	Professions map[string]int `json:"professions"`
}

type SecondTestMLRequest struct {
	TestResult string `json:"testResult"`
}
