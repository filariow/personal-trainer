package types

type Training struct {
	Exercises []Exercise `yaml:"exercises"`
}

type Exercise struct {
	Name          string `yaml:"name"`
	DurationInSec int64  `yaml:"duration"`
	Preparation   int64  `yaml:"preparation"`
}
