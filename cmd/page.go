/*
Copyright © 2022 HIRATA Tomonori <tomonori4565@icloud.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"notion-cli/handler"
	"notion-cli/presenter"

	"github.com/spf13/cobra"
)

var (
	pretty bool
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page [page_id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "not enough arguments in call: get page")
			os.Exit(1)
		}
		pageId := args[0]
		endpoint := "https://api.notion.com/v1/pages/" + pageId

		req := handler.GetHTTPHandler(endpoint, nil)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		presenter.StdOutput(byteArray, pretty)
	},
}

func init() {
	getCmd.AddCommand(pageCmd)
	pageCmd.PersistentFlags().BoolVar(&pretty, "pretty", false, "pretty json output")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
