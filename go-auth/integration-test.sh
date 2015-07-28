#!/bin/sh

function quit {
	docker-compose stop
	docker-compose rm -f
    exit $1
}

docker-compose up -d
# make sure database is rady
sleep 10
# makes sure all containers are started
docker-compose start

sleep 3
docker-compose start


first=$(curl -i -silent -X PUT -d userid=USERNAME -d password=PASSWORD $(boot2docker ip):8080/user | grep "HTTP/1.1")
second=$(curl -i -silent -X PUT -d userid=USERNAME -d password=PASSWORD $(boot2docker ip):8080/user | grep "HTTP/1.1")

status_first=$(echo "$first" | cut -f 2 -d ' ')
status_second=$(echo "$second" | cut -f 2 -d ' ')

if [[ "$status_first" -ne 200 ]]; then
	echo "Expecting 200 OK for first user register"
	quit 1
else
	echo "Pass Register User"
fi


if [[ "$status_second" -ne 409 ]]; then
	echo "Expecting 409 OK for second user register"
	quit 1
else 
	echo "Pass Register User Conflict"
fi

quit 0


