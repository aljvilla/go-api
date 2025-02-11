# 📖 Guía de Instalación y Uso del Proyecto `miapp`

Bienvenido a `miapp`, un proyecto desarrollado en **Go** que utiliza **PostgreSQL** como base de datos.  
Esta guía cubre la instalación de **Go** en distintos sistemas operativos, la configuración de dependencias, el arranque del proyecto y la instalación de **PostgreSQL** con restauración de base de datos.

---

## 🚀 1️⃣ Instalación de Go en tu Sistema Operativo
### 🔹 MacOS
Puedes instalar Go usando **Homebrew** o descargando el instalador oficial.

#### ✅ Instalación con Homebrew (Recomendado)
```sh
brew install go
```
#### ✅ Instalación desde el sitio oficial
1. Descarga el instalador desde 👉 [https://go.dev/dl/](https://go.dev/dl/)
2. Ejecuta el instalador y sigue las instrucciones.
3. Verifica la instalación:
   ```sh
   go version
   ```

---

### 🔹 Windows
1. Descarga el instalador desde 👉 [https://go.dev/dl/](https://go.dev/dl/)
2. Ejecuta el `.msi` y sigue las instrucciones.
3. Verifica la instalación en **PowerShell** o **cmd**:
   ```sh
   go version
   ```

---

### 🔹 Linux (Ubuntu/Debian)
1. Descarga Go desde el sitio oficial:
   ```sh
   wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
   ```
2. Extrae el archivo en `/usr/local`:
   ```sh
   sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
   ```
3. Agrega Go al `PATH` (en `~/.bashrc` o `~/.profile`):
   ```sh
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   source ~/.bashrc
   ```
4. Verifica la instalación:
   ```sh
   go version
   ```

---

## 🛠️ 2️⃣ Instalación de PostgreSQL
### 🔹 MacOS
```sh
brew install postgresql
brew services start postgresql
```

### 🔹 Windows
Descarga el instalador desde 👉 [https://www.postgresql.org/download/](https://www.postgresql.org/download/)  
Sigue las instrucciones y usa **pgAdmin** o la línea de comandos.

### 🔹 Linux (Ubuntu/Debian)
```sh
sudo apt update
sudo apt install postgresql postgresql-contrib
```
Iniciar PostgreSQL:
```sh
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

---

## 🗄️ 3️⃣ Restaurar la Base de Datos
Si tienes un **dump** de la base de datos (`backup.sql`), puedes restaurarlo con:
```sh
psql -U postgres -d DATABASE_NAME -f backup.sql
```
📌 **Ejemplo completo:**
```sh
psql -U postgres -d miapp_db -f backup.sql
```
Si el usuario `postgres` tiene contraseña, usa:
```sh
PGPASSWORD="tu_contraseña" psql -U postgres -d miapp_db -f backup.sql
```

---

## 🔧 4️⃣ Configuración del Proyecto
Antes de iniciar el proyecto, debes configurar las variables de entorno:

### 🔹 Variables de Entorno Requeridas
Crea un archivo **`.env`** en la raíz del proyecto con:
```ini
DATABASE_HOST=localhost
DATABASE_NAME=miapp_db
DATABASE_PASSWORD=mi_secreta
DATABASE_PORT=5432
DATABASE_USERNAME=postgres
```
O exportarlas manualmente en la terminal:
```sh
export DATABASE_HOST="localhost"
export DATABASE_NAME="miapp_db"
export DATABASE_PASSWORD="mi_secreta"
export DATABASE_PORT="5432"
export DATABASE_USERNAME="postgres"
```

En **Windows (PowerShell)**:
```powershell
$env:DATABASE_HOST="localhost"
$env:DATABASE_NAME="miapp_db"
$env:DATABASE_PASSWORD="mi_secreta"
$env:DATABASE_PORT="5432"
$env:DATABASE_USERNAME="postgres"
```

---

## 🚀 5️⃣ Instalación de Dependencias y Arranque del Proyecto
Después de instalar **Go** y **PostgreSQL**, sigue estos pasos:

### 🔹 1. Clona el repositorio
```sh
git clone https://github.com/tu-usuario/miapp.git
cd miapp
```

### 🔹 2. Inicializa y limpia dependencias
```sh
go mod tidy
```

### 🔹 3. Arranca el Servidor
```sh
go run cmd/server/main.go
```
📌 **Si todo está bien, verás algo como:**
```
Servidor corriendo en http://localhost:8080
✅ Conexión a PostgreSQL establecida
```

---

## 🔍 6️⃣ Comandos Útiles
### 📌 Compilar y generar un binario
```sh
go build -o miapp cmd/server/main.go
./miapp
```
### 📌 Ejecutar pruebas
```sh
go test ./...
```
### 📌 Verificar dependencias
```sh
go mod verify
```

---

## 🎯 7️⃣ Endpoints Disponibles
| Método | Ruta | Descripción |
|--------|------|------------|
| `GET`  | `/empresas` | Obtiene todas las empresas |
| `POST` | `/empresas/create` | Crea una nueva empresa |
| `PUT`  | `/empresas/update/{id}` | Actualiza una empresa por ID |

📌 **Ejemplo de solicitud `POST`**
```sh
curl -X POST http://localhost:8080/empresas/create      -H "Content-Type: application/json"      -d '{
           "razon_social": "Mi Empresa S.A.",
           "numero_identificador": "12345678",
           "tipo_numero_identificador": "RUC"
         }'
```

---

## 🎯 Conclusión
🔹 **Go y PostgreSQL están listos para trabajar**.  
🔹 **Las variables de entorno están configuradas**.  
🔹 **El servidor se ejecuta con `go run cmd/server/main.go`**.  
🔹 **Los endpoints están listos para ser consumidos**.  

🚀 **¡Felicidades! Ahora puedes empezar a desarrollar en `miapp`**  
💡 **Si tienes dudas, revisa la documentación o pregunta en la comunidad.** 😊
