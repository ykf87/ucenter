package funcs

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
	"ucenter/app/config"
	"ucenter/app/safety/aess"

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
