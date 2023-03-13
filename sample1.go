package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type Reversiones struct {
	Cuenta        string
	Id            string
	MedioPago     string
	Monto         string
	PagoRevertido PagoRevertido
}
type PagoRevertido struct {
	IdPago            string
	PagoEstado        string
	ReferenciaExterna string
	Items             []ItemsRevertidos
	IntentoPago       IntentoPagoRevertido
}
type ItemsRevertidos struct {
	IdItems       string
	Cantidad      string
	Descripcion   string
	Monto         string
	Identificador string
}
type IntentoPagoRevertido struct {
	IdIntentoPago string
	IdTransaccion string
	FechaPago     string
	ImportePagado string
}

type reversion struct {
	m pdf.Maroto
}

type HeaderReversionData struct {
	Pago    PagoData
	Intento IntentoData
}

type PagoData struct {
	ReferenciaExterna string
	MedioPago         string
	Monto             string
	IdPago            string
	Estado            string
}
type IntentoData struct {
	IdIntento     string
	IdTransaccion string
	FechaPago     string
	Importe       string
}

type DataFooterRecaudacion struct {
	RecuperoComisiones string `json:"recupero_comisiones"`
	IvaRecupero        string `json:"iva_recupero"`
	Totales            string `json:"totales"`
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   51,
		Green: 51,
		Blue:  51,
	}
}

func (r *reversion) buildTitle() {

	fecha_impresion_pdf := fmt.Sprintf(time.Now().Format("02-01-2006"))

	r.m.RegisterHeader(func() {

		r.m.SetBorder(false)

		r.m.Line(1.0,
			props.Line{
				Color: color.Color{
					Red:   0,
					Green: 0,
					Blue:  255,
				},
			})

		r.m.Row(12, func() {

			r.m.Col(6, func() {
				r.m.Text("Informe de Reversiones de Pagos", props.Text{
					Top:   1,
					Align: consts.Center,
				})
			})

			r.m.Col(3, func() {
				r.m.Text("Fecha: "+fecha_impresion_pdf, props.Text{
					Size:   10,
					Style:  consts.BoldItalic,
					Top:    1,
					Family: consts.Helvetica,
				})
			})
		})

		r.m.Line(1.0,
			props.Line{
				Color: color.Color{
					Red:   0,
					Green: 0,
					Blue:  255,
				},
			})
	}) // Fin de RegisterHEader
}

