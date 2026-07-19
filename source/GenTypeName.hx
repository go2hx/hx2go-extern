import go.go.types.*;
import go.Syntax;

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
        } else if (type.isAlias() && (TypeHelper.isNamedType(resolvedTo) || TypeHelper.isBasicType(resolvedTo))) { // TODO: support for func() aswell (see iter.Seq)
            out.typedefStr = Main.genType(resolvedTo);
        } else {
            var isNamed = TypeHelper.isNamedType(type.type());
            var out = getOutput(obj.name());

            if (isNamed) {
                var named = TypeHelper.typeAs(type.type(), Named);

                if (TypeHelper.isStructType(underlying)) {
                    out.isStruct = true;
                    var struct = TypeHelper.typeAs(underlying, Struct);

                    for (i in 0...struct.numFields()) {
                        var field = TypeHelper.typeAs(struct.field(i), Var);
                        if (!field.exported()) {
                            continue;
                        }

                        var fname = field.name();
                        var p = '${Sanitize.name(Main.toHaxeCase(fname))}: ${Main.genType(field.type())}';
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
                        var constraintStr = Main.genType(constraint);
                        params.push('${t.string()}: ${constraintStr}');
                    }

                    out.paramStr = params.length == 0 ? "" : "<" + params.join(", ") + ">";
                }

                var methodSet = go.go.Types.newMethodSet(Syntax.code("types.NewPointer({0})", type.type()));
                for (i in 0...methodSet.value.len()) {
                    var sel = methodSet.value.at(i).value;
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
    }