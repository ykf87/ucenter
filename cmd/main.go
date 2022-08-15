// linux execute file
// env GOOS=linux GOARCH=amd64 go build
// export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && go build
// ./frpc.exe -c frpc.ini
package main

import (
	"flag"
	"log"
	"os"
	"ucenter/app/launch"
)

var configFile = flag.String("c", "config.yaml", "配置文件路径")
var port = flag.Int("port", 0, "端口")

func main() {
	// tk := "R0RLWTN4Q0JTOWRXMW5mVzAtNVBNcV9xd2d0MG1jemFTa05wcWFjdkV4czdaVzhVbFl0UzlRWDBObGxuZ3RsQg=="
	// b, _ := base64.StdEncoding.DecodeString(tk)
	// fmt.Println(aess.EcbDecrypt(string(b), nil))
	// fmt.Println(string(aess.AESKEY), aess.EcbEncrypt(`{"time":111111,"id":1,"sid":1}`, nil))

	flag.Parse()
	if *configFile == "" {
		log.Println("请指定配置文件")
		return
	}

	configExp := "./.example.yaml"
	if _, err := os.Stat(configExp); err != nil {
		if mkexp(configExp) != nil {
			log.Println("示例文件生成失败!", err)
			return
		}
	}

	launch.Start(*configFile, *port)
}

func mkexp(filepath string) error {
	config := `appname: "Your app name"
domain: "ex：http://127.0.0.1:18998"
auther: "ykf"
aeskey: "16位长度的密钥"
port: 服务端口,int类型
limit: 20
static: "./static"
staticname: "/static"
logo: "系统的logo地址,如：http://127.0.0.1:18998/static/logo.png"
mainbanner: "邮件头部图片地址,如：http://127.0.0.1:18998/static/logo.png"
lang: "默认语言: en"
country: "默认国家: US"
universalcaptcha: "万能验证码"
timezone: "默认时区：America/Adak"
datetimefmt: "默认时间格式: 2006-01-02 15:04:05"
datafmt: "默认日期格式2006-01-02"
timefmt: "默认时间格式: 15:04:05"
useim: "Im聊天组件,在下方im中定义"
useoss: "文件存储组件,下方oss中定义"
heterosexual: 搜索用户是否仅显示异性, 0表示搜索出所有性别,1表示排除和搜索账号相同性别
timefmts:
  en:
    datetimefmt: "Jan 2, 2006 3:04 PM"
    datefmt: "Jan 2, 2006"
    timefmt: "3:04 PM"
db:
- type: "mysql"
  dsn: "root:111@tcp(localhost:3306)/ucenter?charset=utf8mb4&parseTime=True&loc=Local"
- type: "sqlite"
  path: ""

redis:
  addr: "0.0.0.0:6379"
  password: "123455"
  dbname: 0

smtp:
  default:
    host: "smtp.qq.com"
    port: 465
    email: "admin@ucenter.com"
    pass: "111111"
    sender: "customer@ucenter.com"
im:
  tencent:
    id: "im方的ID"
    key: "密钥"
oss:
  minio:
    endpoint: "不带http或https开头的地址:1.1.1.1"
    url: "http://1.1.1.1 注意没有结尾反斜杠"
    accesskeyid: ""
    secret: ""
    bucket: ""
    ssl: false
imagethum://缩略图尺寸
  small:
    width: 100
    height: 100
  medium:
    width: 650
    height: 720
payment: "paypal"
payments:
  paypal:
    appid: "11111"
    secret: "2222"
    returnurl: "http://sdfdfsf.rrr"
    cancelurl: "http://sdfdfsf.ccc"
`
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = f.WriteString(config)
	if err != nil {
		return err
	}
	return nil
}

//https://www.tiktok.com/passport/web/get_qrcode/?next=https%3A%2F%2Fwww.tiktok.com&aid=1459&msToken=4W3hUJnG5DDAl994cdGnqAWAEZrK-PYXQOcRgtTQ4iF91kjeQskOF1j39X8VpqdB7VxESNzCUbIw0OahA3fq8QyjYaYShkY9zO7rFwl_cIB6_RL4AJiG1lwy5dkCp2KZLUMfdLUegXOpJ7rGGw==

//https://www.tiktok.com/passport/web/get_qrcode/?next=https%3A%2F%2Fwww.tiktok.com&aid=1459&msToken=4W3hUJnG5DDAl994cdGnqAWAEZrK-PYXQOcRgtTQ4iF91kjeQskOF1j39X8VpqdB7VxESNzCUbIw0OahA3fq8QyjYaYShkY9zO7rFwl_cIB6_RL4AJiG1lwy5dkCp2KZLUMfdLUegXOpJ7rGGw==
