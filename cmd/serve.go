package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the reverse proxy server",
	Long: `Start the reverse proxy server with the specified configuration.
The server will use Let's Encrypt for SSL certificates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		startProxyServer()
	},
}

func init() {
	fmt.Println("serve.go init()")
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
