package lxrunoffline

type Distro struct {
	DistroId              string
	DistroName            string
	WslVersion            uint64
	FileSystemVersion     uint64
	InstallationDirectory string
}

func (lx *LxRunOffline) GetDistroSummary(distro_uid string) (*Distro, error) {
	ds, _, err := lx.GetRegistryValue(addPathPrefix(distro_uid), registry_distro_name)
	if err != nil {
		return &Distro{}, err
	}
	fi, _, err := lx.GetRegistryValueInt(addPathPrefix(distro_uid), registry_version)
	if err != nil {
		return &Distro{}, err
	}
	wv, _, err := lx.GetRegistryValueInt(addPathPrefix(distro_uid), registry_flags)
	if err != nil {
		return &Distro{}, err
	}

	wsl_version := func() uint64 {
		if lx.IsWSL2(wv) {
			return 2
		} else {
			return 1
		}
	}()

	dir, _, err := lx.GetRegistryValue(addPathPrefix(distro_uid), registry_dir)
	if err != nil {
		return &Distro{}, err
	}

	d := &Distro{
		DistroId:              distro_uid,
		DistroName:            ds,
		FileSystemVersion:     fi,
		WslVersion:            wsl_version,
		InstallationDirectory: dir,
	}

	return d, err
}
