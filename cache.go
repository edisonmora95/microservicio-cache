package main
import (
		"github.com/go-redis/redis"
		"fmt"
		)


func PongExample(client *redis.Client) {
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func IncrementarContador(client *redis.Client, key string) {
	client.HIncrBy(key, "contador", 1).Result()
}

func ListarGifs(client *redis.Client, fecha string) {
	cmd := client.Get(fecha)
	fmt.Println(cmd)
	val, err := client.Get(fecha).Result()
	fmt.Println(val, err)
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	
	gif := make(map[string]interface{})
	gif["contenido"] = "Hola"
	gif["contador"] = 0

	hash, err := client.HMSet("gif", gif).Result()
	fmt.Println(hash, err)

	val, err := client.HIncrBy("gif", "contador", 1).Result()
	fmt.Println(val, err)
	//PongExample(client)
	//ExampleClient(client)
	//ListarGifs(client, "foo")
}