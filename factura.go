package main // Define el paquete principal del programa

import ( // Importación de bibliotecas
	"bufio"   // Para leer entradas del usuario de manera eficiente
	"fmt"     // Para imprimir y formatear texto
	"log"     // Para manejar errores y registrar mensajes
	"os"      // Para acceder a las funciones del sistema operativo, como la entrada/salida estándar
	"strconv" // Para convertir cadenas a números y viceversa
	"strings" // Para manipulación de cadenas, como eliminación de espacios en blanco
	"time"    // Para manejar fechas y horas

	"github.com/jung-kurt/gofpdf" // Biblioteca para generar PDFs en Go
)

// Estructura para almacenar los ítems
type Item struct { // Define una estructura para representar un ítem en la factura
	Descripcion    string  // Nombre o descripción del producto
	Cantidad       int     // Cantidad del producto
	PrecioUnitario float64 // Precio por unidad del producto 
}

var cliente string // Variable global para almacenar el nombre del cliente

func main() { // Función principal del programa
	var items []Item                  // Declara un slice para almacenar los productos agregados por el usuario
	fmt.Print("Nombre del cliente: ") // Solicita el nombre del cliente
	fmt.Scanln(&cliente)              // Lee el nombre del cliente ingresado y lo almacena en la variable 'cliente'

	reader := bufio.NewReader(os.Stdin) // Crea un nuevo lector para leer entradas del usuario lector eficiencia
	for {                               // Bucle para agregar productos hasta que el usuario decida salir
		fmt.Println("¿Desea añadir un producto a la factura? (s/n)") // Pregunta al usuario si quiere agregar un producto
		confirm, _ := reader.ReadString('\n')                        // Lee la respuesta del usuario
		confirm = strings.TrimSpace(confirm)                         // Elimina espacios en blanco adicionales de la respuesta
		if confirm == "n" {                                          // Si el usuario responde "n", se rompe el bucle
			break
		}

		// Leer descripción del producto
		fmt.Print("Descripción del producto: ")      // Solicita la descripción del producto
		descripcion, _ := reader.ReadString('\n')    // Lee la descripción del producto ingresada por el usuario
		descripcion = strings.TrimSpace(descripcion) // Elimina espacios en blanco de la descripción

		// Leer cantidad
		fmt.Print("Cantidad: ")                      // Solicita la cantidad del producto
		cantidadStr, _ := reader.ReadString('\n')    // Lee la cantidad ingresada como cadena
		cantidadStr = strings.TrimSpace(cantidadStr) // Elimina espacios en blanco de la cantidad
		cantidad, _ := strconv.Atoi(cantidadStr)     // Convierte la cantidad de cadena a entero

		// Leer precio unitario
		fmt.Print("Precio unitario: ")                         // Solicita el precio unitario del producto
		precioStr, _ := reader.ReadString('\n')                // Lee el precio unitario ingresado como cadena
		precioStr = strings.TrimSpace(precioStr)               // Elimina espacios en blanco del precio
		precioUnitario, _ := strconv.ParseFloat(precioStr, 64) // Convierte el precio unitario de cadena a float64

		// Agregar el producto a la lista de items
		item := Item{ // Crea un nuevo ítem con la descripción, cantidad y precio ingresados
			Descripcion:    descripcion,
			Cantidad:       cantidad,
			PrecioUnitario: precioUnitario,
		}
		items = append(items, item) // Añade el ítem a la lista de productos
	}

	generarFacturaPDF(items) // Llama a la función para generar el PDF con los productos ingresados
}

func generarFacturaPDF(items []Item) { // Función para generar el PDF de la factura
	fechaActual := time.Now().Format("02/01/2006 15:04:05") // Obtiene la fecha y hora actual en el formato especificado

	pdf := gofpdf.New("P", "mm", "A4", "") // Inicializa un nuevo documento PDF en formato A4
	pdf.AddPage()                          // Añade una nueva página al PDF
	pdf.SetFont("Arial", "", 14)           // Establece la fuente del texto en Arial, tamaño 14

	// Encabezado del PDF
	pdf.Cell(0, 10, "-----------------------------------------------------") // Dibuja una línea divisoria
	pdf.Ln(10)                                                               // Salta una línea en el PDF
	pdf.Cell(0, 10, "Factura")                                               // Título del documento
	pdf.Ln(10)                                                               // Salta una línea
	pdf.Cell(0, 10, "Numero de Factura: 12345")                              // Muestra un número de factura fijo
	pdf.Ln(10)                                                               // Salta una línea
	pdf.Cell(0, 10, fmt.Sprintf("Fecha: %s", fechaActual))                   // Agrega la fecha y hora actuales al PDF
	pdf.Ln(10)                                                               // Salta una línea
	pdf.Cell(0, 10, fmt.Sprintf("Cliente: %s ", cliente))                    // Muestra el nombre del cliente en el PDF
	pdf.Ln(10)                                                               // Salta una línea
	pdf.Cell(0, 10, "-----------------------------------------------------") // Dibuja otra línea divisoria
	pdf.Ln(10)                                                               // Salta una línea

	// Encabezado de la tabla de productos
	pdf.Cell(40, 10, "Descripcion")                                          // Columna de descripción del producto
	pdf.Cell(30, 10, "Cantidad")                                             // Columna de cantidad
	pdf.Cell(40, 10, "Precio Unitario")                                      // Columna de precio unitario
	pdf.Cell(40, 10, "Total")                                                // Columna de total por ítem
	pdf.Ln(10)                                                               // Salta una línea
	pdf.Cell(0, 10, "-----------------------------------------------------") // Dibuja una línea divisoria
	pdf.Ln(10)                                                               // Salta una línea

	// Agregar productos a la factura
	var totalFinal float64       // Variable para calcular el total final de la factura d a e
	for _, item := range items { // Recorre cada producto en la lista
		total := float64(item.Cantidad) * item.PrecioUnitario      // Calcula el total por ítem
		totalFinal += total                                        // Suma el total del ítem al total final
		pdf.Cell(40, 10, item.Descripcion)                         // Muestra la descripción del producto
		pdf.Cell(30, 10, strconv.Itoa(item.Cantidad))              // Muestra la cantidad del producto
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", item.PrecioUnitario)) // Muestra el precio unitario del producto
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", total))               // Muestra el total por ítem
		pdf.Ln(10)                                                 // Salta una línea
	}

	// Mostrar el total final de la factura
	pdf.Cell(0, 10, "-----------------------------------------------------") // Dibuja una línea divisoria
	pdf.Ln(10)                                                               // Salta una línea
	pdf.Cell(40, 10, "Total Final")                                          // Muestra la etiqueta de total final
	pdf.Cell(30, 10, "")                                                     // Espacio vacío
	pdf.Cell(40, 10, "")                                                     // Espacio vacío
	pdf.Cell(40, 10, fmt.Sprintf("%.2f", totalFinal))                        // Muestra el total final calculado
	pdf.Ln(10)                                                               // Salta una línea
	pdf.Cell(0, 10, "-----------------------------------------------------") // Dibuja una línea divisoria

	// Guardar el PDF
	err := pdf.OutputFileAndClose("factura.pdf") // Guarda el PDF con el nombre "factura.pdf"
	if err != nil {                              // Verifica si hubo un error al guardar el PDF
		log.Fatal(err) // Muestra un mensaje de error y detiene el programa si ocurre un problema
	}
}
