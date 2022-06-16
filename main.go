package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"gomodss/mypackage"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")

			mypackage.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	err := cmd.Execute()
	if err != nil {
		return
	}
}
