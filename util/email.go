package email

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/config"
	"github.com/go-gomail/gomail"
)

type Email struct {
	to      string "to"
	cc      string "cc"
	subject string "subject"
	msg     string "msg"
	attach  string "attach"
}

func NewEmail(to, cc, subject, msg, attach string) *Email {
	return &Email{to: to, cc: cc, subject: subject, msg: msg, attach: attach}
}

func SendMimeEmail(email *Email) error {
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		fmt.Println(err)
	}
	HOST := iniconf.String("mailhost")
	USER := iniconf.String("mailuser")
	PASSWORD := iniconf.String("mailpwd")
	SERVER_ADDR := iniconf.String("mailhostaddress")

	d := gomail.NewDialer(SERVER_ADDR, 25, USER, PASSWORD)
	m := gomail.NewMessage()
	m.SetHeader("From", USER)
	sendTo := strings.Split(email.to, ";")
	m.SetHeader("To", sendTo)
	cc := strings.Split(email.to, ";")
	m.SetAddressHeader("Cc", cc)
	m.SetHeader("Subject", email.subject)
	m.SetBody("text/html", email.msg)
	m.Attach(email.attach)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return err
	}
	return nil
}
