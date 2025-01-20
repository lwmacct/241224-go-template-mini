package main

import (
	"fmt"
	"os"

	"github.com/lwmacct/241224-go-template-mini/app"
	"github.com/lwmacct/241224-go-template-mini/app/start"
	"github.com/lwmacct/241224-go-template-mini/app/test"
	"github.com/lwmacct/241224-go-template-mini/app/version"
	"github.com/lwmacct/241224-go-template-pkgs/pkgs/m_cmd"
	"github.com/lwmacct/241224-go-template-pkgs/pkgs/m_log"
)

var cmd *m_cmd.Ts

func main() {
	cmd = m_cmd.New(nil)

	{
		// 命令行参数
		cmd.AddCobra(version.Cmd().Cobra())
		cmd.AddCobra(start.Cmd().Cobra())

		// 开发环境中的测试命令
		if os.Getenv("ACF_SHOW_TEST") == "1" {
			cmd.AddCobra(test.Cmd().Cobra())
		}
	}

	{
		// 日志处理
		ml := m_log.NewConfig()
		ml.Level = app.Flag.Log.Level
		if app.Flag.Log.File == "" {
			app.Flag.Log.File = ml.File
		} else {
			ml.File = app.Flag.Log.File
		}
		if version.Workspace != "" {
			ml.CallerClip = version.Workspace
		}
		app.Log = m_log.NewTs(ml)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
