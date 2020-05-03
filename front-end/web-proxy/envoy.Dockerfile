FROM envoyproxy/envoy:v1.14.1

ARG SVC_ADDR
ARG SVC_PORT

RUN apt-get update \
    && apt-get install -y gettext 

ENV SVC_ADDR=${SVC_ADDR}
ENV SVC_PORT=${SVC_PORT}

COPY ./envoy.yaml.tmpl /tmp/envoy/envoy.yaml.tmpl
RUN envsubst < /tmp/envoy/envoy.yaml.tmpl > /etc/envoy/envoy.yaml \
    # cleanup
    && rm -rf /tmp/envoy

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml