## istio安装与使用


#### 一、下载Istio
```
1.从 [下载](https://github.com/istio/istio/releases) 页面下载适用于您操作系统的安装文件，并解压缩

2.移至Istio软件包目录
$ cd istio-1.9.0

安装目录包含：
    示例应用程序 samples/
    客户端二进制文件。bin/

3.配置环境变量
$ export PATH=$PWD/bin:$PATH
```

#### 二、安装Istio
```
1.使用demo配置文件安装，若需其它配置文件，请参考[其它配置](https://istio.io/latest/docs/setup/additional-setup/config-profiles/)
$ istioctl install --set profile=demo -y
✔ Istio core installed
✔ Istiod installed
✔ Egress gateways installed
✔ Ingress gateways installed
✔ Installation complete

2.添加名称空间标签，以指示Istio在以后部署应用程序时自动注入Envoy sidecar代理
$ kubectl label namespace default istio-injection=enabled
namespace/default labeled
```

#### 三、部署示例应用
```
1.部署Bookinfo示例应用程序(istio安装包中samples/目录下的示例)
$ kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
service/details created
serviceaccount/bookinfo-details created
deployment.apps/details-v1 created
service/ratings created
serviceaccount/bookinfo-ratings created
deployment.apps/ratings-v1 created
service/reviews created
serviceaccount/bookinfo-reviews created
deployment.apps/reviews-v1 created
deployment.apps/reviews-v2 created
deployment.apps/reviews-v3 created
service/productpage created
serviceaccount/bookinfo-productpage created
deployment.apps/productpage-v1 created

2.查看服务的状态(第一次部署需拉取镜像，耗时会久一些)
$ kubectl get pods
NAME                              READY   STATUS    RESTARTS   AGE
details-v1-79f774bdb9-dzs29       2/2     Running   0          4m22s
productpage-v1-6b746f74dc-9g6q4   2/2     Running   0          4m21s
ratings-v1-b6994bb9-d4zrs         2/2     Running   0          4m22s
reviews-v1-545db77b95-lfc7f       2/2     Running   0          4m22s
reviews-v2-7bf8c9648f-5kcc6       2/2     Running   0          4m22s
reviews-v3-84779c7bbc-c6gc5       2/2     Running   0          4m22s

$ kubectl get svc
NAME          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
details       ClusterIP   10.103.17.29     <none>        9080/TCP   5m16s
kubernetes    ClusterIP   10.96.0.1        <none>        443/TCP    7d16h
productpage   ClusterIP   10.102.139.194   <none>        9080/TCP   5m15s
ratings       ClusterIP   10.106.170.44    <none>        9080/TCP   5m16s
reviews       ClusterIP   10.101.201.41    <none>        9080/TCP   5m16s

3.验证服务是否正常
$ kubectl exec "$(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}')" -c ratings -- curl -sS productpage:9080/productpage | grep -o "<title>.*</title>"
<title>Simple Bookstore App</title>

或者直接通过service的ip+port访问
$ curl -sS 10.102.139.194:9080/productpage | grep -o "<title>.*</title>"
<title>Simple Bookstore App</title>

4.清除服务
$ kubectl delete -f samples/bookinfo/platform/kube/bookinfo.yaml
service "details" deleted
serviceaccount "bookinfo-details" deleted
deployment.apps "details-v1" deleted
service "ratings" deleted
serviceaccount "bookinfo-ratings" deleted
deployment.apps "ratings-v1" deleted
service "reviews" deleted
serviceaccount "bookinfo-reviews" deleted
deployment.apps "reviews-v1" deleted
deployment.apps "reviews-v2" deleted
deployment.apps "reviews-v3" deleted
service "productpage" deleted
serviceaccount "bookinfo-productpage" deleted
deployment.apps "productpage-v1" deleted
//这时候再通过kubectl get pods/svc 查看，服务已经注销了
```