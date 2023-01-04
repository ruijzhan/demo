package app

import (
	"flag"
	"path/filepath"

	"github.com/ruijzhan/demo/http/framework"
	"github.com/ruijzhan/demo/http/framework/util"
)

type MyApp struct {
	container  framework.Container
	baseFolder string
}

func (m MyApp) BaseFolder() string {
	if m.baseFolder != "" {
		return m.baseFolder
	}

	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder 参数，默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	return util.GetExecDir()
}

func (m MyApp) StorageFolder() string {
	return filepath.Join(m.baseFolder, "storage")
}

func (m MyApp) LogFolder() string {
	return filepath.Join(m.baseFolder, "log")
}

// Version 定义当前版本
func (m MyApp) Version() string {
	panic("not implemented") // TODO: Implement
}

// ConfigFolder 定义了配置文件的路径
func (m MyApp) ConfigFolder() string {
	panic("not implemented") // TODO: Implement
}

// ProviderFolder 定义业务自己的服务提供者地址
func (m MyApp) ProviderFolder() string {
	panic("not implemented") // TODO: Implement
}

// MiddlewareFolder 定义业务自己定义的中间件
func (m MyApp) MiddlewareFolder() string {
	panic("not implemented") // TODO: Implement
}

// CommandFolder 定义业务定义的命令
func (m MyApp) CommandFolder() string {
	panic("not implemented") // TODO: Implement
}

// RuntimeFolder 定义业务的运行中间态信息
func (m MyApp) RuntimeFolder() string {
	panic("not implemented") // TODO: Implement
}

// TestFolder 存放测试所需要的信息
func (m MyApp) TestFolder() string {
	panic("not implemented") // TODO: Implement
}
