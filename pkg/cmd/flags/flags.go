package flags

import (
	"os/user"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// RunOptions ...
type ZitiFlags struct {
	ZConfig string
	Service string

	Debug bool
}

func (f *ZitiFlags) GetUserAndIdentity(input string) (string, string) {
	username := ParseUserName(input)
	f.DebugLog("      username set to: %s", username)
	targetIdentity := ParseTargetIdentity(input)
	f.DebugLog("targetIdentity set to: %s", targetIdentity)
	return username, targetIdentity
}

func ParseUserName(input string) string {
	var username string
	if strings.ContainsAny(input, "@") {
		userServiceName := strings.Split(input, "@")
		username = userServiceName[0]
	} else {
		curUser, err := user.Current()
		if err != nil {
			logrus.Fatal(err)
		}
		username = curUser.Username
		if strings.Contains(username, "\\") && runtime.GOOS == "windows" {
			username = strings.Split(username, "\\")[1]
		}
	}
	return username
}

func ParseTargetIdentity(input string) string {
	var targetIdentity string
	if strings.ContainsAny(input, "@") {
		targetIdentity = strings.Split(input, "@")[1]
	} else {
		targetIdentity = input
	}

	if strings.Contains(targetIdentity, ":") {
		return strings.Split(targetIdentity, ":")[0]
	}
	return targetIdentity
}

func ParseFilePath(input string) string {
	if strings.Contains(input, ":") {
		colPos := strings.Index(input, ":") + 1
		return input[colPos:]
	}
	return input
}

func (f *ZitiFlags) DebugLog(msg string, args ...interface{}) {
	if f.Debug {
		logrus.Infof(msg, args...)
	}
}
