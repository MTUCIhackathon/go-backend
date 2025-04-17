package study

import (
	"errors"

	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

//TODO check all methods

var (
	areaList = map[uint32]string{
		1:  "Биология",
		2:  "География",
		3:  "Геология",
		4:  "Медицина",
		5:  "Легкая и пищевая промышленность",
		6:  "Физика",
		7:  "Химия",
		8:  "Техника",
		9:  "Электро- и радиотехника",
		10: "Металлообработка",
		11: "Деревообработка",
		12: "Строительство",
		13: "Транспорт",
		14: "Авиация",
		15: "Военные специальности",
		16: "История",
		17: "Литература",
		18: "Журналистика",
		19: "Общественная деятельность",
		20: "Педагогика",
		21: "Юриспруденция",
		22: "Сфера обслуживания",
		23: "Математика",
		24: "Экономика",
		25: "Иностранные языки",
		26: "Изобразительное искусство",
		27: "Сценическое искусство",
		28: "Музыка",
		29: "Физкультура и спорт",
	}
)

var (
	ErrWrongArea                       = errors.New("wrong area: area not found")
	ErrWrongNumbersOfMarksForFirstTest = errors.New("wrong numbers of marks: number of marks for first test type should be equal 174")
)

type First struct {
	log      *zap.Logger
	areaList map[uint32]string
}

func NewFirst(log *zap.Logger) *First {
	if log == nil {
		log = zap.NewNop()
		log.Warn("logger for first assay was initialized by nil logger")
	}
	return &First{
		log:      log,
		areaList: areaList,
	}
}

func (f *First) GetAreas(marks []dto.Mark) ([]dto.Area, error) {
	if len(marks) != 174 {
		f.log.Error(
			"wrong number of marks for first test: should be 174",
			zap.Int("marks length", len(marks)),
		)

		return nil, ErrWrongNumbersOfMarksForFirstTest
	}

	sum := f.sumMark(marks)

	sortMarks := f.sortMark(sum)

	priorityAreas := make([]dto.Area, 5)

	for i := 0; i < 5; i++ {
		area, err := f.getArea(sortMarks[i].Order)
		if err != nil {
			f.log.Debug("failed to get area from sort mark", zap.Error(err))
			return nil, err
		}
		priorityAreas[i] = dto.Area{
			Field: area,
			Mark:  sortMarks[i].Mark,
		}
	}

	return priorityAreas, nil
}

func (f *First) getArea(num uint32) (string, error) {
	result, ok := f.areaList[num]
	if !ok {
		return "", ErrWrongArea
	}
	return result, nil
}

func (f *First) sumMark(marks []dto.Mark) [29]dto.Mark {
	res := [29]dto.Mark{}

	for i := 0; i < 29; i++ {
		res[i].Mark = marks[i].Mark + marks[i+29].Mark + marks[i+29*2].Mark + marks[i+29*3].Mark + marks[i+29*4].Mark + marks[i+29*5].Mark
		res[i].Order = uint32(i + 1)
	}

	return res
}

func (f *First) sortMark(marks [29]dto.Mark) [5]dto.Mark {
	for i := 1; i < len(marks); i++ {
		temp := marks[i]
		j := i - 1
		for j >= 0 && marks[j+1].Mark > temp.Mark {
			marks[j+1] = marks[j]
			j--
		}
		marks[j+1] = temp
	}

	res := [5]dto.Mark{marks[len(marks)-1], marks[len(marks)-2], marks[len(marks)-3], marks[len(marks)-4], marks[len(marks)-5]}

	return res
}
