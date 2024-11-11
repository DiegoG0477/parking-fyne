# Simulador de Estacionamiento

Este repositorio contiene un **simulador concurrente de estacionamiento** implementado en Go. El simulador modela el comportamiento de vehículos que entran y salen de un estacionamiento con un número limitado de espacios. La simulación incluye una visualización en consola y en interfaz gráfica utilizando la biblioteca GUI [Fyne](https://fyne.io/).

## Tabla de Contenidos
- [Descripción General](#descripción-general)
- [Arquitectura](#arquitectura)
- [Instalación](#instalación)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Desarrollo](#desarrollo)

## Descripción General

El Simulador de Estacionamiento modela un sistema concurrente simple en el que:
- Los vehículos llegan al azar a la entrada del estacionamiento.
- Si hay un espacio disponible, el vehículo se estaciona por un período de tiempo aleatorio.
- Si el estacionamiento está lleno, los vehículos esperan en una cola hasta que un espacio se libere.
- Una interfaz gráfica muestra el estado del estacionamiento y de la cola de espera.

### Componentes de la Simulación
- **Lógica del Estacionamiento** (`models/parkinglot.go`): Gestiona la entrada, el estacionamiento y la salida de vehículos.
- **Generador de Vehículos**: Crea vehículos en intervalos aleatorios para simular la llegada real.
- **Interfaz Gráfica** (`view/View.go`): Visualiza los espacios del estacionamiento y los vehículos en espera.

## Arquitectura

Este proyecto está organizado en módulos con una clara separación de responsabilidades:
- **`models`**: Contiene la lógica principal de la simulación, incluyendo los modelos de estacionamiento y vehículos.
- **`view`**: Define la interfaz gráfica usando Fyne para visualizar los espacios de estacionamiento y la cola de espera.
- **`main.go`**: El punto de entrada de la aplicación, inicializando el estacionamiento y comenzando la simulación.

### Estructura de Archivos

. ├── assets ├── go.mod ├── go.sum ├── main.go ├── models │ ├── parkinglot.go │ └── vehicles.go ├── simulation └── view └── View.go

### Desarrollo

Durante el desarrollo de este proyecto, comprendí el cómo aplicar canales de go y sincronizarlos a través de channels y mutex, 
lo que me ayudó a comprender la naturaleza de la programación concurrente. Por otro lado, aprender a programar de forma concurrente fue
un reto, para el que tuve que prepararme estudiando las bases.

## Instalación

### Requisitos Previos

- Go 1.16+ instalado
- Biblioteca [Fyne](https://fyne.io/) para la GUI

### Pasos

1. Clonar el repositorio:
   ```bash
   git clone https://github.com/yourusername/parking-sim-go.git
   cd parking-sim-go

2. Inicializar el proyecto:
   ```bash
   go mod tidy
3. Ejecutar el proyecto:
  ```
  go run main.go
