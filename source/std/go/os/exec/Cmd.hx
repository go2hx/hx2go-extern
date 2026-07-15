package std.go.os.exec;

@:go.Type({ name: "Cmd", instanceName: "exec.Cmd", imports: ["os/exec"] })
extern class Cmd {
    function output(): go.Result<go.Slice<go.Byte>>;
}