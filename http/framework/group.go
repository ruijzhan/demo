package framework

type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	Group(string) IGroup
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

type Group struct {
	core   *Core
	prefix string

	parent *Group

	middlewares []ControllerHandler
}

func (g *Group) Group(prefix string) IGroup {
	ng := NewGroup(g.core, prefix)
	ng.parent = g
	return ng
}

func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) getAbsPrefix() (s string) {
	for g != nil {
		s = g.prefix + s
		g = g.parent
	}
	return
}

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

func (g *Group) method(m func(string, ...ControllerHandler), uri string, handlers ...ControllerHandler) {
	url := g.getAbsPrefix() + uri
	all := append(g.getMiddlewares(), handlers...)
	m(url, all...)
}

func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	g.method(g.core.Get, uri, handlers...)
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	g.method(g.core.Post, uri, handlers...)
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	g.method(g.core.Put, uri, handlers...)
}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	g.method(g.core.Delete, uri, handlers...)
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}
