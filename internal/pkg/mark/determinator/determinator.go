package determinator

import (
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"

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
	log      *zap.Logger
	markList map[string]int
}

func NewMark(log *zap.Logger) mark.Marker {
	if log == nil {
		log = zap.NewNop()
	}
	return &Mark{
		log:      log.Named("mark"),
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

func (m *Mark) MarkDecode(answers [][]string) ([][]int, error) {
	res := make([][]int, len(answers))
	length := len(answers)
	for i := 0; i < length; i++ {
		temp := make([]int, len(answers[i]))
		for j := 0; j < len(answers[i]); j++ {
			num, ok := m.markList[strings.ToLower(answers[i][j])]
			if !ok {
				m.log.Debug("failed to convert result into a list", zap.String("answer", answers[i][j]))
				return nil, ErrWrongAnswer
			}
			temp[j] = num
		}
		res[i] = temp
	}
	return res, nil
}
