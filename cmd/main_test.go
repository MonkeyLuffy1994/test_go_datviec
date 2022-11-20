package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"test_go_datviec/pkg/config"
	"test_go_datviec/pkg/routes"
	"test_go_datviec/pkg/service"
	"testing"
)

func TestMainService(t *testing.T) {

	cof := config.Config{Port: "3000"}

	sv := service.Server{Cof: cof}

	r := gin.Default()

	routes.RouteUpload(r, &sv)

	r.GET("/upload", sv.UploadFile)

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("files", "file.json")

	part.Write([]byte(`sample`))

	writer.Close()

	req, _ := http.NewRequest("POST", "/files/upload", body)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
