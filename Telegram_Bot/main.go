package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		reply := "Please, enter a command or 'help'."

		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "help":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, I'm a GitHub bot, I can show information in a GitHub. \nCommands: \n1) git - link to my repository \n2) tasks - all completed tasks")
			bot.Send(msg)

		case "git":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/DimShim/Andersen_DevOps_course")
			bot.Send(msg)

		case "tasks":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "task1 - TIL (Today I've learned) \ntask2 - Ansible+Flask_app+Nginx \ntask3 - Netstat_script \ntask4 - Telegram bot for GitHub \ntask5 - script_github_pr \ntask6 - docker")
			bot.Send(msg)

		case "task1":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/DimShim/Andersen_DevOps_course/tree/main/1_TIL")
			bot.Send(msg)

		case "task2":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/DimShim/Andersen_DevOps_course/tree/main/Ansible")
			bot.Send(msg)

		case "task3":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/DimShim/Andersen_DevOps_course/tree/main/netstat_script")
			bot.Send(msg)

		case "task4":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/DimShim/Andersen_DevOps_course/tree/main/Telegram_Bot")
			bot.Send(msg)
			
		case "task5":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/DimShim/Andersen_DevOps_course/tree/main/script_github_pr")
			bot.Send(msg)
			
		case "task6":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/DimShim/Andersen_DevOps_course/tree/main/docker")
			bot.Send(msg)

		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)

	}
}
