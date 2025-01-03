package main

import (
	"fmt"
	"go-sample/calculator"
	"go-sample/enumeration"
	"go-sample/gomap"
	"go-sample/gostruct"
	"go-sample/slice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. ベーシックなGETリクエスト
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the server!"})
	})

	// 2. 動的パラメータを使ったGETリクエスト
	r.GET("/greet/:name", func(c *gin.Context) {
		name := c.Param("name") // URLパラメータ "name" を取得
		message := fmt.Sprintf("Hello, %s!", name)
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	// 3. POSTリクエストでJSONデータを受け取る
	r.POST("/data", func(c *gin.Context) {
		var requestData map[string]interface{}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 受け取ったデータをそのまま返す
		c.JSON(http.StatusOK, gin.H{
			"received": requestData,
		})
	})
	r.Run(":8080") // サーバーを8080番ポートで起動
}

func sample_run() {
	calculator.RunCalculator()
	slice.RunSlice()
	gomap.RunMap()
	enumeration.RunEnum()
	gostruct.RunStruct()
}
