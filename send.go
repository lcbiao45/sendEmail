package sendMail

import (
	"net/smtp"
	"strings"
)

func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")                    // 分割字符串
	auth := smtp.PlainAuth("", user, password, hp[0]) //返回一个auth的interface的实例。这个interface内容是nil
	//实际上就是一个实现了next和start的一个struct。
	var content_type string
	if mailtype == "html" { //mailtype是邮件的类型。
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8" //如果是HTML那么头文件就是
		// Content-Type: text/html; charset=UTF-8
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8" //
		//Content-Type:text/plain; charset=UTF-8
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";") //使用；作为to中的分隔符。当然返回的是一个slice
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
