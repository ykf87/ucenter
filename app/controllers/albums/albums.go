package albums

import (
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"ucenter/app/controllers"
	"ucenter/app/uploadfile/images"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

//公共相册
func Albums(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	timezoneo, _ := c.Get("_timezone")
	timezone := timezoneo.(string)
	langobj, _ := c.Get("_lang")
	lang := langobj.(string)

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	uid := user.Id

	id := c.Param("id")
	if id != "" {
		id32, _ := strconv.Atoi(id)
		if id32 > 0 {
			uid = int64(id32)
		} else {
			controllers.ErrorNotFound(c)
			return
		}
	}

	list := models.GetAlbumList(uid, page, limit, false)
	if list == nil || len(list) < 1 {
		controllers.Success(c, nil, &controllers.Msg{Str: "No results found"})
	} else {
		for _, v := range list {
			v.Fmt(timezone, lang)
		}
		controllers.Success(c, list, &controllers.Msg{Str: "Success"})
	}
}

//私密相册
func Private(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	timezoneo, _ := c.Get("_timezone")
	timezone := timezoneo.(string)
	langobj, _ := c.Get("_lang")
	lang := langobj.(string)

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	uid := user.Id

	id := strings.Trim(c.Param("id"), "/")
	if id != "" {
		id32, _ := strconv.Atoi(id)
		if id32 < 1 {
			controllers.ErrorNotFound(c)
			return
		}
		uid = int64(id32)

		//如果要开放用户私密相册给其他用户看,则屏蔽下面4行代码
		if user.Id != uid {
			controllers.ErrorNoData(c, "Private content is not allowed to be viewed")
			return
		}
	}

	list := models.GetAlbumList(uid, page, limit, true)
	if list == nil || len(list) < 1 {
		controllers.Success(c, nil, &controllers.Msg{Str: "No results found"})
	} else {
		for _, v := range list {
			v.Fmt(timezone, lang)
		}
		controllers.Success(c, list, &controllers.Msg{Str: "Success"})
	}
}

//上传相册
func UploadAlb(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	timezoneo, _ := c.Get("_timezone")
	timezone := timezoneo.(string)

	langobj, _ := c.Get("_lang")
	lang := langobj.(string)

	path := fmt.Sprintf("%s/%s", models.ALBUMSAVEPATH, user.Invite)
	var isPrivate int

	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
		controllers.ErrorNoData(c, "")
		return
	}
	if form.File == nil {
		controllers.ErrorNoData(c, "Missing editorial content")
	}
	var files []*multipart.FileHeader
	var ok bool
	files, ok = form.File["public"]
	if !ok || len(files) == 0 {
		files, ok = form.File["private"]
		if !ok || len(files) == 0 {
			controllers.Error(c, nil, &controllers.Msg{Str: "Missing editorial content"})
			return
		}
		isPrivate = 1
	}
	filenames := images.UploadFileByFileProcess(path, files)
	ls, err := models.AddAlbumList(user.Id, filenames, isPrivate)
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: "Image upload failed"})
	} else {
		for _, v := range ls {
			v.Fmt(timezone, lang)
		}
		controllers.Success(c, ls, &controllers.Msg{Str: "Success"})
	}
}

//base64上传相册图片
func UploadAlbBase64(c *gin.Context) {
	// fmt.Println(c.Request.PostForm)
	// c.Request.ParseMultipartForm(128)
	// fmt.Println(c.Request.Form)

	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	timezoneo, _ := c.Get("_timezone")
	timezone := timezoneo.(string)

	langobj, _ := c.Get("_lang")
	lang := langobj.(string)

	path := fmt.Sprintf("%s/%s", models.ALBUMSAVEPATH, user.Invite)
	var isPrivate int

	var dts []string
	pub64 := c.PostFormArray("public")

	pri64 := c.PostFormArray("private")
	if len(pub64) > 0 {
		dts = pub64
	} else if len(pri64) > 0 {
		dts = pri64
		isPrivate = 1
	} else {
		controllers.Error(c, nil, &controllers.Msg{Str: "Missing editorial content"})
		return
	}

	filenames := images.UploadFileByBase64Process(path, dts)
	ls, err := models.AddAlbumList(user.Id, filenames, isPrivate)
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: "Image upload failed"})
	} else {
		for _, v := range ls {
			v.Fmt(timezone, lang)
		}
		controllers.Success(c, ls, &controllers.Msg{Str: "Success"})
	}
}

//删除相册
func Remove(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	var ids []string
	id := c.PostForm("id")
	if id == "" {
		idstmp := c.PostFormArray("id")
		if len(idstmp) > 0 {
			ids = idstmp
		} else {
			controllers.Error(c, nil, &controllers.Msg{Str: "Missing editorial content"})
			return
		}
	} else {
		ids = strings.Split(id, ",")
	}
	var aids []string
	for _, v := range ids {
		if v != "" {
			aids = append(aids, v)
		}
	}
	lists := models.GetAlbumListByIds(user.Id, aids)
	var removeIds []int64
	for _, v := range lists {
		removeIds = append(removeIds, v.Id)
		go images.Remove(v.Src)
	}
	err := models.RemoveAlbumByIds(user.Id, removeIds)
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: "System error, please try again later"})
	} else {
		controllers.Success(c, nil, &controllers.Msg{Str: "Success"})
	}
}

//相册私密性互转
func AlbumsExg(c *gin.Context) {
	// rs, _ := c.Get("_user")
	// user, _ := rs.(*models.UserModel)
	// timezoneo, _ := c.Get("_timezone")
	// timezone := timezoneo.(string)
}
