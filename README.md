<img src="https://user-images.githubusercontent.com/16008095/215646733-8fcc1bc6-d2e8-4578-904e-866c8315d943.png" width=200px align=right>

# go-reacto &nbsp; <img height="20px" src="https://img.shields.io/badge/Golang-FFFFFF?logo=go&style=flat">
Reactive in Go
<sub><sub><em>or something</em></sub></sub>

## Usages

```go
p := goreacto.Publisher[int]{}

p.Subscribe(func(i int) { println(i) })

mapped_p := goreacto.Map(&p, func(i int) string {
  return fmt.Sprintf("map block got value: %d", i)
})

mapped_p.Subscribe(func(s string) { println(s) })

p.Publish(1)
p.Publish(2)
p.Publish(3)

// would output
// 1
// map block got value: 1
// 2
// map block got value: 2
// 3
// map block got value: 3
```
