package routes

import (
	"github.com/gin-gonic/gin"
	"test_go_datviec/pkg/service"
)

func RouteUpload(r *gin.Engine, sv *service.Server) {
	routes := r.Group("/files")
	routes.POST("/upload", sv.UploadFile)
}
