import haxe.io.Path;
import sys.FileSystem;

function main() {
    var go2hxPath = Sys.getCwd();
    var go2hxFilePath = Path.join([ go2hxPath, 'output', 'main' ]);
    var bin = Path.join([ go2hxPath, 'output', 'main', executable('main')]);

    if (!FileSystem.exists(go2hxFilePath)) {
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