# How  to run code ?

```go
go run main.go ip --host domain.com
go run main.go servers --host domain.com
```
___

## Google Example

```go
go run main.go ip --host google.com
```

OUTPUT:

```go
142.251.129.78
2800:3f0:4001:81d::200e
```

```go
go run main.go servers --host google.com
```

OUTPUT:

```go
&{ns2.google.com.}
&{ns3.google.com.}
&{ns1.google.com.}
&{ns4.google.com.}
```