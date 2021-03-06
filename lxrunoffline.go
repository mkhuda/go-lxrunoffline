package lxrunoffline

import "errors"

const (
	powershell             = "powershell.exe"
	registry_path          = "Software\\Microsoft\\Windows\\CurrentVersion\\Lxss\\"
	lxRunOffline_libs_main = "LxRunOffline.exe"
	lxRunOffline_libs_path = "libs\\LxRunOffline.exe"
)

var (
	args_powershell     = []string{"-c", "chcp", "65001", ">", "$null", ";"}
	args_list_installed = []string{"list"}
	args_summary        = []string{"sm", "-n"}
	args_get_default    = []string{"gd"}
	args_install        = func(nameOfDistro string, dirToInstall string, tarFile string) []string {
		return []string{"i", "-n", nameOfDistro, "-d", dirToInstall, "-f", tarFile}
	}
	args_export = func(nameOfDistro string, tarDirFile string) []string {
		return []string{"e", "-n", nameOfDistro, "-f", tarDirFile}
	}
	args_duplicate = func(nameOfCurrentDistro string, dirToDuplicate string, newDistroName string) []string {
		return []string{"d", "-n", nameOfCurrentDistro, "-d", dirToDuplicate, "-N", newDistroName}
	}

	registry_default_distro = "DefaultDistribution"
	registry_distro_name    = "DistributionName"
	registry_dir            = "BasePath"
	registry_state          = "State"
	registry_version        = "Version"
	registry_env            = "DefaultEnvironment"
	registry_uid            = "DefaultUid"
	registry_kernel_cmd     = "KernelCommandLine"
	registry_flags          = "Flags"

	wsl2_flags = 8
)

type LxRunOffline struct {
	Options
}
type Options struct {
	LibsPath string
}

// Init(options) can be used to obtain custom location of lxrunoffline.exe
func Init(options Options) *LxRunOffline {
	if options.LibsPath == "" {
		options.LibsPath = lxRunOffline_libs_path
	}

	lx := &LxRunOffline{
		Options: options,
	}

	return lx
}

/* New()
Initialize and find where lxrunoffline.exe is installed to the machine.
Use this initialize method if you have installed lxrunoffline via Chocolatey or Scoop.
Also if you correctly install lxrunoffline.exe manually then added to Windows PATH
*/
func New() (*LxRunOffline, error) {

	lxLocation, err := WhereLx()

	if err != nil {
		return nil, errors.New("no lxrunoffline installed")
	}

	lx := &LxRunOffline{
		Options{
			LibsPath: lxLocation,
		},
	}

	return lx, nil
}

/* lx.ListInstalled()
Return list of distro with some infos. This function read Windows Registrys
*/
func (lx *LxRunOffline) ListInstalled() ([]*Distro, error) {
	var distros []*Distro
	distro_uids, err := lx.GetRegistrySubkey(registry_path, "")

	if err != nil {
		return []*Distro{}, err
	}

	for i := range distro_uids {
		d, err := lx.GetDistroSummary(distro_uids[i])
		if err != nil {
			return []*Distro{}, err
		}
		distros = append(distros, d)
	}

	return distros, nil
}

/* lx.GetDefaultDistro()
Return distro name, distro_uid. Empty string will return if error occured. This function read Windows Registry
*/
func (lx *LxRunOffline) GetDefaultDistro() (string, string, error) {
	distro_uid, _, err := lx.GetRegistryValue("", registry_default_distro)
	if err != nil {
		return "", "", err
	}

	distro_name, _, err := lx.GetRegistryValue(addPathPrefix(distro_uid), registry_distro_name)
	if err != nil {
		return "", "", err
	}

	return distro_name, distro_uid, nil
}
