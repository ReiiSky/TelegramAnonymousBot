package module_test

import (
	"testing"

	"github.com/Satssuki/tele-anon-bot-queue/internal/module"
)

func TestEmptyDeque(t *testing.T) {
	var queue = module.SimpleQueue{}
	queue.Take()
}
