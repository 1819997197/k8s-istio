## HTTP REST
> * order: 底层基础服务，对外提供http接口
> * order_bff: 聚合服务(裁剪、聚合底层服务)，对外提供http接口

#### 1.Run the order service
```
$ cd order
$ CGO_ENABLED=0 GOOS=linux go build -o order-server
$ ./order-server
```

#### 2.Run the order_bff service
```
$ cd order_bff
$ CGO_ENABLED=0 GOOS=linux go build -o order-bff-server
$ ./order-bff-server
```

#### 3.Test the server
```
$ curl localhost:8080/test
hello world!
```