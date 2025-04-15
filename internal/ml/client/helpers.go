package client

import (
	"errors"
	"fmt"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/style/kind"
)

func (cli *PythonClient) getRouteForTests(k kind.Type) (string, error) {
	switch k {
	case kind.FirstOrder:
		return cli.cfg.Route.FirstRoute, nil
	case kind.SecondOrder:
		return cli.cfg.Route.SecondRoute, nil
	case kind.ThirdOrder:
		return cli.cfg.Route.ThirdRoute, nil
	default:
		return "", errors.New("unknown route")

	}
}

func (cli *PythonClient) getRouteForSummarize() (string, error) {
	return cli.cfg.Route.SummarizeRoute, nil
}

func (cli *PythonClient) getServerAddress(dns string, route string) string {
	return fmt.Sprintf("%s/%s", dns, route)
}
