package constant

const(
  MAX            = 100
  internal_count = 1
)

func FooFunc(n int) int{
  return internalFunc(n)
}

func internalFunc(n int) int{
  return n + 1
}
