package main

import "testing"

type CustomTable struct{}

func (c CustomTable) TableName() string {
	return "Fred"
}

func TestTableName(t *testing.T) {

	type User struct{}
	type Something struct{}
	type Assignments struct{}

	for _, v := range []struct {
		val    interface{}
		expect string
	}{
		{
			val:    &User{},
			expect: "User",
		},
		{
			val:    &Something{},
			expect: "Something",
		},
		{
			val:    &Assignments{},
			expect: "Assignments",
		},
		{
			val:    &CustomTable{},
			expect: "Fred",
		},
	} {
		s := Scope{Value: v.val}
		if name := s.TableName(); name != v.expect {
			t.Errorf("Expected %s, got %s", v.expect, name)
		}

	}
}
