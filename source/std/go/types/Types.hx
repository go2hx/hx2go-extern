package std.go.types;

import haxe.Rest;
import std.go.token.Pos;
import std.go.token.FileSet;
import std.go.ast.Ast.Node;
import std.go.ast.Ast.Expr;
import std.go.ast.Ast.Ident;
import std.go.ast.Ast.File;
import std.go.ast.Ast.SelectorExpr;
import std.go.ast.Ast.ImportSpec;
import std.go.constant.Value;

// Externs for the Go standard library `go/types` package.
//
// Types that live outside `go/types` (token.*, ast.*, constant.Value, ...) are
// referenced but intentionally not declared here.
//
// Modelling notes:
//   *types.T -> go.Pointer<T>, []T -> go.Slice<T>, map[K]V -> go.Map<K, V>.
//   `go/types` declares Pointer/Slice/Map/Chan/Array/Struct/Error itself, so the
//   go-interop containers are always written fully-qualified (`go.Pointer<...>`)
//   to avoid clashing with the bare `go/types` names.
//   Type / Object / Sizes / Importer are interfaces and are used bare.
//   Most `go/types` structs are opaque (unexported fields); they expose data
//   through methods rather than fields.
//   The Go 1.23 `iter.Seq` iterator methods (Methods(), Fields(), ...) and the
//   io.Writer / bytes.Buffer helpers (WriteType, Scope.WriteTo, ...) are omitted;
//   the indexed accessors (Method(i)/NumMethods(), ...) and the *String helpers
//   cover the same ground without the exotic outside dependencies.

// ---------------------------------------------------------------------------
// Core interfaces
// ---------------------------------------------------------------------------

@:go.Type({ name: "types.Type", imports: ["go/types"] })
extern class Type {
    function underlying(): Type;
    function string(): String;
}

@:go.Type({ name: "types.Object", imports: ["go/types"] })
extern class Object {
    function parent(): go.Pointer<Scope>;
    function pos(): Pos;
    function pkg(): go.Pointer<Package>;
    function name(): String;
    function type(): Type;
    function exported(): Bool;
    function id(): String;
    function string(): String;
}

@:go.Type({ name: "types.Sizes", imports: ["go/types"] })
extern class Sizes {
    function alignof(T: Type): go.Int64;
    function offsetsof(fields: go.Slice<go.Pointer<Var>>): go.Slice<go.Int64>;
    function sizeof(T: Type): go.Int64;
}

@:go.Type({ name: "types.Importer", imports: ["go/types"] })
extern class Importer {
    @:native("Import") function importPkg(path: String): go.Result<go.Pointer<Package>>;
}

@:go.Type({ name: "types.ImporterFrom", imports: ["go/types"] })
extern class ImporterFrom extends Importer {
    function importFrom(path: String, dir: String, mode: ImportMode): go.Result<go.Pointer<Package>>;
}

// ---------------------------------------------------------------------------
// Types (implement Type)
// ---------------------------------------------------------------------------

@:go.Type({ name: "types.Basic", imports: ["go/types"] })
extern class Basic extends Type {
    function info(): BasicInfo;
    function kind(): BasicKind;
    function name(): String;
}

@:go.Type({ name: "types.Array", imports: ["go/types"] })
extern class Array extends Type {
    function elem(): Type;
    function len(): go.Int64;
}

@:go.Type({ name: "types.Slice", imports: ["go/types"] })
extern class Slice extends Type {
    function elem(): Type;
}

@:go.Type({ name: "types.Struct", imports: ["go/types"] })
extern class Struct extends Type {
    function field(i: Int): go.Pointer<Var>;
    function numFields(): Int;
    function tag(i: Int): String;
}

@:go.Type({ name: "types.Pointer", imports: ["go/types"] })
extern class Pointer extends Type {
    function elem(): Type;
}

@:go.Type({ name: "types.Tuple", imports: ["go/types"] })
extern class Tuple extends Type {
    function at(i: Int): go.Pointer<Var>;
    function len(): Int;
}

