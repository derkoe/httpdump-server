# httpdump-server

Simple HTTP server that dumps the request to stdout. Useful to debug incoming requests :)

Run locally:

    docker run -p 8080:8080 derkoe/httpdump-server

Example output (when hitting it with Chrome):

    GET / HTTP/1.1
    Connection: [keep-alive]
    User-Agent: [Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36]
    Accept: [text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8]
    Accept-Encoding: [gzip, deflate, br]
    Cache-Control: [max-age=0]
    Accept-Language: [en-US;q=0.8,en;q=0.7]
    ================
    GET /favicon.ico HTTP/1.1
    Connection: [keep-alive]
    User-Agent: [Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36]
    Referer: [http://localhost:8080/ ]
    Accept: [image/webp,image/apng,image/*,*/*;q=0.8]
    Accept-Encoding: [gzip, deflate, br]
    Accept-Language: [en-US;q=0.8,en;q=0.7]
    ================
