package internal

import (
	"shared/executor"
)

const processName string = "launcher.hostsEditor.exe"
const ipFlag string = "-ip="
const addFlag string = "-add="

func run(admin bool, ip string, add bool) bool {
	var boolStr string
	if add {
		boolStr = "true"
	} else {
		boolStr = "false"
	}
	args := []string{ipFlag + ip, addFlag + boolStr}
	if admin {
		return executor.RunCustomExecutable(processName, args...)
	}
	return executor.ElevateCustomExecutable(processName, args...)
}

func AddHost(admin bool, ip string) bool {
	return run(admin, ip, true)
}

func RemoveHost(admin bool) bool {
	return run(admin, "", false)
}
