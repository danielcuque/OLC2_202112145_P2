struct Persona{
    var Nombre: String
    var edad = 0 
}

struct Avion {
    var pasajeros = 0 var velocidad = 100
    var nombre: String
    var piloto: Persona
    // metodo dentro de struct 
    mutating func frenar(){
        print("Frenando")
    //al ser mutable sí afecta al struct self.velocidad = 0
    }
// funcion inmutable 
    func mostrarVelocidad(){
        print("Velocidad",self.velocidad) 
    }
}

var avioneta = Avion(nombre: "78496", piloto: Persona(Nombre: "Joel", edad: 43 ))

print("Pasajeros", avioneta.pasajeros)
// modificion de un atributo
avioneta.pasajeros = 5

print("Nuevos pasajeros", avioneta.pasajeros)

avioneta.mostrarVelocidad()
// copia de structs por valor
var avioneta2 = avioneta
avioneta2.pasajeros = 0
//imprime: avioneta.pasajeros: 5 print("avioneta.pasajeros:",avioneta.pasajeros) //avioneta2.pasajeros 0 print("avioneta2.pasajeros:",avioneta2.pasajeros)
print("avioneta.piloto.Nombre:",avioneta2.piloto.Nombre )

struct Fruta{
    let nombre: String = "pera" 
    var precio: Int
}

var pera = Fruta(precio: 10)

struct Verdura{
    let nombre: String
    var precio: Int 
}


var brocoli = Verdura(nombre:"brocoli", precio: 5)

struct Person {
    var name: String 
    var age: Int
}

var personas: [Persona] = []

personas.append(Persona(Nombre: "Celeste", edad: 23)) 
personas.append(Persona(Nombre: "Roel", edad: 32)) 
personas.append(Persona(Nombre: "Flor", edad: 17))

var persona1 = personas[0] 
persona1.Nombre = "Nancy"

print("Debería ser Nancy", persona1.Nombre) //imprime Nancy 
print("Debería ser Celeste", personas[0].Nombre) //imprime Celeste 

// personas[1].edad = 26 // se modifica un array


struct Distro {
    var Nombre: String
    var Version: String 
}

var Distros: [Distro] = [
    Distro(Nombre: "Ubuntu", Version: "22.04"), 
    Distro(Nombre: "Fedora", Version: "38"), 
    Distro(Nombre: "OpenSUSE", Version: "Leap 15")
]

print(Distros[0].Nombre) // Imprime Ubuntu 
print(Distros[1].Version) // Imprime 13

for distro in Distros {
    print(distro.Nombre, distro.Version)
}

func crearVerdura( precioV: Int, nombreV: String ) -> Verdura {
    return Verdura( nombre: nombreV, precio: precioV ) 
}

var apio:Verdura = crearVerdura( precioV: 10, nombreV: "Apio")

print(apio.nombre, apio.precio)
