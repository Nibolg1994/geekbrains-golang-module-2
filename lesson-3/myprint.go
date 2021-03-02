package myprint

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
)

type App struct {
	Port        string `yaml:"port"`
	DbUrl       string `yaml:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppId   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

var expressions = map[string]string{
	"port":         `\d{4}`,
	"db_url":       `\w+`,
	"jaeger_url":   `http://jaeger:\d{5}`,
	"sentry_url":   `http://sentry:\d{4}`,
	"kafka_broker": `\w+:\d{4}`,
	"some_app_id":  `\w{6}`,
	"some_app_key": `\w{7}`,
}

func GetConfig() (App, error) {
	config := App{}
	buffer, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		return config, err
	}
	if !config.validate() {
		return config, fmt.Errorf("inavalid config")
	}
	return config, nil
}

func (config App) validPort() bool {
	re := regexp.MustCompile(expressions["port"])
	return re.MatchString(config.Port)
}

func (config App) validDbUrl() bool {
	re := regexp.MustCompile(expressions["db_url"])
	return re.MatchString(config.DbUrl)
}

func (config App) validJaegerUrl() bool {
	re := regexp.MustCompile(expressions["jaeger_url"])
	return re.MatchString(config.JaegerUrl)
}

func (config App) validSentryUrl() bool {
	re := regexp.MustCompile(expressions["sentry_url"])
	return re.MatchString(config.SentryUrl)
}

func (config App) validKafkaBroker() bool {
	re := regexp.MustCompile(expressions["kafka_broker"])
	return re.MatchString(config.KafkaBroker)
}

func (config App) validSomeAppId() bool {
	re := regexp.MustCompile(expressions["some_app_id"])
	return re.MatchString(config.SomeAppId)
}

func (config App) validSomeAppKey() bool {
	re := regexp.MustCompile(expressions["some_app_key"])
	return re.MatchString(config.SomeAppKey)
}

func (config App) validate() bool {
	return config.validSomeAppKey() &&
		config.validSomeAppId() &&
		config.validKafkaBroker() &&
		config.validSentryUrl() &&
		config.validJaegerUrl() &&
		config.validDbUrl() &&
		config.validPort()
}

func Myprint(str string) error {
	config := App{}
	buffer, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		return err
	}
	if !config.validate() {
		return fmt.Errorf("inavalid config")
	}
	fmt.Println(config)
	return nil
}
