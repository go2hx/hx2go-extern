package main

import "golang.org/x/tools/go/packages"
import "go/types"
import "reflect"
import "os"

var Hx_Obj_main_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "Main",
)

type Hx_Obj_VTable_main interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_main struct {
    VTable Hx_Obj_VTable_main
}

func Hx_Obj_main_CreateEmptyInstance() *Hx_Obj_main {
    obj := &Hx_Obj_main{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_main_CreateInstance() *Hx_Obj_main {
    obj := Hx_Obj_main_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_main) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_main_RTTI
}

func Hx_Field_main_toHaxeCase(input string) string {
    var _hx_tmp_0 string = Hx_Field_go_haxe_hxstring_toLowerCase(Hx_Field_go_haxe_hxstring_charAt(input, 0)); _ = _hx_tmp_0
    var _hx_tmp_1 string = input; _ = _hx_tmp_1
    return (_hx_tmp_0 + Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, 1, struct { Value int; Valid bool }{}))
}

func Hx_Field_main_toPascalCase(input string) string {
    var _hx_tmp_0 string = Hx_Field_go_haxe_hxstring_toUpperCase(Hx_Field_go_haxe_hxstring_charAt(input, 0)); _ = _hx_tmp_0
    var _hx_tmp_1 string = input; _ = _hx_tmp_1
    return (_hx_tmp_0 + Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, 1, struct { Value int; Valid bool }{}))
}

