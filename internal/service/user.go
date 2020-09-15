package service

import (
	"time"

	"github.com/Satssuki/tele-anon-bot-queue/internal/domain/model"
	"github.com/Satssuki/tele-anon-bot-queue/internal/module"
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
		defUserQueue := model.DefaultUserQueue()
		for {
			if defUserQueue.Len() > 1 {
				userID := defUserQueue.Take().(string)
				partnerID := defUserQueue.Take().(string)
				WritePartner(userID, partnerID)
			}
			time.Sleep(time.Millisecond * 15)
		}
	}()
}
