/*
Copyright © 2023 Francesco Ilario
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/filariow/personal-trainer/pkg/spec"
	"github.com/filariow/personal-trainer/pkg/speech"
	"github.com/filariow/personal-trainer/pkg/types"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// this does the trick
		var inputReader io.Reader = cmd.InOrStdin()

		// the argument received looks like a file, we try to open it
		if len(args) > 0 && args[0] != "-" {
			file, err := os.Open(args[0])
			if err != nil {
				return fmt.Errorf("failed open file: %w", err)
			}
			inputReader = file
		}

		// we process the input reader, wherever to be his origin
		t, err := spec.Read(cmd.Context(), inputReader)
		if err != nil {
			return fmt.Errorf("failed process input: %w", err)
		}

		log.Printf("training: %v\n", t)

		s := speech.NewSpeaker()
		for _, e := range t.Exercises {
			getReady(cmd.Context(), s, e)

			doExercise(cmd.Context(), s, e)
		}

		return s.Speak(cmd.Context(), "Hai finito compa'")
	},
}

func getReady(ctx context.Context, speaker speech.Speaker, exercise types.Exercise) {
	_ = speaker.Speak(ctx, fmt.Sprintf("Preparati per %s", exercise.Name))
	time.Sleep(time.Duration(exercise.Preparation-5) * time.Second)

	countdown(ctx, speaker, 5)
}

func doExercise(ctx context.Context, speaker speech.Speaker, exercise types.Exercise) {
	cw := exercise.DurationInSec / 2

	_ = speaker.Speak(ctx, exercise.Name)
	time.Sleep(time.Duration(cw) * time.Second)

	countdown(ctx, speaker, cw)
}

func countdown(ctx context.Context, s speech.Speaker, count int64) {
	t := time.NewTicker(1 * time.Second)

	for c := count; c > 0; c-- {
		<-t.C

		if err := s.Speak(ctx, fmt.Sprintf("%d", c)); err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
