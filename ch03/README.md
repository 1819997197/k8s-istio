## HTTP REST
> * order: �ײ�������񣬶����ṩhttp�ӿ�
> * order_bff: �ۺϷ���(�ü����ۺϵײ����)�������ṩhttp�ӿ�

#### 1.Run the order service
```
cd order
CGO_ENABLED=0 GOOS=linux go build -o order-server
./order-server
```

#### 2.Run the order_bff service
```
cd order_bff
CGO_ENABLED=0 GOOS=linux go build -o order-bff-server
./order-bff-server
```

#### 3.Test the server
```
[vagrant@localhost ~]$ curl localhost:8080/test
hello world!
```