#this is not a good practice to dockerize our go application its done only for testing purposes
FROM alpine  

COPY ./appxposer-controller /usr/local/bin/appxposer-controller 

ENTRYPOINT [ "/usr/local/bin/appxposer-controller" ]