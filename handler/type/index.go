package type

import(
	"retail/handler"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){
	handler.SendResponse(c *gin.Context)
}