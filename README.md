# LxRunOffline on Golang

[LxRunOffline](https://github.com/DDoSolitary/LxRunOffline/) on golang.

> Still on going development

# Installing

Install

```bash
$ go get github.com/mkhuda/go-lxrunoffline
```

# Usage

```go
package main

import (
	"fmt"

	"github.com/mkhuda/go-lxrunoffline"
)

func main() {
	lx := lxrunoffline.New()

	listInstalled, _, err := lx.ListInstalled()
	if err != nil {
		fmt.Println("error listinstalled", err)
		return
	}

	defaultDistro, _, err := lx.GetDefaultDistro()
	if err != nil {
		fmt.Println("error getdefaultdistro", err)
		return
	}

	fmt.Println("List of installed WSL: ")
	for i, distributionName := range listInstalled {
		fmt.Println(i+1, distributionName)
	}

	summaryOfDefaultDistro, cmd, err := lx.GetSummary(defaultDistro)
	if err != nil {
		fmt.Println("error summary", err, cmd)
		return
	}

	fmt.Printf("Summary of default distro: %v\n%v", defaultDistro, summaryOfDefaultDistro)
}

```

# Credits

- LxRunOffline Creator: [DDoSolitary](https://github.com/DDoSolitary/LxRunOffline/)
- CMake: [CMake](https://cmake.org/)
