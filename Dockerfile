# Usar una imagen base oficial de Go 1.23
FROM golang:1.23-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum al contenedor
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod tidy

# Copiar el archivo main.go al contenedor
COPY main.go .

# Exponer el puerto en el que la aplicación Go escuchará (si aplica)
EXPOSE 8080

# Compilar el código Go
RUN go build -o main .

# Definir el comando por defecto para ejecutar la aplicación
CMD ["./main"]
