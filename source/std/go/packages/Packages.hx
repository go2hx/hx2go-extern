package std.go.packages;

import go.Slice;
import go.Pointer;
import go.Map;
import go.Result;
import go.Int64;
import std.go.types.Types.Package as TypesPackage;
import std.go.types.Types.Info as TypesInfo;
import std.go.types.Types.Sizes as TypesSizes;
import std.go.types.Types.Package as TypesPackage;
import std.go.token.FileSet;
import std.go.ast.Ast.File;

@:go.Type({ name: "packages.Package", imports: ["golang.org/x/tools/go/packages"] })
extern class Package {
	var id: String;
	var name: String;
	var pkgPath: String;
	var errors: Slice<PackageError>;
	var goFiles: Slice<String>;
	var compiledGoFiles: Slice<String>;
	var otherFiles: Slice<String>;
	var embedFiles: Slice<String>;
	var embedPatterns: Slice<String>;
	var ignoredFiles: Slice<String>;
	var exportFile: String;
	var imports: Map<String, Pointer<Package>>;
	var types: Pointer<TypesPackage>;
	var fset: Pointer<FileSet>;
	var illTyped: Bool;
	var syntax: Slice<Pointer<File>>;
	var typesInfo: Pointer<TypesInfo>;
	var typesSizes: TypesSizes;
	var module: Pointer<Module>;

	function string(): String;
}

@:go.Type({ name: "packages.PackageError", imports: ["golang.org/x/tools/go/packages"] })
extern class PackageError {
	var pos: String;
	var msg: String;
	var kind: Int;

	function error(): String;
}

@:go.Type({ name: "packages.Module", imports: ["golang.org/x/tools/go/packages"] })
extern class Module {
	var path: String;
	var version: String;
	var main: Bool;
}

@:go.Type({ name: "packages", instanceName: "packages.LoadMode", imports: ["golang.org/x/tools/go/packages"] })
extern class LoadMode {
	static var needName: LoadMode;
	static var needFiles: LoadMode;
	static var needCompiledGoFiles: LoadMode;
	static var needImports: LoadMode;
	static var needDeps: LoadMode;
	static var needExportFile: LoadMode;
	static var needTypes: LoadMode;
	static var needSyntax: LoadMode;
	static var needTypesInfo: LoadMode;
	static var needTypesSizes: LoadMode;
	static var needForTest: LoadMode;
	static var needModule: LoadMode;
	static var needEmbedFiles: LoadMode;
	static var needEmbedPatterns: LoadMode;
	static var needTarget: LoadMode;

	public inline extern function or(other: LoadMode): LoadMode {
		return go.Syntax.code("{0} | {1}", this, other);
	}
}

@:structInit
@:go.Type({ name: "packages.Config", imports: ["golang.org/x/tools/go/packages"] })
extern class Config {
	public var mode: LoadMode;
	public var dir: String;
}

@:go.Type({ name: "packages", imports: ["golang.org/x/tools/go/packages"] })
extern class Packages {
	static function load(cfg: Pointer<Config>, ...patterns: String): Result<Slice<Pointer<Package>>>;
}