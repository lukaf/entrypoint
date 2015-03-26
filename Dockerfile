FROM coreos/etcd

COPY entrypoint /entrypoint

ENTRYPOINT ["/entrypoint", "/etcd"]
