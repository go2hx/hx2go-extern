package main

import "math"
import "math/rand"

var Hx_Obj_math_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "Math",
)

type Hx_Obj_VTable_math interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_math struct {
    VTable Hx_Obj_VTable_math
}

func Hx_Obj_math_CreateEmptyInstance() *Hx_Obj_math {
    obj := &Hx_Obj_math{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_math_CreateInstance() *Hx_Obj_math {
    obj := Hx_Obj_math_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_math) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_math_RTTI
}

var Hx_Field_math_PI float64

var Hx_Field_math_NEGATIVE_INFINITY float64

var Hx_Field_math_POSITIVE_INFINITY float64

var Hx_Field_math_NaN float64

func Hx_Field_math_abs(v float64) float64 {
    return ((float64)(math.Abs(float64(v))))
}

func Hx_Field_math_min(a float64, b float64) float64 {
    var _hx_tmp_0 float64 = float64(a); _ = _hx_tmp_0
    return ((float64)(math.Min(_hx_tmp_0, float64(b))))
}

func Hx_Field_math_max(a float64, b float64) float64 {
    var _hx_tmp_0 float64 = float64(a); _ = _hx_tmp_0
    return ((float64)(math.Max(_hx_tmp_0, float64(b))))
}

func Hx_Field_math_sin(v float64) float64 {
    return ((float64)(math.Sin(float64(v))))
}

func Hx_Field_math_cos(v float64) float64 {
    return ((float64)(math.Cos(float64(v))))
}

func Hx_Field_math_tan(v float64) float64 {
    return ((float64)(math.Tan(float64(v))))
}

func Hx_Field_math_asin(v float64) float64 {
    return ((float64)(math.Asin(float64(v))))
}

func Hx_Field_math_acos(v float64) float64 {
    return ((float64)(math.Acos(float64(v))))
}

func Hx_Field_math_atan(v float64) float64 {
    return ((float64)(math.Atan(float64(v))))
}

func Hx_Field_math_atan2(y float64, x float64) float64 {
    var _hx_tmp_0 float64 = float64(y); _ = _hx_tmp_0
    return ((float64)(math.Atan2(_hx_tmp_0, float64(x))))
}

func Hx_Field_math_exp(v float64) float64 {
    return ((float64)(math.Exp(float64(v))))
}

func Hx_Field_math_log(v float64) float64 {
    return ((float64)(math.Log(float64(v))))
}

func Hx_Field_math_pow(v float64, exp float64) float64 {
    var _hx_tmp_0 float64 = float64(v); _ = _hx_tmp_0
    return ((float64)(math.Pow(_hx_tmp_0, float64(exp))))
}

func Hx_Field_math_sqrt(v float64) float64 {
    return ((float64)(math.Sqrt(float64(v))))
}

func Hx_Field_math_isNaN(f float64) bool {
    return math.IsNaN(float64(f))
}

func Hx_Field_math_ffloor(v float64) float64 {
    return ((float64)(math.Floor(float64(v))))
}

func Hx_Field_math_fceil(v float64) float64 {
    return ((float64)(math.Ceil(float64(v))))
}

func Hx_Field_math_fround(v float64) float64 {
    return ((float64)(math.Round(float64(v))))
}

func Hx_Field_math_round(v float64) int {
    return ((int)(int(Hx_Field_math_fround(v))))
}

func Hx_Field_math_floor(v float64) int {
    return ((int)(int(Hx_Field_math_ffloor(v))))
}

func Hx_Field_math_ceil(v float64) int {
    return ((int)(int(Hx_Field_math_fceil(v))))
}

func Hx_Field_math_random() float64 {
    return ((float64)(rand.Float64()))
}

func Hx_Field_math_isFinite(f float64) bool {
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if (!math.IsNaN(float64(f))) {
        var _hx_tmp_1 float64 = float64(f); _ = _hx_tmp_1
        var _hx_tmp_2 float64 = _hx_tmp_1; _ = _hx_tmp_2
        _hx_tmp_0 = !math.IsInf(_hx_tmp_2, int(0))
    } else {
        _hx_tmp_0 = false
    }

    return _hx_tmp_0
}

func Hx_Field_math_get_PI() float64 {
    return ((float64)(math.Pi))
}

func Hx_Field_math_get_NEGATIVE_INFINITY() float64 {
    return ((float64)(math.Inf(int(-1))))
}

func Hx_Field_math_get_POSITIVE_INFINITY() float64 {
    return ((float64)(math.Inf(int(1))))
}

func Hx_Field_math_get_NaN() float64 {
    return ((float64)(math.NaN()))
}
