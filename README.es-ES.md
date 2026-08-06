

# HumanizerGPT

Este proyecto es un generador de ensayos/código con 0% de detección de IA. Probado con ZeroGPT, Quilbot AI detection, etc. ***AVISO: El autor no recomienda el uso de esta aplicación para ningún propósito poco ético***

Enlace a la entrega del hackathon: https://devpost.com/software/humanizergpt

## Stack Tecnológico

- **Frontend**: Vue.js
- **Backend**: Go

## Estructura del Proyecto

El proyecto se divide en dos directorios principales:

- `frontend`: Este directorio contiene todo el código Vue.js para la interfaz de la aplicación.
- `main`: Este directorio contiene el código Go para el backend de la aplicación.

## APIs

El servidor backend está construido con Go y sirve archivos estáticos desde el directorio `frontend/dist`. Utiliza el paquete `godotenv` para cargar variables de entorno desde un archivo `.env`.

## Cómo Ejecutar

1. Primero, debes compilar el frontend. Navega al directorio `frontend` y ejecuta `yarn install` para instalar las dependencias. Luego, ejecuta `yarn build` para compilar el frontend.
2. A continuación, navega al directorio `main`. Si aún no lo has hecho, ejecuta `go get` para descargar las dependencias de Go.
3. Ejecuta `go run main/server.go` para iniciar el servidor backend. El servidor servirá los archivos estáticos compilados en el paso 1.

Asegúrate de actualizar el archivo `.env` con tu configuración específica antes de ejecutar la aplicación.

## Contribuir

Se agradecen las solicitudes de extracción (pull requests). Para cambios importantes, abre primero un issue para discutir lo que te gustaría cambiar.

## Licencia

[MIT](https://choosealicense.com/licenses/mit/)
