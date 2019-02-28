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

__request_header__  

request header consists:
- info of request
- info of client

if request have a message, Content-Length and Transfer-Encoding is need

common http request headers:
- Accept: Content types that are acceptable by the client as part of the HTTP response.
- Accept-Charset: The character sets required from the server.
- Authorization: This is used to send Basic Authentication credentials to the server.
- Cookie: The client should send back cookies that were set by the calling server.
- Content-Length: The length of the request body in octets.
- Content-Type: The content type of the request body.
- Host: The name of the server, along with the port number.
- Referrer: The address of the previous page that linked to the requested page.
- User-Agent: Describes the calling client.

### response

a response consists of a few line:
- a status line
- zero ro more response headers
- an empty line
- the message body (optional)

status line = status code + reason phrase

__status_code__  
- 1xx: Informational. This tells the client that the server has already received the request and is processing it.
- 2xx: Success. This is what clients want; the server has received the request and has processed it successfully. The standard response in this class is 200 OK.
- 3xx: Redirection. This tells the client that the request is received and processed but the client needs to do more to complete the action.
- 4xx: Client Error. This tells the client that there’s something wrong with the request.
- 5xx: Server Error. This tells the client that there’s something wrong with the request but it’s the server’s fault.

__response_header__  

common http response headers:
- Allow: Tells the client which request methods are supported by the server.
- Content-Length: The length of the response body in octets
- Content-Type: The content type of the response body
- Date: Tells the current time
- Location: This header is used with redirection, to tell the client where to request the next URL.
- Server: Domain name of the server that’s returning the response.
- Set-Cookie: Sets a cookie at the client.
- WWW-Authenticate: Tells header the client what type of authorization clients should supply in their Authorization request header.

---
- rui: name of resource
- rul: location of resource

RUI form:  
    < scheme name> : < hierarchical part> [ ? < query> ][ # < fragment> ]"

http://user:passwd@www.baidu.com/doc/file?name=test&ip=123#sum  

## http2

- focuse on performance
- based on spdy/2
- binary protocol (http1.x base on text)
- full multiplexed
- compress the header
- allow server to push response to client

in go1.6 and after, http2 is default

## define web app

- get a http request form client
- process the requet, do some work
- generate html and return it in an http response message

web app = handlers + template engine

### handlers

- receives and process the http request
- call the template engine to generate the html and something about response 

    mvc(model-view-controller pattern)
    divide a arogram into three parts: model, view, controller
    model - underlaying data
    view - visualization of the model for the user
    controller - use input(from user) to modify the model
    model change --> view update automatically

mvc is a good way for web app, but is not the only way

### template engine

come form ssi technology

there are two type of template with different design philosophies:
- static template or logic-less template
    use placeholder tokens
    no logic
- active template
    placehoder tokens + other programming language 
    eg: jsp asp erb

- [论坛设计](/chitchat/README.md)
