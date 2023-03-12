package speech

import (
	"context"

	htgotts "github.com/hegedustibor/htgo-tts"
)

type Speaker interface {
	Speak(ctx context.Context, text string) error
}

func NewSpeaker() Speaker {
	return &speaker{}
}

type speaker struct{}

func (s *speaker) Speak(ctx context.Context, text string) error {
	speech := htgotts.Speech{Folder: "audio", Language: "it"}
	return speech.Speak(text)
}
