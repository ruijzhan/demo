package cobra

import "github.com/ruijzhan/demo/http/framework"

func (c *Command) SetContainer(container framework.Container) {

}

func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}
