// let notasBuenas: [[Int]] = [
//     [53, 88, 95, 89, 75],
//     [81, 51, 57, 67, 93],
//     [94, 74, 58, 84, 100]
// ]

var arr1: [Int] = [8, 4, 6, 2]
// var arr2: [Int] = []
// print("arr1 cantidad: ", arr1.count)
// print("arr1 vacío: ", arr1.isEmpty)
// for i in 0...arr1.count-1 {
//     print(arr1[i])
// }

func printArray(_ msg: String, _ arr: [Int]) {
    var out = ".["
    for i in 0...arr.count-1 {
        print("i: ", i)
        if i == arr.count - 1 {
            out = out + String(arr[i])
        } else {
            out = out + String(arr[i]) + ", "
        }
    }
    out = out + "]."
    print(msg + out)
}

printArray("arr1: ", arr1)

// print("arr2 cantidad: ", arr2.count)
// print("arr2 vacío: ", arr2.isEmpty)

// func factorial(_ n: Int) -> Int {
//     if n < 2 {
//         return 1
//     } else {
//         return n * factorial(n - 1)
//     }
// }

// print(factorial(suma(3,3)))
// print(factorial(4))
// print(factorial(3))

// func Hanoi(_ discos: Int, _ origen: Int, _ auxiliar: Int, _ destino: Int) {
//     if discos == 1 {
//         print("Mover disco de", origen, "a", destino)
//     } else {
//         Hanoi(discos - 1, origen, destino, auxiliar)
//         print("Mover disco de", origen, "a", destino)
//         Hanoi(discos - 1, auxiliar, origen, destino)
//     }
// }

// print("Hanoi")
// Hanoi(3, 1, 2, 3)

// func suma(_ a: Int, _ b: Int) -> Int {
//     var rs = a+b
//     return rs
// }

// print(suma(1,2))
// print(suma(1,3))