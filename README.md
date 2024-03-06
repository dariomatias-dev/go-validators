# Validation Package

O pacote oferece um conjunto de utilitários de validação para verificar diversas condições em campos de dados. Esses validadores podem ser usados para garantir a integridade dos dados e impor requisitos específicos em suas aplicações Go, utilizando apenas as bibliotecas nativas para realizar as validações.

## Objetivo

O objetivo é oferecer uma maneira simples e flexível de realizar a validação de dados em aplicações Go de maneira custumizavel. Ao fornecer um conjunto de validadores, os desenvolvedores podem facilmente incorporar lógica de validação de dados em suas aplicações para garantir consistência dos dados e atender aos requisitos específicos de validação, sem a necessidade de escrever o código de validação do zero, podendo custumizar de acordo com as necessidades do projeto.

## Instalação

Para utilizar o pacote em seus projetos Go, digite o seguinte comando em seu terminal:

```bash
go get github.com/dariomatias-dev/go-validadores
```

## How to Use

Os validadores devem ser organizados seguindo a seguinte ordem: validador de presença, validador de tipo e validadores de valor. Eles devem seguir essa ordem porque, caso contrário, pode ocorrer um erro se o valor enviado não for aceito por um validador que está colocado posteriormente, mesmo que seja um valor válido.

Essa organização garante que os requisitos básicos, como presença e tipo, sejam validados primeiro antes de validações mais específicas sobre o valor em si. Ao validar nessa ordem, podemos detectar quaisquer erros potenciais no início do processo, levando a um sistema de validação mais robusto e livre de erros.

Assim como não há razões para verificar se o valor é de um tipo específico nos validadores de valor, que necessitam que o valor enviado seja de determinado tipo, pois existe validadores dedicados a esse próposito, diminuindo assim a quantidade de verificações deixando o processo de validação mais eficientes.

### Validators Available

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
