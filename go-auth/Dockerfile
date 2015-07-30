FROM ubuntu

ADD ./go-auth /bin/go-auth

EXPOSE 9000

ENTRYPOINT ["/bin/go-auth","-l","debug","run","-p","9000"]


