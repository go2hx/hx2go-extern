import sys.io.File;
import sys.FileSystem;
import haxe.io.Path;
import haxe.crypto.Md5;
import go.Syntax;
import go.Map;
using StringTools;


import go.go.types.Slice as SliceType;
import go.go.types.Array as ArrayType;
import go.go.types.Pointer as PointerType;
import go.go.types.Signature;
import go.go.types.Func;
import go.go.types.Var;
import go.go.types.Named;
import go.go.types.Interface;
import go.go.types.Alias;
import go.go.types.Struct;
import go.go.types.Chan as ChanType;

import go.golang_org.x.tools.go.Packages;
import go.golang_org.x.tools.go.packages.Package;
import go.golang_org.x.tools.go.packages.Config;
import go.golang_org.x.tools.go.packages.LoadMode;
import go.Pointer;
import go.Os;
import go.os.exec.Cmd;
import go.os.Exec;
import go.Go;
import go.haxe.HxArray;
import go.haxe.HxDynamic;
import go.Reflect;

using StringTools;

class Main {

    public static var topLevelName: String = "go";
    public static var scratchDir: String = null;
    

    static function ensureScratchModule(): String {
        if (scratchDir != null) return scratchDir;

        var dir = Os.tempDir() + "/hx2go-extern-scratch";
        Os.mkdirAll(dir, Syntax.code("0775"));

        var res = Os.stat(dir + "/go.mod");
        if (res.tuple().error != null) {
            var cmd = Exec.command("go", "mod", "init", "scratch");
            cmd.dir = dir;
            cmd.run();
        }

        scratchDir = dir;
        return dir;
    }

    static function ensureDependency(lib: String): Void {
        var dir = ensureScratchModule();
        if (lib.endsWith("/...")) {
            lib = lib.substr(0, lib.length - 4);
        }
        var cmd = Exec.command("go", "get", lib);
        cmd.dir = dir;

        var err = cmd.run();
        if (err != null) {
            Sys.println('failed to fetch $lib: ${err.error()}');
            Sys.exit(1);
        }
    }

    public static function toHaxeCaseWithUnderscore(input: String): String {
        if (input.charAt(0) == input.charAt(0).toLowerCase())
            return "_" + input;
        return toHaxeCase(input);
    }

    public static function toHaxeCase(input: String): String {
        return input == input.toUpperCase() ? input : input.charAt(0).toLowerCase() + input.substr(1);
    }

    public static function toPascalCase(input: String): String {
        var first = input.charAt(0);
        var isLetter = (first >= "a" && first <= "z") || (first >= "A" && first <= "Z");
        if (!isLetter) {
            return "T" + input;
        }
        return first.toUpperCase() + input.substr(1);
    }

    static var initLibs = [];
    static var stdGenBool = false;

    public static function main() {
        var args = Sys.args();
        if (args.length < 2) {
            Sys.println("Usage: hx2go-extern <lib>... <output>");
            Sys.exit(1);
        }
        for (arg in args) {
            switch arg {
                case "-stdgen", "--stdgen":
                    stdGenBool = true;
                    Sys.println("stdgen enabled");
                    args.remove(arg);
            }
        }
        // last arg is the output directory
        Sys.setCwd(args.pop());
        var output = args.pop();
        initLibs = args;
        genLibs(initLibs.copy(), output);
    }

    public static function willGenerate(lib: String): Bool {
        var parts = lib.split("/");
        if (parts.contains("internal")) {
            return false;
        }

        if (stdGenBool && (!parts.contains("term") && !parts.contains("tools") && parts.contains("golang.org"))) {
            return false;
        }
        return true;
    }

    
    // handle net/http.CookieJar and net/http/cookiejar
    static var scopeNames: Map<String, Array<String>> = new Map();

    static function parentPath(lib: String): String {
        var idx = lib.lastIndexOf("/");
        return idx == -1 ? "" : lib.substr(0, idx);
    }

    static function recordScopeNames(entry: Pointer<Package>): Void {
        var lib = entry.value.pkgPath;
        if (scopeNames.exists(lib) || entry.value.types == null) {
            return;
        }

        var names: Array<String> = [];
        var scope = entry.value.types.value.scope().value;
        for (name in scope.names()) {
            names.push(name);
        }

        scopeNames[lib] = names;
    }

    static function libGeneratesTypeName(lib: String, className: String): Bool {
        var names = scopeNames[lib];
        if (names == null) {
            return false;
        }
        var key = className.toLowerCase();
        for (name in names) {
            if (haxeTypeName(lib, name).toLowerCase() == key) {
                return true;
            }
        }
        return false;
    }

