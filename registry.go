package lxrunoffline

import (
	"errors"

	"golang.org/x/sys/windows/registry"
)

func (lx *LxRunOffline) GetRegistryValue(path string, value string) (v string, vtype uint32, err error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, registry_path+path, registry.SZ)
	if err != nil {
		return "", 0, errors.New("Can't OpenKey " + value + err.Error())
	}
	defer k.Close()

	value_string, value_type, err := k.GetStringValue(value)
	if err != nil {
		return "", value_type, errors.New("Can't GetStringValue " + value + err.Error())
	}

	return value_string, value_type, nil
}

func (lx *LxRunOffline) GetRegistryValueInt(path string, value string) (v uint64, vtype uint32, err error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, registry_path+path, registry.ALL_ACCESS)
	if err != nil {
		return 0, 0, errors.New("Can't OpenKey " + value + err.Error())
	}
	defer k.Close()

	value_int, value_type, err := k.GetIntegerValue(value)
	if err != nil {
		return 0, value_type, errors.New("Can't GetIntegerValue " + value + err.Error())
	}

	return value_int, value_type, nil
}

func (lx *LxRunOffline) GetRegistrySubkey(path string, value string) (v []string, err error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, path, registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return []string{""}, errors.New("Can't OpenKey " + value + err.Error())
	}
	defer k.Close()

	values, err := k.ReadSubKeyNames(0)
	if err != nil {
		return []string{""}, errors.New("Can't ReadSubKeyNames " + value + err.Error())
	}

	return values, nil
}

func addPathPrefix(str string) string {
	return str + "\\"
}
