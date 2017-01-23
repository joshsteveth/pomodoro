package main

import (
	"errors"
	"time"

	"gopkg.in/gcfg.v1"
)

type (
	Config struct {
		Message map[string]*MessageConfig
	}

	IntervalConfig struct {
		WorkDuration  string
		PauseDuration string
	}

	//list of icons:
	//face-angel        face-laugh      face-smile
	//face-angry        face-monkey     face-smirk
	//face-cool         face-plain      face-surprise
	//face-crying       face-raspberry  face-tired
	//face-devilish     face-sad        face-uncertain
	//face-embarrassed  face-sick       face-wink
	//face-kiss         face-smile-big  face-worried
	MessageConfig struct {
		Title      string
		Message    string
		UseTimeout bool
		Timeout    int
		UseIcon    bool
		Icon       string
		Duration   string
		duration   time.Duration
	}
)

var (
	ConfigData Config

	noPauseError = errors.New("Pause config is not found")
	noWorkError  = errors.New("Work config is not found")
)

//Read config using gcfg lib
//receive flag config to determine config fila path
//otherwise use "main.ini" as default
func readConfig(filePath string) error {
	if err := gcfg.ReadFileInto(&ConfigData, filePath); err != nil {
		return err
	}

	return ConfigData.validate()
}

//basic validation for config data
//it is necessary to make sure that the program is runnning properly
func (c *Config) validate() error {
	if msg, ok := c.Message[pause]; !ok {
		return noPauseError
	} else {
		if err := msg.validate(); err != nil {
			return err
		}
	}

	if msg, ok := c.Message[work]; !ok {
		return noWorkError
	} else {
		if err := msg.validate(); err != nil {
			return err
		}
	}

	return nil
}

func (msg *MessageConfig) validate() error {
	if dur, err := time.ParseDuration(msg.Duration); err != nil {
		return err
	} else {
		msg.duration = dur
	}

	return nil
}
