package controllers

import (
	//"gindashboard/src/configs"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)
//注入需要使用的配置
type DeploymentCtl struct {
    K8SClient *kubernetes.Clientset `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (this *DeploymentCtl) GetList(ctx *gin.Context) goft.Json  {
	//会报一个错误，不清楚是不是client-go版本的问题
	//自动注入又可以了....
	//config := configs.K8SConfig{}
	//K8SClient := config.InitClient()
	list, err:= this.K8SClient.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	//list,err:=this.K8SClient.CoreV1().Nodes().List(ctx,v1.ListOptions{})
	goft.Error(err)
	//fmt.Println(err.Error())
	return  list

}

func  (this *DeploymentCtl) Build (goft *goft.Goft)  {
	goft.Handle("GET","/deployments",this.GetList)
}


func (this *DeploymentCtl) Name() string  {
	return "DeploymentCtl"
}


