package std.go.os.exec;

@:go.Type({ name: "Cmd", instanceName: "os/exec.Cmd", imports: ["os/exec"] })
extern class Cmd {
    public function output(): Void;
}