package lxrunoffline

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func (lx *LxRunOffline) GetRegistry(path string, value string) (v string, vtype uint32) {
	k, err := registry.OpenKey(registry.CURRENT_USER, registryPath+path, registry.SZ)
	if err != nil {
		fmt.Println("error open registry", err, registry.CURRENT_USER, registryPath+path, registry.SZ)
	}
	defer k.Close()

	value, valueType, err := k.GetStringValue(value)
	if err != nil {
		fmt.Println("error get string", err)
	}
	return value, valueType
}

func addPathPrefix(str string) string {
	return str + "\\"
}
