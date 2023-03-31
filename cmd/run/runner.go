// The run command to start estimation ceremony back-end service
package run

import (
	"estimation/ceremony/internal/log"
	"estimation/ceremony/internal/websocket"
	"os"

	"github.com/spf13/cobra"
)

var (
	host string
	port int

	runCmd = &cobra.Command{
		Use:   "start [OPTIONS]",
		Short: "Starts estimation ceremony back-end service",
		Long:  "Runs estimation ceremony HTTP server to handle client's full-duplex communication",

		Run: func(cmd *cobra.Command, args []string) {
			log.Infow("Executed estimation ceremony command")
			config := websocket.NewConfig(host, port)
			err := websocket.NewServer(config).HandleConnections()
			if err != nil {
				log.Errorw("Unexpected error while running WebSockets server", "err", err)
				os.Exit(1)
			}
		},
	}
)

func Execute() {
	runCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
	runCmd.Flags().StringVarP(&host, "host", "h", "", "Server host name without port")
	runCmd.MarkFlagRequired("host")
	runCmd.Flags().IntVarP(&port, "port", "p", -1, "Server port")
	runCmd.MarkFlagRequired("port")
	runCmd.Execute()
}
