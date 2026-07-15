package main

import "golang.org/x/tools/go/packages"
import "go/types"
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

var Hx_Field_main_didGen map[string]bool = Hx_Init_Hx_Field_main_didGen()
func Hx_Init_Hx_Field_main_didGen() map[string]bool {
    return map[string]bool{}
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
    var args *[]string = Hx_Field_sys_args(); _ = args
    if ((len(*args) < 2)) {
        Hx_Field_sys_println("Usage: go2hx <lib> <output>")
        Hx_Field_sys_exit(1)
    }

    var output string = (*args)[1]; _ = output
    var lib string = (*args)[0]; _ = lib
    Hx_Field_main_genLib(lib, output)
}

func Hx_Field_main_genLib(lib string, output string) {
    var ok bool = false; _ = ok
    _, ok = Hx_Field_main_didGen[lib]
    if (ok) {
        return
    }

    Hx_Field_main_didGen[lib] = true
    var _hx_tmp_3 packages.LoadMode = packages.NeedName; _ = _hx_tmp_3
    var _hx_tmp_2 packages.LoadMode = _hx_tmp_3 | packages.NeedTypes; _ = _hx_tmp_2
    var _hx_tmp_1 packages.LoadMode = _hx_tmp_2 | packages.NeedTypesInfo; _ = _hx_tmp_1
    var _hx_tmp_0 packages.LoadMode = _hx_tmp_1 | packages.NeedSyntax; _ = _hx_tmp_0
    var config packages.Config = packages.Config{ Mode: _hx_tmp_0 | packages.NeedImports }; _ = config
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
            var tmp_ok_1 bool = false; _ = tmp_ok_1
            _, tmp_ok_1 = outputs[name]
            if (!tmp_ok_1) {
                var _hx_tmp_4 any = ((any)(Hx_Obj_stringbuf_CreateInstance())); _ = _hx_tmp_4
                var _hx_tmp_5 any = ((any)(Hx_Obj_stringbuf_CreateInstance())); _ = _hx_tmp_5
                var _hx_tmp_6 any = ((any)(Hx_Obj_stringbuf_CreateInstance())); _ = _hx_tmp_6
                var value any = any(map[string]any{ "staticFunctions": _hx_tmp_4, "instanceFunctions": _hx_tmp_5, "staticVars": _hx_tmp_6, "instanceVars": ((any)(Hx_Obj_stringbuf_CreateInstance())), "paramStr": ((any)("")) }); _ = value
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
            var _hx_tmp_4 int = _g_current; _ = _hx_tmp_4
            if (!((_hx_tmp_4 < len(*_g_array)))) {
                break
            }
        
            var _hx_tmp_5 *[]*packages.Package = _g_array; _ = _hx_tmp_5
            var _hx_tmp_6 int = _g_current; _ = _hx_tmp_6
            _g_current = (_g_current + 1)
            var entry *packages.Package = (*_hx_tmp_5)[_hx_tmp_6]; _ = entry
            var scope types.Scope = (*(*(*entry).Types).Scope()); _ = scope
            {
                var tmp_this1_1 map[string]*packages.Package = (*entry).Imports; _ = tmp_this1_1
                var length struct { Value int; Valid bool } = struct { Value int; Valid bool }{}; _ = length
                var _hx_tmp_7 int; _ = _hx_tmp_7
                if ((length.Valid != false)) {
                    _hx_tmp_7 = length.Value
                } else {
                    _hx_tmp_7 = 0
                }
            
                var keys []string = make([]string, _hx_tmp_7); _ = keys
                for k := range tmp_this1_1 {keys = append(keys, k)}
            
                var tmp_self_1 []string = keys; _ = tmp_self_1
                var dep_current int = 0; _ = dep_current
                var dep_array *[]string = (&tmp_self_1); _ = dep_array
                for  {
                    var _hx_tmp_8 int = dep_current; _ = _hx_tmp_8
                    if (!((_hx_tmp_8 < len(*dep_array)))) {
                        break
                    }
                
                    var _hx_tmp_9 *[]string = dep_array; _ = _hx_tmp_9
                    var _hx_tmp_10 int = dep_current; _ = _hx_tmp_10
                    dep_current = (dep_current + 1)
                    var dep string = (*_hx_tmp_9)[_hx_tmp_10]; _ = dep
                    Hx_Field_main_genLib(dep, output)
                }
            }
        
            {
                var self1 []string = scope.Names(); _ = self1
                var name_current int = 0; _ = name_current
                var name_array *[]string = (&self1); _ = name_array
                for  {
                    var _hx_tmp_7 int = name_current; _ = _hx_tmp_7
                    if (!((_hx_tmp_7 < len(*name_array)))) {
                        break
                    }
                
                    var _hx_tmp_8 *[]string = name_array; _ = _hx_tmp_8
                    var _hx_tmp_9 int = name_current; _ = _hx_tmp_9
                    name_current = (name_current + 1)
                    var name string = (*_hx_tmp_8)[_hx_tmp_9]; _ = name
                    var obj types.Object = scope.Lookup(name); _ = obj
                    if (!obj.Exported()) {
                        continue
                    }
                
                    switch obj.(type) {
                    case *types.TypeName:
                    {
                        var v *types.TypeName = (((any)(obj))).(*types.TypeName); _ = v
                        var _hx_reserved_type types.TypeName = (*v); _ = _hx_reserved_type
                        var isNamed bool = Hx_Field_main_isNamedType(_hx_reserved_type.Type()); _ = isNamed
                        var out any = getOutput(obj.Name()); _ = out
                        if (isNamed) {
                            var tmp_v_1 *types.Named = (((any)(_hx_reserved_type.Type()))).(*types.Named); _ = tmp_v_1
                            var named types.Named = (*tmp_v_1); _ = named
                            var tp *types.TypeParamList = named.TypeParams(); _ = tp
                            if ((tp != nil)) {
                                var params *[]string = &([]string{}); _ = params
                                var tps types.TypeParamList = (*tp); _ = tps
                                {
                                    var _g int = 0; _ = _g
                                    var _g1 int = tps.Len(); _ = _g1
                                    for ((_g < _g1)) {
                                        var _hx_tmp_10 int = _g; _ = _hx_tmp_10
                                        _g = (_g + 1)
                                        var i int = _hx_tmp_10; _ = i
                                        var t types.TypeParam = (*tps.At(i)); _ = t
                                        var constraint types.Type = t.Constraint(); _ = constraint
                                        var constraintStr string = Hx_Field_main_genType(constraint); _ = constraintStr
                                        {
                                            var x string = ((("" + t.String()) + ": ") + constraintStr); _ = x
                                            {
                                                var data []string = (*params); _ = data
                                                var _hx_tmp_11 *[]string = params; _ = _hx_tmp_11
                                                (*_hx_tmp_11) = append(data, x)
                                                var _hx_tmp_12 int = len(data); _ = _hx_tmp_12
                                                var tmp_this1_1 int = (_hx_tmp_12 + int(1)); _ = tmp_this1_1
                                            }
                                        }
                                    }
                                }
                            
                                var tmp string; _ = tmp
                                if ((len(*params) == 0)) {
                                    tmp = ""
                                } else {
                                    var data []string = (*params); _ = data
                                    var length int = len(data); _ = length
                                    var sep string = ", "; _ = sep
                                    var tmp1 string; _ = tmp1
                                    var _hx_tmp_10 int = length; _ = _hx_tmp_10
                                    if ((_hx_tmp_10 == int(0))) {
                                        tmp1 = ""
                                    } else {
                                        var result string = ""; _ = result
                                        var i int = int(0); _ = i
                                        for ((i < length)) {
                                            var _hx_tmp_11 string = result; _ = _hx_tmp_11
                                            result = (_hx_tmp_11 + Hx_Field_std_string(data[((int)(i))]))
                                            var _hx_tmp_12 int = i; _ = _hx_tmp_12
                                            var _hx_tmp_13 int = length; _ = _hx_tmp_13
                                            if ((_hx_tmp_12 < (_hx_tmp_13 - int(1)))) {
                                                result = (result + sep)
                                            }
                                        
                                            i = (i + ((int)(1)))
                                        }
                                    
                                        tmp1 = result
                                    }
                                
                                    tmp = (("<" + tmp1) + ">")
                                }
                            
                                Hx_Field_go_haxe_hxdynamic_setField(out, "paramStr", tmp)
                            }
                        
                            var methodSet *types.MethodSet = types.NewMethodSet(types.NewPointer(_hx_reserved_type.Type())); _ = methodSet
                            {
                                var _g int = 0; _ = _g
                                var _g1 int = (*methodSet).Len(); _ = _g1
                                for ((_g < _g1)) {
                                    var _hx_tmp_10 int = _g; _ = _hx_tmp_10
                                    _g = (_g + 1)
                                    var i int = _hx_tmp_10; _ = i
                                    var sel types.Selection = (*(*methodSet).At(i)); _ = sel
                                    var tmp_v_2 *types.Func = (((any)(sel.Obj()))).(*types.Func); _ = tmp_v_2
                                    var method types.Func = (*tmp_v_2); _ = method
                                    if (!method.Exported()) {
                                        continue
                                    }
                                
                                    var sig types.Signature = (*method.Signature()); _ = sig
                                    {
                                        var _this *Hx_Obj_stringbuf = (Hx_Field_go_haxe_hxdynamic_toClass(Hx_Field_go_haxe_hxdynamic_getField(out, "instanceFunctions"), "Hx_Obj_stringbuf")).(*Hx_Obj_stringbuf); _ = _this
                                        var _hx_tmp_11 string = method.Name(); _ = _hx_tmp_11
                                        var _hx_tmp_12 types.Signature = sig; _ = _hx_tmp_12
                                        var x string = (("    " + Hx_Field_main_genFunc(_hx_tmp_11, _hx_tmp_12, false, struct { Value bool; Valid bool }{}.Value)) + "\n"); _ = x
                                        var _hx_tmp_13 string = _this.Hx_Field_b; _ = _hx_tmp_13
                                        _this.Hx_Field_b = (_hx_tmp_13 + Hx_Field_std_string(x))
                                    }
                                }
                            }
                        }
                    }
                
                    case *types.Func:
                    {
                        var v1 *types.Func = (((any)(obj))).(*types.Func); _ = v1
                        var _hx_reserved_func types.Func = (*v1); _ = _hx_reserved_func
                        var sig types.Signature = (*_hx_reserved_func.Signature()); _ = sig
                        var recv *types.Var = sig.Recv(); _ = recv
                        var buf *Hx_Obj_stringbuf = (Hx_Field_go_haxe_hxdynamic_toClass(Hx_Field_go_haxe_hxdynamic_getField(getOutput((*entry).Name), "staticFunctions"), "Hx_Obj_stringbuf")).(*Hx_Obj_stringbuf); _ = buf
                        var tmp *types.Tuple = sig.Params(); _ = tmp
                        var _hx_tmp_10 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_10
                        if ((tmp != nil)) {
                            _hx_tmp_10 = struct { Value types.Tuple; Valid bool }{ Value: (*tmp), Valid: true }
                        } else {
                            _hx_tmp_10 = struct { Value types.Tuple; Valid bool }{}
                        }
                    
                        var params struct { Value types.Tuple; Valid bool } = _hx_tmp_10; _ = params
                        var _hx_tmp_11 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_11
                        if ((params.Valid != false)) {
                            _hx_tmp_11 = params
                        } else {
                            _hx_tmp_11 = struct { Value types.Tuple; Valid bool }{}
                        }
                    
                        var params1 struct { Value types.Tuple; Valid bool } = _hx_tmp_11; _ = params1
                        var tmp1 *types.Tuple = sig.Results(); _ = tmp1
                        var _hx_tmp_12 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_12
                        if ((tmp1 != nil)) {
                            _hx_tmp_12 = struct { Value types.Tuple; Valid bool }{ Value: (*tmp1), Valid: true }
                        } else {
                            _hx_tmp_12 = struct { Value types.Tuple; Valid bool }{}
                        }
                    
                        var results struct { Value types.Tuple; Valid bool } = _hx_tmp_12; _ = results
                        var _hx_tmp_13 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_13
                        if ((results.Valid != false)) {
                            _hx_tmp_13 = results
                        } else {
                            _hx_tmp_13 = struct { Value types.Tuple; Valid bool }{}
                        }
                    
                        var results1 struct { Value types.Tuple; Valid bool } = _hx_tmp_13; _ = results1
                        var varadic bool = sig.Variadic(); _ = varadic
                        {
                            var _hx_tmp_14 string = name; _ = _hx_tmp_14
                            var _hx_tmp_15 types.Signature = sig; _ = _hx_tmp_15
                            var _hx_tmp_16 bool = (recv == nil); _ = _hx_tmp_16
                            var x string = (("    " + Hx_Field_main_genFunc(_hx_tmp_14, _hx_tmp_15, _hx_tmp_16, struct { Value bool; Valid bool }{}.Value)) + "\n"); _ = x
                            var _hx_tmp_17 string = buf.Hx_Field_b; _ = _hx_tmp_17
                            buf.Hx_Field_b = (_hx_tmp_17 + Hx_Field_std_string(x))
                        }
                    }
                
                    case *types.Var:
                    {
                        var v2 *types.Var = (((any)(obj))).(*types.Var); _ = v2
                        var v3 types.Var = (*v2); _ = v3
                        var buf1 *Hx_Obj_stringbuf = (Hx_Field_go_haxe_hxdynamic_toClass(Hx_Field_go_haxe_hxdynamic_getField(getOutput((*entry).Name), "staticVars"), "Hx_Obj_stringbuf")).(*Hx_Obj_stringbuf); _ = buf1
                        var name1 string = v3.Name(); _ = name1
                        var type1 types.Type = v3.Type(); _ = type1
                        {
                            var _hx_tmp_10 string = (("    static var " + name1) + ": "); _ = _hx_tmp_10
                            var x1 string = ((_hx_tmp_10 + Hx_Field_main_genType(type1)) + ";\n"); _ = x1
                            var _hx_tmp_11 string = buf1.Hx_Field_b; _ = _hx_tmp_11
                            buf1.Hx_Field_b = (_hx_tmp_11 + Hx_Field_std_string(x1))
                        }
                    }
                
                    default:
                    }
                }
            }
        }
    }

    {
        var length struct { Value int; Valid bool } = struct { Value int; Valid bool }{}; _ = length
        var _hx_tmp_4 int; _ = _hx_tmp_4
        if ((length.Valid != false)) {
            _hx_tmp_4 = length.Value
        } else {
            _hx_tmp_4 = 0
        }
    
        var keys []string = make([]string, _hx_tmp_4); _ = keys
        for k := range outputs {keys = append(keys, k)}
    
        var self1 []string = keys; _ = self1
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
            var _hx_tmp_10 any = ((any)(("extern class " + Hx_Field_main_toPascalCase(file)))); _ = _hx_tmp_10
            buf_b = (_hx_tmp_9 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_10, Hx_Field_go_haxe_hxdynamic_getField(out, "paramStr")), ((any)(" {\n\n")))))
            var _hx_tmp_11 string = buf_b; _ = _hx_tmp_11
            buf_b = (_hx_tmp_11 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "staticVars"), "b")))
            var _hx_tmp_12 string = buf_b; _ = _hx_tmp_12
            buf_b = (_hx_tmp_12 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "instanceVars"), "b")))
            var _hx_tmp_13 bool; _ = _hx_tmp_13
            if (Hx_Field_go_haxe_hxdynamic_gt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "staticVars"), "b"), "length"), ((any)(0)))) {
                _hx_tmp_13 = true
            } else {
                _hx_tmp_13 = Hx_Field_go_haxe_hxdynamic_gt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "instanceVars"), "b"), "length"), ((any)(0)))
            }
        
            if (_hx_tmp_13) {
                buf_b = (buf_b + "\n")
            }
        
            var _hx_tmp_14 string = buf_b; _ = _hx_tmp_14
            buf_b = (_hx_tmp_14 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "staticFunctions"), "b")))
            var _hx_tmp_15 string = buf_b; _ = _hx_tmp_15
            buf_b = (_hx_tmp_15 + Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "instanceFunctions"), "b")))
            var _hx_tmp_16 bool; _ = _hx_tmp_16
            if (Hx_Field_go_haxe_hxdynamic_gt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "staticFunctions"), "b"), "length"), ((any)(0)))) {
                _hx_tmp_16 = true
            } else {
                _hx_tmp_16 = Hx_Field_go_haxe_hxdynamic_gt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(out, "instanceFunctions"), "b"), "length"), ((any)(0)))
            }
        
            if (_hx_tmp_16) {
                buf_b = (buf_b + "\n")
            }
        
            buf_b = (buf_b + "}")
            var _hx_tmp_17 string = ((("" + output) + "/go/") + lib); _ = _hx_tmp_17
            os.MkdirAll(_hx_tmp_17, 0775)
            var _hx_tmp_19 string = (((("" + output) + "/go/") + lib) + "/"); _ = _hx_tmp_19
            var _hx_tmp_18 string = ((_hx_tmp_19 + Hx_Field_main_toPascalCase(file)) + ".hx"); _ = _hx_tmp_18
            var _hx_tmp_20 []byte = (([]byte)(buf_b)); _ = _hx_tmp_20
            os.WriteFile(_hx_tmp_18, _hx_tmp_20, 0666)
        }
    }
}