    // http.http2ClientConnPool vs http.http2clientConnPool)
    // handle case insensitive file systems
    public static function haxeTypeName(lib: String, name: String): String {
        var className = toPascalCase(name);
        var names = scopeNames[lib];
        if (names == null) {
            return className;
        }
        var key = className.toLowerCase();
        for (other in names) {
            if (other == name) {
                break;
            }
            if (toPascalCase(other).toLowerCase() == key) {
                className += "_";
            }
        }
        return className;
    }

    static function genLibs(libs: Array<String>, output: String): Void {
        var loadDir = ensureScratchModule();

        var cachePath = Path.join([output, ".hx2go_extern_cache"]);
        var topLevelOut = Path.join([output, topLevelName]);
        var preKey = topLevelCacheKey(libs, loadDir);
        if (preKey != null && FileSystem.exists(topLevelOut)) {
            var prev = FileSystem.exists(cachePath) ? File.getContent(cachePath) : "";
            if (prev == preKey) {
                return;
            }
        }

        for (lib in libs) {
            if (lib.split("/")[0].contains(".")) {
                ensureDependency(lib);
            }
        }
        var cacheKey = topLevelCacheKey(libs, loadDir);

        var config: Config = {
            mode: Syntax.code('packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedFiles | packages.NeedDeps'),
            dir: loadDir,
            tests: false,
            parseFile: null,
            overlay: null,
            logf: null,
            fset: null,
            env: null,
            context: null,
            buildFlags: null,
        };
        var configPtr: Pointer<Config> = config;

        var entries = Packages.load(configPtr, ...libs).sure();

        var all: Map<String, Pointer<Package>> = new Map();
        for (entry in entries) {
            collectPkgs(entry, all);
        }

        for (lib in all.keys()) {
            recordScopeNames(all[lib]);
        }

        for (lib in all.keys()) {
            if (willGenerate(lib)) {
                genPackage(all[lib], output);
            }
        }

        if (cacheKey != null) {
            Os.mkdirAll(output, Syntax.code("0775")).sure();
            File.saveContent(cachePath, cacheKey);
        }
    }
    
    static function topLevelCacheKey(libs: Array<String>, loadDir: String): String {
        var goMod = Path.join([loadDir, "go.mod"]);
        var goSum = Path.join([loadDir, "go.sum"]);
        if (!FileSystem.exists(goMod)) {
            return null;
        }
        var sorted = libs.copy();
        sorted.sort((a, b) -> a < b ? -1 : (a > b ? 1 : 0));
        var parts = [sorted.join("\n")];
        parts.push(Md5.encode(File.getContent(goMod)));
        parts.push(FileSystem.exists(goSum) ? Md5.encode(File.getContent(goSum)) : "");
        return Md5.encode(parts.join("$|"));
    }

    static function collectPkgs(entry: Pointer<Package>, all: Map<String, Pointer<Package>>): Void {
        var lib = entry.value.pkgPath;
        if (all.exists(lib) || entry.value.types == null) {
            return;
        }

        all[lib] = entry;

        if (entry.value.imports == null) {
            return;
        }

        for (dep in entry.value.imports.keys()) {
            collectPkgs(entry.value.imports[dep], all);
        }
    }

