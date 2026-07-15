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
import std.go.os.Os;
import std.go.types.Types.Var;
import std.go.types.Types.Named;

class Main {

    public static var didGen: Map<String, Bool> = new Map();

    public static function toHaxeCase(input: String): String {
        return input.charAt(0).toLowerCase() + input.substr(1);
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

    static function genLib(lib: String, output: String): Void {
        if (didGen.exists(lib)) {
            return;
        }

        didGen.set(lib, true);

        var config: Config = {
            mode: LoadMode.needName.or(LoadMode.needTypes).or(LoadMode.needTypesInfo).or(LoadMode.needSyntax).or(LoadMode.needImports)
        };

        var entries = Packages.load(config, lib).sure();
        var outputs: Map<String, { paramStr: String, staticFunctions: StringBuf, instanceFunctions: StringBuf, staticVars: StringBuf, instanceVars: StringBuf }> = new Map();

        function getOutput(name: String) {
            if (!outputs.exists(name)) {
                outputs[name] = {
                    staticFunctions: new StringBuf(),
                    instanceFunctions: new StringBuf(),
                    staticVars: new StringBuf(),
                    instanceVars: new StringBuf(),
                    paramStr: ''
                };
            }

            return outputs[name];
        }

        for (entry in entries) {
            var scope = entry.value.types.value.scope().value;
            for (dep in entry.value.imports.keys()) {
                genLib(dep, output);
            }

            for (name in scope.names()) {
                var obj = scope.lookup(name);
                if (!obj.exported()) {
                    continue;
                }

                Syntax.code("switch {0}.(type) {", obj); // this is so bad :[
                Syntax.code("case *types.TypeName:"); {
                    var type = typeAs(obj, TypeName);
                    var isNamed = isNamedType(type.type());
                    var out = getOutput(obj.name());

                    if (isNamed) {
                        var named = typeAs(type.type(), Named);

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
                    var buf = getOutput(entry.value.name).staticVars;
                    var name = v.name();
                    var type = v.type();

                    buf.add('    static var ${name}: ${genType(type)};\n');
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

            buf.add('package go.${lib.replace("/", ".")};\n');
            buf.add('\n');
            buf.add('@:go.Type({ name: "${file}", instanceName: "${lib.split("/").pop()}.${file}", imports: ["${lib}"] })\n');
            buf.add('extern class ${toPascalCase(file)}${out.paramStr} {\n\n');
            buf.add(out.staticVars.toString());
            buf.add(out.instanceVars.toString());
            if (out.staticVars.length > 0 || out.instanceVars.length > 0) buf.add('\n');
            buf.add(out.staticFunctions.toString());
            buf.add(out.instanceFunctions.toString());
            if (out.staticFunctions.length > 0 || out.instanceFunctions.length > 0) buf.add("\n");
            buf.add("}");

            Os.mkdirAll('${output}/go/${lib}', Syntax.code("0775"));
            Os.writeFile('${output}/go/${lib}/${toPascalCase(file)}.hx', cast buf.toString(), Syntax.code("0666"));
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

        var tParams = sig.typeParams() != null ? sig.typeParams().value : null;
        var tParamsStr = '';
        if (tParams != null){
            var tParamsLocal = [];
            for (i in 0...tParams.len()) {
                var t = tParams.at(i).value;
                var constraint = t.constraint();
                var constraintStr = genType(constraint);
                if (constraintStr.startsWith("~")) {
                    constraintStr = "Dynamic";
                }

                tParamsLocal.push('${t.string()}: ${constraintStr}');
            }
            tParamsStr = '<' + tParamsLocal.join(", ") + '>';
        }

        return '${meta}${topLevel && !closure ? "static " : ""}${closure ? "" : 'function ${toHaxeCase(name)}'}${tParamsStr}(${params.join(", ")})${closure ? ' -> ' : ': '}${results.len() == 0 ? "Void" :genResults(results)}${closure ? "" : ";"}';
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
                else n + ": " + genType(v.type())
            );
        }

        return items;
    }

    inline extern static function typeAs<V, T>(value: V, as: Class<T>): T {
        var v: Pointer<T> = cast (value : Dynamic);
        return v.value;
    }

    static function resolvePath(path: String): String {
        return 'go.${path.replace("/", ".")}'; // TODO: this is a stub
    }

    static function isNamedType(t:std.go.types.Types.Type): Bool {
        var isNamed = false;
        Syntax.code("if _, ntOk := t.(*types.Named); ntOk { isNamed = true; }");

        return isNamed;
    }

    static function genType(t: std.go.types.Types.Type): String {
        var s = t.string();
        var tParamStr = '';

        if (isNamedType(t) && typeAs(t, Named).typeParams() != null) {
            var typeParams = typeAs(t, Named).typeArgs().value;
            var typeParamStrs = [];
            for (i in 0...typeParams.len()) {
                var r = genType(typeParams.at(i));
                typeParamStrs.push(r.startsWith("~") ? "Dynamic" : r);
            }

            var typeParamStart = s.indexOf('[');
            var typeParamEnd = s.lastIndexOf(']');
            if (typeParamStart != -1 && typeParamEnd != -1 && typeParamEnd > typeParamStart) {
                s = s.substr(0, typeParamStart) + s.substr(typeParamEnd + 1);
            }

            tParamStr = '<' + typeParamStrs.join(", ") + '>';
        }

        var q = switch s {
            case "error": "go.Error";
            case "string": "String";
            case "bool": "Bool";
            case "any": "Dynamic";

            case "int": "Int";
            case "int64": "go.Int64";
            case "int32": "go.Int32";
            case "int16": "go.Int16";
            case "int8": "go.Int8";

            case "uint": "go.UInt";
            case "uint64": "go.UInt64";
            case "uint32": "go.UInt32";
            case "uint16": "go.UInt16";
            case "uint8": "go.UInt8";

            case "float64": "Float";
            case "float32": "go.Float32";
            case "float16": "go.Float16";

            case "byte": "go.Byte";
            case "rune": "go.Rune";

            case "uintptr": "go.UIntPtr";
            case "complex128": "go.Complex128";
            case "complex64": "go.Complex64";
            case "comparable": "go.Comparable";

            case _ if (s.startsWith("chan ")): 'go.Chan<${genType(typeAs(t, Chan).elem())}>';
            case _ if (s.startsWith("chan ")): 'go.Chan<${genType(typeAs(t, Chan).elem())}>';
            case _ if (s.startsWith("[]")): 'go.Slice<${genType(typeAs(t, Slice).elem())}>';
            case _ if (s.startsWith("*")): 'go.Pointer<${genType(typeAs(t, PointerType).elem())}>';
            case _ if (s.startsWith('func')): {
                var sig = typeAs(t, Signature);
                genFunc('', sig, false, true);
            }

            case _ if (s.contains(".")): resolvePath(s);

            case _: t.string();
        }

        return q + tParamStr;
    }
}

interface ElemType {
    public function elem():std.go.types.Types.Type;
}

// walk packages -> files -> types + recv funcs and funcs
