func fibonacci(n: Int) -> Int {
    if n <= 1 {
        return n
    } else {
        return fibonacci(n: n - 1) + fibonacci(n: n - 2)
    }
}

print("Fibonacci de 10:")
print(fibonacci(n: 10))
