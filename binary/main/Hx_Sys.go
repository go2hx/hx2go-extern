package main

import "fmt"
import "os"
import "time"
import "runtime"
import "os/exec"

var Hx_Obj_sys_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "Sys",
)

type Hx_Obj_VTable_sys interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_sys struct {
    VTable Hx_Obj_VTable_sys
}

func Hx_Obj_sys_CreateEmptyInstance() *Hx_Obj_sys {
    obj := &Hx_Obj_sys{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_sys_CreateInstance() *Hx_Obj_sys {
    obj := Hx_Obj_sys_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_sys) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_sys_RTTI
}

func Hx_Field_sys_print(v any) {
    fmt.Print(Hx_Field_std_string(v))
}

func Hx_Field_sys_println(v any) {
    fmt.Println(Hx_Field_std_string(v))
}

func Hx_Field_sys_args() *[]string {
    var args *[]string = &([]string{}); _ = args
    var _hx_tmp_0 *[]string = args; _ = _hx_tmp_0
    (*_hx_tmp_0) = os.Args
    {
        var data []string = (*args); _ = data
        var _hx_tmp_1 int = len(data); _ = _hx_tmp_1
        if ((_hx_tmp_1 != int(0))) {
            var first string = data[0]; _ = first
            var _hx_tmp_2 *[]string = args; _ = _hx_tmp_2
            var _hx_tmp_3 []string = data; _ = _hx_tmp_3
            (*_hx_tmp_2) = _hx_tmp_3[int(1):]
        }
    }

    return args
}

func Hx_Field_sys_getEnv(s string) string {
    return os.Getenv(s)
}

func Hx_Field_sys_putEnv(s string, v struct { Value string; Valid bool }) {
    if ((v.Valid == false)) {
        var _this error = os.Unsetenv(s); _ = _this
        if ((_this != nil)) {
            panic(_this.Error())
        }
    } else {
        var _hx_tmp_0 string = s; _ = _hx_tmp_0
        var _this error = os.Setenv(_hx_tmp_0, v.Value); _ = _this
        if ((_this != nil)) {
            panic(_this.Error())
        }
    }
}

func Hx_Field_sys_sleep(seconds float64) {
    var _hx_tmp_0 int64 = ((int64)(time.Millisecond)); _ = _hx_tmp_0
    time.Sleep(((time.Duration)((_hx_tmp_0 * int64((seconds * ((float64)(1000))))))))
}

func Hx_Field_sys_setTimeLocale(loc string) bool {
    var hx_result_2 struct { Error error; Result *time.Location }; _ = hx_result_2
    hx_result_2.Result, hx_result_2.Error = time.LoadLocation(loc)
    var locale struct { Error error; Result *time.Location } = hx_result_2; _ = locale
    var tmp bool; _ = tmp
    if ((((struct { Error error; Result *time.Location })(locale)).Error != nil)) {
        tmp = false
    } else {
        tmp = true
    }

    if (!tmp) {
        return false
    }

    var tmp1 *time.Location; _ = tmp1
    if ((((struct { Error error; Result *time.Location })(locale)).Error != nil)) {
        var e error = ((struct { Error error; Result *time.Location })(locale)).Error; _ = e
        panic(e)
    } else {
        var r *time.Location = ((struct { Error error; Result *time.Location })(locale)).Result; _ = r
        tmp1 = r
    }

    time.Local = tmp1
    return true
}

func Hx_Field_sys_getCwd() string {
    var hx_result_3 struct { Error error; Result string }; _ = hx_result_3
    hx_result_3.Result, hx_result_3.Error = os.Getwd()
    var this1 struct { Error error; Result string } = ((struct { Error error; Result string })(hx_result_3)); _ = this1
    if ((this1.Error != nil)) {
        var e error = this1.Error; _ = e
        panic(e)
    } else {
        var r string = this1.Result; _ = r
        return r
    }
}

func Hx_Field_sys_setCwd(s string) {
    var _this error = os.Chdir(s); _ = _this
    if ((_this != nil)) {
        panic(_this.Error())
    }
}

func Hx_Field_sys_systemName() string {
    if ((runtime.GOOS == "windows")) {
        return "Windows"
    }

    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if ((runtime.GOOS == "linux")) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = (runtime.GOOS == "android")
    }

    if (_hx_tmp_0) {
        return "Linux"
    }

    var _hx_tmp_1 bool; _ = _hx_tmp_1
    var _hx_tmp_2 bool; _ = _hx_tmp_2
    if ((runtime.GOOS == "freebsd")) {
        _hx_tmp_2 = true
    } else {
        _hx_tmp_2 = (runtime.GOOS == "netbsd")
    }

    if (_hx_tmp_2) {
        _hx_tmp_1 = true
    } else {
        _hx_tmp_1 = (runtime.GOOS == "openbsd")
    }

    if (_hx_tmp_1) {
        return "BSD"
    }

    if ((runtime.GOOS == "darwin")) {
        return "Mac"
    }

    return runtime.GOOS
}

func Hx_Field_sys_command(cmd string, args struct { Value *[]string; Valid bool }) int {
    var arg []string; _ = arg
    if ((args.Valid != false)) {
        var self *[]string = args.Value; _ = self
        arg = (*self)
    } else {
        var length struct { Value int; Valid bool } = struct { Value int; Valid bool }{}; _ = length
        var _hx_tmp_0 int; _ = _hx_tmp_0
        if ((length.Valid != false)) {
            _hx_tmp_0 = length.Value
        } else {
            _hx_tmp_0 = 0
        }
    
        arg = make([]string, _hx_tmp_0)
    }

    var err error = nil; _ = err
    var output []byte = nil; _ = output
    if ((err != nil)) {
        panic(err.Error())
    }

    output, err = exec.Command(cmd, arg...).Output()
    var exitCode int = 0; _ = exitCode
    var exitError *exec.ExitError = nil; _ = exitError
    var ok bool = false; _ = ok
    if ((err != nil)) {
        exitError, ok = err.(*exec.ExitError)
        if (ok) {
            exitCode = exitError.ProcessState.ExitCode()
        } else {
            exitCode = -1
        }
    } else {
        fmt.Println(string(output))
    }

    return 0
}

func Hx_Field_sys_exit(code int) {
    os.Exit(code)
}

func Hx_Field_sys_time() float64 {
    var tn time.Time = time.Now(); _ = tn
    var _hx_tmp_0 int64 = tn.UnixNano(); _ = _hx_tmp_0
    var sec int64 = ((int64)(((float64)((_hx_tmp_0 / ((int64)(time.Second))))))); _ = sec
    var _hx_tmp_1 int64 = sec; _ = _hx_tmp_1
    var hx_tuple_0 struct { Name string; Offset int }; _ = hx_tuple_0
    hx_tuple_0.Name, hx_tuple_0.Offset = tn.Local().Zone()
    sec = (_hx_tmp_1 + int64(((struct { Name string; Offset int })(hx_tuple_0)).Offset))
    var _hx_tmp_2 int64 = time.Now().UnixNano(); _ = _hx_tmp_2
    return ((float64)(((float64)((_hx_tmp_2 / ((int64)(time.Second)))))))
}

func Hx_Field_sys_cpuTime() float64 {
    return Hx_Field_sys_time()
}

func Hx_Field_sys_executablePath() string {
    return Hx_Field_sys_programPath()
}

func Hx_Field_sys_programPath() string {
    return os.Args[0]
}
