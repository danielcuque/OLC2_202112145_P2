// var arr1: [Int] = [8, 4, 6, 2]
// print("arr1 vacío: ", arr1.isEmpty)
// print("arr1 cantidad: ", arr1.count)

func factorial(_ n: Int) -> Int {
    if n < 2 {
        return 1
    } else {
        return n * factorial(n - 1)
    }
}

print(factorial(6))
print(factorial(4))
print(factorial(3))

func Hanoi(_ discos: Int, _ origen: Int, _ auxiliar: Int, _ destino: Int) {
    if discos == 1 {
        print("Mover disco de", origen, "a", destino)
    } else {
        Hanoi(discos - 1, origen, destino, auxiliar)
        print("Mover disco de", origen, "a", destino)
        Hanoi(discos - 1, auxiliar, origen, destino)
    }
}

print("Hanoi")
Hanoi(3, 1, 2, 3)