package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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

		http.HandleFunc("/*", handler)

		addr := fmt.Sprintf(":%s", *port)
		log.Info().Msgf("Server listening on port %s", *port)
		log.Fatal().Err(http.ListenAndServe(addr, nil)).Send()
	},
}

func handler(w http.ResponseWriter, req *http.Request) {
	normalizedOrigin := strings.TrimRight(*origin, "/")
	url := fmt.Sprintf("%s%s", normalizedOrigin, req.URL.Path)
	log.Info().Msgf("Requesting API: %s", url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	res.Write(w)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func init() {
	port = rootCmd.Flags().StringP("port", "p", "3000", "Port to run the proxy on")
	origin = rootCmd.Flags().StringP("origin", "o", "https://dummyjson.com", "Origin server to proxy requests to")
}
