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

/* COLORES */
func getTelCoSoftBlueColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 184,
		Blue:  241,
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

/* OTRAS FUNCIONES */

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

func getHeaderAndContent() (header []string, contents [][]string) {
	header = []string{"Transacción", "Producto", "Cantidad", "Precio"}

	contents = [][]string{}

	contents = append(contents, []string{"389000786014502301031003026613", "30206226", "1", "3026.61"})
	contents = append(contents, []string{"389000786014502301031003012345", "12345678", "1", "1245.65"})

	return header, contents
}

/* FUNCIONES DE ARMADO PRINCIPAL */

func buildHeading(m pdf.Maroto) {
	green := getTelCoGreenColor()
	negro := getHeaderTextColor()
	blanco := color.NewWhite()

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
}

func buildBodyList(m pdf.Maroto) {
	celeste := getTelCoSoftBlueColor()
	blanco := color.NewWhite()
	header, contents := getHeaderAndContent()

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Verás este pago en tu resumen como TelCo Wee!!", props.Text{
				Top:   5,
				Style: consts.Italic,
				Align: consts.Center,
				Color: color.NewBlack(),
			})
		})
	})
	m.Line(5, props.Line{
		Color: blanco,
	})
	m.SetBackgroundColor(celeste)

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			GridSizes: []uint{3, 4, 2, 3},
			Size:      10,
			Style:     consts.Bold,
			Color:     blanco,
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 4, 2, 3},
		},
		Align:                consts.Center,
		HeaderContentSpace:   1,
		AlternatedBackground: &blanco,
		Line:                 true,
		LineProp: props.Line{
			Color: celeste,
		},
		VerticalContentPadding: 7, // alto de fila
	})
}

func buildFooter(m pdf.Maroto) {
	m.RegisterFooter(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/footer_recibo.png", props.Rect{
					Top: 30,
				})
				if err != nil {
					fmt.Println("la imagen no se pudo cargar", err)
				}
			})
		})
	})
}

func main() {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.SetPageMargins(0, 0, 0)

	buildHeading(m)

	m.SetPageMargins(10, 0, 10)

	buildBodyList(m)

	m.SetPageMargins(0, 0, 0)

	buildFooter(m)

	ahora := time.Now().Local().Unix()
	ahoraS := strconv.Itoa(int(ahora))
	err := m.OutputFileAndClose("pdfs/comprobante" + ahoraS + ".pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")

}
