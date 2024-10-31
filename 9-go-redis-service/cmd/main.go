package main

import (
	"fmt"
	"log"
	"net/http"
	"redis-service/internal/cache"
	"redis-service/internal/config"
	"redis-service/internal/handlers"
)

func main() {
	fmt.Println("Starting Go Redis App...")

	redisCache := cache.NewRedisCache(config.RedisAddr)

	cacheHandler := handlers.NewCacheHandler(redisCache)

	http.HandleFunc("/api/cache/set", cacheHandler.SetKey)
	http.HandleFunc("/api/cache/get", cacheHandler.GetKey)
	http.HandleFunc("/api/cache/delete", cacheHandler.DeleteKey)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
