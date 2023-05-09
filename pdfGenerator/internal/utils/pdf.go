package utils

import (
	"fmt"
	"github.com/signintech/gopdf"
	"go.uber.org/zap"
	"image/png"
	"os"
	"pdfGenerator/internal/models"
	"sync"

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

func NewPDF(users []models.User, chatID int64) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go GenerateAllImages(chatID, users, wg)
	pdf := reportHeadline(chatID)
	pdf = newTable(pdf, users, chatID)
}

func reportHeadline(chatID int64) *gopdf.GoPdf {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	widthCenter = gopdf.PageSizeA4.W / 2
	pdf.AddPage()
	err := pdf.AddTTFFont(regularFont, regularFont+":ttf")
	if err != nil {

		Logger.Panic("Couldn't have load the font", zap.String("font", regularFont))
	}
	err = pdf.AddTTFFont(boldFont, boldFont+":ttf")
	if err != nil {
		Logger.Panic("Couldn't have load the bold font", zap.String("font", boldFont))
	}

	pdf.SetFont(boldFont, "", fontHeaderSize)
	headOfFile := fmt.Sprintf("Chat ID: %d - Birthdays", chatID)
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

func newTable(pdf *gopdf.GoPdf, users []models.User, chatID int64) *gopdf.GoPdf {

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

	return newImages(pdf, chatID)
}

func newImages(pdf *gopdf.GoPdf, chatID int64) *gopdf.GoPdf {

	gap := 275
	agesFile, err := os.Open(fmt.Sprintf("pie-ages-%d.png", chatID))
	if err != nil {
		panic(err)
	}
	defer agesFile.Close()
	agesImg, err := png.Decode(agesFile)
	if err != nil {
		panic(err)
	}
	pageWidth := gopdf.PageSizeA4.W
	agesWidth := float64(agesImg.Bounds().Max.X - agesImg.Bounds().Min.X)
	centerX := (pageWidth-agesWidth)/2 + 225
	pdf.ImageFrom(agesImg, centerX, 0, &gopdf.Rect{H: gopdf.PageSizeA4.H / 3, W: gopdf.PageSizeA4.W})

	monthsFile, err := os.Open(fmt.Sprintf("pie-months-%d.png", chatID))
	if err != nil {
		panic(err)
	}
	defer monthsFile.Close()
	monthsImg, err := png.Decode(monthsFile)
	if err != nil {
		panic(err)
	}
	pdf.ImageFrom(monthsImg, centerX, float64(gap), &gopdf.Rect{H: gopdf.PageSizeA4.H / 3, W: gopdf.PageSizeA4.W})

	yearsFIle, err := os.Open(fmt.Sprintf("pie-years-%d.png", chatID))
	if err != nil {
		panic(err)
	}
	defer yearsFIle.Close()
	yearsImg, err := png.Decode(yearsFIle)
	if err != nil {
		panic(err)
	}
	pdf.ImageFrom(yearsImg, centerX, float64(gap*2), &gopdf.Rect{H: gopdf.PageSizeA4.H / 3, W: gopdf.PageSizeA4.W})

	return pdf
}
