package client

import (
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/ml/client/model"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

func (cli *PythonClient) HandlerSendResultsForFirstTest(areas []dto.Area) ([]string, error) {
	var (
		err  error
		resp model.ScientificTestMLResponse
	)

	data := make(map[string]int)
	for _, area := range areas {
		data[area.Field] = int(area.Mark)
	}

	req := model.ScientificTestMLRequest{
		Professions: data,
	}

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(cli.cfg.ML.Bind() + testScientificTestRoute)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Error(err))
		return nil, err
	}
	return resp.Professions, nil
}

func (cli *PythonClient) HandlerSendResultsForSecondTest(kind string) (*model.PersonalityTestMLResponse, error) {
	var (
		err  error
		resp model.PersonalityTestMLResponse
	)

	req := model.PersonalityTestMLRequest{
		TestResult: kind,
	}

	uri := cli.cfg.ML.Bind() + testPersonalityTestRoute
	cli.log.Debug("sending request to ml", zap.Any("req", cli.cfg.ML.Bind()+testPersonalityTestRoute))

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(uri)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Error(err))
		return nil, err
	}

	return &resp, nil
}
