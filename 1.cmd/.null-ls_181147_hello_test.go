package hello

import (
	"testing"
)

func TestSayHello(t *testing.T){
  subtests := []struct{
    items []string
    result string
  }{
    {
    result : "Hello, world!",
  },
  {

  },
  }  

  for _, st := range subtests{
    want := st.result
    got := SayHello(st.items)

    if got != want {
      t.Errorf("want %s , got %s for %v", want, got , st.items)
    }
  }
}
