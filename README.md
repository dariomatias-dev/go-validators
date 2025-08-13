<br>
<div align="center">
  <img src="https://img.shields.io/badge/Go-1.22.3-blue?style=for-the-badge&logo=go&logoColor=white" alt="Go">
</div>
<br>

<h1 align="center">Go Validators</h1>

<p align="center">
  A powerful, lightweight, and dependency-free validation package for Go.
  <br>
  <a href="#about-the-project"><strong>Explore the docs »</strong></a>
  <br><br>
  <a href="https://pkg.go.dev/github.com/dariomatias-dev/go-validators">pkg.go.dev</a>
  ·
  <a href="https://github.com/dariomatias-dev/go-validators/issues">Report Bug</a>
  ·
  <a href="https://github.com/dariomatias-dev/go-validators/issues">Request Feature</a>
</p>

<div align="center">
  <a href="https://pkg.go.dev/github.com/dariomatias-dev/go-validators">
    <img src="https://img.shields.io/badge/Go-1.21-blue?style=flat-square&logo=go&logoColor=white" alt="Go Version">
  </a>
  <a href="https://goreportcard.com/report/github.com/dariomatias-dev/go-validators">
    <img src="https://goreportcard.com/badge/github.com/dariomatias-dev/go-validators" alt="Go Report Card">
  </a>
  <a href="https://pkg.go.dev/github.com/dariomatias-dev/go-validators">
    <img src="https://pkg.go.dev/badge/github.com/dariomatias-dev/go-validators.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/dariomatias-dev/go-validators/blob/main/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT">
  </a>
</div>

## Table of Contents

- [About The Project](#about-the-project)
- [Features](#features)
- [Built With](#built-with)
- [Getting Started](#getting-started)
- [Execution Order](#execution-order)
- [Usage](#usage)
- [Available Validators](#available-validators)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## About The Project

Go Validators is a validation package for Go designed to be powerful, lightweight, and dependency-free. It provides a flexible way to handle data validation, allowing developers to declaratively validate structs from JSON payloads using `validates` tags or use individual validation functions for granular control.

This package is easily integrated into any Go project without introducing external dependencies, offering a simple yet extensible solution for ensuring data integrity.

## Features

- **Struct Validation**: Declaratively validate structs from JSON payloads using `validates` tags.
- **Standalone Functions**: Use individual validation functions for granular control over any value.
- **Zero Dependencies**: A lightweight package that can be integrated into any project without external dependencies.
- **Customizable**: Extend the library with your own logic using the `Custom` validator.

## Built With

This project was developed using the following core technologies:

- **[Go](https://golang.org/)** – An open-source programming language designed for simplicity, reliability, and efficiency in building scalable software.

## Getting Started

To get a local copy up and running, follow these steps.

### Installation

```bash
go get github.com/dariomatias-dev/go-validators
```

## Execution Order

Validators should be organized in a three-stage order within `validates` tags:

1. **Presence Validators (`required`)**: Confirm the field exists.
2. **Type Validators (`isString`, `isInt`, etc.)**: Ensure the value has the correct data type.
3. **Value Validators (`minLength`, `email`, etc.)**: Apply specific rules to the value's content.

This order prevents runtime errors by ensuring that value validators never receive an unexpected type.

## Usage

Import the package using an alias for cleaner code:

```go
import v "github.com/dariomatias-dev/go-validators"
```

### Struct Validation (via Tags)

Validate a JSON payload by mapping it to a struct with `validates` tags. The `Validate` function returns a map of errors or `nil` on success.

```go
type UserModel struct {
    Name  string `json:"name"  validates:"required;isString;minLength=3;maxLength=20"`
    Age   int    `json:"age"   validates:"required;isInt;min=18;max=100"`
    Email string `json:"email" validates:"required;email"`
}

user := UserModel{}

jsonPayload := `{
    "name":  "Name",
    "age":   16,
    "email": "emailexample@gmail.com"
}`

err := v.Validate(&user, jsonPayload)
if err != nil {
    fmt.Println(err)
    // Output: {"age":["The minimum value is 18, but it received 16."]}
    return
}
```

### Standalone Validation

Validators are configurable higher-order functions. The first parentheses configure the rule, the second receives the value.

```go
value := 4
err, _ := v.Min(3, "Value must be at least 3.")(value)
if err != nil {
    fmt.Println(err)
}

value = 2
err, _ = v.Min(3)(value)
if err != nil {
    fmt.Println(err)
    // Output: The minimum value is 3, but it received 2.
    return
}
```

## Available Validators

All parameters marked with (\*) are mandatory. Custom error messages are always optional.

| Validator       | Category | Parameters                  | Applicable Value Type               |
| --------------- | -------- | --------------------------- | ----------------------------------- |
| Required        | Presence |                             | any                                 |
| IsString        | Type     |                             | string                              |
| IsNumber        | Type     |                             | int, float                          |
| IsInt           | Type     |                             | int                                 |
| IsFloat         | Type     |                             | float                               |
| IsBool          | Type     |                             | bool                                |
| IsArray         | Type     | field validators string\*   | slice, array                        |
| IsNullString    | Type     |                             | nil, string                         |
| IsNullNumber    | Type     |                             | nil, int, float                     |
| IsNullInt       | Type     |                             | nil, int                            |
| IsNullFloat     | Type     |                             | nil, float                          |
| IsNullBool      | Type     |                             | nil, bool                           |
| IsNullArray     | Type     | field validators string\*   | nil, slice, array                   |
| Email           | Value    |                             | string                              |
| Password        | Value    |                             | string                              |
| Min             | Value    | minimum value (number)\*    | int, int32, int64, float32, float64 |
| Max             | Value    | maximum value (number)\*    | int, int32, int64, float32, float64 |
| Length          | Value    | size int\*                  | string, slice, array                |
| MinLength       | Value    | minimum size int\*          | string, slice, array                |
| MaxLength       | Value    | maximum size int\*          | string, slice, array                |
| IsAlpha         | Value    |                             | string                              |
| IsAlphaNum      | Value    |                             | string                              |
| IsAlphaSpace    | Value    |                             | string                              |
| IsAlphaNumSpace | Value    |                             | string                              |
| StartsWith      | Value    | prefix string\*             | string                              |
| StartsNotWith   | Value    | prefix string\*             | string                              |
| EndsWith        | Value    | suffix string\*             | string                              |
| EndsNotWith     | Value    | suffix string\*             | string                              |
| Regex           | Value    | regex string\*              | string                              |
| Url             | Value    |                             | string                              |
| Uuid            | Value    | version int (Default: 5)    | string                              |
| OneOf           | Value    | options \[]string\*         | string                              |
| Custom          | Value    | custom validator function\* | any                                 |

## Contributing

Contributions make the open-source community an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch

```sh
git checkout -b feature/AmazingFeature
```

3. Commit your Changes

```sh
git commit -m 'Add some AmazingFeature'
```

4. Push to the Branch

```sh
git push origin feature/AmazingFeature
```

5. Open a Pull Request

## License

Distributed under the **MIT License**. See the [LICENSE](LICENSE) file for more information.

## Contact

I am always open to discussing new projects and ideas. Feel free to get in touch.

- **Portfolio**: [dariomatias-dev](https://dariomatias-dev.com)
- **GitHub**: [dariomatias-dev](https://github.com/dariomatias-dev)
- **Email**: [matiasdario75@gmail.com](mailto:matiasdario75@gmail.com)
- **Instagram**: [@dariomatias_dev](https://instagram.com/dariomatias_dev)
- **LinkedIn**: [linkedin.com/in/dariomatias-dev](https://linkedin.com/in/dariomatias-dev)
