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
wget https://dev.mysql.com/get/mysql80-community-release-el7-1.noarch.rpm

```
Comprobar la firma y comparar con el sitio web (curiosamente usan md5 un hash que ya ha demostrado vulnerabilidades)
```
md5sum https://dev.mysql.com/get/mysql80-community-release-el7-1.noarch.rpm


Output
739dc44566d739c5d7b893de96ee6848  mysql80-community-release-el7-1.noarch.rpm
```
Comprobar con el sitio web
Una vez hecho esto instalar paquete y de ahi podremos acceder a yum para instalar mysql
```
sudo rpm -ivh mysql80-community-release-el7-1.noarch.rpm
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
sudo systemctl enable nginx
```

Además un issue que se tuvo cuando se lo configuro como reverse proxy es que linux no dejaba hacer la redirección . Para arreglarlo
```
sudo setsebool httpd_can_network_connect 1 -P 
```

**Nota**: En el caso de hacer deploy en amazon AWS es necesario instalar firewalld. Para hacerlo seguir instrucciones de https://www.tecmint.com/fix-firewall-cmd-command-not-found-error/ 


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
## Servidor web
El servidor web esta hecho en el framework Gin del lenguaje de programación Go.
Para compilarlo situarse en la ruta ginApp/server y ejecutar
```
go build -o app
```
De ahi correrlo con 
```
./app
```
Por defecto correra en el puerto 3000 pero se redireccionara al puerto 80 con nginx.
Además cabe destacar que debe estar corriendo el microservicio por lo que es necesario abrir otra consola y poner:
```
go run grpc/microservice/main.go
```
## Nginx
Para la configuración de Nginx, y una vez instalado este, se necesitan copiar dos archivos de este repositorio ubicados en la carpeta del mismo nombre: nginx.conf y go-app.conf. 

Para el primero moverlo a la ruta /etc/nginx/ pero después de haber hecho un respaldo
```
sudo cd /etc/nginx/ && mv nginx.conf nginx.conf.backup
sudo mv nginx.conf /etc/nginx/
```
Para go-app.conf moverlo a /etc/nginx/conf.d
```
sudo mv go-app.conf /etc/nginx/conf.d/
```

## Otros
Es necesario en este punto llenar tanto redis y mysql una vez instaladas . 
Para popular redis correr:
```
go run store_gifs_to_redis.go
```
Para popular mysql correr:
```
go run store_gifs_to_mysql.go
```

**Nota** Para mysql es necesario cambiar las credenciales de acceso según las que el usuario haya puesto. Para eso se cambia eso en dos archivos: store_gifs_to_mysql.go y grpc/microservice/main.go 
