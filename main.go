package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)

func main() {
	url := "https://bgg-json.azurewebsites.net/thing/31260"

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error Request:", err)
    return
	}

	// resp.Bodyはクローズすること。クローズしないとTCPコネクションを開きっぱなしになる。
	defer resp.Body.Close()

	// 200 OK 以外の場合はエラーメッセージを表示して終了
	if resp.StatusCode != 200 {
			fmt.Println("Error Response:", resp.Status)
			return
	}

	// Response Body を読み取り
	body, _ := io.ReadAll(resp.Body)

	// JSONを構造体にエンコード
	var boardgames BoardGame
	err = json.Unmarshal([]byte(body), &boardgames)
	if err != nil {
		fmt.Println("Parse Error:", err)
		return
}

	fmt.Println(boardgames.Name)

}

type BoardGame struct {
	GameID            int      `json:"gameId"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Image             string   `json:"image"`
	Thumbnail         string   `json:"thumbnail"`
	MinPlayers        int      `json:"minPlayers"`
	MaxPlayers        int      `json:"maxPlayers"`
	PlayingTime       int      `json:"playingTime"`
	Mechanics         []string `json:"mechanics"`
	IsExpansion       bool     `json:"isExpansion"`
	YearPublished     int      `json:"yearPublished"`
	BggRating         float64  `json:"bggRating"`
	AverageRating     float64  `json:"averageRating"`
	Rank              int      `json:"rank"`
	Designers         []string `json:"designers"`
	Publishers        []string `json:"publishers"`
	Artists           []string `json:"artists"`
	PlayerPollResults []PlayerPoll `json:"playerPollResults"`
}

type PlayerPoll struct {
	NumPlayers            int  `json:"numPlayers"`
	Best                  int  `json:"best"`
	Recommended           int  `json:"recommended"`
	NotRecommended        int  `json:"notRecommended"`
	NumPlayersIsAndHigher bool `json:"numPlayersIsAndHigher"`
}
