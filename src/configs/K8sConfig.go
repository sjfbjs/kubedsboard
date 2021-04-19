package configs

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type K8SConfig struct {

}

func NewK8SConfig() *K8SConfig {
	return &K8SConfig{}
}

func (*K8SConfig) InitClient() *kubernetes.Clientset{
	//k8s地址
	config := &rest.Config{
		//使用代理
		// nohup  kubectl proxy --port=9999 --address='0.0.0.0' --accept-hosts='^*$'  &
		Host: "http://192.168.0.211:9999",
		//APIPath: "api",
	}
	c,err:=kubernetes.NewForConfig(config)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(c)
	if c == nil {
		fmt.Println("未获取到k8s链接")
	}
	//list,err:=c.AppsV1().Deployments("default").List(context.Background(),v1.ListOptions{})
	//fmt.Println(list)
	return  c
}
