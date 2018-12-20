package opts

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Opts contains the cmd args values
type Opts struct {
	GitlabAPIToken string
	GitlabAPIURL   string
	Verbose        bool
	ListProjects   bool
	User           string
}

// New initializes a new Opts struct
func New() *Opts {
	return &Opts{}
}

// ParseCmdArgs initializes vars, parses cmd args and validates them
func (o *Opts) ParseCmdArgs() error {
	flag.StringVar(&o.User, "user", "", "Gitlab user")
	flag.StringVar(&o.GitlabAPIToken, "token", "", "Gitlab api token")
	flag.StringVar(&o.GitlabAPIURL, "url", "", "Gitlab api url. Example: https://gitlab.com/api/v4")
	flag.BoolVar(&o.ListProjects, "listprojects", false, "lists all gitlab projects for current user")
	flag.BoolVar(&o.Verbose, "verbose", true, "Increase verbosity level to info")

	flag.Parse()

	if o.GitlabAPIToken == "" {
		return fmt.Errorf("No Gitlab API token provided")
	}

	if o.GitlabAPIURL == "" {
		return fmt.Errorf("No Gitlab URL provided")
	}

	if o.User == "" {
		return fmt.Errorf("No gitlab user specified")
	}

	if !o.Verbose {
		log.SetLevel(log.WarnLevel)
	}

	return nil
}
