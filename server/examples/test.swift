var num1 = 10; 
var num2 = 20; 
var num3 = 30; 

var arr1: [Int] = [8, 4, 6, 2]

for i in 0...3 {
    arr1[i] = arr1[i] * 2
    print(arr1[i])
}


func suma(_ a: Int, _ b: Int) {
    var c = a + b
    print(c)
}


suma(num1, num3)

