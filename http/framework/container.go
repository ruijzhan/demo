package framework

import (
	"errors"
	"fmt"
	"sync"
)

func NewMyContainer() Container {
	return &MyContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]any{},
	}
}

type Container interface {

	// Bind 绑定一个服务提供者
	Bind(provider ServiceProvider) error

	// IsBind 关键字是否已经绑定服务提供者
	IsBind(key string) bool

	Make(key string) (any, error)

	MustMake(key string) any

	MakeNew(key string, params []any) (any, error)
}

type MyContainer struct {
	Container
	providers map[string]ServiceProvider
	instances map[string]any
	lock      sync.RWMutex
}

// Bind 绑定一个服务提供者
func (c *MyContainer) Bind(provider ServiceProvider) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	key := provider.Name()
	c.providers[key] = provider

	if !provider.IsDefer() {
		if err := provider.Boot(c); err != nil {
			return err
		}

		params := provider.Params(c)
		method := provider.Reginster(c)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		c.instances[key] = instance
	}

	return nil
}

// IsBind 关键字是否已经绑定服务提供者
func (c *MyContainer) IsBind(key string) bool {
	return c.findServiceProvider(key) != nil
}

func (c *MyContainer) Make(key string) (any, error) {
	return c.make(key, nil, false)
}

func (c *MyContainer) MustMake(key string) any {
	if ins, err := c.make(key, nil, false); err != nil {
		panic(err)
	} else {
		return ins
	}
}

func (c *MyContainer) MakeNew(key string, params []any) (any, error) {
	return c.make(key, params, true)
}

func (c *MyContainer) make(key string, params []any, forceNew bool) (any, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	sp := c.findServiceProvider(key)
	if sp == nil {
		return nil, fmt.Errorf("contract %s was not registered", key)
	}

	if forceNew {
		return c.newInstance(sp, params)
	}

	if ins, ok := c.instances[key]; ok {
		return ins, nil
	}

	ins, err := c.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	c.instances[key] = ins

	return ins, nil
}

func (c *MyContainer) findServiceProvider(key string) ServiceProvider {
	c.lock.Lock()
	defer c.lock.Unlock()
	if sp, ok := c.providers[key]; ok {
		return sp
	}
	return nil
}

func (c *MyContainer) newInstance(sp ServiceProvider, params []any) (any, error) {
	if err := sp.Boot(c); err != nil {
		return nil, err
	}

	if params == nil {
		params = sp.Params(c)
	}

	method := sp.Reginster(c)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, nil
}
