package starter

import "fmt"

// init 依赖于文档名称的词法顺序，按照开头的a-z的顺序加载
func init() {
	// 注册配置文件初始化启动器
	Register(new(ConfigStarter))
}

// ConfigStarter 配置文件启动器
type ConfigStarter struct {
	BaseStarter
}

// Init 初始化配置文件启动器
func (c *ConfigStarter) Init(ctx StarterContext) {
	fmt.Println("配置文件启动器已经被初始化")
}
