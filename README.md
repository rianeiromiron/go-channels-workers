# Go Channels Workers

Este proyecto demuestra el uso de canales y workers en Go para procesar tareas concurrentemente.

## Descripción

Se crea un pool de workers que consume trabajos desde un canal, simula procesamiento con una pausa de 700ms y envía resultados a otro canal. La función principal espera a que todos los workers finalicen y luego imprime los resultados recolectados.

## Características

- Uso de goroutines
- Comunicación entre goroutines mediante canales
- Coordinación con sync.WaitGroup
- Manejo de producción y consumo de tareas
- Simulación de trabajo concurrente

## Requisitos

- Go 1.20 o superior

## Ejecución

```bash
go run .
```

## Ejemplo de salida

```text
🚀 Levantando 3 trabajadores concurrentes...

📊 [Main]: Esperando y recolectando resultados...
📥 Recibido en Main: [Trabajo #1] por [Worker 1] -> Procesado exitosamente 'Imagen-1.png'
...
🏁 ¡Todos los trabajos fueron procesados por el pool!
```

## Autor

rianeiromiron
