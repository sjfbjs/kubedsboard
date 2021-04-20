package main

import (
	"gindashboard/src/configs"
	"gindashboard/src/controllers"
	"github.com/shenyisyn/goft-gin/goft"
)

func main() {

	//configure注解
	goft.Ignite().Config(configs.NewK8SConfig()).Mount(
		"v1", controllers.NewDeploymentCtl()).Config().Launch()
	//goft.Ignite().Config(configs.NewK8SConfig()).Mount("v1").Config(routers.NewRouterConfig()).Launch()
}
