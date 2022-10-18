package main

import (
    "github.com/deatil/doak-cron/pkg/db"
    "os"
    "fmt"
    "log"
    "time"

    "github.com/urfave/cli/v2"

    "github.com/deatil/doak-cron/pkg/cron"
    "github.com/deatil/doak-cron/pkg/parse"
    "github.com/deatil/doak-cron/pkg/table"
    "github.com/deatil/doak-cron/pkg/logger"

    "github.com/kataras/iris/v12"
    "github.com/google/uuid"
)

// 版本号
var version = "1.0.6"

/**
 * go版本的通用计划任务
 *
 * > go run main.go cron --conf="./cron.json" --debug
 * > go run main.go cron --conf="./cron.json" --log="./cron.log" --debug
 * > go run main.go cron ver
 *
 * > main.exe cron --conf="./cron.json" --debug
 * > main.exe cron --conf="./cron.json" --log="./cron.log" --debug
 * > main.exe cron ver
 *
 * @create 2022-6-29
 * @author deatil
 */
func main() {
    //Lris
    httpapp := iris.Default()
    httpapp.Logger().SetLevel("error")
    httpapp.OnErrorCode(iris.StatusInternalServerError, internalServerError)
    httpapp.Use(myMiddleware)
    // 加载视图模板地址
    httpapp.RegisterView(iris.HTML("./views", ".html"))
    //提供静态文件服务
    httpapp.HandleDir("/static", "./static")
    httpapp.Get("/", func(ctx iris.Context) {
        var crons []db.CronModel
        //db.Db.Where("status=1").Find(&crons)
        db.Db.Find(&crons)
        ctx.ViewData("crons", crons)
        // Render template file: ./views/index.html
        ctx.View("index.html")
    })
    httpapp.Get("/add", func(ctx iris.Context) {
        ctx.View("add.html")
    })
    httpapp.Post("/save", func(ctx iris.Context) {
        CronTime := ctx.FormValue("CronTime")
        command := ctx.FormValue("command")
        cron_type, _ := ctx.PostValueInt("type")
        fmt.Printf("CronTime:%s, cron_type:%d, command:%s", CronTime, cron_type, command)
        db.Db.Create(&db.CronModel{
            UniueCode:uuid.NewString(),
            CronType:cron_type,
            CronTime:CronTime,
            Command:command,
            RunStatus:0,
            Status:1,
            CreatedAt:time.Now().Format("2006-01-02 15:04:05"),
        })
        ctx.Redirect("/")
    })
    httpapp.Post("/modify", func(ctx iris.Context) {
        uniue_code := ctx.PostValue("uniue_code")
        status,err := ctx.PostValueInt("status")
        if err!=nil{
            ctx.JSON(iris.Map{"code":1,"msg":"缺少参数"})
        }
        var cron db.CronModel
        db.Db.First(&cron, "uniue_code=?",uniue_code)
        db.Db.Model(&cron).Update("status", status)
        ctx.JSON(iris.Map{"code":0,"msg":"修改成功"})
    })

    // Listens and serves incoming http requests
    // on http://localhost:8080.
    httpapp.Run(iris.Addr(":8080"))

    app := cli.NewApp()
    app.EnableBashCompletion = true
    app.Commands = []*cli.Command{
        {
            Name:    "cron",
            Aliases: []string{"c"},
            Usage:   "go版本的通用计划任务",
            Flags: []cli.Flag{
                &cli.BoolFlag{Name: "debug", Aliases: []string{"d"}},
                &cli.StringFlag{Name: "conf", Aliases: []string{"c"}},
                &cli.StringFlag{Name: "log", Aliases: []string{"l"}},
            },
            Action: func(ctx *cli.Context) error {
                // 设置日志存储文件
                log := ctx.String("log")
                if log != "" {
                    logger.WithLogFile(log)
                }

                conf := ctx.String("conf")
                debug := ctx.Bool("debug")

                crons, settings := parse.MakeCron(conf, debug)
                if crons == nil {
                    fmt.Println("配置文件错误")
                    return nil
                }

                // 执行计划任务
                cron.AddCron(func(croner *cron.Cron) {
                    if len(crons) > 0 {
                        for k, c := range crons {
                            cronId, err := croner.AddFunc(c.Spec, c.Cmd)
                            if err != nil{
                                 logger.Log().Error().Msg("[cron]" + err.Error())
                            }

                            settings[k]["cron_id"] = cronId
                        }
                    }

                    fmt.Println("")

                    // 显示详情
                    title := "Doak Cron v" + version
                    table.ShowTable(title, settings)
                })

                return nil
            },
            Subcommands: []*cli.Command{
                {
                    Name:  "ver",
                    Usage: "显示计划任务版本号",
                    Action: func(ctx *cli.Context) error {
                        fmt.Println("计划任务当前版本号为: ", version)

                        return nil
                    },
                },
                {
                    Name:  "version",
                    Usage: "显示计划任务版本号",
                    Action: func(ctx *cli.Context) error {
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

func myMiddleware(ctx iris.Context) {
    ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
    ctx.Next()
}

func internalServerError(ctx iris.Context) {
    ctx.WriteString("Oups something went wrong, try again")
}
