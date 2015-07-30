#!/bin/bash

function quit {
	docker-compose stop
	docker-compose rm -f
	exit $1
}

set -x

docker-compose up -d
# make sure database is ready
docker-compose start Database

sleep 20
# makes sure all containers are started
docker-compose start Goauth

sleep 10

service_container=$(docker ps -a | awk '{ print $1,$2 }' | grep go-auth | awk '{print $1 }')

echo $service_container

service_ip=$(docker inspect -f '{{ .NetworkSettings.IPAddress }}' ${service_container})

first=$(curl -i -silent -X PUT -d userid=USERNAME -d password=PASSWORD ${service_ip}:9000/user | grep "HTTP/1.1")
second=$(curl -i -silent -X PUT -d userid=USERNAME -d password=PASSWORD ${service_ip}:9000/user | grep "HTTP/1.1")

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


