package std.go.os.exec;

@:go.Type({ name: "exec", instanceName: "exec.exec", imports: ["os/exec"] })
extern class Exec {

    static var ErrDot: go.Error;
    static var ErrNotFound: go.Error;
    static var ErrWaitDelay: go.Error;

    static function command(name: String, arg: haxe.Rest<String>): go.Pointer<std.go.os.exec.Cmd>;
    static function lookPath(file: String): go.Result<String>;

}