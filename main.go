package main

import (
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


	fmt.Printf("%-v", string(body))

}
