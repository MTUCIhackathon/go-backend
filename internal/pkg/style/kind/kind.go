package kind

import (
	"encoding/json"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/style"
)

type Type string

const (
	FirstOrder  Type = "first_order"
	SecondOrder Type = "second_order"
	ThirdOrder  Type = "third_order"
)

func Parse(raw string) (Type, error) {
	kind := Type(raw)
	switch kind {
	case FirstOrder, SecondOrder, ThirdOrder:
		return kind, nil
	}
	return "", style.ErrUnknownType
}

func (t *Type) ownUnmarshalJSON(data []byte) error {
	switch string(data) {
	case "\"" + string(FirstOrder) + "\"":
		*t = FirstOrder
	case "\"" + string(SecondOrder) + "\"":
		*t = SecondOrder
	case "\"" + string(ThirdOrder) + "\"":
		*t = ThirdOrder
	default:
		return style.ErrUnknownType
	}
	return nil
}

func (t *Type) jsonBasedUnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t, err = Parse(s)
	return err
}

func (t *Type) UnmarshalJSON(data []byte) error {
	return t.jsonBasedUnmarshalJSON(data)
}

func (t *Type) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, style.ErrUnknownType
	}
	switch s := *t; s {
	case FirstOrder, SecondOrder, ThirdOrder:
		return []byte("\"" + string(s) + "\""), nil
	}
	return nil, style.ErrUnknownType
}
