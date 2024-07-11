package main

import (
	"SquarePosSystem/configurations"
	"fmt"
)

func main() {
	conf := configurations.LoadConfigurations()

	fmt.Println(conf.AppConfig.AppPort)

}
