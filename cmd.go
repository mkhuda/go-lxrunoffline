package lxrunoffline

import (
	"os/exec"
	"strings"
)

/* WhereLx()
Show current lxrunoffline installation from system. */
func WhereLx() (string, error) {
	cmd := exec.Command("where", lxRunOffline_libs_main)
	out, err := cmd.Output()

	lx_location := strings.Split(string(out), "\r")
	if len(lx_location) > 0 {
		lx_location = lx_location[:len(lx_location)-1]
	}

	return strings.TrimSuffix(lx_location[0], "\r"), err
}

/* lx.ListInstalledCmd()
Get list of installed distro on your machine. This function provided by LxRunOffline.exe and return list of distroName string */
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

/* lx.GetSummaryCmd(distributionName string)
Will show current summary by distributionName. This function provided by LxRunOffline.exe */
func (lx *LxRunOffline) GetSummaryCmd(distributionName string) (string, *exec.Cmd, error) {
	args := append(args_powershell, lx.LibsPath)
	summary_args := append(args_summary, distributionName)
	start_command := append(args, summary_args...)

	cmd := exec.Command(powershell, start_command...)
	output, err := cmd.Output()

	return string(output), cmd, err
}

/* lx.GetDefaultDistroCmd()
Will show your default distro if you run `wsl` on command prompt.
This function provided by LxRunOffline.exe. Return is string of distroName */
func (lx *LxRunOffline) GetDefaultDistroCmd() (string, *exec.Cmd, error) {
	args := append(args_powershell, lx.LibsPath)
	start_command := append(args, args_get_default...)

	cmd := exec.Command(powershell, start_command...)
	out, err := cmd.Output()

	output := lx.ClearASCII(out, true)

	return output, cmd, err
}

/* lx.ExportDistro(distributionName string, tarDirFile string)
Will export a distro into `*.tar.gz` file. You should provide the full path with name for `tarDirFile` params, e.g: "G:\WSL_Backup\debian.tar.gz". This function provided by LxRunOffline.exe */
func (lx *LxRunOffline) ExportDistro(distributionName string, tarDirFile string) error {
	args := append(args_powershell, lx.LibsPath)
	export_args := args_export(distributionName, tarDirFile)
	start_command := append(args, export_args...)

	cmd := exec.Command(powershell, start_command...)
	_, err := cmd.Output()
	cmd.Wait()

	return err
}
