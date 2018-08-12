# microservicio-cache

Proyecto de Sistemas Distribuidos de ESPOL
Implementación de un microservicio de una caché para reducir latencia de acceso a una base de datos


## Microservicio
Es un microservicio escrito en Go usando la implementación de grpc.
Basado en el proto que se encuentra en grpc/proto/redisapp.proto

Se puede volver a generar el código con:
```
protoc -I proto/ proto/redisapp.proto --go_out=plugins=grpc:proto
```

Para correr el servidor
```
go run grpc/microservice/main.go
```
Para probar con un cliente stub
```
go run grpc/client/main.go
```

Para poder correr el servidor debe estar instalado redis y go-redis