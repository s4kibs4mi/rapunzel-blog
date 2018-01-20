package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

/**
 * := Coded with love by Sakib Sami on 20/1/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

var RootCmd = cobra.Command{}

func Execute() {
	RootCmd.AddCommand(&ServeCmd)
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
