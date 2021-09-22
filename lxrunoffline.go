package lxrunoffline

import (
	"log"
)

const (
	powershell             = "powershell.exe"
	registry_path          = "Software\\Microsoft\\Windows\\CurrentVersion\\Lxss\\"
	lxRunOffline_libs_path = "libs\\LxRunOffline.exe"
)

var (
	args_powershell     = []string{"-c", "chcp", "65001", ">", "$null", ";"}
	args_list_installed = []string{"list"}
	args_summary        = []string{"sm", "-n"}
	args_get_default    = []string{"gd"}

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
	libsPath string
}

type Distro struct {
	distroId              string
	distroName            string
	wslVersion            uint64
	fileSystemVersion     uint64
	installationDirectory string
	configurationFlags    uint64
}

func Init(options Options) *LxRunOffline {
	if options.libsPath == "" {
		options.libsPath = lxRunOffline_libs_path
	}

	lx := &LxRunOffline{
		Options: options,
	}

	return lx
}

func New() *LxRunOffline {
	lx := &LxRunOffline{
		Options{
			libsPath: lxRunOffline_libs_path,
		},
	}

	return lx
}

func (lx *LxRunOffline) ListInstalled() ([]Distro, error) {
	var distros []Distro
	distro_uids, err := lx.GetRegistrySubkey(registry_path, "")

	if err != nil {
		log.Println(err)
	}

	for i := range distro_uids {
		_, _, err := lx.GetRegistryValue(addPathPrefix(distro_uids[i]), registry_distro_name)
		if err != nil {
			log.Println(err)
		}
		distros = append(distros, *lx.GetDistroSummary(distro_uids[i]))
	}

	return distros, nil
}

func (lx *LxRunOffline) GetDefaultDistro() (string, error) {
	distro_uid, _, err := lx.GetRegistryValue("", registry_default_distro)
	if err != nil {
		log.Println(err)
	}

	distro_name, _, err := lx.GetRegistryValue(addPathPrefix(distro_uid), registry_distro_name)
	if err != nil {
		log.Println(err)
	}

	return distro_name, nil
}

func (lx *LxRunOffline) GetDistroSummary(distro string) *Distro {
	ds, _, err := lx.GetRegistryValue(addPathPrefix(distro), registry_distro_name)
	if err != nil {
		log.Println("boo", err)
	}
	fi, _, err := lx.GetRegistryValueInt(addPathPrefix(distro), registry_version)
	if err != nil {
		log.Println("wa", err, registry_version)
	}
	wv, _, err := lx.GetRegistryValueInt(addPathPrefix(distro), registry_flags)

	if err != nil {
		log.Println("wa", err, registry_version)
	}

	wsl_version := func() uint64 {
		if lx.IsWSL2(wv) {
			return 2
		} else {
			return 1
		}
	}()

	d := &Distro{
		distroId:          distro,
		distroName:        ds,
		fileSystemVersion: fi,
		wslVersion:        wsl_version,
	}

	return d
}
