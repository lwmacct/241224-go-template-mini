package start

import (
	"github.com/lwmacct/241224-go-template-mini/app"
	"github.com/lwmacct/241224-go-template-pkgs/m_cmd"
	"github.com/lwmacct/241224-go-template-pkgs/m_log"
	"github.com/spf13/cobra"
)

func Cmd() *m_cmd.Ts {
	mc := m_cmd.New(app.Flag).UsePackageName("")
	mc.AddCmd(func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	}, "run", "", "m_log")
	return mc
}

func run(cmd *cobra.Command, args []string) {
	_ = map[string]any{"cmd": cmd, "args": args}
	m_log.Info(m_log.H{"msg": "app.Flag", "data": app.Flag})

}
