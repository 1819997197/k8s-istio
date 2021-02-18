## HTTP REST + docker + k8s
> * order: 底层基础服务，对外提供http接口
> * order_bff: 聚合服务(裁剪、聚合底层服务)，对外提供http接口

#### 1.Build the order service
```
cd order
CGO_ENABLED=0 GOOS=linux go build -o order-server

docker build -t order-server:0.3 .

kubectl apply -f order-deployment.yaml
kubectl apply -f order-svc.yaml
```

#### 2.Build the order_bff service
```
cd order_bff
CGO_ENABLED=0 GOOS=linux go build -o order-bff-server

docker build -t order-bff-server:0.2 .

kubectl apply -f order-bff-deployment.yaml
kubectl apply -f order-bff-svc.yaml
```

#### 3.Test the server
```
[root@will ~]# kubectl get pods
NAME                                    READY   STATUS    RESTARTS   AGE
order-bff-deployment-744989c994-dnr29   1/1     Running   0          15m
order-deployment-6d7d9b64cd-lvdc8       1/1     Running   1          35m

[root@will ~]# kubectl get svc
NAME           TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)     AGE
kubernetes     ClusterIP   10.96.0.1        <none>        443/TCP     4d17h
order-bff-f1   ClusterIP   10.105.143.236   <none>        8080/TCP    14m
order-svc      ClusterIP   10.96.78.144     <none>        50050/TCP   32m
[root@will ~]# curl 10.105.143.236:8080/test
hello world!
```