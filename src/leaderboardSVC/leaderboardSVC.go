package LeaderboardSVC

import (
	"context"
	"fmt"
	"log"
	"net/http"

	DB "Microservice/db"

	"github.com/gin-gonic/gin"

	"github.com/redis/go-redis/v9"
)

type LeaderboardType int

const (
	None LeaderboardType = iota
	BlackrockCaverns
	ThroneoftheTides
)

type ResponsePlayers struct {
	PlayerID int             `json:"id"`
	Type     LeaderboardType `json:"type"`
	Time     int             `json:"time"`
}

type UpdateRequestPlayers struct {
	PlayerID int             `json:"id"`
	Type     LeaderboardType `json:"type"`
	Time     int             `json:"time"`
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

func getName(leaderboard LeaderboardType) string {
	return []string{
		"None",
		"BlackrockCaverns",
		"ThroneoftheTides",
	}[leaderboard]
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func UpdateScore(c *gin.Context) {

	ctx := context.Background()

	//get post json
	var req UpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{"Bad request"})
		return
	}
	var resp Response
	for _, player := range req.Players {
		member := fmt.Sprint(player.PlayerID)
		leaderboard_key := "LB_" + getName(player.Type)

		_, err := DB.Rdb.ZAdd(ctx, leaderboard_key, redis.Z{
			Member: member,
			Score:  float64(player.Time),
		}).Result()

		if err != nil {
			//response failed json
			resp.Message = err.Error()
			continue
		}

		resp.Players = append(resp.Players, ResponsePlayers{
			PlayerID: player.PlayerID,
			Type:     player.Type,
			Time:     int(player.Time),
		})

	}

	if resp.Message == "" {
		resp.Success = true
		resp.Message = "Score updated successfully"
	}
	c.JSON(http.StatusOK, resp)
}

type GetScoreRequestPlayers struct {
	Type     LeaderboardType `json:"type"`
	PlayerID int             `json:"id"`
}

type GetScoreRequest struct {
	Players []GetScoreRequestPlayers `json:"players"`
}

func GetScore(c *gin.Context) {

	ctx := context.Background()

	//get post json
	var req GetScoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{"Bad request"})
		return
	}

	var resp Response

	for _, player := range req.Players {
		leaderboard_key := "LB_" + getName(player.Type)

		score, err := DB.Rdb.ZScore(ctx, leaderboard_key, fmt.Sprint(player.PlayerID)).Result()
		if err != nil {
			continue
		}

		resp.Players = append(resp.Players, ResponsePlayers{
			PlayerID: player.PlayerID,
			Type:     player.Type,
			Time:     int(score),
		})

	}

	if resp.Message == "" {
		resp.Success = true
		resp.Message = "Score retrieved successfully"
	}

	c.JSON(http.StatusOK, resp)
}
