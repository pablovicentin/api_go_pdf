package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/johnfercher/maroto/pkg/color"
// 	"github.com/johnfercher/maroto/pkg/consts"
// 	"github.com/johnfercher/maroto/pkg/pdf"
// 	"github.com/johnfercher/maroto/pkg/props"
// )

// type Reversiones struct {
// 	Cuenta        string
// 	Id            string
// 	MedioPago     string
// 	Monto         string
// 	PagoRevertido PagoRevertido
// }
// type PagoRevertido struct {
// 	IdPago            string
// 	PagoEstado        string
// 	ReferenciaExterna string
// 	Items             []ItemsRevertidos
// 	IntentoPago       IntentoPagoRevertido
// }
// type ItemsRevertidos struct {
// 	IdItems       string
// 	Cantidad      string
// 	Descripcion   string
// 	Monto         string
// 	Identificador string
// }
// type IntentoPagoRevertido struct {
// 	IdIntentoPago string
// 	IdTransaccion string
// 	FechaPago     string
// 	ImportePagado string
// }

// type reversion struct {
// 	m pdf.Maroto
// }

// type ReversionData struct {
// 	Pago    PagoData
// 	Intento IntentoData
// 	Items   []ItemsRevertidos
// }

// type ItemsReversionData struct {
// 	Cantidad      string
// 	Descripcion   string
// 	Identificador string
// 	Monto         string
// }

// type PagoData struct {
// 	ReferenciaExterna string
// 	MedioPago         string
// 	Monto             string
// 	IdPago            string
// 	Estado            string
// }

// type IntentoData struct {
// 	IdIntento     string
// 	IdTransaccion string
// 	FechaPago     string
// 	Importe       string
// }

// type DataFooterRecaudacion struct {
// 	RecuperoComisiones string `json:"recupero_comisiones"`
// 	IvaRecupero        string `json:"iva_recupero"`
// 	Totales            string `json:"totales"`
// }

// func getDarkGrayColor() color.Color {
// 	return color.Color{
// 		Red:   51,
// 		Green: 51,
// 		Blue:  51,
// 	}
// }

// func (r *reversion) buildTitle() {

// 	fecha_impresion_pdf := fmt.Sprintf(time.Now().Format("02-01-2006"))

// 	r.m.RegisterHeader(func() {

// 		r.m.SetBorder(false)

// 		r.m.Line(1.0,
// 			props.Line{
// 				Color: color.Color{
// 					Red:   0,
// 					Green: 0,
// 					Blue:  255,
// 				},
// 			})

// 		r.m.Row(12, func() {

// 			r.m.Col(6, func() {
// 				r.m.Text("Reversiones de Pagos", props.Text{
// 					Top:   1,
// 					Align: consts.Center,
// 					Size:  12,
// 				})
// 			})

// 			r.m.Col(3, func() {
// 				r.m.Text("Fecha: "+fecha_impresion_pdf, props.Text{
// 					Size:   10,
// 					Style:  consts.BoldItalic,
// 					Top:    1,
// 					Family: consts.Helvetica,
// 				})
// 			})
// 		})

// 		r.m.Line(1.0,
// 			props.Line{
// 				Color: color.Color{
// 					Red:   0,
// 					Green: 0,
// 					Blue:  255,
// 				},
// 			})
// 	}) // Fin de RegisterHEader
// }

// func (r *reversion) buildHeadingsReversiones(data []ReversionData) {

// 	for _, dato := range data {

// 		// fila 1 Datos on Top
// 		r.m.Row(7, func() {

// 			r.m.SetBorder(true)

// 			// col1 Ref Externa
// 			r.m.Col(4, func() {

// 				r.m.Text("Referencia Externa: "+dato.Pago.ReferenciaExterna, props.Text{
// 					// Top:    3,
// 					Left:   2,
// 					Style:  consts.Bold,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})
// 			})

// 			// col2 Medio de pago
// 			r.m.Col(4, func() {
// 				r.m.Text("Medio de Pago: "+dato.Pago.MedioPago, props.Text{
// 					// Top:    3,
// 					Left:   2,
// 					Style:  consts.Bold,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})
// 			})

// 			// col3 Monto
// 			r.m.Col(4, func() {
// 				r.m.Text("Monto: "+dato.Pago.Monto, props.Text{
// 					// Top:    3,
// 					Left:   2,
// 					Style:  consts.Bold,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})
// 			})
// 		})

// 		// fila 2 datos Pago e Intento
// 		r.m.Row(35, func() {

// 			// col 1 Pago
// 			r.m.Col(4, func() {

// 				r.m.Text("Pago", props.Text{

// 					Style:  consts.Bold,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Align:  consts.Center,
// 					Color:  getDarkGrayColor(),
// 				})

// 				r.m.Text("Id: "+dato.Pago.IdPago, props.Text{
// 					Top:    7,
// 					Left:   2,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})

// 				r.m.Text("Estado: "+dato.Pago.Estado, props.Text{
// 					Top:  14,
// 					Left: 2,

// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})

// 			})

// 			// col 2 Intento
// 			r.m.Col(8, func() {

// 				r.m.Text("Intento", props.Text{
// 					Align:  consts.Center,
// 					Style:  consts.Bold,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})

// 				r.m.Text("Id Intento: "+dato.Intento.IdIntento, props.Text{
// 					Top:    7,
// 					Left:   2,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})

