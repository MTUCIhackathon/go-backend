package client

import (
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/ml/client/model"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

func (cli *PythonClient) HandlerSendResultsForFirstTest(areas []dto.Area) ([]string, error) {
	var (
		err  error
		resp model.FirstTestMLResponse
	)

	data := make(map[string]int)
	for _, area := range areas {
		data[area.Field] = int(area.Mark)
	}

	req := model.FirstTestMLRequest{
		Professions: data,
	}

	dns := cli.cfg.ML.Bind()

	//TODO logic with url
	url := dns

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(url)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Any("error", err))
		return nil, err
	}
	return resp.Professions, nil
}

func (cli *PythonClient) HandlerSendResultsForSecondTest(kind string) (*model.SecondTestMLResponse, error) {
	var (
		err  error
		resp model.SecondTestMLResponse
	)
	req := model.SecondTestMLRequest{
		TestResult: kind,
	}
	//here will be logic with url

	url := ""

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(url)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Any("error", err))
		return nil, err
	}

	return &resp, nil
}
