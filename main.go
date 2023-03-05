package main

import (
	"context"
	"log"
)

func main() {
	t := "Sato hai rotto il cazzo"

	s := NewSpeaker()
	if err := s.Speak(context.TODO(), t); err != nil {
		log.Fatal(err)
	}

}
