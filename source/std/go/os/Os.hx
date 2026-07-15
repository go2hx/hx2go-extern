package std.go.os;

@:go.Type({ name: "os", instanceName: "os.os", imports: ["os"] })
extern class Os {

    static function writeFile(name: String, data: go.Slice<go.Byte>, perm: std.go.os.FileMode): go.Error;
    static function mkdirAll(path: String, perm: std.go.os.FileMode): go.Error;
}