package test

import (
	"fmt"

	"github.com/dll02/webgo/framework/cobra"
)

func InitTest() *cobra.Command {
	 
	return TestCommand
}

var TestCommand = &cobra.Command{
	Use:   "test",
	Short: "test",
	RunE: func(c *cobra.Command, args []string) error {
        container := c.GetContainer()
		fmt.Println(container)
		return nil
	},
}

