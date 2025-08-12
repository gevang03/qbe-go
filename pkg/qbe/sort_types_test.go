package qbe_test

import (
	"strings"
	"testing"

	"github.com/gevang03/qbe-go/pkg/qbe"
)

func TestTypeOrder(t *testing.T) {
	mod := qbe.NewModule("type_order")
	foo := mod.DefineStruct("foo")
	bar := mod.DefineStruct("bar")
	{
		barFields := qbe.NewFieldList()
		barFields.InsertField(foo, 1)
		bar.SetFields(barFields)
	}
	baz := mod.DefineStruct("baz")
	{
		bazFields := qbe.NewFieldList()
		bazFields.InsertField(bar, 2)
		baz.SetFields(bazFields)
	}
	sb := strings.Builder{}
	mod.ToIL(&sb)
	output := sb.String()
	fooIndex := strings.Index(output, "type :foo")
	if fooIndex == -1 {
		t.Fatal("foo type must be emitted")
	}
	barIndex := strings.Index(output, "type :bar")
	if barIndex == -1 {
		t.Fatal("bar type must be emitted")
	}
	bazIndex := strings.Index(output, "type :baz")
	if bazIndex == -1 {
		t.Fatal("baz type must be emitted")
	}

	if fooIndex >= barIndex {
		t.Fatal("foo must be emitted before bar")
	}
	if barIndex >= bazIndex {
		t.Fatal("bar must be emitted before baz")
	}
}
