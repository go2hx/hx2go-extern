package main

type Hx_Enum_haxe_stackitem interface {
    Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
    M_Hx_Enum_haxe_stackitem()
}

type Hx_Enum_haxe_stackitem_Module struct {
    Hx_Field_m string
}

func (this Hx_Enum_haxe_stackitem_Module) M_Hx_Enum_haxe_stackitem() {}
func (this Hx_Enum_haxe_stackitem_Module) Hx_Field_enumIndex() int { return 1 }
func (this Hx_Enum_haxe_stackitem_Module) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum { return Hx_Enum_haxe_stackitem_RTTI }
func (this Hx_Enum_haxe_stackitem_Module) Hx_Field_enumParams() any { return &([]any{ any(this.Hx_Field_m) }) }
func (this Hx_Enum_haxe_stackitem_Module) Hx_Field_enumParameter(index int) any {
    switch index {
        case 0: return any(this.Hx_Field_m)
        default: return nil
    }
}

type Hx_Enum_haxe_stackitem_Method struct {
    Hx_Field_classname struct { Value string; Valid bool }
    Hx_Field_method string
}

func (this Hx_Enum_haxe_stackitem_Method) M_Hx_Enum_haxe_stackitem() {}
func (this Hx_Enum_haxe_stackitem_Method) Hx_Field_enumIndex() int { return 3 }
func (this Hx_Enum_haxe_stackitem_Method) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum { return Hx_Enum_haxe_stackitem_RTTI }
func (this Hx_Enum_haxe_stackitem_Method) Hx_Field_enumParams() any { return &([]any{ any(this.Hx_Field_classname), any(this.Hx_Field_method) }) }
func (this Hx_Enum_haxe_stackitem_Method) Hx_Field_enumParameter(index int) any {
    switch index {
        case 0: return any(this.Hx_Field_classname)
        case 1: return any(this.Hx_Field_method)
        default: return nil
    }
}

type Hx_Enum_haxe_stackitem_LocalFunction struct {
    Hx_Field_v struct { Value int; Valid bool }
}

func (this Hx_Enum_haxe_stackitem_LocalFunction) M_Hx_Enum_haxe_stackitem() {}
func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field_enumIndex() int { return 4 }
func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum { return Hx_Enum_haxe_stackitem_RTTI }
func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field_enumParams() any { return &([]any{ any(this.Hx_Field_v) }) }
func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field_enumParameter(index int) any {
    switch index {
        case 0: return any(this.Hx_Field_v)
        default: return nil
    }
}

type Hx_Enum_haxe_stackitem_FilePos struct {
    Hx_Field_s struct { Value Hx_Enum_haxe_stackitem; Valid bool }
    Hx_Field_file string
    Hx_Field_line int
    Hx_Field_column struct { Value int; Valid bool }
}

func (this Hx_Enum_haxe_stackitem_FilePos) M_Hx_Enum_haxe_stackitem() {}
func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field_enumIndex() int { return 2 }
func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum { return Hx_Enum_haxe_stackitem_RTTI }
func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field_enumParams() any { return &([]any{ any(this.Hx_Field_s), any(this.Hx_Field_file), any(this.Hx_Field_line), any(this.Hx_Field_column) }) }
func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field_enumParameter(index int) any {
    switch index {
        case 0: return any(this.Hx_Field_s)
        case 1: return any(this.Hx_Field_file)
        case 2: return any(this.Hx_Field_line)
        case 3: return any(this.Hx_Field_column)
        default: return nil
    }
}

type Hx_Enum_haxe_stackitem_CFunction struct {}

func (this Hx_Enum_haxe_stackitem_CFunction) M_Hx_Enum_haxe_stackitem() {}
func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field_enumIndex() int { return 0 }
func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum { return Hx_Enum_haxe_stackitem_RTTI }
func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field_enumParams() any { return &([]any{ }) }
func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field_enumParameter(index int) any {
    switch index {
        default: return nil
    }
}

