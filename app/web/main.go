package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

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
	r.POST("/submit", func(c *gin.Context) {
		content := c.PostForm("content")
		captcha := c.PostForm("captcha")
		if content == "" || captcha == "" {
			c.String(http.StatusForbidden, "请填写您的站点并通过机器人验证")
			return
		}
		_, y := recaptcha(captcha, c.ClientIP())
		if !y {
			c.String(http.StatusForbidden, "未能通过机器人验证")
			return
		}
		serverChan(content)
	})
	r.Run()
}

func serverChan(content string) {
	params := url.Values{
		"text": {"「" + sdw.C.Name + "」有新提交"},
		"desp": {content},
	}
	http.PostForm("https://sc.ftqq.com/"+sdw.C.ServerChan+".send", params)
}

type recaptchaResp struct {
	Success  bool
	Hostname string
}

func recaptcha(gresp, ip string) (host string, flag bool) {
	resp, err := http.Post("https://www.recaptcha.net/recaptcha/api/siteverify",
		"application/x-www-form-urlencoded",
		strings.NewReader("secret="+sdw.C.Recaptcha.Secret+"&response="+gresp+"&remoteip="+ip))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var rp recaptchaResp
	err = json.Unmarshal(body, &rp)
	if err != nil {
		return
	}
	return rp.Hostname, rp.Success
}
