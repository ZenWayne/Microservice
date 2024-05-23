package main

import (
	LeaderboardSVC "Microservice/leaderboardSVC"
	RedisClient "Microservice/redis"
	"log"
	_log "log"
	"net/http"
	"os"
)

func main() {
	logger := _log.New(os.Stdout, "MSVC: ", _log.LstdFlags)
	RedisClient.Init()

	leaderboardSVC := LeaderboardSVC.LeaderboardSVC{L: logger}
	http.HandleFunc("/leaderboard/update", leaderboardSVC.UpdateScore)
	http.HandleFunc("/leaderboard/get", leaderboardSVC.GetScore)

	logger.Println("Starting Server")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
