package main

import "strconv"
import "reflect"
import "fmt"

var Hx_Obj_std_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "Std",
)

type Hx_Obj_VTable_std interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_std struct {
    VTable Hx_Obj_VTable_std
}

func Hx_Obj_std_CreateEmptyInstance() *Hx_Obj_std {
    obj := &Hx_Obj_std{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_std_CreateInstance() *Hx_Obj_std {
    obj := Hx_Obj_std_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_std) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_std_RTTI
}

func Hx_Field_std_int(x float64) int {
    return ((int)(int(x)))
}

func Hx_Field_std_parseInt(x string) int {
    var hx_result_0 struct { Error error; Result int }; _ = hx_result_0
    hx_result_0.Result, hx_result_0.Error = strconv.Atoi(x)
    var this1 struct { Error error; Result int } = ((struct { Error error; Result int })(hx_result_0)); _ = this1
    if ((this1.Error != nil)) {
        var e error = this1.Error; _ = e
        panic(e)
    } else {
        var r int = this1.Result; _ = r
        return r
    }
}

func Hx_Field_std_parseFloat(x string) float64 {
    var hx_result_1 struct { Error error; Result float64 }; _ = hx_result_1
    hx_result_1.Result, hx_result_1.Error = strconv.ParseFloat(x, 64)
    var this1 struct { Error error; Result float64 } = ((struct { Error error; Result float64 })(hx_result_1)); _ = this1
    if ((this1.Error != nil)) {
        var e error = this1.Error; _ = e
        panic(e)
    } else {
        var r float64 = this1.Result; _ = r
        return r
    }
}

func Hx_Field_std_string(s any) string {
    if (Hx_Field_go_haxe_hxdynamic_equals(s, nil)) {
        return "null"
    }

    var value reflect.Value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(s); _ = value
    if (!value.IsValid()) {
        return "null"
    }

    var kind reflect.Kind = value.Kind(); _ = kind
    var _hx_tmp_0 reflect.Kind = kind; _ = _hx_tmp_0
    if ((_hx_tmp_0 == reflect.Ptr)) {
        return Hx_Field_std_string(value.Elem())
    }

    var _hx_tmp_1 bool; _ = _hx_tmp_1
    var _hx_tmp_2 reflect.Kind = kind; _ = _hx_tmp_2
    if ((_hx_tmp_2 == reflect.Array)) {
        _hx_tmp_1 = true
    } else {
        var _hx_tmp_3 reflect.Kind = kind; _ = _hx_tmp_3
        var _hx_tmp_4 reflect.Kind = _hx_tmp_3; _ = _hx_tmp_4
        _hx_tmp_1 = (_hx_tmp_4 == reflect.Slice)
    }

    if (_hx_tmp_1) {
        var buf_b string = ""; _ = buf_b
        buf_b = (buf_b + "[")
        {
            var arr any = value.Interface(); _ = arr
            var f func(any) string = func(s any) string {
                return Hx_Field_std_string(s)
            }; _ = f
            var output *[]string = &([]string{}); _ = output
            {
                var _g int = 0; _ = _g
                for  {
                    var _hx_tmp_3 int = _g; _ = _hx_tmp_3
                    if (!((_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)))) {
                        break
                    }
                
                    var x any = Hx_Field_go_haxe_hxdynamic_getArrayIndex(arr, _g); _ = x
                    _g++
                    {
                        var x1 string = f(x); _ = x1
                        {
                            var data []string = (*output); _ = data
                            var _hx_tmp_4 *[]string = output; _ = _hx_tmp_4
                            (*_hx_tmp_4) = append(data, x1)
                            var _hx_tmp_5 int = len(data); _ = _hx_tmp_5
                            var this1 int = (_hx_tmp_5 + int(1)); _ = this1
                        }
                    }
                }
            }
        
            var data []string = (*output); _ = data
            var length int = len(data); _ = length
            var sep string = ","; _ = sep
            var x string; _ = x
            var _hx_tmp_3 int = length; _ = _hx_tmp_3
            if ((_hx_tmp_3 == int(0))) {
                x = ""
            } else {
                var result string = ""; _ = result
                var i int = int(0); _ = i
                for ((i < length)) {
                    var _hx_tmp_4 string = result; _ = _hx_tmp_4
                    result = (_hx_tmp_4 + Hx_Field_std_string(data[((int)(i))]))
                    var _hx_tmp_5 int = i; _ = _hx_tmp_5
                    var _hx_tmp_6 int = length; _ = _hx_tmp_6
                    if ((_hx_tmp_5 < (_hx_tmp_6 - int(1)))) {
                        result = (result + sep)
                    }
                
                    i = (i + ((int)(1)))
                }
            
                x = result
            }
        
            var _hx_tmp_4 string = buf_b; _ = _hx_tmp_4
            buf_b = (_hx_tmp_4 + Hx_Field_std_string(x))
        }
    
        buf_b = (buf_b + "]")
        return buf_b
    }

    var _hx_tmp_3 reflect.Kind = kind; _ = _hx_tmp_3
    if ((_hx_tmp_3 == reflect.Map)) {
        var buf_b string = ""; _ = buf_b
        var keys []reflect.Value = value.MapKeys(); _ = keys
        buf_b = (buf_b + "[")
        {
            var self []reflect.Value = keys; _ = self
            var _this *[]reflect.Value = (&self); _ = _this
            var output *[]string = &([]string{}); _ = output
            {
                var _g int = 0; _ = _g
                for  {
                    var _hx_tmp_4 int = _g; _ = _hx_tmp_4
                    if (!((_hx_tmp_4 < len(*_this)))) {
                        break
                    }
                
                    var x reflect.Value = (*_this)[_g]; _ = x
                    _g++
                    {
                        var _hx_tmp_5 string = (("" + Hx_Field_std_string(x)) + " => "); _ = _hx_tmp_5
                        var x1 string = (_hx_tmp_5 + Hx_Field_std_string(value.MapIndex(x))); _ = x1
                        {
                            var data []string = (*output); _ = data
                            var _hx_tmp_6 *[]string = output; _ = _hx_tmp_6
                            (*_hx_tmp_6) = append(data, x1)
                            var _hx_tmp_7 int = len(data); _ = _hx_tmp_7
                            var this1 int = (_hx_tmp_7 + int(1)); _ = this1
                        }
                    }
                }
            }
        
            var data []string = (*output); _ = data
            var length int = len(data); _ = length
            var sep string = ", "; _ = sep
            var x string; _ = x
            var _hx_tmp_4 int = length; _ = _hx_tmp_4
            if ((_hx_tmp_4 == int(0))) {
                x = ""
            } else {
                var result string = ""; _ = result
                var i int = int(0); _ = i
                for ((i < length)) {
                    var _hx_tmp_5 string = result; _ = _hx_tmp_5
                    result = (_hx_tmp_5 + Hx_Field_std_string(data[((int)(i))]))
                    var _hx_tmp_6 int = i; _ = _hx_tmp_6
                    var _hx_tmp_7 int = length; _ = _hx_tmp_7
                    if ((_hx_tmp_6 < (_hx_tmp_7 - int(1)))) {
                        result = (result + sep)
                    }
                
                    i = (i + ((int)(1)))
                }
            
                x = result
            }
        
            var _hx_tmp_5 string = buf_b; _ = _hx_tmp_5
            buf_b = (_hx_tmp_5 + Hx_Field_std_string(x))
        }
    
        buf_b = (buf_b + "]")
        return buf_b
    }

    var _hx_tmp_4 reflect.Kind = kind; _ = _hx_tmp_4
    if ((_hx_tmp_4 == reflect.Struct)) {
        var enumIndexMethod reflect.Value = value.MethodByName("Hx_Field_enumIndex"); _ = enumIndexMethod
        var enumTypeMethod reflect.Value = value.MethodByName("Hx_Field_enumType"); _ = enumTypeMethod
        var _hx_tmp_5 bool; _ = _hx_tmp_5
        if (enumIndexMethod.IsValid()) {
            _hx_tmp_5 = enumTypeMethod.IsValid()
        } else {
            _hx_tmp_5 = false
        }
    
        if (_hx_tmp_5) {
            var self *[]reflect.Value = &([]reflect.Value{}); _ = self
            var enumIndex int = ((int)(Hx_Field_go_haxe_hxdynamic_toInt(enumIndexMethod.Call((*self))[0].Interface()))); _ = enumIndex
            var self1 *[]reflect.Value = &([]reflect.Value{}); _ = self1
            var enumType *Hx_Obj_go_haxe_hxenum = (Hx_Field_go_haxe_hxdynamic_toClass(enumTypeMethod.Call((*self1))[0].Interface(), "Hx_Obj_go_haxe_hxenum")).(*Hx_Obj_go_haxe_hxenum); _ = enumType
            var enumCtorName string = (*enumType.Hx_Field_constructorNames)[enumIndex]; _ = enumCtorName
            var enumCtorCount int = (*enumType.Hx_Field_constructorArgCounts)[enumIndex]; _ = enumCtorCount
            var values *[]string = &([]string{}); _ = values
            {
                var _g int = 0; _ = _g
                var _g1 int = ((int)(value.NumField())); _ = _g1
                for ((_g < _g1)) {
                    var _hx_tmp_6 int = _g; _ = _hx_tmp_6
                    _g = (_g + 1)
                    var i int = _hx_tmp_6; _ = i
                    {
                        var x string = Hx_Field_std_string(value.Field(int(i))); _ = x
                        {
                            var data []string = (*values); _ = data
                            var _hx_tmp_7 *[]string = values; _ = _hx_tmp_7
                            (*_hx_tmp_7) = append(data, x)
                            var _hx_tmp_8 int = len(data); _ = _hx_tmp_8
                            var this1 int = (_hx_tmp_8 + int(1)); _ = this1
                        }
                    }
                }
            }
        
            if ((enumCtorCount == 0)) {
                return enumCtorName
            } else {
                var tmp string = (("" + enumCtorName) + "("); _ = tmp
                var data []string = (*values); _ = data
                var length int = len(data); _ = length
                var sep string = ","; _ = sep
                var tmp1 string; _ = tmp1
                var _hx_tmp_6 int = length; _ = _hx_tmp_6
                if ((_hx_tmp_6 == int(0))) {
                    tmp1 = ""
                } else {
                    var result string = ""; _ = result
                    var i int = int(0); _ = i
                    for ((i < length)) {
                        var _hx_tmp_7 string = result; _ = _hx_tmp_7
                        result = (_hx_tmp_7 + Hx_Field_std_string(data[((int)(i))]))
                        var _hx_tmp_8 int = i; _ = _hx_tmp_8
                        var _hx_tmp_9 int = length; _ = _hx_tmp_9
                        if ((_hx_tmp_8 < (_hx_tmp_9 - int(1)))) {
                            result = (result + sep)
                        }
                    
                        i = (i + ((int)(1)))
                    }
                
                    tmp1 = result
                }
            
                return ((tmp + tmp1) + ")")
            }
        }
    
        var valid reflect.Value = value.FieldByName("Valid"); _ = valid
        if (valid.IsValid()) {
            var val reflect.Value = value.FieldByName("Value"); _ = val
            var _hx_tmp_6 bool; _ = _hx_tmp_6
            if (Hx_Field_go_haxe_hxdynamic_equals(valid.Interface(), ((any)(false)))) {
                _hx_tmp_6 = true
            } else {
                _hx_tmp_6 = !val.IsValid()
            }
        
            if (_hx_tmp_6) {
                return "null"
            } else {
                return Hx_Field_std_string(val.Interface())
            }
        }
    
        var vt reflect.Value = value.FieldByName("VTable"); _ = vt
        if (vt.IsValid()) {
            var toStr reflect.Value = vt.MethodByName("Hx_Field_toString"); _ = toStr
            if (toStr.IsValid()) {
                var self *[]reflect.Value = &([]reflect.Value{}); _ = self
                return Hx_Field_std_string(toStr.Call((*self))[0])
            }
        }
    
        return fmt.Sprintf("%v", value.Interface())
    }

    var _hx_tmp_5 bool; _ = _hx_tmp_5
    var _hx_tmp_6 reflect.Kind = kind; _ = _hx_tmp_6
    if ((_hx_tmp_6 == reflect.Interface)) {
        _hx_tmp_5 = value.IsNil()
    } else {
        _hx_tmp_5 = false
    }

    if (_hx_tmp_5) {
        return "null"
    }

    return fmt.Sprintf("%v", value.Interface())
}

func Hx_Field_std_isOfType(v any, t any) bool {
    return false
}