    static function genPackage(entry: Pointer<Package>, output: String): Void {
        var lib = entry.value.pkgPath;

        // skip regenerating an already existing package
        var checkSumPath = Path.join([getPackageDir(output, topLevelName, lib), ".hx2go_cache"]);
        var prevCheckSum = FileSystem.exists(checkSumPath) ? File.getContent(checkSumPath) : "";
        var checkSum = Cache.getPackageCheckSum(entry);
        if (prevCheckSum == checkSum) {
            return;
        } else {
            Os.mkdirAll(Path.directory(checkSumPath), Syntax.code("0775")).sure();
            File.saveContent(checkSumPath, checkSum);
        }
        Sys.println('generating "$lib"');

        var outputs: Map<String, GenOutput> = new Map();

        function getOutput(name: String): GenOutput {
            if (name == lib.split("/").pop()) {
                name = "@package";
            }

            if (!outputs.exists(name)) {
                outputs[name] = {
                    staticFunctions: new StringBuf(),
                    instanceFunctions: new StringBuf(),
                    staticVars: new StringBuf(),
                    instanceVars: new StringBuf(),
                    consts: new StringBuf(),
                    paramStr: '',
                    isInterface: false,
                    isStruct: false,
                    ctorParams: [],
                    ctorValues: [],
                    typedefStr: null
                };
            }

            return outputs[name];
        }

        var scope = entry.value.types.value.scope().value;

        for (name in scope.names()) {
            var obj = scope.lookup(name);

            Syntax.code("switch {0}.(type) {", obj); // this is so bad :[
            Syntax.code("case *types.TypeName:");
                GenTypeName.gen(obj, getOutput);
            Syntax.code("case *types.Func:"); {
                if (!obj.exported()) {
                    continue;
                }
                var func = TypeHelper.typeAs(obj, Func);
                var sig = func.signature().value;
                var recv = sig.recv();
                var buf = getOutput(entry.value.name).staticFunctions;

                // var typeParams = sig.typeParams().value;
                var params = sig.params()?.value ?? null;
                var results = sig.results()?.value ?? null;
                var varadic = sig.variadic();

                buf.add('    ' + genFunc(name, sig, recv == null) + '\n');
            }

            Syntax.code("case *types.Var:"); {
                if (!obj.exported()) {
                    continue;
                }
                var v = TypeHelper.typeAs(obj, Var);
                if (!TypeHelper.isExportedType(v.type())) {
                    continue;
                }

                var buf = getOutput(entry.value.name).staticVars;
                var name = v.name();
                var type = v.type();

                buf.add('    @:native("${name}") static var ${Sanitize.name(toHaxeCaseWithUnderscore(name))}: ${genType(type)};\n');
            }

            Syntax.code("case *types.Const:"); {
                var c = TypeHelper.typeAs(obj, go.go.types.Const);

                if (!c.exported()) {
                    continue;
                }

                var buf = getOutput(entry.value.name).consts;

                var name = c.name();
                var type = c.type();

                buf.add('    @:native("${name}") static var ${Sanitize.name(toHaxeCaseWithUnderscore(name))}: ${genType(type)};\n');
            }

            Syntax.code("default:"); {
                // trace(Reflect.typeOf(obj).string());
            }

            Syntax.code("}");
            // trace(obj.name(), obj.type().string());
        }

        for (file in outputs.keys()) {
            var buf = new StringBuf();
            var out = outputs[file];
            var isPkg = false;
            var relLib = lib;

            if (file == "@package") {
                var parts = lib.split("/");
                file = (parts.pop() : String);
                relLib = parts.join("/");
                isPkg = true;
            }

            // if os.File and os.file add "_" suffix to unexported type variant
            var className = isPkg ? toPascalCase(file) : haxeTypeName(lib, file);
            if (className != file) {
                // the parent's types own the directory this package class lands in
                if (isPkg && libGeneratesTypeName(relLib, className)) {
                    className += "_";
                }
                // unexported empty, skip
                // if (out.instanceFunctions.toString() == "" && out.instanceVars.toString() == "") {
                //     continue;
                // }
            }
            var pkgName = topLevelName + (relLib.length > 0 ? "." + Sanitize.packagePath(relLib) : "");

            // recursive typedef, avoid collision
            if (out.typedefStr == className) {
                out.typedefStr = switch className {
                    case "String": "std.String";
                    case "Bool": "StdTypes.Bool";
                    case "Float": "StdTypes.Float";
                    case _: out.typedefStr;
                }
            }

            buf.add('package $pkgName;\n');
            buf.add('\n');
            if (out.isStruct == true) {
                buf.add('@:structInit\n'); // TODO: generate constructor where everything is optional
            }
            var packageName = entry.value.name;
            buf.add('@:go.Type({ name: "${file}", instanceName: "${packageName}.${file}", imports: ["${lib}"] })\n');
            buf.add('extern ${out.isInterface || out.typedefStr != null ? "typedef" : "class"} ${className}${out.paramStr}${out.isInterface || out.typedefStr != null ? " = " : " "}');

            if (out.typedefStr != null) {
                if (out.instanceFunctions.toString() != "") {
                    var methodsTypeName = className + "Methods" + out.paramStr;

                    buf.add('haxe.extern.EitherType<${out.typedefStr}, {\n');
                    buf.add(out.instanceFunctions.toString());
                    if (out.staticFunctions.length > 0 || out.instanceFunctions.length > 0) buf.add("\n");
                    buf.add("}>");
                }else{
                    buf.add(out.typedefStr);
                }
            } else {
                buf.add('{\n\n');
                buf.add(out.consts.toString());
                if (out.consts.length > 0) buf.add('\n');
                buf.add(out.staticVars.toString());
                buf.add(out.instanceVars.toString());
                if (out.staticVars.length > 0 || out.instanceVars.length > 0) buf.add('\n');
                if (out.ctorParams.length > 0) {
                    var ctorParams = out.ctorParams.copy();
                    for (i in 0...ctorParams.length) {
                        var value = out.ctorValues[i];
                        if (value == "") {
                            continue;
                        }
                        ctorParams[i] += "=" + value;
                    }
                    buf.add('    function new(${ctorParams.join(", ")});\n\n');
                }

                buf.add(out.staticFunctions.toString());
                buf.add(out.instanceFunctions.toString());
                if (out.staticFunctions.length > 0 || out.instanceFunctions.length > 0) buf.add("\n");
                buf.add("}");
            }

            var dir = '${output}/${topLevelName}/${Sanitize.packageDir(relLib)}';

            Os.mkdirAll(dir, Syntax.code("0775")).sure();
            Os.writeFile('$dir/${className}.hx', cast buf.toString(), Syntax.code("0666")).sure();
        }
    }

