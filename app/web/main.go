package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naiba/sdw"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("data/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&sdw.C)
	if err != nil {
		panic(err)
	}
	log.Println(sdw.C)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("resource/template/*")
	r.Static("/static", "resource/static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Conf": sdw.C,
		})
	})
	r.Run()
}
