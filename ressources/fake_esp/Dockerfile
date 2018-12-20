FROM node:8.11.4

ENV http_proxy=http://PITC-Zscaler-EMEA-Amsterdam3PR.proxy.corporate.ge.com:80

ENV https_proxy=http://PITC-Zscaler-EMEA-Amsterdam3PR.proxy.corporate.ge.com:80

ADD ./app /app

RUN mkdir /data

ADD ./iot_passport_example.json /data

WORKDIR /app

CMD npm install && npm start