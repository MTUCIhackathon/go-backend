package inmemory

import (
	"os"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type Option interface {
	onStart(*Cache) error
	onStop(*Cache) error
}

type optionFunc func(*Cache) error

func (f optionFunc) onStart(c *Cache) error {
	if c == nil || f == nil {
		return ErrNilReference
	}

	c.mu.Lock()
	err := f(c)
	c.mu.Unlock()

	return err
}

func (f optionFunc) onStop(c *Cache) error {
	if c == nil || f == nil {
		return ErrNilReference
	}

	c.mu.Lock()
	err := f(c)
	c.mu.Unlock()

	return err
}

type Loader struct{}

func (l Loader) onStart(c *Cache) error {
	data, err := os.ReadFile(c.config.CachePath)
	if err != nil {
		c.log.Error(
			"error reading cache file",
			zap.Error(err),
		)
		return ErrReadingFile
	}

	m := make(map[string]map[int]string, testsLength)

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		c.log.Error(
			"error while unmarshalling into data",
			zap.Error(err),
		)
		return ErrUnmarshalling
	}

	for k, v := range m {
		var (
			test dto.Test
		)

		test.Questions = make([]dto.TestQuestion, 0, len(v))
		test.Name = k
		test.Description = v[0]
		test.ID = uuid.New()

		for i := 1; i <= len(v)-1; i++ {
			text, _ := v[i]
			test.Questions = append(test.Questions, dto.TestQuestion{
				Order:    i,
				Question: text,
			})
		}
		c.data[test.ID] = test
	}

	return nil
}

func (l Loader) onStop(c *Cache) error {
	return nil
}

func WithLoader() Option {
	return Loader{}
}
