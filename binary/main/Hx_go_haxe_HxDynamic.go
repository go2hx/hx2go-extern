package main

import "reflect"
import "unicode/utf8"

var Hx_Obj_go_haxe_hxdynamic_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "go.haxe.HxDynamic",
)

type Hx_Obj_VTable_go_haxe_hxdynamic interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxdynamic struct {
    VTable Hx_Obj_VTable_go_haxe_hxdynamic
}

func Hx_Obj_go_haxe_hxdynamic_CreateEmptyInstance() *Hx_Obj_go_haxe_hxdynamic {
    obj := &Hx_Obj_go_haxe_hxdynamic{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_go_haxe_hxdynamic_CreateInstance() *Hx_Obj_go_haxe_hxdynamic {
    obj := Hx_Obj_go_haxe_hxdynamic_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_go_haxe_hxdynamic) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_go_haxe_hxdynamic_RTTI
}

var Hx_Field_go_haxe_hxdynamic_Null reflect.Value = Hx_Init_Hx_Field_go_haxe_hxdynamic_Null()
func Hx_Init_Hx_Field_go_haxe_hxdynamic_Null() reflect.Value {
    return reflect.ValueOf(nil)
}

func Hx_Field_go_haxe_hxdynamic_not(d any) bool {
    var dV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d); _ = dV
    var _hx_tmp_0 reflect.Kind = dV.Kind(); _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Bool)) {
        var dB bool = dV.Bool(); _ = dB
        return !dB
    }

    panic(("runtime.HxDynamic.not value invalid: " + Hx_Field_std_string(d)))
}

func Hx_Field_go_haxe_hxdynamic_increment(d any) any {
    var dV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d); _ = dV
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (dV.CanInt()) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = dV.CanUint()
    }

    if (_hx_tmp_0) {
        return ((any)((Hx_Field_go_haxe_hxdynamic_valueToInt(dV) + 1)))
    } else {
        if (dV.CanFloat()) {
            return ((any)((Hx_Field_go_haxe_hxdynamic_valueToFloat(dV) + 1.0)))
        }
    }

    panic(("runtime.HxDynamic.increment value invalid: " + Hx_Field_std_string(d)))
}

func Hx_Field_go_haxe_hxdynamic_decrement(d any) any {
    var dV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d); _ = dV
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (dV.CanInt()) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = dV.CanUint()
    }

    if (_hx_tmp_0) {
        return ((any)((Hx_Field_go_haxe_hxdynamic_valueToInt(dV) - 1)))
    } else {
        if (dV.CanFloat()) {
            return ((any)((Hx_Field_go_haxe_hxdynamic_valueToFloat(dV) - 1.0)))
        }
    }

    panic(("runtime.HxDynamic.decrement value invalid: " + Hx_Field_std_string(d)))
}

func Hx_Field_go_haxe_hxdynamic_negate(d any) any {
    var dV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d); _ = dV
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (dV.CanInt()) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = dV.CanUint()
    }

    if (_hx_tmp_0) {
        return ((any)((0 - Hx_Field_go_haxe_hxdynamic_valueToInt(dV))))
    } else {
        if (dV.CanFloat()) {
            return ((any)((0.0 - Hx_Field_go_haxe_hxdynamic_valueToFloat(dV))))
        }
    }

    panic(("runtime.HxDynamic.negate value invalid: " + Hx_Field_std_string(d)))
}

func Hx_Field_go_haxe_hxdynamic_bitnot(d any) any {
    var dV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d); _ = dV
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (dV.CanInt()) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = dV.CanUint()
    }

    if (_hx_tmp_0) {
        return ((any)(^Hx_Field_go_haxe_hxdynamic_valueToInt(dV)))
    }

    panic(("runtime.HxDynamic.bitnot value invalid: " + Hx_Field_std_string(d)))
}

