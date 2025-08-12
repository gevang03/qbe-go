package qbe_test

import (
	"testing"

	"github.com/gevang03/qbe-go/pkg/qbe"
)

func TestOffsetOf(t *testing.T) {
	mod := qbe.NewModule("offset_of")
	foo := mod.DefineStruct("foo")
	fooFields := qbe.NewFieldList()
	f0 := fooFields.InsertField(qbe.ByteType(), 1)
	f1 := fooFields.InsertField(qbe.WordType(), 1)
	foo.SetFields(fooFields)
	idx0 := foo.OffsetOf(f0)
	if idx0 != 0 {
		t.Fatalf("first field of struct is at offset %#v", idx0)
	}
	idx1 := foo.OffsetOf(f1)
	if idx1 != 1 {
		t.Fatalf("field after a byte at offset %#v", idx1)
	}
}

func TestOffsetOfAligned(t *testing.T) {
	mod := qbe.NewModule("offset_of")
	foo := mod.DefineStruct("foo")
	fooFields := qbe.NewFieldList()
	f0 := fooFields.InsertField(qbe.ByteType(), 1)
	f1 := fooFields.InsertFieldAligned(qbe.WordType(), 1)
	foo.SetFields(fooFields)
	idx0 := foo.OffsetOf(f0)
	if idx0 != 0 {
		t.Fatalf("first field of struct is at offset %#v", idx0)
	}
	idx1 := foo.OffsetOf(f1)
	if idx1 != 4 {
		t.Fatalf("aligned word field after a byte at offset %#v", idx1)
	}
}

func TestAlign(t *testing.T) {
	mod := qbe.NewModule("align")
	foo := mod.DefineStruct("foo")
	fooFields := qbe.NewFieldList()
	fooFields.InsertField(qbe.LongType(), 1)
	foo.SetFields(fooFields)
	fooAlign := foo.AlignOf()
	if fooAlign != qbe.LongType().AlignOf() {
		t.Fatalf("alignment of default aligned struct is max field alignment, got %#v", foo.Align)
	}
}
