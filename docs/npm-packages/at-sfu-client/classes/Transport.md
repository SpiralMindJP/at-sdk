[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / Transport

# Class: Transport

Transport は、WebRTC 接続のトランスポートを表します。

## Table of contents

### Constructors

- [constructor](Transport.md#constructor)

### Properties

- [apiChannel](Transport.md#apichannel)
- [candidates](Transport.md#candidates)
- [motChannel](Transport.md#motchannel)
- [msgChannel](Transport.md#msgchannel)
- [pc](Transport.md#pc)
- [role](Transport.md#role)
- [signal](Transport.md#signal)

## Constructors

### constructor

• **new Transport**(`role`, `signal`, `config`)

指定された Role、シグナリングクライアント、構成で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type |
| :------ | :------ |
| `role` | [`Role`](../enums/Role.md) |
| `signal` | [`Signal`](Signal.md) |
| `config` | `RTCConfiguration` |

## Properties

### apiChannel

• **apiChannel**: `Option`<`RTCDataChannel`\>

API 用の DataChannel オブジェクト。。

___

### candidates

• **candidates**: `RTCIceCandidateInit`[]

ICE 接続候補のリスト。

___

### motChannel

• **motChannel**: `Option`<`RTCDataChannel`\>

モーション用の DataChannel オブジェクト。。

___

### msgChannel

• **msgChannel**: `Option`<`RTCDataChannel`\>

メッセージ用の DataChannel オブジェクト。。

___

### pc

• `Readonly` **pc**: `RTCPeerConnection`

WebRTC のコネクションオブジェクト。

___

### role

• `Readonly` **role**: [`Role`](../enums/Role.md)

___

### signal

• `Readonly` **signal**: [`Signal`](Signal.md)
