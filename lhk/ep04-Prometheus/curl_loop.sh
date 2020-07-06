#!/bin/bash

IP_1=161.35.254.131 
IP_2=161.35.254.114  
counter=1
while [ $counter -le 360 ]
    do
        echo "request $counter";
        curl -s "http://$IP_1" && echo " app1 ";
        curl -s "http://$IP_2" && echo " app2 ";
        ((counter++));
        sleep 2;
done