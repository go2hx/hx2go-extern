package std.go.os.exec;

@:go.Type({ name: "exec", instanceName: "os/exec.exec", imports: ["os/exec"] })
extern class Exec {

    static function command(name: String, arg: ...String): std.go.os.exec.Cmd;
    static function lookPath(file: String): go.Result<String>;

}