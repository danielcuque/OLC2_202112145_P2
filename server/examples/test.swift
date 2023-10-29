// let notas: [[Int]] = [
//     [53, 88, 95, 89, 75],
//     [81, 51, 57, 67, 93],
//     [94, 74, 58, 84, 100]
// ]

// print(notas[2][4]);


// var stringArray: [String] = ["ab", "cd", "ef", "gh", "ij"]
// for i in 0...stringArray.count-1 {
//     print(stringArray[i])
// }

// var arr1: [Int] = [8, 4, 6, 2]
// var arr2: [Int] = [40, 21, 1, 3, 14, 4]
// var arr3: [Int] = [90, 3, 40, 10, 8, 5]

// printArray("arr1: ", arr1)
// arr1.append(9)
// printArray("arr1: ", arr1)

// printArray("arr2: ", arr2)
// arr2.removeLast()
// printArray("arr2: ", arr2)

// printArray("arr3: ", arr3)
// arr3.remove(at: 3)
// printArray("arr3: ", arr3)

// func printArray(_ msg: String, _ arr: [Int]) {
//     print(arr.count)
//     print(arr.isEmpty)
//     var out = ".["
//     for i in 0...arr.count-1 {
//         if i == arr.count - 1 {
//             out = out + String(arr[i])
//         } else {
//             out = out + String(arr[i]) + ", "
//         }
//     }
//     out = out + "]."
//     print(msg + out)
// }

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
