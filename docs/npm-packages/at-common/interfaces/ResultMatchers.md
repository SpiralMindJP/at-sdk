[@spiralmindjp/at-common](../README.md) / [Exports](../modules.md) / ResultMatchers

# Interface: ResultMatchers<T, E, R1, R2\>

ResultMathcers<T, E, R1, R2> は、matchResult でマッチしたときの処理を定義します。

## Type parameters

| Name | Type |
| :------ | :------ |
| `T` | `T` |
| `E` | extends `Error` |
| `R1` | `R1` |
| `R2` | `R2` |

## Table of contents

### Methods

- [err](ResultMatchers.md#err)
- [ok](ResultMatchers.md#ok)

## Methods

### err

▸ **err**(`error`): `R2`

#### Parameters

| Name | Type |
| :------ | :------ |
| `error` | `E` |

#### Returns

`R2`

___

### ok

▸ **ok**(`value`): `R1`

#### Parameters

| Name | Type |
| :------ | :------ |
| `value` | `T` |

#### Returns

`R1`