// 				r.m.Text("Transaccion: "+dato.Intento.IdTransaccion, props.Text{
// 					Top:    14,
// 					Left:   2,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})
// 				r.m.Text("Fecha: "+dato.Intento.FechaPago, props.Text{
// 					Left:   2,
// 					Top:    21,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})
// 				r.m.Text("Importe: "+dato.Intento.Importe, props.Text{
// 					Left:   2,
// 					Top:    28,
// 					Family: consts.Courier,
// 					Size:   10,
// 					Color:  getDarkGrayColor(),
// 				})

// 			})
// 		})

// 		r.m.SetBorder(false)

// 		// Mostrar los items de cada pago revertido
// 		buildBodyItems(r.m, dato.Items)

// 		// Agregar una nueva pagina para separar las reversiones
// 		r.m.AddPage()

// 	} // fin de for range

// }

// func buildBodyItems(m pdf.Maroto, items []ItemsRevertidos) {
// 	header, contents := getMediumContent(items)
// 	m.Line(1)
// 	m.Row(7, func() {
// 		m.Col(12, func() {
// 			m.Text("Items del Pago", props.Text{Align: consts.Left, Left: 2, Size: 10, Top: 2})
// 		})
// 	})
// 	m.Line(1)

// 	m.TableList(header, contents, props.TableList{
// 		ContentProp: props.TableListContent{
// 			Family:    consts.Courier,
// 			Style:     consts.Italic,
// 			GridSizes: []uint{2, 4, 4, 2},
// 		},
// 		HeaderProp: props.TableListContent{
// 			GridSizes: []uint{2, 4, 4, 2},
// 			Family:    consts.Courier,
// 			Style:     consts.BoldItalic,
// 			Color:     color.Color{100, 0, 0},
// 		},
// 		Line: true,
// 		LineProp: props.Line{
// 			Color: color.Color{
// 				Red:   128,
// 				Green: 221,
// 				Blue:  205,
// 			},
// 			Style: consts.Dashed,
// 		},
// 	})
// }

// func (r *reversion) buildFooter() {
// 	r.m.SetFirstPageNb(1)
// 	r.m.RegisterFooter(func() {
// 		r.m.Row(5, func() {
// 			r.m.Col(12, func() {
// 				r.m.Text(strconv.Itoa(r.m.GetCurrentPage()), props.Text{
// 					Align: consts.Right,
// 					Size:  8,
// 					Top:   10,
// 				})
// 			})
// 		})
// 	})
// }

// func getFakeReversiones() []Reversiones {
// 	// Open our jsonFile
// 	jsonFile, err := os.Open("response_reversiones.json")
// 	// if we os.Open returns an error then handle it
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// defer the closing of our jsonFile so that we can parse it later on
// 	defer jsonFile.Close()

// 	byteValue, _ := ioutil.ReadAll(jsonFile)

// 	var rever []Reversiones
// 	json.Unmarshal([]byte(byteValue), &rever)
// 	return rever
// }

// func getMediumContent(items []ItemsRevertidos) ([]string, [][]string) {
// 	header := []string{"Cantidad", "Descripcion", "Identificador", "Monto"}

// 	contents := [][]string{}

// 	for _, item := range items {
// 		contents = append(contents, []string{item.Cantidad, item.Descripcion, item.Identificador, item.Monto})
// 	}

// 	return header, contents
// }

// func main2() {

// 	reversionPdf := pdf.NewMaroto(consts.Portrait, consts.A4)
// 	reversionPdf.SetPageMargins(10, 15, 10)

// 	var rev reversion
// 	rev.m = reversionPdf

// 	rev.buildTitle()

// 	//Obtener reversiones fake. reversiones es un slice de reportedtos.Reversiones
// 	reversiones := getFakeReversiones()
// 	var sliceHeaderData []ReversionData

// 	// Parsear manualmente las reversiones a la data struct para presentar el pdf
// 	for _, oneRever := range reversiones {
// 		var data ReversionData
// 		// Datos del pago en header
// 		data.Pago.ReferenciaExterna = oneRever.PagoRevertido.ReferenciaExterna
// 		data.Pago.MedioPago = oneRever.MedioPago
// 		data.Pago.Monto = oneRever.Monto
// 		data.Pago.IdPago = oneRever.PagoRevertido.IdPago
// 		data.Pago.Estado = oneRever.PagoRevertido.PagoEstado
// 		// Datos del intento de pago en herader
// 		data.Intento.IdIntento = oneRever.PagoRevertido.IntentoPago.IdIntentoPago
// 		data.Intento.IdTransaccion = oneRever.PagoRevertido.IntentoPago.IdTransaccion
// 		fecha := strings.Split(oneRever.PagoRevertido.IntentoPago.FechaPago, " ")
// 		data.Intento.FechaPago = fecha[0]
// 		data.Intento.Importe = oneRever.PagoRevertido.IntentoPago.ImportePagado
// 		data.Items = oneRever.PagoRevertido.Items
// 		sliceHeaderData = append(sliceHeaderData, data)
// 	}

// 	// registrar el footer
// 	rev.buildFooter()
// 	// Las cabeceras de cada pago revertidos y los items de cada pago
// 	rev.buildHeadingsReversiones(sliceHeaderData)

// 	err := rev.m.OutputFileAndClose("pdfs/reversiones.pdf")
// 	if err != nil {
// 		fmt.Println("Could not save PDF:", err)
// 		os.Exit(1)
// 	}

// }
