package albums

import (
	"fmt"
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

	list := models.GetAlbumList(user.Id, page, limit, false)
	if list == nil || len(list) < 1 {
		controllers.Resp(c, nil, &controllers.Msg{Str: "No results found"}, 404)
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

	list := models.GetAlbumList(user.Id, page, limit, true)
	if list == nil || len(list) < 1 {
		controllers.Resp(c, nil, &controllers.Msg{Str: "No results found"}, 404)
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

	form, _ := c.MultipartForm()
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