var Hx_Enum_haxe_stackitem_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
    "haxe.StackItem",
    &([]string{ "CFunction", "Module", "FilePos", "Method", "LocalFunction" }),
    &([]int{ 0, 1, 4, 2, 1 }),
    func (index int, params any) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
        return nil
    },
)

var Hx_Obj_haxe__callstack_callstack_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe._CallStack.CallStack_Impl_",
)

type Hx_Obj_VTable_haxe__callstack_callstack_impl_ interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__callstack_callstack_impl_ struct {
    VTable Hx_Obj_VTable_haxe__callstack_callstack_impl_
}

func Hx_Obj_haxe__callstack_callstack_impl__CreateEmptyInstance() *Hx_Obj_haxe__callstack_callstack_impl_ {
    obj := &Hx_Obj_haxe__callstack_callstack_impl_{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe__callstack_callstack_impl__CreateInstance() *Hx_Obj_haxe__callstack_callstack_impl_ {
    obj := Hx_Obj_haxe__callstack_callstack_impl__CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_haxe__callstack_callstack_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe__callstack_callstack_impl__RTTI
}

func Hx_Field_haxe__callstack_callstack_impl__get_length(this *[]Hx_Enum_haxe_stackitem) int {
    return len(*this)
}

func Hx_Field_haxe__callstack_callstack_impl__callStack() *[]Hx_Enum_haxe_stackitem {
    Hx_Field_sys_println("not implemented")
    Hx_Field_sys_println("not implemented")
    return &([]Hx_Enum_haxe_stackitem{})
}

func Hx_Field_haxe__callstack_callstack_impl__exceptionStack(fullStack bool) *[]Hx_Enum_haxe_stackitem {
    Hx_Field_sys_println("not implemented")
    Hx_Field_sys_println("not implemented")
    var eStack *[]Hx_Enum_haxe_stackitem = &([]Hx_Enum_haxe_stackitem{}); _ = eStack
    var _hx_tmp_0 *[]Hx_Enum_haxe_stackitem; _ = _hx_tmp_0
    if (fullStack) {
        _hx_tmp_0 = eStack
    } else {
        var _hx_tmp_1 *[]Hx_Enum_haxe_stackitem = eStack; _ = _hx_tmp_1
        var _hx_tmp_2 *[]Hx_Enum_haxe_stackitem = _hx_tmp_1; _ = _hx_tmp_2
        _hx_tmp_0 = Hx_Field_haxe__callstack_callstack_impl__subtract(_hx_tmp_2, Hx_Field_haxe__callstack_callstack_impl__callStack())
    }

    return _hx_tmp_0
}

func Hx_Field_haxe__callstack_callstack_impl__toString(stack *[]Hx_Enum_haxe_stackitem) string {
    var b *Hx_Obj_stringbuf = Hx_Obj_stringbuf_CreateInstance(); _ = b
    {
        var _g int = 0; _ = _g
        var _g1 *[]Hx_Enum_haxe_stackitem = stack; _ = _g1
        for  {
            var _hx_tmp_0 int = _g; _ = _hx_tmp_0
            if (!((_hx_tmp_0 < len(*_g1)))) {
                break
            }
        
            var s Hx_Enum_haxe_stackitem = (*_g1)[_g]; _ = s
            _g++
            b.Hx_Field_b = (b.Hx_Field_b + "\nCalled from ")
            Hx_Field_haxe__callstack_callstack_impl__itemToString(b, s)
        }
    }

    return b.Hx_Field_b
}

func Hx_Field_haxe__callstack_callstack_impl__subtract(this *[]Hx_Enum_haxe_stackitem, stack *[]Hx_Enum_haxe_stackitem) *[]Hx_Enum_haxe_stackitem {
    var startIndex int = -1; _ = startIndex
    var i int = -1; _ = i
    for  {
        i = (i + 1)
        var _hx_tmp_0 int = i; _ = _hx_tmp_0
        if (!((_hx_tmp_0 < len(*this)))) {
            break
        }
    
        {
            var _g int = 0; _ = _g
            var _g1 int = len(*stack); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_1 int = _g; _ = _hx_tmp_1
                _g = (_g + 1)
                var j int = _hx_tmp_1; _ = j
                var _hx_tmp_2 struct { Value Hx_Enum_haxe_stackitem; Valid bool } = struct { Value Hx_Enum_haxe_stackitem; Valid bool }{ Value: (*this)[i], Valid: true }; _ = _hx_tmp_2
                if (Hx_Field_haxe__callstack_callstack_impl__equalItems(_hx_tmp_2, struct { Value Hx_Enum_haxe_stackitem; Valid bool }{ Value: (*stack)[j], Valid: true })) {
                    if ((startIndex < 0)) {
                        startIndex = i
                    }
                
                    i++
                    var _hx_tmp_3 int = i; _ = _hx_tmp_3
                    if ((_hx_tmp_3 >= len(*this))) {
                        break
                    }
                } else {
                    startIndex = -1
                }
            }
        }
    
        if ((startIndex >= 0)) {
            break
        }
    }

    if ((startIndex >= 0)) {
        var data []Hx_Enum_haxe_stackitem = (*this); _ = data
        var length int = len(data); _ = length
        var start int = int(0); _ = start
        var clampedEnd int; _ = clampedEnd
        if ((struct { Value int; Valid bool }{ Value: startIndex, Valid: true }.Valid == false)) {
            clampedEnd = length
        } else {
            if ((startIndex < 0)) {
                var _hx_tmp_0 int = length; _ = _hx_tmp_0
                var e int = (_hx_tmp_0 + int(startIndex)); _ = e
                var _hx_tmp_1 int; _ = _hx_tmp_1
                var _hx_tmp_2 int = e; _ = _hx_tmp_2
                if ((_hx_tmp_2 < int(0))) {
                    _hx_tmp_1 = int(0)
                } else {
                    _hx_tmp_1 = e
                }
            
                clampedEnd = _hx_tmp_1
            } else {
                clampedEnd = int(startIndex)
            }
        }
    
        var _hx_tmp_0 int; _ = _hx_tmp_0
        if ((clampedEnd > length)) {
            _hx_tmp_0 = length
        } else {
            _hx_tmp_0 = clampedEnd
        }
    
        var stop int = _hx_tmp_0; _ = stop
        if (((start > length) || (stop <= start))) {
            return &([]Hx_Enum_haxe_stackitem{})
        } else {
            var result *[]Hx_Enum_haxe_stackitem = &([]Hx_Enum_haxe_stackitem{}); _ = result
            var _hx_tmp_1 *[]Hx_Enum_haxe_stackitem = result; _ = _hx_tmp_1
            var _hx_tmp_2 []Hx_Enum_haxe_stackitem = (*result); _ = _hx_tmp_2
            (*_hx_tmp_1) = append(_hx_tmp_2, (*this)[start:stop]...)
            return result
        }
    } else {
        return this
    }
}

func Hx_Field_haxe__callstack_callstack_impl__copy(this *[]Hx_Enum_haxe_stackitem) *[]Hx_Enum_haxe_stackitem {
    var newArr *[]Hx_Enum_haxe_stackitem = &([]Hx_Enum_haxe_stackitem{}); _ = newArr
    var _hx_tmp_0 *[]Hx_Enum_haxe_stackitem = newArr; _ = _hx_tmp_0
    var _hx_tmp_1 []Hx_Enum_haxe_stackitem = (*newArr); _ = _hx_tmp_1
    (*_hx_tmp_0) = append(_hx_tmp_1, (*this)...)
    return newArr
}

func Hx_Field_haxe__callstack_callstack_impl__get(this *[]Hx_Enum_haxe_stackitem, index int) Hx_Enum_haxe_stackitem {
    return (*this)[index]
}

func Hx_Field_haxe__callstack_callstack_impl__asArray(this *[]Hx_Enum_haxe_stackitem) *[]Hx_Enum_haxe_stackitem {
    return this
}

func Hx_Field_haxe__callstack_callstack_impl__equalItems(item1 struct { Value Hx_Enum_haxe_stackitem; Valid bool }, item2 struct { Value Hx_Enum_haxe_stackitem; Valid bool }) bool {
    if ((item1.Valid == false)) {
        return (item2.Valid == false)
    } else {
        switch (item1.Value.Hx_Field_enumIndex()) {
            case 0:
                if ((item2.Valid == false)) {
                    return false
                } else {
                    return (item2.Value.Hx_Field_enumIndex() == 0)
                }
        
            case 1:
                if ((item2.Valid == false)) {
                    return false
                } else {
                    if ((item2.Value.Hx_Field_enumIndex() == 1)) {
                        var m2 string = (item2.Value.Hx_Field_enumParameter(0)).(string); _ = m2
                        var m1 string = (item1.Value.Hx_Field_enumParameter(0)).(string); _ = m1
                        return (m1 == m2)
                    } else {
                        return false
                    }
                }
        
            case 2:
                if ((item2.Valid == false)) {
                    return false
                } else {
                    if ((item2.Value.Hx_Field_enumIndex() == 2)) {
                        var item21 struct { Value Hx_Enum_haxe_stackitem; Valid bool } = (item2.Value.Hx_Field_enumParameter(0)).(struct { Value Hx_Enum_haxe_stackitem; Valid bool }); _ = item21
                        var file2 string = (item2.Value.Hx_Field_enumParameter(1)).(string); _ = file2
                        var line2 int = (item2.Value.Hx_Field_enumParameter(2)).(int); _ = line2
                        var col2 struct { Value int; Valid bool } = (item2.Value.Hx_Field_enumParameter(3)).(struct { Value int; Valid bool }); _ = col2
                        var col1 struct { Value int; Valid bool } = (item1.Value.Hx_Field_enumParameter(3)).(struct { Value int; Valid bool }); _ = col1
                        var line1 int = (item1.Value.Hx_Field_enumParameter(2)).(int); _ = line1
                        var file1 string = (item1.Value.Hx_Field_enumParameter(1)).(string); _ = file1
                        var item11 struct { Value Hx_Enum_haxe_stackitem; Valid bool } = (item1.Value.Hx_Field_enumParameter(0)).(struct { Value Hx_Enum_haxe_stackitem; Valid bool }); _ = item11
                        var _hx_tmp_0 bool; _ = _hx_tmp_0
                        var _hx_tmp_1 bool; _ = _hx_tmp_1
                        if (((file1 == file2) && (line1 == line2))) {
                            var _hx_tmp_2 int = col1.Value; _ = _hx_tmp_2
                            var _hx_tmp_3 int = _hx_tmp_2; _ = _hx_tmp_3
                            _hx_tmp_1 = (_hx_tmp_3 == col2.Value)
                        } else {
                            _hx_tmp_1 = false
                        }
                    
                        if (_hx_tmp_1) {
                            _hx_tmp_0 = Hx_Field_haxe__callstack_callstack_impl__equalItems(item11, item21)
                        } else {
                            _hx_tmp_0 = false
                        }
                    
                        return _hx_tmp_0
                    } else {
                        return false
                    }
                }
        
            case 3:
                if ((item2.Valid == false)) {
                    return false
                } else {
                    if ((item2.Value.Hx_Field_enumIndex() == 3)) {
                        var class2 struct { Value string; Valid bool } = (item2.Value.Hx_Field_enumParameter(0)).(struct { Value string; Valid bool }); _ = class2
                        var method2 string = (item2.Value.Hx_Field_enumParameter(1)).(string); _ = method2
                        var method1 string = (item1.Value.Hx_Field_enumParameter(1)).(string); _ = method1
                        var class1 struct { Value string; Valid bool } = (item1.Value.Hx_Field_enumParameter(0)).(struct { Value string; Valid bool }); _ = class1
                        var _hx_tmp_0 bool; _ = _hx_tmp_0
                        var _hx_tmp_1 string = class1.Value; _ = _hx_tmp_1
                        if ((_hx_tmp_1 == class2.Value)) {
                            _hx_tmp_0 = (method1 == method2)
                        } else {
                            _hx_tmp_0 = false
                        }
                    
                        return _hx_tmp_0
                    } else {
                        return false
                    }
                }
        
            case 4:
                if ((item2.Valid == false)) {
                    return false
                } else {
                    if ((item2.Value.Hx_Field_enumIndex() == 4)) {
                        var v2 struct { Value int; Valid bool } = (item2.Value.Hx_Field_enumParameter(0)).(struct { Value int; Valid bool }); _ = v2
                        var v1 struct { Value int; Valid bool } = (item1.Value.Hx_Field_enumParameter(0)).(struct { Value int; Valid bool }); _ = v1
                        var _hx_tmp_0 int = v1.Value; _ = _hx_tmp_0
                        return (_hx_tmp_0 == v2.Value)
                    } else {
                        return false
                    }
                }
        
            default: 
                panic("exhaustiveness check mismatch, you shouldn't be able to reach this! please report!")
        }
    }
}

func Hx_Field_haxe__callstack_callstack_impl__exceptionToString(e *Hx_Obj_haxe_exception) string {
    if ((e.VTable.Hx_Field_get_previous().Valid == false)) {
        var tmp string = ("Exception: " + e.VTable.Hx_Field_toString()); _ = tmp
        var tmp1 *[]Hx_Enum_haxe_stackitem = e.VTable.Hx_Field_get_stack(); _ = tmp1
        var _hx_tmp_0 string = tmp; _ = _hx_tmp_0
        var _hx_tmp_1 string; _ = _hx_tmp_1
        if ((tmp1 == nil)) {
            _hx_tmp_1 = "null"
        } else {
            _hx_tmp_1 = Hx_Field_haxe__callstack_callstack_impl__toString(tmp1)
        }
    
        return (_hx_tmp_0 + (_hx_tmp_1))
    }

    var result string = ""; _ = result
    var e1 struct { Value *Hx_Obj_haxe_exception; Valid bool } = struct { Value *Hx_Obj_haxe_exception; Valid bool }{ Value: e, Valid: true }; _ = e1
    var prev struct { Value *Hx_Obj_haxe_exception; Valid bool } = struct { Value *Hx_Obj_haxe_exception; Valid bool }{}; _ = prev
    for  {
        if (!((e1.Valid != false))) {
            break
        }
    
        if ((prev.Valid == false)) {
            var result1 string = ("Exception: " + e1.Value.VTable.Hx_Field_get_message()); _ = result1
            var tmp *[]Hx_Enum_haxe_stackitem = e1.Value.VTable.Hx_Field_get_stack(); _ = tmp
            var _hx_tmp_0 string = result1; _ = _hx_tmp_0
            var _hx_tmp_1 string; _ = _hx_tmp_1
            if ((tmp == nil)) {
                _hx_tmp_1 = "null"
            } else {
                _hx_tmp_1 = Hx_Field_haxe__callstack_callstack_impl__toString(tmp)
            }
        
            result = ((_hx_tmp_0 + (_hx_tmp_1)) + result)
        } else {
            var _hx_tmp_0 *[]Hx_Enum_haxe_stackitem = e1.Value.VTable.Hx_Field_get_stack(); _ = _hx_tmp_0
            var prevStack *[]Hx_Enum_haxe_stackitem = Hx_Field_haxe__callstack_callstack_impl__subtract(_hx_tmp_0, prev.Value.VTable.Hx_Field_get_stack()); _ = prevStack
            var _hx_tmp_1 string = ("Exception: " + e1.Value.VTable.Hx_Field_get_message()); _ = _hx_tmp_1
            var _hx_tmp_2 string; _ = _hx_tmp_2
            if ((prevStack == nil)) {
                _hx_tmp_2 = "null"
            } else {
                _hx_tmp_2 = Hx_Field_haxe__callstack_callstack_impl__toString(prevStack)
            }
        
            result = (((_hx_tmp_1 + (_hx_tmp_2)) + "\n\nNext ") + result)
        }
    
        prev = e1
        e1 = e1.Value.VTable.Hx_Field_get_previous()
    }

    return result
}

func Hx_Field_haxe__callstack_callstack_impl__itemToString(b *Hx_Obj_stringbuf, s Hx_Enum_haxe_stackitem) {
    switch (s.Hx_Field_enumIndex()) {
        case 0:
            b.Hx_Field_b = (b.Hx_Field_b + "a C function")
    
        case 1:
            {
                var m string = (s.Hx_Field_enumParameter(0)).(string); _ = m
                {
                    b.Hx_Field_b = (b.Hx_Field_b + "module ")
                    var _hx_tmp_0 string = b.Hx_Field_b; _ = _hx_tmp_0
                    b.Hx_Field_b = (_hx_tmp_0 + Hx_Field_std_string(m))
                }
            }
    
        case 2:
            {
                var s1 struct { Value Hx_Enum_haxe_stackitem; Valid bool } = (s.Hx_Field_enumParameter(0)).(struct { Value Hx_Enum_haxe_stackitem; Valid bool }); _ = s1
                var file string = (s.Hx_Field_enumParameter(1)).(string); _ = file
                var line int = (s.Hx_Field_enumParameter(2)).(int); _ = line
                var col struct { Value int; Valid bool } = (s.Hx_Field_enumParameter(3)).(struct { Value int; Valid bool }); _ = col
                {
                    if ((s1.Valid != false)) {
                        var _hx_tmp_0 *Hx_Obj_stringbuf = b; _ = _hx_tmp_0
                        Hx_Field_haxe__callstack_callstack_impl__itemToString(_hx_tmp_0, s1.Value)
                        b.Hx_Field_b = (b.Hx_Field_b + " (")
                    }
                
                    var _hx_tmp_0 string = b.Hx_Field_b; _ = _hx_tmp_0
                    b.Hx_Field_b = (_hx_tmp_0 + Hx_Field_std_string(file))
                    b.Hx_Field_b = (b.Hx_Field_b + " line ")
                    var _hx_tmp_1 string = b.Hx_Field_b; _ = _hx_tmp_1
                    b.Hx_Field_b = (_hx_tmp_1 + Hx_Field_std_string(line))
                    if ((col.Valid != false)) {
                        b.Hx_Field_b = (b.Hx_Field_b + " column ")
                        var _hx_tmp_2 string = b.Hx_Field_b; _ = _hx_tmp_2
                        b.Hx_Field_b = (_hx_tmp_2 + Hx_Field_std_string(col))
                    }
                
                    if ((s1.Valid != false)) {
                        b.Hx_Field_b = (b.Hx_Field_b + ")")
                    }
                }
            }
    
        case 3:
            {
                var cname struct { Value string; Valid bool } = (s.Hx_Field_enumParameter(0)).(struct { Value string; Valid bool }); _ = cname
                var meth string = (s.Hx_Field_enumParameter(1)).(string); _ = meth
                {
                    var _hx_tmp_0 string = b.Hx_Field_b; _ = _hx_tmp_0
                    var _hx_tmp_1 struct { Value string; Valid bool }; _ = _hx_tmp_1
                    if ((cname.Valid == false)) {
                        _hx_tmp_1 = struct { Value string; Valid bool }{ Value: "<unknown>", Valid: true }
                    } else {
                        _hx_tmp_1 = cname
                    }
                
                    b.Hx_Field_b = (_hx_tmp_0 + Hx_Field_std_string(_hx_tmp_1))
                    b.Hx_Field_b = (b.Hx_Field_b + ".")
                    var _hx_tmp_2 string = b.Hx_Field_b; _ = _hx_tmp_2
                    b.Hx_Field_b = (_hx_tmp_2 + Hx_Field_std_string(meth))
                }
            }
    
        case 4:
            {
                var n struct { Value int; Valid bool } = (s.Hx_Field_enumParameter(0)).(struct { Value int; Valid bool }); _ = n
                {
                    b.Hx_Field_b = (b.Hx_Field_b + "local function #")
                    var _hx_tmp_0 string = b.Hx_Field_b; _ = _hx_tmp_0
                    b.Hx_Field_b = (_hx_tmp_0 + Hx_Field_std_string(n))
                }
            }
    
    }
}
