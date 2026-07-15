package main

type Hx_Typedef_haxe__int64___int64 = int64

var Hx_Obj_haxe__int64_int64_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe._Int64.Int64_Impl_",
)

type Hx_Obj_VTable_haxe__int64_int64_impl_ interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__int64_int64_impl_ struct {
    VTable Hx_Obj_VTable_haxe__int64_int64_impl_
}

func Hx_Obj_haxe__int64_int64_impl__CreateEmptyInstance() *Hx_Obj_haxe__int64_int64_impl_ {
    obj := &Hx_Obj_haxe__int64_int64_impl_{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe__int64_int64_impl__CreateInstance() *Hx_Obj_haxe__int64_int64_impl_ {
    obj := Hx_Obj_haxe__int64_int64_impl__CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_haxe__int64_int64_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe__int64_int64_impl__RTTI
}

var Hx_Field_haxe__int64_int64_impl__MASK int64 = Hx_Init_Hx_Field_haxe__int64_int64_impl__MASK()
func Hx_Init_Hx_Field_haxe__int64_int64_impl__MASK() int64 {
    var v int64 = int64(65535); _ = v
    var _hx_tmp_0 int64 = v; _ = _hx_tmp_0
    var _hx_tmp_1 int64 = v; _ = _hx_tmp_1
    var _hx_tmp_2 int64 = _hx_tmp_0; _ = _hx_tmp_2
    var _hx_tmp_3 int64 = _hx_tmp_1; _ = _hx_tmp_3
    return (_hx_tmp_2 | (_hx_tmp_3 << int64(16)))
}

func Hx_Field_haxe__int64_int64_impl__make(high int, low int) Hx_Typedef_haxe__int64___int64 {
    var h int64 = int64(high); _ = h
    var l int64 = int64(low); _ = l
    var _hx_tmp_0 int64 = h; _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_0 << int64(32)) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
}

func Hx_Field_haxe__int64_int64_impl___hx_new(x Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)(x))
}

func Hx_Field_haxe__int64_int64_impl__get_val(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return this
}

func Hx_Field_haxe__int64_int64_impl__set_val(this Hx_Typedef_haxe__int64___int64, x Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    this = x
    return this
}

func Hx_Field_haxe__int64_int64_impl__get_high(this Hx_Typedef_haxe__int64___int64) int {
    var _hx_tmp_0 int64 = ((int64)(this)); _ = _hx_tmp_0
    return ((int)((_hx_tmp_0 >> int64(32))))
}

func Hx_Field_haxe__int64_int64_impl__get_low(this Hx_Typedef_haxe__int64___int64) int {
    return ((int)(this))
}

func Hx_Field_haxe__int64_int64_impl__copy(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)(this))
}

func Hx_Field_haxe__int64_int64_impl__ofInt(x int) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)(x))
}

func Hx_Field_haxe__int64_int64_impl__fromInt(x int) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)(x))
}

func Hx_Field_haxe__int64_int64_impl__toFloat(this Hx_Typedef_haxe__int64___int64) float64 {
    return Hx_Field_haxe_int64helper_toFloat(this)
}

func Hx_Field_haxe__int64_int64_impl__is(val any) bool {
    return false
}

func Hx_Field_haxe__int64_int64_impl__isInt64(val any) bool {
    return false
}

func Hx_Field_haxe__int64_int64_impl__toInt(x Hx_Typedef_haxe__int64___int64) int {
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    var _hx_tmp_1 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(x)); _ = _hx_tmp_1
    if ((_hx_tmp_1 < int64(-2147483648))) {
        _hx_tmp_0 = true
    } else {
        var _hx_tmp_2 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(x)); _ = _hx_tmp_2
        var _hx_tmp_3 Hx_Typedef_haxe__int64___int64 = _hx_tmp_2; _ = _hx_tmp_3
        _hx_tmp_0 = (_hx_tmp_3 > int64(2147483647))
    }

    if (_hx_tmp_0) {
        panic("Overflow")
    }

    return ((int)(((Hx_Typedef_haxe__int64___int64)(x))))
}

func Hx_Field_haxe__int64_int64_impl__getHigh(x Hx_Typedef_haxe__int64___int64) int {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(x)))); _ = _hx_tmp_0
    return ((int)((_hx_tmp_0 >> int64(32))))
}

func Hx_Field_haxe__int64_int64_impl__getLow(x Hx_Typedef_haxe__int64___int64) int {
    return ((int)(((Hx_Typedef_haxe__int64___int64)(x))))
}

