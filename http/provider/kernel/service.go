package kernel

import (
	"net/http"

	"github.com/ruijzhan/demo/http/framework/gin"
)

type MyKernelService struct {
	engine *gin.Engine
}

func NewMyKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &MyKernelService{engine: httpEngine}, nil
}

func (s *MyKernelService) HttpEngine() http.Handler {
	return s.engine
}