@:go.Type({ name: "types.Signature", imports: ["go/types"] })
extern class Signature extends Type {
    function params(): go.Pointer<Tuple>;
    function recv(): go.Pointer<Var>;
    function recvTypeParams(): go.Pointer<TypeParamList>;
    function results(): go.Pointer<Tuple>;
    function typeParams(): go.Pointer<TypeParamList>;
    function variadic(): Bool;
}

@:go.Type({ name: "types.Union", imports: ["go/types"] })
extern class Union extends Type {
    function len(): Int;
    function term(i: Int): go.Pointer<Term>;
}

@:go.Type({ name: "types.Interface", imports: ["go/types"] })
extern class Interface extends Type {
    function complete(): go.Pointer<Interface>;
    function embedded(i: Int): go.Pointer<Named>;
    function embeddedType(i: Int): Type;
    function empty(): Bool;
    function explicitMethod(i: Int): go.Pointer<Func>;
    function isComparable(): Bool;
    function isImplicit(): Bool;
    function isMethodSet(): Bool;
    function markImplicit(): Void;
    function method(i: Int): go.Pointer<Func>;
    function numEmbeddeds(): Int;
    function numExplicitMethods(): Int;
    function numMethods(): Int;
}

@:go.Type({ name: "types.Map", imports: ["go/types"] })
extern class Map extends Type {
    function elem(): Type;
    function key(): Type;
}

@:go.Type({ name: "types.Chan", imports: ["go/types"] })
extern class Chan extends Type {
    function dir(): ChanDir;
    function elem(): Type;
}

@:go.Type({ name: "types.Named", imports: ["go/types"] })
extern class Named extends Type {
    function addMethod(m: go.Pointer<Func>): Void;
    function method(i: Int): go.Pointer<Func>;
    function numMethods(): Int;
    function obj(): go.Pointer<TypeName>;
    function origin(): go.Pointer<Named>;
    function setTypeParams(tparams: go.Slice<go.Pointer<TypeParam>>): Void;
    function setUnderlying(u: Type): Void;
    function typeArgs(): go.Pointer<TypeList>;
    function typeParams(): go.Pointer<TypeParamList>;
}

@:go.Type({ name: "types.TypeParam", imports: ["go/types"] })
extern class TypeParam extends Type {
    function constraint(): Type;
    function index(): Int;
    function obj(): go.Pointer<TypeName>;
    function setConstraint(bound: Type): Void;
}

@:go.Type({ name: "types.Alias", imports: ["go/types"] })
extern class Alias extends Type {
    function obj(): go.Pointer<TypeName>;
    function origin(): go.Pointer<Alias>;
    function rhs(): Type;
    function setTypeParams(tparams: go.Slice<go.Pointer<TypeParam>>): Void;
    function typeArgs(): go.Pointer<TypeList>;
    function typeParams(): go.Pointer<TypeParamList>;
}

// ---------------------------------------------------------------------------
// Objects (implement Object)
// ---------------------------------------------------------------------------

@:go.Type({ name: "types.PkgName", imports: ["go/types"] })
extern class PkgName extends Object {
    function imported(): go.Pointer<Package>;
}

@:go.Type({ name: "types.Const", imports: ["go/types"] })
extern class Const extends Object {
    function val(): Value;
}

@:go.Type({ name: "types.TypeName", imports: ["go/types"] })
extern class TypeName extends Object {
    function isAlias(): Bool;
}

@:go.Type({ name: "types.Var", imports: ["go/types"] })
extern class Var extends Object {
    function anonymous(): Bool;
    function embedded(): Bool;
    function isField(): Bool;
    function kind(): VarKind;
    function origin(): go.Pointer<Var>;
    function setKind(kind: VarKind): Void;
}

@:go.Type({ name: "types.Func", imports: ["go/types"] })
extern class Func extends Object {
    function fullName(): String;
    function origin(): go.Pointer<Func>;
    function scope(): go.Pointer<Scope>;
    function signature(): go.Pointer<Signature>;
}

