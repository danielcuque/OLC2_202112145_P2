# Proyecto 2 - T-Swift Compiler

Ejecutar el siguiente comando para generar el parser y el lexer de ANTLR4

```bash
java org.antlr.v4.Tool -Dlanguage=Go -o ../core/parser/ Swift.g4 SwiftLexer.g4 -visitor -no-listener
```

Para ejecutar el servidor, se debe ejecutar el siguiente comando:

```bash
cd /server
go get
go run main.go
```

Para ejecutar el cliente, se debe ejecutar el siguiente comando:

```bash
cd /client
pnpm i
pnpm run dev
```