package std.go.ast;

import go.Slice;
import go.Pointer;
import go.Map;
import go.Error;
import std.go.token.Pos;
import std.go.token.FileSet;
import std.go.token.Token;

// Externs for the Go standard library `go/ast` package.
//
// Types that live outside `go/ast` (token.Pos, token.Token, token.FileSet, ...)
// are referenced but intentionally not declared here.
//
// Modelling notes:
//   *ast.T  -> Pointer<T>
//   []T     -> Slice<T>
//   Expr / Stmt / Decl / Spec / Node are interfaces and are used bare.
//   Field / method names are written in lowerCamelCase; the codegen re-pascalises
//   them (e.g. `tokPos` -> `TokPos`). @:native pins names the pascaliser can't
//   reach (Go all-caps constants) or Haxe keywords (`for`, `if`, `else`, ...).

// ---------------------------------------------------------------------------
// Interfaces
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.Node", imports: ["go/ast"] })
extern class Node {
    function pos(): Pos;
    function end(): Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.Expr", imports: ["go/ast"] })
extern class Expr extends Node {}

@:go.Type({ name: "ast", instanceName: "ast.Stmt", imports: ["go/ast"] })
extern class Stmt extends Node {}

@:go.Type({ name: "ast", instanceName: "ast.Decl", imports: ["go/ast"] })
extern class Decl extends Node {}

@:go.Type({ name: "ast", instanceName: "ast.Spec", imports: ["go/ast"] })
extern class Spec extends Node {}

@:go.Type({ name: "ast", instanceName: "ast.Visitor", imports: ["go/ast"] })
extern class Visitor {
    function visit(node: Node): Visitor;
}

// ---------------------------------------------------------------------------
// Expressions
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.BadExpr", imports: ["go/ast"] })
extern class BadExpr extends Expr {
    var from: Pos;
    var to: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.Ident", imports: ["go/ast"] })
extern class Ident extends Expr {
    var namePos: Pos;
    var name: String;
    var obj: Pointer<Object>;

    function isExported(): Bool;
    function string(): String;
}

@:go.Type({ name: "ast", instanceName: "ast.Ellipsis", imports: ["go/ast"] })
extern class Ellipsis extends Expr {
    var ellipsis: Pos;
    var elt: Expr;
}

@:go.Type({ name: "ast", instanceName: "ast.BasicLit", imports: ["go/ast"] })
extern class BasicLit extends Expr {
    var valuePos: Pos;
    var valueEnd: Pos;
    var kind: Token;
    var value: String;
}

@:go.Type({ name: "ast", instanceName: "ast.FuncLit", imports: ["go/ast"] })
extern class FuncLit extends Expr {
    var type: Pointer<FuncType>;
    var body: Pointer<BlockStmt>;
}

@:go.Type({ name: "ast", instanceName: "ast.CompositeLit", imports: ["go/ast"] })
extern class CompositeLit extends Expr {
    var type: Expr;
    var lbrace: Pos;
    var elts: Slice<Expr>;
    var rbrace: Pos;
    var incomplete: Bool;
}

@:go.Type({ name: "ast", instanceName: "ast.ParenExpr", imports: ["go/ast"] })
extern class ParenExpr extends Expr {
    var lparen: Pos;
    var x: Expr;
    var rparen: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.SelectorExpr", imports: ["go/ast"] })
extern class SelectorExpr extends Expr {
    var x: Expr;
    var sel: Pointer<Ident>;
}

@:go.Type({ name: "ast", instanceName: "ast.IndexExpr", imports: ["go/ast"] })
extern class IndexExpr extends Expr {
    var x: Expr;
    var lbrack: Pos;
    var index: Expr;
    var rbrack: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.IndexListExpr", imports: ["go/ast"] })
extern class IndexListExpr extends Expr {
    var x: Expr;
    var lbrack: Pos;
    var indices: Slice<Expr>;
    var rbrack: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.SliceExpr", imports: ["go/ast"] })
extern class SliceExpr extends Expr {
    var x: Expr;
    var lbrack: Pos;
    var low: Expr;
    var high: Expr;
    var max: Expr;
    var slice3: Bool;
    var rbrack: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.TypeAssertExpr", imports: ["go/ast"] })
extern class TypeAssertExpr extends Expr {
    var x: Expr;
    var lparen: Pos;
    var type: Expr;
    var rparen: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.CallExpr", imports: ["go/ast"] })
extern class CallExpr extends Expr {
    var fun: Expr;
    var lparen: Pos;
    var args: Slice<Expr>;
    var ellipsis: Pos;
    var rparen: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.StarExpr", imports: ["go/ast"] })
extern class StarExpr extends Expr {
    var star: Pos;
    var x: Expr;
}

@:go.Type({ name: "ast", instanceName: "ast.UnaryExpr", imports: ["go/ast"] })
extern class UnaryExpr extends Expr {
    var opPos: Pos;
    var op: Token;
    var x: Expr;
}

@:go.Type({ name: "ast", instanceName: "ast.BinaryExpr", imports: ["go/ast"] })
extern class BinaryExpr extends Expr {
    var x: Expr;
    var opPos: Pos;
    var op: Token;
    var y: Expr;
}

@:go.Type({ name: "ast", instanceName: "ast.KeyValueExpr", imports: ["go/ast"] })
extern class KeyValueExpr extends Expr {
    var key: Expr;
    var colon: Pos;
    var value: Expr;
}

// ---------------------------------------------------------------------------
// Type expressions (these also implement Expr)
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.ArrayType", imports: ["go/ast"] })
extern class ArrayType extends Expr {
    var lbrack: Pos;
    var len: Expr;
    var elt: Expr;
}

@:go.Type({ name: "ast", instanceName: "ast.StructType", imports: ["go/ast"] })
extern class StructType extends Expr {
    var struct: Pos;
    var fields: Pointer<FieldList>;
    var incomplete: Bool;
}

@:go.Type({ name: "ast", instanceName: "ast.FuncType", imports: ["go/ast"] })
extern class FuncType extends Expr {
    var func: Pos;
    var typeParams: Pointer<FieldList>;
    var params: Pointer<FieldList>;
    var results: Pointer<FieldList>;
}

@:go.Type({ name: "ast", instanceName: "ast.InterfaceType", imports: ["go/ast"] })
extern class InterfaceType extends Expr {
    @:native("Interface") var interfacePos: Pos;
    var methods: Pointer<FieldList>;
    var incomplete: Bool;
}

@:go.Type({ name: "ast", instanceName: "ast.MapType", imports: ["go/ast"] })
extern class MapType extends Expr {
    var map: Pos;
    var key: Expr;
    var value: Expr;
}

@:go.Type({ name: "ast", instanceName: "ast.ChanType", imports: ["go/ast"] })
extern class ChanType extends Expr {
    var begin: Pos;
    var arrow: Pos;
    var dir: ChanDir;
    var value: Expr;
}

// ---------------------------------------------------------------------------
// Statements
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.BadStmt", imports: ["go/ast"] })
extern class BadStmt extends Stmt {
    var from: Pos;
    var to: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.DeclStmt", imports: ["go/ast"] })
extern class DeclStmt extends Stmt {
    var decl: Decl;
}

@:go.Type({ name: "ast", instanceName: "ast.EmptyStmt", imports: ["go/ast"] })
extern class EmptyStmt extends Stmt {
    var semicolon: Pos;
    var implicit: Bool;
}

@:go.Type({ name: "ast", instanceName: "ast.LabeledStmt", imports: ["go/ast"] })
extern class LabeledStmt extends Stmt {
    var label: Pointer<Ident>;
    var colon: Pos;
    var stmt: Stmt;
}

@:go.Type({ name: "ast", instanceName: "ast.ExprStmt", imports: ["go/ast"] })
extern class ExprStmt extends Stmt {
    var x: Expr;
}

@:go.Type({ name: "ast", instanceName: "ast.SendStmt", imports: ["go/ast"] })
extern class SendStmt extends Stmt {
    var chan: Expr;
    var arrow: Pos;
    var value: Expr;
}

@:go.Type({ name: "ast", instanceName: "ast.IncDecStmt", imports: ["go/ast"] })
extern class IncDecStmt extends Stmt {
    var x: Expr;
    var tokPos: Pos;
    var tok: Token;
}

@:go.Type({ name: "ast", instanceName: "ast.AssignStmt", imports: ["go/ast"] })
extern class AssignStmt extends Stmt {
    var lhs: Slice<Expr>;
    var tokPos: Pos;
    var tok: Token;
    var rhs: Slice<Expr>;
}

@:go.Type({ name: "ast", instanceName: "ast.GoStmt", imports: ["go/ast"] })
extern class GoStmt extends Stmt {
    var go: Pos;
    var call: Pointer<CallExpr>;
}

@:go.Type({ name: "ast", instanceName: "ast.DeferStmt", imports: ["go/ast"] })
extern class DeferStmt extends Stmt {
    var defer: Pos;
    var call: Pointer<CallExpr>;
}

@:go.Type({ name: "ast", instanceName: "ast.ReturnStmt", imports: ["go/ast"] })
extern class ReturnStmt extends Stmt {
    @:native("Return") var returnPos: Pos;
    var results: Slice<Expr>;
}

@:go.Type({ name: "ast", instanceName: "ast.BranchStmt", imports: ["go/ast"] })
extern class BranchStmt extends Stmt {
    var tokPos: Pos;
    var tok: Token;
    var label: Pointer<Ident>;
}

@:go.Type({ name: "ast", instanceName: "ast.BlockStmt", imports: ["go/ast"] })
extern class BlockStmt extends Stmt {
    var lbrace: Pos;
    var list: Slice<Stmt>;
    var rbrace: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.IfStmt", imports: ["go/ast"] })
extern class IfStmt extends Stmt {
    @:native("If") var ifPos: Pos;
    var init: Stmt;
    var cond: Expr;
    var body: Pointer<BlockStmt>;
    @:native("Else") var elseStmt: Stmt;
}

@:go.Type({ name: "ast", instanceName: "ast.CaseClause", imports: ["go/ast"] })
extern class CaseClause extends Stmt {
    @:native("Case") var casePos: Pos;
    var list: Slice<Expr>;
    var colon: Pos;
    var body: Slice<Stmt>;
}

@:go.Type({ name: "ast", instanceName: "ast.SwitchStmt", imports: ["go/ast"] })
extern class SwitchStmt extends Stmt {
    @:native("Switch") var switchPos: Pos;
    var init: Stmt;
    var tag: Expr;
    var body: Pointer<BlockStmt>;
}

@:go.Type({ name: "ast", instanceName: "ast.TypeSwitchStmt", imports: ["go/ast"] })
extern class TypeSwitchStmt extends Stmt {
    @:native("Switch") var switchPos: Pos;
    var init: Stmt;
    var assign: Stmt;
    var body: Pointer<BlockStmt>;
}

@:go.Type({ name: "ast", instanceName: "ast.CommClause", imports: ["go/ast"] })
extern class CommClause extends Stmt {
    @:native("Case") var casePos: Pos;
    var comm: Stmt;
    var colon: Pos;
    var body: Slice<Stmt>;
}

@:go.Type({ name: "ast", instanceName: "ast.SelectStmt", imports: ["go/ast"] })
extern class SelectStmt extends Stmt {
    var select: Pos;
    var body: Pointer<BlockStmt>;
}

@:go.Type({ name: "ast", instanceName: "ast.ForStmt", imports: ["go/ast"] })
extern class ForStmt extends Stmt {
    @:native("For") var forPos: Pos;
    var init: Stmt;
    var cond: Expr;
    var post: Stmt;
    var body: Pointer<BlockStmt>;
}

@:go.Type({ name: "ast", instanceName: "ast.RangeStmt", imports: ["go/ast"] })
extern class RangeStmt extends Stmt {
    @:native("For") var forPos: Pos;
    var key: Expr;
    var value: Expr;
    var tokPos: Pos;
    var tok: Token;
    var range: Pos;
    var x: Expr;
    var body: Pointer<BlockStmt>;
}

// ---------------------------------------------------------------------------
// Declarations
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.BadDecl", imports: ["go/ast"] })
extern class BadDecl extends Decl {
    var from: Pos;
    var to: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.GenDecl", imports: ["go/ast"] })
extern class GenDecl extends Decl {
    var doc: Pointer<CommentGroup>;
    var tokPos: Pos;
    var tok: Token;
    var lparen: Pos;
    var specs: Slice<Spec>;
    var rparen: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.FuncDecl", imports: ["go/ast"] })
extern class FuncDecl extends Decl {
    var doc: Pointer<CommentGroup>;
    var recv: Pointer<FieldList>;
    var name: Pointer<Ident>;
    var type: Pointer<FuncType>;
    var body: Pointer<BlockStmt>;
}

// ---------------------------------------------------------------------------
// Specs
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.ImportSpec", imports: ["go/ast"] })
extern class ImportSpec extends Spec {
    var doc: Pointer<CommentGroup>;
    var name: Pointer<Ident>;
    var path: Pointer<BasicLit>;
    var comment: Pointer<CommentGroup>;
    var endPos: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.ValueSpec", imports: ["go/ast"] })
extern class ValueSpec extends Spec {
    var doc: Pointer<CommentGroup>;
    var names: Slice<Pointer<Ident>>;
    var type: Expr;
    var values: Slice<Expr>;
    var comment: Pointer<CommentGroup>;
}

@:go.Type({ name: "ast", instanceName: "ast.TypeSpec", imports: ["go/ast"] })
extern class TypeSpec extends Spec {
    var doc: Pointer<CommentGroup>;
    var name: Pointer<Ident>;
    var typeParams: Pointer<FieldList>;
    var assign: Pos;
    var type: Expr;
    var comment: Pointer<CommentGroup>;
}

// ---------------------------------------------------------------------------
// Comments, fields and files
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.Comment", imports: ["go/ast"] })
extern class Comment extends Node {
    var slash: Pos;
    var text: String;
}

@:go.Type({ name: "ast", instanceName: "ast.CommentGroup", imports: ["go/ast"] })
extern class CommentGroup extends Node {
    var list: Slice<Pointer<Comment>>;

    function text(): String;
}

@:go.Type({ name: "ast", instanceName: "ast.Field", imports: ["go/ast"] })
extern class Field extends Node {
    var doc: Pointer<CommentGroup>;
    var names: Slice<Pointer<Ident>>;
    var type: Expr;
    var tag: Pointer<BasicLit>;
    var comment: Pointer<CommentGroup>;
}

@:go.Type({ name: "ast", instanceName: "ast.FieldList", imports: ["go/ast"] })
extern class FieldList extends Node {
    var opening: Pos;
    var list: Slice<Pointer<Field>>;
    var closing: Pos;

    function numFields(): Int;
}

@:go.Type({ name: "ast", instanceName: "ast.File", imports: ["go/ast"] })
extern class File extends Node {
    var doc: Pointer<CommentGroup>;
    @:native("Package") var packagePos: Pos;
    var name: Pointer<Ident>;
    var decls: Slice<Decl>;
    var fileStart: Pos;
    var fileEnd: Pos;
    var scope: Pointer<Scope>;
    var imports: Slice<Pointer<ImportSpec>>;
    var unresolved: Slice<Pointer<Ident>>;
    var comments: Slice<Pointer<CommentGroup>>;
    var goVersion: String;
}

@:go.Type({ name: "ast", instanceName: "ast.Package", imports: ["go/ast"] })
extern class Package extends Node {
    var name: String;
    var scope: Pointer<Scope>;
    var imports: Map<String, Pointer<Object>>;
    var files: Map<String, Pointer<File>>;
}

// ---------------------------------------------------------------------------
// Directives
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.Directive", imports: ["go/ast"] })
extern class Directive extends Node {
    var tool: String;
    var name: String;
    var args: String;
    var slash: Pos;
    var argsPos: Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.DirectiveArg", imports: ["go/ast"] })
extern class DirectiveArg {
    var arg: String;
    var pos: Pos;
}

// ---------------------------------------------------------------------------
// Objects and scopes
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.Object", imports: ["go/ast"] })
extern class Object {
    var kind: ObjKind;
    var name: String;
    var decl: Dynamic;
    var data: Dynamic;
    var type: Dynamic;

    function pos(): Pos;
}

@:go.Type({ name: "ast", instanceName: "ast.Scope", imports: ["go/ast"] })
extern class Scope {
    var outer: Pointer<Scope>;
    var objects: Map<String, Pointer<Object>>;

    function insert(obj: Pointer<Object>): Pointer<Object>;
    function lookup(name: String): Pointer<Object>;
    function string(): String;
}

// ---------------------------------------------------------------------------
// Named integer types and their package-level constants
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.ChanDir", imports: ["go/ast"] })
extern class ChanDir {
    @:native("SEND") static var send: ChanDir;
    @:native("RECV") static var recv: ChanDir;
}

@:go.Type({ name: "ast", instanceName: "ast.ObjKind", imports: ["go/ast"] })
extern class ObjKind {
    static var bad: ObjKind;
    static var pkg: ObjKind;
    static var con: ObjKind;
    static var typ: ObjKind;
    @:native("Var") static var varKind: ObjKind;
    static var fun: ObjKind;
    static var lbl: ObjKind;

    function string(): String;
}

@:go.Type({ name: "ast", instanceName: "ast.MergeMode", imports: ["go/ast"] })
extern class MergeMode {
    static var filterFuncDuplicates: MergeMode;
    static var filterUnassociatedComments: MergeMode;
    static var filterImportDuplicates: MergeMode;
}

// ---------------------------------------------------------------------------
// Comment map
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", instanceName: "ast.CommentMap", imports: ["go/ast"] })
extern class CommentMap {
    function comments(): Slice<Pointer<CommentGroup>>;
    function filter(node: Node): CommentMap;
    function string(): String;
    function update(oldNode: Node, newNode: Node): Node;
}

// ---------------------------------------------------------------------------
// Package-level functions and constructors
// ---------------------------------------------------------------------------

@:go.Type({ name: "ast", imports: ["go/ast"] })
extern class Ast {
    static function newIdent(name: String): Pointer<Ident>;
    static function newObj(kind: ObjKind, name: String): Pointer<Object>;
    static function newScope(outer: Pointer<Scope>): Pointer<Scope>;
    static function newCommentMap(fset: Pointer<FileSet>, node: Node, comments: Slice<Pointer<CommentGroup>>): CommentMap;

    static function fileExports(src: Pointer<File>): Bool;
    static function packageExports(pkg: Pointer<Package>): Bool;
    static function filterDecl(decl: Decl, f: String -> Bool): Bool;
    static function filterFile(src: Pointer<File>, f: String -> Bool): Bool;
    static function filterPackage(pkg: Pointer<Package>, f: String -> Bool): Bool;
    static function inspect(node: Node, f: Node -> Bool): Void;
    static function isExported(name: String): Bool;
    static function isGenerated(file: Pointer<File>): Bool;
    static function walk(v: Visitor, node: Node): Void;
    static function unparen(e: Expr): Expr;
    static function sortImports(fset: Pointer<FileSet>, f: Pointer<File>): Void;
    static function mergePackageFiles(pkg: Pointer<Package>, mode: MergeMode): Pointer<File>;
    static function print(fset: Pointer<FileSet>, x: Dynamic): Error;
    static function preorderStack(root: Node, stack: Slice<Node>, f: (Node, Slice<Node>) -> Bool): Void;
}
