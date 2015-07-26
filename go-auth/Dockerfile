FROM ubuntu

ADD ./go-auth /bin/go-auth

EXPOSE 8080

ENTRYPOINT ["/bin/go-auth","-l","debug","run","-p","8080"]