func Hx_Field_haxe__int64_int64_impl__isNeg(x Hx_Typedef_haxe__int64___int64) bool {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(x)); _ = _hx_tmp_0
    return (_hx_tmp_0 < int64(0))
}

func Hx_Field_haxe__int64_int64_impl__isZero(x Hx_Typedef_haxe__int64___int64) bool {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(x)); _ = _hx_tmp_0
    return (_hx_tmp_0 == int64(0))
}

func Hx_Field_haxe__int64_int64_impl__compare(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) int {
    if ((((Hx_Typedef_haxe__int64___int64)(a)) < ((Hx_Typedef_haxe__int64___int64)(b)))) {
        return -1
    }

    if ((((Hx_Typedef_haxe__int64___int64)(a)) > ((Hx_Typedef_haxe__int64___int64)(b)))) {
        return 1
    }

    return 0
}

func Hx_Field_haxe__int64_int64_impl__ucompare(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) int {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(a)); _ = _hx_tmp_0
    if ((_hx_tmp_0 < int64(0))) {
        var _hx_tmp_1 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(b)); _ = _hx_tmp_1
        if ((_hx_tmp_1 < int64(0))) {
            if ((((Hx_Typedef_haxe__int64___int64)(a)) < ((Hx_Typedef_haxe__int64___int64)(b)))) {
                return -1
            } else {
                if ((((Hx_Typedef_haxe__int64___int64)(a)) > ((Hx_Typedef_haxe__int64___int64)(b)))) {
                    return 1
                } else {
                    return 0
                }
            }
        } else {
            return 1
        }
    }

    var _hx_tmp_1 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(b)); _ = _hx_tmp_1
    if ((_hx_tmp_1 < int64(0))) {
        return -1
    } else {
        if ((((Hx_Typedef_haxe__int64___int64)(a)) < ((Hx_Typedef_haxe__int64___int64)(b)))) {
            return -1
        } else {
            if ((((Hx_Typedef_haxe__int64___int64)(a)) > ((Hx_Typedef_haxe__int64___int64)(b)))) {
                return 1
            } else {
                return 0
            }
        }
    }
}

func Hx_Field_haxe__int64_int64_impl__toStr(x Hx_Typedef_haxe__int64___int64) string {
    return ("" + Hx_Field_std_string(((Hx_Typedef_haxe__int64___int64)(x))))
}

func Hx_Field_haxe__int64_int64_impl__divMod(dividend Hx_Typedef_haxe__int64___int64, divisor Hx_Typedef_haxe__int64___int64) any {
    return any(map[string]any{ "quotient": ((any)(((Hx_Typedef_haxe__int64___int64)(((float64)((((int64)(((Hx_Typedef_haxe__int64___int64)(dividend)))) / ((int64)(((Hx_Typedef_haxe__int64___int64)(divisor))))))))))), "modulus": ((any)(((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(dividend)))) % ((int64)(((Hx_Typedef_haxe__int64___int64)(divisor))))))))) })
}

func Hx_Field_haxe__int64_int64_impl__toString(this Hx_Typedef_haxe__int64___int64) string {
    return ("" + Hx_Field_std_string(this))
}

func Hx_Field_haxe__int64_int64_impl__parseString(sParam string) Hx_Typedef_haxe__int64___int64 {
    return Hx_Field_haxe_int64helper_parseString(sParam)
}

func Hx_Field_haxe__int64_int64_impl__fromFloat(f float64) Hx_Typedef_haxe__int64___int64 {
    return Hx_Field_haxe_int64helper_fromFloat(f)
}

func Hx_Field_haxe__int64_int64_impl__neg(x Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)(-((int64)(((Hx_Typedef_haxe__int64___int64)(x))))))
}

func Hx_Field_haxe__int64_int64_impl__preIncrement(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    this = (this + ((int64)(1)))
    return this
}

func Hx_Field_haxe__int64_int64_impl__postIncrement(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    this = (this + ((int64)(1)))
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(this)) - ((int64)(1)))))
}

func Hx_Field_haxe__int64_int64_impl__preDecrement(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    this = (this - ((int64)(1)))
    return this
}

func Hx_Field_haxe__int64_int64_impl__postDecrement(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    this = (this - ((int64)(1)))
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(this)) + ((int64)(1)))))
}

func Hx_Field_haxe__int64_int64_impl__add(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(a)))) + ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__addInt(a Hx_Typedef_haxe__int64___int64, b int) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(a)))); _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 + int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__sub(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(a)))) - ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__subInt(a Hx_Typedef_haxe__int64___int64, b int) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(a)))); _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 - int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__intSub(a int, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((int64(a) - ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__mul(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(a)))) * ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__mulInt(a Hx_Typedef_haxe__int64___int64, b int) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(a)))); _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 * int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__div(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)(((float64)((((int64)(((Hx_Typedef_haxe__int64___int64)(a)))) / ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))))
}

