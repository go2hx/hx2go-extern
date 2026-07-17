class Sanitize {
    static var sanitizeList = [
        "abstract",
            "break",
            "case",
            "cast",
            "catch",
            "class",
            "continue",
            "default",
            "do",
            "else",
            "enum", 
            "extends",
            "extern",
            "false",
            "final",
            "for",
            "function",
            "if",
            "implements",
            "import",
            "in",
            "interface",
            "inline",
            "is",
            "macro",
            "new",
            "null",
            "override",
            "package",
            "private",
            "public",
            "return",
            "static",
            "super",
            "switch",
            "this",
            "throw",
            "true",
            "try",
            "typedef",
            "untyped",
            "using",
            "var",
            "while",
            "from",
            "to",
            "dynamic",
    ];

    public static function name(name: String): String {
        return if (sanitizeList.contains(name)) {
            "_" + name;
        }else{
            name;
        }
    }

    public static function segment(seg: String): String {
        seg = seg.split(".").join("_");
        seg = seg.split("-").join("_");
        return seg;
    }

    public static function packagePath(path: String): String {
        return path.split("/").map(Sanitize.segment).join(".");
    }

    public static function packageDir(path: String): String {
        return path.split("/").map(Sanitize.segment).join("/");
    }
}