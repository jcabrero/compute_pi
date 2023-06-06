# Calculate $pi$

The code for computing $pi$ can be found in `internal/pi/pi.go`.

## How to run:
To execute run:

```
docker run --rm -it -v $PWD:/usr/src/myapp -w /usr/src/myapp golang:1.20
make build
./bin/pi
```

