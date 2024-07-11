package main

import (
	"SquarePOS/configurations"
	"fmt"
)

func main() {
	conf := configurations.LoadConfigurations()

	fmt.Println(conf.AppConfig.AppPort)

}
