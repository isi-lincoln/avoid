# Building coredns

```
git submodule update --init --recursive
```

The main code fore coredns is in our avoid plugin located in `coredns/plugin/avoid/avoid.go`.  Reference material for how to build a plugin is [here](https://coredns.io/2017/03/01/how-to-add-plugins-to-coredns/).

## Updating avoid plugin

First thing is that in the top level avoid directory we want to make sure we have a `git tag`.  You could use the git commit, but its not a resilient as git tags for long term deployments.  Once our git tag has been created and pushed.  We go back into coredns:

```
cd coredns; go get github.com/isi-lincoln/avoid@v0.0.4; make; cd ..
```

In place of the tag `v0.0.4` use whichever tag has been pushed.

## Updating coredns dockerfile

After we've updated our code and made it, we can use the modified makefile to create a new docker image.

```
cd coredns; sudo make docker; cd ..
```

One can use `ENV` variables such as `REGISTRY`, `REPO`, `TAG`, and `PUSH` to modify the image name and if the container should be pushed to that remote registry.

## Avoid plugin parameters

The avoid coredns plugin takes 2 optional parameters

```
example.com:53 {
    avoid [hostname] [port]
}
```

Here the `hostname` and `port` represent the hostname and port of the avoid-dns-service container.  Our avoid plugin uses the avoid protobuf protocol to reach the dns-service to access the etcd values.

## Testing dns

```
dig +nocmd +noall +answer +ttlid a gw0.example.com @localhost
```


# Using the avoid-dns-cli

```
avoid-dns-cli update cli 127.0.0.1 gw0.example.com 0 DOCKERTEST --a 10.0.0.1 --aaaa aaaa::0001
```
