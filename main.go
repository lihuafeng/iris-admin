package main

import (
    "os"
    "fmt"
    "log"

    "github.com/urfave/cli/v2"

    "github.com/deatil/doak-cron/pkg/cron"
    "github.com/deatil/doak-cron/pkg/parse"
)

// 版本号
var version = "1.0.1"

// go run main.go cron --conf="./cron.json" --debug
// go run main.go cron ver
func main() {
    app := cli.NewApp()
    app.EnableBashCompletion = true
    app.Commands = []*cli.Command{
        {
            Name:    "cron",
            Aliases: []string{"c"},
            Usage:   "cron",
            Flags: []cli.Flag{
                &cli.BoolFlag{Name: "debug", Aliases: []string{"d"}},
                &cli.StringFlag{Name: "conf", Aliases: []string{"c"}},
            },
            Action: func(c *cli.Context) error {
                conf := c.String("conf")
                debug := c.Bool("debug")

                crons := parse.MakeCron(conf, debug)
                if crons == nil {
                    fmt.Println("配置文件错误")
                    return nil
                }

                newCrons := make([]cron.Option, 0)
                for _, v := range crons {
                    for kk, vv := range v {
                        newCrons = append(newCrons, cron.Option{
                            Spec: kk,
                            Cmd:  vv,
                        })
                    }
                }

                fmt.Println("任务开始执行...")

                cron.AddCrons(newCrons)

                return nil
            },
            Subcommands: []*cli.Command{
                {
                    Name:  "ver",
                    Usage: "显示计划任务版本号",
                    Action: func(c *cli.Context) error {
                        fmt.Println("计划任务当前版本号为: ", version)

                        return nil
                    },
                },
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}