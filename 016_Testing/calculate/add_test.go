package calculate

import "testing"

func TestAdd(t *testing.T) {
	have := Add(10, 10)
	expect := 20
	if have != expect {
		t.Fatalf("expected %d but have %d", 20, have)
	}

	have = Add(11, 10)
	expect = 22
	if have != expect {
		t.Fatalf("expected %d but have %d", 22, have)
	}
}
