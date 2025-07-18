// cmd/artificialmasterpiecegeneratorprotocolpro/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratorprotocolpro/internal/artificialmasterpiecegeneratorprotocolpro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratorprotocolpro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
