package study

import (
	"go.uber.org/zap"
)

var (
	typeList = map[int]string{
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

type Second struct {
	log      *zap.Logger
	typeList map[int]string
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
