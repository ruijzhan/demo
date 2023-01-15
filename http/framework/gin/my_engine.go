package gin

import "github.com/ruijzhan/demo/http/framework"

func (e *Engine) SetContainer(c framework.Container) {
	e.container = c
}
