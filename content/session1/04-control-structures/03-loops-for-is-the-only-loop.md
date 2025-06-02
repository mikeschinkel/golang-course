# Loops (`for` is the only loop)

```go
// Traditional for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style loop
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// Infinite loop
for {
    // Do something forever
    if condition {
        break
    }
}

// Range over slice
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Range with index only
for index := range numbers {
    fmt.Println(index)
}

// Range with value only
for _, value := range numbers {
    fmt.Println(value)
}

// Range over map
m := map[string]int{"a": 1, "b": 2, "c": 3}
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Range over string (runes)
for index, runeValue := range "Hello, 世界" {
    fmt.Printf("%d: %c\n", index, runeValue)
}

// Range over channel
ch := make(chan int)
go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}()

for value := range ch {
    fmt.Println(value)
}
```