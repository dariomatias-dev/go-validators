# Validation Package

Para que as validações sejam aplicadas corretamente, utilize os validadores na seguinte ordem: `IsRequired` ou `IsOptional`, validador de tipo e em seguida os validadores do valor.
É importante que sejam passados nessa ordem para que não seja gerado erros, pois por exemplo, o validador `IsOptional` permitirá que o campo seja nulo, logo não faz sentido aplicar as validações seguintes nele, contudo, se não passar esse validador como primeiro, poderá ocorrer um comportamento inesperado ou um erro.

# Validadores em Go

O pacote oferece um conjunto de utilitários de validação para verificar diversas condições em campos de dados. Esses validadores podem ser usados para garantir a integridade dos dados e impor requisitos específicos em suas aplicações Go, utilizando apenas as bibliotecas nativas para realizar as validações.

## Objetivo

O objetivo é oferecer uma maneira simples e flexível de realizar a validação de dados em aplicações Go de maneira custumizavel. Ao fornecer um conjunto de validadores, os desenvolvedores podem facilmente incorporar lógica de validação de dados em suas aplicações para garantir consistência dos dados e atender aos requisitos específicos de validação, sem a necessidade de escrever o codigo de validacao do zero, podendo custumizar de acordo com as necessidades do projeto.

## Instalação

Para utilizar o pacote Validadores em Go em seu projeto Go, pode instalá-lo usando `go get`:

```bash
go get github.com/dariomatias-dev/go-validadores
```

## Validators

<table>
    <tr>
        <th>Validators</th>
        <th>Type</th>
        <th>Input</th>
    </tr>
    <tr>
        <td>IsRequired</td>
        <td>Presença</td>
        <td>Error message</td>
    </tr>
    <tr>
        <td>IsOptional</td>
        <td>Presença</td>
        <td>Nenhum</td>
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
