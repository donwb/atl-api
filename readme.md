# ATl-API
Demo app for Atlanta Code Camp, used in both the REDIS and Goji sessions

## To Run
Use the default make target `all` by typing `make all` in the source folder.  This will start the web server.  You will also need REDIS running locally on the default port.

_**Disclaimer**_
There's not much error handling in this code.  This is purposeful to keep the code that's core to REDIS and Goji easier to read if you don't already know Golang.

### Protobuf generator:
protoc --go_out=. user.proto
