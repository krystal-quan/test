package main

import (
	"fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "1340551640:AAHtwcA4PsGd_Tt0bULVv25g6s6QRunwH6Y",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		menuLan = &tb.ReplyMarkup{ResizeReplyKeyboard: true,
			OneTimeKeyboard: true}

		btnVi  = menuLan.Text("Tiếng Việt")
		btnEng = menuLan.Text("English")
	)

	menuLan.Reply(
		menuLan.Row(btnVi),
		menuLan.Row(btnEng),
	)

	var listMem ListMem
	var content Content
	b.Handle("/setup", func(m *tb.Message) {
		content.setup()
		listAd, err := b.AdminsOf(m.Chat)
		if err != nil {
			log.Println(err)
		}

		var founderId int
		for _, v := range listAd {
			if v.Role == "creator" {
				founderId = v.User.ID
			}
		}

		isOwner := false
		if m.Sender.ID == founderId {
			isOwner = true
		}

		if isOwner {
			listMem = newListMem()
			listMem.AddNew(m.Sender.ID, Lead, English)
		} else {
			b.Send(m.Sender, "You don't have permission to do this action.")
		}
	})

	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		listMem.AddMem(m.UserJoined.ID)
		b.Send(m.UserJoined, "Xin chào, mời chọn ngôn ngữ hiển thị\nHello, Please choose your language", menuLan)
	})

	b.Handle(tb.OnUserLeft, func(m *tb.Message) {
		listMem.delMem(m.UserLeft.ID)
	})

	b.Handle(&btnVi, func(m *tb.Message) {
		b.Send(m.Sender, content.loadContent(Vietnamese, "changeLan", "Result"))

		check, user := listMem.isInList(m.Sender.ID)
		if check {
			user.lan = Vietnamese
		} else {
			listMem.AddMem(m.Sender.ID)
		}
	})

	b.Handle(&btnEng, func(m *tb.Message) {
		b.Send(m.Sender, content.loadContent(English, "changeLan", "Result"))

		check, user := listMem.isInList(m.Sender.ID)
		if check {
			user.lan = English
		} else {
			listMem.AddMem(m.Sender.ID)
		}
	})

	b.Handle("/changeLan", func(m *tb.Message) {
		b.Send(m.Sender, content.loadContent(Vietnamese, "changeLan", "string"))
		b.Send(m.Sender, content.loadContent(English, "changeLan", "string"), menuLan)
	})

	b.Handle("/settime", func(m *tb.Message) {
		mes, err := b.Send(m.Sender, "Start the timer")
		if err != nil {
			log.Println(err)
		}

		for i := 0; i < 100; i++ {
			time.Sleep(1 * time.Second)
			hello := fmt.Sprintf("Time left: %d", i)
			b.Edit(mes, hello)
		}
	})

	b.Handle("/listMem", func(m *tb.Message) {
		// listAd, err := b.AdminsOf(m.Chat)
		// if err != nil {
		// 	log.Println(err)
		// }

		// isAdmin := false
		// for _, v := range listAd {
		// 	if m.Sender.ID == v.User.ID {
		// 		isAdmin = true
		// 	}
		// }

		// if listMem.isAdmin(&m.Sender.ID) {
		// 	for _, mem := range listMem {
		// 		content := fmt.Sprintf("id: %d\nrole: %s\n", mem.id, mem.role.toString())
		// 		b.Send(m.Sender, content)
		// 	}
		// } else {
		// 	b.Send(m.Sender, "You do not have permission to do this action")
		// }
		b.Send(m.Sender, listMem.String())
	})

	b.Handle("/comWithArgs", func(m *tb.Message) {
		fmt.Println(m.Payload)
	})

	b.Start()

	// fmt.Print("Enter text: ")
	// var input string
	// fmt.Scanln(&input)
	// if input == "exit" {

	// 	listMemJson, err := json.Marshal(listMem)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}

	// 	err = ioutil.WriteFile("Users.json", listMemJson, 0644)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}

	// 	b.Stop()
	// }

}
