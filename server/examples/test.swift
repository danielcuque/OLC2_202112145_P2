// let numero = 1
// switch numero {
//     case 1:
//         print("Uno")
//     case 2:
//         print("Dos")
//     case 3:
//         print("Tres")
//     default:
//         print("Otro")
// }

i = 2
while (i <= 10) {
    // la sentencia guard verifica si i es impar
    guard i % 2 == 0 else {
        i = i + 1
        continue
    }
    print(i)
    i = i + 1
}
