package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	p := NewParser(strings.NewReader("hp>100 mana<=50 name=Morte"))

	res, err := p.Parse()

	if err != nil {
		t.Errorf("Parser failed on valid string: %s", err)
	}
	fmt.Printf("%v", res)
}