@:go.Type({ name: "types.Label", imports: ["go/types"] })
extern class Label extends Object {}

@:go.Type({ name: "types.Builtin", imports: ["go/types"] })
extern class Builtin extends Object {}

@:go.Type({ name: "types.Nil", imports: ["go/types"] })
extern class Nil extends Object {}

// ---------------------------------------------------------------------------
// Type/object lists and selections
// ---------------------------------------------------------------------------

@:go.Type({ name: "types.TypeList", imports: ["go/types"] })
extern class TypeList {
    function at(i: Int): Type;
    function len(): Int;
}

@:go.Type({ name: "types.TypeParamList", imports: ["go/types"] })
extern class TypeParamList {
    function at(i: Int): go.Pointer<TypeParam>;
    function len(): Int;
}

@:go.Type({ name: "types.MethodSet", imports: ["go/types"] })
extern class MethodSet {
    function at(i: Int): go.Pointer<Selection>;
    function len(): Int;
    function lookup(pkg: go.Pointer<Package>, name: String): go.Pointer<Selection>;
    function string(): String;
}

@:go.Type({ name: "types.Selection", imports: ["go/types"] })
extern class Selection {
    function index(): go.Slice<Int>;
    function indirect(): Bool;
    function kind(): SelectionKind;
    function obj(): Object;
    function recv(): Type;
    function string(): String;
    function type(): Type;
}

@:go.Type({ name: "types.Term", imports: ["go/types"] })
extern class Term {
    function string(): String;
    function tilde(): Bool;
    function type(): Type;
}

// ---------------------------------------------------------------------------
// Packages and scopes
// ---------------------------------------------------------------------------

@:go.Type({ name: "types.Package", imports: ["go/types"] })
extern class Package {
    function complete(): Bool;
    function goVersion(): String;
    function imports(): go.Slice<go.Pointer<Package>>;
    function markComplete(): Void;
    function name(): String;
    function path(): String;
    function scope(): go.Pointer<Scope>;
    function setImports(list: go.Slice<go.Pointer<Package>>): Void;
    function setName(name: String): Void;
    function string(): String;
}

@:go.Type({ name: "types.Scope", imports: ["go/types"] })
extern class Scope {
    function child(i: Int): go.Pointer<Scope>;
    function contains(pos: Pos): Bool;
    function end(): Pos;
    function innermost(pos: Pos): go.Pointer<Scope>;
    function insert(obj: Object): Object;
    function len(): Int;
    function lookup(name: String): Object;
    @:go.Tuple("scope", "obj")
    function lookupParent(name: String, pos: Pos): go.Tuple<{ scope: go.Pointer<Scope>, obj: Object }>;
    function names(): go.Slice<String>;
    function numChildren(): Int;
    function parent(): go.Pointer<Scope>;
    function pos(): Pos;
    function string(): String;
}

// ---------------------------------------------------------------------------
// Type checker
// ---------------------------------------------------------------------------

@:go.Type({ name: "types.Config", imports: ["go/types"] })
extern class Config {
    var context: go.Pointer<Context>;
    var goVersion: String;
    var ignoreFuncBodies: Bool;
    var fakeImportC: Bool;
    var error: go.Error -> Void;
    var importer: Importer;
    var sizes: Sizes;
    var disableUnusedImportCheck: Bool;

    function check(path: String, fset: go.Pointer<FileSet>, files: go.Slice<go.Pointer<File>>, info: go.Pointer<Info>): go.Result<go.Pointer<Package>>;
}

@:go.Type({ name: "types.Checker", imports: ["go/types"] })
extern class Checker {
    function files(files: go.Slice<go.Pointer<File>>): go.Error;
}

@:go.Type({ name: "types.Context", imports: ["go/types"] })
extern class Context {}

