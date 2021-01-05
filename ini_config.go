/*
List 2.11
*/
package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {
		fmt.Printf("Failed to parse config file: %s\n", err)
	}
	fmt.Println(config.Section)
	fmt.Println(config.Section.Path)
	fmt.Println(config.Section.Enabled)
}