func Hx_Field_main_main() {
    var _hx_tmp_2 packages.LoadMode = packages.NeedName; _ = _hx_tmp_2
    var _hx_tmp_1 packages.LoadMode = _hx_tmp_2 | packages.NeedTypes; _ = _hx_tmp_1
    var _hx_tmp_0 packages.LoadMode = _hx_tmp_1 | packages.NeedTypesInfo; _ = _hx_tmp_0
    var config packages.Config = packages.Config{ Mode: _hx_tmp_0 | packages.NeedSyntax }; _ = config
    var args *[]string = Hx_Field_sys_args(); _ = args
    if ((len(*args) < 2)) {
        Hx_Field_sys_println("Usage: go2hx <lib> <output>")
        Hx_Field_sys_exit(1)
    }

    var output string = (*args)[1]; _ = output
    var lib string = (*args)[0]; _ = lib
    Hx_Field_sys_println((((("Writing \"" + lib) + "\" to \"") + output) + "\""))
    var hx_result_4 struct { Error error; Result []*packages.Package }; _ = hx_result_4
    hx_result_4.Result, hx_result_4.Error = packages.Load((&config), lib)
    var this1 struct { Error error; Result []*packages.Package } = ((struct { Error error; Result []*packages.Package })(hx_result_4)); _ = this1
    var entries []*packages.Package; _ = entries
    if ((this1.Error != nil)) {
        var e error = this1.Error; _ = e
        panic(e)
    } else {
        var r []*packages.Package = this1.Result; _ = r
        entries = r
    }

    var outputs map[string]any = map[string]any{}; _ = outputs
    var getOutput func(string) any = func(name string) any {
        return func(name string) any {
            var ok bool = false; _ = ok
            _, ok = outputs[name]
            if (!ok) {
                var _hx_tmp_3 any = ((any)(Hx_Obj_stringbuf_CreateInstance())); _ = _hx_tmp_3
                var _hx_tmp_4 any = ((any)(Hx_Obj_stringbuf_CreateInstance())); _ = _hx_tmp_4
                var _hx_tmp_5 any = ((any)(Hx_Obj_stringbuf_CreateInstance())); _ = _hx_tmp_5
                var value any = any(map[string]any{ "staticFunctions": _hx_tmp_3, "instanceFunctions": _hx_tmp_4, "staticVars": _hx_tmp_5, "instanceVars": ((any)(Hx_Obj_stringbuf_CreateInstance())) }); _ = value
                outputs[name] = value
            }
        
            return outputs[name]
        }(name)
    }; _ = getOutput
    {
        var self []*packages.Package = entries; _ = self
        var _g_current int = 0; _ = _g_current
        var _g_array *[]*packages.Package = (&self); _ = _g_array
        for  {
            var _hx_tmp_3 int = _g_current; _ = _hx_tmp_3
            if (!((_hx_tmp_3 < len(*_g_array)))) {
                break
            }
        
            var _hx_tmp_4 *[]*packages.Package = _g_array; _ = _hx_tmp_4
            var _hx_tmp_5 int = _g_current; _ = _hx_tmp_5
            _g_current = (_g_current + 1)
            var entry *packages.Package = (*_hx_tmp_4)[_hx_tmp_5]; _ = entry
            var scope types.Scope = (*(*(*entry).Types).Scope()); _ = scope
            {
                var tmp_self_1 []string = scope.Names(); _ = tmp_self_1
                var name_current int = 0; _ = name_current
                var name_array *[]string = (&tmp_self_1); _ = name_array
                for  {
                    var _hx_tmp_6 int = name_current; _ = _hx_tmp_6
                    if (!((_hx_tmp_6 < len(*name_array)))) {
                        break
                    }
                
                    var _hx_tmp_7 *[]string = name_array; _ = _hx_tmp_7
                    var _hx_tmp_8 int = name_current; _ = _hx_tmp_8
                    name_current = (name_current + 1)
                    var name string = (*_hx_tmp_7)[_hx_tmp_8]; _ = name
                    var obj types.Object = scope.Lookup(name); _ = obj
                    if (!obj.Exported()) {
                        continue
                    }
                
                    switch obj.(type) {
                    case *types.TypeName:
                    var v *types.TypeName = (((any)(obj))).(*types.TypeName); _ = v
                    var _hx_reserved_type types.TypeName = (*v); _ = _hx_reserved_type
                    case *types.Func:
                    {
                        var v1 *types.Func = (((any)(obj))).(*types.Func); _ = v1
                        var _hx_reserved_func types.Func = (*v1); _ = _hx_reserved_func
                        var buf *Hx_Obj_stringbuf = (Hx_Field_go_haxe_hxdynamic_toClass(Hx_Field_go_haxe_hxdynamic_getField(getOutput((*entry).Name), "staticFunctions"), "Hx_Obj_stringbuf")).(*Hx_Obj_stringbuf); _ = buf
                        var sig types.Signature = (*_hx_reserved_func.Signature()); _ = sig
                        var recv *types.Var = sig.Recv(); _ = recv
                        var tmp *types.Tuple = sig.Params(); _ = tmp
                        var _hx_tmp_9 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_9
                        if ((tmp != nil)) {
                            _hx_tmp_9 = struct { Value types.Tuple; Valid bool }{ Value: (*tmp), Valid: true }
                        } else {
                            _hx_tmp_9 = struct { Value types.Tuple; Valid bool }{}
                        }
                    
                        var params struct { Value types.Tuple; Valid bool } = _hx_tmp_9; _ = params
                        var _hx_tmp_10 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_10
                        if ((params.Valid != false)) {
                            _hx_tmp_10 = params
                        } else {
                            _hx_tmp_10 = struct { Value types.Tuple; Valid bool }{}
                        }
                    
                        var params1 struct { Value types.Tuple; Valid bool } = _hx_tmp_10; _ = params1
                        var tmp1 *types.Tuple = sig.Results(); _ = tmp1
                        var _hx_tmp_11 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_11
                        if ((tmp1 != nil)) {
                            _hx_tmp_11 = struct { Value types.Tuple; Valid bool }{ Value: (*tmp1), Valid: true }
                        } else {
                            _hx_tmp_11 = struct { Value types.Tuple; Valid bool }{}
                        }
                    
                        var results struct { Value types.Tuple; Valid bool } = _hx_tmp_11; _ = results
                        var _hx_tmp_12 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_12
                        if ((results.Valid != false)) {
                            _hx_tmp_12 = results
                        } else {
                            _hx_tmp_12 = struct { Value types.Tuple; Valid bool }{}
                        }
                    
                        var results1 struct { Value types.Tuple; Valid bool } = _hx_tmp_12; _ = results1
                        var varadic bool = sig.Variadic(); _ = varadic
                        {
                            var _hx_tmp_13 string = name; _ = _hx_tmp_13
                            var _hx_tmp_14 *types.Var = recv; _ = _hx_tmp_14
                            var _hx_tmp_15 types.Tuple = params1.Value; _ = _hx_tmp_15
                            var _hx_tmp_16 types.Tuple = results1.Value; _ = _hx_tmp_16
                            var _hx_tmp_17 bool = varadic; _ = _hx_tmp_17
                            var _hx_tmp_18 bool = (recv == nil); _ = _hx_tmp_18
                            var x string = (("    " + Hx_Field_main_genFunc(_hx_tmp_13, _hx_tmp_14, _hx_tmp_15, _hx_tmp_16, _hx_tmp_17, _hx_tmp_18, struct { Value bool; Valid bool }{}.Value)) + "\n"); _ = x
                            var _hx_tmp_19 string = buf.Hx_Field_b; _ = _hx_tmp_19
                            buf.Hx_Field_b = (_hx_tmp_19 + Hx_Field_std_string(x))
                        }
                    }
                
                    case *types.Var:
                    {
                        var v2 *types.Var = (((any)(obj))).(*types.Var); _ = v2
                        var v3 types.Var = (*v2); _ = v3
                        var buf1 *Hx_Obj_stringbuf = (Hx_Field_go_haxe_hxdynamic_toClass(Hx_Field_go_haxe_hxdynamic_getField(getOutput((*entry).Name), "staticVars"), "Hx_Obj_stringbuf")).(*Hx_Obj_stringbuf); _ = buf1
                        var name1 string = v3.Name(); _ = name1
                        var _hx_tmp_9 string = buf1.Hx_Field_b; _ = _hx_tmp_9
                        buf1.Hx_Field_b = (_hx_tmp_9 + Hx_Field_std_string((("    static var " + name1) + ";")))
                    }
                
                    default:
                    var _hx_tmp_9 string = reflect.TypeOf(obj).String(); _ = _hx_tmp_9
                    Hx_Field_haxe_log_trace(_hx_tmp_9, any(map[string]any{ "fileName": ((any)("source/Main.hx")), "lineNumber": ((any)(95)), "className": ((any)("Main")), "methodName": ((any)("main")) }))
                    }
                }
            }
        }
    }

    var tmp func(any, any) = func(v any, infos any) {
        Hx_Field_haxe_log_trace(v, infos)
    }; _ = tmp
    var length struct { Value int; Valid bool } = struct { Value int; Valid bool }{}; _ = length
    var _hx_tmp_3 int; _ = _hx_tmp_3
    if ((length.Valid != false)) {
        _hx_tmp_3 = length.Value
    } else {
        _hx_tmp_3 = 0
    }

    var keys []string = make([]string, _hx_tmp_3); _ = keys
    for k := range outputs {keys = append(keys, k)}

    tmp("keys", any(map[string]any{ "fileName": ((any)("source/Main.hx")), "lineNumber": ((any)(104)), "className": ((any)("Main")), "methodName": ((any)("main")), "customParams": &([]any{ ((any)(keys)) }) }))
    {
        var length1 struct { Value int; Valid bool } = struct { Value int; Valid bool }{}; _ = length1
        var _hx_tmp_4 int; _ = _hx_tmp_4
        if ((length1.Valid != false)) {
            _hx_tmp_4 = length1.Value
        } else {
            _hx_tmp_4 = 0
        }
    
        var keys1 []string = make([]string, _hx_tmp_4); _ = keys1
        for k := range outputs {keys1 = append(keys1, k)}
    
        var self1 []string = keys1; _ = self1
        var file_current int = 0; _ = file_current
        var file_array *[]string = (&self1); _ = file_array
        for  {
            var _hx_tmp_5 int = file_current; _ = _hx_tmp_5
            if (!((_hx_tmp_5 < len(*file_array)))) {
                break
            }
        
            var _hx_tmp_6 *[]string = file_array; _ = _hx_tmp_6
            var _hx_tmp_7 int = file_current; _ = _hx_tmp_7
            file_current = (file_current + 1)
            var file string = (*_hx_tmp_6)[_hx_tmp_7]; _ = file
            var buf_b string = ""; _ = buf_b
            var out any = outputs[file]; _ = out
            var _hx_tmp_8 string = buf_b; _ = _hx_tmp_8
            buf_b = (_hx_tmp_8 + Hx_Field_std_string((("package go." + Hx_Field_stringtools_replace(lib, "/", ".")) + ";\n")))
            buf_b = (buf_b + "\n")
            {
                var _this *[]string = Hx_Field_go_haxe_hxstring_split(lib, "/"); _ = _this
                var data []string = (*_this); _ = data
                var _hx_tmp_9 int = len(data); _ = _hx_tmp_9
                var lastIdx int = (_hx_tmp_9 - int(1)); _ = lastIdx
                var x struct { Value string; Valid bool }; _ = x
                var _hx_tmp_10 int = lastIdx; _ = _hx_tmp_10
                if ((_hx_tmp_10 < int(0))) {
                    x = struct { Value string; Valid bool }{}
                } else {
                    var last string = data[((int)(lastIdx))]; _ = last
                    var _hx_tmp_11 *[]string = _this; _ = _hx_tmp_11
                    (*_hx_tmp_11) = data[:lastIdx]
                    x = struct { Value string; Valid bool }{ Value: last, Valid: true }
                }
            
                var _hx_tmp_11 string = buf_b; _ = _hx_tmp_11
                var _hx_tmp_12 string = (("@:go.Type({ name: \"" + file) + "\", instanceName: \""); _ = _hx_tmp_12
                buf_b = (_hx_tmp_11 + Hx_Field_std_string(((((((_hx_tmp_12 + x.Value) + ".") + file) + "\", imports: [\"") + lib) + "\"] })\n")))
            }
        
            var _hx_tmp_9 string = buf_b; _ = _hx_tmp_9
            buf_b = (_hx_tmp_9 + Hx_Field_std_string((("extern class " + Hx_Field_main_toPascalCase(file)) + " {\n\n")))
            var _hx_tmp_10 string = buf_b; _ = _hx_tmp_10
            buf_b = (_hx_tmp_10 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "staticVars"), "b")))
            var _hx_tmp_11 string = buf_b; _ = _hx_tmp_11
            buf_b = (_hx_tmp_11 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "instanceVars"), "b")))
            buf_b = (buf_b + "\n")
            var _hx_tmp_12 string = buf_b; _ = _hx_tmp_12
            buf_b = (_hx_tmp_12 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "staticFunctions"), "b")))
            var _hx_tmp_13 string = buf_b; _ = _hx_tmp_13
            buf_b = (_hx_tmp_13 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "instanceFunctions"), "b")))
            buf_b = (buf_b + "\n}")
            var _hx_tmp_15 string = ((("" + output) + "/go/") + lib); _ = _hx_tmp_15
            var _hx_tmp_14 error = os.MkdirAll(_hx_tmp_15, 0775); _ = _hx_tmp_14
            Hx_Field_haxe_log_trace(_hx_tmp_14, any(map[string]any{ "fileName": ((any)("source/Main.hx")), "lineNumber": ((any)(121)), "className": ((any)("Main")), "methodName": ((any)("main")) }))
            var _hx_tmp_18 string = (((("" + output) + "/go/") + lib) + "/"); _ = _hx_tmp_18
            var _hx_tmp_17 string = ((_hx_tmp_18 + Hx_Field_main_toPascalCase(file)) + ".hx"); _ = _hx_tmp_17
            var _hx_tmp_19 []byte = (([]byte)(buf_b)); _ = _hx_tmp_19
            var _hx_tmp_16 error = os.WriteFile(_hx_tmp_17, _hx_tmp_19, 0666); _ = _hx_tmp_16
            Hx_Field_haxe_log_trace(_hx_tmp_16, any(map[string]any{ "fileName": ((any)("source/Main.hx")), "lineNumber": ((any)(122)), "className": ((any)("Main")), "methodName": ((any)("main")) }))
        }
    }
}

