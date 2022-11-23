package framework

type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	Group(string) IGroup
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

type Group struct {
	core   *Core
	prefix string
}

func (g *Group) Group(prefix string) IGroup {
	return NewGroup(g.core, g.prefix+prefix)
}

func (g *Group) Get(uri string, h ControllerHandler) {
	url := g.prefix + uri
	g.core.Get(url, h)
}

func (g *Group) Post(uri string, h ControllerHandler) {
	url := g.prefix + uri
	g.core.Post(url, h)
}

func (g *Group) Put(uri string, h ControllerHandler) {
	url := g.prefix + uri
	g.core.Put(url, h)
}

func (g *Group) Delete(uri string, h ControllerHandler) {
	url := g.prefix + uri
	g.core.Delete(url, h)
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}
