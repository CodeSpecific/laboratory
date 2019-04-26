package starter

/*
 * 启动器相关处理逻辑
 */

// StarterContext 启动器的上下文环境，用于存储各个启动器之间的依赖关系
type StarterContext map[string]interface{}

// Starter 启动器接口
type Starter interface {
	Init(ctx StarterContext)
	SetUp(ctx StarterContext)
	Start(ctx StarterContext)
	StartBlocking() bool
	Stop(ctx StarterContext)
}

// 验证基础启动器是否实现了 Starter 接口，类型断言
var _ Starter = new(BaseStarter)

// BaseStarter 基础启动器提供默认实现，目的是为了有些启动器不需要实现所有的接口方法
type BaseStarter struct {
}

// Init 初始化
func (s *BaseStarter) Init(ctx StarterContext) {}

// SetUp 搭建启动器基础设施
func (s *BaseStarter) SetUp(ctx StarterContext) {}

// Start 启动启动器
func (s *BaseStarter) Start(ctx StarterContext) {}

// StartBlocking 启动器的是否需要阻塞
func (s *BaseStarter) StartBlocking() bool { return false }

// Stop 优雅的停止启动器
func (s *BaseStarter) Stop(ctx StarterContext) {}
