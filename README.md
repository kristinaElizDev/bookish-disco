# bookish-disco
Azure AMR Testing - for go! 


- Initialize your repo: go mod init github.com/my/repo

- Add go-redis: go get github.com/redis/go-redis/v9
- go get github.com/redis/go-redis-entraid
- go get github.com/github/go-redis/crdb

Go Proxy: 
- Updates to devcontainer to use go proxy and pull 
    - Resources: 
        https://github.com/github/goproxy/blob/main/doc/user.md
        https://github.com/github/features/tree/main/src/goproxy
        https://github.com/github/goproxy/blob/main/doc/user.md
- if you get auth errors with this, replace your PAT in netrc


https://redis.io/docs/latest/develop/clients/go/amr/