@:go.Type({ name: "types.Info", imports: ["go/types"] })
extern class Info {
    var types: go.Map<Expr, TypeAndValue>;
    var instances: go.Map<go.Pointer<Ident>, Instance>;
    var defs: go.Map<go.Pointer<Ident>, Object>;
    var uses: go.Map<go.Pointer<Ident>, Object>;
    var implicits: go.Map<Node, Object>;
    var selections: go.Map<go.Pointer<SelectorExpr>, go.Pointer<Selection>>;
    var scopes: go.Map<Node, go.Pointer<Scope>>;
    var initOrder: go.Slice<go.Pointer<Initializer>>;
    var fileVersions: go.Map<go.Pointer<File>, String>;

    function objectOf(id: go.Pointer<Ident>): Object;
    function pkgNameOf(imp: go.Pointer<ImportSpec>): go.Pointer<PkgName>;
    function typeOf(e: Expr): Type;
}

@:go.Type({ name: "types.Initializer", imports: ["go/types"] })
extern class Initializer {
    var lhs: go.Slice<go.Pointer<Var>>;
    var rhs: Expr;

    function string(): String;
}

@:go.Type({ name: "types.Instance", imports: ["go/types"] })
extern class Instance {
    var typeArgs: go.Pointer<TypeList>;
    var type: Type;
}

@:go.Type({ name: "types.TypeAndValue", imports: ["go/types"] })
extern class TypeAndValue {
    var type: Type;
    var value: Value;

    function addressable(): Bool;
    function assignable(): Bool;
    function hasOk(): Bool;
    function isBuiltin(): Bool;
    function isNil(): Bool;
    function isType(): Bool;
    function isValue(): Bool;
    function isVoid(): Bool;
}

// ---------------------------------------------------------------------------
// Errors and sizing
// ---------------------------------------------------------------------------

@:go.Type({ name: "types.ArgumentError", imports: ["go/types"] })
extern class ArgumentError {
    var index: Int;
    var err: go.Error;

    function error(): String;
    function unwrap(): go.Error;
}

@:go.Type({ name: "types.Error", imports: ["go/types"] })
extern class Error {
    var fset: go.Pointer<FileSet>;
    var pos: Pos;
    var msg: String;
    var soft: Bool;

    function error(): String;
}

@:go.Type({ name: "types.StdSizes", imports: ["go/types"] })
extern class StdSizes {
    var wordSize: go.Int64;
    var maxAlign: go.Int64;

    function alignof(T: Type): go.Int64;
    function offsetsof(fields: go.Slice<go.Pointer<Var>>): go.Slice<go.Int64>;
    function sizeof(T: Type): go.Int64;
}

// ---------------------------------------------------------------------------
// Named integer/flag types and their package-level constants
// ---------------------------------------------------------------------------

@:go.Type({ name: "types.BasicKind", imports: ["go/types"] })
extern class BasicKind {
    static var invalid: BasicKind;
    static var bool: BasicKind;
    static var int: BasicKind;
    static var int8: BasicKind;
    static var int16: BasicKind;
    static var int32: BasicKind;
    static var int64: BasicKind;
    static var uint: BasicKind;
    static var uint8: BasicKind;
    static var uint16: BasicKind;
    static var uint32: BasicKind;
    static var uint64: BasicKind;
    static var uintptr: BasicKind;
    static var float32: BasicKind;
    static var float64: BasicKind;
    static var complex64: BasicKind;
    static var complex128: BasicKind;
    static var string: BasicKind;
    static var unsafePointer: BasicKind;
    static var untypedBool: BasicKind;
    static var untypedInt: BasicKind;
    static var untypedRune: BasicKind;
    static var untypedFloat: BasicKind;
    static var untypedComplex: BasicKind;
    static var untypedString: BasicKind;
    static var untypedNil: BasicKind;
    static var byte: BasicKind;
    static var rune: BasicKind;
}

@:go.Type({ name: "types.BasicInfo", imports: ["go/types"] })
extern class BasicInfo {
    static var isBoolean: BasicInfo;
    static var isInteger: BasicInfo;
    static var isUnsigned: BasicInfo;
    static var isFloat: BasicInfo;
    static var isComplex: BasicInfo;
    static var isString: BasicInfo;
    static var isUntyped: BasicInfo;
    static var isOrdered: BasicInfo;
    static var isNumeric: BasicInfo;
    static var isConstType: BasicInfo;
}

@:go.Type({ name: "types.ChanDir", imports: ["go/types"] })
extern class ChanDir {
    static var sendRecv: ChanDir;
    static var sendOnly: ChanDir;
    static var recvOnly: ChanDir;
}

@:go.Type({ name: "types.SelectionKind", imports: ["go/types"] })
extern class SelectionKind {
    static var fieldVal: SelectionKind;
    static var methodVal: SelectionKind;
    static var methodExpr: SelectionKind;
}

@:go.Type({ name: "types.VarKind", imports: ["go/types"] })
extern class VarKind {
    static var packageVar: VarKind;
    static var localVar: VarKind;
    static var recvVar: VarKind;
    static var paramVar: VarKind;
    static var resultVar: VarKind;
    static var fieldVar: VarKind;

    function string(): String;
}

@:go.Type({ name: "types.ImportMode", imports: ["go/types"] })
extern class ImportMode {}

// ---------------------------------------------------------------------------
// Package-level variables, constructors and functions
// ---------------------------------------------------------------------------

