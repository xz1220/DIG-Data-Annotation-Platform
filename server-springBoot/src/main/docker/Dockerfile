FROM adoptopenjdk/openjdk11:ubuntu

VOLUME /tmp

ADD *.jar app.jar

EXPOSE 80

ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/app.jar"]

# Ubuntu 时区
#RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime