package LeaderboardSVC

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	RedisClient "Microservice/redis"
	"time"
)

type LeaderboardType int

const (
	Global LeaderboardType = iota
	Personal
)

type ResponsePlayers struct {
	PlayerID int             `json:"id"`
	Type     LeaderboardType `json:"type"`
	Score    int             `json:"score"`
}

type UpdateRequestPlayers struct {
	PlayerID int             `json:"id"`
	Type     LeaderboardType `json:"type"`
	Score    int             `json:"score"`
}

type UpdateRequest struct {
	Players []UpdateRequestPlayers `json:"players"`
}

type Response struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Players []ResponsePlayers `json:"players"`
}

type LeaderboardSVC struct {
	L *log.Logger
}

var LeaderboardSVCInstance *LeaderboardSVC

func (leaderboard *LeaderboardSVC) UpdateScore(rw http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	if r.Method != "PUT" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//get post json
	var req UpdateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		leaderboard.L.Printf("Error decoding request body: %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var resp Response
	for _, player := range req.Players {
		leaderboard_key := ""
		member := ""
		if player.Type == Global {
			leaderboard_key = "LB_Global"
			member = fmt.Sprint(player.PlayerID)
		} else {
			leaderboard_key = fmt.Sprintf("LB_Personal_%d", player.PlayerID)
			member = fmt.Sprint(time.Now().Unix())
		}

		leaderboard.L.Printf("UpdateScore leaderboard_key: %s, member: %s", leaderboard_key, member)

		// Process the request data
		// ...

		var resp Response
		increment_res, err := RedisClient.Rdb.ZIncrBy(ctx, leaderboard_key, float64(player.Score), member).Result()

		if err != nil {
			//response failed json
			leaderboard.L.Printf("failed to update score player id %d err: %s", player.PlayerID, err.Error())
			resp.Message = err.Error()
			continue
		}

		resp.Players = append(resp.Players, ResponsePlayers{
			PlayerID: player.PlayerID,
			Type:     player.Type,
			Score:    int(increment_res),
		})
	}

	rw.Header().Set("Content-Type", "application/json")
	if resp.Message == "" {
		resp.Success = true
		resp.Message = "Score updated successfully"
	}
	json.NewEncoder(rw).Encode(resp)
}

type GetScoreRequestPlayers struct {
	PlayerID int             `json:"id"`
	Type     LeaderboardType `json:"type"`
}

type GetScoreRequest struct {
	Players []GetScoreRequestPlayers `json:"players"`
}

func (leaderboard *LeaderboardSVC) GetScore(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	//get post json
	var req GetScoreRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		leaderboard.L.Printf("Error decoding request body: %v", err)
	}

	var resp Response

	for _, player := range req.Players {
		leaderboard_key := ""
		member := ""
		if player.Type == Global {
			leaderboard_key = "LB_Global"
			member = fmt.Sprint(player.PlayerID)
		} else {
			leaderboard_key = fmt.Sprintf("LB_Personal_%d", player.PlayerID)
			member = fmt.Sprint(time.Now().Unix())
		}
		leaderboard.L.Printf("GetScore leaderboard_key: %s, member: %s", leaderboard_key, member)

		// Process the request data
		score, err := RedisClient.Rdb.ZScore(ctx, leaderboard_key, member).Result()

		if err != nil {
			leaderboard.L.Printf("failed to get score player id %d", player.PlayerID)
			resp.Message = err.Error()
			continue
		}

		resp.Players = append(resp.Players, ResponsePlayers{
			PlayerID: player.PlayerID,
			Type:     player.Type,
			Score:    int(score),
		})
	}

	if resp.Message == "" {
		resp.Success = true
		resp.Message = "Score retrieved successfully"
	}

	rw.Header().Set("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(resp)
}
