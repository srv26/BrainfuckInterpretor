package interpretor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/srv26/brainfucklibrary"
)

var rootCmd *cobra.Command

func NewRootCmd() *cobra.Command {
	return &cobra.Command{}
}

func Initialize(r *strings.Reader) {
	// Create a fresh rootCmd
	rootCmd = NewRootCmd()
	addExecuteCommand(r)
}

func addExecuteCommand(r *strings.Reader) {
	interpretCmd := &cobra.Command{
		Use:   "interpret [scan path]",
		Short: "Brain Fuck interpreter accept one command at a time and executes",
		RunE: func(cmd *cobra.Command, arg []string) error {

			output := new(bytes.Buffer)
			mydata := brainfucklibrary.GetData()
			mydata.Input = os.Stdin
			mydata.Output = output
			mydata.Memory = make(map[int]byte, 0)
			var err error
			buf := make([]byte, 1)
			for {
				_, err := r.Read(buf)
				if err == io.EOF {
					break
				}
				mydata.Program = buf[0]
				// Reads and executes one command at a time
				err = mydata.Run()
				if err != nil {
					fmt.Println(err)
				}
			}
			if err == nil {
				fmt.Println(string(output.Bytes()))
			}
			return nil
		},
	}

	rootCmd.AddCommand(interpretCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
