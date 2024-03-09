# Validation Package

The package provides a set of validation utilities to check various conditions in data fields. These validators can be used to ensure data integrity and enforce specific requirements in your Go applications, making use only of native libraries to perform the validations.

## Objective

The aim is to offer a simple and flexible way to perform data validation in Go applications in a customizable manner. By providing a set of validators, developers can easily incorporate data validation logic into their applications to ensure data consistency and meet specific validation requirements, without the need to write validation code from scratch, customizing validations according to the project's needs.

## Installation

To use the package in your Go projects, type the following command in your terminal:

```bash
go get github.com/dariomatias-dev/go-validadores
```

## Order of Validators

Validators should be organized following the following order: presence validator, type validator, and value validators. They should follow this order because otherwise, an error may occur if the sent value is not accepted by a validator that is placed later, even if it is a valid value.

This organization ensures that basic requirements, such as presence and type, are validated first before more specific validations about the value itself. By validating in this order, we can detect any potential errors early in the process, leading to a more robust and error-free validation system.

Just as there are no reasons to check if the value is of a specific type in value validators, which require the sent value to be of a certain type, as there are dedicated validators for this purpose, thus reducing the number of checks, making the validation process more efficient.

## How to Use

To use the package, first import it into your project with the following import statement:

```go
import "github.com/dariomatias-dev/go-validators"
```

I advise you to give it an alias to make package usage simpler and more efficient, like this:

```go
import v "github.com/dariomatias-dev/go-validators"
```

### Functionality of Validators

The validators have been created based on configurable functions, where the first set of parameters within the first pair of parentheses is used to configure the behavior of the validator, while the second set of parentheses receives the value to be validated.
In the table of [validadores disponíveis](#validators-available), it's referenced which validators require which value to be provided in order to perform validation.

### Usage

To use the validations, use `v.` followed by the desired validation. In the first parenthesis, provide what is being requested, and if you don't want the default error message, insert the desired message afterwards.
The validators will return two values: the first will be the error message if the provided value did not pass validation, and the second will be a boolean value indicating whether the validations should be halted or not. The second value is used in situations where, if the value did not pass the validator, subsequent validations cannot be executed because they will result in an error.

Validations can be performed in three distinct ways: [individually](#validate-individual-value), in [combination](#validate-value-with-multiple-validators), or within a [map](#validate-map).

</br>

#### Validate Individual Value

Um único validador é aplicado a um valor específico.

**Example:**

```go
value := 4

if err, _ := v.Min(3)(value); err != nil {
    log.Fatal(err)
}
```

#### Validate Value with Multiple Validators

Vários validadores são combinados para validar um único valor.

**Example:**

```go
value := 4

errors := v.Validate(
    value,
    v.IsInt(),
    v.Min(3),
    v.Max(10),
)

if len(*errors) != 0 {
    log.Fatal(*errors)
}
```

#### Validate Map

Cada chave do mapa é validada separadamente com seus próprios conjuntos de validadores.

**Example:**

```go
data := map[string]interface{}{
    "name":  "Name",
    "age":   15,
    "email": "emailexample@gmail.com",
}

validations := Validators{
    "name": []Validator{
        IsString(),
        MinLength(3),
        MaxLength(20),
    },
    "age": []Validator{
        IsInt(),
        Min(18),
        Max(100),
    },
    "email": []Validator{
        Email(),
    },
}

errors := ValidateMap(data, validations)
if errors != nil {
    fmt.Println(*errors)
    return
}
```

### Validators Available

<table>
    <tr>
        <th>Validators</th>
        <th>Type</th>
        <th>Input</th>
    </tr>
    <tr>
        <td>IsRequired</td>
        <td>Presence</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsOptional</td>
        <td>Presence</td>
        <td>None</td>
    </tr>
    <tr>
        <td>IsString</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsNumber</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsInt</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsFloat</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsBool</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsArray</td>
        <td>Type</td>
        <td>
            <ul>
                <li>Type of values*</li>
                <li>Array settings*</li>
                <li>Field validators*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsNullString</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsNullNumber</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsNullInt</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsNullFloat</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsNullBool</td>
        <td>Type</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsNullArray</td>
        <td>Type</td>
        <td>
            <ul>
                <li>Type of values*</li>
                <li>Array settings*</li>
                <li>Field validators*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>Min</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Minimum value*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>Max</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Maximum value*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>MinLength</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Minimum size*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>MaxLength</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Maximum size*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsAlpha</td>
        <td>Value</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsAlphaSpace</td>
        <td>Value</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>Regex</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Regex*</li>
                <li>Error message*</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>Custom</td>
        <td>Value</td>
        <td>Custom validator*</td>
    </tr>
    <tr>
        <td>Password</td>
        <td>Value</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>URL</td>
        <td>Value</td>
        <td>Error message</td>
    </tr>
</table>

</br>

**All entries marked with (\*) are mandatory.**
