package starter

/*
 * 启动器管理相关处理逻辑
 */

type registerStarter struct {
	starters []Starter
}

func (r *registerStarter) Register(starter Starter) {
	r.starters = append(r.starters, starter)
}

func (r *registerStarter) GetAllStarters() []Starter {
	return r.starters
}

// RegisterStarter 供包外部使用的注册器变量
var RegisterStarter = new(registerStarter)

// Register 供包外部使用的注册器函数
func Register(starter Starter) {
	RegisterStarter.Register(starter)
}

// SystemRun 运行所有的启动器
func SystemRun() {
	var startContext = make(StarterContext)
	//初始化各个启动器
	for _, starter := range RegisterStarter.GetAllStarters() {
		starter.Init(startContext)
	}
	//建立各个启动器的基础环境
	for _, starter := range RegisterStarter.GetAllStarters() {
		starter.SetUp(startContext)
	}
	//启动各个启动器
	for _, starter := range RegisterStarter.GetAllStarters() {
		starter.Start(startContext)
	}
}
