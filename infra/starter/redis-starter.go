package starter

import "fmt"

func init() {
	// 注册Redis初始化启动器
	Register(new(RedisStarter))
}

// RedisStarter 配置Redis启动器
type RedisStarter struct {
	BaseStarter
}

// Init 初始化Redis启动器
func (c *RedisStarter) Init(ctx StarterContext) {
	fmt.Println("Redis启动器已经被初始化")
}

// SetUp 搭建Redis启动器基础环境
func (c *RedisStarter) SetUp(ctx StarterContext) {
	fmt.Println("Redis启动器基础环境已经搭建完毕")
}

// Start 启动Redis启动器
func (c *RedisStarter) Start(ctx StarterContext) {
	fmt.Println("Redis启动器正在启动")
}