func Hx_Field_main_genPackage(pkg any) string {
    return ""
}

func Hx_Field_main_genFile(file any) string {
    return ""
}

func Hx_Field_main_genFunc(name string, recv *types.Var, params types.Tuple, results types.Tuple, varadic bool, topLevel bool, closure bool) string {
    var _hx_tmp_0 types.Tuple = params; _ = _hx_tmp_0
    var _hx_tmp_1 bool = varadic; _ = _hx_tmp_1
    var params1 *[]string = Hx_Field_main_genTuple(_hx_tmp_0, _hx_tmp_1, struct { Value bool; Valid bool }{}.Value); _ = params1
    var meta string = ""; _ = meta
    var _hx_tmp_2 bool; _ = _hx_tmp_2
    var _hx_tmp_3 bool; _ = _hx_tmp_3
    if ((results.Len() > 1)) {
        _hx_tmp_3 = !Hx_Field_main_isResultType(results)
    } else {
        _hx_tmp_3 = false
    }

    if (_hx_tmp_3) {
        _hx_tmp_2 = !closure
    } else {
        _hx_tmp_2 = false
    }

    if (_hx_tmp_2) {
        var names *[]string = &([]string{}); _ = names
        var unnamed int = 0; _ = unnamed
        {
            var _g int = 0; _ = _g
            var _g1 int = results.Len(); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_4 int = _g; _ = _hx_tmp_4
                _g = (_g + 1)
                var idx int = _hx_tmp_4; _ = idx
                var name string = (*results.At(idx)).Name(); _ = name
                if ((name == "")) {
                    var _hx_tmp_5 int = unnamed; _ = _hx_tmp_5
                    unnamed = (unnamed + 1)
                    name = ("p" + Hx_Field_std_string(_hx_tmp_5))
                }
            
                {
                    var data []string = (*names); _ = data
                    var _hx_tmp_5 *[]string = names; _ = _hx_tmp_5
                    (*_hx_tmp_5) = append(data, (("\"" + name) + "\""))
                    var _hx_tmp_6 int = len(data); _ = _hx_tmp_6
                    var this1 int = (_hx_tmp_6 + int(1)); _ = this1
                }
            }
        }
    
        var data []string = (*names); _ = data
        var length int = len(data); _ = length
        var sep string = ", "; _ = sep
        var meta1 string; _ = meta1
        var _hx_tmp_4 int = length; _ = _hx_tmp_4
        if ((_hx_tmp_4 == int(0))) {
            meta1 = ""
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
        
            meta1 = result
        }
    
        meta = (("@:go.Tuple(" + meta1) + ") ")
    }

    var _hx_tmp_4 string; _ = _hx_tmp_4
    if (closure) {
        _hx_tmp_4 = ""
    } else {
        _hx_tmp_4 = ("function " + Hx_Field_main_toHaxeCase(name))
    }

    var tmp string = _hx_tmp_4; _ = tmp
    var _hx_tmp_5 string = ("" + meta); _ = _hx_tmp_5
    var _hx_tmp_6 string; _ = _hx_tmp_6
    if ((topLevel && !closure)) {
        _hx_tmp_6 = "static "
    } else {
        _hx_tmp_6 = ""
    }

    var tmp1 string = (((_hx_tmp_5 + (_hx_tmp_6)) + tmp) + "("); _ = tmp1
    var data []string = (*params1); _ = data
    var length int = len(data); _ = length
    var sep string = ", "; _ = sep
    var tmp2 string; _ = tmp2
    var _hx_tmp_7 int = length; _ = _hx_tmp_7
    if ((_hx_tmp_7 == int(0))) {
        tmp2 = ""
    } else {
        var result string = ""; _ = result
        var i int = int(0); _ = i
        for ((i < length)) {
            var _hx_tmp_8 string = result; _ = _hx_tmp_8
            result = (_hx_tmp_8 + Hx_Field_std_string(data[((int)(i))]))
            var _hx_tmp_9 int = i; _ = _hx_tmp_9
            var _hx_tmp_10 int = length; _ = _hx_tmp_10
            if ((_hx_tmp_9 < (_hx_tmp_10 - int(1)))) {
                result = (result + sep)
            }
        
            i = (i + ((int)(1)))
        }
    
        tmp2 = result
    }

    var _hx_tmp_8 string; _ = _hx_tmp_8
    if ((results.Len() == 0)) {
        _hx_tmp_8 = "Void"
    } else {
        _hx_tmp_8 = Hx_Field_main_genResults(results)
    }

    var tmp3 string = _hx_tmp_8; _ = tmp3
    var _hx_tmp_9 string = ((tmp1 + tmp2) + ")"); _ = _hx_tmp_9
    var _hx_tmp_10 string; _ = _hx_tmp_10
    if (closure) {
        _hx_tmp_10 = " -> "
    } else {
        _hx_tmp_10 = ": "
    }

    var _hx_tmp_11 string; _ = _hx_tmp_11
    if (closure) {
        _hx_tmp_11 = ""
    } else {
        _hx_tmp_11 = ";"
    }

    return (((_hx_tmp_9 + (_hx_tmp_10)) + tmp3) + (_hx_tmp_11))
}

