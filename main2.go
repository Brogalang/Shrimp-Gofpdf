package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	pdf := gofpdf.New("p", "mm", "A4", "")
	pdf.AddPage()

	// Font files and their identifiers
	fontFiles := map[string]string{
		"OpenSansRegular": filepath.Join("OpenSans-Regular.ttf"),
		"OpenSansItalic":  filepath.Join("OpenSans-Italic.ttf"),
		"OpenSansBold":    filepath.Join("OpenSans-Bold.ttf"),
	}

	// Read and add the TTF fonts from files
	for fontName, fontPath := range fontFiles {
		fontBytes, err := ioutil.ReadFile(fontPath)
		if err != nil {
			log.Fatalf("Failed to read font file %s: %v", fontPath, err)
		}
		if len(fontBytes) == 0 {
			log.Fatalf("Font file %s is empty or could not be read correctly", fontPath)
		}
		pdf.AddUTF8FontFromBytes(fontName, "", fontBytes)
		if pdf.Error() != nil {
			log.Fatalf("Error adding font %s: %v", fontPath, pdf.Error())
		}
	}

	// Tambahkan gambar
	imgX := 0.0
	imgY := -25.0
	imgWidth := 120.0
	pdf.Image("logo_shrimp.png", imgX, imgY, imgWidth, 0, false, "", 0, "")

	pdf.SetTextColor(226, 32, 40) // RGB untuk #E22028
	textX := 13.0
	textY := imgY + 95 // Atur Y untuk teks di bawah gambar
	pdf.SetXY(textX, textY)
	pdf.SetFont("OpenSansBold", "", 35)
	pdf.CellFormat(0, 10, "Report PCR Aquamarine", "", 0, "L", false, 0, "")

	pdf.Ln(15)                       // Jarak tambahan di bawah teks
	pdf.SetDrawColor(226, 32, 40)    // RGB: untuk #E22028
	lineY := pdf.GetY()              // Ambil posisi Y saat ini
	pdf.Line(-10, lineY, 162, lineY) // Gambar garis di bawah teks

	pdf.SetTextColor(226, 32, 40) // RGB untuk #E22028
	text2X := textX
	text3Y := textY + 20
	pdf.SetXY(text2X, text3Y)
	pdf.SetFont("OpenSansBold", "", 15)
	namaTambak := strings.ToUpper("Tambak Aquamarine")
	pdf.CellFormat(0, 10, "Nama Tambak: "+fmt.Sprintf("%s", namaTambak), "", 0, "L", false, 0, "")

	imgFtX := 1.0
	imgFtY := 116.0
	imgFtWidth := 300.0
	pdf.Image("icon_shrimp.png", imgFtX, imgFtY, imgFtWidth, 0, false, "", 0, "")

	// for page := 0; page < 2; page++ {
	// ## Halaman kedua ##
	pdf.AddPage()
	imgPg2X := 5.0
	imgPg2Y := -20.0
	imgPg2Width := 80.0
	pdf.Image("logo_shrimp.png", imgPg2X, imgPg2Y, imgPg2Width, 0, false, "", 0, "")

	pdf.Ln(15)
	imgH2X := 10.0
	imgH2Y := 0.0
	imgH2Width := 200.0
	pdf.Image("header_shrimp.png", imgH2X, imgH2Y, imgH2Width, 0, false, "", 0, "")

	pdf.SetTextColor(226, 32, 40) // RGB untuk #E22028
	textN1X := 13.0
	textN1Y := imgPg2Y + 55 // Atur Y untuk teks di bawah gambar
	pdf.SetXY(textN1X, textN1Y)
	pdf.SetFont("OpenSansBold", "", 18)
	pdf.CellFormat(0, 10, "Rangkuman Hasil PCR", "", 0, "L", false, 0, "")

	pdf.Ln(6)
	pdf.SetX(textN1X)
	pdf.SetFont("OpenSansItalic", "", 12)
	pdf.CellFormat(0, 10, "PCR Report Summary", "", 0, "L", false, 0, "")

	pdf.SetTextColor(0, 0, 0) // Set warna hitam
	// Baris pertama
	pdf.SetXY(textN1X, 55)
	pdf.SetFont("OpenSansRegular", "", 10)
	pdf.CellFormat(35, 1, "Nama Perusahaan", "", 0, "L", false, 0, "")
	pdf.CellFormat(2, 2, ":", "0", 0, "L", false, 0, "")
	pdf.CellFormat(0, 2, "GSI LAB", "0", 1, "L", false, 0, "")

	pdf.SetXY(120, 55)
	pdf.SetFont("OpenSansRegular", "", 10)
	pdf.CellFormat(40, 1, "Tanggal Pengambilan", "", 0, "L", false, 0, "")
	pdf.CellFormat(2, 2, ":", "0", 0, "L", false, 0, "")
	pdf.CellFormat(0, 2, "16-07-2024", "0", 1, "L", false, 0, "")

	pdf.SetXY(textN1X, 56.5)
	pdf.SetFont("OpenSansItalic", "", 8)
	pdf.CellFormat(0, 6, "Company Name", "", 0, "L", false, 0, "")

	pdf.SetXY(120, 56.5)
	pdf.SetFont("OpenSansItalic", "", 8)
	pdf.CellFormat(0, 6, "Date Collected", "", 0, "L", false, 0, "")

	// Baris kedua
	pdf.SetXY(textN1X, 65)
	pdf.SetFont("OpenSansRegular", "", 10)
	pdf.CellFormat(35, 1, "Alamat E-mail", "", 0, "L", false, 0, "")
	pdf.CellFormat(2, 2, ":", "0", 0, "L", false, 0, "")
	pdf.CellFormat(0, 2, "Galang.wibowo@gsilab.id", "0", 1, "L", false, 0, "")

	pdf.SetXY(120, 65)
	pdf.SetFont("OpenSansRegular", "", 10)
	pdf.CellFormat(40, 1, "Tanggal Diterima", "", 0, "L", false, 0, "")
	pdf.CellFormat(2, 2, ":", "0", 0, "L", false, 0, "")
	pdf.CellFormat(0, 2, "17-07-2024", "0", 1, "L", false, 0, "")

	pdf.SetXY(textN1X, 66.5)
	pdf.SetFont("OpenSansItalic", "", 8)
	pdf.CellFormat(0, 6, "E-mail Address", "", 0, "L", false, 0, "")

	pdf.SetXY(120, 66.5)
	pdf.SetFont("OpenSansItalic", "", 8)
	pdf.CellFormat(0, 6, "Date Received", "", 0, "L", false, 0, "")

	// Baris ketiga
	pdf.SetXY(textN1X, 75)
	pdf.SetFont("OpenSansRegular", "", 10)
	pdf.CellFormat(35, 1, "ID Booking", "", 0, "L", false, 0, "")
	pdf.CellFormat(2, 2, ":", "0", 0, "L", false, 0, "")
	pdf.CellFormat(0, 2, "ITB-976-289-9185", "0", 1, "L", false, 0, "")

	pdf.SetXY(120, 75)
	pdf.SetFont("OpenSansRegular", "", 10)
	pdf.CellFormat(40, 1, "Tanggal Dilaporkan", "", 0, "L", false, 0, "")
	pdf.CellFormat(2, 2, ":", "0", 0, "L", false, 0, "")
	pdf.CellFormat(0, 2, "19-07-2024", "0", 1, "L", false, 0, "")

	pdf.SetXY(textN1X, 76.5)
	pdf.SetFont("OpenSansItalic", "", 8)
	pdf.CellFormat(0, 6, "Booking ID", "", 0, "L", false, 0, "")

	pdf.SetXY(120, 76.5)
	pdf.SetFont("OpenSansItalic", "", 8)
	pdf.CellFormat(0, 6, "Date Reported", "", 0, "L", false, 0, "")

	pdf.SetTextColor(226, 32, 40) // RGB untuk #E22028
	textN2X := 13.0
	textN2Y := 85.0
	pdf.SetXY(textN2X, textN2Y)
	pdf.SetFont("OpenSansBold", "", 18)
	pdf.CellFormat(0, 10, "Hasil Pemeriksaan", "", 0, "L", false, 0, "")

	pdf.Ln(6)
	pdf.SetX(textN2X)
	pdf.SetFont("OpenSansItalic", "", 12)
	pdf.CellFormat(0, 10, "Test Result", "", 0, "L", false, 0, "")

	// Tabel
	pdf.SetTextColor(255, 255, 255) // RGB untuk Putih
	pdf.SetDrawColor(226, 32, 40)   // Warna merah (RGB: 226, 32, 40)
	// Atur ketebalan garis menjadi 1 px
	pdf.SetLineWidth(0.3)
	pdf.SetFillColor(226, 32, 40) // Warna abu-abu (RGB: 200, 200, 200)

	// Margin atau jarak antar border
	margin := 1.0
	cellWidth := 32.0
	cellWidthMini := 12.0
	cellHeightH := 20.0
	cellHeight := 10.0
	cellHeighCol := 8.0

	pdf.SetFont("OpenSansRegular", "", 8)
	pdf.SetXY(13, 103)
	pdf.CellFormat(cellWidth-margin*1.5, cellHeightH-margin*2, "Specimen ID", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+cellWidth, 103)
	pdf.CellFormat(cellWidth-margin*1.5, cellHeightH-margin*2, "Sample Name", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+2*cellWidth, 103)
	pdf.CellFormat(cellWidth-margin*1.5, cellHeightH-margin*2, "Sampel Source", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+3*cellWidth, 103)
	pdf.CellFormat(cellWidth-margin*1.5, cellHeightH-margin*2, "Sampel Type", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+4*cellWidth, 103)
	pdf.CellFormat(cellWidth+28-margin*1.5, cellHeight+1-margin*2, "CT", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+4*cellWidth, 113)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "WSSV", "1", 0, "C", true, 0, "")
	pdf.SetXY(25+4*cellWidth, 113)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "IMNV", "1", 0, "C", true, 0, "")
	pdf.SetXY(37+4*cellWidth, 113)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "EHP", "1", 0, "C", true, 0, "")
	pdf.SetXY(49+4*cellWidth, 113)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "AHPND", "1", 0, "C", true, 0, "")
	pdf.SetXY(61+4*cellWidth, 113)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "IHHNV", "1", 0, "C", true, 0, "")
	pdf.Ln(4)

	// Menambahkan padding antar sel dengan menggambar border manual
	pdf.SetTextColor(0, 0, 0) // RGB untuk hitam
	yPos := pdf.GetY() + margin + 5
	for row := 0; row < 17; row++ { // termasuk header
		xPos := 13.0
		for col := 0; col < 9; col++ {
			pdf.SetXY(xPos, yPos)
			if col > 3 {
				pdf.CellFormat(cellWidthMini-margin*1.5, cellHeighCol-margin*2, fmt.Sprintf("%.2f", yPos), "1", 0, "C", false, 0, "")
				xPos += cellWidthMini
			} else if col <= 3 {
				pdf.CellFormat(cellWidth-margin*1.5, cellHeighCol-margin*2, fmt.Sprintf("%.2f", yPos), "1", 0, "C", false, 0, "")
				xPos += cellWidth
			}

		}
		yPos += cellHeighCol
	}

	// Deskripsi Resiko
	pdf.SetTextColor(0, 0, 0) // RGB untuk hitam
	pdf.SetFont("OpenSansRegular", "", 10.5)
	pdf.SetXY(13, 260)
	pdf.CellFormat(0, 10, "Deskripsi Resiko :", "", 0, "L", false, 0, "")
	pdf.SetFont("OpenSansItalic", "", 7.5)
	pdf.SetXY(13, 265)
	pdf.CellFormat(0, 10, "Risk Description", "", 0, "L", false, 0, "")

	imgIconRed2X := 58.0
	imgIconRed2Y := 260.5
	imgIconRed2Width := 14.5
	pdf.Image("mini_red_shrimp.png", imgIconRed2X, imgIconRed2Y, imgIconRed2Width, 0, false, "", 0, "")

	pdf.SetXY(70, 260.5)
	pdf.SetFont("OpenSansRegular", "", 10.5)
	pdf.CellFormat(0, 10, "Positif", "", 0, "L", false, 0, "")

	pdf.SetXY(70, 265)
	pdf.SetFont("OpenSansItalic", "", 7.5)
	pdf.CellFormat(0, 10, "Positive", "", 0, "L", false, 0, "")

	imgIconGreen2X := 100.0
	imgIconGreen2Y := 260.5
	imgIconGreen2Width := 12.5
	pdf.Image("mini_green_shrimp.png", imgIconGreen2X, imgIconGreen2Y, imgIconGreen2Width, 0, false, "", 0, "")

	pdf.SetXY(112, 260.5)
	pdf.SetFont("OpenSansRegular", "", 10.5)
	pdf.CellFormat(0, 10, "Negatif", "", 0, "L", false, 0, "")

	pdf.SetXY(112, 265)
	pdf.SetFont("OpenSansItalic", "", 7.5)
	pdf.CellFormat(0, 10, "Negative", "", 0, "L", false, 0, "")

	imgIconBlack2X := 145.0
	imgIconBlack2Y := 260.5
	imgIconBlack2Width := 13.5
	pdf.Image("black_shrimp.png", imgIconBlack2X, imgIconBlack2Y, imgIconBlack2Width, 0, false, "", 0, "")

	pdf.SetXY(160, 260.5)
	pdf.SetFont("OpenSansRegular", "", 10.5)
	pdf.CellFormat(0, 10, "Tidak Diperiksa", "", 0, "L", false, 0, "")

	pdf.SetXY(160, 265)
	pdf.SetFont("OpenSansItalic", "", 7.5)
	pdf.CellFormat(0, 10, "Not Examined", "", 0, "L", false, 0, "")

	// Dapatkan dimensi halaman
	_, pageHeightN := pdf.GetPageSize()
	imgWidthN := 350.0        // Lebar gambar
	xPosN := 0.0              // Posisi X (pojok kiri)
	yPosN := pageHeightN - 12 // Posisi Y (pojok bawah)
	pdf.Image("footer_shrimp.png", xPosN, yPosN, imgWidthN, 0, false, "", 0, "")
	// }

	// ## Halaman ketiga ##
	pdf.AddPage()
	imgPg2X = 5.0
	imgPg2Y = -20.0
	imgPg2Width = 80.0
	pdf.Image("logo_shrimp.png", imgPg2X, imgPg2Y, imgPg2Width, 0, false, "", 0, "")

	pdf.Ln(15)
	imgH2X = 10.0
	imgH2Y = 0.0
	imgH2Width = 200.0
	pdf.Image("header_shrimp.png", imgH2X, imgH2Y, imgH2Width, 0, false, "", 0, "")

	pdf.SetTextColor(226, 32, 40) // RGB untuk #E22028
	textN3X := 13.0
	textN3Y := 35.0
	pdf.SetXY(textN3X, textN3Y)
	pdf.SetFont("OpenSansBold", "", 18)
	pdf.CellFormat(0, 10, "Hasil Pemeriksaan", "", 0, "L", false, 0, "")

	pdf.Ln(6)
	pdf.SetX(textN3X)
	pdf.SetFont("OpenSansItalic", "", 12)
	pdf.CellFormat(0, 10, "Test Result", "", 0, "L", false, 0, "")

	// Tabel
	pdf.SetTextColor(255, 255, 255) // RGB untuk Putih
	pdf.SetDrawColor(226, 32, 40)   // Warna merah (RGB: 226, 32, 40)
	// Atur ketebalan garis menjadi 1 px
	pdf.SetLineWidth(0.3)
	pdf.SetFillColor(226, 32, 40) // Warna abu-abu (RGB: 200, 200, 200)

	// Margin atau jarak antar border
	margin = 1.0
	cellWidth = 32.0
	cellWidthMini = 12.0
	cellHeightH = 20.0
	cellHeight = 10.0
	cellHeighCol = 8.0

	pdf.SetFont("OpenSansRegular", "", 8)
	pdf.SetXY(13, 53)
	pdf.CellFormat(cellWidth-margin*1.5, cellHeightH-margin*2, "Specimen ID", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+cellWidth, 53)
	pdf.CellFormat(cellWidth-margin*1.5, cellHeightH-margin*2, "Sample Name", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+2*cellWidth, 53)
	pdf.CellFormat(cellWidth-margin*1.5, cellHeightH-margin*2, "Sampel Source", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+3*cellWidth, 53)
	pdf.CellFormat(cellWidth-margin*1.5, cellHeightH-margin*2, "Sampel Type", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+4*cellWidth, 53)
	pdf.CellFormat(cellWidth+28-margin*1.5, cellHeight+1-margin*2, "CT", "1", 0, "C", true, 0, "")
	pdf.SetXY(13+4*cellWidth, 63)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "WSSV", "1", 0, "C", true, 0, "")
	pdf.SetXY(25+4*cellWidth, 63)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "IMNV", "1", 0, "C", true, 0, "")
	pdf.SetXY(37+4*cellWidth, 63)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "EHP", "1", 0, "C", true, 0, "")
	pdf.SetXY(49+4*cellWidth, 63)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "AHPND", "1", 0, "C", true, 0, "")
	pdf.SetXY(61+4*cellWidth, 63)
	pdf.CellFormat(cellWidthMini-margin*1.5, cellHeight-margin*2, "IHHNV", "1", 0, "C", true, 0, "")
	pdf.Ln(4)

	// Menambahkan padding antar sel dengan menggambar border manual
	pdf.SetTextColor(0, 0, 0) // RGB untuk hitam
	yPos = pdf.GetY() + margin + 5
	for row := 0; row < 16; row++ { // termasuk header
		xPos := 13.0
		for col := 0; col < 9; col++ {
			pdf.SetXY(xPos, yPos)
			if col > 3 {
				pdf.CellFormat(cellWidthMini-margin*1.5, cellHeighCol-margin*2, fmt.Sprintf("%.2f", yPos), "1", 0, "C", false, 0, "")
				xPos += cellWidthMini
			} else if col <= 3 {
				pdf.CellFormat(cellWidth-margin*1.5, cellHeighCol-margin*2, fmt.Sprintf("%.2f", yPos), "1", 0, "C", false, 0, "")
				xPos += cellWidth
			}

		}
		yPos += cellHeighCol
	}

	// Deskripsi Resiko
	pdf.SetTextColor(0, 0, 0) // RGB untuk hitam
	pdf.SetFont("OpenSansRegular", "", 10.5)
	pdf.SetXY(13, 203)
	pdf.CellFormat(0, 10, "Deskripsi Resiko :", "", 0, "L", false, 0, "")
	pdf.SetFont("OpenSansItalic", "", 7.5)
	pdf.SetXY(13, 207)
	pdf.CellFormat(0, 10, "Risk Description", "", 0, "L", false, 0, "")

	imgIconRedX := 58.0
	imgIconRedY := 203.5
	imgIconRedWidth := 14.5
	pdf.Image("mini_red_shrimp.png", imgIconRedX, imgIconRedY, imgIconRedWidth, 0, false, "", 0, "")

	pdf.SetXY(70, 203.5)
	pdf.SetFont("OpenSansRegular", "", 10.5)
	pdf.CellFormat(0, 10, "Positif", "", 0, "L", false, 0, "")

	pdf.SetXY(70, 207)
	pdf.SetFont("OpenSansItalic", "", 7.5)
	pdf.CellFormat(0, 10, "Positive", "", 0, "L", false, 0, "")

	imgIconGreenX := 100.0
	imgIconGreenY := 203.5
	imgIconGreenWidth := 12.5
	pdf.Image("mini_green_shrimp.png", imgIconGreenX, imgIconGreenY, imgIconGreenWidth, 0, false, "", 0, "")

	pdf.SetXY(112, 203.5)
	pdf.SetFont("OpenSansRegular", "", 10.5)
	pdf.CellFormat(0, 10, "Negatif", "", 0, "L", false, 0, "")

	pdf.SetXY(112, 207)
	pdf.SetFont("OpenSansItalic", "", 7.5)
	pdf.CellFormat(0, 10, "Negative", "", 0, "L", false, 0, "")

	imgIconBlackX := 145.0
	imgIconBlackY := 203.5
	imgIconBlackWidth := 13.5
	pdf.Image("black_shrimp.png", imgIconBlackX, imgIconBlackY, imgIconBlackWidth, 0, false, "", 0, "")

	pdf.SetXY(160, 203.5)
	pdf.SetFont("OpenSansRegular", "", 10.5)
	pdf.CellFormat(0, 10, "Tidak Diperiksa", "", 0, "L", false, 0, "")

	pdf.SetXY(160, 207)
	pdf.SetFont("OpenSansItalic", "", 7.5)
	pdf.CellFormat(0, 10, "Not Examined", "", 0, "L", false, 0, "")

	// Catatan
	pdf.SetXY(13, 218)
	pdf.SetFont("OpenSansRegular", "", 7.5)
	pdf.CellFormat(0, 10, "Catatan / Note", "", 0, "L", false, 0, "")
	pdf.SetFont("OpenSansRegular", "", 7.5)
	pdf.SetXY(13, 226)
	pdf.CellFormat(5, 1, "1.", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(18, 225)
	pdf.MultiCell(0, 3, "Pemeriksaan ini hanya menggambarkan kondisi saat pengambilan specimen. The above examination only \ndescribes the condition during the specimen collection.", "0", "L", false)
	pdf.SetXY(13, 232)
	pdf.CellFormat(5, 1, "2.", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(18, 231)
	pdf.MultiCell(0, 3, "Pemeriksaan yang dilakukan merupakan pemeriksaan yang mendeteksi materi genetic patogen, tanpa \nmembedakan viabilitas patogen tersebut (hidup atau mati). The above examination only detects genetic \nmaterials of pathogens, without the ability to distinguish viability of pathogens.", "0", "L", false)
	pdf.SetXY(13, 241)
	pdf.CellFormat(5, 1, "3.", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(18, 240)
	pdf.MultiCell(0, 3, "Hasil qRT-PCR serta nilai Ct tidak dapat dibandingkan antara satu pemeriksaan dengan pemeriksaan lain; \nmengingat adanya variasi nilai Ct untuk alat, metode, reagen yang digunakan serta gen target yang \ndiperiksa. qRT-PCR result nor Ct value should not be compared between PCR test since a wide variety of \nCt values have been reported due to different tools, methods, reagents, and targeted genes used.", "0", "L", false)
	pdf.SetXY(13, 253)
	pdf.CellFormat(5, 1, "4.", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(18, 252)
	pdf.MultiCell(0, 3, "Nilai cut-off Ct masing-masing pemeriksaan:", "0", "L", false)

	pdf.SetXY(18, 256.5)
	pdf.SetFont("OpenSansRegular", "", 15)
	pdf.CellFormat(5, 1, ".", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(22, 257.7)
	pdf.SetFont("OpenSansRegular", "", 7.5)
	pdf.CellFormat(0, 1, "WSSV <= 41.5", "0", 1, "L", false, 0, "")

	pdf.SetXY(45, 256.5)
	pdf.SetFont("OpenSansRegular", "", 15)
	pdf.CellFormat(5, 1, ".", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(49, 257.7)
	pdf.SetFont("OpenSansRegular", "", 7.5)
	pdf.CellFormat(0, 1, "AHPND <= 41", "0", 1, "L", false, 0, "")

	pdf.SetXY(18, 259.5)
	pdf.SetFont("OpenSansRegular", "", 15)
	pdf.CellFormat(5, 1, ".", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(22, 260.7)
	pdf.SetFont("OpenSansRegular", "", 7.5)
	pdf.CellFormat(0, 1, "IMNV <= 41", "0", 1, "L", false, 0, "")

	pdf.SetXY(45, 259.5)
	pdf.SetFont("OpenSansRegular", "", 15)
	pdf.CellFormat(5, 1, ".", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(49, 260.7)
	pdf.SetFont("OpenSansRegular", "", 7.5)
	pdf.CellFormat(0, 1, "IHHNV <= 39.09", "0", 1, "L", false, 0, "")

	pdf.SetXY(18, 262.5)
	pdf.SetFont("OpenSansRegular", "", 15)
	pdf.CellFormat(5, 1, ".", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 1, "", "0", 0, "L", false, 0, "")
	pdf.SetXY(22, 263.7)
	pdf.SetFont("OpenSansRegular", "", 7.5)
	pdf.CellFormat(0, 1, "EHP <= 42", "0", 1, "L", false, 0, "")

	pdf.Ln(10)                                 // Jarak tambahan di bawah teks
	pdf.SetDrawColor(0, 0, 0)                  // RGB: untuk hita,
	lineBlackY := pdf.GetY()                   // Ambil posisi Y saat ini
	pdf.Line(200, lineBlackY, 160, lineBlackY) // Gambar garis di bawah teks

	// Dapatkan dimensi halaman
	_, pageHeightN = pdf.GetPageSize()
	imgWidthN = 350.0        // Lebar gambar
	xPosN = 0.0              // Posisi X (pojok kiri)
	yPosN = pageHeightN - 12 // Posisi Y (pojok bawah)
	pdf.Image("footer_shrimp.png", xPosN, yPosN, imgWidthN, 0, false, "", 0, "")

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}
