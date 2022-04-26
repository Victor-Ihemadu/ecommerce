package routes

import(
	"github.com/akhil/e-commerce-yt/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup")
	incomingRoutes.POST("user/login")
	incomingRoutes.POST("/adim/addproduct")
}