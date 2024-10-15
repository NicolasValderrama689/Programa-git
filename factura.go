package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Estructura para almacenar los ítems
type Item struct {
	Descripcion    string
	Cantidad       int
	PrecioUnitario float64
}

var cliente string

func main() {
	var items []Item
	fmt.Print("Nombre del cliente: ")
	fmt.Scanln(&cliente)
	// Entrada del usuario para agregar productos
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("¿Desea añadir un producto a la factura? (s/n)")
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(confirm)
		if confirm == "n" {
			break
		}
		// Leer descripción del producto
		fmt.Print("Descripción del producto: ")
		descripcion, _ := reader.ReadString('\n')
		descripcion = strings.TrimSpace(descripcion)
		// Leer cantidad
		fmt.Print("Cantidad: ")
		cantidadStr, _ := reader.ReadString('\n')
		cantidadStr = strings.TrimSpace(cantidadStr)
		cantidad, _ := strconv.Atoi(cantidadStr)
		// Leer precio unitario
		fmt.Print("Precio unitario: ")
		precioStr, _ := reader.ReadString('\n')
		precioStr = strings.TrimSpace(precioStr)
		precioUnitario, _ := strconv.ParseFloat(precioStr, 64)
		// Agregar el producto a la lista de items
		item := Item{
			Descripcion:    descripcion,
			Cantidad:       cantidad,
			PrecioUnitario: precioUnitario,
		}
		items = append(items, item)
	}
	// Generar el PDF de la factura
	generarFacturaPDF(items)
}
func generarFacturaPDF(items []Item) {
	// Crear un nuevo PDF
	pdf := gofpdf.New("P", "mm", "A4", "") // Inicializa el PDF
	pdf.AddPage()
	// Establecer fuente
	pdf.SetFont("Arial", "", 14)
	// Mostrar detalles de la factura
	pdf.Cell(0, 10, "-----------------------------------------------------")
	pdf.Ln(10)
	pdf.Cell(0, 10, "Factura")
	pdf.Ln(10)
	pdf.Cell(0, 10, "Numero de Factura: 12345")
	pdf.Ln(10)
	pdf.Cell(0, 10, "Fecha: 01/10/2024")
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Cliente: %s ", cliente))
	pdf.Ln(10)
	pdf.Cell(0, 10, "-----------------------------------------------------")
	pdf.Ln(10)
	pdf.Cell(40, 10, "Descripcion")
	pdf.Cell(30, 10, "Cantidad")
	pdf.Cell(40, 10, "Precio Unitario")
	pdf.Cell(40, 10, "Total")
	pdf.Ln(10)
	pdf.Cell(0, 10, "-----------------------------------------------------")
	pdf.Ln(10)
	// Agregar los ítems y calcular total
	var totalFinal float64
	for _, item := range items {
		total := float64(item.Cantidad) * item.PrecioUnitario
		totalFinal += total
		pdf.Cell(40, 10, item.Descripcion)
		pdf.Cell(30, 10, strconv.Itoa(item.Cantidad))
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", item.PrecioUnitario))
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", total))
		pdf.Ln(10)
	}
	// Total final de la factura
	pdf.Cell(0, 10, "-----------------------------------------------------")
	pdf.Ln(10)
	pdf.Cell(40, 10, "Total Final")
	pdf.Cell(30, 10, "")
	pdf.Cell(40, 10, "")
	pdf.Cell(40, 10, fmt.Sprintf("%.2f", totalFinal))
	pdf.Ln(10)
	pdf.Cell(0, 10, "-----------------------------------------------------")
	// Guardar el PDF
	err := pdf.OutputFileAndClose("factura.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
