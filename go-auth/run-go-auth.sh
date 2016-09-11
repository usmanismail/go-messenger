#!/bin/sh


# Add Entry for mysql to etc hosts
echo "$GO_AUTH_MYSQL_SERVICE_HOST mysql" >> /etc/hosts

# Create the database
echo "create database messenger" | mysql -u root -p${MYSQL_ROOT_PASSWORD} -h mysql || echo "Already exists"

ln -s /tmp/log /dev/log
/bin/go-auth $@
