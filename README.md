# Validation Package

The package provides a set of validation utilities to check various conditions in data fields. These validators can be used to ensure data integrity and enforce specific requirements in your Go applications, making use only of native libraries to perform the validations.

## Objective

The aim is to offer a simple and flexible way to perform data validation in Go applications in a customizable manner. By providing a set of validators, developers can easily incorporate data validation logic into their applications to ensure data consistency and meet specific validation requirements, without the need to write validation code from scratch, customizing validations according to the project's needs.

## Installation

To use the package in your Go projects, type the following command in your terminal:

```bash
go get github.com/dariomatias-dev/go-validadores
```

## How to Use

Validators should be organized following the following order: presence validator, type validator, and value validators. They should follow this order because otherwise, an error may occur if the sent value is not accepted by a validator that is placed later, even if it is a valid value.

This organization ensures that basic requirements, such as presence and type, are validated first before more specific validations about the value itself. By validating in this order, we can detect any potential errors early in the process, leading to a more robust and error-free validation system.

Just as there are no reasons to check if the value is of a specific type in value validators, which require the sent value to be of a certain type, as there are dedicated validators for this purpose, thus reducing the number of checks, making the validation process more efficient.

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

**All entries marked with (*) are mandatory.**
