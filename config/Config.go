package config

import (
	"fmt"
)

type Config struct {
	Mysql   Mysql   `yaml:"mysql"`
	System  System  `yaml:"system"`
	Logger  Loggers `yaml:"logger"`
	Uploads Uploads `yaml:"uploads"`
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
