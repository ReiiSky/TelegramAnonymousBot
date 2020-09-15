package service

import (
	"time"

	"github.com/Satssuki/tele-anon-bot-queue/internal/domain/model"
	"github.com/Satssuki/tele-anon-bot-queue/internal/module"
	"github.com/Satssuki/tele-anon-bot-queue/pkg"
)

var userSessionCache = module.SafeCache{}

// ReadUserPartner ..
func ReadUserPartner(userID string) (string, bool) {
	return userSessionCache.Read(userID)
}

// CountOfPartner ..
func CountOfPartner() int {
	return userSessionCache.Length() / 2
}

// WritePartner ..
func WritePartner(userID, partnerID string) {
	userSessionCache.Insert(userID, partnerID)
	userSessionCache.Insert(partnerID, userID)
}

// RemovePartner ..
func RemovePartner(userID, partnerID string) {
	userSessionCache.Delete(userID)
	userSessionCache.Delete(partnerID)
}

// UserWorkerQueue ..
func UserWorkerQueue() {
	go func() {
		compiledHiMessage := pkg.
			GetDefaultTeleClient().
			TextMessageBuilder().
			Content("You get the partner, say hii!!")
		defUserQueue := model.DefaultUserQueue()
		for {
			if defUserQueue.Len() > 1 {

				userID := defUserQueue.Take().(string)
				compiledHiMessage.ChatID(userID)
				pkg.GetDefaultTeleClient().Push(compiledHiMessage)

				partnerID := defUserQueue.Take().(string)
				compiledHiMessage.ChatID(partnerID)
				pkg.GetDefaultTeleClient().Push(compiledHiMessage)

				WritePartner(userID, partnerID)
			}
			time.Sleep(time.Millisecond * 15)
		}
	}()
}
