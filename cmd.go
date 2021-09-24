package lxrunoffline

import (
	"os/exec"
	"strings"
)

func WhereLx() (string, error) {
	cmd := exec.Command("where", lxRunOffline_libs_main)
	out, err := cmd.Output()

	lx_location := strings.Split(string(out), "\r")
	if len(lx_location) > 0 {
		lx_location = lx_location[:len(lx_location)-1]
	}

	return strings.TrimSuffix(lx_location[0], "\r"), err
}

func (lx *LxRunOffline) ListInstalledCmd() ([]string, *exec.Cmd, error) {
	args := append(args_powershell, lx.LibsPath)
	start_command := append(args, args_list_installed...)
	cmd := exec.Command(powershell, start_command...)
	out, err := cmd.Output()

	sOutput := strings.Split(string(out), "\r")
	if len(sOutput) > 0 {
		sOutput = sOutput[:len(sOutput)-1]
	}

	return sOutput, cmd, err
}

func (lx *LxRunOffline) GetSummaryCmd(distributionName string) (string, *exec.Cmd, error) {
	args := append(args_powershell, lx.LibsPath)
	summary_args := append(args_summary, distributionName)
	start_command := append(args, summary_args...)

	cmd := exec.Command(powershell, start_command...)
	output, err := cmd.Output()

	return string(output), cmd, err
}

func (lx *LxRunOffline) GetDefaultDistroCmd() (string, *exec.Cmd, error) {
	args := append(args_powershell, lx.LibsPath)
	start_command := append(args, args_get_default...)

	cmd := exec.Command(powershell, start_command...)
	out, err := cmd.Output()

	output := lx.ClearASCII(out, true)

	return output, cmd, err
}
