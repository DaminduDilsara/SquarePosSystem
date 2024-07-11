package configurations

var AppConf AppConfigurations

type AppConfigurations struct {
	AppPort int `yaml:"app_port"`
}
