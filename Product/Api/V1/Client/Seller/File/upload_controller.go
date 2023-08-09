package File

import (
	"Backend/Core/Constants/Keys/FileKeys"
	"Backend/Core/Utilities/Responses"
	"Backend/Product/Model/File"
	"Backend/Product/Services/External/CDN"
	File2 "Backend/Product/Services/Internal/File"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"time"
)

type UploadController struct {
	CDN         CDN.ICDNService[*s3.S3]
	FileService File2.IFileService
}

func (self *UploadController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if err := request.ParseForm(); err != nil {
		Responses.InvokeBadRequest(writer, err.Error())
		return
	}

	file, header, err := request.FormFile(FileKeys.FileKey)
	if err != nil {
		Responses.InvokeBadRequest(writer, err.Error())
		return
	}

	defer file.Close()

	var bytes []byte
	if _, err := file.Read(bytes); err != nil {
		Responses.InvokeBadRequest(writer, err.Error())
		return
	}

	url, err := self.CDN.UploadFile(bytes, header.Filename)
	if err != nil {
		Responses.InvokeInternalServerError(writer, err.Error())
		return
	}

	model := File.File{
		URL:       url,
		Filename:  header.Filename,
		Size:      header.Size,
		CreatedAt: time.Now(),
	}

	if err := self.FileService.CreateFile(&model); err != nil {
		Responses.InvokeInternalServerError(writer, err.Error())
		return
	}

	Responses.InvokeSuccess(writer, model)

}
