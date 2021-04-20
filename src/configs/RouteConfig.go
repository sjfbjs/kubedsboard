package configs

import (
	"gindashboard/src/controllers"
	"github.com/shenyisyn/goft-gin/goft"
)

type RouterConfig struct {
	Goft *goft.Goft `inject:"-"`
	ctl  *controllers.DeploymentCtl
	//IndexClass * `inject:"-"`
}

func NewRouterConfig() *RouterConfig {
	return &RouterConfig{}
}
func (this *RouterConfig) DeploymentsRoutes() interface{} {
	this.Goft.Handle("GET", "/deployments", this.ctl.GetList)
	//this.Goft.Handle("GET", "/b", this.IndexClass.TestA)
	return goft.Empty
}
