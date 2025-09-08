/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
)

var timefromzoneCmd = &cobra.Command {
	Use: "timefromzone",
	Short: "Return time from inputted geo zone",
	Long: `Command return time from inputted geo zone. Accepts only one argument - zone, which time need to known.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timefromzone := args[0]

		timeNow, err := getTimeFromZone(timefromzone)
		if err != nil {
			log.Fatalln("Incorect time zone")
		}

		fmt.Println(timeNow)
	},
}


func init() {
	rootCmd.AddCommand(timefromzoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timefromzoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timefromzoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
