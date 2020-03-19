package cmd

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/suzuki-shunsuke/github-comment/pkg/api"
	"github.com/suzuki-shunsuke/github-comment/pkg/config"
	"github.com/suzuki-shunsuke/github-comment/pkg/option"
	"github.com/urfave/cli/v2"
)

var postCommand = &cli.Command{
	Name:   "post",
	Usage:  "post a comment",
	Action: postAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "org",
			Usage: "GitHub organization name",
		},
		&cli.StringFlag{
			Name:  "repo",
			Usage: "GitHub repository name",
		},
		&cli.StringFlag{
			Name:    "token",
			Usage:   "GitHub API token",
			EnvVars: []string{"GITHUB_TOKEN", "GITHUB_ACCESS_TOKEN"},
		},
		&cli.StringFlag{
			Name:  "sha1",
			Usage: "commit sha1",
		},
		&cli.StringFlag{
			Name:  "template",
			Usage: "comment template",
		},
		&cli.StringFlag{
			Name:    "template-key",
			Aliases: []string{"k"},
			Usage:   "comment template key",
			Value:   "default",
		},
		&cli.StringFlag{
			Name:  "config",
			Usage: "configuration file path",
		},
		&cli.IntFlag{
			Name:  "pr",
			Usage: "GitHub pull request number",
		},
	},
}

func parsePostOptions(opts *option.PostOptions, c *cli.Context) {
	opts.Org = c.String("org")
	opts.Repo = c.String("repo")
	opts.Token = c.String("token")
	opts.SHA1 = c.String("sha1")
	opts.Template = c.String("template")
	opts.TemplateKey = c.String("template-key")
	opts.ConfigPath = c.String("config")
	opts.PRNumber = c.Int("pr")
}

func postAction(c *cli.Context) error {
	opts := &option.PostOptions{}
	parsePostOptions(opts, c)
	ctx := context.Background()
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	return api.Post(ctx, wd, opts, os.Getenv, func() bool {
		return terminal.IsTerminal(0)
	}, os.Stdin, &http.Client{}, existFile, config.Read)
}