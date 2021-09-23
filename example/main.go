package main

import (
	"fmt"

	"github.com/mkhuda/go-lxrunoffline"
)

func main() {
	lx := lxrunoffline.New()

	listInstalled, err := lx.ListInstalled()
	if err != nil {
		fmt.Println("error listinstalled", err)
		return
	}

	defaultDistroName, _ := lx.GetDefaultDistro()

	fmt.Println("List of installed WSL: ")
	for i, distributionName := range listInstalled {
		fmt.Println(i+1, distributionName.DistroName)
	}

	summaryOfDefaultDistro, cmd, err := lx.GetSummaryCmd(defaultDistroName)
	if err != nil {
		fmt.Println("error summary", err, cmd)
		return
	}

	fmt.Printf("Summary of default distro: %v\n%v", defaultDistroName, summaryOfDefaultDistro)
}
