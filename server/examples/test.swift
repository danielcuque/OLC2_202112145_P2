let a = true
let aa = false
var b =  2
var c = 1.1 + 2
var d = -c * -1
var e = false + 1.1
var f = m + 1
var evalWhile = 10
var num: Float = 1
var num3: Int = 10.1
var num2: Int = false
evalWhile = 5
evalWhile = false
// b = 1

var g = 1 > 1.1

var h = "hola" == "hola"

var i = "ho" + "la"

let j = 1 > false

let k = 10/0
let l = 10%0

let m: Float?
var n: Int?

// Test simple if with nested if
if a {
    let testIfB = 6
	evalWhile = testIfB
    if !aa {
        evalWhile = 7
        if b < c{
            evalWhile = 8
        }
    }

    if !!a{
        evalWhile = 9
    }
}

// Test if else
var testIf = 0
if !a {
    testIf = 10
} else {
    testIf = 0
}

// Test if else

var testIfElse = 0

if a {
    testIfElse = 100
} else {
    testIfElse = 0
}

// Test if else if

var testIfElseIf = 0

var o = "b"

if b == "a" {
    let am = 10
    testIfElseIf = am
} else if o == "b" {
    testIfElseIf = 12
} else if o == "c" {
    testIfElseIf = 13
} else {
    testIfElseIf = 14
}

// Test switch

switch o {
    case "a":
        let testSwitch = 10
        testIfElseIf -= testSwitch
    case "b":
        testIfElseIf += 11
    case "c":
        testIfElseIf = 13
    default:
        testIfElseIf = 14
}

// Test While
var evalWhile = 20
var a = 0

while evalWhile > 0 {
    if evalWhile > 10 {
        while evalWhile > 10 {
            evalWhile -= 1
            break
        }
    }
    
    if evalWhile == 20 {
        break
    } else if evalWhile == 10 {
        evalWhile -= 10
        continue
    }

    switch evalWhile {
        case 20:
            let rest = 100
            evalWhile -= rest
        case 10:
            evalWhile -= 10
        default:
            evalWhile -= 1
    }
    evalWhile -= 1

}

/*
1. While {
    validProps: ["break", "continue"]
}
2. While {
    validProps: ["break", "continue"]
}
*/
var nums = 100

var testContinue = 0

while nums >= 0 {
    if nums % 2 == 0 {
        nums -= 1
        break
    }
    testContinue += 1
    nums -= 1
}

var testFor = 0

for i in 6...10 {
    testFor += i
    break
}

var testStringFor = "hola"

for i in testStringFor {
    let bbb = 1
    if i == "a" {
        break
    }

    testFor += bbb
}