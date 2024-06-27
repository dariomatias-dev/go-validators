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
In the table of [validators-available](#validators-available), it's referenced which validators require which value to be provided in order to perform validation.

### Usage

To use the validations, use `v.` followed by the desired validation. In the first parenthesis, provide what is being requested, and if you don't want the default error message, insert the desired message afterwards.
The validators will return two values: the first will be the error message if the provided value did not pass validation, and the second will be a boolean value indicating whether the validations should be halted or not. The second value is used in situations where, if the value did not pass the validator, subsequent validations cannot be executed because they will result in an error.

Validations can be performed in three distinct ways: [individually](#validate-individual-value), or within a [json](#validate-json).

</br>

#### Validate Individual Value

A single validator is applied to a specific value.

**Examples:**

```go
// Success
value := 4

err, _ := v.Min(3)(value) // nil, false
if err != nil {
    fmt.Println(err)
    return
}

// Error
value = 2

err, _ = v.Min(3)(value) // The minimum value is 3, but it received 2, false
if err != nil {
    fmt.Println(err)
    return
}
```

#### Validate JSON

Validates the provided JSON based on the validators defined for each field, returning an error if there is an inconsistency, or assigning the values from the JSON to the provided struct if there are no errors.

**Examples:**

```go
type UserModel struct {
    Name  string `json:"name"  validates:"required;isString;minLength=3;maxLength=20"`
    Age   int    `json:"age"   validates:"required;isInt;min=18;max=100"`
    Email string `json:"email" validates:"required;email"`
}
user := UserModel{}

json := `{
    "name":  "Name",
    "age":   18,
    "email": "emailexample@gmail.com"
}`

// Success
err := Validate(&user, json) // Output: nil
if err != nil {
    fmt.Println(err)
    return
}

// Error
json := `{
    "name":  "Name",
    "age":   16,
    "email": "emailexample@gmail.com"
}`
err := Validate(&user, json) // Output: {"age":["The minimum value is 18, but it received 16."]}
if err != nil {
    fmt.Println(err)
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
        <td>Required</td>
        <td>Presence</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsString</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsNumber</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsInt</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsFloat</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsBool</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsArray</td>
        <td>Type</td>
        <td>
            <ul>
                <li>Field validators*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsNullString</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsNullNumber</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsNullInt</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsNullFloat</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsNullBool</td>
        <td>Type</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsNullArray</td>
        <td>Type</td>
        <td>
            <ul>
                <li>Field validators*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>Email</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Minimum value*</li>
                <li>Error messages</li>
                <ul>
                    <li>Invalid email</li>
                    <li>Value is not string</li>
                </ul>
            </ul>
        </td>
    </tr>
    <tr>
        <td>Password</td>
        <td>Value</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
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
        <td>Length</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Size*</li>
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
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>IsAlphaSpace</td>
        <td>Value</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>StartsWith</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Starts with*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>StartsNotWith</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Starts not with*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>EndsWith</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Ends with*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>EndsNotWith</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Ends not with*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>Regex</td>
        <td>Value</td>
        <td>
            <ul>
                <li>Regex*</li>
                <li>Error message</li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>URL</td>
        <td>Value</td>
        <td>
            <ul>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>OneOf</td>
        <td>Value</td>
        <td>
            <ul>
                <li>
	                Options*
                </li>
                <li>
                    Error message
                </li>
            </ul>
        </td>
    </tr>
    <tr>
        <td>Custom</td>
        <td>Value</td>
        <td>
            <ul>
                <li>
                    Custom validator*
                </li>
            </ul>
        </td>
    </tr>
</table>

### **All entries marked with (\*) are mandatory.**

</br>

# Donations

Help maintain the project with donations.

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dariomatias)
