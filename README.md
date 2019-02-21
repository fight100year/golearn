# go

- simple
- efficient
- back end system

for web app and *-as-a-server system

large-scale web app features:
- scalable
  - vertical scaling: increasing cpu
    goroutine
  - horizontal scaling: increasing machines
    layer: add a proxy layer
- modular
  easy to add/remove/modify feature
  easy to reusing modular components
- maintainable
- high performance


## web app and web server

app: software program

web app: 
* return html,(client render html and display to user)
* transport data by http

web service:
* not return html
* server for other program

## http

app-level communication protocol

* stateless
* text-based
* request-response
* c/s

### request

a request consists of a few line:
- request-line
- zero or more request headers
- an empty line
- the message body (optional)

GET /Protocols/rfc2616/rfc2616.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0
(empty line)
abc

request-line = request method + uri + http version
request headers = a:b c:d

__request method__  
- get: Tells the server to return the specified resource.
- head: The same as GET except that the server must not return a message body.
- post: Tells the server that the data in the message body should be passed to the resource identified by the URI .
- put: Tells the server that the data in the message body should be the resource at the given URI . 
- delete: Tells the server to remove the resource identified by the URI .
- trace: Tells the server to return the request
- options: Tells the server to return a list of HTTP methods that the server supports.
- connect: Tells the server to set up a network connection with the client.
- patch: Tells the server that the data in the message body modifies the resource identified by the URI .

__idempotent__
the result of 1 call = the result of 1000 times call


