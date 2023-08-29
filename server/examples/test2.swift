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


func negar (_ a: inout Bool) {
    a = !a
}

func negar2 (_ a: Bool) {
    a = !a
}


var b = true
print("Debería ser true -> ", String(b))  // True

print("Debería ser false ->", String(b)) // False

negar2(b)
print("Debería ser false ->", String(b)) // False
