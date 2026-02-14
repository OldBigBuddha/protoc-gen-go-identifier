# protoc-gen-go-identifier

A buf/protoc plugin that generates Go helper methods for identifier-type protobuf messages.

## Usage

1. Import the options proto in your `.proto` file:

```protobuf
import "plugins/identifier/v1/options.proto";
```

2. Add the option to your identifier message:

```protobuf
message FooID {
  option (plugins.identifier.v1.identifier) = {};
  string id = 1;
}
```

3. Run code generation:

```bash
buf generate
```

4. Use the generated methods:

```go
id := AsFooID("foo-123")
fmt.Println(id.Unwrap())        // "foo-123"
fmt.Println(id.Equal(other))    // true/false
clone := id.Clone()
```

## Generated Methods

| Method | Description |
|--------|-------------|
| `As<Type>(v)` | Constructor |
| `Unwrap()` | Nil-safe value extraction |
| `Equal(other)` | Value comparison |
| `Clone()` | Deep copy via proto.Clone |

## Constraints

- Message MUST have exactly one field
- Field MUST be a scalar type (string, intN, uintN, bytes)
- Field CANNOT be repeated or map

## Skip Options

```protobuf
option (plugins.identifier.v1.identifier) = {
  skip_constructor: true
  skip_unwrap: true
  skip_equal: true
  skip_clone: true
};
```

## Maintenance

If you modify `v1/options.proto`, regenerate the Go code:

```bash
buf generate
```

