import haxe.io.Path;
import sys.FileSystem;

function main() {
    var cwd = Sys.getCwd();
    var filePath = Path.join([ cwd, 'output', 'main' ]);
    var bin = Path.join([ filePath, executable('main')]);

    var args = Sys.args();

    if (!FileSystem.exists(bin) || (args.length > 0 && args[0] == "rebuild")) {
        args.shift();
        Sys.println('hx2go-extern building...');
        Sys.command("haxe Compile.hxml");
    }

    Sys.command(bin, [].concat(args));
}

private function executable(path: String): String {
	return if (Sys.systemName().toLowerCase() == "windows") {
		path + '.exe';
	}else{
		path;
	}
}