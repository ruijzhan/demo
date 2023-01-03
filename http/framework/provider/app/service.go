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
