package images

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"regexp"
	"strings"
	"time"
)

//上传图片的保存
func SaveFileFromUpload(savePath, saveName string, f *multipart.FileHeader) (string, error) {
	if saveName == "" {
		fns := strings.Split(f.Filename, ".")
		saveName = fmt.Sprintf("%d%06v.%s", time.Now().Unix(), rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000), fns[len(fns)-1])
	} else if strings.Contains(saveName, ".") == false {
		fns := strings.Split(f.Filename, ".")
		saveName = saveName + "." + fns[len(fns)-1]
	}
	savePath = strings.TrimRight(savePath, "/\\") + "/"
	if _, err := os.Stat(savePath); err != nil {
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	saveTo := savePath + saveName
	if _, err := os.Stat(saveTo); err == nil {
		os.Remove(saveTo)
	}

	src, err := f.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(saveTo)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}
	return saveTo, nil
}

//保存base64格式内容
func SaveFileBase64(savePath, saveName string, base64_image_content string) (string, error) {
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64_image_content)
	if !b {
		return "", errors.New("Please set the content to be modified")
	}
	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(base64_image_content), 2)
	fileType := string(allData[0][1]) //png ，jpeg 后缀获取

	if saveName == "" {
		saveName = fmt.Sprintf("%d%06v.%s", time.Now().Unix(), rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000), fileType)
	} else if strings.Contains(saveName, ".") == false {
		saveName = saveName + "." + fileType
	}
	savePath = strings.TrimRight(savePath, "/\\") + "/"
	if _, err := os.Stat(savePath); err != nil {
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	saveTo := savePath + saveName
	if _, err := os.Stat(saveTo); err == nil {
		os.Remove(saveTo)
	}
	base64Str := re.ReplaceAllString(base64_image_content, "")
	byte, _ := base64.StdEncoding.DecodeString(base64Str)

	err := ioutil.WriteFile(saveTo, byte, 0666)
	if err != nil {
		return "", err
	}
	return saveTo, nil
}
