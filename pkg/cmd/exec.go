package cmd

import (
	"context"
	"os"

	"github.com/suzuki-shunsuke/github-comment/pkg/api"
	"github.com/suzuki-shunsuke/github-comment/pkg/config"
	"github.com/suzuki-shunsuke/github-comment/pkg/option"
	"github.com/urfave/cli/v2"
)

var execCommand = &cli.Command{
	Name:   "exec",
	Usage:  "post a command result as a comment",
	Action: execAction,
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

func parseExecOptions(opts *option.ExecOptions, c *cli.Context) {
	opts.Org = c.String("org")
	opts.Repo = c.String("repo")
	opts.Token = c.String("token")
	opts.SHA1 = c.String("sha1")
	opts.TemplateKey = c.String("template-key")
	opts.ConfigPath = c.String("config")
	opts.PRNumber = c.Int("pr")
	opts.Args = c.Args().Slice()
}

func existFile(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func execAction(c *cli.Context) error {
	opts := &option.ExecOptions{}
	parseExecOptions(opts, c)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	ctx := context.Background()
	return api.Exec(ctx, wd, opts, os.Getenv, existFile, config.Read)
}