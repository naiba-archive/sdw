package sdw

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
	Name      string
	Slogan    string
	Aside     string
	AsideCats []Cat `mapstructure:"aside_cats"`
	MainCats  []Cat `mapstructure:"main_cats"`
}

var C Config
