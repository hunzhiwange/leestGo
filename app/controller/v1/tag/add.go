package tag

import (
	"leest/app/service/v1/tag"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	tag.Add(c)
}
