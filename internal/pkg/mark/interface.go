package mark

type Marker interface {
	MarkResult(answer string) (int, error)
}
