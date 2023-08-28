// func imprimirArray (_ simpleArray: [Int] ) {
//     for i in 0...simpleArray.count - 1 {
//         print(simpleArray[i])
//     }
// }

// func duplicarConPuntero1 (_ array1: inout [Int] ) {
//     var i = 0
//     while (i < array1.count ) {
//         array1[i] += array1[i]
//         i += 1
//     }
// }

// func duplicarConPuntero2 (_ array2: inout [Int] ) {
//     duplicarConPuntero1(&array2)
// }

// func duplicarSinPuntero (_ arrayS: [Int] ) {
//     for i in 0...arrayS.count - 1 {
//         arrayS[i] += arrayS[i]
//     }
// }


// var array: [Int] = [1,2,3,4,5,6]

// print("Array original")
// imprimirArray(array)

// print("Array sin puntero")
// duplicarSinPuntero(array)
// imprimirArray(array)

// print("Array con puntero")
// duplicarConPuntero1(&array)
// imprimirArray(array)

// print("Array con puntero doble")
// duplicarConPuntero2(&array)
// imprimirArray(array)


// func negar (_ a: inout Bool) {
//     a = !a
// }

// func negar2 (_ a: Bool) {
//     a = !a
// }


// var b = true
// print("Debería ser true -> ", String(b))  // True

// negar(&true)
// print("Debería ser false ->", String(b)) // False

// negar2(b)
// print("Debería ser false ->", String(b)) // False

// func fibonacci(_ n: Int) -> Int {
//     if n > 1 {
//         return fibonacci(n - 1) + fibonacci(n - 2)
//     } else if n == 1 {
//         return 1
//     } else if n == 0 {
//         return 0
//     } else {
//         print("error")
//         return 0
//     }
// }

// func sumar2 (_ a: inout Int, _ b: inout Int) {
//     a = a + b
// }

// let fibonacci5 = 5
// sumar2(&fibonacci5, &fibonacci5)

// print("\nDebería ser 55")
// print(fibonacci(fibonacci5)) 
// print("El valor de fibonacci10 es", fibonacci5)

// func append(_ a: Int) {
//     print(a)
// }

// append(1)


// Matrix declaration

var vec1: [Int] = [20, 30, 40] 

var mtx1: [[Int]] = [[1], [4, 5], [7, 8, 9]]

var mtx2 : [[[Int]]] = [[[1,2,3],[4,5,6],[7,8,9]], 
                        [[10,11,12],[13,14,15],[16,17,18]], 
                        [[19,20,21],[22,23,24],[25,26,27]]]

// Using repeating

// Verificar tipo de matriz declarada y creada
// Verificar que tenga el mismo tamaño que el especificado
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