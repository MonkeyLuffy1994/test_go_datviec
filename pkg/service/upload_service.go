package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"test_go_datviec/pkg/config"
	"test_go_datviec/pkg/model"
)

type Server struct {
	Cof config.Config
}

func (s *Server) UploadFile(ctx *gin.Context) {
	fmt.Println("File Upload Endpoint Hit")
	r := ctx.Request

	err := r.ParseMultipartForm(10 * 1024)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	file, handler, err := r.FormFile("files")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	go saveFile(file)

	var response = model.MessageResponse{
		Message:   "Upload success.",
		ErrorCode: 0,
	}

	ctx.JSON(http.StatusOK, response)
}

func saveFile(file multipart.File) {

	tempFile, err := ioutil.TempFile("./files/", "file-*.json")

	if err != nil {
		fmt.Println(err)
	}

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Upload Success ....")
	tempFile.Write(fileBytes)
}
