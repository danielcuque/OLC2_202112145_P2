// let a = true
// let aa = false
// var b =  2
// var c = 1.1 + 2
// var d = -c * -1
// var e = false + 1.1
// var f = m + 1
// var evalWhile = 10
// var num: Float = 1
// var num3: Int = 10.1
// var num2: Int = false
// evalWhile = 5
// evalWhile = false
// // b = 1

// var g = 1 > 1.1

// var h = "hola" == "hola"

// var i = "ho" + "la"

// let j = 1 > false

// let k = 10/0
// let l = 10%0

// let m: Float?
// var n: Int?

// declaración de variables
var valor: String?                      //correcto, declaración sin valor

// correcto, declaración de una variable tipo Int con valor var valor1 = 10

var valor1:Int = 10.01                  // Error: no se puede asignar un Float a un Int var valor2:Float = 10.2 // correcto
var valor2_1:Float = 10 + 1             // correcto será 11.0
var valor3 = "esto es una variable";    // correcto variable tipo String

// Test += -=

var var1:Int = 10 
var1 += 10                              //var1 tendrá el valor de 20
var1 += 10.0                            // error, no puede asignar un valor de tipo Float a un Int

var var2:Float = 0.0
var2 += 10                              //var2 tendrá el valor de 10.0
var2 += 10.0                            //var2 tendrá el valor de 20.0

var str:String = "cad"
str += "cad"                            //str tendrá el valor de "cadcad"
str += 10                               //operación inválida String + Int

// if (a) {
// 	evalWhile = 5
// } else {
//     evalWhile = 10
// }

// while (b < evalWhile) {
// 	b = b + 1
// }