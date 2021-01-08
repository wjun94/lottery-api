#!bin/bash

echo "第1步"
docker rmi -f $(docker images | grep "none" | awk '{print $3}')
docker build -t buildgo:0.5 . 
echo "第2步"
docker run -d --name redis1 -p 6379:6379 redis --requirepass "19941121"
echo "第3步"
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=19941121 -d mysql
#!bin/bash
echo "第4步"
sleep 6s
docker run -dit --link mysql:mysql --link redis1:redis -p 8080:8080 buildgo:0.5