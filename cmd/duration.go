/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/filariow/personal-trainer/pkg/spec"
	"github.com/spf13/cobra"
)

// durationCmd represents the duration command
var durationCmd = &cobra.Command{
	Use:   "duration",
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

		var td int64
		for _, e := range t.Exercises {
			td += e.DurationInSec
			td += e.Preparation
		}

		m := time.Duration(td) * time.Second
		fmt.Printf("Training duration is %s\n", m)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(durationCmd)
}
