# golang-study

## Util

- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)

## Executando

1. Install [asdf](https://asdf-vm.com/guide/getting-started.html#_3-install-asdf)

2. Adicionar plugin **golang**

```bash
asdf plugin-add golang https://github.com/kennyp/asdf-golang.git
```

1. Instalar **golang**

```bash
asdf install
```

## Commands

## go mod

### Inicializando modulo

```bash
go mod init <name-module>
```

&nbsp;

## go test

### Executando test

```bash
# Entre na pasta contendo o teste

go test

# Executando de forma recursiva, é necessário ter modulo go
# Flags
# -v: verboso
# --cover: gera porcentagem de cobertura
# --coverprofile <nome-file>.txt: gera relatório de cobertura 
relatório de cobertura

go test ./...
```

### Executando em paralelo

```go
// Adicionar no início de cado função do teste

func TestParapelo(t *testing.T) {
    t.Parallel()

    return "true"
}
```

## go tool

### Obter informações gerado pelo `--coverprofile`

```bash
# Informações genéricas
go tool cover --func=<nome-file>.txt

# Gerando html
go tool cover --html=<nome-file>.txt
```
