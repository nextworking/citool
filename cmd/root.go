package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)



var Path string

func init(){

	rootCmd.Flags().StringVar(&Path, "path",  "", "path")

}


var rootCmd = &cobra.Command{
	Use:   "citool",
	Short: "CI tools",
	Long: `Tools package to support a CI ecosystem`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
