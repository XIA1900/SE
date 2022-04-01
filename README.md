# Gator Forum Backend

## Developer

- Bowei Wu(bowei.wu@ufl.edu)
- Yingjie Chen(yingjie.chen@ufl.edu)

## Components

- **gin** for web server
- **gorm** for database operation
  - MySQL
- **Redis** for cache
- uber/**zap** for log
- **JWT** for user authentication
- configuration information stored in yaml file
- **Casbin** for role management
- **Elasticsearch** for forum user searching
- **gin-swagger** for api information
- **wire** for dependency injection

## File Structure

- Directory `GFBackend` is this project backend, when some features or functions have been accomplished, we will merge it to main branch




## Docker

### Initial Deployment

- Database Deployment Steps
  - Download `MySQL`, `Redis`, `ElasticSearch` images;

    ```shell
    docker pull mysql
    docker pull redis
    docker pull elasticsearch:7.16.2
    ```

  - create a folders

    ```shell
    mkdir docker_volumes
    cd docker_volumes
    mkdir es mysql redis
    ```

    - edit `mysql` folder

      ```shell
      cd mysql
      mkdir conf data log
      cd conf
      touch my.cnf
      ```

      ```shell
      # edition of my.cnf
      default_character_set=utf8
      [mysqld]
      collation_server=utf8_general_ci
      character_set_server=utf8
      ```
  
    - edit `redis` folder
  
      ```shell
      cd redis
      touch redis.conf
      ```
  
      ```shell
      # edition of redis.conf
      protected-mode yes
      port 6379
      tcp-backlog 511
      timeout 0
      tcp-keepalive 300
      daemonize no
      pidfile /var/run/redis_6379.pid
      loglevel notice
      logfile ""
      databases 16
      always-show-logo no
      set-proc-title yes
      proc-title-template "{title} {listen-addr} {server-mode}"
      stop-writes-on-bgsave-error yes
      rdbcompression yes
      rdbchecksum yes
      dbfilename dump.rdb
      rdb-del-sync-files no
      dir ./
      replica-serve-stale-data yes
      replica-read-only yes
      repl-diskless-sync no
      repl-diskless-sync-delay 5
      repl-diskless-load disabled
      repl-disable-tcp-nodelay no
      replica-priority 100
      acllog-max-len 128
      requirepass Redis6!
      lazyfree-lazy-eviction no
      lazyfree-lazy-expire no
      lazyfree-lazy-server-del no
      replica-lazy-flush no
      lazyfree-lazy-user-del no
      lazyfree-lazy-user-flush no
      oom-score-adj no
      oom-score-adj-values 0 200 800
      disable-thp yes
      appendonly no
      appendfilename "appendonly.aof"
      appendfsync everysec
      no-appendfsync-on-rewrite no
      auto-aof-rewrite-percentage 100
      auto-aof-rewrite-min-size 64mb
      aof-load-truncated yes
      aof-use-rdb-preamble yes
      lua-time-limit 5000
      slowlog-log-slower-than 10000
      slowlog-max-len 128
      latency-monitor-threshold 0
      notify-keyspace-events ""
      hash-max-ziplist-entries 512
      hash-max-ziplist-value 64
      list-max-ziplist-size -2
      list-compress-depth 0
      set-max-intset-entries 512
      zset-max-ziplist-entries 128
      zset-max-ziplist-value 64
      hll-sparse-max-bytes 3000
      stream-node-max-bytes 4096
      stream-node-max-entries 100
      activerehashing yes
      client-output-buffer-limit normal 0 0 0
      client-output-buffer-limit replica 256mb 64mb 60
      client-output-buffer-limit pubsub 32mb 8mb 60
      hz 10
      dynamic-hz yes
      aof-rewrite-incremental-fsync yes
      rdb-save-incremental-fsync yes
      jemalloc-bg-thread yes
      ```
  
    - edit `es` folder
  
      ```shell
      mkdir backup conf data logs plugins
      ```
  
  - docker start script
  
    - start `MySQL` docker container
  
      ```shell
      touch mysql_start.sh
      ```
  
      ```shell
      echo "docker start MySQL8"
      docker run -d \
      -p 3306:3306 \
      --privileged=true \
      -v /root/docker_volumns/mysql/log:/var/log/mysql \
      -v /root/docker_volumns/mysql/data:/var/lib/mysql \
      -v /root/docker_volumns/mysql/conf:/etc/mysql/conf.d \
      -e MYSQL_ROOT_PASSWORD=MySQL8.0! \
      --name mysql 826efd84393b
      ```
    - start `Redis` docker container
  
      ```shell
      touch redis_start.sh
      ```
  
      ```shell
      echo "docker start redis"
      docker run -d \
      -p 6379:6379 \
      --privileged=true \
      -v /root/docker_volumns/redis/redis.conf:/etc/redis/redis.conf \
      -v /root/docker_volumns/redis/data:/data \
      --name redis 0e403e3816e8 \
      redis-server /etc/redis/redis.conf
      ```
      
    - start `ElasticSearch` docker container
    
      ```shell
      touch es_start.sh
      ```
    
      ```shell
      docker run -d \
      --privileged=true \
      -v /root/docker_volumns/es/data:/usr/share/elasticsearch/data \
      -v /root/docker_volumns/es/conf/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml \
      -v /root/docker_volumns/es/logs:/user/share/elasticsearch/logs \
      -v /root/docker_volumns/es/backup:/usr/share/elasticsearch/backup \
      -v /root/docker_volumns/es/plugins:/usr/share/elasticsearch/plugins \
      -p 9200:9200 -p 9300:9300 \
      -e ES_JAVA_OPTS="-Xms512m -Xmx512m" \
      -e "discovery.type=single-node" \
      --name es 66c29cde15ce
      ```
      
    
    