func Hx_Field_main_isResultType(args types.Tuple) bool {
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if ((args.Len() == 2)) {
        _hx_tmp_0 = ((*args.At(1)).Type().String() == "error")
    } else {
        _hx_tmp_0 = false
    }

    return _hx_tmp_0
}

func Hx_Field_main_genResults(args types.Tuple) string {
    if ((args.Len() > 1)) {
        if (Hx_Field_main_isResultType(args)) {
            return (("go.Result<" + Hx_Field_main_genType((*args.At(0)).Type())) + ">")
        }
    
        var data []string = (*Hx_Field_main_genTuple(args, false, false)); _ = data
        var length int = len(data); _ = length
        var sep string = ", "; _ = sep
        var tmp string; _ = tmp
        var _hx_tmp_0 int = length; _ = _hx_tmp_0
        if ((_hx_tmp_0 == int(0))) {
            tmp = ""
        } else {
            var result string = ""; _ = result
            var i int = int(0); _ = i
            for ((i < length)) {
                var _hx_tmp_1 string = result; _ = _hx_tmp_1
                result = (_hx_tmp_1 + Hx_Field_std_string(data[((int)(i))]))
                var _hx_tmp_2 int = i; _ = _hx_tmp_2
                var _hx_tmp_3 int = length; _ = _hx_tmp_3
                if ((_hx_tmp_2 < (_hx_tmp_3 - int(1)))) {
                    result = (result + sep)
                }
            
                i = (i + ((int)(1)))
            }
        
            tmp = result
        }
    
        return (("go.Tuple<{ " + tmp) + " }>")
    } else {
        var data []string = (*Hx_Field_main_genTuple(args, false, true)); _ = data
        var length int = len(data); _ = length
        var sep string = ", "; _ = sep
        var _hx_tmp_0 int = length; _ = _hx_tmp_0
        if ((_hx_tmp_0 == int(0))) {
            return ""
        } else {
            var result string = ""; _ = result
            var i int = int(0); _ = i
            for ((i < length)) {
                var _hx_tmp_1 string = result; _ = _hx_tmp_1
                result = (_hx_tmp_1 + Hx_Field_std_string(data[((int)(i))]))
                var _hx_tmp_2 int = i; _ = _hx_tmp_2
                var _hx_tmp_3 int = length; _ = _hx_tmp_3
                if ((_hx_tmp_2 < (_hx_tmp_3 - int(1)))) {
                    result = (result + sep)
                }
            
                i = (i + ((int)(1)))
            }
        
            return result
        }
    }
}

