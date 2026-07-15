package main

import "unicode/utf8"

var Hx_Obj_haxe_iterators_stringkeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.iterators.StringKeyValueIterator",
)

type Hx_Obj_VTable_haxe_iterators_stringkeyvalueiterator interface {
    Hx_Field_next() any
    Hx_Field_hasNext() bool
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_stringkeyvalueiterator struct {
    VTable Hx_Obj_VTable_haxe_iterators_stringkeyvalueiterator
    Hx_Field_offset int
    Hx_Field_s string
}

func Hx_Obj_haxe_iterators_stringkeyvalueiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_stringkeyvalueiterator {
    obj := &Hx_Obj_haxe_iterators_stringkeyvalueiterator{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_iterators_stringkeyvalueiterator_CreateInstance(s string) *Hx_Obj_haxe_iterators_stringkeyvalueiterator {
    obj := Hx_Obj_haxe_iterators_stringkeyvalueiterator_CreateEmptyInstance()
    obj.Hx_New(s)
    return obj
}

func (this *Hx_Obj_haxe_iterators_stringkeyvalueiterator) Hx_New(s string) {
    this.Hx_Field_s = s
}

func (this *Hx_Obj_haxe_iterators_stringkeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_iterators_stringkeyvalueiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_stringkeyvalueiterator) Hx_Field_hasNext() bool {
    var _hx_tmp_0 int = this.Hx_Field_offset; _ = _hx_tmp_0
    return (_hx_tmp_0 < utf8.RuneCountInString(this.Hx_Field_s))
}

func (this *Hx_Obj_haxe_iterators_stringkeyvalueiterator) Hx_Field_next() any {
    var _hx_tmp_0 any = ((any)(this.Hx_Field_offset)); _ = _hx_tmp_0
    var _hx_tmp_1 string = this.Hx_Field_s; _ = _hx_tmp_1
    var _hx_tmp_2 int = this.Hx_Field_offset; _ = _hx_tmp_2
    this.Hx_Field_offset = (this.Hx_Field_offset + 1)
    return any(map[string]any{ "key": _hx_tmp_0, "value": ((any)(Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_1, _hx_tmp_2).Value)) })
}
