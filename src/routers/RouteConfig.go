package routers

import (
	"gindashboard/src/controllers"
	"github.com/shenyisyn/goft-gin/goft"
)

type RouterConfig struct {
	Goft          *goft.Goft                 `inject:"-"`
	DeploymentCtl *controllers.DeploymentCtl `inject:"-"`
	PodCtl        *controllers.PodCtl        `inject:"-"`
	//IndexClass * `inject:"-"`
}

func NewRouterConfig() *RouterConfig {
	return &RouterConfig{}
}
func (this *RouterConfig) DeploymentsRoutes() interface{} {
	this.Goft.Handle("GET", "/deployments", this.DeploymentCtl.GetList)
	this.Goft.Handle("GET", "/pods", this.PodCtl.GetList)
	//this.Goft.Handle("GET", "/b", this.IndexClass.TestA)
	return goft.Empty
}
