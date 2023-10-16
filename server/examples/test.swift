for i in 0...9 {
    var output = " "
    for j in 0...(10 - i) {
        output = output + "|"
    }
    for k in 0...i {
        output = output + "*"
    }
    print(output)
}