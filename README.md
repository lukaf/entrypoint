```bash
$ go build entrypoint.go
$ docker build -t myetcd .
[...]

$ docker run --rm -t -i myetcd -addr %ip%:4001 -peer-addr %ip%:7001
[etcd] Mar 26 20:29:50.281 WARNING   | Using the directory 004b0f9dc8b0.etcd as the etcd curation directory because a directory was not specified. 
[etcd] Mar 26 20:29:50.282 INFO      | 004b0f9dc8b0 is starting a new cluster
[etcd] Mar 26 20:29:50.285 INFO      | etcd server [name 004b0f9dc8b0, listen on :4001, advertised url http://172.17.0.12:4001]
[etcd] Mar 26 20:29:50.285 INFO      | peer server [name 004b0f9dc8b0, listen on :7001, advertised url http://172.17.0.12:7001]
[etcd] Mar 26 20:29:50.285 INFO      | 004b0f9dc8b0 starting in peer mode
[etcd] Mar 26 20:29:50.285 INFO      | 004b0f9dc8b0: state changed from 'initialized' to 'follower'.
[etcd] Mar 26 20:29:50.285 INFO      | 004b0f9dc8b0: state changed from 'follower' to 'leader'.
[etcd] Mar 26 20:29:50.285 INFO      | 004b0f9dc8b0: leader changed from '' to '004b0f9dc8b0'.
```
