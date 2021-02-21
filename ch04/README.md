## HTTP REST + docker + k8s
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

#### 3.Test the server
```
1)查看服务运行状态
$ kubectl get pods
NAME                                    READY   STATUS    RESTARTS   AGE
order-bff-deployment-744989c994-dnr29   1/1     Running   0          15m
order-deployment-6d7d9b64cd-lvdc8       1/1     Running   1          35m

$ kubectl get svc
NAME           TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)     AGE
kubernetes     ClusterIP   10.96.0.1        <none>        443/TCP     4d17h
order-bff-f1   ClusterIP   10.105.143.236   <none>        8080/TCP    14m
order-svc      ClusterIP   10.96.78.144     <none>        50050/TCP   32m

2)验证服务是否可用
$ curl 10.105.143.236:8080/test
hello world!
```

#### 4.kubectl 常用命令总结
```
1)申明式资源管理
**apply**

2)命令式资源管理
创建: **create**
更新: **scale**
删除: **delete**

3)资源查看
**get、describe**

4)容器管理
**log、exec、cp**

5)集群管理
**cluster-info、version**
​```