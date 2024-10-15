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
	Descripcion    string  
	Cantidad       int     
	PrecioUnitario float64 
}

var cliente string 

func main() { // Función principal del programa
	var items []Item                  // Declara un slice para almacenar los productos agregados por el usuario
	fmt.Print("Nombre del cliente: ") 
	fmt.Scanln(&cliente)              // Lee el nombre del cliente ingresado y lo almacena en la var

	reader := bufio.NewReader(os.Stdin) // Crea un nuevo lector para leer entradas del usuario lector eficiencia
	for {                               // Bucle para agregar productos hasta que el usuario decida salir
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
	fechaActual := time.Now().Format("02/01/2006 15:04:05") 

	pdf := gofpdf.New("P", "mm", "A4", "") // formato
	pdf.AddPage()                         
	pdf.SetFont("Arial", "", 14)        

	// Encabezado del PDF
	pdf.Cell(0, 10, "-----------------------------------------------------") 
	pdf.Ln(10)                                                            
	pdf.Cell(0, 10, "Factura")                                               
	pdf.Ln(10)                                                          
	pdf.Cell(0, 10, "Numero de Factura: 12345")                      
	pdf.Ln(10)                                                             
	pdf.Cell(0, 10, fmt.Sprintf("Fecha: %s", fechaActual))                   
	pdf.Ln(10)                                                               
	pdf.Cell(0, 10, fmt.Sprintf("Cliente: %s ", cliente))                    
	pdf.Ln(10)                                                               
	pdf.Cell(0, 10, "-----------------------------------------------------") 
	pdf.Ln(10)                                                               

	// Encabezado de la tabla de productos
	pdf.Cell(40, 10, "Descripcion")                                         
	pdf.Cell(30, 10, "Cantidad")                                        
	pdf.Cell(40, 10, "Precio Unitario")                                    
	pdf.Cell(40, 10, "Total")                                             
	pdf.Ln(10)                                                              
	pdf.Cell(0, 10, "-----------------------------------------------------") 
	pdf.Ln(10)                                                               

	// Agregar productos a la factura
	var totalFinal float64       // Variable para calcular el total final de la factura d a e
	for _, item := range items { // Recorre cada producto en la lista
		total := float64(item.Cantidad) * item.PrecioUnitario      
		totalFinal += total                                        
		pdf.Cell(40, 10, item.Descripcion)                        
		pdf.Cell(30, 10, strconv.Itoa(item.Cantidad))              
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", item.PrecioUnitario)) // Muestra el precio unitario del producto
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", total))               // Muestra el total por ítem
		pdf.Ln(10)                                                
	}

	// Mostrar el total final de la factura
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
