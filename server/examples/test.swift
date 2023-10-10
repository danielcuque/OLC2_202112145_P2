var a = -1
while a < 5 {
    a = a + 1
    if a == 3 {
        print("a")
        continue
    } else if a == 4 {
        print("b")
        break
    }
    print("El valor de a es:", a)
}