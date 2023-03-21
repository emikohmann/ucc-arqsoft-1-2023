# Guia Práctica

## Comandos de GO

### Corroborar la instalación y versión de GO

```bash
go version
```

### Correr un programa GO

> Pararse siempre en la carpeta que contiene el archivo `main.go`

```bash
go run main.go
```

### Compilar y correr el programa en 2 pasos

> Pararse siempre en la carpeta que contiene el archivo `main.go`<br/>Para Linux o macOS, quitar la extensión `.exe`

```bash
go build
./nombre_del_ejecutable.exe
```

### Crear un módulo para poder usar paquetes

> Pararse siempre en la carpeta que contiene el archivo `main.go`

```bash
go mod init
go mod tidy
```

> Para todos los imports, usar el prefijo `nombre_del_modulo`, por ejemplo `import "go-api/my/package"`

### Formatear todos los archivos del proyecto

> Pararse siempre en la carpeta que contiene el archivo `main.go`

```bash
go fmt ./...
```

## Comandos de GIT

### Iniciar un nuevo repositorio

> Crear una carpeta con el nombre del repositorio <br /> Pararse en la carpeta del repositorio

```bash
git init
````

### Clonar un repositorio

> Pararse en la carpeta de proyectos

```bash
git clone https://github.com/{username}/{repo_name}.git
```

### Crear un nuevo branch

```bash
git checkout -b {branch_name}
```

### Configurar un remote

> No es necesario si se clona el repositorio

```bash
git remote add {remote_name} https://github.com/{username}/{repo_name}
```

### Ver el listado de remotes

```bash
git remote -v
```

### Hacer un commit

```bash
git add .
git commit -m "Commit message"
git push origin {branch_name}
```

### Bajar los cambios del remote

```bash
git pull origin {branch_name}
```
