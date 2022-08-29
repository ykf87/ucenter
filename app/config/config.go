package config

import (
	"github.com/jinzhu/configor"
)

var Och chan bool = make(chan bool)

type SmtpConf struct {
	Host   string
	Port   int
	Email  string
	Pass   string
	Sender string
}

type OssConf struct {
	Endpoint    string
	Accesskeyid string
	Secret      string
	Bucket      string
	Url         string
	Ssl         bool
}

type ImConf struct {
	Id  string
	Key string
}

type RedisConf struct {
	Addr     string
	Password string
	Dbname   int
}

type PaymentConf struct {
	Appid     string
	Secret    string
	ReturnUrl string
	CancelUrl string
	IsPro     int
}

type ConfigStruct struct {
	APPName          string `required:"true"`
	Domain           string `required:"true"`
	Port             int    `required:"true"`
	Lang             string `default:"en"` //默认语言
	Auther           string `default:"blandal.com@gmail.com"`
	Copyright        string
	Logo             string
	Mainbanner       string
	Static           string
	Staticname       string
	Aeskey           string
	Universalcaptcha string //万能验证码
	Limit            int    `default:20`
	Country          string `default:"US"`
	Timezone         string `default:"America/Adak"`
	Datetimefmt      string `default:"2006-01-02 15:04:05"`
	Datefmt          string `default:"2006-01-02"`
	Timefmt          string `default:"15:04:05"`
	Useim            string `required:"true"` //使用的im
	Useoss           string `required:"true"`
	Heterosexual     int    `default: 1` //对于用户搜索结果,是否仅显示异性
	ShowPay          int    `default:0`  //是否开启支付

	DB []struct {
		Type string
		Dsn  string
		Path string
	} `required:"true"`

	Redis *RedisConf           `required:"true"`
	Smtp  map[string]*SmtpConf `required:"true"`
	Oss   map[string]*OssConf  `required:"true"`

	Timefmts map[string]struct {
		Datetimefmt string
		Datefmt     string
		Timefmt     string
	}

	Payment  string                  `required:"true"`
	Payments map[string]*PaymentConf `required:"true"`

	Im        map[string]*ImConf `required:"true"`
	Imagethum map[string]map[string]int
}

var Config = new(ConfigStruct)
var Cpath string

func init() {
	cc := new(ConfigStruct)
	err := configor.Load(cc, "config.yaml")
	if err == nil {
		Config = cc
	}
}

func Init(path string) (err error) {
	if Config != nil {
		return
	}
	cc := new(ConfigStruct)
	err = configor.Load(cc, path)
	if err != nil {
		return
	}
	Config = cc
	return
}
