struct StructArr {
    var datos: Float
}

struct CentroTuristico {
    var nombre: String
}

struct Carro {
    var placa: String
    var color: String
    var tipo: String
}

struct Personaje {
    var nombre: String
    var edad: Int
    var descripcion: String
    var carro: Carro
    var numeros: StructArr
}

let newCarro = Carro(placa: "090PLO", color: "gris", tipo: "mecanico")
var nums = StructArr(datos: 2.12)

var p1 = Personaje(
    nombre: "Jose",
    edad: 18,
    descripcion: "No hace nada",
    carro: newCarro,
    numeros: nums
)

print(p1.carro.tipo, p1.numeros.datos)