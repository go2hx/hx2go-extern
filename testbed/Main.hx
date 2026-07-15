import go2hx.net.http.ResponseWriter;
import go2hx.net.http.Request;
import go2hx.net.Http;
import go2hx.encoding.Json;
import go2hx.Bytes;
import go.Pointer;
import go.Map;
import go.os.Os;
import go2hx.Image;
import go2hx.bytes.Buffer;
import go2hx.image.Png;
import go2hx.image.color.RGBA;
import go.Syntax;
import go2hx.Net;
import go2hx.net.Conn;
import go.Slice;
import go.Byte;
import go.Go;
import go2hx.os.exec.Cmd;
import go2hx.os.Exec;

function handler(w: ResponseWriter, req: Pointer<Request>): Void {
    w.write(Std.string(req.userAgent()));
}

function httpTest() {
    Http.handleFunc("/", handler);
    Http.listenAndServe(":8080", null);
}

function jsonTest() {
    var v = '{ "hello": "world", "num": 123, "boolean": true, "nest": { "a": 1, "b": 2 } }';
    var dec = Json.newDecoder(
        Bytes.newReader(v)
    );

    var data = new Map<String, Dynamic>();
    dec.decode(Pointer.addressOf(data)).sure();

    trace(data);
}

function imageTest() {
    var width = 512;
    var height = 512;
    var img = Image.newRGBA(Image.rect(0, 0, width, height));

    for (x in 0...width) {
        var t = x / width - 1;
        var r = Std.int(255 * (1 - t));
        var g = Std.int(255 * t);

        for (y in 0...height) {
            img.set(x, y, new RGBA(r, g, 0, 255));
        }
    }

    var buf: Buffer = null;
    Png.encode(Pointer.addressOf(buf), img).sure();

    Os.writeFile("./output.png", buf.bytes(), Syntax.code("0664"));
}

function handleConn(conn: Conn) {
    var buf = new Slice<Byte>(1024);
    while (true) {
        var result = conn.read(buf);
        if (!result.isOk()) {
            conn.close();
            break;
        }

        conn.write(buf.slice(0, result.sure()));
    }
}

function tcpServer() {
    var ln = Net.listen("tcp", ":9000").sure();
    while (true) {
        var conn = ln.accept().sure();
        Syntax.go(() -> handleConn(conn));
    }
}

function tcpClient() {
    var conn = Net.dial("tcp", "localhost:9000").sure();
    conn.write("hello go2hx");

    var buf = new Slice<Byte>(1024);
    var result = conn.read(buf);
    trace(Go.string(buf.slice(0, result.sure())));

    conn.close();
}


function tcpTest() {
    Syntax.go(tcpServer);
    Sys.sleep(4);
    tcpClient();
}

function commandTest() {
    var cmd: Pointer<Cmd> = Exec.command("echo", "hello", "from", "go2hx");
    var out = Go.string(cmd.value.output().sure());
    trace(out);
}

function main() {
//    jsonTest();
//    httpTest();
//    imageTest();
//    tcpTest();
//    commandTest();
}