func Hx_Field_main_genPackage(pkg any) string {
    return ""
}

func Hx_Field_main_genFile(file any) string {
    return ""
}

func Hx_Field_main_genFunc(name string, sig types.Signature, topLevel bool, closure bool) string {
    var recv *types.Var = sig.Recv(); _ = recv
    var tmp *types.Tuple = sig.Params(); _ = tmp
    var _hx_tmp_0 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_0
    if ((tmp != nil)) {
        _hx_tmp_0 = struct { Value types.Tuple; Valid bool }{ Value: (*tmp), Valid: true }
    } else {
        _hx_tmp_0 = struct { Value types.Tuple; Valid bool }{}
    }

    var params struct { Value types.Tuple; Valid bool } = _hx_tmp_0; _ = params
    var _hx_tmp_1 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_1
    if ((params.Valid != false)) {
        _hx_tmp_1 = params
    } else {
        _hx_tmp_1 = struct { Value types.Tuple; Valid bool }{}
    }

    var params1 struct { Value types.Tuple; Valid bool } = _hx_tmp_1; _ = params1
    var tmp1 *types.Tuple = sig.Results(); _ = tmp1
    var _hx_tmp_2 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_2
    if ((tmp1 != nil)) {
        _hx_tmp_2 = struct { Value types.Tuple; Valid bool }{ Value: (*tmp1), Valid: true }
    } else {
        _hx_tmp_2 = struct { Value types.Tuple; Valid bool }{}
    }

    var results struct { Value types.Tuple; Valid bool } = _hx_tmp_2; _ = results
    var _hx_tmp_3 struct { Value types.Tuple; Valid bool }; _ = _hx_tmp_3
    if ((results.Valid != false)) {
        _hx_tmp_3 = results
    } else {
        _hx_tmp_3 = struct { Value types.Tuple; Valid bool }{}
    }

    var results1 struct { Value types.Tuple; Valid bool } = _hx_tmp_3; _ = results1
    var varadic bool = sig.Variadic(); _ = varadic
    var _hx_tmp_4 types.Tuple = params1.Value; _ = _hx_tmp_4
    var _hx_tmp_5 bool = varadic; _ = _hx_tmp_5
    var params2 *[]string = Hx_Field_main_genTuple(_hx_tmp_4, _hx_tmp_5, struct { Value bool; Valid bool }{}.Value); _ = params2
    var meta string = ""; _ = meta
    var _hx_tmp_6 bool; _ = _hx_tmp_6
    var _hx_tmp_7 bool; _ = _hx_tmp_7
    if ((results1.Value.Len() > 1)) {
        _hx_tmp_7 = !Hx_Field_main_isResultType(results1.Value)
    } else {
        _hx_tmp_7 = false
    }

    if (_hx_tmp_7) {
        _hx_tmp_6 = !closure
    } else {
        _hx_tmp_6 = false
    }

    if (_hx_tmp_6) {
        var names *[]string = &([]string{}); _ = names
        var unnamed int = 0; _ = unnamed
        {
            var _g int = 0; _ = _g
            var _g1 int = results1.Value.Len(); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_8 int = _g; _ = _hx_tmp_8
                _g = (_g + 1)
                var idx int = _hx_tmp_8; _ = idx
                var name string = (*results1.Value.At(idx)).Name(); _ = name
                if ((name == "")) {
                    var _hx_tmp_9 int = unnamed; _ = _hx_tmp_9
                    unnamed = (unnamed + 1)
                    name = ("p" + Hx_Field_std_string(_hx_tmp_9))
                }
            
                {
                    var data []string = (*names); _ = data
                    var _hx_tmp_9 *[]string = names; _ = _hx_tmp_9
                    (*_hx_tmp_9) = append(data, (("\"" + name) + "\""))
                    var _hx_tmp_10 int = len(data); _ = _hx_tmp_10
                    var this1 int = (_hx_tmp_10 + int(1)); _ = this1
                }
            }
        }
    
        var data []string = (*names); _ = data
        var length int = len(data); _ = length
        var sep string = ", "; _ = sep
        var meta1 string; _ = meta1
        var _hx_tmp_8 int = length; _ = _hx_tmp_8
        if ((_hx_tmp_8 == int(0))) {
            meta1 = ""
        } else {
            var result string = ""; _ = result
            var i int = int(0); _ = i
            for ((i < length)) {
                var _hx_tmp_9 string = result; _ = _hx_tmp_9
                result = (_hx_tmp_9 + Hx_Field_std_string(data[((int)(i))]))
                var _hx_tmp_10 int = i; _ = _hx_tmp_10
                var _hx_tmp_11 int = length; _ = _hx_tmp_11
                if ((_hx_tmp_10 < (_hx_tmp_11 - int(1)))) {
                    result = (result + sep)
                }
            
                i = (i + ((int)(1)))
            }
        
            meta1 = result
        }
    
        meta = (("@:go.Tuple(" + meta1) + ") ")
    }

    var _hx_tmp_8 struct { Value types.TypeParamList; Valid bool }; _ = _hx_tmp_8
    if ((sig.TypeParams() != nil)) {
        _hx_tmp_8 = struct { Value types.TypeParamList; Valid bool }{ Value: (*sig.TypeParams()), Valid: true }
    } else {
        _hx_tmp_8 = struct { Value types.TypeParamList; Valid bool }{}
    }

    var tParams struct { Value types.TypeParamList; Valid bool } = _hx_tmp_8; _ = tParams
    var tParamsStr string = ""; _ = tParamsStr
    if ((tParams.Valid != false)) {
        var tParamsLocal *[]string = &([]string{}); _ = tParamsLocal
        {
            var _g int = 0; _ = _g
            var _g1 int = tParams.Value.Len(); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_9 int = _g; _ = _hx_tmp_9
                _g = (_g + 1)
                var i int = _hx_tmp_9; _ = i
                var t types.TypeParam = (*tParams.Value.At(i)); _ = t
                var constraint types.Type = t.Constraint(); _ = constraint
                var constraintStr string = Hx_Field_main_genType(constraint); _ = constraintStr
                if (Hx_Field_stringtools_startsWith(constraintStr, "~")) {
                    constraintStr = "Dynamic"
                }
            
                {
                    var x string = ((("" + t.String()) + ": ") + constraintStr); _ = x
                    {
                        var data []string = (*tParamsLocal); _ = data
                        var _hx_tmp_10 *[]string = tParamsLocal; _ = _hx_tmp_10
                        (*_hx_tmp_10) = append(data, x)
                        var _hx_tmp_11 int = len(data); _ = _hx_tmp_11
                        var this1 int = (_hx_tmp_11 + int(1)); _ = this1
                    }
                }
            }
        }
    
        var data []string = (*tParamsLocal); _ = data
        var length int = len(data); _ = length
        var sep string = ", "; _ = sep
        var tParamsStr1 string; _ = tParamsStr1
        var _hx_tmp_9 int = length; _ = _hx_tmp_9
        if ((_hx_tmp_9 == int(0))) {
            tParamsStr1 = ""
        } else {
            var result string = ""; _ = result
            var i int = int(0); _ = i
            for ((i < length)) {
                var _hx_tmp_10 string = result; _ = _hx_tmp_10
                result = (_hx_tmp_10 + Hx_Field_std_string(data[((int)(i))]))
                var _hx_tmp_11 int = i; _ = _hx_tmp_11
                var _hx_tmp_12 int = length; _ = _hx_tmp_12
                if ((_hx_tmp_11 < (_hx_tmp_12 - int(1)))) {
                    result = (result + sep)
                }
            
                i = (i + ((int)(1)))
            }
        
            tParamsStr1 = result
        }
    
        tParamsStr = (("<" + tParamsStr1) + ">")
    }

    var _hx_tmp_9 string; _ = _hx_tmp_9
    if (closure) {
        _hx_tmp_9 = ""
    } else {
        _hx_tmp_9 = ("function " + Hx_Field_main_toHaxeCase(name))
    }

    var tmp2 string = _hx_tmp_9; _ = tmp2
    var _hx_tmp_10 string = ("" + meta); _ = _hx_tmp_10
    var _hx_tmp_11 string; _ = _hx_tmp_11
    if ((topLevel && !closure)) {
        _hx_tmp_11 = "static "
    } else {
        _hx_tmp_11 = ""
    }

    var tmp3 string = ((((_hx_tmp_10 + (_hx_tmp_11)) + tmp2) + tParamsStr) + "("); _ = tmp3
    var data []string = (*params2); _ = data
    var length int = len(data); _ = length
    var sep string = ", "; _ = sep
    var tmp4 string; _ = tmp4
    var _hx_tmp_12 int = length; _ = _hx_tmp_12
    if ((_hx_tmp_12 == int(0))) {
        tmp4 = ""
    } else {
        var result string = ""; _ = result
        var i int = int(0); _ = i
        for ((i < length)) {
            var _hx_tmp_13 string = result; _ = _hx_tmp_13
            result = (_hx_tmp_13 + Hx_Field_std_string(data[((int)(i))]))
            var _hx_tmp_14 int = i; _ = _hx_tmp_14
            var _hx_tmp_15 int = length; _ = _hx_tmp_15
            if ((_hx_tmp_14 < (_hx_tmp_15 - int(1)))) {
                result = (result + sep)
            }
        
            i = (i + ((int)(1)))
        }
    
        tmp4 = result
    }

    var _hx_tmp_13 string; _ = _hx_tmp_13
    if ((results1.Value.Len() == 0)) {
        _hx_tmp_13 = "Void"
    } else {
        _hx_tmp_13 = Hx_Field_main_genResults(results1.Value)
    }

    var tmp5 string = _hx_tmp_13; _ = tmp5
    var _hx_tmp_14 string = ((tmp3 + tmp4) + ")"); _ = _hx_tmp_14
    var _hx_tmp_15 string; _ = _hx_tmp_15
    if (closure) {
        _hx_tmp_15 = " -> "
    } else {
        _hx_tmp_15 = ": "
    }

    var _hx_tmp_16 string; _ = _hx_tmp_16
    if (closure) {
        _hx_tmp_16 = ""
    } else {
        _hx_tmp_16 = ";"
    }

    return (((_hx_tmp_14 + (_hx_tmp_15)) + tmp5) + (_hx_tmp_16))
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
                        var _hx_tmp_2 string = (n + ": haxe.Rest<"); _ = _hx_tmp_2
                        x = ((_hx_tmp_2 + Hx_Field_main_genType((*v1).Elem())) + ">")
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

