import haxe.io.Path;
import sys.FileSystem;

function main() {
    var go2hxPath = Sys.getCwd();
    var go2hxFilePath = Path.join([ go2hxPath, 'output', 'main' ]);
    var bin = Path.join([ go2hxPath, 'output', 'main', executable('main')]);

    var args = Sys.args();

    if (!FileSystem.exists(go2hxFilePath) || (args.length > 0 && args[0] == "rebuild")) {
        Sys.println('go2hx has not been built yet, doing so now');
        Sys.command("haxe Compile.hxml");
    }

    Sys.command(bin, [].concat(Sys.args()));
}

private function executable(path: String): String {
	return if (Sys.systemName().toLowerCase() == "windows") {
		path + '.exe';
	}else{
		path;
	}
}