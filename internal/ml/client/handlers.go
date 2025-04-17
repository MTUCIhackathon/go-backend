package client

import (
	"encoding/base64"
	"errors"

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

	cli.log.Debug("received response from ml ml", zap.Any("resp", resp))

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

	cli.log.Debug("get uri address", zap.Any("uri", uri))

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(uri)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Error(err))
		return nil, err
	}

	cli.log.Debug("received response from ml ml", zap.Any("resp", resp))

	return &resp, nil
}

func (cli *PythonClient) HandlerSendResultsForThirdTest(questions dto.ThirdTestAnswers) (*dto.ThirdTestQuestions, error) {
	var (
		err  error
		resp *dto.ThirdTestQuestions
	)

	req := model.AITestMLRequest{
		AQ: questions.QA,
	}

	uri := cli.cfg.ML.Bind() + aiTestGenerateRoute

	cli.log.Debug("get uri address", zap.Any("uri", uri))

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(uri)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Error(err))
		return nil, err
	}

	cli.log.Debug("received response from ml ml", zap.Any("resp", resp))

	return resp, nil
}

func (cli *PythonClient) HandlerGetResultByThirdTest(qa map[string]string) ([]string, error) {
	var (
		err  error
		resp model.AITestMLProfessionsResponse
	)
	cli.log.Debug("qa", zap.Any("map", qa))
	req := model.AITestMLProfessionsRequest{
		AQ: qa,
	}

	uri := cli.cfg.ML.Bind() + aiTestSummarizeRoute

	cli.log.Debug("get uri address", zap.Any("uri", uri))

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(uri)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Error(err))
		return nil, err
	}

	cli.log.Debug("received response from ml ml", zap.Any("resp", resp))

	return resp.Professions, nil
}

func (cli *PythonClient) HandlerGetCommonResultByML(professions [][]string) ([]string, error) {
	var (
		err  error
		resp model.AICommonProfessionsResponse
	)

	if len(professions) != 3 {
		cli.log.Debug("professions length is not correct", zap.Any("professions", professions))
		return nil, errors.New("professions length is not correct")
	}

	req := model.AICommonProfessionsRequest{
		FirstTest:  professions[0],
		SecondTest: professions[1],
		ThirdTest:  professions[2],
	}

	cli.log.Debug("prepare request to send", zap.Any("request", req))

	uri := cli.cfg.ML.Bind() + summarySummarizeRoute

	cli.log.Debug("get uri address", zap.Any("uri", uri))

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(uri)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Error(err))
		return nil, err
	}

	cli.log.Debug("received response from ml ml", zap.Any("resp", resp))

	return resp.Professions, nil

}

func (cli *PythonClient) HandlerGenerateImage(profession string) ([]byte, error) {
	var (
		err  error
		resp model.ImageGenerateResponse
	)

	req := model.ImageGenerateRequest{
		Profession: profession,
	}

	uri := cli.cfg.ML.Bind() + imageGenerateImage

	cli.log.Debug("get uri address", zap.Any("uri", uri))

	_, err = cli.cli.R().SetBody(req).SetResult(&resp).Post(uri)
	if err != nil {
		cli.log.Debug("failed to send request to ml", zap.Error(err))
		return nil, err
	}

	encResult, err := base64.StdEncoding.DecodeString(resp.ImageData)
	if err != nil {
		cli.log.Debug("failed to decode base64 image", zap.Error(err))
		return nil, err
	}

	return encResult, nil
}
