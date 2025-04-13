package style

type Interface interface {
	UnmarshalJSON(data []byte) error
	MarshalJSON() ([]byte, error)
}
