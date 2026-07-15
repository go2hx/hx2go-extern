import go.net.http.ResponseWriter;
import go.net.http.Request;
import go.net.http.Http;
import go.encoding.json.Json;
import go.bytes.Bytes;
import go.Pointer;
import go.Map;

function handler(w: ResponseWriter, req: Pointer<Request>): Void {
    w.write("Hello, World!");
}

function httpServer() {
    Http.handleFunc("/", handler);
    Http.listenAndServe(":8080", null);
}

function jsonDecoder() {
    var v = '{ "hello": "world", "num": 123, "boolean": true, "nest": { "a": 1, "b": 2 } }';
    var dec = Json.newDecoder(
        Bytes.newReader(v)
    );

    var data = new Map<String, Dynamic>();
    dec.decode(Pointer.addressOf(data)).sure();

    trace(data);
}

function main() {
    jsonDecoder();
}