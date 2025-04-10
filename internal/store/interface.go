package store

type (
	Interface interface {
		Forms() FormsRepository
		Consumers() ConsumersRepository
	}

	ConsumersRepository interface {
		// methods
	}
	FormsRepository interface {
		// methods
	}
)