func Hx_Field_haxe__int64_int64_impl__divInt(a Hx_Typedef_haxe__int64___int64, b int) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(a)))); _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)(((float64)((_hx_tmp_0 / int64(b))))))
}

func Hx_Field_haxe__int64_int64_impl__intDiv(a int, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)(((float64)((int64(a) / ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))))
}

func Hx_Field_haxe__int64_int64_impl__mod(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(a)))) % ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__modInt(a Hx_Typedef_haxe__int64___int64, b int) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(a)))); _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 % int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__intMod(a int, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((int64(a) % ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__eq(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
    return (((Hx_Typedef_haxe__int64___int64)(a)) == ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__eqInt(a Hx_Typedef_haxe__int64___int64, b int) bool {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(a)); _ = _hx_tmp_0
    return (_hx_tmp_0 == int64(b))
}

func Hx_Field_haxe__int64_int64_impl__neq(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
    return (((Hx_Typedef_haxe__int64___int64)(a)) != ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__neqInt(a Hx_Typedef_haxe__int64___int64, b int) bool {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(a)); _ = _hx_tmp_0
    return (_hx_tmp_0 != int64(b))
}

func Hx_Field_haxe__int64_int64_impl__lt(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
    return (((Hx_Typedef_haxe__int64___int64)(a)) < ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__ltInt(a Hx_Typedef_haxe__int64___int64, b int) bool {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(a)); _ = _hx_tmp_0
    return (_hx_tmp_0 < int64(b))
}

func Hx_Field_haxe__int64_int64_impl__intLt(a int, b Hx_Typedef_haxe__int64___int64) bool {
    return (int64(a) < ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__lte(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
    return (((Hx_Typedef_haxe__int64___int64)(a)) <= ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__lteInt(a Hx_Typedef_haxe__int64___int64, b int) bool {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(a)); _ = _hx_tmp_0
    return (_hx_tmp_0 <= int64(b))
}

func Hx_Field_haxe__int64_int64_impl__intLte(a int, b Hx_Typedef_haxe__int64___int64) bool {
    return (int64(a) <= ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__gt(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
    return (((Hx_Typedef_haxe__int64___int64)(a)) > ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__gtInt(a Hx_Typedef_haxe__int64___int64, b int) bool {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(a)); _ = _hx_tmp_0
    return (_hx_tmp_0 > int64(b))
}

func Hx_Field_haxe__int64_int64_impl__intGt(a int, b Hx_Typedef_haxe__int64___int64) bool {
    return (int64(a) > ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__gte(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
    return (((Hx_Typedef_haxe__int64___int64)(a)) >= ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__gteInt(a Hx_Typedef_haxe__int64___int64, b int) bool {
    var _hx_tmp_0 Hx_Typedef_haxe__int64___int64 = ((Hx_Typedef_haxe__int64___int64)(a)); _ = _hx_tmp_0
    return (_hx_tmp_0 >= int64(b))
}

func Hx_Field_haxe__int64_int64_impl__intGte(a int, b Hx_Typedef_haxe__int64___int64) bool {
    return (int64(a) >= ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_Field_haxe__int64_int64_impl__complement(x Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)(^((int64)(((Hx_Typedef_haxe__int64___int64)(x))))))
}

func Hx_Field_haxe__int64_int64_impl__and(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(a)))) & ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__or(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(a)))) | ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__xor(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
    return ((Hx_Typedef_haxe__int64___int64)((((int64)(((Hx_Typedef_haxe__int64___int64)(a)))) ^ ((int64)(((Hx_Typedef_haxe__int64___int64)(b)))))))
}

func Hx_Field_haxe__int64_int64_impl__shl(a Hx_Typedef_haxe__int64___int64, b int) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(a)))); _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 << int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__shr(a Hx_Typedef_haxe__int64___int64, b int) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(a)))); _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 >> int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__ushr(a Hx_Typedef_haxe__int64___int64, b int) Hx_Typedef_haxe__int64___int64 {
    var _hx_tmp_0 int64 = ((int64)(((Hx_Typedef_haxe__int64___int64)(a)))); _ = _hx_tmp_0
    return ((Hx_Typedef_haxe__int64___int64)(((int64)((((uint64)(_hx_tmp_0)) >> int64(b))))))
}
