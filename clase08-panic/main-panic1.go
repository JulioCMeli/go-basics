/*
Ejercicio 1 - Datos de clientes
Un estudio contable necesita acceder a los datos de sus empleados para poder
realizar distintas liquidaciones. Para ello, cuentan con todo el detalle necesario
en un archivo .txt.

Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt que nos
indica el cliente, sin embargo, no han pasado el archivo a leer por nuestro programa.
Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt”
(recordá lo visto sobre el pkg “os”).
Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso,
el programa deberá arrojar un panic al intentar leer un archivo que no existe,
mostrando el mensaje “The indicated file was not found or is damaged”.

Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.

*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	//Esto se va a ejecutar a pesar de un panic
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("ejecución finalizada")
	}()

	// open a file
	var fileName = "customers.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("The indicated file was not found or is damaged")
		return
	}
	fmt.Println(file.Name())

}
