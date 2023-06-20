FROM amazonlinux:latest

RUN yum -y update
RUN yum -y install httpd
RUN yum -y install php

COPY httpd.conf /etc/httpd/conf/httpd.conf
COPY index.php /var/www/html/index.php

CMD ["/usr/sbin/httpd","-D","FOREGROUND"]

EXPOSE 80
