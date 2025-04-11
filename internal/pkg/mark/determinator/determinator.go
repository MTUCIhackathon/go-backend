package determinator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/mark"
)

var (
	markList = map[string]int{
		"да":         2,
		"скорее да":  1,
		"возможно":   0,
		"скорее нет": -1,
		"нет":        -2,
	}
)

var (
	ErrWrongAnswer = errors.New("wrong answer: answer not found")
)

type Mark struct {
	markList map[string]int
}

func NewMark() mark.Marker {
	return &Mark{
		markList: markList,
	}
}

func (m *Mark) MarkResult(answer string) (int, error) {
	result, ok := m.markList[strings.ToLower(answer)]
	fmt.Println(strings.ToLower(answer))
	if !ok {
		return 0, ErrWrongAnswer
	}
	return result, nil
}
