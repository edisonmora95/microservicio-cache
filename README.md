# microservicio-cache

Proyecto de Sistemas Distribuidos de ESPOL
Implementación de un microservicio de una caché para reducir latencia de acceso a una base de datos

## Pre-requisitos

### Go y sus dependencias

El proyecto esta hecho en Go tanto en el servidor web como en el microservicio por lo que es necesario instalar go. Para ello seguir instrucciones de la página oficial

https://golang.org/doc/install 

Una vez instalado seguir con las dependencias

Instalar net
```
go get -u golang.org/x/net
```
Instalar go-sql-driver
```
go get -u github.com/go-sql-driver/mysql
```
Instalar go-redis
```
go get -u github.com/go-redis/redis
```
Instalar grpc-go
```
go get -u google.golang.org/grpc
```
Instalar plugin de protocol buffers para go
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
Instalar gin
```
go get -u github.com/gin-gonic/gin
```

## Otras dependencias

### Protocol buffers

Para instalar protocol buffers en Centos7 linux seguir las instrucciones de:
https://github.com/google/protobuf/blob/master/src/README.md 

### Redis
Para la instalacion de Redis seguir:
```
sudo yum install epel-release
sudo yum update
sudo yum install redis
```
Para habilitarlo y que se inicie al booteo
```
sudo systemctl start redis
sudo systemctl enable redis
```

### Mysql
Para instalar mysql primero visitar https://dev.mysql.com/downloads/repo/yum/ 
Una vez ahi descargar la version para centos7 (linux7)
```
wget https://dev.mysql.com/get/mysql57-community-release-el7-9.noarch.rpm
```
Comprobar la firma y comparar con el sitio web (curiosamente usan md5 un hash que ya ha demostrado vulnerabilidades)
```
md5sum mysql57-community-release-el7-9.noarch.rpm

Output
1a29601dc380ef2c7bc25e2a0e25d31e  mysql57-community-release-el7-9.noarch.rpm
```
Comprobar con el sitio web
Una vez hecho esto instalar paquete y de ahi podremos acceder a yum para instalar mysql
```
sudo rpm -ivh mysql57-community-release-el7-9.noarch.rpm
sudo yum install mysql-server
```

Iniciar Mysql
```
sudo systemctl start mysqld
```
Nota: no es necesario hacer que se inicie con el boot, ya mysql lo hace por defecto

### Nginx
Para instalar Nginx
```
sudo yum install epel-release //en el caso de no haberlo instalado con redis
sudo yum install nginx
```
Para iniciarlo
```
sudo systemctl start nginx
```
Para permitir en el firewall
```
sudo firewall-cmd --permanent --zone=public --add-service=http 
sudo firewall-cmd --permanent --zone=public --add-service=https
sudo firewall-cmd --reload
```
Para que se inicie con el booteo
```
sudo systemctl start nginx
```

Además un issue que se tuvo cuando se lo configuro como reverse proxy es que linux no dejaba hacer la redirección . Para arreglarlo
```
sudo setsebool httpd_can_network_connect 1 -P 
```


## Microservicio
Es un microservicio escrito en Go usando la implementación de grpc.
Basado en el proto que se encuentra en grpc/proto/redisapp.proto

Se puede volver a generar el código con:
```
protoc -I grpc/proto/ grpc/proto/redisapp.proto --go_out=plugins=grpc:grpc/proto
```

Para correr el servidor
```
go run grpc/microservice/main.go
```
Para probar con un cliente stub
```
go run grpc/client/main.go
```
