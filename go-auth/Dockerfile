FROM ubuntu

ADD ./go-auth /bin/go-auth

EXPOSE 9000

ADD run-go-auth.sh /bin/start-go-auth.sh

ENTRYPOINT ["/bin/start-go-auth.sh"]