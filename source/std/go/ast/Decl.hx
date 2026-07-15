package std.go.ast;

import std.go.ast.Ast.Node;


@:go.Type({ name: "ArrayType", instanceName: "ast.Node", imports: ["go/ast"] })
extern class Decl extends Node {}