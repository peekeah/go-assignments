package main

import "fmt"

func RunChatSimulator() {
	existingUsers := []string{
		"Alice", "Bob", "Charlie", "Diana", "Ethan",
		"Fiona", "George", "Hannah", "Ian", "Julia",
	}

	chats := InitializeMsgStore()
	users := InitializeUserStore()

	// 1. Add users
	for _, val := range existingUsers {
		users.AddUser(val)
	}

	// 2. Get all users
	allUsers := users.GetUsers()
	fmt.Println("Users:", allUsers)

	// Pick 2 users for chat
	sender, _ := users.GetUser(allUsers[0].Id)
	receiver, _ := users.GetUser(allUsers[1].Id)

	// 3. Send messages
	chats.SendMessage(sender, receiver, "Hello there!")
	chats.SendMessage(sender, receiver, "How are you?")
	chats.SendMessage(receiver, sender, "I’m good, thanks!")

	// 4. Get messages
	fmt.Println("Messages after sending:", chats.GetMessages())

	// 5. Test DeleteMessage
	msgs := chats.GetMessages()
	if len(msgs) > 0 {
		firstMsgId := msgs[0].Id
		err := chats.DeleteMessage(firstMsgId)
		if err != nil {
			fmt.Println("Delete failed:", err)
		} else {
			fmt.Println("Message deleted, new chat history:", chats.GetMessages())
		}
	}

	// 6. Test GetUser (failure case)
	_, err := users.GetUser(999999) // Non-existent user
	if err != nil {
		fmt.Println("GetUser error:", err)
	}
}
