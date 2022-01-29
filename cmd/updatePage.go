/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"notion-cli/config"
	"notion-cli/handler"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// updatePageCmd represents the updatePage command
var updatePageCmd = &cobra.Command{
	Use:   "page",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "not enough arguments in call: post page")
			os.Exit(1)
		}

		pageId := args[0]
		path := args[1]
		endpoint := config.PAGE_URL + pageId
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "not found url: "+path)
			os.Exit(1)
		}
		payload := strings.NewReader(string(bytes))
		req := handler.PatchHTTPRequester(endpoint, payload)
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))
	},
}

func init() {
	updateCmd.AddCommand(updatePageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatePageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatePageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
