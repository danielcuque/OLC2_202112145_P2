// declaración de variables
var valor: String?                      //correcto, declaración sin valor

// correcto, declaración de una variable tipo Int con valor var valor1 = 10

var valor1:Int = 10.01                  // Error: no se puede asignar un Float a un Int 
var valor2:Float = 10.2                 // correcto
var valor2_1:Float = 10 + 1             // correcto será 11.0
var valor3 = "esto es una variable";    // correcto variable tipo String

var char:Character = 'A';               //correcto variable tipo Character 
var valor4:Bool = true                  //correcto

// debe ser un error ya que los tipos no son compatibles 
var valor41:String = true

// debe ser un error ya que existe otra variable valor3 definida previamente
var valor3:Int = 10

// var .58 = 4;                         // debe ser error porque .58 no es un nombre válido
// var if = "10"                        // debe ser sun error porque "if" es una palabra reservada // ejemplo de asignaciones
valor1 = 200                            // correcto
valor3 = "otra cadena"                  // correcto
valor4 = 10                             // error tipos incompatibles (Bool, Int)
valor2 = 200                            // error tipos incompatibles (Float, Int)
char = "otra cadena"                    // error tipos incompatibles (Character, String)

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

var i = 10 // variable global es accesible desde este ámbito 
if i == 10 {
    var j:Int = 10 + i // i es accesible desde este ámbito 
    if i==10 {
        var k:Int=j+1 //iyjsonaccesiblesdesdeesteámbito 
    }
    j = k // error k ya no es accesible en este ámbito 
}

func suma( num1 x : Int, num2 y: Int) -> Int {
    return x + y
}

func resta(_ x : Int, _ y: Int) -> Int {
    return x - y
}

//función mul
// Nombres externos: x, y
// Nombres internos: x, y
func mul(x: Int, y: Int) -> Int {
    return x * y
}

// // func duplicar(_ x: inout Int){
// //     x += x
// // }

// // func duplicarA (_ array: inout [Int] ) {
// //     var i = 0
// //     while (i < array.count ) {
// //         array[i] += array[i]
// //         i += 1
// //     }
// // }


var numero1 = 3
var numero2 = 2
//llamada con nombres externos
print("La suma de", numero1, "y", numero2, "es ->", suma(num1: numero1, num2: numero2))
//llamada sin nombres externos
print(resta(numero1, numero2))
//llamada con nombres externos e internos idénticos
print(mul(x: numero1, y: numero2))



func func1() -> Int  {
    return 1
}

print(Int("10"))
print(Int("12.001"))
print(Int(10.9999))
print(Int("Q10"))

print(Float("10"))
print(Float("10.001"))
print(Float("Q10.00"))


print(String(10) + String(3.51)) //imprime 103.5000 
print( String( true )) //true
var cadena = String(true) + "->" + String(3.504) 
print(cadena) // imprime true->3.50400000