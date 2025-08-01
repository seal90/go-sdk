# start rocketmq
```shell
# start rocketmq nameserver
docker run -d --name rmqnamesrv --net host apache/rocketmq:5.3.3 sh mqnamesrv

# check NameServer start success or not
docker logs -f rmqnamesrv

# config Broker IP
echo "brokerIP1=127.0.0.1" > broker.conf

# start Broker and Proxy
docker run -d \
--name rmqbroker \
--net host         \
-e "NAMESRV_ADDR=127.0.0.1:9876" \
-v ./broker.conf:/home/rocketmq/rocketmq-5.3.3/conf/broker.conf \
apache/rocketmq:5.3.3 sh mqbroker --enable-proxy \
-c /home/rocketmq/rocketmq-5.3.3/conf/broker.conf

# check Broker start success or not
docker exec -it rmqbroker bash -c "tail -n 10 /home/rocketmq/logs/rocketmqlogs/proxy.log"
# When see 'The broker boot success..' will success

# create topic
docker exec -it rmqbroker bash
sh mqadmin updatetopic -t virtualnamespace -n 127.0.0.1:9876 -c DefaultCluster

# create group
sh mqadmin updateSubGroup -g virtualnamespace_g  -n 127.0.0.1:9876  -c DefaultCluster
sh mqadmin updateSubGroup -g virtualnamespace_g2  -n 127.0.0.1:9876  -c DefaultCluster

# dashboard
docker pull apacherocketmq/rocketmq-dashboard:2.0.1

docker run -d --name rocketmq-dashboard --net=host  -e "JAVA_OPTS=-Drocketmq.namesrv.addr=127.0.0.1:9876 -Dserver.port=10000" -t apacherocketmq/rocketmq-dashboard:2.0.1
```

# cmd

```shell
../dapr/cmd/daprd/debug_bin run --app-id sub          --app-protocol http          --app-port 8085          --dapr-http-port 3500          --log-level debug          --resources-path ./examples/virtual-namespace/pubsub/config/ --config ~/dapr_dev/config.yaml
```

# use same consumerGroup, all virtual_namesapce IS NOT NULL, consume message by command
# use different consumerGroup, main namespace virtual_namesapce IS NOT NULL, the other virtual_namesapce IS NOT NULL AND virtual_namesapce = 'VIRTUAL_NAMESPACE'

# TODO
* message with rocketmq key `virtual_namespace=NAMESPACE` for search

