[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / EventStream

# Class: EventStream

EventStream は、イベント受信のストリームを提供します。

## Hierarchy

- `JSONRPC2Base`

  ↳ **`EventStream`**

## Table of contents

### Properties

- [config](EventStream.md#config)
- [socket](EventStream.md#socket)

### Accessors

- [onclose](EventStream.md#onclose)
- [onerror](EventStream.md#onerror)
- [onnotify](EventStream.md#onnotify)
- [onroomevent](EventStream.md#onroomevent)

### Methods

- [call](EventStream.md#call)
- [close](EventStream.md#close)
- [closeSocket](EventStream.md#closesocket)
- [notify](EventStream.md#notify)
- [onClose](EventStream.md#onclose-1)
- [onError](EventStream.md#onerror-1)
- [onNotify](EventStream.md#onnotify-1)
- [onRoomEvent](EventStream.md#onroomevent-1)
- [ready](EventStream.md#ready)
- [connect](EventStream.md#connect)
- [connectBase](EventStream.md#connectbase)

## Properties

### config

• `Readonly` **config**: [`EventStreamConfig`](../interfaces/EventStreamConfig.md)

___

### socket

• `Protected` **socket**: `WebSocket`

#### Inherited from

JSONRPC2Base.socket

## Accessors

### onclose

• `set` **onclose**(`handler`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`ev`: `Event`) => `void` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.onclose

___

### onerror

• `set` **onerror**(`handler`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`error`: `Event`) => `void` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.onerror

___

### onnotify

• `set` **onnotify**(`handler`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`method`: `string`, `params`: `any`) => `void` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.onnotify

___

### onroomevent

• `set` **onroomevent**(`handler`): `void`

onroomevent プロパティは、roomevent イベントを処理するための event handler です。

#### Parameters

| Name | Type |
| :------ | :------ |
| `handler` | (`event`: [`RoomEvent`](../interfaces/RoomEvent.md)) => `void` |

#### Returns

`void`

## Methods

### call

▸ `Protected` **call**<`T`\>(`method`, `params`): `PromiseResult`<`T`, `JSONRPC2Error`\>

#### Type parameters

| Name |
| :------ |
| `T` |

#### Parameters

| Name | Type |
| :------ | :------ |
| `method` | `string` |
| `params` | `any` |

#### Returns

`PromiseResult`<`T`, `JSONRPC2Error`\>

#### Inherited from

JSONRPC2Base.call

___

### close

▸ **close**(): `void`

#### Returns

`void`

#### Inherited from

JSONRPC2Base.close

___

### closeSocket

▸ **closeSocket**(`code?`, `reason?`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `code?` | `number` |
| `reason?` | `string` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.closeSocket

___

### notify

▸ `Protected` **notify**(`method`, `params`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `method` | `string` |
| `params` | `any` |

#### Returns

`void`

#### Inherited from

JSONRPC2Base.notify

___

### onClose

▸ `Protected` **onClose**(`ev`): `void`

onClose は、WebSocket が閉じられた場合の処理を行ないます。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `ev` | `CloseEvent` | イベントデータ。 |

#### Returns

`void`

#### Overrides

JSONRPC2Base.onClose

___

### onError

▸ `Protected` **onError**(`ev`): `void`

onClose は、WebSocket でエラーが発生した場合の処理を行ないます。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `ev` | `Event` | イベントデータ。 |

#### Returns

`void`

#### Overrides

JSONRPC2Base.onError

___

### onNotify

▸ `Protected` **onNotify**(`method`, `params`): `void`

onNotify は、サーバーから notify を受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `method` | `string` | メソッドの名前。 |
| `params` | `any` | パラメーター。 |

#### Returns

`void`

#### Overrides

JSONRPC2Base.onNotify

___

### onRoomEvent

▸ `Protected` **onRoomEvent**(`event`): `void`

onRoomEvent は、RoomEvent を受信した場合の処理を行います。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `event` | [`RoomEvent`](../interfaces/RoomEvent.md) | 受信した RoomEvent。 |

#### Returns

`void`

___

### ready

▸ **ready**(): `PromiseResult`<[`DeviceInfo`](../interfaces/DeviceInfo.md), `JSONRPC2Error`\>

ready は、ready API をコールします。

#### Returns

`PromiseResult`<[`DeviceInfo`](../interfaces/DeviceInfo.md), `JSONRPC2Error`\>

サーバーから返された DeviceInfo を待機する Promise。

___

### connect

▸ `Static` **connect**(`config?`): `PromiseResult`<[`EventStream`](EventStream.md), `Error`\>

connect は、EventStream API に接続します。

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `config` | [`EventStreamConfig`](../interfaces/EventStreamConfig.md) | EventStream の構成。 |

#### Returns

`PromiseResult`<[`EventStream`](EventStream.md), `Error`\>

EventStream を待機する Promise。

___

### connectBase

▸ `Static` `Protected` **connectBase**(`url`): `PromiseResult`<`WebSocket`, `Error`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `url` | `string` |

#### Returns

`PromiseResult`<`WebSocket`, `Error`\>

#### Inherited from

JSONRPC2Base.connectBase
