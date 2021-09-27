package lxrunoffline

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Distro struct {
	DistroId              string
	DistroName            string
	WslVersion            uint64
	FileSystemVersion     uint64
	InstallationDirectory string
}

/* lx.GetDistroSummary(distro_uid string)
Return a distro with the info it has. This function read Windows Registry
*/
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

/* *Distro.DirSize()
Return a prettier string of size of distro installation folder on disk
*/
func (ds *Distro) DirSize() string {
	var stringSize string
	size := diskUsage(ds.InstallationDirectory)

	switch {
	case size > 1024*1024*1024:
		stringSize = fmt.Sprintf("%.1fG", float64(size)/(1024*1024*1024))
	case size > 1024*1024:
		stringSize = fmt.Sprintf("%.1fM", float64(size)/(1024*1024))
	case size > 1024:
		stringSize = fmt.Sprintf("%.1fK", float64(size)/1024)
	default:
		stringSize = fmt.Sprintf("%d", size)
	}
	return stringSize
}

/* diskUsage(path string)
Return int64 of disk usage of current distro. This function used by *Distro.DirSize().
*/
func diskUsage(path string) int64 {
	var size int64
	pathString := strings.TrimLeft(path, `\?`)
	dir, err := os.Open(pathString)
	if err != nil {
		return size
	}
	defer dir.Close()

	filesOnDir, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for _, file := range filesOnDir {
		if !file.IsDir() {
			size += file.Size()
		}
		if file.IsDir() {
			size += diskUsage(fmt.Sprintf("%s/%s", pathString, file.Name()))
		} else {
			size += file.Size()
		}
	}
	return size
}
