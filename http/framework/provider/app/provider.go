package app

import (
	"errors"

	"github.com/ruijzhan/demo/http/framework"
	"github.com/ruijzhan/demo/http/framework/contract"
)

type MyAppProvider struct {
	BaseFolder string
}

func (h *MyAppProvider) Params(container framework.Container) []any {
	return []any{container, h.BaseFolder}
}

func NewMyApp(params ...any) (any, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	return &MyApp{container: container, baseFolder: baseFolder}, nil
}

func (h *MyAppProvider) Name() string {
	return contract.AppKey
}

func (h *MyAppProvider) Reginster(_ framework.Container) framework.NewInstance {
	return NewMyApp
}

func (h *MyAppProvider) Boot(_ framework.Container) error {
	return nil
}

func (h *MyAppProvider) IsDefer() bool {
	return false
}
