package lib

import (
	"fmt"
	"os"

	"github.com/Satssuki/tele-anon-bot-queue/internal/service"
	"github.com/Satssuki/tele-anon-bot-queue/pkg"
)

// MessageDeliverQueue queue to delivering message through worker
// i using channel to prevent goroutine ownership race
var MessageDeliverQueue = make(chan *pkg.Message)

// StartWorker ..
func StartWorker(workerCount int) {
	for i := 0; i < workerCount; i++ {
		workerFunction()
	}
}

func workerFunction() {
	c := pkg.GetDefaultTeleClient()
	username := os.Getenv("BOTUSERNAME")
	go func() {
		for {
			message := <-MessageDeliverQueue
			partnerID, isExist := service.ReadUserPartner(fmt.Sprint(message.Content.From.ID))
			if isExist && message.Content.Reply.From.Username != username {

				textBuilder := c.TextMessageBuilder().ChatID(partnerID).Content(message.Content.Text)
				replyID := message.Content.Reply.MessageID

				if replyID > 0 {
					textBuilder.ReplyTo(fmt.Sprint(replyID))
				}
				_, err := c.Push(textBuilder)

				if err != nil {
					MessageDeliverQueue <- message
				}
			}
		}
	}()
}
