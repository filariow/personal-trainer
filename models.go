package main

type Exercise struct {
	Name          string `yaml:"name"`
	DurationInSec int64  `yaml:"duration"`
	Preparation   int64  `yaml:"preparation"`
}
