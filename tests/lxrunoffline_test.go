package test

import (
	"fmt"
	"testing"

	"github.com/mkhuda/go-lxrunoffline"
)

func TestExportDistro(t *testing.T) {
	lxTest, err := lxrunoffline.New()
	if err != nil {
		fmt.Println(err)
	}

	err = lxTest.ExportDistro("Debian", "G:\\WSL_Backup\\go-lxrunoffline-test\\testbackupdebian.tar.gz")
	if err != nil {
		t.Errorf(err.Error())
	}
}
