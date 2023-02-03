package Admin

import (
	"fmt"
	"github.com/deatil/doak-cron/pkg/email"
	"github.com/kataras/iris/v12"
	"mime/multipart"
	"time"
)

type UploadFileController struct {}

func (file *UploadFileController) UploadImg(ctx iris.Context)  {

	res, err := ctx.UploadFormFiles("./uploads", beforeSave)
	fmt.Printf("res:%v, err:%v \n", res, err)
	ctx.JSON(iris.Map{"code":0,"msg":"上传成功"})
	return
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	//ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	//ip = strings.Replace(ip, ".", "_", -1)
	//ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=    time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.

	file.Filename = time.Now().Format("20060102") + email.RandCode(5) + file.Filename
}
