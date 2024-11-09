# pointer

Pointer is a library that makes getting arbitrary pointers in Go easier.

## Usage

```shell
go get github.com/verygoodsoftwarenotvirus/pointer@latest
```

```go
func someFunctionThatRequiresManyStringPointers(
    arg1,
    arg2,
    arg3,
    arg4,
    arg5,
    arg6,
    arg7,
    arg8 *string
) error {
    return nil
}

// the old and busted way:

first := "first"
second := "second"
third := "third"
fourth := "fourth"
fifth := "fifth"
sixth := "sixth"
seventh := "seventh"
eighth := "eighth"

// the old busted way:
someFunctionThatRequiresManyStringPointers(
    &first,
    &second,
    &third,
    &fourth,
    &fifth,
    &sixth,
    &seventh,
    &eighth,
)

// with this library:

someFunctionThatRequiresManyStringPointers(
    pointer.To("first"),
    pointer.To("second"),
    pointer.To("third"),
    pointer.To("fourth"),
    pointer.To("fifth"),
    pointer.To("sixth"),
    pointer.To("seventh"),
    pointer.To("eighth"),
)

```

There's also a corresponding `Deref` method that deferences a pointer, and those same functions but applied to every element of a slice or a map.
