package presenters

import "testing"

func TestBuildingANewAttachment(t *testing.T) {
    attachment := NewAttachment("Bazinga", "foo", "bar", "#fff", "footer")

    if attachment.Title != "Bazinga" {
        t.Error("Failed to set attachments title")
    }

    if attachment.Text != "bar" {
        t.Error("Failed to set attachments text")
    }

    if attachment.Footer != "footer" {
        t.Error("Failed to set attachemnts footer")
    }
}
