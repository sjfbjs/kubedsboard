package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//注入需要使用的配置
type PodCtl struct {
	K8SClient *kubernetes.Clientset `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (this *PodCtl) GetList(ctx *gin.Context) goft.Json {
	//会报一个错误，不清楚是不是client-go版本的问题
	//自动注入又可以了....
	//config := configs.K8SConfig{}
	//K8SClient := config.InitClient()
	list, err := this.K8SClient.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	//list,err:=this.K8SClient.CoreV1().Nodes().List(ctx,v1.ListOptions{})
	goft.Error(err)
	//fmt.Println(err.Error())
	return list

}

func (this *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", this.GetList)
}

func (this *PodCtl) Name() string {
	return "PodCtl"
}
