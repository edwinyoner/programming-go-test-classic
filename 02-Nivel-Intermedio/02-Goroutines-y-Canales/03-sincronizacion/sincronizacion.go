/*
Archivo: sincronizacion.go
Autor: Edwin Yoner
Fecha: 26/10/2024
Versión: 1.0

Este archivo muestra cómo sincronizar goroutines usando WaitGroup.
*/

package main

import (
    "fmt"
    "sync"
)

func tarea(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Marca la tarea como completada al finalizar.
    fmt.Printf("Tarea %d completada\n", id)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go tarea(i, &wg)
    }
    wg.Wait() // Espera a que todas las goroutines terminen.
    fmt.Println("Todas las tareas completadas")
}
