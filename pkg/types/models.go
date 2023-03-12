package types

import "time"

type Training struct {
	Exercises []Exercise `yaml:"exercises"`
}

func (t Training) Duration() time.Duration {
	var td int64
	for _, e := range t.Exercises {
		td += e.DurationInSec
		td += e.Preparation
	}

	return time.Duration(td) * time.Second
}

type Exercise struct {
	Name          string `yaml:"name"`
	DurationInSec int64  `yaml:"duration"`
	Preparation   int64  `yaml:"preparation"`
}