func Hx_Field_main_genTuple(args types.Tuple, varadic bool, ret bool) *[]string {
    var items *[]string = &([]string{}); _ = items
    var idx int = 0; _ = idx
    {
        var _g int = 0; _ = _g
        var _g1 int = args.Len(); _ = _g1
        for ((_g < _g1)) {
            var _hx_tmp_0 int = _g; _ = _hx_tmp_0
            _g = (_g + 1)
            var i int = _hx_tmp_0; _ = i
            var _hx_tmp_1 int = i; _ = _hx_tmp_1
            var isLastArg bool = (_hx_tmp_1 == (args.Len() - 1)); _ = isLastArg
            var v types.Var = (*args.At(i)); _ = v
            var n string = v.Name(); _ = n
            if ((n == "")) {
                var _hx_tmp_2 int = idx; _ = _hx_tmp_2
                idx = (idx + 1)
                n = ("p" + Hx_Field_std_string(_hx_tmp_2))
            }
        
            {
                var x string; _ = x
                if (ret) {
                    x = Hx_Field_main_genType(v.Type())
                } else {
                    if ((varadic && isLastArg)) {
                        var v1 *types.Slice = (((any)(v.Type()))).(*types.Slice); _ = v1
                        var _hx_tmp_2 string = (n + ": ..."); _ = _hx_tmp_2
                        x = (_hx_tmp_2 + Hx_Field_main_genType((*v1).Elem()))
                    } else {
                        var _hx_tmp_2 string = (n + ": "); _ = _hx_tmp_2
                        x = (_hx_tmp_2 + Hx_Field_main_genType(v.Type()))
                    }
                }
            
                {
                    var data []string = (*items); _ = data
                    var _hx_tmp_2 *[]string = items; _ = _hx_tmp_2
                    (*_hx_tmp_2) = append(data, x)
                    var _hx_tmp_3 int = len(data); _ = _hx_tmp_3
                    var this1 int = (_hx_tmp_3 + int(1)); _ = this1
                }
            }
        }
    }

    return items
}

