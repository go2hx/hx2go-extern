import haxe.io.Path;
import sys.FileSystem;

function main() {
    var go2hxPath = Sys.getCwd();
    var go2hxFilePath = Path.join([ go2hxPath, 'binary', 'go2hx' ]);

    if (!FileSystem.exists(go2hxFilePath)) {
        Sys.println('go2hx has not been built yet, doing so now');
        Sys.command("haxe Compile.hxml");
    }

    Sys.command("go", ["-C", "./binary/go2hx/main", "run", "."].concat(Sys.args()));
}