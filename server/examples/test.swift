// var arr1: [Int] = [8, 4, 6, 2]
// print("arr1 vacío: ", arr1.isEmpty)
// print("arr1 cantidad: ", arr1.count)

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