// let a = true
// let aa = false
// var b =  2
// var c = 1.1 + 2
// var d = -c * -1
// var e = false + 1.1 // Error
// var f = m + 1 // Error
// var evalWhile = 10
// var num: Float = 1
// var num3: Int = 10.1
// var num2: Int = false
// evalWhile = 5
// evalWhile = false // Error



// // var g = 1 > 1.1

// // var h = "hola" == "hola"

// // var i = "ho" + "la"

// // let j = 1 > false // Error

// // let k = 10/0
// // let l = 10%0

// // let m: Float?
// // var n: Int?

// // Test simple if with nested if
// if a {
//     let testIfB = 6
// 	evalWhile = testIfB
//     if !aa {
//         evalWhile = 7
//         if b < c{
//             evalWhile = 8
//         }
//     }

//     if !!a{
//         evalWhile = 9
//     }
// }

// // Test if else
// var testIf = 0
// if !a {
//     testIf = 10
// } else {
//     testIf = 0
// }

// // Test if else

// var testIfElse = 0

// if a {
//     testIfElse = 100
// } else {
//     testIfElse = 0
// }

// // Test if else if

// var testIfElseIf = 0

// var o = "b"

// if b == "a" {
//     let am = 10
//     testIfElseIf = am
// } else if o == "b" {
//     testIfElseIf = 12
// } else if o == "c" {
//     testIfElseIf = 13
// } else {
//     testIfElseIf = 14
// }

// // Test switch

// switch o {
//     case "a":
//         let testSwitch = 10
//         testIfElseIf -= testSwitch
//     case "b":
//         testIfElseIf += 11
//     case "c":
//         testIfElseIf = 13
//     default:
//         testIfElseIf = 14
// }

// // Test While
// var evalWhile = 20
// var a = 0

// while evalWhile > 0 {
//     if evalWhile > 10 {
//         while evalWhile > 10 {
//             evalWhile -= 1
//             break
//         }
//     }
    
//     if evalWhile == 20 {
//         break
//     } else if evalWhile == 10 {
//         evalWhile -= 10
//         continue
//     }

//     switch evalWhile {
//         case 20:
//             let rest = 100
//             evalWhile -= rest
//         case 10:
//             evalWhile -= 10
//         default:
//             evalWhile -= 1
//     }
//     evalWhile -= 1

// }


// var nums = 100

// var testContinue = 0

// while nums >= 0 {
//     if nums % 2 == 0 {
//         nums -= 1
//         break
//     }
//     testContinue += 1
//     nums -= 1
// }

// var testFor = 0

// for i in 6...10 {
//     print(i)
// }

// print(testFor, "testFor")

// var testStringFor = "hola"

// for i in testStringFor {
//     let bbb = 1
//     if i == "a" {
//         break
//     }

//     testFor += bbb
// }

// var testGuard = 2
// var num = 0

// while (testGuard <= 10){
//     guard testGuard % 2 == 0 else {
//         print(testGuard, "es impar")
//         testGuard += 1
//         continue
//     }
//     print(testGuard, "es par")
//     num += testGuard
//     testGuard = testGuard + 1
// }

// print(testGuard, "|", num)

// print(Int("10"))
// print(Int("12.001"))
// print(Int(10.9999))
// print(Int("Q10"))

// print(Float("10"))
// print(Float("10.001"))
// print(Float("Q10.00"))

// print(1.0001)
// print(true)
// print(nil)
// print("\tcadena1 \ncadena2") // mostraría cadena1 y cadena2 en líneas separadas
// Test double quote
// print("cadena1 \"cadena2\" cadena3")

// Vectors

//vector con valores
var vec1: [Int] = [20, 30, 40] 
// //vector vacío
// var vec2: [Float] = [1.1, 2.2, 3.3]
// //vector vacío
// var vec3: [String] = []

// var copiaVec: [Int] = vec1

// print("La posicion 1 de vec1 es ->", vec1[0])

let number = vec1[0]
// // let number2: String = vec1[0] // Error
// let number3: Int = vec1[1]

// print("La posicion 0 de vec1 es ->", number)
// print("La posicion 1 de vec1 es ->", number3)

vec1[0] = 10 + number
// print("El nuevo valor de la posicion 0 de vec1 es ->", vec1[0]) // Debería ser 30
print("La cantidad de objetos en el vector1 es de:", vec1.count, "está vacío:", vec1.isEmpty)
// print("La cantidad de objetos en el vector2 es de:", vec2.count, "está vacío:", vec2.isEmpty)
// print("La cantidad de objetos en el vector3 es de:", vec3.count, "está vacío:", vec3.isEmpty)
vec1.append(50)


// // persona.vec1.append(50) // Error


// print("\n")
// while vec1.count > 0 {
//     print("vec1 tiene", vec1.count, "elementos")
//     vec1.removeLast()
// }
// print("\n")

// vec2.remove(at: 3) // Error
// vec2.remove(asds: 2) // Error


print("La nueva cantidad de vec1 es de:", vec1.count, "está vacío:", vec1.isEmpty)


func fibonacci(_ n: Int) -> Int {
    if n > 1 {
        return fibonacci(n - 1) + fibonacci(n - 2)
    } else if n == 1 {
        return 1
    } else if n == 0 {
        return 0
    } else {
        print("error")
        return 0
    }
}

func sumar2 (_ a: inout Int, _ b: inout Int) {
    a = a + b
}

let fibonacci5 = 5
sumar2(&fibonacci5, &fibonacci5)

print("Debería ser 55")
print(fibonacci(fibonacci5)) 
print("El valor de fibonacci10 es", fibonacci5)

