package lxrunoffline

import (
	"os/exec"
	"strings"
)

func (lx *LxRunOffline) ListInstalledCmd() ([]string, *exec.Cmd, error) {
	args := append(args_powershell, lx.libsPath)
	start_command := append(args, args_list_installed...)
	cmd := exec.Command(powershell, start_command...)
	out, err := cmd.Output()

	sOutput := strings.Split(string(out), "\n")
	if len(sOutput) > 0 {
		sOutput = sOutput[:len(sOutput)-1]
	}

	return sOutput, cmd, err
}

func (lx *LxRunOffline) GetSummaryCmd(distributionName string) (string, *exec.Cmd, error) {
	args := append(args_powershell, lx.libsPath)
	summary_args := append(args_summary, distributionName)
	start_command := append(args, summary_args...)

	cmd := exec.Command(powershell, start_command...)
	output, err := cmd.Output()

	return string(output), cmd, err
}

func (lx *LxRunOffline) GetDefaultDistroCmd() (string, *exec.Cmd, error) {
	args := append(args_powershell, lx.libsPath)
	start_command := append(args, args_get_default...)

	cmd := exec.Command(powershell, start_command...)
	out, _ := cmd.Output()

	output := lx.ClearASCII(out, true)

	return output, cmd, nil
}
