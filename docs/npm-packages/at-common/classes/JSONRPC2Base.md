[@spiralmindjp/at-common](../README.md) / [Exports](../modules.md) / JSONRPC2Base

# Class: JSONRPC2Base

JSONRPC2Base クラスは、WebSocket 上での JSON-RPC2 の基本機能を提供する抽象クラスです。

## Table of contents

### Constructors

- [constructor](JSONRPC2Base.md#constructor)

### Properties

- [socket](JSONRPC2Base.md#socket)

### Accessors

- [onclose](JSONRPC2Base.md#onclose)
- [onerror](JSONRPC2Base.md#onerror)
- [onnotify](JSONRPC2Base.md#onnotify)

### Methods

- [call](JSONRPC2Base.md#call)
- [close](JSONRPC2Base.md#close)
- [closeSocket](JSONRPC2Base.md#closesocket)
- [notify](JSONRPC2Base.md#notify)
- [onClose](JSONRPC2Base.md#onclose-1)
- [onError](JSONRPC2Base.md#onerror-1)
- [onNotify](JSONRPC2Base.md#onnotify-1)
- [connectBase](JSONRPC2Base.md#connectbase)

## Constructors

### constructor

• `Protected` **new JSONRPC2Base**(`socket`)

指定された WebSocket でインスタンスを初期化します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `socket` | `WebSocket` | JSON-RPC2 の通信を行う WebSocket。 |

## Properties

### socket

• `Protected` **socket**: `WebSocket`

## Accessors

### onclose

• `set` **onclose**(`handler`): `void`

onclose プロパティは、WebSocket の close イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`ev`: `Event`) => `void` |

#### Returns

`void`

___

### onerror

• `set` **onerror**(`handler`): `void`

onerror プロパティは、WebSocket の error イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`error`: `Event`) => `void` |

#### Returns

`void`

___

### onnotify

• `set` **onnotify**(`handler`): `void`

onnotify プロパティは、notify イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`method`: `string`, `params`: `any`) => `void` |

#### Returns

`void`

## Methods

### call

▸ `Protected` **call**<`T`\>(`method`, `params`): [`PromiseResult`](../modules.md#promiseresult)<`T`, [`JSONRPC2Error`](JSONRPC2Error.md)\>

call は、指定されたメソッドをコールします。

#### Type parameters

| Name |
| :------ |
| `T` |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `method` | `string` | メソッド名。 |
| `params` | `any` | - |

#### Returns

[`PromiseResult`](../modules.md#promiseresult)<`T`, [`JSONRPC2Error`](JSONRPC2Error.md)\>

メソッドの戻り値を待機する Promise。

___

### close

▸ **close**(): `void`

close は、JSON-RPC2 のストリームをクローズします。

#### Returns

`void`

___

### closeSocket

▸ **closeSocket**(`code?`, `reason?`): `void`

close は、WebSocket をクローズします。

#### Parameters

| Name | Type |
| :------ | :------ |
| `code?` | `number` |
| `reason?` | `string` |

#### Returns

`void`

___

### notify

▸ `Protected` **notify**(`method`, `params`): `void`

notify は、指定されたメソッド名の通知を送信します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `method` | `string` | メソッド名。 |
| `params` | `any` | - |

#### Returns

`void`

___

### onClose

▸ `Protected` **onClose**(`e`): `void`

onClose は、WebSocket がクローズした場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `e` | `Event` | イベントデータ。 |

#### Returns

`void`

___

### onError

▸ `Protected` **onError**(`e`): `void`

onError は、WebSocket でエラーが発生した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `e` | `Event` | イベントデータ。 |

#### Returns

`void`

___

### onNotify

▸ `Protected` **onNotify**(`method`, `params`): `void`

onNotify は、リモートから notify を受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `method` | `string` | メソッド名。 |
| `params` | `any` | - |

#### Returns

`void`

___

### connectBase

▸ `Static` `Protected` **connectBase**(`url`): [`PromiseResult`](../modules.md#promiseresult)<`WebSocket`, `Error`\>

connectBase は、指定されたエンドポイントと WebSocket による接続を確立します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `url` | `string` | エンドポイントの URL。 |

#### Returns

[`PromiseResult`](../modules.md#promiseresult)<`WebSocket`, `Error`\>

WebSocket との接続結果を待機する Promise。
