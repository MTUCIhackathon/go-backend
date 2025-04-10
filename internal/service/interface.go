package service

type Interface interface {
	// Methods for test

	GetTestByName
	GetManyTest

	// Methods for resolved

	CreateResolved

	GetResolvedByID(id)
	GetManyResolved

	// ?
	GetOldResolvedByID

	// Methods for consumers

	CreateConsumer
	GetConsumerByID
	UpdateConsumerPassword
	SendResultOnEmail

	// Methods for
}

// Client

// GetSummary
// GetProfessions
// GetTest
