package main

import "unicode/utf8"

var Hx_Obj_stringbuf_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "StringBuf",
)

type Hx_Obj_VTable_stringbuf interface {
    Hx_Field_toString() string
    Hx_Field_get_length() int
    Hx_Field_clear()
    Hx_Field_addSub(s string, pos int, _hx_reserved_len struct { Value int; Valid bool })
    Hx_Field_addChar(c int)
    Hx_Field_add(x any)
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_stringbuf struct {
    VTable Hx_Obj_VTable_stringbuf
    Hx_Field_b string
}

func Hx_Obj_stringbuf_CreateEmptyInstance() *Hx_Obj_stringbuf {
    obj := &Hx_Obj_stringbuf{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_stringbuf_CreateInstance() *Hx_Obj_stringbuf {
    obj := Hx_Obj_stringbuf_CreateEmptyInstance()
    obj.Hx_New()
    return obj
}

func (this *Hx_Obj_stringbuf) Hx_New() {
    this.Hx_Field_b = ""
}

func (this *Hx_Obj_stringbuf) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_stringbuf_RTTI
}

func (this *Hx_Obj_stringbuf) Hx_Field_get_length() int {
    return utf8.RuneCountInString(this.Hx_Field_b)
}

func (this *Hx_Obj_stringbuf) Hx_Field_add(x any) {
    var _hx_tmp_0 string = this.Hx_Field_b; _ = _hx_tmp_0
    this.Hx_Field_b = (_hx_tmp_0 + Hx_Field_std_string(x))
}

func (this *Hx_Obj_stringbuf) Hx_Field_addChar(c int) {
    var _hx_tmp_0 string = this.Hx_Field_b; _ = _hx_tmp_0
    this.Hx_Field_b = (_hx_tmp_0 + Hx_Field_go_haxe_hxstring_fromCharCode(c))
}

func (this *Hx_Obj_stringbuf) Hx_Field_addSub(s string, pos int, _hx_reserved_len struct { Value int; Valid bool }) {
    var _hx_tmp_0 string = this.Hx_Field_b; _ = _hx_tmp_0
    var _hx_tmp_1 string; _ = _hx_tmp_1
    if ((_hx_reserved_len.Valid == false)) {
        var _hx_tmp_2 string = s; _ = _hx_tmp_2
        var _hx_tmp_3 int = pos; _ = _hx_tmp_3
        var _hx_tmp_4 string = _hx_tmp_2; _ = _hx_tmp_4
        var _hx_tmp_5 int = _hx_tmp_3; _ = _hx_tmp_5
        _hx_tmp_1 = Hx_Field_go_haxe_hxstring_substr(_hx_tmp_4, _hx_tmp_5, struct { Value int; Valid bool }{})
    } else {
        _hx_tmp_1 = Hx_Field_go_haxe_hxstring_substr(s, pos, _hx_reserved_len)
    }

    this.Hx_Field_b = (_hx_tmp_0 + _hx_tmp_1)
}

func (this *Hx_Obj_stringbuf) Hx_Field_clear() {
    this.Hx_Field_b = ""
}

func (this *Hx_Obj_stringbuf) Hx_Field_toString() string {
    return this.Hx_Field_b
}