    static function getPackageDir(output:String, topLevelName:String, relLib:String) {
        return '${output}/${topLevelName}/${Sanitize.packageDir(relLib)}';
    }

    public static function genFunc(name: String, sig: Signature, topLevel:Bool, closure: Bool = false) {
        var recv = sig.recv();
        var params = sig.params()?.value ?? null;
        var results = sig.results()?.value ?? null;
        var varadic = sig.variadic();

        var params = genTuple(params, varadic);
        var meta = "";

        if (results.len() > 1 && !isResultType(results) && !closure) {
            var names = [];
            var unnamed = 0;

            for (idx in 0...results.len()) {
                var name = results.at(idx).value.name();
                if (name == "") {
                    name = 'p${unnamed++}';
                }

                names.push('"$name"');
            }

            meta = '@:go.Tuple(${names.join(", ")}) ';
        }

        if (!closure) {
            meta += '@:native("${name}") ';
        }

        var tParams = sig.typeParams() != null ? sig.typeParams().value : null;
        var tParamsStr = '';
        if (tParams != null){
            var tParamsLocal = [];
            for (i in 0...tParams.len()) {
                var t = tParams.at(i).value;
                var constraint = t.constraint();
                var constraintStr = genType(constraint);

                tParamsLocal.push('${toPascalCase(t.string())}: ${constraintStr}');
            }
            tParamsStr = '<' + tParamsLocal.join(", ") + '>';
        }
        
        return '${meta}${topLevel && !closure ? "static " : ""}${closure ? "" : 'function ${Sanitize.name(toHaxeCaseWithUnderscore(name))}'}${tParamsStr}(${params.join(", ")})${closure ? ' -> ' : ': '}${results.len() == 0 ? "Void" :('(' + genResults(results) + ')')}${closure ? "" : ";"}';
    }

    static function isResultType(args:go.go.types.Tuple) {
        return args.len() == 2 && args.at(1).value.type().string() == "error";
    }

    static function genResults(args:go.go.types.Tuple) {
        if (args.len() > 1) {
            if (isResultType(args)) {
                return 'go.Result<${genType(args.at(0).value.type())}>';
            }

            return 'go.Tuple<{ ' + genTuple(args, false, false).join(", ") + ' }>';
        }else{
            return genTuple(args, false, true).join(", ");
        }
    }

    static function genTuple(args:go.go.types.Tuple, varadic: Bool = false, ret: Bool = false) {
        var items = [];
        var idx = 0;

        for (i in 0...args.len()) {
            var isLastArg = i == args.len() - 1;

            var v = args.at(i).value;
            var n = v.name();
            if (n == "") {
                n = 'p${idx++}';
            }

            items.push(
                if (ret) genType(v.type());
                else if (varadic && isLastArg) n + ": haxe.Rest<" + genType(TypeHelper.typeAs(v.type(), SliceType).elem()) + ">";
                else Sanitize.name(n) + ": " + genType(v.type())
            );
        }

        return items;
    }

    static function resolvePath(path: String): String {
        var lastSlash = path.lastIndexOf("/");
        var pkgPath = lastSlash == -1 ? "" : path.substr(0, lastSlash);
        var rest = lastSlash == -1 ? path : path.substr(lastSlash + 1);
        var dotIdx = rest.indexOf(".");
        var pkgBase = rest.substr(0, dotIdx);
        var typeName = rest.substr(dotIdx + 1);
        var fullPkgPath = pkgPath.length > 0 ? '${pkgPath}/${pkgBase}' : pkgBase;

        return '${topLevelName}.${Sanitize.packagePath(fullPkgPath)}.${haxeTypeName(fullPkgPath, typeName)}';
    }

