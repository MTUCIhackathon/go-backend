package client

import (
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/ml/client/model"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/style/kind"
)

// TODO need to think how to realize input python routers, maybe put in config and add func for switch type test
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

	route, err := cli.getRouteForTests(kind.FirstOrder)
	if err != nil {
		cli.log.Debug("failed to get route for ml", zap.Any("error", err))
		return nil, err
	}

	url := cli.getServerAddress(dns, route)

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(url)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Any("error", err))
		return nil, err
	}
	return resp.Professions, nil
}
