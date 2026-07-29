import go.Syntax;
import go.Pointer;
using StringTools;
import go.go.types.Pointer as PointerType;
import go.go.types.Chan as ChanType;
import go.go.types.Named;
import go.go.types.Slice as SliceType;


class TypeHelper {
    public inline extern static function typeAs<V, T>(value: V, as: Class<T>): T {
        var v: Pointer<T> = cast (value : Dynamic);
        return v.value;
    }

    public static function isNamedType(t:go.go.types.Type): Bool {
        var isNamed = false;
        Syntax.code("if _, ntOk := t.(*types.Named); ntOk { isNamed = true; }");

        return isNamed;
    }

    public static function isExportedType(t:go.go.types.Type):Bool {
        if (isNamedType(t)) {
            var named = TypeHelper.typeAs(t, Named);
            return named.obj().value.exported();
        }

        var s = t.string();

        if (s.startsWith("*")) {
            return isExportedType(TypeHelper.typeAs(t, PointerType).elem());
        }

        if (s.startsWith("[]")) {
            return isExportedType(TypeHelper.typeAs(t, SliceType).elem());
        }

        if (s.startsWith("chan ")) {
            return isExportedType(TypeHelper.typeAs(t, ChanType).elem());
        }

        return true;
    }

    public static function isInterfaceType(t:go.go.types.Type):Bool {
        var result = false;

        Syntax.code("
            if _, ok := t.(*types.Interface); ok {
                result = true;
            }
        ");

        return result;
    }

    public static function isTypeParamType(t:go.go.types.Type):Bool {
        var result = false;

        Syntax.code("
            if _, ok := t.(*types.TypeParam); ok {
                result = true;
            }
        ");

        return result;
    }

    public static function getUnderlying(t:go.go.types.Type):go.go.types.Type {
        if (TypeHelper.isNamedType(t)) {
            // TODO: change out temp fix
            return Syntax.code("((types.Type)(((t).Underlying())))", (t : go.go.types.Type));
            // return TypeHelper.typeAs(t, Named).underlying();
        }

        return t;
    }

    public static function isBasicType(t:go.go.types.Type):Bool {
        var result = false;

        Syntax.code("
            if _, ok := t.(*types.Basic); ok {
                result = true;
            }
        ");

        return result;
    }

    public static function isStructType(t:go.go.types.Type):Bool {
        var result = false;

        Syntax.code("
        if _, ok := t.(*types.Struct); ok {
            result = true;
        }
    ");

        return result;
    }

    public static function isArrayType(t:go.go.types.Type):Bool {
        var result = false;

        Syntax.code("
        if _, ok := t.(*types.Array); ok {
            result = true;
        }
    ");

        return result;
    }
}