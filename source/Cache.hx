import haxe.crypto.Md5;
import sys.io.File;
import go.Pointer;
import go.golang_org.x.tools.go.packages.Package;

class Cache {
    public static function getPackageCheckSum(entry:Pointer<Package>):String {
        var checksums = [for (file in entry.goFiles) {
            Md5.encode(File.getContent(file));
        }];
        checksums.sort((a, b) -> {
            return a > b ? -1 : 1;
        });
        return Md5.encode(checksums.join("$|"));
    }
}