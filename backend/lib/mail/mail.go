package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	SendMail()
}

func SendMail() {
	addr := "smtp.gmail.com:587"

	auth := smtp.PlainAuth("", "n3o.bar@gmail.com", "suofnjcqapctltfp", "smtp.gmail.com")

	from := "n3o.bar@gmail.com"

	to := []string{"davidsgv98@gmail.com"}
	concatenate := strings.Join(to, ", ")
	toMsg := fmt.Sprintf("To: %v\r\n", concatenate)

	subject := "Test Subject"
	subjectMsg := fmt.Sprintf("Subject: %v\r\n", subject)

	body := "Example body."
	bodyMsg := fmt.Sprintf("%v\r\n", body)

	msg := []byte(toMsg + subjectMsg + "\r\n" + bodyMsg)

	resultado := sendEmail(addr, auth, from, to, msg)
	fmt.Println(resultado)
}

func sendEmail(addr string, auth smtp.Auth, from string, to []string, message []byte) string {
	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		fmt.Printf("Error!: %s", err)
		return "Error! sending mail"
	}
	return "email sent successfully"
}
