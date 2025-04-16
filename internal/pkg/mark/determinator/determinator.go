package determinator

import (
	"errors"
	"strings"

	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/mark"
)

var (
	markList = map[string]int8{
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
	markList map[string]int8
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

func (m *Mark) MarkResult(answer string) (int8, error) {
	result, ok := m.markList[strings.ToLower(answer)]
	if !ok {
		m.log.Debug(
			"failed to mark result: negative process result",
			zap.String("answer", answer),
			zap.Int8("result", result),
		)
		return 0, ErrWrongAnswer
	}

	m.log.Debug(
		"marked result",
		zap.String("answer", answer),
		zap.Int8("result", result),
	)

	return result, nil
}

func (m *Mark) MarkDecode(answers [][]string) ([][]int8, error) {
	res := make([][]int8, len(answers))
	length := len(answers)
	for i := 0; i < length; i++ {
		temp := make([]int8, len(answers[i]))
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
