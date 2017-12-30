package main

import(
  "testing"
)

func TestAppName(t *testing.T){
  expect := "Zoo Application"
  actual := AppName()

  if expect != actual{
    t.Errorf("%s == %s", expect, actual)
  }
}

// ➜  zoo git:(master) ✗ go test -v
// === RUN   TestAppName
// --- PASS: TestAppName (0.00s)
// PASS
// ok    _/Users/umeki/Desktop/Go/GoAlgorithm/review/zoo 0.005s