func Hx_Field_go_haxe_hxdynamic_and(a any, b any) bool {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    var _hx_tmp_1 reflect.Kind = aV.Kind(); _ = _hx_tmp_1
    if ((_hx_tmp_1 == reflect.Bool)) {
        var _hx_tmp_2 reflect.Kind = bV.Kind(); _ = _hx_tmp_2
        var _hx_tmp_3 reflect.Kind = _hx_tmp_2; _ = _hx_tmp_3
        _hx_tmp_0 = (_hx_tmp_3 == reflect.Bool)
    } else {
        _hx_tmp_0 = false
    }

    if (_hx_tmp_0) {
        var _hx_tmp_2 bool; _ = _hx_tmp_2
        if (aV.Bool()) {
            _hx_tmp_2 = bV.Bool()
        } else {
            _hx_tmp_2 = false
        }
    
        return _hx_tmp_2
    } else {
        var _hx_tmp_2 string = (("runtime.HxDynamic.and invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_2
        panic((_hx_tmp_2 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_or(a any, b any) bool {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    var _hx_tmp_1 reflect.Kind = aV.Kind(); _ = _hx_tmp_1
    if ((_hx_tmp_1 == reflect.Bool)) {
        var _hx_tmp_2 reflect.Kind = bV.Kind(); _ = _hx_tmp_2
        var _hx_tmp_3 reflect.Kind = _hx_tmp_2; _ = _hx_tmp_3
        _hx_tmp_0 = (_hx_tmp_3 == reflect.Bool)
    } else {
        _hx_tmp_0 = false
    }

    if (_hx_tmp_0) {
        var _hx_tmp_2 bool; _ = _hx_tmp_2
        if (aV.Bool()) {
            _hx_tmp_2 = true
        } else {
            _hx_tmp_2 = bV.Bool()
        }
    
        return _hx_tmp_2
    } else {
        var _hx_tmp_2 string = (("runtime.HxDynamic.or invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_2
        panic((_hx_tmp_2 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_jointKind(aV reflect.Value, bV reflect.Value) reflect.Kind {
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (aV.CanInt()) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = aV.CanUint()
    }

    var avCi bool = _hx_tmp_0; _ = avCi
    var _hx_tmp_1 bool; _ = _hx_tmp_1
    if (bV.CanInt()) {
        _hx_tmp_1 = true
    } else {
        _hx_tmp_1 = bV.CanUint()
    }

    var bvCi bool = _hx_tmp_1; _ = bvCi
    if ((avCi && bvCi)) {
        return reflect.Int
    } else {
        var _hx_tmp_2 bool; _ = _hx_tmp_2
        var _hx_tmp_3 bool; _ = _hx_tmp_3
        if (aV.CanFloat()) {
            _hx_tmp_3 = true
        } else {
            _hx_tmp_3 = avCi
        }
    
        if (_hx_tmp_3) {
            var _hx_tmp_4 bool; _ = _hx_tmp_4
            if (bV.CanFloat()) {
                _hx_tmp_4 = true
            } else {
                _hx_tmp_4 = bvCi
            }
        
            _hx_tmp_2 = (_hx_tmp_4)
        } else {
            _hx_tmp_2 = false
        }
    
        if (_hx_tmp_2) {
            return reflect.Float64
        } else {
            var _hx_tmp_4 bool; _ = _hx_tmp_4
            var _hx_tmp_5 reflect.Kind = aV.Kind(); _ = _hx_tmp_5
            if ((_hx_tmp_5 == reflect.String)) {
                _hx_tmp_4 = true
            } else {
                var _hx_tmp_6 reflect.Kind = bV.Kind(); _ = _hx_tmp_6
                var _hx_tmp_7 reflect.Kind = _hx_tmp_6; _ = _hx_tmp_7
                _hx_tmp_4 = (_hx_tmp_7 == reflect.String)
            }
        
            if (_hx_tmp_4) {
                return reflect.String
            }
        }
    }

    return reflect.Invalid
}

func Hx_Field_go_haxe_hxdynamic_toAnySlice(v any) []any {
    var _hx_reserved_len int = Hx_Field_go_haxe_hxdynamic_getArrayLength(v); _ = _hx_reserved_len
    var length struct { Value int; Valid bool } = struct { Value int; Valid bool }{}; _ = length
    var this1 []any; _ = this1
    var length1 struct { Value int; Valid bool } = length; _ = length1
    var _hx_tmp_0 int; _ = _hx_tmp_0
    if ((length1.Valid != false)) {
        _hx_tmp_0 = length1.Value
    } else {
        _hx_tmp_0 = 0
    }

    this1 = make([]any, _hx_tmp_0)
    var slice []any = this1; _ = slice
    {
        var _g int = 0; _ = _g
        var _g1 int = _hx_reserved_len; _ = _g1
        for ((_g < _g1)) {
            var _hx_tmp_1 int = _g; _ = _hx_tmp_1
            _g = (_g + 1)
            var i int = _hx_tmp_1; _ = i
            var v1 any = Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getArrayIndex(v, i)); _ = v1
            slice = append(slice, v1)
        }
    }

    return slice
}

func Hx_Field_go_haxe_hxdynamic_formatField(name string) string {
    if ((utf8.RuneCountInString(name) == 0)) {
        return name
    }

    return ("Hx_Field_" + name)
}

func Hx_Field_go_haxe_hxdynamic_isNull(x any) bool {
    return x == nil
}

func Hx_Field_go_haxe_hxdynamic_equals(a any, b any) bool {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var aN bool = !aV.IsValid(); _ = aN
    var bN bool = !bV.IsValid(); _ = bN
    if ((aN || bN)) {
        if ((aN && bN)) {
            return true
        } else {
            return false
        }
    }

    var aK reflect.Kind = aV.Kind(); _ = aK
    var bK reflect.Kind = bV.Kind(); _ = bK
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    var _hx_tmp_1 reflect.Kind = aK; _ = _hx_tmp_1
    if ((_hx_tmp_1 == reflect.Bool)) {
        _hx_tmp_0 = true
    } else {
        var _hx_tmp_2 reflect.Kind = bK; _ = _hx_tmp_2
        var _hx_tmp_3 reflect.Kind = _hx_tmp_2; _ = _hx_tmp_3
        _hx_tmp_0 = (_hx_tmp_3 == reflect.Bool)
    }

    if (_hx_tmp_0) {
        var _hx_tmp_2 bool; _ = _hx_tmp_2
        var _hx_tmp_3 reflect.Kind = aK; _ = _hx_tmp_3
        if ((_hx_tmp_3 == reflect.Bool)) {
            var _hx_tmp_4 reflect.Kind = bK; _ = _hx_tmp_4
            var _hx_tmp_5 reflect.Kind = _hx_tmp_4; _ = _hx_tmp_5
            _hx_tmp_2 = (_hx_tmp_5 == reflect.Bool)
        } else {
            _hx_tmp_2 = false
        }
    
        if (_hx_tmp_2) {
            var _hx_tmp_4 bool = aV.Bool(); _ = _hx_tmp_4
            return (_hx_tmp_4 == bV.Bool())
        } else {
            return false
        }
    }

    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_2 reflect.Kind = k; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.Int)) {
        var _hx_tmp_3 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_3
        return (_hx_tmp_3 == Hx_Field_go_haxe_hxdynamic_valueToInt(bV))
    } else {
        var _hx_tmp_3 reflect.Kind = k; _ = _hx_tmp_3
        if ((_hx_tmp_3 == reflect.Float64)) {
            var _hx_tmp_4 float64 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV); _ = _hx_tmp_4
            return (_hx_tmp_4 == Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))
        } else {
            var _hx_tmp_4 reflect.Kind = k; _ = _hx_tmp_4
            if ((_hx_tmp_4 == reflect.String)) {
                var _hx_tmp_5 string = Hx_Field_go_haxe_hxdynamic_toString(a); _ = _hx_tmp_5
                return (_hx_tmp_5 == Hx_Field_go_haxe_hxdynamic_toString(b))
            } else {
                var _hx_tmp_5 string = (("runtime.HxDynamic.equals invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_5
                panic((_hx_tmp_5 + bV.String()))
            }
        }
    }
}

func Hx_Field_go_haxe_hxdynamic_nequals(a any, b any) bool {
    return !Hx_Field_go_haxe_hxdynamic_equals(a, b)
}

func Hx_Field_go_haxe_hxdynamic_lt(a any, b any) bool {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return (_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_valueToInt(bV))
    } else {
        var _hx_tmp_1 reflect.Kind = k; _ = _hx_tmp_1
        if ((_hx_tmp_1 == reflect.Float64)) {
            var _hx_tmp_2 float64 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV); _ = _hx_tmp_2
            return (_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))
        } else {
            var _hx_tmp_2 reflect.Kind = k; _ = _hx_tmp_2
            if ((_hx_tmp_2 == reflect.String)) {
                var _hx_tmp_3 string = Hx_Field_go_haxe_hxdynamic_toString(a); _ = _hx_tmp_3
                return (_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_toString(b))
            } else {
                var _hx_tmp_3 string = (("runtime.HxDynamic.lt invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_3
                panic((_hx_tmp_3 + bV.String()))
            }
        }
    }
}

func Hx_Field_go_haxe_hxdynamic_gtequals(a any, b any) bool {
    return !Hx_Field_go_haxe_hxdynamic_lt(a, b)
}

func Hx_Field_go_haxe_hxdynamic_gt(a any, b any) bool {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return (_hx_tmp_1 > Hx_Field_go_haxe_hxdynamic_valueToInt(bV))
    } else {
        var _hx_tmp_1 reflect.Kind = k; _ = _hx_tmp_1
        if ((_hx_tmp_1 == reflect.Float64)) {
            var _hx_tmp_2 float64 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV); _ = _hx_tmp_2
            return (_hx_tmp_2 > Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))
        } else {
            var _hx_tmp_2 reflect.Kind = k; _ = _hx_tmp_2
            if ((_hx_tmp_2 == reflect.String)) {
                var _hx_tmp_3 string = Hx_Field_go_haxe_hxdynamic_toString(a); _ = _hx_tmp_3
                return (_hx_tmp_3 > Hx_Field_go_haxe_hxdynamic_toString(b))
            } else {
                var _hx_tmp_3 string = (("runtime.HxDynamic.gt invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_3
                panic((_hx_tmp_3 + bV.String()))
            }
        }
    }
}

func Hx_Field_go_haxe_hxdynamic_ltequals(a any, b any) bool {
    return !Hx_Field_go_haxe_hxdynamic_gt(a, b)
}

func Hx_Field_go_haxe_hxdynamic_add(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 + Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 reflect.Kind = k; _ = _hx_tmp_1
        if ((_hx_tmp_1 == reflect.Float64)) {
            var _hx_tmp_2 float64 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV); _ = _hx_tmp_2
            return ((any)((_hx_tmp_2 + Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))))
        } else {
            var _hx_tmp_2 reflect.Kind = k; _ = _hx_tmp_2
            if ((_hx_tmp_2 == reflect.String)) {
                var _hx_tmp_3 string = Hx_Field_go_haxe_hxdynamic_toString(a); _ = _hx_tmp_3
                return ((any)((_hx_tmp_3 + Hx_Field_go_haxe_hxdynamic_toString(b))))
            } else {
                var _hx_tmp_3 string = (("runtime.HxDynamic.add invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_3
                panic((_hx_tmp_3 + bV.String()))
            }
        }
    }
}

func Hx_Field_go_haxe_hxdynamic_subtract(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 - Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 reflect.Kind = k; _ = _hx_tmp_1
        if ((_hx_tmp_1 == reflect.Float64)) {
            var _hx_tmp_2 float64 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV); _ = _hx_tmp_2
            return ((any)((_hx_tmp_2 - Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))))
        } else {
            var _hx_tmp_2 string = (("runtime.HxDynamic.subtract invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_2
            panic((_hx_tmp_2 + bV.String()))
        }
    }
}

func Hx_Field_go_haxe_hxdynamic_multiply(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 * Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 reflect.Kind = k; _ = _hx_tmp_1
        if ((_hx_tmp_1 == reflect.Float64)) {
            var _hx_tmp_2 float64 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV); _ = _hx_tmp_2
            return ((any)((_hx_tmp_2 * Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))))
        } else {
            var _hx_tmp_2 string = (("runtime.HxDynamic.multiply invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_2
            panic((_hx_tmp_2 + bV.String()))
        }
    }
}

func Hx_Field_go_haxe_hxdynamic_divide(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    var _hx_tmp_1 reflect.Kind = k; _ = _hx_tmp_1
    if ((_hx_tmp_1 == reflect.Int)) {
        _hx_tmp_0 = true
    } else {
        var _hx_tmp_2 reflect.Kind = k; _ = _hx_tmp_2
        var _hx_tmp_3 reflect.Kind = _hx_tmp_2; _ = _hx_tmp_3
        _hx_tmp_0 = (_hx_tmp_3 == reflect.Float64)
    }

    if (_hx_tmp_0) {
        var _hx_tmp_2 float64 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV); _ = _hx_tmp_2
        return ((any)(((float64)((_hx_tmp_2 / Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))))))
    } else {
        var _hx_tmp_2 string = (("runtime.HxDynamic.divide invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_2
        panic((_hx_tmp_2 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_modulo(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 % Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 string = (("runtime.HxDynamic.modulo invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_1
        panic((_hx_tmp_1 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_bitand(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 & Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 string = (("runtime.HxDynamic.bitand invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_1
        panic((_hx_tmp_1 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_bitor(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 | Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 string = (("runtime.HxDynamic.bitor invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_1
        panic((_hx_tmp_1 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_bitxor(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 ^ Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 string = (("runtime.HxDynamic.bitxor invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_1
        panic((_hx_tmp_1 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_lbitshift(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 << Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 string = (("runtime.HxDynamic.lbitshift invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_1
        panic((_hx_tmp_1 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_rbitshift(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)((_hx_tmp_1 >> Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
    } else {
        var _hx_tmp_1 string = (("runtime.HxDynamic.rbitshift invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_1
        panic((_hx_tmp_1 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_urbitshift(a any, b any) any {
    var aV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a); _ = aV
    var bV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b); _ = bV
    var k reflect.Kind = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        var _hx_tmp_1 int = Hx_Field_go_haxe_hxdynamic_valueToInt(aV); _ = _hx_tmp_1
        return ((any)(((int)((((uint32)(_hx_tmp_1)) >> Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))))
    } else {
        var _hx_tmp_1 string = (("runtime.HxDynamic.urbitshift invalid operands: " + aV.String()) + " and "); _ = _hx_tmp_1
        panic((_hx_tmp_1 + bV.String()))
    }
}

func Hx_Field_go_haxe_hxdynamic_toString(d any) string {
    return Hx_Field_std_string(d)
}

func Hx_Field_go_haxe_hxdynamic_toBool(d any) bool {
    var dV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d); _ = dV
    var _hx_tmp_0 reflect.Kind = dV.Kind(); _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Bool)) {
        return dV.Bool()
    }

    return (Hx_Field_go_haxe_hxdynamic_valueToInt(dV) != 0)
}

func Hx_Field_go_haxe_hxdynamic_valueToInt(dV reflect.Value) int {
    if (dV.CanUint()) {
        var this1 uint64 = dV.Uint(); _ = this1
        return ((int)(((uint64)(this1))))
    } else {
        if (dV.CanInt()) {
            var this1 int64 = dV.Int(); _ = this1
            return ((int)(((int64)(this1))))
        } else {
            if (dV.CanFloat()) {
                var this1 float64 = dV.Float(); _ = this1
                return Hx_Field_math_round(((float64)(((float64)(this1)))))
            }
        }
    }

    return 0
}

func Hx_Field_go_haxe_hxdynamic_isClass(v reflect.Value) bool {
    var kind reflect.Kind = v.Kind(); _ = kind
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    var _hx_tmp_1 reflect.Kind = kind; _ = _hx_tmp_1
    if ((_hx_tmp_1 == reflect.Ptr)) {
        _hx_tmp_0 = true
    } else {
        var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
        var _hx_tmp_3 reflect.Kind = _hx_tmp_2; _ = _hx_tmp_3
        _hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
    }

    if (_hx_tmp_0) {
        return Hx_Field_go_haxe_hxdynamic_isClass(v.Elem())
    }

    var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
    if ((_hx_tmp_2 != reflect.Struct)) {
        return false
    }

    return v.FieldByName("VTable").IsValid()
}

func Hx_Field_go_haxe_hxdynamic_convertToType(v reflect.Value, t reflect.Type) reflect.Value {
    var cv reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(v.Interface()); _ = cv
    var k reflect.Kind = t.Kind(); _ = k
    var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Int)) {
        return reflect.ValueOf(Hx_Field_go_haxe_hxdynamic_valueToInt(cv))
    }

    var _hx_tmp_1 reflect.Kind = k; _ = _hx_tmp_1
    if ((_hx_tmp_1 == reflect.Float64)) {
        return reflect.ValueOf(Hx_Field_go_haxe_hxdynamic_valueToFloat(cv))
    }

    var _hx_tmp_2 reflect.Kind = k; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.String)) {
        return reflect.ValueOf(Hx_Field_go_haxe_hxdynamic_toString(cv.Interface()))
    }

    var _hx_tmp_3 reflect.Kind = k; _ = _hx_tmp_3
    if ((_hx_tmp_3 == reflect.Bool)) {
        return reflect.ValueOf(Hx_Field_go_haxe_hxdynamic_toBool(cv))
    }

    return cv
}

func Hx_Field_go_haxe_hxdynamic_call(fn any, args any) any {
    var fV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(fn); _ = fV
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (Hx_Field_go_haxe_hxdynamic_isNull(fn)) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = !fV.IsValid()
    }

    if (_hx_tmp_0) {
        panic("runtime.HxDynamic.call null function value")
    }

    var _hx_tmp_1 reflect.Kind = fV.Kind(); _ = _hx_tmp_1
    if ((_hx_tmp_1 != reflect.Func)) {
        panic(("runtime.HxDynamic.call value not callable: " + Hx_Field_std_string(fn)))
    }

    var fType reflect.Type = fV.Type(); _ = fType
    var numIn int = fType.NumIn(); _ = numIn
    var isVariadic bool = fType.IsVariadic(); _ = isVariadic
    var argVals *[]reflect.Value = &([]reflect.Value{}); _ = argVals
    {
        var _g int = 0; _ = _g
        var _g1 int = numIn; _ = _g1
        for ((_g < _g1)) {
            var _hx_tmp_2 int = _g; _ = _hx_tmp_2
            _g = (_g + 1)
            var i int = _hx_tmp_2; _ = i
            if ((isVariadic && (i == (numIn - 1)))) {
                break
            }
        
            var pt reflect.Type = fType.In(i); _ = pt
            var _hx_tmp_3 any; _ = _hx_tmp_3
            var _hx_tmp_4 any = ((any)(i)); _ = _hx_tmp_4
            if (Hx_Field_go_haxe_hxdynamic_lt(_hx_tmp_4, Hx_Field_go_haxe_hxdynamic_getField(args, "length"))) {
                _hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getArrayIndex(args, i)
            } else {
                _hx_tmp_3 = nil
            }
        
            var av any = _hx_tmp_3; _ = av
            {
                var x reflect.Value = Hx_Field_go_haxe_hxdynamic_convertToType(Hx_Field_go_haxe_hxdynamic_ensureValue(av), pt); _ = x
                {
                    var data []reflect.Value = (*argVals); _ = data
                    var _hx_tmp_5 *[]reflect.Value = argVals; _ = _hx_tmp_5
                    (*_hx_tmp_5) = append(data, x)
                    var _hx_tmp_6 int = len(data); _ = _hx_tmp_6
                    var _hx_tmp_7 int = _hx_tmp_6; _ = _hx_tmp_7
                    var this1 int = (_hx_tmp_7 + int(1)); _ = this1
                    _ = ((int)(((int)(this1))))
                }
            }
        }
    }

    var self *[]reflect.Value = argVals; _ = self
    var results []reflect.Value = fV.Call((*self)); _ = results
    var _hx_tmp_2 int = len(results); _ = _hx_tmp_2
    if ((_hx_tmp_2 == int(0))) {
        return nil
    }

    return Hx_Field_go_haxe_hxdynamic_ensureInterface(results[0])
}

func Hx_Field_go_haxe_hxdynamic_toInt(d any) int {
    var dV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d); _ = dV
    return Hx_Field_go_haxe_hxdynamic_valueToInt(dV)
}

func Hx_Field_go_haxe_hxdynamic_valueToFloat(dV reflect.Value) float64 {
    if (dV.CanUint()) {
        var this1 uint64 = dV.Uint(); _ = this1
        return ((float64)(((uint64)(this1))))
    } else {
        if (dV.CanInt()) {
            var this1 int64 = dV.Int(); _ = this1
            return ((float64)(((int64)(this1))))
        } else {
            if (dV.CanFloat()) {
                var this1 float64 = dV.Float(); _ = this1
                return ((float64)(((float64)(this1))))
            }
        }
    }

    return 0.0
}

func Hx_Field_go_haxe_hxdynamic_toFloat(d any) float64 {
    var dV reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d); _ = dV
    return Hx_Field_go_haxe_hxdynamic_valueToFloat(dV)
}

func Hx_Field_go_haxe_hxdynamic_valueToClass(value reflect.Value, className string) reflect.Value {
    var kind reflect.Kind = value.Kind(); _ = kind
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (Hx_Field_go_haxe_hxdynamic_isNull(value)) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = !value.IsValid()
    }

    if (_hx_tmp_0) {
        panic(("runtime.HxDynamic.toClass: dynamic is null, cannot cast to " + className))
    }

    var _hx_tmp_1 bool; _ = _hx_tmp_1
    var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.Ptr)) {
        _hx_tmp_1 = true
    } else {
        var _hx_tmp_3 reflect.Kind = kind; _ = _hx_tmp_3
        var _hx_tmp_4 reflect.Kind = _hx_tmp_3; _ = _hx_tmp_4
        _hx_tmp_1 = (_hx_tmp_4 == reflect.Interface)
    }

    if (_hx_tmp_1) {
        value = Hx_Field_go_haxe_hxdynamic_valueToClass(value.Elem(), className)
    }

    var _hx_tmp_3 bool; _ = _hx_tmp_3
    var _hx_tmp_4 reflect.Kind = kind; _ = _hx_tmp_4
    if ((_hx_tmp_4 == reflect.Struct)) {
        _hx_tmp_3 = (value.Type().Name() != className)
    } else {
        _hx_tmp_3 = false
    }

    if (_hx_tmp_3) {
        value = value.FieldByName(className)
    }

    return value
}

func Hx_Field_go_haxe_hxdynamic_toClass(d any, className string) any {
    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureValue(d); _ = value
    var cls reflect.Value = Hx_Field_go_haxe_hxdynamic_valueToClass(value, className); _ = cls
    return cls.Addr().Interface()
}

func Hx_Field_go_haxe_hxdynamic_getField(dyn any, fieldName string) any {
    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn); _ = value
    var kind reflect.Kind = value.Kind(); _ = kind
    var found bool = false; _ = found
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (Hx_Field_go_haxe_hxdynamic_isNull(dyn)) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = !value.IsValid()
    }

    if (_hx_tmp_0) {
        panic(("runtime.HxDynamic.field null field access: " + fieldName))
    }

    var _hx_tmp_1 bool; _ = _hx_tmp_1
    var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.Ptr)) {
        _hx_tmp_1 = true
    } else {
        var _hx_tmp_3 reflect.Kind = kind; _ = _hx_tmp_3
        var _hx_tmp_4 reflect.Kind = _hx_tmp_3; _ = _hx_tmp_4
        _hx_tmp_1 = (_hx_tmp_4 == reflect.Interface)
    }

    if (_hx_tmp_1) {
        value = (Hx_Field_go_haxe_hxdynamic_getField(value.Elem(), fieldName)).(reflect.Value)
        found = true
    }

    var _hx_tmp_3 reflect.Kind = kind; _ = _hx_tmp_3
    if ((_hx_tmp_3 == reflect.Struct)) {
        var f reflect.Value = value.FieldByName(Hx_Field_go_haxe_hxdynamic_formatField(fieldName)); _ = f
        if (f.IsValid()) {
            found = true
        } else {
            var vtable reflect.Value = value.FieldByName("VTable"); _ = vtable
            if (vtable.IsValid()) {
                f = vtable.MethodByName(Hx_Field_go_haxe_hxdynamic_formatField(fieldName))
                found = true
            } else {
                var m reflect.Value = value.MethodByName(Hx_Field_go_haxe_hxdynamic_formatField(fieldName)); _ = m
                if (m.IsValid()) {
                    f = m
                    found = true
                }
            }
        }
    
        value = f
    }

    var _hx_tmp_4 reflect.Kind = kind; _ = _hx_tmp_4
    if ((_hx_tmp_4 == reflect.Map)) {
        found = true
        value = value.MapIndex(Hx_Field_go_haxe_hxdynamic_ensureValue(fieldName))
    }

    var _hx_tmp_5 bool; _ = _hx_tmp_5
    var _hx_tmp_6 reflect.Kind = kind; _ = _hx_tmp_6
    if ((_hx_tmp_6 == reflect.Array)) {
        _hx_tmp_5 = true
    } else {
        var _hx_tmp_7 reflect.Kind = kind; _ = _hx_tmp_7
        var _hx_tmp_8 reflect.Kind = _hx_tmp_7; _ = _hx_tmp_8
        _hx_tmp_5 = (_hx_tmp_8 == reflect.Slice)
    }

    if (_hx_tmp_5) {
        if ((fieldName == "length")) {
            value = Hx_Field_go_haxe_hxdynamic_ensureValue(Hx_Field_go_haxe_hxdynamic_getArrayLength(dyn))
            found = true
        }
    }

    var _hx_tmp_7 reflect.Kind = kind; _ = _hx_tmp_7
    if ((_hx_tmp_7 == reflect.String)) {
        if ((fieldName == "length")) {
            value = Hx_Field_go_haxe_hxdynamic_ensureValue(utf8.RuneCountInString(Hx_Field_go_haxe_hxdynamic_toString(dyn)))
            found = true
        }
    }

    var _hx_tmp_8 any; _ = _hx_tmp_8
    if (found) {
        _hx_tmp_8 = ((any)(value))
    } else {
        _hx_tmp_8 = nil
    }

    return _hx_tmp_8
}

func Hx_Field_go_haxe_hxdynamic_setField(dyn any, fieldName string, v any) any {
    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn); _ = value
    var kind reflect.Kind = value.Kind(); _ = kind
    var _hx_tmp_0 reflect.Kind = kind; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Interface)) {
        return Hx_Field_go_haxe_hxdynamic_setField(value.Elem(), fieldName, v)
    }

    var _hx_tmp_1 reflect.Kind = kind; _ = _hx_tmp_1
    if ((_hx_tmp_1 == reflect.Ptr)) {
        value = value.Elem()
        kind = value.Kind()
    }

    var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.Struct)) {
        var field reflect.Value = value.FieldByName(Hx_Field_go_haxe_hxdynamic_formatField(fieldName)); _ = field
        if (!field.IsValid()) {
            var _hx_tmp_3 string = (("runtime.HxDynamic.setField field \"" + fieldName) + "\" not present on \""); _ = _hx_tmp_3
            panic(((_hx_tmp_3 + Hx_Field_std_string(value)) + "\""))
        }
    
        if (!field.CanSet()) {
            var _hx_tmp_3 string = (("runtime.HxDynamic.setField cannot set \"" + fieldName) + "\" on \""); _ = _hx_tmp_3
            panic(((_hx_tmp_3 + Hx_Field_std_string(value)) + "\""))
        }
    
        field.Set(Hx_Field_go_haxe_hxdynamic_ensureValue(v))
    }

    var _hx_tmp_3 reflect.Kind = kind; _ = _hx_tmp_3
    if ((_hx_tmp_3 == reflect.Map)) {
        var fn reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureValue(fieldName); _ = fn
        var mi reflect.Value = value.MapIndex(fn); _ = mi
        if (!mi.IsValid()) {
            var _hx_tmp_4 string = (("runtime.HxDynamic.setField field \"" + fieldName) + "\" not present on \""); _ = _hx_tmp_4
            panic(((_hx_tmp_4 + Hx_Field_std_string(value)) + "\""))
        }
    
        var _hx_tmp_4 reflect.Value = fn; _ = _hx_tmp_4
        value.SetMapIndex(_hx_tmp_4, Hx_Field_go_haxe_hxdynamic_ensureValue(v))
    }

    return v
}

func Hx_Field_go_haxe_hxdynamic_setArrayIndex(dyn any, index int, v any) {
    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn); _ = value
    var kind reflect.Kind = value.Kind(); _ = kind
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (Hx_Field_go_haxe_hxdynamic_isNull(dyn)) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = !value.IsValid()
    }

    if (_hx_tmp_0) {
        panic("runtime.HxDynamic.setArrayIndex null array access")
    }

    var _hx_tmp_1 reflect.Kind = kind; _ = _hx_tmp_1
    if ((_hx_tmp_1 == reflect.Interface)) {
        value = value.Elem()
        kind = value.Kind()
    }

    var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.Ptr)) {
        value = value.Elem()
        kind = value.Kind()
    }

    var _hx_tmp_3 bool; _ = _hx_tmp_3
    var _hx_tmp_4 reflect.Kind = kind; _ = _hx_tmp_4
    if ((_hx_tmp_4 == reflect.Slice)) {
        _hx_tmp_3 = true
    } else {
        var _hx_tmp_5 reflect.Kind = kind; _ = _hx_tmp_5
        var _hx_tmp_6 reflect.Kind = _hx_tmp_5; _ = _hx_tmp_6
        _hx_tmp_3 = (_hx_tmp_6 == reflect.Array)
    }

    if (_hx_tmp_3) {
        var length int = value.Len(); _ = length
        var _hx_tmp_5 float64 = float64(index); _ = _hx_tmp_5
        if ((_hx_tmp_5 >= float64(length))) {
            var _hx_tmp_6 reflect.Kind = kind; _ = _hx_tmp_6
            if ((_hx_tmp_6 == reflect.Array)) {
                panic("runtime.HxDynamic.setArrayIndex out of bounds exception, cannot grow go array")
            }
        
            var _hx_tmp_7 int = (int(index) - length); _ = _hx_tmp_7
            value.Grow((_hx_tmp_7 + int(1)))
            value.SetLen(int((index + 1)))
        }
    
        var item reflect.Value = value.Index(int(index)); _ = item
        item.Set(Hx_Field_go_haxe_hxdynamic_ensureValue(v))
    }
}

func Hx_Field_go_haxe_hxdynamic_getArrayIndex(dyn any, index int) any {
    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn); _ = value
    var kind reflect.Kind = value.Kind(); _ = kind
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (Hx_Field_go_haxe_hxdynamic_isNull(dyn)) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = !value.IsValid()
    }

    if (_hx_tmp_0) {
        panic("runtime.HxDynamic.getArrayIndex null array access")
    }

    var _hx_tmp_1 bool; _ = _hx_tmp_1
    var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.Ptr)) {
        _hx_tmp_1 = true
    } else {
        var _hx_tmp_3 reflect.Kind = kind; _ = _hx_tmp_3
        var _hx_tmp_4 reflect.Kind = _hx_tmp_3; _ = _hx_tmp_4
        _hx_tmp_1 = (_hx_tmp_4 == reflect.Interface)
    }

    if (_hx_tmp_1) {
        return Hx_Field_go_haxe_hxdynamic_getArrayIndex(value.Elem(), index)
    }

    var _hx_tmp_3 bool; _ = _hx_tmp_3
    var _hx_tmp_4 reflect.Kind = kind; _ = _hx_tmp_4
    if ((_hx_tmp_4 == reflect.Slice)) {
        _hx_tmp_3 = true
    } else {
        var _hx_tmp_5 reflect.Kind = kind; _ = _hx_tmp_5
        var _hx_tmp_6 reflect.Kind = _hx_tmp_5; _ = _hx_tmp_6
        _hx_tmp_3 = (_hx_tmp_6 == reflect.Array)
    }

    if (_hx_tmp_3) {
        var length int = value.Len(); _ = length
        var _hx_tmp_5 float64 = float64(index); _ = _hx_tmp_5
        if ((_hx_tmp_5 >= float64(length))) {
            panic("runtime.HxDynamic.getArrayIndex out of bounds exception")
        }
    
        return value.Index(int(index)).Interface()
    }

    return value.Interface()
}

func Hx_Field_go_haxe_hxdynamic_getArrayLength(dyn any) int {
    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn); _ = value
    var kind reflect.Kind = value.Kind(); _ = kind
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (Hx_Field_go_haxe_hxdynamic_isNull(value)) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = !value.IsValid()
    }

    if (_hx_tmp_0) {
        return 0
    }

    var _hx_tmp_1 bool; _ = _hx_tmp_1
    var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.Ptr)) {
        _hx_tmp_1 = true
    } else {
        var _hx_tmp_3 reflect.Kind = kind; _ = _hx_tmp_3
        var _hx_tmp_4 reflect.Kind = _hx_tmp_3; _ = _hx_tmp_4
        _hx_tmp_1 = (_hx_tmp_4 == reflect.Interface)
    }

    if (_hx_tmp_1) {
        return Hx_Field_go_haxe_hxdynamic_getArrayLength(value.Elem())
    }

    var _hx_tmp_3 bool; _ = _hx_tmp_3
    var _hx_tmp_4 reflect.Kind = kind; _ = _hx_tmp_4
    if ((_hx_tmp_4 == reflect.Slice)) {
        _hx_tmp_3 = true
    } else {
        var _hx_tmp_5 reflect.Kind = kind; _ = _hx_tmp_5
        var _hx_tmp_6 reflect.Kind = _hx_tmp_5; _ = _hx_tmp_6
        _hx_tmp_3 = (_hx_tmp_6 == reflect.Array)
    }

    if (_hx_tmp_3) {
        var this1 int = value.Len(); _ = this1
        return ((int)(((int)(this1))))
    }

    panic("runtime.HxDynamic.getArrayLength invalid type")
}

func Hx_Field_go_haxe_hxdynamic_ensureValue(dyn any) reflect.Value {
    var ok bool = false; _ = ok
    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_Null; _ = value
    value, ok = dyn.(reflect.Value)
    if (!ok) {
        return reflect.ValueOf(dyn)
    }

    return value
}

func Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(dyn any) reflect.Value {
    var v reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn); _ = v
    var k reflect.Kind = v.Kind(); _ = k
    for  {
        var _hx_tmp_0 reflect.Kind = k; _ = _hx_tmp_0
        if (!((_hx_tmp_0 == reflect.Interface))) {
            break
        }
    
        v = v.Elem()
        k = v.Kind()
    }

    return v
}

func Hx_Field_go_haxe_hxdynamic_ensureInterface(dyn any) any {
    var ok bool = false; _ = ok
    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_Null; _ = value
    value, ok = dyn.(reflect.Value)
    if (!ok) {
        return dyn
    }

    if (!value.CanInterface()) {
        panic("runtime.HxDynamic.ensureInterface cannot convert to iface")
    }

    return value.Interface()
}
