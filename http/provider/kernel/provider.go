package kernel

import (
	"github.com/ruijzhan/demo/http/framework"
	"github.com/ruijzhan/demo/http/framework/contract"
	"github.com/ruijzhan/demo/http/framework/gin"
)

type MyKernelProvider struct {
	HttpEngine *gin.Engine
}

func (p *MyKernelProvider) Name() string {
	return contract.KernelKey
}

func (p *MyKernelProvider) Reginster(_ framework.Container) framework.NewInstance {
	return NewMyKernelService
}

func (p *MyKernelProvider) Boot(c framework.Container) error {
	if p.HttpEngine == nil {
		p.HttpEngine = gin.Default()
	}
	p.HttpEngine.SetContainer(c)
	return nil
}

func (p *MyKernelProvider) IsDefer() bool {
	return false
}

func (p *MyKernelProvider) Params(_ framework.Container) []interface{} {
	return []interface{}{p.HttpEngine}
}