    public static function genType(t: go.go.types.Type): String {
        var s = t.string();
        var tParamStr = '';

        if (TypeHelper.isNamedType(t)) {
            var namedObj = TypeHelper.typeAs(t, Named).obj();
            if (namedObj != null) {
                var pkg = namedObj.value.pkg();
                if (pkg != null && !willGenerate(pkg.value.path())) {
                    return "Dynamic";
                }
            }
        }

        if (TypeHelper.isNamedType(t) && TypeHelper.typeAs(t, Named).typeArgs() != null) {
            var typeParams = TypeHelper.typeAs(t, Named).typeArgs().value;
            var typeParamStrs = [];
            for (i in 0...typeParams.len()) {
                typeParamStrs.push(genType(typeParams.at(i)));
            }

            var typeParamStart = s.indexOf('[');
            var typeParamEnd = s.lastIndexOf(']');
            if (typeParamStart != -1 && typeParamEnd != -1 && typeParamEnd > typeParamStart) {
                s = s.substr(0, typeParamStart) + s.substr(typeParamEnd + 1);
            }

            tParamStr = '<' + typeParamStrs.join(", ") + '>';
        }

        if (s.startsWith("untyped ")) {
            s = s.substr(8);
        }

        if (s.startsWith("invalid type")) {
            return "Dynamic";
        }

        if (TypeHelper.isInterfaceType(t) && !TypeHelper.isNamedType(t)) {
            var iface = TypeHelper.typeAs(t, Interface);
            if (iface.numMethods() == 0) {
                return "Dynamic" + tParamStr;
            }

            return "Dynamic" + tParamStr;
        }

        var q = switch s {
            case "error": "go.Error";
            case "string": "String";
            case "bool": "Bool";
            case "any": "Dynamic";

            case "int": "go.GoInt";
            case "int64": "go.Int64";
            case "int32": "go.Int32";
            case "int16": "go.Int16";
            case "int8": "go.Int8";

            case "uint": "go.GoUInt";
            case "uint64": "go.UInt64";
            case "uint32": "go.UInt32";
            case "uint16": "go.UInt16";
            case "uint8": "go.UInt8";

            case "float": "Float"; // can happen in the case of untyped types apparently
            case "float64": "Float";
            case "float32": "go.Float32";
            case "float16": "go.Float16";

            case "byte": "go.Byte";
            case "rune": "go.Rune";

            case "uintptr": "go.UIntPtr";
            case "complex128": "go.Complex128";
            case "complex64": "go.Complex64";
            case "comparable": "go.Comparable";

            case _ if (s.substr(0, 7) == "chan<- "): 'go.Chan<${genType(TypeHelper.typeAs(t, ChanType).elem())}>';
            case _ if (s.substr(0, 5) == "chan "): 'go.Chan<${genType(TypeHelper.typeAs(t, ChanType).elem())}>';
            case _ if (s.substr(0, 7) == "<-chan "): 'go.Chan<${genType(TypeHelper.typeAs(t, ChanType).elem())}>';
            case _ if (s.substr(0, 2) == "[]"): 'go.Slice<${genType(TypeHelper.typeAs(t, SliceType).elem())}>';
            case _ if (s.substr(0, 1) == "*"): 'go.Pointer<${genType(TypeHelper.typeAs(t, PointerType).elem())}>';

            case _ if (s.replace(" ", "").startsWith("struct{")): {
                // TODO: generate struct type
                'Dynamic';
            }

            case _ if (s.startsWith("map[")): {
                var keyType = TypeHelper.typeAs(t, go.go.types.Map).key();
                var elemType = TypeHelper.typeAs(t, go.go.types.Map).elem();
                'go.Map<${genType(keyType)}, ${genType(elemType)}>';
            }

            case _ if (s.startsWith('func')): {
                var sig = TypeHelper.typeAs(t, Signature);
                genFunc('', sig, false, true);
            }

            case _ if (TypeHelper.isArrayType(t)): {
                var arr = TypeHelper.typeAs(t, ArrayType);
                'go.GoArray<${genType(arr.elem())}, ${arr.len()}>';
            }

            case _ if (TypeHelper.isTypeParamType(t)): toPascalCase(s);
            case _ if (s.contains(".")): resolvePath(s);
            case _: s;
        }

        if (q.charAt(0) == "~") {
            q = "Dynamic"; // Todo: ~T[] -> go.Slice<Dynamic> and not go.Slice<T>
        }

        return q + tParamStr;
    }
}

interface ElemType {
    public function elem():go.go.types.Type;
}
