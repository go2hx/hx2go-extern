package main

import "strings"
import "unicode/utf8"

var Hx_Obj_go_haxe_hxstring_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "go.haxe.HxString",
)

type Hx_Obj_VTable_go_haxe_hxstring interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxstring struct {
    VTable Hx_Obj_VTable_go_haxe_hxstring
}

func Hx_Obj_go_haxe_hxstring_CreateEmptyInstance() *Hx_Obj_go_haxe_hxstring {
    obj := &Hx_Obj_go_haxe_hxstring{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_go_haxe_hxstring_CreateInstance() *Hx_Obj_go_haxe_hxstring {
    obj := Hx_Obj_go_haxe_hxstring_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_go_haxe_hxstring) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_go_haxe_hxstring_RTTI
}

func Hx_Field_go_haxe_hxstring_toUpperCase(s string) string {
    return strings.ToUpper(s)
}

func Hx_Field_go_haxe_hxstring_toLowerCase(s string) string {
    return strings.ToLower(s)
}

func Hx_Field_go_haxe_hxstring_charAt(s string, index int) string {
    var _hx_tmp_0 string = Hx_Field_go_haxe_hxstring_toLowerCase(""); _ = _hx_tmp_0
    return (_hx_tmp_0 + Hx_Field_std_string(string(([]rune)(s)[index])))
}

func Hx_Field_go_haxe_hxstring_charCodeAt(s string, index int) struct { Value int; Valid bool } {
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if ((index < 0)) {
        _hx_tmp_0 = true
    } else {
        var _hx_tmp_1 int = index; _ = _hx_tmp_1
        var _hx_tmp_2 int = _hx_tmp_1; _ = _hx_tmp_2
        _hx_tmp_0 = (_hx_tmp_2 > utf8.RuneCountInString(s))
    }

    if (_hx_tmp_0) {
        return struct { Value int; Valid bool }{}
    }

    var code int32 = ([]rune)(s)[index]; _ = code
    return struct { Value int; Valid bool }{ Value: ((int)(code)), Valid: true }
}

func Hx_Field_go_haxe_hxstring_indexOf(s string, str string, startIndex struct { Value int; Valid bool }) int {
    if ((startIndex.Valid != false)) {
        var _hx_tmp_0 int = startIndex.Value; _ = _hx_tmp_0
        if ((_hx_tmp_0 >= utf8.RuneCountInString(s))) {
            return -1
        }
    
        var _hx_tmp_1 string = s; _ = _hx_tmp_1
        s = string(([]rune)(_hx_tmp_1)[startIndex.Value:])
    }

    return strings.Index(s, str)
}

func Hx_Field_go_haxe_hxstring_lastIndexOf(s string, str string, startIndex struct { Value int; Valid bool }) int {
    if ((startIndex.Valid != false)) {
        var _hx_tmp_0 int = startIndex.Value; _ = _hx_tmp_0
        if ((_hx_tmp_0 >= utf8.RuneCountInString(s))) {
            return -1
        }
    
        var _hx_tmp_1 string = s; _ = _hx_tmp_1
        s = string(([]rune)(_hx_tmp_1)[startIndex.Value:])
    }

    return strings.LastIndex(s, str)
}

func Hx_Field_go_haxe_hxstring_split(s string, delimiter string) *[]string {
    var x []string = strings.Split(s, delimiter); _ = x
    var self []string = x; _ = self
    return (&self)
}

func Hx_Field_go_haxe_hxstring_substr(s string, pos int, _hx_reserved_len struct { Value int; Valid bool }) string {
    var _hx_tmp_0 int; _ = _hx_tmp_0
    if ((pos < 0)) {
        _hx_tmp_0 = (utf8.RuneCountInString(s) + pos)
    } else {
        _hx_tmp_0 = pos
    }

    var startIndex int = _hx_tmp_0; _ = startIndex
    if ((startIndex < 0)) {
        startIndex = 0
    }

    var _hx_tmp_1 int; _ = _hx_tmp_1
    if ((_hx_reserved_len.Valid != false)) {
        var _hx_tmp_2 int = startIndex; _ = _hx_tmp_2
        var _hx_tmp_3 int = _hx_tmp_2; _ = _hx_tmp_3
        _hx_tmp_1 = (_hx_tmp_3 + _hx_reserved_len.Value)
    } else {
        _hx_tmp_1 = utf8.RuneCountInString(s)
    }

    var end int = _hx_tmp_1; _ = end
    var _hx_tmp_2 bool; _ = _hx_tmp_2
    if ((_hx_reserved_len.Valid != false)) {
        _hx_tmp_2 = (_hx_reserved_len.Value < 0)
    } else {
        _hx_tmp_2 = false
    }

    if (_hx_tmp_2) {
        return ""
    }

    var _hx_tmp_3 int = end; _ = _hx_tmp_3
    if ((_hx_tmp_3 > utf8.RuneCountInString(s))) {
        end = utf8.RuneCountInString(s)
    }

    return string(([]rune)(s)[startIndex:end])
}

func Hx_Field_go_haxe_hxstring_substring(s string, startIndex int, endIndex struct { Value int; Valid bool }) string {
    if ((endIndex.Valid == false)) {
        endIndex = struct { Value int; Valid bool }{ Value: utf8.RuneCountInString(s), Valid: true }
    }

    if ((startIndex < 0)) {
        startIndex = 0
    }

    if ((endIndex.Value < 0)) {
        endIndex = struct { Value int; Valid bool }{ Value: 0, Valid: true }
    }

    var _hx_tmp_0 int = startIndex; _ = _hx_tmp_0
    if ((_hx_tmp_0 > endIndex.Value)) {
        var tmp struct { Value int; Valid bool } = endIndex; _ = tmp
        endIndex = struct { Value int; Valid bool }{ Value: startIndex, Valid: true }
    
        startIndex = tmp.Value
    }

    var _hx_tmp_1 int = endIndex.Value; _ = _hx_tmp_1
    if ((_hx_tmp_1 > utf8.RuneCountInString(s))) {
        endIndex = struct { Value int; Valid bool }{ Value: utf8.RuneCountInString(s), Valid: true }
    }

    var _hx_tmp_2 int = startIndex; _ = _hx_tmp_2
    if ((_hx_tmp_2 > utf8.RuneCountInString(s))) {
        return ""
    }

    var _hx_tmp_3 string = s; _ = _hx_tmp_3
    var _hx_tmp_4 int = startIndex; _ = _hx_tmp_4
    return string(([]rune)(_hx_tmp_3)[_hx_tmp_4:endIndex.Value])
}

func Hx_Field_go_haxe_hxstring_toString(s string) string {
    return s
}

func Hx_Field_go_haxe_hxstring_fromCharCode(code int) string {
    return string(rune(code))
}
