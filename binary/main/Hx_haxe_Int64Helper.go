package main

import "unicode/utf8"
import "math"

var Hx_Obj_haxe_int64helper_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.Int64Helper",
)

type Hx_Obj_VTable_haxe_int64helper interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_int64helper struct {
    VTable Hx_Obj_VTable_haxe_int64helper
}

func Hx_Obj_haxe_int64helper_CreateEmptyInstance() *Hx_Obj_haxe_int64helper {
    obj := &Hx_Obj_haxe_int64helper{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_int64helper_CreateInstance() *Hx_Obj_haxe_int64helper {
    obj := Hx_Obj_haxe_int64helper_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_haxe_int64helper) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_int64helper_RTTI
}

func Hx_Field_haxe_int64helper_parseString(sParam string) Hx_Typedef_haxe__int64___int64 {
    var base Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)(10)))); _ = base
    var current Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)(0)))); _ = current
    var multiplier Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)(1)))); _ = multiplier
    var sIsNegative bool = false; _ = sIsNegative
    var s string = Hx_Field_stringtools_trim(sParam); _ = s
    if ((Hx_Field_go_haxe_hxstring_charAt(s, 0) == "-")) {
        sIsNegative = true
        var _hx_tmp_0 string = s; _ = _hx_tmp_0
        s = Hx_Field_go_haxe_hxstring_substring(_hx_tmp_0, 1, struct { Value int; Valid bool }{ Value: utf8.RuneCountInString(s), Valid: true })
    }

    var _hx_reserved_len int = utf8.RuneCountInString(s); _ = _hx_reserved_len
    {
        var _g int = 0; _ = _g
        var _g1 int = _hx_reserved_len; _ = _g1
        for ((_g < _g1)) {
            var _hx_tmp_0 int = _g; _ = _hx_tmp_0
            _g = (_g + 1)
            var i int = _hx_tmp_0; _ = i
            var digitInt int = (Hx_Field_go_haxe_hxstring_charCodeAt(s, ((_hx_reserved_len - 1) - i)).Value - 48); _ = digitInt
            if (((digitInt < 0) || (digitInt > 9))) {
                panic("NumberFormatError")
            }
        
            if ((digitInt != 0)) {
                var digit Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)(digitInt)))); _ = digit
                if (sIsNegative) {
                    current = ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(current)))) - ((int64)(((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(multiplier)))) * ((int64)(((Hx_Typedef_haxe__int64___int64)(digit))))))))))))))
                    var _hx_tmp_1 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(current)); _ = _hx_tmp_1
                    if (!((_hx_tmp_1 < int64(0)))) {
                        panic("NumberFormatError: Underflow")
                    }
                } else {
                    current = ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(current)))) + ((int64)(((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(multiplier)))) * ((int64)(((Hx_Typedef_haxe__int64___int64)(digit))))))))))))))
                    var _hx_tmp_1 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(current)); _ = _hx_tmp_1
                    if ((_hx_tmp_1 < int64(0))) {
                        panic("NumberFormatError: Overflow")
                    }
                }
            }
        
            multiplier = ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(multiplier)))) * ((int64)(((Hx_Typedef_haxe__int64___int64)(base)))))))
        }
    }

    return current
}

func Hx_Field_haxe_int64helper_fromFloat(f float64) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (Hx_Field_math_isNaN(f)) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = !Hx_Field_math_isFinite(f)
    }

    if (_hx_tmp_0) {
        panic("Number is NaN or Infinite")
    }

    var _hx_tmp_1 float64 = f; _ = _hx_tmp_1
    var noFractions float64 = (_hx_tmp_1 - math.Mod(f, ((float64)(1)))); _ = noFractions
    if ((noFractions > 9007199254740991)) {
        panic("Conversion overflow")
    }

    if ((noFractions < -9007199254740991)) {
        panic("Conversion underflow")
    }

    var result Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)(0)))); _ = result
    var neg bool = (noFractions < 0); _ = neg
    var _hx_tmp_2 float64; _ = _hx_tmp_2
    if (neg) {
        _hx_tmp_2 = -noFractions
    } else {
        _hx_tmp_2 = noFractions
    }

    var rest float64 = _hx_tmp_2; _ = rest
    var i int = 0; _ = i
    for ((rest >= 1)) {
        var curr float64 = math.Mod(rest, ((float64)(2))); _ = curr
        rest = (rest / ((float64)(2)))
        if ((curr >= 1)) {
            var _hx_tmp_3 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(result)))); _ = _hx_tmp_3
            var _hx_tmp_4 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)(1)))))); _ = _hx_tmp_4
            result = ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_3 + ((int64)((((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)((_hx_tmp_4 << int64(i))))))))))))
        }
    
        i++
    }

    if (neg) {
        result = Hx_Field_haxe__int64_int64_impl__neg(((Hx_Typedef_haxe__int64___int64)(result)))
    }

    return result
}

func Hx_Field_haxe_int64helper_toFloat(x Hx_Typedef_haxe__int64___int64) float64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(x)))); _ = _hx_tmp_0
    var high int = ((int)(((int)((_hx_tmp_0 >> int64(32)))))); _ = high
    var low int = ((int)(((int)(((Hx_Typedef_haxe__int64___int64)(x)))))); _ = low
    var _hx_tmp_1 float64 = (((float64)(high)) * 4294967296.); _ = _hx_tmp_1
    var _hx_tmp_2 float64; _ = _hx_tmp_2
    if ((low < 0)) {
        _hx_tmp_2 = (((float64)(low)) + 4294967296.)
    } else {
        _hx_tmp_2 = ((float64)(low))
    }

    return (_hx_tmp_1 + (_hx_tmp_2))
}
