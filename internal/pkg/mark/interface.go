package mark

type Marker interface {
	MarkResult(answer string) (int8, error)
}