func (r *reversion) buildHeadingsReversiones(data []HeaderReversionData) {

	for _, header := range data {

		// fila 1 Datos on Top
		r.m.Row(7, func() {

			r.m.SetBorder(true)

			// col1 Ref Externa
			r.m.Col(4, func() {

				r.m.Text("Referencia Externa: "+header.Pago.ReferenciaExterna, props.Text{
					// Top:    3,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
			})

			// col2 Medio de pago
			r.m.Col(4, func() {
				r.m.Text("Medio de Pago: "+header.Pago.MedioPago, props.Text{
					// Top:    3,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
			})

			// col3 Monto
			r.m.Col(4, func() {
				r.m.Text("Monto: "+header.Pago.Monto, props.Text{
					// Top:    3,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
			})
		})

		// fila 2 datos Pago e Intento
		r.m.Row(35, func() {

			// col 1 Pago
			r.m.Col(4, func() {

				r.m.Text("Pago", props.Text{

					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Align:  consts.Center,
					Color:  getDarkGrayColor(),
				})

				r.m.Text("Id: "+header.Pago.IdPago, props.Text{
					Top:    7,
					Left:   2,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

				r.m.Text("Estado: "+header.Pago.Estado, props.Text{
					Top:  14,
					Left: 2,

					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

			})

			// col 2 Intento
			r.m.Col(8, func() {

				r.m.Text("Intento", props.Text{
					Align:  consts.Center,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

				r.m.Text("Id Intento: "+header.Intento.IdIntento, props.Text{
					Top:    7,
					Left:   2,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

				r.m.Text("Transaccion: "+header.Intento.IdTransaccion, props.Text{
					Top:    14,
					Left:   2,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
				r.m.Text("Fecha: "+header.Intento.FechaPago, props.Text{
					Left:   2,
					Top:    21,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
				r.m.Text("Importe: "+header.Intento.Importe, props.Text{
					Left:   2,
					Top:    28,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

			})
		})

		// Agregar una nueva pagina para separar las reversiones
		r.m.AddPage()

	} // fin de for range

}

func getFakeReversiones() []Reversiones {
	// Open our jsonFile
	jsonFile, err := os.Open("response_reversiones.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var rever []Reversiones
	json.Unmarshal([]byte(byteValue), &rever)
	return rever
}
func getSmallContent() ([]string, [][]string) {
	header := []string{"Origin", "Destiny", "Cost"}

	contents := [][]string{}
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "R$ 20,00"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "R$ 25,00"})
	contents = append(contents, []string{"São José do Vale do Rio Preto", "Osasco", "R$ 20,00"})
	contents = append(contents, []string{"Osasco", "São Paulo", "R$ 5,00"})
	contents = append(contents, []string{"Congonhas", "Fortaleza", "R$ 100,00"})
	contents = append(contents, []string{"Natal", "Santo André", "R$ 200,00"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 44,00"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 56,00"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "R$ 35,00"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "R$ 82,00"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "R$ 62,00"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "R$ 21,00"})
	contents = append(contents, []string{"Natal", "Jundiaí", "R$ 12,00"})
	contents = append(contents, []string{"Niterói", "Itapevi", "R$ 21,00"})
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "R$ 31,00"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "R$ 42,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "R$ 19,00"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 42,00"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 58,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "R$ 21,00"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 56,00"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "R$ 35,00"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "R$ 82,00"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "R$ 62,00"})

	return header, contents
}

func getMediumContent() ([]string, [][]string) {
	header := []string{"Origin", "Destiny", "Cost per Hour"}

	contents := [][]string{}
	contents = append(contents, []string{"São José do Vale do Rio Preto", "Osasco", "R$ 12,00"})
	contents = append(contents, []string{"Niterói", "Itapevi", "R$ 2,10"})
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "R$ 3,10"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "R$ 4,20"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "R$ 1,90"})
	contents = append(contents, []string{"Osasco", "São Paulo", "R$ 0,70"})
	contents = append(contents, []string{"Congonhas", "Fortaleza", "R$ 11,30"})
	contents = append(contents, []string{"Natal", "Santo André", "R$ 19,80"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 4,20"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 5,80"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "R$ 3,90"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "R$ 2,30"})
	contents = append(contents, []string{"Natal", "Jundiaí", "R$ 1,10"})
	contents = append(contents, []string{"Natal", "Santo André", "R$ 19,80"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 4,20"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 5,80"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "R$ 3,90"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "R$ 7,70"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "R$ 6,40"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "R$ 2,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "R$ 1,90"})
	contents = append(contents, []string{"Osasco", "São Paulo", "R$ 0,70"})
	contents = append(contents, []string{"Congonhas", "São José do Vale do Rio Preto", "R$ 11,30"})
	contents = append(contents, []string{"Natal", "Santo André", "R$ 19,80"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 4,20"})

	return header, contents
}

func main() {
	begin := time.Now()
	reversionPdf := pdf.NewMaroto(consts.Portrait, consts.A4)
	reversionPdf.SetPageMargins(10, 15, 10)

	var rev reversion
	rev.m = reversionPdf

	// headerSmall, smallContent := getSmallContent()
	headerMedium, mediumContent := getMediumContent()

	// rev.m.SetAliasNbPages("algo")
	// rev.m.SetFirstPageNb(1)

	rev.buildTitle()

	//Obtener reversiones fake. reversiones es un slice de reportedtos.Reversiones
	reversiones := getFakeReversiones()
	var sliceHeaderData []HeaderReversionData

	// Parsear manualmente las reversiones a la data struct para presentar el pdf
	for _, oneRever := range reversiones {
		var header_data HeaderReversionData
		// Datos del pago en header
		header_data.Pago.ReferenciaExterna = oneRever.PagoRevertido.ReferenciaExterna
		header_data.Pago.MedioPago = oneRever.MedioPago
		header_data.Pago.Monto = oneRever.Monto
		header_data.Pago.IdPago = oneRever.PagoRevertido.IdPago
		header_data.Pago.Estado = oneRever.PagoRevertido.PagoEstado
		// Datos del intento de pago en herader
		header_data.Intento.IdIntento = oneRever.PagoRevertido.IntentoPago.IdIntentoPago
		header_data.Intento.IdTransaccion = oneRever.PagoRevertido.IntentoPago.IdTransaccion
		header_data.Intento.FechaPago = oneRever.PagoRevertido.IntentoPago.FechaPago
		header_data.Intento.Importe = oneRever.PagoRevertido.IntentoPago.ImportePagado

		sliceHeaderData = append(sliceHeaderData, header_data)
	}

	rev.buildHeadingsReversiones(sliceHeaderData)

	rev.m.AddPage()

	rev.m.Row(15, func() {
		rev.m.Col(12, func() {
			rev.m.Text(fmt.Sprintf("Medium Packages / %du.", len(mediumContent)), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	rev.m.TableList(headerMedium, mediumContent, props.TableList{
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Italic,
			GridSizes: []uint{5, 5, 2},
		},
		HeaderProp: props.TableListContent{
			GridSizes: []uint{5, 5, 2},
			Family:    consts.Courier,
			Style:     consts.BoldItalic,
			Color:     color.Color{100, 0, 0},
		},
		Align: consts.Center,
		Line:  true,
		LineProp: props.Line{
			Color: color.Color{
				Red:   128,
				Green: 221,
				Blue:  205,
			},
			Style: consts.Dashed,
		},
	})

	err := rev.m.OutputFileAndClose("pdfs/sample1.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
