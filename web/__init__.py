import http.server
import socketserver

PORT = 8123


def start():
    Handler = http.server.SimpleHTTPRequestHandler

    httpd = socketserver.TCPServer(("", PORT), Handler)

    print("serving at port", PORT)
    httpd.serve_forever()

if __name__ == "__main__":
    start()