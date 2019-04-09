package sdw

import "github.com/spf13/viper"

//Link 链接
type Link struct {
	Name string
	URL  string `mapstructure:"url"`
	Desc string
}

//Cat 链接分类
type Cat struct {
	Name  string
	Links []Link
}

//Config 网站配置
type Config struct {
	Name         string
	Slogan       string
	BuildVersion string `mapstructure:"-"`
	Aside        string
	ServerChan   string `mapstructure:"server_chan"`
	Recaptcha    struct {
		SiteKey string `mapstructure:"site_key"`
		Secret  string
	}
	Stats     string
	AsideCats []Cat `mapstructure:"aside_cats"`
	MainCats  []Cat `mapstructure:"main_cats"`
}

//C 全站配置
var C Config

//BuildVersion 构建版本
var BuildVersion = "_BuildVersion_"

func init() {
	viper.SetConfigFile("data/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
	C.BuildVersion = BuildVersion[:8]
}
