package main

import "unicode/utf8"

var Hx_Obj_haxe_iterators_stringiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.iterators.StringIterator",
)

type Hx_Obj_VTable_haxe_iterators_stringiterator interface {
    Hx_Field_next() int
    Hx_Field_hasNext() bool
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_stringiterator struct {
    VTable Hx_Obj_VTable_haxe_iterators_stringiterator
    Hx_Field_offset int
    Hx_Field_s string
}

func Hx_Obj_haxe_iterators_stringiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_stringiterator {
    obj := &Hx_Obj_haxe_iterators_stringiterator{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_iterators_stringiterator_CreateInstance(s string) *Hx_Obj_haxe_iterators_stringiterator {
    obj := Hx_Obj_haxe_iterators_stringiterator_CreateEmptyInstance()
    obj.Hx_New(s)
    return obj
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_New(s string) {
    this.Hx_Field_s = s
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_iterators_stringiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_Field_hasNext() bool {
    var _hx_tmp_0 int = this.Hx_Field_offset; _ = _hx_tmp_0
    return (_hx_tmp_0 < utf8.RuneCountInString(this.Hx_Field_s))
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_Field_next() int {
    var _hx_tmp_0 any = Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_s, "cca"); _ = _hx_tmp_0
    var _hx_tmp_1 int = this.Hx_Field_offset; _ = _hx_tmp_1
    this.Hx_Field_offset = (this.Hx_Field_offset + 1)
    return ((int)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, &([]any{ ((any)(_hx_tmp_1)) })))))
}
