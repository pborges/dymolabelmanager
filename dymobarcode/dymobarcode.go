package main
import (
	"os"
	"os/exec"
	"fmt"
	"github.com/pborges/dymolabelmanager"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:", os.Args[0], "<print|pdf> <data to be encoded>")
		return
	}

	pdf, err := dymolabelmanager.Barcode(os.Args[2])
	if err != nil {
		panic(err)
	}

	b, err := ioutil.TempFile("", "dymobarcode")
	if err != nil {
		panic(err)
	}
	b.Close()

	fileName := b.Name()

	if os.Args[1] == "pdf" {
		fileName = os.Args[2] + ".pdf"
	}

	err = pdf.OutputFileAndClose(fileName)
	if err != nil {
		panic(err)
	}

	if os.Args[1] == "print" {
		c := exec.Command("lpr", b.Name())
		err = c.Run()
		c.Wait()
		os.Remove(b.Name())
		if err != nil {
			panic(err)
		}
	}
}