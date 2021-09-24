package main

import (
	"encoding/json"
	"fmt"

	"github.com/mkhuda/go-lxrunoffline"
)

func main() {
	lx, err := lxrunoffline.New()
	if err != nil {
		fmt.Println(err)
	}

	listInstalled, err := lx.ListInstalled()
	if err != nil {
		fmt.Println("error listinstalled", err)
		return
	}

	defaultDistroName, defaultDistroUid, _ := lx.GetDefaultDistro()

	fmt.Println("List of installed WSL: ")
	for i, distros := range listInstalled {
		fmt.Println(i+1, distros.DistroName)
	}

	distroSummary, err := lx.GetDistroSummary(defaultDistroUid)
	if err != nil {
		fmt.Println("error summary", err)
		return
	}
	distroJson, err := json.Marshal(distroSummary)
	if err != nil {
		fmt.Println("error summary", err)
		return
	}

	fmt.Printf("Summary of default distro (marshalled): %v\n", string(distroJson))

	summaryOfDefaultDistro, cmd, err := lx.GetSummaryCmd(defaultDistroName)
	if err != nil {
		fmt.Println("error summary cmd", err, cmd)
		return
	}

	fmt.Printf("Summary of default distro: %v\n%v", defaultDistroName, summaryOfDefaultDistro)
}
