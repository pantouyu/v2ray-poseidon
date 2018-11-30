package v2ray_ssrpanel_plugin

import (
	"os"
	"time"
	"v2ray.com/core/common/errors"
)

func Run() {
	err := run()
	if err != nil {
		fatal(err)
	}
}

func run() error {
	commandLine.Parse(os.Args[1:])
	if *test {
		return testConfig()
	}

	cfg, err := getConfig()
	if err != nil {
		return err
	}
	db, err := NewMySQLConn(cfg.myPluginConfig.MySQL)
	if err != nil {
		return err
	}

	db.GetAllUsers()

	return nil
}

func newError(values ...interface{}) *errors.Error {
	values = append([]interface{}{"SSRPanelPlugin: "}, values...)
	return errors.New(values...)
}

func fatal(values ...interface{}) {
	newError(values...).AtError().WriteToLog()
	// Wait log
	time.Sleep(1 * time.Second)
	os.Exit(-2)
}
