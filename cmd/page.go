/*
Copyright © 2022 HIRATA Tomonori <tomonori4565@icloud.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"notion-cli/config"
	"notion-cli/handler"
	"notion-cli/presenter"

	"github.com/spf13/cobra"
)

var (
	pretty  bool
	maximum int
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page [page_id1] [page_id2] ...",
	Short: "get page information",
	Long:  `get page information`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "not enough arguments in call: get page")
			os.Exit(1)
		}
		if len(args) > maximum {
			fmt.Fprintln(os.Stderr, "number of pages must be less than "+strconv.Itoa(maximum))
			fmt.Fprintln(os.Stderr, "if you want to enlarge maximum number of pages, use --maximum option.")
			os.Exit(1)
		}

		var byteArray []byte
		byteArray = append(byteArray, []byte("[")...)
		for i := 0; i < len(args); i++ {
			pageId := args[i]
			endpoint := config.PAGE_URL + pageId

			req := handler.GetHTTPRequester(endpoint, nil)
			client := new(http.Client)
			resp, _ := client.Do(req)
			defer resp.Body.Close()

			tmp, _ := ioutil.ReadAll(resp.Body)
			byteArray = append(byteArray, tmp...)
			if i != (len(args) - 1) {
				byteArray = append(byteArray, []byte(",")...)
			}
		}
		byteArray = append(byteArray, []byte("]")...)
		presenter.StdOutput(byteArray, pretty)
	},
}

func init() {
	getCmd.AddCommand(pageCmd)
	pageCmd.PersistentFlags().BoolVar(&pretty, "pretty", false, "format the json output")
	pageCmd.PersistentFlags().IntVar(&maximum, "maximum", 5, "change the upper bound of pages number")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
