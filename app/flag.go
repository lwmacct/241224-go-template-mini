package app

import "github.com/lwmacct/241224-go-template-pkgs/m_log"

type TsFlag struct {
	Log    m_log.Config
	Start  struct{} `group:"start" note:"默认配置"`
	Client struct {
		Name string `group:"client" note:"客户端名称(ID)" default:""`
	}

	Server struct {
		Listener string `group:"server" note:"监听地址" default:"0.0.0.0:8080"`
	}
}
