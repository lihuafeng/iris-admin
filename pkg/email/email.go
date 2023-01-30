package email

import (
	"fmt"
	Config "github.com/deatil/doak-cron/config"
	"io/ioutil"
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
	instance.Port = 25
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

func SendMail(username ,email_account string, title string, blade_type string) (err error) {
	mailConf := EmailConfInit()

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", Config.SERVER_NAME, mailConf.User)
	e.To = []string{email_account}
	//e.Bcc = []string{"771831851@qq.com"}
	//e.Cc = []string{"771831851@qq.com"}
	e.Subject = title
	//e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = emailBlade(username, blade_type)
	er := e.Send(fmt.Sprintf("%s:%d", mailConf.Host, mailConf.Port), smtp.PlainAuth("", mailConf.User, mailConf.Password, mailConf.Host))
	if er != nil{
		return er
	}
	return
}

//读取模板内容 使用go语言中的内置包，buffio和ioutil来读取
func emailBlade(username string, blade_type string) []byte  {
	txt,err := ioutil.ReadFile(fmt.Sprintf("./views/emailblade/%s.html", blade_type)) //读取文件
	if blade_type =="code"{
		code := RandCode(6)
		txt_str := fmt.Sprintf(string(txt), username, code)
		txt = []byte(txt_str)
	}
	if err != nil{
		return txt
	}
	return txt
}
//获取随机数字
func RandCode(length int) string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
