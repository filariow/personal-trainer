package spec

import (
	"context"
	"io"

	"github.com/filariow/personal-trainer/pkg/types"
	"gopkg.in/yaml.v3"
)

func Read(ctx context.Context, reader io.Reader) (*types.Training, error) {
	t := &types.Training{}
	if err := yaml.NewDecoder(reader).Decode(t); err != nil {
		return nil, err
	}

	return t, nil
}
