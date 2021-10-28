package sendmail

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)


func Send(to []string, host, from, subject, emailbody string) (emailerror string) {
	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")
	fmt.Println("r", r)
	c, err := smtp.Dial(host)
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	defer c.Close()
	if err = c.Mail(r.Replace(from)); err != nil {
		fmt.Println(err)
		return "err"
	}
	for i := range to {
		to[i] = r.Replace(to[i])
		if err = c.Rcpt(to[i]); err != nil {
			fmt.Println(err)
			return "err"
		}
	}

	w, err := c.Data()
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	msg := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(emailbody))

	_, err = w.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
	}
	err = w.Close()
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	c.Quit()
	return "ok"
}
