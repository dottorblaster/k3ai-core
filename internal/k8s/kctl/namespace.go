package kctl

import (
	"fmt"
)

func nameSpaceExists(config Config, nameSpace string) bool {
	cmd, args := prepareCommand(config, "get", "namespace", nameSpace, "--no-headers")
	err := execute(config, cmd, args...)
	if err != nil {
		return false
	}
	return true
}

func createNameSpace(config Config, nameSpace string) error {
	if !nameSpaceExists(config, nameSpace) {
		cmd, args := prepareCommand(config, create, "namespace", nameSpace)
		return execute(config, cmd, args...)
	}
	return nil
}

func deleteNameSpace(config Config, nameSpace string) error {
	if nameSpaceExists(config, nameSpace) {
		cmd, args := prepareCommand(config, delete, "namespace", nameSpace)
		return execute(config, cmd, args...)
	}
	return nil
}

func IsNameSpaceEmpty(config Config, nameSpace string) bool {
	if nameSpaceExists(config, nameSpace) {
		cmd, args := prepareCommand(config, "get", "all", "-n", nameSpace)
		err := execute(config, cmd, args...)
		fmt.Printf("%s\n", err)
		if err != nil {
			fmt.Printf("%s", err)
			return true
		}
	}
	return false
}
