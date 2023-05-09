package handlers

import (
	"bytes"
	"fmt"
	"fontstudios/config"
	"fontstudios/utils"
	"net/http"
	"net/mail"
	"net/smtp"
	"strings"
	"text/template"
)

func render(w http.ResponseWriter, data interface{}) {
	utils.MainRenderTemplate(w, "order", data)
}

func sendEmail(name string, email string, workType string, workDescription string, taskLink string, budget string, deadlines string) {
	cfg := config.LoadConfig()

	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	// Инициализируем данные, которые будут использоваться в шаблоне
	data := struct {
		Name        string
		Email       string
		Description string
		Type        string
		Link        string
		Budget      string
		Deadlines   string
	}{
		Name:        name,
		Email:       email,
		Description: workDescription,
		Type:        workType,
		Link:        taskLink,
		Budget:      budget,
		Deadlines:   deadlines,
	}

	// Инициализируем шаблон
	tpl, err := template.ParseFiles("templates/letter.html")
	if err != nil {
		panic(err)
	}

	// Создаем буфер для результата применения шаблона
	buf := new(bytes.Buffer)

	// Применяем шаблон к данным, записываем результат в буфер
	if err := tpl.Execute(buf, data); err != nil {
		panic(err)
	}

	// Получаем содержимое письма в виде строки
	body := buf.String()

	// Создайте сообщение
	message := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Новый заказ в FrontStudios\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=\"utf-8\"\r\n"+
		"\r\n"+
		"%s", cfg.UserEmail, body))

	// Подключитесь к SMTP-серверу
	auth := smtp.PlainAuth("", cfg.SmtpUsername, cfg.SmtpPassword, smtpServer)
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, cfg.SmtpUsername, []string{cfg.UserEmail}, message)
	if err != nil {
		panic(err)
	}
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		render(w, nil)
	} else if r.Method == http.MethodPost {

		taskLink := r.FormValue("task_link")
		if !strings.HasPrefix(taskLink, "https://") {
			render(w, "Ссылка на ТЗ указана неправильно")
			return
		}

		email := r.FormValue("email")
		_, err := mail.ParseAddress(email)
		if err != nil {
			render(w, "Email указан неправильно")
			return
		}

		name := r.FormValue("name")
		workType := r.FormValue("work_type")
		workDescription := r.FormValue("work_description")
		budget := r.FormValue("budget")
		deadlines := r.FormValue("deadlines")

		sendEmail(name, email, workType, workDescription, taskLink, budget, deadlines)

		render(w, "Заказ отправлен успешно")
	}
}
