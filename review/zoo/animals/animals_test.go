package animals

import(
  "testing"
)

func TestElephantFeed(t *testing.T){
  expect := "Grass"
  actual := ElephantFeed()

  if expect != actual{
    t.Errorf("%s != %s", expect, actual)
  }
}

func TestMonkeyFeed(t *testing.T){
  expect := "Banana"
  actual := MonkeyFeed()

  if expect != actual{
    t.Errorf("%s != %s", expect, actual)
  }
}

func TestRabbitFeed(t *testing.T){
  expect := "Carrot"
  actual := RabbitFeed()

  if expect != actual{
    t.Errorf("%s != %s", expect, actual)
  }
}

// ➜  zoo git:(master) ✗ go test ./animals -v
// === RUN   TestElephantFeed
// --- PASS: TestElephantFeed (0.00s)
// === RUN   TestMonkeyFeed
// --- PASS: TestMonkeyFeed (0.00s)
// === RUN   TestRabbitFeed
// --- PASS: TestRabbitFeed (0.00s)
// PASS
// ok    _/Users/umeki/Desktop/Go/GoAlgorithm/review/zoo/animals 0.005s

