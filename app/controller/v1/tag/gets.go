package tag

import (
	"github.com/gin-gonic/gin"
	"leest/app/service/v1/tag"
)

func Gets(c *gin.Context) {
	tag.Gets(c)
}
