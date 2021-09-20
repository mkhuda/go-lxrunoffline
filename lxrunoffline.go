package lxrunoffline

import (
	"os/exec"
	"strings"
)

const (
	powershell           = "powershell.exe"
	lxRunOfflinelibsPath = "libs\\LxRunOffline.exe"
)

var (
	args_powershell    = []string{"-c", "chcp", "65001", ">", "$null", ";"}
	args_listInstalled = []string{"list"}
	args_Summary       = []string{"sm", "-n"}
	args_GetDefault    = []string{"gd"}
)

type Options struct {
	libsPath string
}

type LxRunOffline struct {
	Options
}

func Init(options Options) *LxRunOffline {
	if options.libsPath == "" {
		options.libsPath = lxRunOfflinelibsPath
	}

	lx := &LxRunOffline{
		Options: options,
	}

	return lx
}

func New() *LxRunOffline {
	lx := &LxRunOffline{
		Options{
			libsPath: lxRunOfflinelibsPath,
		},
	}

	return lx
}

func (lx *LxRunOffline) ListInstalled() ([]string, *exec.Cmd, error) {
	args := append(args_powershell, lx.libsPath)
	startCommand := append(args, args_listInstalled...)
	cmd := exec.Command(powershell, startCommand...)
	out, err := cmd.Output()

	sOutput := strings.Split(string(out), "\n")
	if len(sOutput) > 0 {
		sOutput = sOutput[:len(sOutput)-1]
	}

	return sOutput, cmd, err
}

func (lx *LxRunOffline) GetSummary(distributionName string) (string, *exec.Cmd, error) {
	args := append(args_powershell, lx.libsPath)
	summaryArgs := append(args_Summary, distributionName)
	startCommand := append(args, summaryArgs...)

	cmd := exec.Command(powershell, startCommand...)
	output, err := cmd.Output()

	return string(output), cmd, err
}

func (lx *LxRunOffline) GetDefaultDistro() (string, *exec.Cmd, error) {
	args := append(args_powershell, lx.libsPath)
	startCommand := append(args, args_GetDefault...)

	cmd := exec.Command(powershell, startCommand...)
	out, _ := cmd.Output()

	output := lx.ClearASCII(out, true)

	return output, cmd, nil
}
