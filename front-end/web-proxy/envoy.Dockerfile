FROM envoyproxy/envoy:v1.14.1

COPY ./envoy.yaml /etc/envoy/envoy.yaml.tmpl
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml