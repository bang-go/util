# util

Small, dependency-light helpers for Go services.

## Design

- Keep APIs explicit. Panic and error forms are named differently.
- Preserve caller types in generic helpers whenever possible.
- Prefer tiny focused utilities over catch-all abstractions.

## Packages

- `util`: generic helpers for pointers, retries, files, strings, numbers, time, and UUIDs.
- `cipher`: SHA-256 hashing and HMAC helpers.
- `errcode`: typed business errors with explicit code lookups.
- `constraint`: reusable generic type constraints.

## Core Helpers

### Pointer helpers

```go
name := util.Ptr("bang")
copy := util.ClonePtr(name)
normalized := util.NilIfZero(util.Ptr(uint64(42)))

v1 := util.MustDeref(name)
v2 := util.DerefOr[string](nil, "fallback")
v3 := util.DerefZero[int64](nil)

_ = normalized
```

### Retry helpers

```go
err := util.Retry(3, fn)
err = util.RetryWithInterval(3, time.Second, fn)
```

### Numeric helpers

```go
value, err := util.IntRandRange[int](1, 10)

sum := util.FloatAdd(0.1, 0.2)
quotient, err := util.FloatDiv(10.0, 2.0)
must := util.MustFloatDiv(10.0, 2.0)
```

### Cryptographic helpers

```go
digest := cipher.SHA256Hex("payload")
signature := cipher.HMACSHA256Hex("secret", "payload")
```

### Error code helpers

```go
businessErr := errcode.New(403, "forbidden")

code, ok := errcode.Code(businessErr)
message, ok := errcode.Message(businessErr)
typed, ok := errcode.As(businessErr)

_ = code
_ = message
_ = typed
```

### File helpers

```go
caller := util.Caller()
exists, err := util.FileExists("/tmp/demo.txt")
err = util.EnsureDir("/tmp/demo")
```

### Time helpers

```go
loc := util.LoadLocationOrFixed("Asia/Shanghai", 8*3600)
```

## Notes

- `IntRandRange` is inclusive on both ends.
- `FloatDiv` returns an error on division by zero; `MustFloatDiv` panics.
- `cipher` exposes only SHA-256 based helpers.
- `LoadLocationOrFixed` falls back to `time.FixedZone` when tzdata is unavailable.
