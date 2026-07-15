package main

var Hx_Obj_haxe_log_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.Log",
)

type Hx_Obj_VTable_haxe_log interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_log struct {
    VTable Hx_Obj_VTable_haxe_log
}

func Hx_Obj_haxe_log_CreateEmptyInstance() *Hx_Obj_haxe_log {
    obj := &Hx_Obj_haxe_log{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_log_CreateInstance() *Hx_Obj_haxe_log {
    obj := Hx_Obj_haxe_log_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_haxe_log) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_log_RTTI
}

func Hx_Field_haxe_log_formatOutput(v any, infos any) string {
    var str string = Hx_Field_std_string(v); _ = str
    if (Hx_Field_go_haxe_hxdynamic_equals(infos, nil)) {
        return str
    }

    var _hx_tmp_0 any = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(infos, "fileName"), ((any)(":"))); _ = _hx_tmp_0
    var pstr string = Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(infos, "lineNumber"))); _ = pstr
    if (Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(infos, "customParams"), nil)) {
        var _g int = 0; _ = _g
        var _g1 any = Hx_Field_go_haxe_hxdynamic_getField(infos, "customParams"); _ = _g1
        for  {
            var _hx_tmp_1 any = ((any)(_g)); _ = _hx_tmp_1
            if (!(Hx_Field_go_haxe_hxdynamic_lt(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_getField(_g1, "length")))) {
                break
            }
        
            var v any = Hx_Field_go_haxe_hxdynamic_getArrayIndex(_g1, _g); _ = v
            _g++
            var _hx_tmp_2 string = str; _ = _hx_tmp_2
            str = (_hx_tmp_2 + (", " + Hx_Field_std_string(v)))
        }
    }

    return ((pstr + ": ") + str)
}

var Hx_Field_haxe_log_trace = func (v any, infos any) {
    var str string = Hx_Field_haxe_log_formatOutput(v, infos); _ = str
    Hx_Field_sys_println(str)
}
