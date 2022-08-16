package funcs

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
	"ucenter/app/config"
	"ucenter/app/safety/aess"

	"github.com/gin-gonic/gin"

	"github.com/tidwall/gjson"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//删除slice重复值
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

//获取文件的 Content-Type
func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

//随机数
func Random(min, max int64) int64 {
	return rand.Int63n(max-min-1) + min + 1
}

//邀请url生成
func InviUrl(uid string) string {
	tk := aess.EcbEncrypt(fmt.Sprintf(`{"time":%d,"code":"%s"}`, time.Now().Unix(), uid), nil)
	url := fmt.Sprintf("%s/invitation?f=%s", config.Config.Domain, tk)
	return url
}

//解析邀请url
//timeout 邀请过期时间,单位秒
func DeInviUrl(str string, timeout int64) (invocode string, err error) {
	if str == "" {
		err = errors.New("Parse error")
		return
	}
	jsonStr := aess.EcbDecrypt(str, nil)
	if jsonStr == "" {
		err = errors.New("Parse error")
		return
	}
	gjsons := gjson.Parse(jsonStr).Map()
	if gjsons["code"].Exists() == false {
		err = errors.New("Parse error")
		return
	}
	if timeout > 0 {
		getTime := gjsons["time"].Int()
		if (time.Now().Unix() - getTime) > timeout {
			err = errors.New("Expired link address")
			return
		}
	}
	invocode = gjsons["code"].String()
	return
}

//数据库ip转ipv4
func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

//IPV4转数据库ip
func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

//根据头部信息生成md5信息,作为每次不同登录状态的记录依据
func UserDeviceMd5(c *gin.Context) string {
	deviceid := c.GetHeader("deviceid")
	if deviceid == "" {
		return ""
	}

	ip := c.ClientIP()
	return string(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", deviceid, ip)))))
}

//网络请求
//发起网络请求
//如果发起的是get请求,uri请自行拼接
//uri为完整的http连接地址
func Request(method, uri string, data []byte, header map[string]string, proxy string) ([]byte, error) {
	var body io.Reader
	if method == "POST" && data != nil {
		// cont, err := json.Marshal(data)
		// if err == nil {
		// 	body = bytes.NewBuffer(cont)
		// }
		body = bytes.NewBuffer(data)
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	}}

	if proxy != "" {
		proxyUrl, err := url.Parse(proxy)
		if err == nil { //使用传入代理
			tr.Proxy = http.ProxyURL(proxyUrl)
		}
	}

	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest(method, uri, body)
	if header == nil {
		header = make(map[string]string)
		header["Content-Type"] = "application/json"
		header["accept-language"] = "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7"
		header["pragma"] = "no-cache"
		header["cache-control"] = "no-cache"
		header["upgrade-insecure-requests"] = "1"
		header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36"
	} else {
		// header["Content-Type"] = "application/json"
		// header["accept-language"] = "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7"
		// header["pragma"] = "no-cache"
		// header["cache-control"] = "no-cache"
		// header["upgrade-insecure-requests"] = "1"
		// header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36"
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, errors.New(fmt.Sprintf("请求发生错误:\r\n\turi: %s.\r\n\thttpcode: %d.\r\n\tmessage: %s", uri, resp.StatusCode, string(respbody)))
	}
	return respbody, nil
}
