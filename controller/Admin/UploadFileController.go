package Admin

import (
	"fmt"
	"github.com/deatil/doak-cron/pkg/email"
	"github.com/kataras/iris/v12"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type UploadFileController struct {}
//
func (uploadFile *UploadFileController) UploadImg(ctx iris.Context)  {

	res, err := ctx.UploadFormFiles("./uploads", beforeSave)
	fmt.Printf("res:%v, err:%v \n", res, err)
	if err==nil{
		ctx.JSON(iris.Map{"code":0,"msg":"上传成功"})
		return
	}
	ctx.JSON(iris.Map{"code":1,"msg":"上传失败"})
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

	//file.Filename = time.Now().Format("20060102") + email.RandCode(5) + file.Filename
	file.Filename = time.Now().Format("20060102") + email.RandCode(5)
}
//接收处理
func (uploadFile *UploadFileController) UploadManual(ctx iris.Context)  {
	// Get the file from the request.
	file, info, err := ctx.FormFile("uploadfile")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename
	//fmt.Printf("mime:%v", info.Header)
	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	uploadDir := "./uploads/"+ time.Now().Format("20060102")
	_,errs:=os.Stat(uploadDir)
	if errs!= nil {
		os.MkdirAll(uploadDir,0777)
	}
	fileUrl := "/uploads/"+ time.Now().Format("20060102") + "/" +fname
	out, err := os.OpenFile( uploadDir + "/" +fname, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	_,cerr := io.Copy(out, file)
	if cerr == nil {
		ctx.JSON(iris.Map{"code":0,"msg":"上传成功","fileUrl":fileUrl})
		return
	}
	ctx.JSON(iris.Map{"code":1,"msg":"上传失败"})
	return
}