func Hx_Field_main_genStmt(stmt any) string {
    return ""
}

func Hx_Field_main_genExpr(expr any) string {
    return ""
}

func Hx_Field_main_resolvePath(path string) string {
    return ("go." + Hx_Field_stringtools_replace(path, "/", "."))
}

func Hx_Field_main_genType(t types.Type) string {
    var s string = t.String(); _ = s
    switch (s) {
        case "any":
            return "Dynamic"
    
        case "bool":
            return "Bool"
    
        case "byte":
            return "go.Byte"
    
        case "comparable":
            return "go.Comparable"
    
        case "complex128":
            return "go.Complex128"
    
        case "complex64":
            return "go.Complex64"
    
        case "error":
            return "go.Error"
    
        case "float16":
            return "go.Float16"
    
        case "float32":
            return "go.Float32"
    
        case "float64":
            return "Float"
    
        case "int":
            return "Int"
    
        case "int16":
            return "go.Int16"
    
        case "int32":
            return "go.Int32"
    
        case "int64":
            return "go.Int64"
    
        case "int8":
            return "go.Int8"
    
        case "rune":
            return "go.Rune"
    
        case "string":
            return "String"
    
        case "uint":
            return "go.UInt"
    
        case "uint16":
            return "go.UInt16"
    
        case "uint32":
            return "go.UInt32"
    
        case "uint64":
            return "go.UInt64"
    
        case "uint8":
            return "go.UInt8"
    
        case "uintptr":
            return "go.UIntPtr"
    
        default: 
            if (Hx_Field_stringtools_startsWith(s, "chan ")) {
                var v *types.Chan = (((any)(t))).(*types.Chan); _ = v
                return (("go.Chan<" + Hx_Field_main_genType((*v).Elem())) + ">")
            } else {
                if (Hx_Field_stringtools_startsWith(s, "chan ")) {
                    var v *types.Chan = (((any)(t))).(*types.Chan); _ = v
                    return (("go.Chan<" + Hx_Field_main_genType((*v).Elem())) + ">")
                } else {
                    if (Hx_Field_stringtools_startsWith(s, "[]")) {
                        var v *types.Slice = (((any)(t))).(*types.Slice); _ = v
                        return (("go.Slice<" + Hx_Field_main_genType((*v).Elem())) + ">")
                    } else {
                        if (Hx_Field_stringtools_startsWith(s, "*")) {
                            var v *types.Pointer = (((any)(t))).(*types.Pointer); _ = v
                            return (("go.Pointer<" + Hx_Field_main_genType((*v).Elem())) + ">")
                        } else {
                            if (Hx_Field_stringtools_startsWith(s, "func")) {
                                var v *types.Signature = (((any)(t))).(*types.Signature); _ = v
                                var sig types.Signature = (*v); _ = sig
                                var tmp *types.Var = sig.Recv(); _ = tmp
                                var tmp1 *types.Tuple = sig.Params(); _ = tmp1
                                var _hx_tmp_0 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_0
                                if ((tmp1 != nil)) {
                                    _hx_tmp_0 = struct { Value types.Tuple; Valid bool }{ Value: (*tmp1), Valid: true }
                                } else {
                                    _hx_tmp_0 = struct { Value types.Tuple; Valid bool }{}
                                }
                            
                                var tmp2 struct { Value types.Tuple; Valid bool } = _hx_tmp_0; _ = tmp2
                                var tmp3 *types.Tuple = sig.Results(); _ = tmp3
                                var _hx_tmp_1 *types.Var = tmp; _ = _hx_tmp_1
                                var _hx_tmp_2 types.Tuple = tmp2.Value; _ = _hx_tmp_2
                                var _hx_tmp_3 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_3
                                if ((tmp3 != nil)) {
                                    _hx_tmp_3 = struct { Value types.Tuple; Valid bool }{ Value: (*tmp3), Valid: true }
                                } else {
                                    _hx_tmp_3 = struct { Value types.Tuple; Valid bool }{}
                                }
                            
                                return Hx_Field_main_genFunc("", _hx_tmp_1, _hx_tmp_2, _hx_tmp_3.Value, sig.Variadic(), false, true)
                            } else {
                                var _hx_tmp_0 string = s; _ = _hx_tmp_0
                                if ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, ".", struct { Value int; Valid bool }{}) != -1)) {
                                    return Hx_Field_main_resolvePath(s)
                                } else {
                                    return t.String()
                                }
                            }
                        }
                    }
                }
            }
    }
}

var Hx_Obj_elemtype_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "ElemType",
)

type Hx_Obj_VTable_elemtype interface {
    Hx_Field_elem() types.Type
}

type Hx_Obj_elemtype struct {
    VTable Hx_Obj_VTable_elemtype
}
