package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	port   *string
	origin *string
)

var rootCmd = &cobra.Command{
	Use:   "cache-proxy",
	Short: "Cache proxy is a simple HTTP proxy that caches responses.",
	Long:  "Cache proxy is a simple HTTP proxy that caches responses.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		http.HandleFunc("/*", func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "Hello, world!\n")
		})

		addr := fmt.Sprintf(":%s", *port)
		log.Info().Msgf("Server listening on port %s", *port)
		log.Fatal().Err(http.ListenAndServe(addr, nil)).Send()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func init() {
	port = rootCmd.Flags().StringP("port", "p", "3000", "Port to run the proxy on")
	origin = rootCmd.Flags().StringP("origin", "o", "http://dummyjson.com", "Origin server to proxy requests to")
}
