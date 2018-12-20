package sync

import (
	"os"
	"os/exec"

	"github.com/ben-st/rettungsring/pkg/opts"
	log "github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
)

// ListUserProjects returns all found projects for given User
func ListUserProjects(gc *gitlab.Client, o *opts.Opts) {

	projects := getProjects(gc, o)
	for _, p := range projects {
		log.Infof("Project \"%s\" (ID: %d) at %s", p.Name, p.ID, p.WebURL)
	}

	log.Infoln("Listing done. No changes were made until now.")
}

// InitGitlabClient inits the gitlab client
func InitGitlabClient(o *opts.Opts) *gitlab.Client {
	git := gitlab.NewClient(nil, o.GitlabAPIToken)
	git.SetBaseURL(o.GitlabAPIURL)

	return git
}

// getProjectPages returns the total pages
func getProjectPages(gc *gitlab.Client, o *opts.Opts) int {
	var userName interface{}
	userName = o.User

	opt := &gitlab.ListProjectsOptions{}
	_, resp, err := gc.Projects.ListUserProjects(userName, opt)

	if err != nil {
		log.Fatalf("Cannot get project pages: %s", err)
	}

	log.Infof("Found %d project pages", resp.TotalPages)

	return resp.TotalPages
}

// getProjectsForPage returns all projects for all pages
func getProjectsForPage(gc *gitlab.Client, page int, o *opts.Opts) []*gitlab.Project {
	opt := &gitlab.ListProjectsOptions{}
	opt.Page = page
	opt.PerPage = 100

	var userName interface{}
	userName = o.User

	projects, _, err := gc.Projects.ListUserProjects(userName, opt)

	if err != nil {
		log.Fatalf("Cannot get projects for page: %d %s", page, err)
	}

	return projects
}

func getProjects(gc *gitlab.Client, o *opts.Opts) []gitlab.Project {
	projects := []gitlab.Project{}

	pages := getProjectPages(gc, o)
	for page := 1; page <= pages; page++ {
		projectsPerPage := getProjectsForPage(gc, page, o)

		for _, project := range projectsPerPage {
			projects = append(projects, *project)
		}
	}

	return projects
}

// DownloadProjects will download all projects to current folder
func DownloadProjects(gc *gitlab.Client, o *opts.Opts) {

	projects := getProjects(gc, o)

	log.Infof("Found %d projects", len(projects))
	dirName := "repos"
	createDirIfNotExist(dirName)

	err := os.Chdir(dirName)
	if err != nil {
		panic(err)
	}

	for _, p := range projects {
		log.Infof("cloning repo: %s into folder %s", p.Name, dirName)
		cmd := exec.Command("git", "clone", p.SSHURLToRepo)
		err = cmd.Run()
		if err != nil {
			log.Errorf("clone failed for project %s, with error: %v", p.Name, err)
		}
	}
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
