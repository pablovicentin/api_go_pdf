package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/color"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// func getHeader() []string {
// 	return []string{"", "Producto", "Cantidad", "Precio"}
// }

// func formatFechaString(fecha time.Time, formatoFecha string) string {
// 	fechaStr := fecha.Format(formatoFecha)
// 	fechaArrayStr := strings.Split(fechaStr[0:10], "-")
// 	fechaVto := fmt.Sprintf("%v-%v-%v", fechaArrayStr[2], fechaArrayStr[1], fechaArrayStr[0])
// 	return fechaVto
// }

// func getContents(pagoItems *[]entities.Pagoitems) [][]string {
// 	size := len(*pagoItems)
// 	items := make([][]string, size)
// 	for i, x := range *pagoItems {
// 		identificador := ""

// 		if x.Identifier != "" {
// 			identificador = x.Identifier
// 		}

// 		items[i] = []string{
// 			identificador, x.Description, fmt.Sprint(x.Quantity), strconv.FormatFloat(x.Amount.Float64(), 'f', 2, 64),
// 		}

// 		// items[i] = []string{
// 		// 	"", x.Description, fmt.Sprint(x.Quantity), strconv.FormatFloat(x.Amount.Float64(), 'f', 2, 64),
// 		// }
// 	}
// 	return items
// }

// func getDarkGrayColor() color.Color {
// 	return color.Color{
// 		Red:   55,
// 		Green: 55,
// 		Blue:  55,
// 	}
// }

// func getGrayColor() color.Color {
// 	return color.Color{
// 		Red:   200,
// 		Green: 200,
// 		Blue:  200,
// 	}
// }

// func getBlackColor() color.Color {
// 	return color.Color{
// 		Red:   0,
// 		Green: 0,
// 		Blue:  0,
// 	}
// }
// func getColorWeePrimary() color.Color {
// 	return color.Color{
// 		Red:   000,
// 		Green: 156,
// 		Blue:  222,
// 	}
// }

// func getColorWeeSecondary() color.Color {
// 	return color.Color{
// 		Red:   149,
// 		Green: 201,
// 		Blue:  223,
// 	}
// }

func buildHeading(m pdf.Maroto) {
	green := getTelCoGreenColor()
	negro := getHeaderTextColor()
	blanco := color.NewWhite()

	fmt.Print(green)
	// RegisterHeader
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/cabecera_recibo.png", props.Rect{})

				if err != nil {
					fmt.Println("la imagen no se pudo cargar", err)
				}
			})
		})
	})
	// texto:= "MUNICIP. DE SANTO TOME CTES - CUIT 30546676427"
	texto := "MUNICIPALIDAD DE VIRASO - CUIT 30546676427"
	// colSize, colSpaceSize := _resolveColumnWidthSize("TELCO SAPEM - CUIT 30546676427")
	colSize, colSpaceSize := _resolveColumnWidthSize(texto)

	m.Row(10, func() {
		m.ColSpace(uint(colSpaceSize))
		m.SetBackgroundColor(green)
		m.Col(uint(colSize), func() {
			m.Text(texto, props.Text{
				Style: consts.Normal,
				Color: negro,
				Align: consts.Right,
				Size:  8,
				Top:   3,
				Right: 2,
			})
		})
		m.SetBackgroundColor(blanco)

	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Verás este pago en tu resumen como TelCo Wee!!", props.Text{
				Top:   3,
				Style: consts.Italic,
				Align: consts.Center,
				Color: negro,
			})
		})
	})
}

/* COLORES */
func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getWhiteColor() color.Color {
	return color.NewWhite()
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getHeaderTextColor() color.Color {
	return color.NewBlack()
}

func getTelCoGreenColor() color.Color {
	return color.Color{
		Red:   195,
		Green: 216,
		Blue:  46,
	}
}

// size dinamico segun long del texto que se recibe
func _resolveColumnWidthSize(texto string) (colSize, colSpaceSize int) {
	long := len(texto)
	var columMaxSize int = 12
	switch {
	case long <= 30:
		colSize = 3
		colSpaceSize = columMaxSize - colSize
	case long > 30 && long < 43:
		colSize = 4
		colSpaceSize = columMaxSize - colSize
	case long >= 43:
		colSize = 5
		colSpaceSize = columMaxSize - colSize
	default:
		colSize = 4
		colSpaceSize = columMaxSize - colSize
	}
	return
}

func main() {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(0, 0, 0)

	buildHeading(m)
	// buildFruitList(m)

	ahora := time.Now().Local().Unix()
	ahoraS := strconv.Itoa(int(ahora))
	err := m.OutputFileAndClose("pdfs/comprobante" + ahoraS + ".pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")

}
