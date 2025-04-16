package determinator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestMark_New(t *testing.T) {
	marker := NewMark(zap.L())
	tt := []struct {
		name   string
		answer string
		result int
		expect error
	}{
		{name: "case1", answer: "Да", result: 2, expect: nil},
		{name: "case2", answer: "да", result: 2, expect: nil},

		{name: "case3", answer: "Скорее да", result: 1, expect: nil},
		{name: "case4", answer: "скорее да", result: 1, expect: nil},
		{name: "case5", answer: "Скорее Да", result: 1, expect: nil},

		{name: "case6", answer: "Возможно", result: 0, expect: nil},
		{name: "case7", answer: "возможно", result: 0, expect: nil},

		{name: "case8", answer: "Скорее нет", result: -1, expect: nil},
		{name: "case9", answer: "скорее нет", result: -1, expect: nil},
		{name: "case10", answer: "скорее Нет", result: -1, expect: nil},

		{name: "case11", answer: "Нет", result: -2, expect: nil},
		{name: "case12", answer: "нет", result: -2, expect: nil},

		{name: "case13", answer: "нте", result: 0, expect: ErrWrongAnswer},
		{name: "case14", answer: "скоере да", result: 0, expect: ErrWrongAnswer},
		{name: "case15", answer: "возможон", result: 0, expect: ErrWrongAnswer},
		{name: "case16", answer: "д", result: 0, expect: ErrWrongAnswer},
		{name: "case17", answer: "сркоее нет", result: 0, expect: ErrWrongAnswer},
		{name: "case18", answer: "А", result: 0, expect: ErrWrongAnswer},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result, err := marker.MarkResult(tc.answer)
			assert.Equal(t, tc.result, result)
			if err != nil {
				assert.ErrorIs(t, err, tc.expect)
			} else {
				assert.Equal(t, err, nil)
			}
		})
	}
}
