# img-proxy-server
An image proxy server using Golang and support sock5

## Browser Usage:
origin img url: `http://xx.com/1.jpg`
#### using server
proxy img url: `http://0.0.0.0:10880/img?url=http://xx.com/1.jpg`
#### using sock5 proxy on the server
proxy with sock5 img url: `http://0.0.0.0:10880/img_proxy?url=http://xx.com/1.jpg`

## Server Usage:
```
./build_for_linux_amd64_linux -h

Usage of ./build_for_linux_amd64_linux:
  -port string
    	http listen port
  -proxy string
    	sock5 proxy add:port
    	
Example:
./build_for_mac -port=":80"  -proxy="0.0.0.0:1086"
```

## Run without args:
```
You can not use proxy without address..
Using random port..

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.0.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:33775

```

## Source Code

[main.go](main/main.go)