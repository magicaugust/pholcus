// +build windows
package exec

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/henrylee2cn/pholcus/app/scheduler"
	"github.com/henrylee2cn/pholcus/config"
	"github.com/henrylee2cn/pholcus/runtime/cache"

	"github.com/henrylee2cn/pholcus/cmd" // cmd版
	"github.com/henrylee2cn/pholcus/gui" // gui版
	"github.com/henrylee2cn/pholcus/web" // web版
)

func run(which string) {
	exec.Command("cmd.exe", "/c", "title", config.APP_FULL_NAME).Start()
	defer func() {
		if cache.Task.InheritDeduplication {
			scheduler.SaveDeduplication()
		}
	}()

	// 选择运行界面
	switch which {
	case "gui":
		gui.Run()

	case "cmd":
		cmd.Run()

	case "web":
		fallthrough
	default:
		ctrl := make(chan os.Signal, 1)
		signal.Notify(ctrl, os.Interrupt, os.Kill)
		go web.Run()
		<-ctrl
	}
}
