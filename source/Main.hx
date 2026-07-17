import std.go.types.Types.TypeName;
import go.Syntax;
import go.Map;
import std.go.types.Types.Signature;
import std.go.types.Types.Func;
using StringTools;

import std.go.types.Types.Chan;
import std.go.types.Types.Slice;
import std.go.types.Types.Pointer as PointerType;
import std.go.packages.Packages;
import go.Pointer;
import go.Os;
import std.go.types.Types.Var;
import std.go.types.Types.Named;
import go.os.exec.Cmd;
import go.os.Exec;
import go.Go;
import std.go.types.Types.Interface;
import std.go.types.Types.Alias;
import std.go.types.Types.Struct;
import go.haxe.HxArray;
import go.haxe.HxDynamic;
import std.go.types.Types.Array as ArrayType;

using StringTools;

class Main {

    public static var didGen: Map<String, Bool> = new Map();
    static var mutex = new sys.thread.Mutex();
    public static var topLevelName: String = "go";
    public static var scratchDir: String = null;

    static function ensureScratchModule(): String {
        if (scratchDir != null) return scratchDir;

        var dir = Os.tempDir() + "/go2hx-scratch";
        Os.mkdirAll(dir, Syntax.code("0775"));

        var res = Os.stat(dir + "/go.mod");
        if (res.tuple().error != null) {
            var cmd = Exec.command("go", "mod", "init", "go2hxscratch");
            cmd.dir = dir;
            cmd.run();
        }

        scratchDir = dir;
        return dir;
    }

    static function ensureDependency(lib: String): Void {
        var dir = ensureScratchModule();
        var cmd = Exec.command("go", "get", lib);
        cmd.dir = dir;

        var err = cmd.run();
        if (err != null) {
            Sys.println('failed to fetch $lib: ${err.error()}');
            Sys.exit(1);
        }
    }

    public static function toHaxeCase(input: String): String {
        return input == input.toUpperCase() ? input : input.charAt(0).toLowerCase() + input.substr(1);
    }

    public static function toPascalCase(input: String): String {
        return input.charAt(0).toUpperCase() + input.substr(1);
    }

    public static function main() {
        var args = Sys.args();
        if (args.length < 2) {
            Sys.println("Usage: go2hx <lib> <output>");
            Sys.exit(1);
        }

        var output = args[1];
        var lib = args[0];
        genLib(lib, output);
    }

    static function isExportedType(t:std.go.types.Types.Type):Bool {
        if (isNamedType(t)) {
            var named = typeAs(t, Named);
            return named.obj().value.exported();
        }

        var s = t.string();

        if (s.startsWith("*")) {
            return isExportedType(typeAs(t, PointerType).elem());
        }

        if (s.startsWith("[]")) {
            return isExportedType(typeAs(t, Slice).elem());
        }

        if (s.startsWith("chan ")) {
            return isExportedType(typeAs(t, Chan).elem());
        }

        return true;
    }

