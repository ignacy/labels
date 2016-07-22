package presenters

import "testing"

func TestBuildingANewMessage(t *testing.T) {
    message := NewMessage("Bazinga")

    if message.Text != "Bazinga" {
        t.Error("Failed to set Message text")
    }

    if message.Username != "github-update" {
        t.Error("Failed to set default username")
    }

    if message.IconEmoji != ":ghost:" {
        t.Error("Failed to set default emoji")
    }
}
