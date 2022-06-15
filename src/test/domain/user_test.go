package domain

import (
	"testing"
	"time"

	"github.com/kristiansantos/learning/src/core/domain"
)

func TestUser(t *testing.T) {
	checkUser := func(t *testing.T, user domain.User) {
		got := user
		expected := "1"

		if got.ID != expected {
			t.Errorf("got %v want %v", got, expected)
		}
	}

	t.Run("Sucess: Instance user struct'", func(t *testing.T) {
		timeStamp := time.Now()
		user := domain.User{"1", "testName", "test@email.com", "testPassword", true, timeStamp, timeStamp}

		checkUser(t, user)
	})
}