    static function isInterfaceType(t:std.go.types.Types.Type):Bool {
        var result = false;

        Syntax.code("
            if _, ok := t.(*types.Interface); ok {
                result = true;
            }
        ");

        return result;
    }

    static function getUnderlying(t:std.go.types.Types.Type):std.go.types.Types.Type {
        if (isNamedType(t)) {
            return typeAs(t, Named).underlying();
        }

        return t;
    }

    static function isBasicType(t:std.go.types.Types.Type):Bool {
        var result = false;

        Syntax.code("
            if _, ok := t.(*types.Basic); ok {
                result = true;
            }
        ");

        return result;
    }

    static function isStructType(t:std.go.types.Types.Type):Bool {
        var result = false;

        Syntax.code("
        if _, ok := t.(*types.Struct); ok {
            result = true;
        }
    ");

        return result;
    }

    static function isArrayType(t:std.go.types.Types.Type):Bool {
        var result = false;

        Syntax.code("
        if _, ok := t.(*types.Array); ok {
            result = true;
        }
    ");

        return result;
    }

    static function genLib(lib: String, output: String): Void {
        mutex.acquire();
        if (didGen.exists(lib) || lib.split("/").contains("internal")) {
            mutex.release();
            return;
        }

        didGen.set(lib, true);
        mutex.release();

        Sys.println('generating "$lib"');

        var loadDir = "";
        if (lib.split("/")[0].contains(".")) {
            ensureDependency(lib);
            loadDir = ensureScratchModule();
        }

        var config: Config = {
            mode: LoadMode.needName.or(LoadMode.needTypes).or(LoadMode.needTypesInfo).or(LoadMode.needSyntax).or(LoadMode.needImports),
            dir: loadDir
        };

        var entries = Packages.load(config, lib).sure();
        var outputs: Map<String, { ctorParams: Array<String>, isStruct: Bool, typedefStr: Null<String>, consts: StringBuf, isInterface: Bool, paramStr: String, staticFunctions: StringBuf, instanceFunctions: StringBuf, staticVars: StringBuf, instanceVars: StringBuf }> = new Map();

        function getOutput(name: String) {
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
                    typedefStr: null
                };
            }

            return outputs[name];
        }

        for (entry in entries) {
            var scope = entry.value.types.value.scope().value;
            for (dep in entry.value.imports.keys()) {
                Syntax.go(() -> genLib(dep, output));
            }

            for (name in scope.names()) {
                var obj = scope.lookup(name);
//                if (!obj.exported()) {
//                    continue;
//                }

                Syntax.code("switch {0}.(type) {", obj); // this is so bad :[
                Syntax.code("case *types.TypeName:"); {
                    var type = typeAs(obj, TypeName);
                    var out = getOutput(obj.name());
                    var underlying = getUnderlying(type.type());

                    var resolvedTo = std.go.types.Types.Types.unalias(type.type());
                    if (resolvedTo == type.type() && isNamedType(resolvedTo)) {
                        var named = typeAs(resolvedTo, Named);
                        resolvedTo = named.underlying();
                    }

                    if (isInterfaceType(underlying)) {
                        var iface = typeAs(underlying, Interface);
                        out.isInterface = true;

                        iface.complete();

                        for (i in 0...iface.numMethods()) {
                            var method = typeAs(iface.method(i), Func);
                            if (!method.exported()) {
                                continue;
                            }

                            var sig = method.signature().value;
                            out.instanceFunctions.add('    ${genFunc(method.name(), sig, false, false)}\n');
                        }
                    } else if (type.isAlias() && (isNamedType(resolvedTo) || isBasicType(resolvedTo))) { // TODO: support for func() aswell (see iter.Seq)
                        out.typedefStr = genType(resolvedTo);
                    } else {
                        var isNamed = isNamedType(type.type());
                        var out = getOutput(obj.name());

                        if (isNamed) {
                            var named = typeAs(type.type(), Named);

                            if (isStructType(underlying)) {
                                out.isStruct = true;
                                var struct = typeAs(underlying, Struct);

                                for (i in 0...struct.numFields()) {
                                    var field = typeAs(struct.field(i), Var);
                                    if (!field.exported()) {
                                        continue;
                                    }

                                    var fname = field.name();
                                    var p = '${Sanitize.name(toHaxeCase(fname))}: ${genType(field.type())}';
                                    out.instanceVars.add('    @:native("${fname}") var ${p};\n');
                                    out.ctorParams.push(p);
                                }
                            }

                            var tp = named.typeParams();
                            if (tp != null) {
                                var params = [];
                                var tps = tp.value;

                                for (i in 0...tps.len()) {
                                    var t = tps.at(i).value;
                                    var constraint = t.constraint();
                                    var constraintStr = genType(constraint);
                                    params.push('${t.string()}: ${constraintStr}');
                                }

                                out.paramStr = params.length == 0 ? "" : "<" + params.join(", ") + ">";
                            }

                            var methodSet = std.go.types.Types.newMethodSet(Syntax.code("types.NewPointer({0})", type.type()));
                            for (i in 0...methodSet.value.len()) {
                                var sel = methodSet.value.at(i).value;
                                var method = typeAs(sel.obj(), Func);

                                if (!method.exported()) {
                                    continue;
                                }

                                var sig = method.signature().value;
                                out.instanceFunctions.add('    ${genFunc(method.name(), sig, false)}\n');
                            }
                        }
                    }
                }

                Syntax.code("case *types.Func:"); {
                    var func = typeAs(obj, Func);
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
                    var v = typeAs(obj, Var);
                    if (!isExportedType(v.type())) {
                        continue;
                    }

                    var buf = getOutput(entry.value.name).staticVars;
                    var name = v.name();
                    var type = v.type();

                    buf.add('    @:native("${name}") static var ${Sanitize.name(name)}: ${genType(type)};\n');
                }

                Syntax.code("case *types.Const:"); {
                    var c = typeAs(obj, std.go.types.Types.Const);

                    if (!c.exported()) {
                        continue;
                    }

                    var buf = getOutput(entry.value.name).consts;

                    var name = c.name();
                    var type = c.type();

                    buf.add('    @:native("${name}") static var ${Sanitize.name(name)}: ${genType(type)};\n');
                }

                Syntax.code("default:"); {
                    // trace(go.reflect.Reflect.typeOf(obj).string());
                }

                Syntax.code("}");
                // trace(obj.name(), obj.type().string());
            }

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

            buf.add('package ${topLevelName}${relLib.length > 0 ? "." + Sanitize.packagePath(relLib) : ""};\n');
            buf.add('\n');
            if (out.isStruct == true) {
                buf.add('@:structInit\n'); // TODO: generate constructor where everything is optional
            }
            buf.add('@:go.Type({ name: "${file}", instanceName: "${lib.split("/").pop()}.${file}", imports: ["${lib}"] })\n');
            buf.add('extern ${out.isInterface || out.typedefStr != null ? "typedef" : "class"} ${toPascalCase(file)}${out.paramStr}${out.isInterface || out.typedefStr != null ? " = " : " "}');

            if (out.typedefStr != null) {
                buf.add(out.typedefStr);
            } else {
                buf.add('{\n\n');
                buf.add(out.consts.toString());
                if (out.consts.length > 0) buf.add('\n');
                buf.add(out.staticVars.toString());
                buf.add(out.instanceVars.toString());
                if (out.staticVars.length > 0 || out.instanceVars.length > 0) buf.add('\n');
                if (out.ctorParams.length > 0) {
                    buf.add('    function new(${out.ctorParams.join(", ")});\n\n');
                }

                buf.add(out.staticFunctions.toString());
                buf.add(out.instanceFunctions.toString());
                if (out.staticFunctions.length > 0 || out.instanceFunctions.length > 0) buf.add("\n");
                buf.add("}");
            }

            Os.mkdirAll('${output}/${topLevelName}/${Sanitize.packageDir(relLib)}', Syntax.code("0775"));
            Os.writeFile('${output}/${topLevelName}/${Sanitize.packageDir(relLib)}/${toPascalCase(file)}.hx', cast buf.toString(), Syntax.code("0666"));
        }
    }

    static function genPackage(pkg) {
        return '';
    }

    static function genFile(file) {
        return '';
    }

    static function genFunc(name: String, sig: Signature, topLevel:Bool, closure: Bool = false) {
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

                tParamsLocal.push('${t.string()}: ${constraintStr}');
            }
            tParamsStr = '<' + tParamsLocal.join(", ") + '>';
        }

        return '${meta}${topLevel && !closure ? "static " : ""}${closure ? "" : 'function ${Sanitize.name(toHaxeCase(name))}'}${tParamsStr}(${params.join(", ")})${closure ? ' -> ' : ': '}${results.len() == 0 ? "Void" :genResults(results)}${closure ? "" : ";"}';
    }

