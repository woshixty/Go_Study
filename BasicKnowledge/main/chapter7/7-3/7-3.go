package main

/**
7-3 理解接口与类型的关系
*/

/**
7.3.1 一个类型可以实现多个接口
*/
type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

/**
7.3.2 多个类型可以实现相同的接口
一个接口的方法不一定需要一个类型完全实现，可以通过类型中嵌入的其他类型或者结构体来实现
*/
// 一个服务器需要满足能够开启和写日志的功能
type Service interface {
	Start()
	Log(string)
}

// 日志器
type Logger struct {
}

// 实现Service的Log()方法
func (g *Logger) Log(l string) {
}

// 游戏服务
type GameService struct {
	Logger
}

// 实现Service的Start方法
func (g *GameService) Start() {

}

func main() {

}
