FROM envoyproxy/envoy:v1.17.0
# Use config file from docker-compose deployment
COPY docker/docker-compose/envoy-config.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log
