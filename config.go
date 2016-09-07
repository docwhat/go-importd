package main

import (
	"log"
	"os"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type appConfig struct {
	listenAddress string
	githubUserURL string
	importDomain  string
}

func parseFlags(args []string) appConfig {
	config := appConfig{}

	app := kingpin.New("go-importd", "Redirector for go imports for custom domains.")
	app.Writer(os.Stdout)
	app.HelpFlag.Short('h')
	app.Author("Christian Höltje https://docwhat.org/")
	app.Version(appVersion)
	app.VersionFlag.Short('v')

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Unable to get hostname: %s", err)
	}

	app.Flag("listen", "The address to serve HTTP on").
		Short('l').
		Default(":http").
		OverrideDefaultFromEnvar("GO_IMPORTD_LISTEN").
		StringVar(&config.listenAddress)

	app.Flag("import-domain", "The domain for imports. Usually this hostname.").
		Short('i').
		Default(hostname).
		OverrideDefaultFromEnvar("GO_IMPORTD_IMPORT_DOMAIN").
		StringVar(&config.importDomain)

	app.Flag("github-user-url", "The base URL on github for your projects.").
		Short('g').
		Default("https://github.com/docwhat").
		OverrideDefaultFromEnvar("GO_IMPORTD_GITHUB_USER_URL").
		StringVar(&config.githubUserURL)

	if command, err := app.Parse(args); err != nil {
		app.Usage(nil)
	} else {
		if command != "" {
			app.Usage(nil)
		}
	}

	config.githubUserURL = strings.TrimSuffix(config.githubUserURL, "/") + "/"
	config.importDomain = strings.TrimSuffix(config.importDomain, "/") + "/"

	return config
}
