# Preparación de la base de datos con docker

MySQL:`docker run --name mysql5 --hostname mysql5 -v C:\Users\MaQuiNa1995\workspace\docker\mysql-go:/var/lib/mysql --network bridge -e MYSQL_ROOT_PASSWORD=pass -d -p 3306:3306 mysql:5`

PHPMyAdmin: `docker run --name phpmyadmin --hostname phpmyadmin -d --network bridge -p 8080:80 -e PMA_HOSTS=mysql5 --link mysql5:db -e PMA_VERBOSES="Interna" phpmyadmin/phpmyadmin:latest`

