Recursion is a practice of a function calling itself. When using recursion you
need a stop point. An example is:

```go
func main() {
    n := factorial(4)
    fmt.Println(n)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

This function runs and with each run reduces n. The if statement breakes the
recursion when n becomes zero. This function will recure and expand to become:

```go
return n * factorial(4-1) * factorial(3-1) * factorial(2-1) * 1
```

And once expanded to that in memory it is calculated and returned.
