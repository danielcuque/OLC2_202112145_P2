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


// var g = 1 > 1.1

// var h = "hola" == "hola"

// var i = "ho" + "la"

// let j = 1 > false // Error

// let k = 10/0
// let l = 10%0

// let m: Float?
// var n: Int?

// Test simple if with nested if
if a {
    let testIfB = 6
	evalWhile = testIfB
    if !aa {
        evalWhile = 7
        if b < c{
            evalWhile = 8
        }
    }

    if !!a{
        evalWhile = 9
    }
}

// Test if else
var testIf = 0
if !a {
    testIf = 10
} else {
    testIf = 0
}

// Test if else

var testIfElse = 0

if a {
    testIfElse = 100
} else {
    testIfElse = 0
}

// Test if else if

var testIfElseIf = 0

var o = "b"

if b == "a" {
    let am = 10
    testIfElseIf = am
} else if o == "b" {
    testIfElseIf = 12
} else if o == "c" {
    testIfElseIf = 13
} else {
    testIfElseIf = 14
}

// Test switch

switch o {
    case "a":
        let testSwitch = 10
        testIfElseIf -= testSwitch
    case "b":
        testIfElseIf += 11
    case "c":
        testIfElseIf = 13
    default:
        testIfElseIf = 14
}

// Test While
var evalWhile = 20
var a = 0

while evalWhile > 0 {
    if evalWhile > 10 {
        while evalWhile > 10 {
            evalWhile -= 1
            break
        }
    }
    
    if evalWhile == 20 {
        break
    } else if evalWhile == 10 {
        evalWhile -= 10
        continue
    }

    switch evalWhile {
        case 20:
            let rest = 100
            evalWhile -= rest
        case 10:
            evalWhile -= 10
        default:
            evalWhile -= 1
    }
    evalWhile -= 1

}


var nums = 100

var testContinue = 0

while nums >= 0 {
    if nums % 2 == 0 {
        nums -= 1
        break
    }
    testContinue += 1
    nums -= 1
}

var testFor = 0

for i in 6...10 {
    testFor += i
    break
}

print(testFor, "testFor")

var testStringFor = "hola"

for i in testStringFor {
    let bbb = 1
    if i == "a" {
        break
    }

    testFor += bbb
}

var testGuard = 2
var num = 0

while (testGuard <= 10){
    guard testGuard % 2 == 0 else {
        print(testGuard, "es impar")
        testGuard += 1
        continue
    }
    print(testGuard, "es par")
    num += testGuard
    testGuard = testGuard + 1
}

// print(testGuard, "|", num)
print(1.0001)
print(true)
print(nil)
print("\tcadena1 \ncadena2") // mostraría cadena1 y cadena2 en líneas separadas
// Test double quote
print("cadena1 \"cadena2\" cadena3")

// Vectors

//vector con valores
var vec1: [Int] = [20, 30, 40] 
//vector vacío
var vec2: [Float] = [1.1]
//vector vacío
var vec3: [String] = []

var copiaVec: [Int] = vec1

print("La cantidad de objetos en el vector es de:", vec1.count, "está vacío:", vec1.isEmpty)
print("La cantidad de objetos en el vector es de:", vec2.count, "está vacío:", vec2.isEmpty)
print("La cantidad de objetos en el vector es de:", vec3.count, "está vacío:", vec3.isEmpty)


persona.vec1.append(50) // Error
vec1.append(50)

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

print("Debería ser 55")
print(fibonacci(10)) 

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