# Universidad de San Carlos de Guatemala

## Facultad de Ingeniería

## Escuela de Ciencias y Sistemas

## Organización de Lenguajes y Compiladores 2

## 202112145

## Daniel Estuardo Cuque Ruíz

## **Manual técnico**

### Índice

- [Introducción](#introducción)
- [Instalación](#instalación)
- [Estructura del proyecto](#estructura-del-proyecto)
- [Estructura del servidor](#estructura-del-servidor)
- [Estructura de la gramática](#estructura-de-la-gramática)
- [Estructura de la API](#estructura-de-la-api)
- [Estructura de las estructuras](#estructura-de-las-estructuras)
- [Estructura del cliente](#estructura-del-cliente)

### Introducción

El presente documento tiene como objetivo describir el funcionamiento interno del intérprete de lenguaje de programación **T-Swift**. Se describirán los componentes principales del intérprete, así como su funcionamiento interno y la forma en que se relacionan entre sí.

### Instalación

Para poder ejecutar el intérprete, es necesario contar con las siguientes herramientas:

- **Node.js**: Entorno de ejecución para JavaScript. Puede ser descargado desde el siguiente enlace: https://nodejs.org/es/download/

- **Golang**: Lenguaje de programación. Puede ser descargado desde el siguiente enlace: https://golang.org/dl/

- **ANTLR4**: Herramienta para generar analizadores léxicos y sintácticos. Puede ser descargado desde el siguiente enlace: https://www.antlr.org/download.html

### Estructura del proyecto

El proyecto está dividido en dos carpetas:
- **client**: Contiene el código fuente del cliente web, el cual se encarga de la interfaz gráfica del intérprete.	

- **server**: Contiene el código fuente del servidor, el cual se encarga de la lógica del intérprete.

### Estructura del servidor

Para poder ejecutar el servidor, es necesario ejecutar el siguiente comando en la carpeta raíz del proyecto:

```bash
cd server/

go run main.go
```

El servidor está dividido en 4 paquetes principales:

- **api**: Contiene la lógica de la API REST del intérprete.

- **grammar**: Contiene la lógica del analizador léxico y sintáctico del intérprete.

- **chore**: Contiene las estructuras que manejan los objetos dentro del intérprete.

Para compilar la gramática del intérprete, es necesario ejecutar el siguiente comando en la carpeta raíz del proyecto:

```bash
cd server/grammar/

java org.antlr.v4.Tool -Dlanguage=Go -o ../chore/parser/ Swift.g4 SwiftLexer.g4 -visitor -no-listener
```

### Estructura de la gramática

La gramática del intérprete está dividida en 2 archivos:

- **Swift.g4**: Contiene la gramática del intérprete.

- **SwiftLexer.g4**: Contiene la gramática del analizador léxico del intérprete.

### Estructura de la API

La API del intérprete está dividida en 2 archivos:

- **api.go**: Contiene la lógica de la API REST del intérprete.

- **api_test.go**: Contiene las pruebas unitarias de la API REST del intérprete.

### Estructura de las estructuras

Para el manejo de los valores dentro del intérprete, se utiliza la interfaz IVALUE, donde cualquier tipo de dato, ya sea booleano, entero, flotante, cadena, arreglo, diccionario, etc., implementa dicha interfaz.

```go	
type IValue interface {
	GetValue() interface{}
	GetType() string
	String() string
	Copy() IValue
}

type RangeV struct {
	Value []IValue
}

func (a *RangeV) GetValue() interface{} {
	return a.Value
}

func (a *RangeV) GetType() string {
	return IntType
}

func (a *RangeV) String() string {
	return ""
}

func (a *RangeV) Copy() IValue {
	return NewRangeValue(a.Value)
}

func NewRangeValue(value []IValue) *RangeV {
	return &RangeV{Value: value}
}

func IsBaseType(value IValue) bool {
	switch value.GetType() {
	case StringType:
		return true
	case IntType:
		return true
	case FloatType:
		return true
	case BooleanType:
		return true
	case CharType:
		return true
	case NilType:
		return true
	default:
		return false
	}
}

func IsBaseTypeString(value string) bool {
	switch value {
	case StringType:
		return true
	case IntType:
		return true
	case FloatType:
		return true
	case BooleanType:
		return true
	case CharType:
		return true
	case NilType:
		return true
	default:
		return false
	}
}
```


Para el manejo de los objetos dentro del intérprete, se utiliza la estructura Object, la cual contiene un valor y un tipo de dato. Además, se utiliza la estructura Environment, la cual contiene un mapa de objetos, el cual se utiliza para almacenar las variables dentro del intérprete.

```go
package interpreter

import (
	V "OLC2/core/values"
	"fmt"
)

const (
	RootEnv   = "Root"
	FuncEnv   = "Func"
	WhileEnv  = "While"
	ForEnv    = "For"
	IfEnv     = "If"
	ElseEnv   = "Else"
	SwitchEnv = "Switch"
	StructEnv = "Struct"
)

type TokenSymbol struct {
	Env      string
	Type     string
	Name     string
	DataType string
	Value    string
	Params   []string
}

func NewTokenSymbol(env, tokenType, name, dataType, value string, params []string) *TokenSymbol {
	return &TokenSymbol{
		Env:      env,
		Type:     tokenType,
		Name:     name,
		DataType: dataType,
		Value:    value,
		Params:   params,
	}
}

type EnvNode struct {
	Parent    *EnvNode
	Child     []*EnvNode
	Level     int
	ScopeType string
	Variables map[string]*Variable
	Functions map[string]*Function
	Structs   map[string]*ObjectV
}

func NewEnvNode(parent *EnvNode, envType string, Level int) *EnvNode {
	return &EnvNode{
		Parent:    parent,
		Child:     make([]*EnvNode, 0),
		Level:     Level,
		ScopeType: envType,
		Variables: make(map[string]*Variable),
		Functions: make(map[string]*Function),
		Structs:   make(map[string]*ObjectV),
	}
}

func (s *EnvNode) AddVariable(name string, value *Variable) {
	s.Variables[name] = value
}

func (s *EnvNode) GetVariable(name string) *Variable {
	if val, ok := s.Variables[name]; ok {
		return val
	}
	return nil
}

func (s *EnvNode) SetVariable(name string, value V.IValue) {
	if val, ok := s.Variables[name]; ok {
		val.SetValue(value)
	}
}

func (s *EnvNode) ResetEnvNode() {
	s.Variables = make(map[string]*Variable)
}

func (s *EnvNode) GetType() string {
	return string(s.ScopeType)
}

func (s *EnvNode) String() string {
	result := ""
	for i := 0; i < s.Level; i++ {
		result += "\t"
	}
	result += fmt.Sprintf("%s\n", s.ScopeType)
	for _, v := range s.Child {
		result += v.String()
	}
	return result
}

func (s *EnvNode) Copy() *EnvNode {
	newNode := NewEnvNode(nil, s.ScopeType, s.Level)
	for _, child := range s.Child {
		newNode.Child = append(newNode.Child, child.Copy())
	}
	for key, variable := range s.Variables {
		newNode.Variables[key] = variable.CopyVar()
	}
	for key, function := range s.Functions {
		newNode.Functions[key] = function
	}
	for key, object := range s.Structs {
		newNode.Structs[key] = object.Copy().(*ObjectV)
	}
	return newNode
}

// EnvTree is a nary tree to represent scopes
type EnvTree struct {
	Root    *EnvNode
	Current *EnvNode
}

func NewEnvTree() *EnvTree {
	root := NewEnvNode(nil, RootEnv, 0)
	return &EnvTree{
		Root:    root,
		Current: root,
	}
}

func (s *EnvTree) AddVariable(name string, value *Variable) {
	s.Current.AddVariable(name, value)
}

func (s *EnvTree) GetVariable(name string) *Variable {
	for node := s.Current; node != nil; node = node.Parent {
		if val, ok := node.Variables[name]; ok {
			return val
		}
	}

	return nil
}

func (s *EnvTree) SetVariable(name string, value V.IValue) {
	s.Current.SetVariable(name, value)
}

func (s *EnvTree) AddFunction(name string, value *Function) {
	s.Current.Functions[name] = value
}

func (s *EnvTree) GetFunction(name string) *Function {
	for node := s.Current; node != nil; node = node.Parent {
		if _, ok := node.Functions[name]; ok {
			return node.Functions[name]
		}
	}
	return nil
}

func (s *EnvTree) AddStruct(name string, value *ObjectV) {
	s.Current.Structs[name] = value
}

func (s *EnvTree) GetStruct(name string) *ObjectV {
	for node := s.Current; node != nil; node = node.Parent {
		if _, ok := node.Structs[name]; ok {
			return node.Structs[name]
		}
	}
	return nil
}

func (s *EnvTree) PushEnv(scopeType string) *EnvNode {
	node := NewEnvNode(s.Current, scopeType, s.Current.Level+1)
	s.Current.Child = append(s.Current.Child, node)
	s.Current = node
	return node
}

func (s *EnvTree) PopEnv() {
	s.Current = s.Current.Parent
}

func (s *EnvTree) GetSymbolTable() []ApiVariable {
	// Traverse tree to get symbol table
	return s.Root.GetAllVariables()
}

func (s *EnvTree) ResetEnv() {
	// Clean all variables inside current scope
	s.Current.ResetEnvNode()
}

func (s *EnvTree) GetCurrentScope() *EnvNode {
	return s.Current
}

func (s *EnvTree) String() string {
	return s.Root.String()
}

// ApiVariable is a struct to represent variables in api
type ApiVariable struct {
	Name    string
	IsConst bool
	Value   interface{}
	Type    string
	Line    int
	Column  int
	Scope   string
}

func (s *EnvNode) GetAllVariables() []ApiVariable {
	allVariables := make([]ApiVariable, 0)
	s.collectVariables(&allVariables)
	return allVariables
}

func (s *EnvNode) collectVariables(allVariables *[]ApiVariable) {
	for _, variable := range s.Variables {
		apiVar := ApiVariable{
			Name:    variable.GetName(),
			IsConst: variable.IsConstant(),
			Value:   variable.GetValue(),
			Type:    variable.GetType(),
			Line:    variable.GetLine(),
			Column:  variable.GetColumn(),
			Scope:   s.GetType(),
		}
		*allVariables = append(*allVariables, apiVar)
	}
	for _, child := range s.Child {
		child.collectVariables(allVariables)
	}
}

```

Para guardar el árbol de ambientes, se utiliza la estructura Tree, la cual contiene un ambiente, un padre y un hijo. Además, se utiliza la estructura TreeList, la cual contiene una lista de árboles.

A partir del árbol de ambientes, se extrae la tabla de símbolos, la cual contiene el nombre, el tipo, el valor, la línea, la columna y el ámbito de cada variable.

También tiene métodos que funcionan para agregar, obtener y modificar variables, funciones y estructuras dentro del árbol de ambientes.


También se utiliza la estructura Function, la cual contiene un nombre, una lista de parámetros, una lista de instrucciones y un ambiente. Además, se utiliza la estructura FunctionList, la cual contiene una lista de funciones.


La estructura Variable contiene un nombre y un valor. Además, se utiliza la estructura VariableList, la cual contiene una lista de variables.


### Estructura del cliente

El cliente está desarrollado con React.js

Para poder ejecutar el cliente, es necesario ejecutar el siguiente comando en la carpeta raíz del proyecto:

```bash
cd client/
pnpm install
pnpm run dev
```

### Tecnologías utilizadas

- [Pnpm](https://pnpm.io/) - Gestor de paquetes
- [React](https://reactjs.org/) - Biblioteca de JavaScript para construir interfaces de usuario
- [Jison](https://github.com/zaach/jison) - Generador de analizadores léxicos y sintácticos
- [Tailwind CSS](https://tailwindcss.com/) - Framework de CSS
- [Vite](https://vitejs.dev/) - Herramienta de desarrollo web
- [TypeScript](https://www.typescriptlang.org/) - Lenguaje de programación
- [ESLint](https://eslint.org/) - Herramienta de análisis de código estático para identificar patrones problemáticos encontrados en el código JavaScript
- [Prettier](https://prettier.io/) - Formateador de código





