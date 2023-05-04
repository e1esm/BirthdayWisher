package utils

import (
	"fmt"
	"github.com/signintech/gopdf"
	"go.uber.org/zap"
	"pdfGenerator/internal/models"
	"pdfGenerator/utils"
	"time"
)

var widthCenter float64
var headerLength float64 = 400
var widthConstraint float64 = 200
var heightConstraint float64 = 25
var fontHeaderSize float64 = 24
var fontMainSize float64 = 16
var regularFont string = "times"
var boldFont string = "times-bold"

func NewPDF(users []models.User) {
	pdf := reportHeadline()
	pdf = newTable(pdf, users)
}

func reportHeadline() *gopdf.GoPdf {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	widthCenter = gopdf.PageSizeA4.W / 2
	pdf.AddPage()
	err := pdf.AddTTFFont(regularFont, regularFont+":ttf")
	if err != nil {
		utils.Logger.Panic("Couldn't have load the font", zap.String("font", regularFont))
	}
	err = pdf.AddTTFFont(boldFont, boldFont+":ttf")
	if err != nil {
		utils.Logger.Panic("Couldn't have load the bold font", zap.String("font", boldFont))
	}

	pdf.SetFont(boldFont, "", fontHeaderSize)
	headOfFile := fmt.Sprintf("Chat ID: %d - Birthdays", 123)
	linelength, _ := pdf.MeasureTextWidth(headOfFile)

	pdf.SetX(widthCenter - linelength/2)
	pdf.Cell(nil, headOfFile)
	pdf.Br(32)

	pdf.SetFont(regularFont, "", fontMainSize)
	date := fmt.Sprintf("%s", time.Now().Format(time.DateTime))

	pdf.Cell(nil, date)
	pdf.Br(32)
	return newTableHeader(pdf)
}

func newTableHeader(pdf *gopdf.GoPdf) *gopdf.GoPdf {

	pdf.SetX(widthCenter - headerLength/2)
	pdf.CellWithOption(&gopdf.Rect{H: heightConstraint, W: widthConstraint}, "Users", gopdf.CellOption{Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
		Align: gopdf.Middle | gopdf.Center})
	pdf.CellWithOption(&gopdf.Rect{H: heightConstraint, W: widthConstraint}, "Dates", gopdf.CellOption{Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
		Align: gopdf.Middle | gopdf.Center})
	pdf.Br(32)
	return pdf
}

func newTable(pdf *gopdf.GoPdf, users []models.User) *gopdf.GoPdf {

	pdf.SetX(widthCenter - headerLength/2)
	fmt.Print(users)
	for i := 0; i < 10; i++ {
		for _, v := range users {
			if pdf.GetY() > gopdf.PageSizeA4.W+200 {
				pdf.AddPage()
			}
			pdf.SetX(widthCenter - headerLength/2)
			pdf.CellWithOption(&gopdf.Rect{H: heightConstraint, W: widthConstraint}, v.Username, gopdf.CellOption{Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
				Align: gopdf.Middle | gopdf.Center})
			pdf.CellWithOption(&gopdf.Rect{H: heightConstraint, W: widthConstraint}, v.Date[:10], gopdf.CellOption{Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
				Align: gopdf.Middle | gopdf.Center})
			pdf.Br(25)
		}
	}

	return pdf
}
