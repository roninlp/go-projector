package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/theprimeagen/projector/pkg/cli"
)

func main() {
	opts, err := cli.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	config, err := cli.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v", err)
	}

	proj := cli.NewProjector(config)
	if config.Operation == cli.Print {
		if len(config.Args) == 0 {
			data := proj.GetValueAll()
			jsonString, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("this line should never be reached %v", err)
			}

			fmt.Printf("%v", string(jsonString))
		} else if value, ok := proj.GetValue(config.Args[0]); ok {
			fmt.Printf("%v", value)
		}
	}

	if config.Operation == cli.Add {
		proj.SetValue(config.Args[0], config.Args[1])
		proj.Save()
	}
	if config.Operation == cli.Remove {
		proj.RemoveValue(config.Args[0])
		proj.Save()
	}
}
