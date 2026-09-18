package main

import (
	"fmt"
	"sync"
	"time"
)

// Definimos la estructura de un Trabajo y su Resultado
type Trabajo struct {
	ID        int
	Contenido string
}

type Resultado struct {
	TrabajoID int
	Mensaje   string
	WorkerID  int
}

// Función Worker: Cada trabajador toma tareas del canal 'trabajos'
func worker(id int, trabajos <-chan Trabajo, resultados chan<- Resultado, wg *sync.WaitGroup) {
	defer wg.Done() // Al terminar la función, avisa al WaitGroup (-1)

	// El bucle lee tareas automáticamente hasta que el canal 'trabajos' se cierre
	for t := range trabajos {
		fmt.Printf("👷 Worker %d: INICIANDO trabajo #%d ('%s')\n", id, t.ID, t.Contenido)

		// Simula tiempo de procesamiento (entre 500ms y 1s)
		time.Sleep(700 * time.Millisecond)

		fmt.Printf("   ✅ Worker %d: COMPLETÓ trabajo #%d\n", id, t.ID)

		// Enviamos el resultado al canal de resultados
		resultados <- Resultado{
			TrabajoID: t.ID,
			Mensaje:   fmt.Sprintf("Procesado exitosamente '%s'", t.Contenido),
			WorkerID:  id,
		}
	}
}

func main() {
	const numTrabajos = 8
	const numWorkers = 3

	// Canales con búfer para los trabajos y resultados
	canalTrabajos := make(chan Trabajo, numTrabajos)
	canalResultados := make(chan Resultado, numTrabajos)

	// WaitGroup para coordinar la finalización de los 3 workers
	var wg sync.WaitGroup

	// 1. Levantamos a los 3 trabajadores en segundo plano
	fmt.Printf("🚀 Levantando %d trabajadores concurrentes...\n\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) // Sumamos 1 al contador por cada worker
		go worker(w, canalTrabajos, canalResultados, &wg)
	}

	// 2. Enviamos los 8 trabajos a la cola
	for j := 1; j <= numTrabajos; j++ {
		canalTrabajos <- Trabajo{
			ID:        j,
			Contenido: fmt.Sprintf("Imagen-%d.png", j),
		}
	}
	// Cerramos el canal de trabajos para avisar que ya no hay más tareas por asignar
	close(canalTrabajos)

	// 3. Lanzamos una goroutine para esperar que los workers terminen y cerrar resultados
	go func() {
		wg.Wait()              // Se bloquea hasta que el contador llegue a 0
		close(canalResultados) // Ya nadie más escribirá en resultados
	}()

	// 4. Recolectamos todos los resultados en 'main'
	fmt.Println("\n📊 [Main]: Esperando y recolectando resultados...")
	for res := range canalResultados {
		fmt.Printf("📥 Recibido en Main: [Trabajo #%d] por [Worker %d] -> %s\n",
			res.TrabajoID, res.WorkerID, res.Mensaje)
	}

	fmt.Println("\n🏁 ¡Todos los trabajos fueron procesados por el pool!")
}
