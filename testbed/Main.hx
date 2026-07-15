import go.net.http.ResponseWriter;
import go.net.http.Request;
import go.Pointer;
import go.net.http.Http;
import go.fmt.Fmt;

function handler(w: ResponseWriter, req: Pointer<Request>): Void {
    w.write(cast "Hello, World!");
}

function main() {
    Http.handleFunc("/", handler);
    Http.listenAndServe(":8080", null);
}