package main
import (
	tgbotapi "Projects/awesomeProject/main/telegram-bot-api-master"
  "log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1901931911:AAFINPGIJ4psmAvUJFrpTm4dqaoj-zcog9Q")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	array_of_commands:=[...]string{
		"/hello",
		"/freedom",
		"/statusBender",
		"/statusBenderOff",
	}
	var array_of_users_name  []string
	array_of_users_name = append(array_of_users_name,"defolt")
	status_write:=false
	for update := range updates {

		//name := update.Message.Chat
		//name_st :=name.UserName
		//chek:=false
		//for new_name := range array_of_users_name{
		//	if array_of_users_name[new_name]==name_st{
		//		chek=true
		//		break
		//	}
		//}
		//if chek==false{
		//
		//	array_of_users_name:=append(array_of_users_name,name_st)
		//	log.Println(array_of_users_name)
		//}

		//if update.Message == nil { // ignore any non-Message Updates
		//	continue
		//}
		//
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		//
		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID
		//
		//bot.Send(msg)

		if len(update.Message.Text)>0{
			ChatID := update.Message.Chat.ID
			st:=update.Message.Text
			if st[0]=='/'{
				if st == array_of_commands[0] {
					Text := "Hello world"
					msg := tgbotapi.NewMessage(ChatID, Text)
					bot.Send(msg)
				}
				if st == array_of_commands[1] {
					Text := "Freedom for Bender!!!"
					msg := tgbotapi.NewMessage(ChatID, Text)
					bot.Send(msg)
				}
				if st == array_of_commands[2]{
					bot.Send(Message_replay(update,array_of_commands[2],status_write))
					status_write=true
				}
				if st == array_of_commands[3]{
					bot.Send(Message_replay(update,array_of_commands[3],status_write))
					status_write=false
				}
			}else{
				bot.Send(Message_replay(update,"no",true))
			}
		}
	}
}
func Message_replay(update tgbotapi.Update,status string,status_write bool) tgbotapi.MessageConfig{

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "no")
	if update.Message != nil && status=="no"{
		// Пользователь, который написал боту
		UserName := update.Message.From.UserName

		// ID чата/диалога.
		// Может быть идентификатором как чата с пользователем
		// (тогда он равен UserID) так и публичного чата/канала
		ChatID := update.Message.Chat.ID

		// Текст сообщения
		Text := update.Message.Text

		log.Printf("[%s] %d %s", UserName, ChatID, Text)

		// Ответим пользователю его же сообщением
		reply := Text
		// Созадаем сообщение
		msg := tgbotapi.NewMessage(ChatID, reply)
		// и отправляем его
		return msg
	}
	if update.Message != nil && status=="/statusBender" && status_write==false{

		ChatID := update.Message.Chat.ID
		status_write=true
		reply := "O no you activate Bender" +
			"\nAnd Bender went to get a beer!!!"
		msg := tgbotapi.NewMessage(ChatID, reply)
		return msg
	}
	if update.Message != nil && status=="/statusBenderOff" && status_write==true{

		ChatID := update.Message.Chat.ID
		status_write=true
		reply := "Bender is leaving because Bender deserves more :( \n *angry bot"
		msg := tgbotapi.NewMessage(ChatID, reply)
		return msg
	}
	return msg
}