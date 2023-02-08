package core

import (
	"fmt"
	Config "github.com/deatil/doak-cron/config"
	Controller "github.com/deatil/doak-cron/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gobwas"
	"log"
	"net/http"
	"strconv"
)

const (
	addr      = "localhost:8080"
	endpoint  = "/echo"
	namespace = "default"
	// false if client sends a join request.
	serverJoinRoom = false
	// if the above is true then this field should be filled, it's the room name that server force-joins a namespace connection.
	serverRoomName = "room1"
)

// userMessage implements the `MessageBodyUnmarshaler` and `MessageBodyMarshaler`.
type userMessage struct {
	From string `json:"from"`
	Text string `json:"text"`
}

var serverAndClientEvents = neffos.Namespaces{
	namespace: neffos.Events{
		neffos.OnNamespaceConnected: func(c *neffos.NSConn, msg neffos.Message) error {
			log.Printf("[%s] connected to namespace [%s].", c, msg.Namespace)

			if !c.Conn.IsClient() && serverJoinRoom {
				c.JoinRoom(nil, serverRoomName)
			}

			return nil
		},
		neffos.OnNamespaceDisconnect: func(c *neffos.NSConn, msg neffos.Message) error {
			log.Printf("[%s] disconnected from namespace [%s].", c, msg.Namespace)
			return nil
		},
		neffos.OnRoomJoined: func(c *neffos.NSConn, msg neffos.Message) error {
			text := fmt.Sprintf("[%s] joined to room [%s].", c, msg.Room)
			log.Println(text)

			// notify others.
			if !c.Conn.IsClient() {
				c.Conn.Server().Broadcast(c, neffos.Message{
					Namespace: msg.Namespace,
					Room:      msg.Room,
					Event:     "notify",
					Body:      []byte(text),
				})
			}

			return nil
		},
		neffos.OnRoomLeft: func(c *neffos.NSConn, msg neffos.Message) error {
			text := fmt.Sprintf("[%s] left from room [%s].", c, msg.Room)
			log.Println(text)

			// notify others.
			if !c.Conn.IsClient() {
				c.Conn.Server().Broadcast(c, neffos.Message{
					Namespace: msg.Namespace,
					Room:      msg.Room,
					Event:     "notify",
					Body:      []byte(text),
				})
			}

			return nil
		},
		"chat": func(c *neffos.NSConn, msg neffos.Message) error {
			if !c.Conn.IsClient() {
				c.Conn.Server().Broadcast(c, msg)
			} else {
				var userMsg userMessage
				err := msg.Unmarshal(&userMsg)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s >> [%s] says: %s\n", msg.Room, userMsg.From, userMsg.Text)
			}
			return nil
		},
		// client-side only event to catch any server messages comes from the custom "notify" event.
		"notify": func(c *neffos.NSConn, msg neffos.Message) error {
			if !c.Conn.IsClient() {
				return nil
			}

			fmt.Println(string(msg.Body))
			return nil
		},
	},
}


func Run()  {
	//Lris
	httpapp := iris.Default()
	httpapp.Logger().SetLevel("error")
	httpapp.Use(myMiddleware)
	// 加载视图模板地址 Reload(true)重新加载本地文件更改,线上环境可以去掉，避免重复启动
	httpapp.RegisterView(iris.HTML("./views", ".html").Reload(true))
	//提供静态文件服务
	httpapp.HandleDir("/static", "./static")
	httpapp.HandleDir("/uploads", "./uploads")

	//加载路由
	Controller.RouterHandler(httpapp)

	//前置操作
	httpapp.Use(before)
	//后置操作
	//httpapp.Use(after)

	//websocket
	ws := neffos.New(gobwas.DefaultUpgrader, serverAndClientEvents)
	ws.IDGenerator = func(w http.ResponseWriter, r *http.Request) string {
		if userID := r.Header.Get("X-Username"); userID != "" {
			return userID
		}

		return neffos.DefaultIDGenerator(w, r)
	}
	ws.OnUpgradeError = func(err error) {
		log.Printf("ERROR: %v", err)
	}
	ws.OnConnect = func(c *neffos.Conn) error {
		log.Printf("[%s] connected to the server.", c)
		return nil
	}
	ws.OnDisconnect = func(c *neffos.Conn) {
		log.Printf("[%s] disconnected from the server.", c)
	}

	httpapp.Get("/msg", websocket.Handler(ws))

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	httpapp.Run(iris.Addr(":"+strconv.Itoa(Config.SERVER_PORT)))
}

func before(ctx iris.Context)  {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // execute the next handler, in this case the main one.
}

func after(ctx iris.Context)  {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
