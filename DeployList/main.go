package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os"
	"strings"
)

var namespace = os.Getenv("namespace")
var url = os.Getenv("webhook")
var scheduled = os.Getenv("scheduled")


func Robot(content string){
	dataJsonStr := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s"}}`,content)
	resp, err := http.Post(
		url,
		"application/json",
		//strings.NewReader("saas1命名空间fin-srv-dev下：\n"+dataJsonStr))
		strings.NewReader(dataJsonStr))

	if err != nil {
		fmt.Println(err)
	}
	var a []byte
	json.Unmarshal(a, resp.Body)
	fmt.Println(resp.StatusCode,a)
}


func main() {
	//config ,err := clientcmd.BuildConfigFromFlags("",clientcmd.RecommendedHomeFile)
	kubeconfig ,err := ioutil.ReadFile("/conf/config")
	if err != nil{
		panic(err)
	}
	config ,err := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	clientset , err := kubernetes.NewForConfig(config)
	if err != nil{
		panic(err)
	}

	//coreV1 := clientset.CoreV1()
	fmt.Printf("Listing deployments in namespace %q:\n", "fin-srv-dev")
	deploy ,err := clientset.AppsV1().Deployments(namespace).List(context.TODO(),metav1.ListOptions{})
	if err !=nil{
		panic(err)
		}

	var content string

	for _,d := range deploy.Items{
		//fmt.Println(d.Name,*d.Spec.Replicas)
		if *d.Spec.Replicas >= 2{
			//fmt.Printf(" * %s (副本个数：%d)\n",d.Name, *d.Spec.Replicas)
			content += fmt.Sprintf("* deployment名称: %s 副本数: %d\n", d.Name, *d.Spec.Replicas)
			//content += fmt.Sprintf("%s：%d\n",d.Name,*d.Spec.Replicas)
		}


	}
	content2 := "命名空间"+namespace+"下："
	contentsum := fmt.Sprintf("%s\n%s",content2,content)

	//fmt.Println(contentsum)

	c := cron.New()
	c.AddFunc(scheduled,func(){
		Robot(contentsum)
		//log.Println("hello word")
	})
	c.Start()
	select {

	}


	//Robot(contentsum)
}
