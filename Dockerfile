# Utiliza la imagen base oficial de Go para compilar la aplicación
FROM golang:1.21 as builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del módulo Go y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto de los archivos de la aplicación
COPY . .

# Compila la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-jwt .

# Utiliza la imagen base de alpine por ser una imagen ligera
FROM alpine:latest

# Instalar ca-certificates en caso de que tu aplicación realice llamadas HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia el ejecutable compilado desde la imagen de compilación
COPY --from=builder /app/go-jwt .

# Copia el archivo de base de datos SQLite, si ya existe
# Si el archivo de la base de datos se va a generar en tiempo de ejecución o se monta un volumen, esta línea no es necesaria
COPY --from=builder /app/*.db ./

# Expone el puerto que Gin está escuchando
EXPOSE 8080

# Ejecuta la aplicación
CMD ["./go-jwt"]