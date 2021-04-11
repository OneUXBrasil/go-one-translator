# ![one_translator](https://github.com/OneUXBrasil/go-one-translator) para Golang!
# Usando
## Instalando o modulo e importando
instalar o modulo: `go get github.com/OneUXBrasil/go-one-translator`
### no seu arquivo main
```go
import one_translator "github.com/OneUXBrasil/go-one-translator"
```
## exemplo
```go
package main

import one_translator "github.com/OneUXBrasil/go-one-translator"

func main() {
	result, _ := one_translator.Translate("Hello,world", "en", "pt")
	println(result)
}
```
