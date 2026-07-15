import go.net.http.ResponseWriter;
import go.net.http.Request;
import go.Pointer;
import go.net.http.Http;
import go.encoding.json.Json;
import go.bytes.Bytes;
import go.Map;

function handler(w: ResponseWriter, req: Pointer<Request>): Void {
    w.write(cast "Hello, World!");
}

function httpServer() {
    Http.handleFunc("/", handler);
    Http.listenAndServe(":8080", null);
}

function jsonDecoder() {
    var v = '{ "hello": "world", "num": 123, "boolean": true, "nest": { "a": 1, "b": 2 } }';
    var dec = Json.newDecoder(
        cast Bytes.newReader(v)
    );

    var data = new Map<String, Dynamic>();
    dec.value.decode(Pointer.addressOf(data)).sure();

    trace(data);
}

function main() {
    jsonDecoder();
}