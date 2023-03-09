// Billing example
package main

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type Actor struct {
	RazonSocial    string `json:"razonSocial"`
	Domicilio      string `json:"domicilio"`
	Cuit           string `json:"cuit"`
	IngresosBrutos string `json:"ingresosBrutos"`
	Iva            string `json:"iva"`
}

type Fechas struct {
	Cobro    string `json:"cobro"`
	Deposito string `json:"deposito"`
	Proceso  string `json:"proceso"`
}

type DataBody struct {
	ChannelName       string `json:"channel_name"`
	ImporteCobrado    string `json:"importe_cobrado"`
	ImporteDepositado string `json:"importe_depositado"`
	CantidadBoletas   string `json:"cantidad_boletas"`
	Comisiones        string `json:"comisiones"`
	IvaComision       string `json:"iva_comision"`
	RetIva            string `json:"ret_iva"`
}
type DataFooter struct {
	RecuperoComisiones string `json:"recupero_comisiones"`
	IvaRecupero        string `json:"iva_recupero"`
	Totales            string `json:"totales"`
}

func buildHeading(m pdf.Maroto, emisor, receptor Actor, fechas Fechas, fileName string) {
	m.RegisterHeader(func() {

		// fila 1: nombre del archivo
		m.Row(10, func() {

			// columna 1
			m.Col(12, func() {
				m.Text("Recaudación WEE!", props.Text{
					Top:    3,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   12,
					Align:  consts.Center,
					Color:  getDarkGrayColor(),
				})
			})
		})

		m.SetBorder(true)

		// fila 1 datos emisor y receptor
		m.Row(60, func() {

			// col 1 emisor
			m.Col(6, func() {

				m.Text("Entidad Cobradora", props.Text{
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   12,
					Color:  getDarkGrayColor(),
				})

				m.Text(emisor.RazonSocial, props.Text{
					Top:    10,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

				m.Text("Domicilio: "+emisor.Domicilio, props.Text{
					Top:    20,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
				m.Text("C.U.I.T.: "+emisor.Cuit, props.Text{
					Left:   2,
					Top:    30,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
				m.Text("Ing. Brutos: "+emisor.IngresosBrutos, props.Text{
					Left:   2,
					Top:    40,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
				m.Text("I.V.A.: "+emisor.Iva, props.Text{
					Left:   2,
					Top:    50,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

			})

			// col 2 receptor
			m.Col(6, func() {

				m.Text("Empresa", props.Text{
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   12,
					Color:  getDarkGrayColor(),
				})

				m.Text(receptor.RazonSocial, props.Text{
					Top:    10,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

				m.Text("Domicilio: "+receptor.Domicilio, props.Text{
					Top:    20,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
				m.Text("C.U.I.T.: "+receptor.Cuit, props.Text{
					Left:   2,
					Top:    30,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
				m.Text("Ing. Brutos: "+receptor.IngresosBrutos, props.Text{
					Left:   2,
					Top:    40,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
				m.Text("I.V.A.: "+receptor.Iva, props.Text{
					Left:   2,
					Top:    50,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})

			})
		})

		m.SetBorder(false)

		// fila 2 fechas
		m.Row(10, func() {

			// col1 fecha de cobro
			m.Col(4, func() {

				m.Text("Fecha Cobro: "+fechas.Cobro, props.Text{
					Top:    3,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
			})

			// col2 fecha de deposito
			m.Col(4, func() {
				m.Text("Fecha Depósito: "+fechas.Deposito, props.Text{
					Top:    3,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
			})

			// col3 fecha de proceso
			m.Col(4, func() {
				m.Text("Fechas Proceso: "+fechas.Proceso, props.Text{
					Top:    3,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
			})
		})

		m.Row(10, func() {
			m.Col(12, func() {
				m.Text("Nombre de Archivo Transferido: "+fileName, props.Text{
					Top:    3,
					Left:   2,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   10,
					Color:  getDarkGrayColor(),
				})
			})
		})

	})
}

func buildContent(m pdf.Maroto, contents []DataBody) {

	// tableHeadings := []string{"Fruit", "Description", "Price"}
	// contents := [][]string{{"Apple", "Red and juicy", "2.00"}, {"Orange", "Orange and juicy", "3.00"}}
	for _, data := range contents {

		m.Line(2)
		m.Row(5, func() {
			m.Col(12, func() {
				m.Text(data.ChannelName, props.Text{Align: consts.Left, Left: 5})
			})
		})
		m.Line(2)

		// m.SetBorder(true)

		m.Row(7, func() {
			m.Col(3, func() {
				m.Text("Importe Cobrado: ", props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
			m.Col(9, func() {
				m.Text(data.ImporteCobrado, props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
		})
		m.Row(7, func() {
			m.Col(3, func() {
				m.Text("Importe Depositado: ", props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
			m.Col(9, func() {
				m.Text(data.ImporteDepositado, props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
		})
		m.Row(7, func() {
			m.Col(3, func() {
				m.Text("Cantidad Boletas: ", props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
			m.Col(9, func() {
				m.Text(data.CantidadBoletas, props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
		})
		m.Row(7, func() {
			m.Col(3, func() {
				m.Text("Comisiones: ", props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
			m.Col(9, func() {
				m.Text(data.Comisiones, props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
		})
		m.Row(7, func() {
			m.Col(3, func() {
				m.Text("IVA Comisión: ", props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
			m.Col(9, func() {
				m.Text(data.IvaComision, props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
		})
		m.Row(7, func() {
			m.Col(3, func() {
				m.Text("Ret IVA 3130: ", props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
			m.Col(9, func() {
				m.Text(data.RetIva, props.Text{Align: consts.Left, Left: 2, Size: 8})
			})
		})

		// m.Row(10, func() { m.Text("") })
	}
}

func buildFooter(m pdf.Maroto, dataFooter DataFooter) {

	m.RegisterFooter(func() {

		m.Row(7, func() {
			m.Col(12, func() {
				m.Text("Recupero Comisiones Bancarias: "+dataFooter.RecuperoComisiones, props.Text{Align: consts.Left, Left: 2, Size: 8, Top: 13})
			})
		})
		m.Row(7, func() {
			m.Col(12, func() {
				m.Text("IVA Recupero Com Bancarias: "+dataFooter.IvaRecupero, props.Text{Align: consts.Left, Left: 2, Size: 8, Top: 16})
			})
		})
		m.Row(7, func() {
			m.Col(12, func() {
				m.Text("Totales: "+dataFooter.Totales, props.Text{Align: consts.Left, Left: 2, Size: 8, Top: 19})
			})
		})

	})
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   51,
		Green: 51,
		Blue:  51,
	}
}

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	/************* Cabecera  *************/
	var emisor, receptor Actor
	var fechas Fechas

	emisor.RazonSocial = "Corrientes Telecomunicaciones SAPEM"
	emisor.Domicilio = "Dr. R. Carillo 444"
	emisor.Cuit = " 30716550849"
	emisor.IngresosBrutos = ""
	emisor.Iva = ""

	receptor.RazonSocial = "Direccion Prov de Energia de Corrientes"
	receptor.Domicilio = "Junin 1240"
	receptor.Cuit = " 30608090181"
	receptor.IngresosBrutos = ""
	receptor.Iva = "Responsable Inscripto"

	fechas.Cobro = "24/11/22"
	fechas.Deposito = "29/11/22"
	fechas.Proceso = "25/11/22"

	fileName := "archivo de rendicion"

	// encabezdo
	buildHeading(m, emisor, receptor, fechas, fileName)

	/************* Cuerpo *************/
	var data, data2 DataBody
	var contents []DataBody

	data.ChannelName = "Canal de Pago 1"
	data.ImporteCobrado = "100"
	data.ImporteDepositado = "100"
	data.CantidadBoletas = "10"
	data.Comisiones = "5"
	data.IvaComision = "6"
	data.RetIva = "3"

	data2.ChannelName = "Canal de Pago 2"
	data2.ImporteCobrado = "100"
	data2.ImporteDepositado = "100"
	data2.CantidadBoletas = "10"
	data2.Comisiones = "5"
	data2.IvaComision = "6"
	data2.RetIva = "3"

	contents = append(contents, data, data2)

	// contenido
	buildContent(m, contents)

	m.Line(1)

	/************* Footer  *************/

	var dataFooter DataFooter
	dataFooter.RecuperoComisiones = "200"
	dataFooter.IvaRecupero = "21"
	dataFooter.Totales = "221"

	// pie de informe
	buildFooter(m, dataFooter)

	// crear archivo en la carpeta pdfs
	err := m.OutputFileAndClose("pdfs/div_rhino_fruit.pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}