func Hx_Field_main_resolvePath(path string) string {
    return ("go." + Hx_Field_stringtools_replace(path, "/", "."))
}

func Hx_Field_main_isNamedType(t types.Type) bool {
    var isNamed bool = false; _ = isNamed
    if _, ntOk := t.(*types.Named); ntOk { isNamed = true; }

    return isNamed
}

func Hx_Field_main_genType(t types.Type) string {
    var s string = t.String(); _ = s
    var tParamStr string = ""; _ = tParamStr
    var tmp bool; _ = tmp
    if (Hx_Field_main_isNamedType(t)) {
        var v *types.Named = (((any)(t))).(*types.Named); _ = v
        tmp = ((*v).TypeParams() != nil)
    } else {
        tmp = false
    }

    if (tmp) {
        var v *types.Named = (((any)(t))).(*types.Named); _ = v
        var typeParams types.TypeList = (*(*v).TypeArgs()); _ = typeParams
        var typeParamStrs *[]string = &([]string{}); _ = typeParamStrs
        {
            var _g int = 0; _ = _g
            var _g1 int = typeParams.Len(); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_0 int = _g; _ = _hx_tmp_0
                _g = (_g + 1)
                var i int = _hx_tmp_0; _ = i
                var r string = Hx_Field_main_genType(typeParams.At(i)); _ = r
                {
                    var _hx_tmp_1 string; _ = _hx_tmp_1
                    if (Hx_Field_stringtools_startsWith(r, "~")) {
                        _hx_tmp_1 = "Dynamic"
                    } else {
                        _hx_tmp_1 = r
                    }
                
                    var x string = _hx_tmp_1; _ = x
                    {
                        var data []string = (*typeParamStrs); _ = data
                        var _hx_tmp_2 *[]string = typeParamStrs; _ = _hx_tmp_2
                        (*_hx_tmp_2) = append(data, x)
                        var _hx_tmp_3 int = len(data); _ = _hx_tmp_3
                        var this1 int = (_hx_tmp_3 + int(1)); _ = this1
                    }
                }
            }
        }
    
        var _hx_tmp_0 string = s; _ = _hx_tmp_0
        var typeParamStart int = Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "[", struct { Value int; Valid bool }{}); _ = typeParamStart
        var _hx_tmp_1 string = s; _ = _hx_tmp_1
        var typeParamEnd int = Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_1, "]", struct { Value int; Valid bool }{}); _ = typeParamEnd
        if ((((typeParamStart != -1) && (typeParamEnd != -1)) && (typeParamEnd > typeParamStart))) {
            var _hx_tmp_3 string = s; _ = _hx_tmp_3
            var _hx_tmp_2 string = Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 0, struct { Value int; Valid bool }{ Value: typeParamStart, Valid: true }); _ = _hx_tmp_2
            var _hx_tmp_4 string = s; _ = _hx_tmp_4
            var _hx_tmp_5 int = (typeParamEnd + 1); _ = _hx_tmp_5
            s = (_hx_tmp_2 + Hx_Field_go_haxe_hxstring_substr(_hx_tmp_4, _hx_tmp_5, struct { Value int; Valid bool }{}))
        }
    
        var data []string = (*typeParamStrs); _ = data
        var length int = len(data); _ = length
        var sep string = ", "; _ = sep
        var tParamStr1 string; _ = tParamStr1
        var _hx_tmp_2 int = length; _ = _hx_tmp_2
        if ((_hx_tmp_2 == int(0))) {
            tParamStr1 = ""
        } else {
            var result string = ""; _ = result
            var i int = int(0); _ = i
            for ((i < length)) {
                var _hx_tmp_3 string = result; _ = _hx_tmp_3
                result = (_hx_tmp_3 + Hx_Field_std_string(data[((int)(i))]))
                var _hx_tmp_4 int = i; _ = _hx_tmp_4
                var _hx_tmp_5 int = length; _ = _hx_tmp_5
                if ((_hx_tmp_4 < (_hx_tmp_5 - int(1)))) {
                    result = (result + sep)
                }
            
                i = (i + ((int)(1)))
            }
        
            tParamStr1 = result
        }
    
        tParamStr = (("<" + tParamStr1) + ">")
    }

    var q string; _ = q
    switch (s) {
        case "any":
            q = "Dynamic"
    
        case "bool":
            q = "Bool"
    
        case "byte":
            q = "go.Byte"
    
        case "comparable":
            q = "go.Comparable"
    
        case "complex128":
            q = "go.Complex128"
    
        case "complex64":
            q = "go.Complex64"
    
        case "error":
            q = "go.Error"
    
        case "float16":
            q = "go.Float16"
    
        case "float32":
            q = "go.Float32"
    
        case "float64":
            q = "Float"
    
        case "int":
            q = "Int"
    
        case "int16":
            q = "go.Int16"
    
        case "int32":
            q = "go.Int32"
    
        case "int64":
            q = "go.Int64"
    
        case "int8":
            q = "go.Int8"
    
        case "rune":
            q = "go.Rune"
    
        case "string":
            q = "String"
    
        case "uint":
            q = "go.UInt"
    
        case "uint16":
            q = "go.UInt16"
    
        case "uint32":
            q = "go.UInt32"
    
        case "uint64":
            q = "go.UInt64"
    
        case "uint8":
            q = "go.UInt8"
    
        case "uintptr":
            q = "go.UIntPtr"
    
        default: 
            if (Hx_Field_stringtools_startsWith(s, "chan ")) {
                var v *types.Chan = (((any)(t))).(*types.Chan); _ = v
                q = (("go.Chan<" + Hx_Field_main_genType((*v).Elem())) + ">")
            } else {
                if (Hx_Field_stringtools_startsWith(s, "chan ")) {
                    var v *types.Chan = (((any)(t))).(*types.Chan); _ = v
                    q = (("go.Chan<" + Hx_Field_main_genType((*v).Elem())) + ">")
                } else {
                    if (Hx_Field_stringtools_startsWith(s, "[]")) {
                        var v *types.Slice = (((any)(t))).(*types.Slice); _ = v
                        q = (("go.Slice<" + Hx_Field_main_genType((*v).Elem())) + ">")
                    } else {
                        if (Hx_Field_stringtools_startsWith(s, "*")) {
                            var v *types.Pointer = (((any)(t))).(*types.Pointer); _ = v
                            q = (("go.Pointer<" + Hx_Field_main_genType((*v).Elem())) + ">")
                        } else {
                            if (Hx_Field_stringtools_startsWith(s, "func")) {
                                var v *types.Signature = (((any)(t))).(*types.Signature); _ = v
                                var sig types.Signature = (*v); _ = sig
                                q = Hx_Field_main_genFunc("", sig, false, true)
                            } else {
                                var _hx_tmp_0 string; _ = _hx_tmp_0
                                var _hx_tmp_1 string = s; _ = _hx_tmp_1
                                if ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_1, ".", struct { Value int; Valid bool }{}) != -1)) {
                                    _hx_tmp_0 = Hx_Field_main_resolvePath(s)
                                } else {
                                    _hx_tmp_0 = t.String()
                                }
                            
                                q = _hx_tmp_0
                            }
                        }
                    }
                }
            }
    }

    return (q + tParamStr)
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
