.PHONY:	client server
all:	client server
client:
	cd client && go install ${LDFLAGS} template-client.go `ls -1 *.go|grep -v template-client.go`
server:
	cd server && go install ${LDFLAGS} template-server.go `ls -1 *.go|grep -v template-server.go`
