# Arquitectura de Software 2 - Universidad Nacional de Quilmes

[Enunciado del Trabajo Práctico](https://github.com/cassa10/arq2-tp1/blob/main/doc/Arq2%20-%20Trabajo%20pr%C3%A1ctico.pdf)

## Tecnologías:

- [Golang](https://go.dev/)
- [Gin (WEB API)](https://gin-gonic.com/)
- [MongoDB](https://www.mongodb.com/)

## Prerequisitos:

- Go 1.20 or up / Docker


## Test y coverage

### Pasos:

1) Ir al folder root del repositorio

2) Ejecutar los comandos

```
> go test -coverprofile="coverage.out" -covermode=atomic ./src/domain/...
> go install gitlab.com/fgmarand/gocoverstats@latest
> gocoverstats -v -f coverage.out -percent > coverage_rates.out
```

3) Se generará el file coverage.out el cual contedrá el info del coverage del folder ./src/domain, y 
overage_rates.out que contendrá los porcentajes dicho coverages.

**Nota**: En la construcción de la imagen, se realiza la ejecucion de test y la generacion del coverage.
Los archivos de coverage son almacenados dentro del container en la carpeta "/app". Es decir, que encontraremos:
- /app/coverage.out
- /app/coverage_rates.out

Dichos archivos se pueden acceder desde un volume vinculado o sino ejecutando bash dentro del container los comando que se menciono anteriormente.


## Swagger

Instalar swag localmente (se necesita go 1.20 or up)

```
go install github.com/swaggo/swag/cmd/swag@v1.8.10
```

Para actualizar la api doc de swagger, ejecutar en el folder root del repo:

```
swag init -g src/infrastructure/api/app.go
```

Luego de levantar la api e ir al endpoint:

```
http://localhost:<port>/docs/index.html
```


## Inicialización y ejecución del proyecto (docker)

### Pasos:

1) Ir a la carpeta root del repositorio

2) Construir el Dockerfile (imagen) del servicio

```
docker build -t arq2-tp1 .
```

3) Ejecutar la imagen construida

Importante: Se requiere configurar env var "MONGO_URI" dentro de ./resources/local.env con `"mongodb+srv://<user>:<password>@cluster0.80ymcdr.mongodb.net/<database>?retryWrites=true&w=majority"`

database = "arq-soft-2-meli"

Nota: Pedir credenciales por privado.

Tambien, si se desea se puede cambiar las envs por otras de las que estan. Se recomienda utilizar el mismo puerto externo e interno para que funcione correctamente swagger.

```
docker run -p <port>:8080 --env-file ./resources/local.env --name arq2-tp1 arq2-tp1
```

Nota: agregar "-d" si se quiere ejecutar como deamon

```
docker run -d -p <port>:8080 --env-file ./resources/local.env --name arq2-tp1 arq2-tp1
```

Ejemplo:

```
docker run -d -p 8080:8080 --env-file ./resources/local.env --name arq2-tp1 arq2-tp1
```

4) En un browser, abrir swagger del servicio en el siguiente url:

`http://localhost:<port>/docs/index.html`

Segun el ejemplo:

`http://localhost:8080/docs/index.html`

5) Probar el endpoint health check y debe retornar ok

6) La API esta disponible para ser utilizada

