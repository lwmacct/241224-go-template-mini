package app

import "github.com/lwmacct/241224-go-template-pkgs/m_log"

type TsFlag struct {
	Log   m_log.Config
	Start struct{} `group:"start" note:"默认配置"`

	DB struct {
		PGSQL string `group:"db" note:"PostgreSQL URL" default:"postgres://username:password@localhost:5432/dbname?sslmode=disable&TimeZone=Asia/Shanghai"`
		Redis string `group:"db" note:"Redis URL" default:"redis://user:password@localhost:6379/0?protocol=3"` // https://github.com/redis/go-redis/blob/91dddc2e1108c779e8c5b85fd667029873c95172/options.go#L247
	}

	Task struct {
		Bitmask int `group:"task" note:"任务位掩码" default:"0"`
	}
}
