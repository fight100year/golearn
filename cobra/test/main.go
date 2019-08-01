package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	s string
	i int
)

func init() {
	rootCmd.Flags().StringVarP(&s, "string", "s", "", "string val")
	rootCmd.Flags().IntVarP(&i, "int", "i", 0, "int val")
}

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "for test",
	Long:  `just for test`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(s) == 0 {
			fmt.Println("no string")
			return
		}
		fmt.Printf("string:%s, int:%d\n", s, i)
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
