# Usa la imagen oficial de Go en la versión 1.22.3
FROM golang:1.22.3-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Descarga las dependencias y construye el binario de la aplicación
RUN go mod tidy
RUN go build -o main .

# Expone el puerto que usará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
