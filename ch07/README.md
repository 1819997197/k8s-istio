## HTTP REST + k8s + Istio(流量转移)
> * order: 底层基础服务
> * order_bff: 聚合服务(裁剪、聚合底层服务)
> * order、order_bff服务部署在k8s集群上面，最终对外提供http接口

#### 1.Build the order service
```
$ cd order
$ CGO_ENABLED=0 GOOS=linux go build -o order-server

$ docker build -t order-server:0.3 .

$ kubectl apply -f order-deployment.yaml
$ kubectl apply -f order-svc.yaml
```

#### 2.Build the order_bff service
```
$ cd order_bff
$ CGO_ENABLED=0 GOOS=linux go build -o order-bff-server

$ docker build -t order-bff-server:0.2 .

$ kubectl apply -f order-bff-deployment.yaml
$ kubectl apply -f order-bff-svc.yaml
```

#### 3.Run the gateway
```
$ kubectl apply -f order-gateway.yaml
```

#### 4.Test the server
```
1)查看服务运行状态
$ kubectl get pods
NAME                                    READY   STATUS    RESTARTS   AGE
order-bff-deployment-744989c994-dnr29   1/1     Running   0          15m
order-deployment-6d7d9b64cd-lvdc8       1/1     Running   1          35m

$ kubectl get svc
NAME           TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)     AGE
NAME            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)     AGE
kubernetes      ClusterIP   10.96.0.1        <none>        443/TCP     165d
order-bff-svc   ClusterIP   10.108.147.101   <none>        8080/TCP    26m
order-svc       ClusterIP   10.108.168.120   <none>        50050/TCP   35m

$ kubectl get gateway
NAME            AGE
order-gateway   12s

$ kubectl get virtualservice
NAME           GATEWAYS          HOSTS   AGE
order-bff-vs   [order-gateway]   [*]     20s

2)验证服务是否可用
// 集群外部访问(使用服务的node port访问网关)
$ export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
$ export INGRESS_HOST=$(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0].status.hostIP}')
// 打开浏览器访问 http://$INGRESS_HOST:$INGRESS_PORT/test 即可(或者curl访问)

3)查看服务流量
// 进入集群容器，打开日志
$ kubectl exec -it order-bff-deployment-v1-5c6fd84c4c-8wrvm bash
$ tail -f log/api.log

// curl 测试
curl http://$INGRESS_HOST:$INGRESS_PORT/test //路由到v2版本
curl http://$INGRESS_HOST:$INGRESS_PORT/will //路由到v1版本
```

#### 5.清除
```
$ kubectl delete -f order-gateway.yaml
$ kubectl delete -f order/order-svc.yaml
$ kubectl delete -f order/order-deployment.yaml
$ kubectl delete -f order_bff/order-bff-svc.yaml
$ kubectl delete -f order_bff/order-bff-deployment.yaml
```