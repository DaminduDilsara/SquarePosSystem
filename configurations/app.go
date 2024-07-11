package configurations

import "time"

type AppConfigurations struct {
	AppPort      int           `yaml:"app_port"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	ReadTimeOut  time.Duration `yaml:"read_time_out"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}
