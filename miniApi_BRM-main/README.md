# Mini api BRM

API  básica desarrollada en Go para la gestión de usuarios.


## Características

- CRUD completo de usuarios
- Arquitectura usada:  en capas
- Patrón de diseño:  Repository
- Base de datos MySQL:  Implementada con Aiven datos de conexión en el file .env
- Dockerfile
- Código documentado, tanto con uso como aprendizajes del proyecto


## Arquitectura del proyecto


El proyecto sigue una arquitectura en capas, esto con el fin de mejorar el orden e identificar más facilmente elementos:

```
miniApi_BRM/
├── cmd/api/          
├── internal/
│   ├── domain/       # Modelos de dominio
│   ├── repository/   # Capa de acceso a datos (Patrón Repository)
│   ├── service/      # Lógica de negocio
│   ├── http/         
│   └── db/           # Configuración de base de datos
└── scripts/          
```

## Descargar API

1. Clonar el repositorio:
```bash
git clone <https://github.com/sergiosocha/miniApi_BRM.git>
cd miniApi_BRM
```

2. Instalar dependencias:
```bash
go mod download
```

3. Ejecutar la aplicación:
```bash

    go run cmd/api/main.go
```

3. Opcional ejecutar desde docker:
```bash
    docker build -t miniapi .
    docker run --rm -p 8080:8080 miniapi
```

3. Visualizar desde Postman o navegador a traves de:
```bash
    http://localhost:8080/users
```
## Documentación uso de API

### URL de despligue local
```
http://localhost:8080
```

### Endpoints

#### Listado de los endpoints disponibles al usar la API

| # | Método | Ruta           | Descripción              |
|---|--------|----------------|--------------------------|
| 1 | POST   | `/users`       | Crear usuario            |
| 2 | GET    | `/users`       | Listar usuarios          |
| 3 | GET    | `/users/{id}`  | Obtener usuario por ID   |
| 4 | PUT    | `/users/{id}`  | Actualizar usuario       |
| 5 | DELETE | `/users/{id}`  | Eliminar usuario         |


### COMO USAR:
#### 1. Crear Usuario
```http
POST /users
Content-Type: application/json

{
  "name": "Sergio Socha",
  "email": "sergi@test.com"
}
```

**Respuesta:**
```json
{
  "id": 1,
  "name": "Sergio Socha",
  "email": "sergi@test.com",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:00:00Z"
}
```

#### 2. Obtener todos los usuarios
```http
GET /users
```

**Respuesta:**
```json
[
  {
    "id": 1,
    "name": "Sergio Socha",
    "email": "sergi@test.com",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
]
```

#### 3. Obtener usuario por ID
```http
GET /users/{id}
```

**Respuesta:**
```json
{
  "id": 1,
  "name": "Sergio Socha",
  "email": "sergi@test.com",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:00:00Z"
}
```

#### 4. Actualizar usuario
```http
PUT /users/{id}


{
  "name": "Sergio Socha actualizado",
  "email": "sergi_update@test.com",
}
```

**Respuesta:**
```json
{
  "id": 1,
  "name": "Sergio Socha actualizado",
  "email": "sergi_update@test.com",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:30:00Z"
}
```

#### 5. Eliminar usuario
```http
DELETE /users/{id}
```

**Respuesta:**
```json
{
  "message": "Usuario eliminado exitosamente"
}
```



## Base de Datos

### Configuración Aiven MySQL

- **Host**: usersapi-miniapi01.j.aivencloud.com
- **Port**: 20927
- **Database**: defaultdb
- **User**: avnadmin
- **SSL**: Required

Credenciales db son suministrados en correo todo se trabaja en .env para evitar
exponer datos sensibles 

### Estructura de la tabla

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

##  Tecnologías Utilizadas

- **Go 1.21**: Lenguaje de programación
- **Gorilla Mux**: Router HTTP
- **MySQL Driver**: Conexión a base de datos
- **Docker**: Containerización
- **Postman**: Prueba de endpoints
- **Github**: Repositorio
- **Aiven MySQL**: Base de datos en la nube
##  Autor

* Sergio Eduardo Socha 
* Ingeniero informatico
* email: sergio_socha@yahoo.com