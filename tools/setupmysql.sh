#!/bin/bash

mysql -h 127.0.0.1 -P 3306 -u root -p -e "CREATE USER 'ethadmin'@'%' IDENTIFIED BY '123';"
mysql -h 127.0.0.1 -P 3306 -u root -p -e "GRANT ALL PRIVILEGES ON emsvc.* TO ethadmin@'%';"