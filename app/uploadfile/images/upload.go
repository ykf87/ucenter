package images

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"regexp"
	"strings"
	"time"
	"ucenter/app/config"
	"ucenter/app/funcs"
	"ucenter/app/oss"
)

//上传图片的保存
func SaveFileFromUpload(savePath, saveName string, f *multipart.FileHeader, ch chan string, maxch chan byte) (string, error) {
	if maxch != nil {
		maxch <- 0
	}
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
			if ch != nil {
				ch <- ""
			}
			log.Println(err)
			return "", err
		}
	}

	saveTo := savePath + saveName
	if _, err := os.Stat(saveTo); err == nil {
		os.Remove(saveTo)
	}

	src, err := f.Open()
	if err != nil {
		if ch != nil {
			ch <- ""
		}
		log.Println(err)
		return "", err
	}

	out, err := os.Create(saveTo)
	if err != nil {
		if ch != nil {
			ch <- ""
		}
		log.Println(err)
		return "", err
	}

	_, err = io.Copy(out, src)
	if err != nil {
		if ch != nil {
			ch <- ""
		}
		log.Println(err)
		return "", err
	}
	src.Close()
	out.Close()

	clouldPath := strings.Trim(saveTo, "./ ")
	clouldPaths := strings.Split(clouldPath, "/")
	clouldPaths = clouldPaths[1:]
	clouldPath = strings.Join(clouldPaths, "/")
	newFileName, err := OssUploadByFileName(saveTo, clouldPath)
	if err == nil {
		os.Remove(saveTo)
		saveTo = newFileName
	}
	if ch != nil {
		ch <- saveTo
	}
	return saveTo, nil
}

//保存base64格式内容
func SaveFileBase64(savePath, saveName string, base64_image_content string, ch chan string, maxch chan byte) (string, error) {
	if maxch != nil {
		maxch <- 0
	}
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64_image_content)
	if !b {
		if ch != nil {
			ch <- ""
		}
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
			if ch != nil {
				ch <- ""
			}
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
		if ch != nil {
			ch <- ""
		}
		return "", err
	}

	newFileName, err := OssUploadByFileName(saveTo, strings.Trim(saveTo, "./ "))
	if err == nil {
		os.Remove(saveTo)
		saveTo = newFileName
	}
	if ch != nil {
		ch <- saveTo
	}
	return saveTo, nil
}

//协程上传文件
func UploadFileByFileProcess(savePath string, f []*multipart.FileHeader) (filenames []string) {
	max := len(f)
	if max < 1 {
		return
	}
	maxCh := make(chan byte, 10)
	ch := make(chan string)
	for _, v := range f {
		name := fmt.Sprintf("%d-%d-%d-%d-%d", funcs.Random(10000, 99999), funcs.Random(1000, 9999), funcs.Random(1000, 9999), funcs.Random(1000, 9999), time.Now().Unix())
		go SaveFileFromUpload(savePath, name, v, ch, maxCh)
	}
	for {
		max = max - 1
		if max < 0 {
			break
		}
		str := <-ch
		<-maxCh
		if str != "" {
			filenames = append(filenames, str)
		}
	}
	return
}

//协程上传文件-base64格式
func UploadFileByBase64Process(savePath string, f []string) (filenames []string) {
	max := len(f)
	if max < 1 {
		return
	}
	maxCh := make(chan byte, 10)
	ch := make(chan string)
	for _, v := range f {
		name := fmt.Sprintf("%d-%d-%d-%d-%d", funcs.Random(10000, 99999), funcs.Random(1000, 9999), funcs.Random(1000, 9999), funcs.Random(1000, 9999), time.Now().Unix())
		go SaveFileBase64(savePath, name, v, ch, maxCh)
	}
	for {
		max = max - 1
		if max < 0 {
			break
		}
		str := <-ch
		<-maxCh
		if str != "" {
			filenames = append(filenames, str)
		}
	}
	return
}

//上传到oss
func OssUploadByFileName(filename, ossSaveFileNameWithPath string) (string, error) {
	object, err := oss.GetOss(config.Config.Useoss)
	if err != nil {
		return "", err
	}
	return object.Upload(filename, ossSaveFileNameWithPath)
}

//获取文件全路径
func FullPath(filename string) string {
	if filename == "" {
		return ""
	}
	object, err := oss.GetOss(config.Config.Useoss)
	if err != nil {
		return strings.TrimRight(config.Config.Domain, "/") + "/" + filename
	}
	return object.Url(filename)
}

//删除文件
func Remove(filename string) error {
	object, err := oss.GetOss(config.Config.Useoss)
	if err != nil {
		return os.Remove(filename)
	}
	return object.Remove(filename)
}
