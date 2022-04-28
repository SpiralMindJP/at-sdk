[@spiralmindjp/at-common](README.md) / Exports

# @spiralmindjp/at-common

## Table of contents

### Classes

- [JSONRPC2Base](classes/JSONRPC2Base.md)
- [JSONRPC2Error](classes/JSONRPC2Error.md)

### Interfaces

- [ErrorMessage](interfaces/ErrorMessage.md)
- [OptionMatchers](interfaces/OptionMatchers.md)
- [ResultMatchers](interfaces/ResultMatchers.md)

### Type aliases

- [Failure](modules.md#failure)
- [None](modules.md#none)
- [Option](modules.md#option)
- [PromiseOption](modules.md#promiseoption)
- [PromiseResult](modules.md#promiseresult)
- [Result](modules.md#result)
- [Some](modules.md#some)
- [Success](modules.md#success)

### Functions

- [failure](modules.md#failure-1)
- [failureException](modules.md#failureexception)
- [matchOption](modules.md#matchoption)
- [matchResult](modules.md#matchresult)
- [none](modules.md#none-1)
- [some](modules.md#some-1)
- [someOrNone](modules.md#someornone)
- [success](modules.md#success-1)
- [successVoid](modules.md#successvoid)

## Type aliases

### Failure

Ƭ **Failure**<`E`\>: `Object`

Failure<E> は、失敗を表し、エラーを格納します。

#### Type parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `E` | extends `Error` | エラーを表す Error クラスを継承する型。 |

#### Type declaration

| Name | Type |
| :------ | :------ |
| `error` | `E` |
| `ok` | ``false`` |

___

### None

Ƭ **None**: `Object`

None は、値が存在しない型を表します。

#### Type declaration

| Name | Type |
| :------ | :------ |
| `some` | ``false`` |

___

### Option

Ƭ **Option**<`T`\>: [`Some`](modules.md#some)<`T`\> \| [`None`](modules.md#none)

Option<T> は値が存在する場合と存在しない場合を表します。

**`remarks`**
Option<T> は値が存在する Some<T> と、値が存在しない None の直和型です。
Some<T> は、`{ some: true; value: T }` を表します。
None は、`{ some: false }` を表します。

#### Type parameters

| Name | Description |
| :------ | :------ |
| `T` | 値の型。 |

___

### PromiseOption

Ƭ **PromiseOption**<`T`\>: `Promise`<[`Option`](modules.md#option)<`T`\>\>

PromiseOption<T> は、Promise<Option<T>> の型エイリアスです。

#### Type parameters

| Name |
| :------ |
| `T` |

___

### PromiseResult

Ƭ **PromiseResult**<`T`, `E`\>: `Promise`<[`Result`](modules.md#result)<`T`, `E`\>\>

PromiseResult<T, E> は、Promise<Result<T, E>> の型エイリアスです。

#### Type parameters

| Name | Type |
| :------ | :------ |
| `T` | `void` |
| `E` | extends `Error` = `Error` |

___

### Result

Ƭ **Result**<`T`, `E`\>: [`Success`](modules.md#success)<`T`\> \| [`Failure`](modules.md#failure)<`E`\>

Result<T, E> は結果を格納します。

**`remarks`**
Result<T, E> は結果を格納する Success<T> と Failure<E> の直和型です。
Success<T> は、`{ ok: true; value: T }` を表します。
Failure<E> は、`{ ok: false; error: E }` を表します。

#### Type parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `T` | `void` | 結果を表す型。 |
| `E` | extends `Error` = `Error` | エラーを表す Error クラスを継承する型。 |

___

### Some

Ƭ **Some**<`T`\>: `Object`

Some<T> は、値が存在する型を表し、その値を格納します。

#### Type parameters

| Name | Description |
| :------ | :------ |
| `T` | 値の型。 |

#### Type declaration

| Name | Type |
| :------ | :------ |
| `some` | ``true`` |
| `value` | `T` |

___

### Success

Ƭ **Success**<`T`\>: `Object`

Success<T> は、成功を表し、結果の値を格納します。

#### Type parameters

| Name | Description |
| :------ | :------ |
| `T` | 結果を表す型。 |

#### Type declaration

| Name | Type |
| :------ | :------ |
| `ok` | ``true`` |
| `value` | `T` |

## Functions

### failure

▸ **failure**<`T`, `E`\>(`error`): [`Result`](modules.md#result)<`T`, `E`\>

success は、Error を格納した失敗を表す Result<T, E> を返します。

#### Type parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `T` | `T` | 結果を表す型。 |
| `E` | extends `Error` | エラーを表す Error クラスを継承する型。 |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `error` | `E` | Error。 |

#### Returns

[`Result`](modules.md#result)<`T`, `E`\>

Error を格納した失敗を表す Result<T, E>。

___

### failureException

▸ **failureException**<`T`\>(`e`): [`Result`](modules.md#result)<`T`, `Error`\>

failureException は、任意の例外エラーから失敗を表す Result<T, E> を表します。

#### Type parameters

| Name | Description |
| :------ | :------ |
| `T` | 結果を表す型。 |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `e` | `any` | 例外エラー。 |

#### Returns

[`Result`](modules.md#result)<`T`, `Error`\>

例外 Error または例外の値をラップした Error を格納した失敗を表す Result<T, E>。

___

### matchOption

▸ **matchOption**<`T`, `R1`, `R2`\>(`option`, `matchers`): `R1` \| `R2`

matchOption は、指定された Option<T> の値が存在する場合と値が存在しない場合にマッチしたときに、各処理を実行します。

#### Type parameters

| Name | Description |
| :------ | :------ |
| `T` | 値の型。 |
| `R1` | - |
| `R2` | - |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `option` | [`Option`](modules.md#option)<`T`\> | マッチさせる Option<T>。 |
| `matchers` | [`OptionMatchers`](interfaces/OptionMatchers.md)<`T`, `R1`, `R2`\> | 値が存在したときまたは値が存在しないときの処理を行う OptionMatchers<T, R1, R2>。 |

#### Returns

`R1` \| `R2`

マッチした結果の戻り値。

___

### matchResult

▸ **matchResult**<`T`, `E`, `R1`, `R2`\>(`result`, `matchers`): `R1` \| `R2`

matchResult は、指定された Result<T, E> が成功または失敗にマッチしたときに、各処理を実行します。

#### Type parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `T` | `T` | 結果を表す型。 |
| `E` | extends `Error` | エラーを表す Error クラスを継承する型。 |
| `R1` | `R1` | - |
| `R2` | `R2` | - |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `result` | [`Result`](modules.md#result)<`T`, `E`\> | マッチさせる Result<T, E>。 |
| `matchers` | [`ResultMatchers`](interfaces/ResultMatchers.md)<`T`, `E`, `R1`, `R2`\> | 成功または失敗時の処理を行う ResultMatchers<T, E, R1, R2>。 |

#### Returns

`R1` \| `R2`

マッチした結果の戻り値。

___

### none

▸ **none**<`T`\>(): [`Option`](modules.md#option)<`T`\>

some は、値が存在しない Option<T> を返します。

#### Type parameters

| Name | Description |
| :------ | :------ |
| `T` | 値の型。 |

#### Returns

[`Option`](modules.md#option)<`T`\>

値が存在しない Option<T>。

___

### some

▸ **some**<`T`\>(`value`): [`Option`](modules.md#option)<`T`\>

some は、値が存在する Option<T> を返します。

#### Type parameters

| Name | Description |
| :------ | :------ |
| `T` | 値の型。 |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `value` | `T` | 値。 |

#### Returns

[`Option`](modules.md#option)<`T`\>

値が存在する Option<T>。

___

### someOrNone

▸ **someOrNone**<`T`\>(`value`): [`Option`](modules.md#option)<`T`\>

someOrNone は、任意の値を受け取り、該当する Option<T> を返します。

**`remarks`**
someOrNone は、value が undefined または null の場合に、値が存在しない Option<T> を返します。
それ以外の場合は、value を値とする値が存在する Option<T> を返します。

#### Type parameters

| Name | Description |
| :------ | :------ |
| `T` | 値の型。 |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `value` | `undefined` \| ``null`` \| `T` | 値。 |

#### Returns

[`Option`](modules.md#option)<`T`\>

値が存在する Option<T> または値が存在しない Option<T>。

___

### success

▸ **success**<`T`, `E`\>(`value`): [`Result`](modules.md#result)<`T`, `E`\>

success は、結果の値を格納した成功を表す Result<T, E> を返します。

#### Type parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `T` | `T` | 結果を表す型。 |
| `E` | extends `Error` | エラーを表す Error クラスを継承する型。 |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `value` | `T` | 結果の値。 |

#### Returns

[`Result`](modules.md#result)<`T`, `E`\>

結果の値 value を格納した成功を表す Result<T, E>。

___

### successVoid

▸ **successVoid**<`E`\>(): [`Result`](modules.md#result)<`void`, `E`\>

success は、結果の値を持たない成功を表す Result<T, E> を返します。

#### Type parameters

| Name | Type |
| :------ | :------ |
| `E` | extends `Error` |

#### Returns

[`Result`](modules.md#result)<`void`, `E`\>

結果の値 value を持たない成功を表す Result<T, E>。
