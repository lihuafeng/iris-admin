package email

import (
	"fmt"
	Config "github.com/deatil/doak-cron/config"
	"math/rand"
	"net/smtp"
	"github.com/jordan-wright/email"
	"net/textproto"
	"time"
)

type EmailConf struct {
	Host       string
	Port 	   int64
	User       string
	Password   string
}


func EmailConfInit() *EmailConf {
	instance := new(EmailConf)
	instance.Host = "smtp.yeah.net"
	instance.Port = 465
	instance.User = "lhf2008@yeah.net"
	instance.Password = "NZWMEVSHIFXLRZTO"
	return instance
}

func SendMailTest()  {
	mailConf := EmailConfInit()

	e := &email.Email {
		To: []string{"771831851@qq.com"},
		From: "Jordan Wright <lhf2008@yeah.net>",
		Subject: "Awesome Subject",
		Text: []byte("Text Body is, of course, supported!"),
		HTML: []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	fmt.Println("**1**")
	e.Send(fmt.Sprintf("%s:%d", mailConf.Host, mailConf.Port), smtp.PlainAuth("", mailConf.User, mailConf.Password, mailConf.Host))
	fmt.Println("**2**")
}

func SendMail() (err error) {
	mailConf := EmailConfInit()

	e := email.NewEmail()
	e.From = "Jordan Wright <lhf2008@yeah.net>"
	e.To = []string{"771831851@qq.com"}
	e.Bcc = []string{"771831851@qq.com"}
	e.Cc = []string{"771831851@qq.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")

	fmt.Println(fmt.Sprintf("%s:%d", mailConf.Host, mailConf.Port))
	fmt.Println(mailConf.User)
	fmt.Println(mailConf.Password)
	fmt.Println(mailConf.Host)
	er := e.Send(fmt.Sprintf("%s:%d", mailConf.Host, mailConf.Port), smtp.PlainAuth("", mailConf.User, mailConf.Password, mailConf.Host))
	fmt.Println("##")
	fmt.Print(er)
	if er != nil{
		return er
	}
	return
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Config.EMAIL_CODE_LENGTH; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
