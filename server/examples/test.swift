struct StructArr {
    var datos: Float
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
var nums = StructArr(datos: 0.0)

var p1 = Personaje(
    nombre: "Jose",
    edad: 18,
    descripcion: "Nada",
    carro: newCarro,
    numeros: nums
)

var nums2 = StructArr(datos: Float("23.43"))

print(p1.numeros.datos)
p1.numeros = nums2
print(p1.numeros.datos)

/*
--------------------------
----------STRUCT----------
----------12 pts----------
--------------------------
 
=============================================
================DEFINICIÓN===================
=============================================
 
=============================================
==============INSTANCIACIÓN==================
=============================================
 
=============================================
========ASIGNACIÓN Y ACCESO==================
=============================================
El nombre del Centro turistico 1 es:  Volcan de pacaya
El nombre del Centro turistico 2 es:  Rio dulce
El nombre del Centro turistico 3 es:  Laguna Luchoa
El nombre del Centro turistico 4 es:  Playa Blanca
El nombre del Centro turistico 5 es:  Antigua Guatemala
El nombre del Centro turistico 6 es:  Lago de Atitlan
Persona nombre:  Jose , edad:  18 , carroTipo:  mecanico , numeros:  0.0
Persona nombre:  Jose , edad:  18 , carroTipo:  mecanico , nuevos numeros:  23.43
*/