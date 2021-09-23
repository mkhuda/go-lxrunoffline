package lxrunoffline

import (
	"strings"
)

func (lx *LxRunOffline) GetFirstLine(s string) string {
	sOutput := strings.Split(s, "\n")
	if len(sOutput) > 0 {
		sOutput = sOutput[:len(sOutput)-1]
	}

	return sOutput[0]
}

func (lx *LxRunOffline) ClearASCII(values []byte, shouldGetFirstLine bool) string {
	var ret []byte
	var output string
	for _, x := range values {
		if x != 0 {
			ret = append(ret, x)
		}
	}

	cleared := string(ret)

	if shouldGetFirstLine {
		output = lx.GetFirstLine(cleared)
	} else {
		output = cleared
	}

	return output
}

func (lx *LxRunOffline) IsWSL2(flag uint64) bool {
	return flag == uint64(wsl2_flags)
}
