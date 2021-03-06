package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	prowconfig "k8s.io/test-infra/prow/config"
	prowplugins "k8s.io/test-infra/prow/plugins"

	"github.com/openshift/ci-tools/pkg/deprecatetemplates"
)

type options struct {
	prowJobConfigDir     string
	prowConfigPath       string
	prowPluginConfigPath string
	allowlistPath        string
	prune                bool

	help bool
}

func bindOptions(fs *flag.FlagSet) *options {
	opt := &options{}

	fs.StringVar(&opt.prowJobConfigDir, "prow-jobs-dir", "", "Path to a root of directory structure with Prow job config files (ci-operator/jobs in openshift/release)")
	fs.StringVar(&opt.prowConfigPath, "prow-config-path", "", "Path to the Prow configuration file")
	fs.StringVar(&opt.prowPluginConfigPath, "prow-plugin-config-path", "", "Path to the Prow plugin configuration file")
	fs.StringVar(&opt.allowlistPath, "allowlist-path", "", "Path to template deprecation allowlist")
	fs.BoolVar(&opt.prune, "prune", false, "If set, remove from allowlist all jobs that either no longer exist or no longer use a template")

	return opt
}

func (o *options) validate() error {
	for param, value := range map[string]string{
		"--prow-jobs-dir":           o.prowJobConfigDir,
		"--prow-config-path":        o.prowConfigPath,
		"--prow-plugin-config-path": o.prowPluginConfigPath,
		"--allowlist-path":          o.allowlistPath,
	} {
		if value == "" {
			return fmt.Errorf("mandatory argument %s was not set", param)
		}
	}

	return nil
}

func main() {
	opt := bindOptions(flag.CommandLine)
	flag.Parse()

	if opt.help {
		flag.Usage()
		os.Exit(0)
	}

	if err := opt.validate(); err != nil {
		logrus.WithError(err).Fatal("Invalid parameters")
	}

	agent := prowplugins.ConfigAgent{}
	if err := agent.Load(opt.prowPluginConfigPath, true); err != nil {
		logrus.WithError(err).Fatal("Failed to read Prow plugin configuration")
	}
	pluginCfg := agent.Config().ConfigUpdater

	prowCfg, err := prowconfig.Load(opt.prowConfigPath, opt.prowJobConfigDir)
	if err != nil {
		logrus.WithError(err).Fatal("failed to load Prow configuration")
	}

	enforcer, err := deprecatetemplates.NewEnforcer(opt.allowlistPath)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize template deprecator")
	}

	enforcer.LoadTemplates(pluginCfg)
	enforcer.ProcessJobs(prowCfg)

	if opt.prune {
		enforcer.Prune()
	}

	if err := enforcer.SaveAllowlist(opt.allowlistPath); err != nil {
		logrus.WithError(err).Fatal("Failed to save template deprecation allowlist")
	}
}