@:go.Type({ imports: ["go/types"] })
extern class Types {
    // Package-level variables.
    static var typ: go.Slice<go.Pointer<Basic>>;
    static var universe: go.Pointer<Scope>;
    static var unsafe: go.Pointer<Package>;

    // Constructors.
    static function newArray(elem: Type, len: go.Int64): go.Pointer<Array>;
    static function newChan(dir: ChanDir, elem: Type): go.Pointer<Chan>;
    static function newInterface(methods: go.Slice<go.Pointer<Func>>, embeddeds: go.Slice<go.Pointer<Named>>): go.Pointer<Interface>;
    static function newInterfaceType(methods: go.Slice<go.Pointer<Func>>, embeddeds: go.Slice<Type>): go.Pointer<Interface>;
    static function newMap(key: Type, elem: Type): go.Pointer<Map>;
    static function newNamed(obj: go.Pointer<TypeName>, underlying: Type, methods: go.Slice<go.Pointer<Func>>): go.Pointer<Named>;
    static function newPointer(elem: Type): go.Pointer<Pointer>;
    static function newSignature(recv: go.Pointer<Var>, params: go.Pointer<Tuple>, results: go.Pointer<Tuple>, variadic: Bool): go.Pointer<Signature>;
    static function newSignatureType(recv: go.Pointer<Var>, recvTypeParams: go.Slice<go.Pointer<TypeParam>>, typeParams: go.Slice<go.Pointer<TypeParam>>, params: go.Pointer<Tuple>, results: go.Pointer<Tuple>, variadic: Bool): go.Pointer<Signature>;
    static function newSlice(elem: Type): go.Pointer<Slice>;
    static function newStruct(fields: go.Slice<go.Pointer<Var>>, tags: go.Slice<String>): go.Pointer<Struct>;
    static function newTuple(vars: Rest<go.Pointer<Var>>): go.Pointer<Tuple>;
    static function newUnion(terms: go.Slice<go.Pointer<Term>>): go.Pointer<Union>;
    static function newTerm(tilde: Bool, typ: Type): go.Pointer<Term>;
    static function newTypeParam(obj: go.Pointer<TypeName>, constraint: Type): go.Pointer<TypeParam>;
    static function newAlias(obj: go.Pointer<TypeName>, rhs: Type): go.Pointer<Alias>;
    static function newConst(pos: Pos, pkg: go.Pointer<Package>, name: String, typ: Type, val: Value): go.Pointer<Const>;
    static function newFunc(pos: Pos, pkg: go.Pointer<Package>, name: String, sig: go.Pointer<Signature>): go.Pointer<Func>;
    static function newLabel(pos: Pos, pkg: go.Pointer<Package>, name: String): go.Pointer<Label>;
    static function newPkgName(pos: Pos, pkg: go.Pointer<Package>, name: String, imported: go.Pointer<Package>): go.Pointer<PkgName>;
    static function newTypeName(pos: Pos, pkg: go.Pointer<Package>, name: String, typ: Type): go.Pointer<TypeName>;
    static function newVar(pos: Pos, pkg: go.Pointer<Package>, name: String, typ: Type): go.Pointer<Var>;
    static function newField(pos: Pos, pkg: go.Pointer<Package>, name: String, typ: Type, embedded: Bool): go.Pointer<Var>;
    static function newParam(pos: Pos, pkg: go.Pointer<Package>, name: String, typ: Type): go.Pointer<Var>;
    static function newMethodSet(T: Type): go.Pointer<MethodSet>;
    static function newContext(): go.Pointer<Context>;
    static function newChecker(conf: go.Pointer<Config>, fset: go.Pointer<FileSet>, pkg: go.Pointer<Package>, info: go.Pointer<Info>): go.Pointer<Checker>;
    static function newPackage(path: String, name: String): go.Pointer<Package>;
    static function newScope(parent: go.Pointer<Scope>, pos: Pos, end: Pos, comment: String): go.Pointer<Scope>;

    // Predicates and helpers.
    static function assertableTo(V: go.Pointer<Interface>, T: Type): Bool;
    static function assignableTo(V: Type, T: Type): Bool;
    static function comparable(T: Type): Bool;
    static function convertibleTo(V: Type, T: Type): Bool;
    @:native("Default") static function defaultType(t: Type): Type;
    static function identical(x: Type, y: Type): Bool;
    static function identicalIgnoreTags(x: Type, y: Type): Bool;
    @:native("Implements") static function implementsInterface(V: Type, T: go.Pointer<Interface>): Bool;
    static function isInterface(t: Type): Bool;
    static function satisfies(V: Type, T: go.Pointer<Interface>): Bool;
    static function unalias(t: Type): Type;
    static function sizesFor(compiler: String, arch: String): Sizes;

    // Instantiation and lookup.
    static function instantiate(ctxt: go.Pointer<Context>, orig: Type, targs: go.Slice<Type>, validate: Bool): go.Result<Type>;
    @:go.Tuple("obj", "index", "indirect")
    static function lookupFieldOrMethod(T: Type, addressable: Bool, pkg: go.Pointer<Package>, name: String): go.Tuple<{ obj: Object, index: go.Slice<Int>, indirect: Bool }>;
    @:go.Tuple("selection", "ok")
    static function lookupSelection(T: Type, addressable: Bool, pkg: go.Pointer<Package>, name: String): go.Tuple<{ selection: Selection, ok: Bool }>;
    @:go.Tuple("method", "wrongType")
    static function missingMethod(V: Type, T: go.Pointer<Interface>, staticOnly: Bool): go.Tuple<{ method: go.Pointer<Func>, wrongType: Bool }>;

    // String formatting.
    static function id(pkg: go.Pointer<Package>, name: String): String;
    static function exprString(x: Expr): String;
    static function objectString(obj: Object, qf: go.Pointer<Package> -> String): String;
    static function selectionString(s: go.Pointer<Selection>, qf: go.Pointer<Package> -> String): String;
    static function typeString(typ: Type, qf: go.Pointer<Package> -> String): String;
    static function relativeTo(pkg: go.Pointer<Package>): go.Pointer<Package> -> String;

    // Type-checking entry points.
    static function checkExpr(fset: go.Pointer<FileSet>, pkg: go.Pointer<Package>, pos: Pos, expr: Expr, info: go.Pointer<Info>): go.Error;
    static function eval(fset: go.Pointer<FileSet>, pkg: go.Pointer<Package>, pos: Pos, expr: String): go.Result<TypeAndValue>;
    static function defPredeclaredTestFuncs(): Void;
}
