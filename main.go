package main

import (
	"github.com/ben-st/rettungsring/pkg/opts"
	"github.com/ben-st/rettungsring/pkg/sync"
	log "github.com/sirupsen/logrus"
)

func main() {
	o := opts.New()

	if err := o.ParseCmdArgs(); err != nil {
		log.Fatalln(err)
	}

	gc := sync.InitGitlabClient(o)

	if o.ListProjects {
		sync.ListUserProjects(gc, o)
	} else {
		sync.DownloadProjects(gc, o)
	}
}
