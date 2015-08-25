package dymolabelmanager
import (
	"github.com/jung-kurt/gofpdf"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func Barcode(data string) (pdf *gofpdf.Fpdf, err error) {
	b, err := code128.Encode(data)
	if err != nil { return }

	b, err = barcode.Scale(b, b.Bounds().Size().X, 10)
	if err != nil {
		panic(err)
	}

	out, err := ioutil.TempFile("", "lblmgrpnpbarcode")
	if err != nil {
		panic(err)
	}
	var opt jpeg.Options

	opt.Quality = 100
	err = jpeg.Encode(out, b, &opt)
	if err != nil { return }

	imgHeight := float64(b.Bounds().Size().Y)
	imgWidth := float64(b.Bounds().Size().X / 3)

	pdf = gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:    "mm",
		Size:       gofpdf.SizeType{Wd: imgWidth, Ht: imgHeight},
	})

	pdf.AddPage()

	height, width := pdf.GetPageSize()
	pdf.Image(out.Name(), 0, 0, height, width, false, "jpg", 0, "")

	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(0, 7, imgWidth, 5, "F")
	pdf.SetFont("Arial", "B", 6)
	pdf.Text(0, 9.1, data)
	out.Close()
	os.Remove(out.Name())
	return
}