import haxe.crypto.Sha256;
import sys.io.File;
import go.Pointer;
import std.go.packages.Packages;

class Cache {
    public static function getPackageCheckSum(entry:Pointer<Package>):String {
        var checksums = [for (file in entry.goFiles) {
            Sha256.encode(File.getContent(file));
        }];
        checksums.sort((a, b) -> {
            return a > b ? -1 : 1;
        });
        return Sha256.encode(checksums.join("$|"));
    }
}