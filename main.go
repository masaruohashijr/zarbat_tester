package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"zarbat_test/godog/dial"
	"zarbat_test/godog/hangup"
	"zarbat_test/godog/number"
	"zarbat_test/godog/pause"
	"zarbat_test/godog/ping"
	"zarbat_test/godog/play"
	"zarbat_test/godog/record"
	"zarbat_test/godog/redirect"
	"zarbat_test/godog/reject"
	"zarbat_test/godog/say"
	"zarbat_test/godog/sms"

	"github.com/cucumber/godog"
)

var RegMap map[string]FeatureTest

func initRegister() {
	RegMap = make(map[string]FeatureTest)
	RegMap["play"] = FeatureTest{
		Path:                "features/play",
		ScenarioInitializer: play.InitializeScenario,
	}
	RegMap["ping"] = FeatureTest{
		Path:                "features/ping",
		ScenarioInitializer: ping.InitializeScenario,
	}
	RegMap["pause"] = FeatureTest{
		Path:                "features/pause",
		ScenarioInitializer: pause.InitializeScenario,
	}
	RegMap["dial"] = FeatureTest{
		Path:                "features/dial",
		ScenarioInitializer: dial.InitializeScenario,
	}
	RegMap["redirect"] = FeatureTest{
		Path:                "features/redirect",
		ScenarioInitializer: redirect.InitializeScenario,
	}
	RegMap["reject"] = FeatureTest{
		Path:                "features/reject",
		ScenarioInitializer: reject.InitializeScenario,
	}
	RegMap["hangup"] = FeatureTest{
		Path:                "features/hangup",
		ScenarioInitializer: hangup.InitializeScenario,
	}
	RegMap["say"] = FeatureTest{
		Path:                "features/say",
		ScenarioInitializer: say.InitializeScenario,
	}
	RegMap["record"] = FeatureTest{
		Path:                "features/record",
		ScenarioInitializer: record.InitializeScenario,
	}
	RegMap["buy"] = FeatureTest{
		Path:                "features/number",
		ScenarioInitializer: number.InitializeScenario,
	}
	RegMap["sms"] = FeatureTest{
		Path:                "features/sms",
		ScenarioInitializer: sms.InitializeScenario,
	}
}

func main() {
	initRegister()
	tests := initArgs()
	status := 0
	for _, a := range tests {
		ft := RegMap[a]
		opts := godog.Options{
			Format:    "progress",
			Paths:     []string{ft.Path},
			Randomize: time.Now().UTC().UnixNano(),
		}
		println(opts.Paths[0], a)
		status = godog.TestSuite{
			Name:                "zarbat_test",
			ScenarioInitializer: ft.ScenarioInitializer,
			Options:             &opts,
		}.Run()
	}
	os.Exit(status)
}

type FeatureTest struct {
	Path                 string
	ScenarioInitializer  func(ctx *godog.ScenarioContext)
	TestSuiteInitializer func(ctx *godog.TestSuiteContext)
}

func initArgs() []string {
	var tests []string
	configPtr := flag.String("config", "config/config.txt", "a configuration file")
	triesPtr := flag.String("n", "2", "number of tries")
	logPtr := flag.String("l", "log/zarbat.log", "log location")
	logLevelPtr := flag.String("level", "info", "logging level")
	testPtr := flag.String("test", "buy", "ctlang")
	flag.Parse()
	addons := flag.Args()
	tests = append(tests, *testPtr)
	for _, a := range addons {
		tests = append(tests, a)
	}
	fmt.Println("************************************************")
	fmt.Println("*** Config:", *configPtr)
	fmt.Println("*** Number of Tries:", *triesPtr)
	fmt.Println("*** Log:", *logPtr)
	fmt.Println("*** Logging Level:", *logLevelPtr)
	fmt.Printf("*** Tests: %+q\n", tests)
	fmt.Println("************************************************")
	return tests
}
