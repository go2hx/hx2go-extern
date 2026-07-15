package go.os.exec;

@:go.Type({ name: "exec", instanceName: "exec.exec", imports: ["os/exec"] })
extern class Exec {

    static function command(name: String, arg: ...String): go.Pointer<go.os.exec.Cmd>;
    static function commandContext(ctx: go.context.Context, name: String, arg: ...String): go.Pointer<go.os.exec.Cmd>;
    static function lookPath(file: String): go.Result<String>;

}