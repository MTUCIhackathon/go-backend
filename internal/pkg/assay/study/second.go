package study

import (
	"errors"

	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

var (
	typeList = map[uint32]string{
		1:  "S",
		2:  "E",
		3:  "F",
		4:  "P",
		5:  "S",
		6:  "E",
		7:  "F",
		8:  "P",
		9:  "I",
		10: "T",
		11: "E",
		12: "J",
		13: "S",
		14: "F",
		15: "I",
		16: "S",
		17: "F",
		18: "E",
		19: "J",
		20: "T",
		21: "N",
		22: "F",
		23: "T",
		24: "N",
		25: "J",
		26: "E",
		27: "P",
		28: "N",
		29: "S",
		30: "I",
		31: "J",
		32: "T",
		33: "T",
		34: "N",
		35: "J",
		36: "S",
		37: "E",
		38: "P",
		39: "I",
		40: "P",
	}
)

var (
	ErrWrongPersonality                 = errors.New("wrong personality")
	ErrWrongNumbersOfMarksForSecondTest = errors.New("wrong numbers of marks: number of marks for second test type should be equal 40")
)

type Second struct {
	log      *zap.Logger
	typeList map[uint32]string
}

func NewSecond(log *zap.Logger) *Second {
	if log == nil {
		log = zap.NewNop()
		log.Named("second")
		log.Warn("logger for second assay was initialized by nil logger")
	}
	return &Second{
		log:      log,
		typeList: typeList,
	}
}

//func (s *Second) GetPersonality(marks []dto.Mark) (string, error) {
//	m := make(map[string]int8)
//	res := ""
//
//	for i := 0; i < len(marks); i++ {
//		name, err := s.getPersonality(marks[i].Order)
//		if err != nil {
//			return "", ErrWrongPersonality
//		}
//		m[name] += marks[i].Mark
//	}
//
//	for len(res) != 4 {
//		maxi := int8(-127)
//		str := ""
//		for k, v := range m {
//			if v > maxi {
//				str = k
//			}
//		}
//		res += str
//		delete(m, str)
//	}
//
//	return res, nil
//}

func (s *Second) GetPersonality(marks []dto.Mark) (string, error) {
	if len(marks) != 40 {
		s.log.Error(
			"wrong number of marks for second test: should be 40",
			zap.Int("marks length", len(marks)),
		)

		return "", ErrWrongNumbersOfMarksForSecondTest
	}

	m := make(map[string]int8)
	res := ""

	for i := 0; i < len(marks); i++ {
		name, err := s.getPersonality(marks[i].Order)
		if err != nil {
			return "", ErrWrongPersonality
		}
		m[name] += marks[i].Mark
	}

	for len(res) != 4 {
		maxi := int8(-127)
		str := ""
		for k, v := range m {
			if v > maxi {
				str = k
			}
		}
		res += str
		delete(m, str)
	}

	return res, nil
}

func (s *Second) getPersonality(num uint32) (string, error) {
	result, ok := s.typeList[num]
	if !ok {
		return "", ErrWrongPersonality
	}
	return result, nil

}
