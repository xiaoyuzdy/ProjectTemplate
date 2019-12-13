package main

import (
	"github.com/urfave/cli"
	"go-web/cmd"
	"go-web/component"
	"os"
	"runtime"
	"sort"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	component.InitComponent()
	// 配置cli参数
	app := cli.NewApp()
	app.Name = "go-web"
	app.Usage = "go-web"
	app.Version = "1.0.0"
	app.Action = allIn
	app.Commands = []cli.Command{
		{
			Name:  "app",
			Usage: "run",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开启运行api服务",
					Action: cmd.Api,
				},
			},
		}, {
			Name:  "createTable",
			Usage: "初始化数据库表",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开启初始化数据库表程序",
					Action: cmd.CreateTable,
				},
			},
		},
		{
			Name:    "all",
			Aliases: []string{"all"},
			Usage:   "全部服务",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "所有服务同时运行",
					Action: allIn,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}

func allIn(c *cli.Context) {
	go cmd.CreateTable()
	go cmd.Grpc(c)
	cmd.Api(c)
}
