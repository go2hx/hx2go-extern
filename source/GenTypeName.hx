import go.go.types.*;
import go.Syntax;
import go.Map;

class GenTypeName {
    public static function gen(obj:Object, getOutput:String->GenOutput) {
        var type = TypeHelper.typeAs(obj, TypeName);

        var out = getOutput(obj.name());
        var underlying = TypeHelper.getUnderlying(type.type());

        var resolvedTo = go.go.Types.unalias(type.type());
        if (resolvedTo == type.type() && TypeHelper.isNamedType(resolvedTo)) {
            var named = TypeHelper.typeAs(resolvedTo, Named);
            resolvedTo = named.underlying();
        }

        if (TypeHelper.isInterfaceType(underlying)) {
            var iface = TypeHelper.typeAs(underlying, Interface);
            out.isInterface = true;

            iface.complete();

            for (i in 0...iface.numMethods()) {
                var method = TypeHelper.typeAs(iface.method(i), Func);
                if (!method.exported()) {
                    continue;
                }

                var sig = method.signature().value;
                out.instanceFunctions.add('    ${Main.genFunc(method.name(), sig, false, false)}\n');
            }
        } else {
            if (type.isAlias() && (TypeHelper.isNamedType(resolvedTo) || TypeHelper.isBasicType(resolvedTo))) { // TODO: support for func() aswell (see iter.Seq)
                out.typedefStr = Main.genType(resolvedTo);
            }
            var isNamed = TypeHelper.isNamedType(type.type());
            var out = getOutput(obj.name());

            if (isNamed) {
                var named = TypeHelper.typeAs(type.type(), Named);

                var methodSet = go.go.Types.newMethodSet(Syntax.code("types.NewPointer({0})", type.type()));
                var seen = new Map<String, Bool>();
                for (i in 0...methodSet.len()) {
                    var method = TypeHelper.typeAs(methodSet.at(i).obj(), Func);
                    if (method.exported()) {
                        seen.set(Sanitize.name(Main.toHaxeCase(method.name())), true);
                    }
                }

                if (TypeHelper.isStructType(underlying)) {
                    out.isStruct = true;
                    addStructFields(underlying, out, seen);
                }

                var tp = named.typeParams();
                if (tp != null) {
                    var params = [];
                    var tps = tp.value;

                    for (i in 0...tps.len()) {
                        var t = tps.at(i).value;
                        var constraint = t.constraint();
                        var constraintStr = Main.genType(constraint);
                        params.push('${t.string()}: ${constraintStr}');
                    }

                    out.paramStr = params.length == 0 ? "" : "<" + params.join(", ") + ">";
                }

                for (i in 0...methodSet.len()) {
                    var sel = methodSet.at(i);
                    var method = TypeHelper.typeAs(sel.obj(), Func);

                    if (!method.exported()) {
                        continue;
                    }

                    var sig = method.signature().value;
                    out.instanceFunctions.add('    ${Main.genFunc(method.name(), sig, false)}\n');
                }
            }
        }
    }

    static function addStructFields(t: go.go.types.Type, out: GenOutput, seen: Map<String, Bool>) {
        addStructFieldsRec(t, out, 0, seen);
    }

    static function addStructFieldsRec(t: go.go.types.Type, out: GenOutput, depth: Int, seen: Map<String, Bool>) {
        // guard against cyclic embeds e.g. mutual *T
        if (depth > 16) {
            return;
        }

        var struct = TypeHelper.typeAs(t, Struct);

        for (i in 0...struct.numFields()) {
            var field = TypeHelper.typeAs(struct.field(i), Var);
            var fname = field.name();
            var hname = Sanitize.name(Main.toHaxeCase(fname));
            if (!field.exported() || seen.exists(hname)) {
                continue;
            }
            seen.set(hname, true);

            var p = '${hname}: ${Main.genType(field.type())}';
            out.instanceVars.add('    @:native("${fname}") var ${p};\n');
            if (depth == 0) {
                out.ctorParams.push(p);
                out.ctorValues.push(getTypeDefaultValue(field.type()));
            }
        }

        for (i in 0...struct.numFields()) {
            var field = TypeHelper.typeAs(struct.field(i), Var);
            if (!field.embedded()) {
                continue;
            }
            var est = embeddedStructType(field.type());
            if (est != null) {
                addStructFieldsRec(est, out, depth + 1, seen);
            }
        }
    }

    static function getTypeDefaultValue(t:go.go.types.Type):String {
        var isNamed = false;
        Syntax.code("if _, ok := {0}.(*types.Named); ok { {1} = true }", t, isNamed);

        t.underlying();
        var s = "";
        Syntax.code("switch u := {0}.(type) {", t.underlying());
        Syntax.code("case *types.Pointer, *types.Slice, *types.Map, *types.Chan, *types.Signature, *types.Interface:"); 
            s = 'null';
        Syntax.code("case *types.Basic:");
            Syntax.code("switch u.Kind() {");
            Syntax.code("case types.UnsafePointer:");
                s = "null";
            Syntax.code("case types.Bool:");
                s = isNamed ? "cast false" : "false";
            Syntax.code("case types.String:");
                s = isNamed ? 'cast ""' : '""';
            Syntax.code("case types.Float32, types.Float64:");
                s = isNamed ? 'cast 0.0' : '0.0';
            Syntax.code("default:");
                Syntax.code("_ = u");
                s = isNamed ? "cast 0" : "0";
            Syntax.code("}");
        Syntax.code("case *types.TypeParam:");
        Syntax.code("}");
        return s;
    }

    static function embeddedStructType(t: go.go.types.Type): go.go.types.Type {
        if (t.string().charAt(0) == "*") {
            t = TypeHelper.typeAs(t, Pointer).elem();
        }

        var u = TypeHelper.getUnderlying(t);
        return TypeHelper.isStructType(u) ? u : null;
    }
}