    static function isResultType(args:std.go.types.Types.Tuple) {
        return args.len() == 2 && args.at(1).value.type().string() == "error";
    }

    static function genResults(args:std.go.types.Types.Tuple) {
        if (args.len() > 1) {
            if (isResultType(args)) {
                return 'go.Result<${genType(args.at(0).value.type())}>';
            }

            return 'go.Tuple<{ ' + genTuple(args, false, false).join(", ") + ' }>';
        }else{
            return genTuple(args, false, true).join(", ");
        }
    }

    static function genTuple(args:std.go.types.Types.Tuple, varadic: Bool = false, ret: Bool = false) {
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
                else if (varadic && isLastArg) n + ": haxe.Rest<" + genType(typeAs(v.type(), Slice).elem()) + ">";
                else Sanitize.name(n) + ": " + genType(v.type())
            );
        }

        return items;
    }

    inline extern static function typeAs<V, T>(value: V, as: Class<T>): T {
        var v: Pointer<T> = cast (value : Dynamic);
        return v.value;
    }

    static function resolvePath(path: String): String {
        var lastSlash = path.lastIndexOf("/");
        var pkgPath = lastSlash == -1 ? "" : path.substr(0, lastSlash);
        var rest = lastSlash == -1 ? path : path.substr(lastSlash + 1);
        var dotIdx = rest.indexOf(".");
        var pkgBase = rest.substr(0, dotIdx);
        var typeName = rest.substr(dotIdx + 1);
        var fullPkgPath = pkgPath.length > 0 ? '${pkgPath}/${pkgBase}' : pkgBase;

        return '${topLevelName}.${Sanitize.packagePath(fullPkgPath)}.${toPascalCase(typeName)}';
    }

    static function isNamedType(t:std.go.types.Types.Type): Bool {
        var isNamed = false;
        Syntax.code("if _, ntOk := t.(*types.Named); ntOk { isNamed = true; }");

        return isNamed;
    }

    static function genType(t: std.go.types.Types.Type): String {
        var s = t.string();
        var tParamStr = '';

        if (isNamedType(t) && typeAs(t, Named).typeArgs() != null) {
            var typeParams = typeAs(t, Named).typeArgs().value;
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

        if (isInterfaceType(t) && !isNamedType(t)) {
            var iface = typeAs(t, Interface);
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

            case _ if (s.substr(0, 7) == "chan<- "): 'go.Chan<${genType(typeAs(t, Chan).elem())}>';
            case _ if (s.substr(0, 5) == "chan "): 'go.Chan<${genType(typeAs(t, Chan).elem())}>';
            case _ if (s.substr(0, 7) == "<-chan "): 'go.Chan<${genType(typeAs(t, Chan).elem())}>';
            case _ if (s.substr(0, 2) == "[]"): 'go.Slice<${genType(typeAs(t, Slice).elem())}>';
            case _ if (s.substr(0, 1) == "*"): 'go.Pointer<${genType(typeAs(t, PointerType).elem())}>';

            case _ if (s.replace(" ", "").startsWith("struct{")): {
                // TODO: generate struct type
                'Dynamic';
            }

            case _ if (s.startsWith("map[")): {
                var keyType = typeAs(t, std.go.types.Types.Map).key();
                var elemType = typeAs(t, std.go.types.Types.Map).elem();
                'go.Map<${genType(keyType)}, ${genType(elemType)}>';
            }

            case _ if (s.startsWith('func')): {
                var sig = typeAs(t, Signature);
                genFunc('', sig, false, true);
            }

            case _ if (isArrayType(t)): {
                var arr = typeAs(t, ArrayType);
                'go.GoArray<${genType(arr.elem())}, ${arr.len()}>';
            }

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
    public function elem():std.go.types.Types.Type;
}

// walk packages -> files -> types + recv funcs and funcs
