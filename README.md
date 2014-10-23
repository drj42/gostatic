Static File Server
==================

This is a bare bones simple static file server, written in Go.  Its about as
basic as it can get. It logs all requests to STDOUT.  I use this as a Python
SimpleHTTPServer replacement.

Install
=======

Recommended to have $GOPATH/bin in your $PATH

    $ go get github.com/drj42/gostatic

Example Usage
=======

Serve $HOME on 192.168.0.10:8000

    $ gostatic -port=8000 -host=192.168.0.10 -root=$HOME 

Serve current dir on 127.0.0.1:8080

    $ gostatic

