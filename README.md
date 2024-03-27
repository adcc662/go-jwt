# Go-jwt

## Objetivo
Crear un servicio de Go que permita autenticar usuarios y generar un token JWT.

## Tecnologias usadas
- **Go:** Elegido por ser el stack que se usa en la empresa.
- **Gin:** Al ser un framework de web te ofrece cosas que ya vienen para implementarse de una manera mas sencilla.
- **SQLite:** Base de datos que se eligió por ser un archivo y no tener que instalar un motor de base de datos y hacer un desarrollo mas rápido.
- **Gorm:** ORM que se eligió por ser muy sencillo de usar y tener una buena documentación.

## Estructura del proyecto
```
.
├── Dockerfile
├── README.md
├── common
│   ├── database.go
│   └── utils.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
└── users
    ├── middlewares.go
    ├── models.go
    ├── routers.go
    ├── serializers.go
    └── validators.go
```

## Instalación
1. Clonar el repositorio
```sh
git clone git@github.com:adcc662/go-jwt.git
cd go-jwt
docker build -t go-jwt .
docker run -p 8080:8080 go-jwt

#Si se require usar sin docker podemos usar:
go build
go run main.go
```

## Endpoints
- **POST /api/users**
  - **Descripción:** Registra un usuario
    - **Body:**
      ```json
      {
           "username": "testUser",
           "email": "user@example.com",
           "password": "Password123!",
           "phone": "1234567890"
    }
    ```

- **POST /api/users/login**
- **Descripción:** Inicia sesión y devuelve un token JWT
  - **Body:**
    ```json
    {
        "email": "user@example.com",
        "password": "Password123!"
    }
    ```
- **GET /api/users**
- **Descripción:** Devuelve todos los usuarios
- **Headers:**
  - **Authorization:** Bearer token
  - **Content-Type:** application/json

- **PUT /api/user**
- **Descripción:** Actualiza un usuario
  - **Headers:**
    - **Authorization:** Bearer token
    - **Content-Type:** application/json
  - **Body:**
    ```json
    {
          "username": "updatedUser",
          "phone": "0987654321"
    }
    ```
## Respuestas al cuestionario
El archivo con las respuestas al cuestionario se encuentra en el archivo [ANSWERS.md](ANSWERS.md)