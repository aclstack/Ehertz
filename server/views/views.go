package views

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gorm.io/gorm"
)

type AllMethod struct {
	conn *gorm.DB
}

func (all AllMethod) All(ctx *app.RequestContext, exec, check func()) (resp map[string]interface{}) {
	method := string(ctx.Request.Method())
	switch method {
	case "GET":
		return utils.H{"ping": "GET"}
	case "POST":
		return utils.H{"ping": "POST"}
	case "DELETE":
		return utils.H{"ping": "DELETE"}
	case "PUT":
		return utils.H{"ping": "PUT"}
	case "OPTIONS":
		return utils.H{"ping": "OPTIONS"}
	default:
		return utils.H{"Message": "fuck"}
	}
	
}