func ackerman(_ m: Int, _ n: Int) -> Int {
    if m == 0 {
        return n + 1
    } else if m > 0 && n == 0 {
        return ackerman(m - 1, 1)
    } else {
        return ackerman(m - 1, ackerman(m, n - 1))
    }
}

print("Debería ser 125")
print(ackerman(3,4))

func Hanoi(_ discos: Int, _ origen: Int, _ auxiliar: Int, _ destino: Int) {
    if discos == 1 {
        print("Mover disco de", origen, "a", destino)
    } else {
        Hanoi(discos - 1, origen, destino, auxiliar)
        print("Mover disco de", origen, "a", destino)
        Hanoi(discos - 1, auxiliar, origen, destino)
    }
}

Hanoi(3, 1, 2, 3)


// Matrix declaration

var mtx1: [[Int]] = [[1], [4, 5], [7, 8, 9]]

var mtx2 : [[[Int]]] = [[[1,2,3],[4,5,6],[7,8,9]], 
                        [[10,11,12],[13,14,15],[16,17,18]], 
                        [[19,20,21],[22,23,24],[25,26,27]]]

// Using repeating

var matrix2 : [[[Int]]] = [[[Int]]] (repeating: [[Int]] (repeating: [Int] (repeating: 0, count:2), count:3), count:4)

var matrix0 : [[[String]]] = [[[String]]] (repeating: [[String]] (repeating: [String] (repeating:"OLC2", count:2), count:1), count:3)


print(vec1[0], "======") //imprime 20
vec1[0] = mtx1[0][0] //cambia 20 por 1

print(vec1[0], "======") //imprime 1

print(mtx1[1][1], "^^^^") //imprime 5
mtx1[1][1] = 10 //cambia 5 por 10 
print(mtx1[1][1], "^^^^") //imprime 10

print(mtx1[0][0]) //imprime 1 

print(mtx2[0][1][2], "-----") //imprime 6
mtx2[0][1][2] = 10 //cambia 6 por 10
print(mtx2[0][1][2], "-----") //imprime 10

// error indices fuera de rango - error
// mtx1[100][100] = 10


func negar(_ x: inout Int){ 
    let a = -x
    print("negado:",a)
    // x = -x
}

let numero1 = 10

negar(&numero1)
print("numero1:", numero1) // imprime -10

let numero2 = 20
let numero3 = 30

func sumar(x: Int, y: Int) -> Int{
    x += 100
    return x + y
}

print("suma:", sumar(x: numero2, y: numero3)) // imprime 50

print("numero2:", numero2) // imprime 20
print("numero3:", numero3) // imprime 30


// Vector por referencia

func imprimirArray (_ simpleArray: [Int] ) {
    for i in 0...simpleArray.count - 1 {
        print(simpleArray[i])
    }
}

func duplicarConPuntero1 (_ array1: inout [Int] ) {
    var i = 0
    while (i < array1.count ) {
        array1[i] += array1[i]
        i += 1
    }
}

func duplicarConPuntero2 (_ array2: inout [Int] ) {
    duplicarConPuntero1(&array2)
}

func duplicarSinPuntero (_ arrayS: [Int] ) {
    for i in 0...arrayS.count - 1 {
        arrayS[i] += arrayS[i]
    }
}


var array: [Int] = [1,2,3,4,5,6]

print("Array original")
imprimirArray(array)

print("Array sin puntero")
duplicarSinPuntero(array)
imprimirArray(array)

print("Array con puntero")
duplicarConPuntero1(&array)
imprimirArray(array)

print("Array con puntero doble")
duplicarConPuntero2(&array)
imprimirArray(array)

// Structs

struct Persona{
    var Nombre: String
    var edad = 0
    var apellido = "García"

    func mostrarNombre(){
        print("Mi nombre es", self.Nombre, self.apellido)
    }

}

var persona1 = Persona(Nombre:"Daniel", edad: 21, apellido: "Cuque")

print("Mi nombre es", persona1.Nombre, persona1.apellido, "y tengo", persona1.edad, "años")


let pruebaGlobal = 10
// struct con funciones 
struct Avion {
    var pasajeros = 0 
    var velocidad = 100
    var nombre: String?
    var piloto: Persona

    // metodo dentro de struct 

    mutating func frenar(){
        print("El avion", self.nombre, "esta frenando", "velocidad actual", self.velocidad)
        self.velocidad = pruebaGlobal
        print("La nueva velocidad es de:", self.velocidad)
    }
        
    // funcion inmutable 
    func mostrarVelocidad(velocidad: Int){
        print("La velocidad actual es", self.velocidad) 
        print("La nueva velocidad es", velocidad)
    }

    func getVelocidad() -> Int {
        return self.velocidad
    }

    func PrintPiloto() {
        print("El piloto es", self.piloto.Nombre, self.piloto.apellido)
    }
}

var avion1 = Avion(pasajeros: 100, velocidad: 1000, nombre: "Avion1", piloto: persona1)

print("El avion", avion1.nombre, "tiene", avion1.pasajeros, "pasajeros y va a", avion1.velocidad, "km/h")

avion1.mostrarVelocidad(velocidad: 2000)
avion1.frenar()

print(avion1.getVelocidad())


persona1.Nombre = "Daniel2"
persona1.mostrarNombre()
avion1.piloto.mostrarNombre()

struct Distro {
    var Nombre: String?
    var Version: String? 
}

var Distros: [Distro] = [
    Distro(Nombre: "Fedora", Version: "38"),
    Distro(Nombre: "Ubuntu", Version: "22.04"),
    Distro(Nombre: "OpenSUSE", Version: "Leap 15")
]

for distro in Distros {
    print(distro.Nombre, distro.Version)
}

var numeros: [Int] = [1,2,3,4,5,6,7,8,9,10]
