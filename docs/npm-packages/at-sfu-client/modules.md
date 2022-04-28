[@spiralmindjp/at-sfu-client](README.md) / Exports

# @spiralmindjp/at-sfu-client

## Table of contents

### Enumerations

- [OperationType](enums/OperationType.md)
- [Role](enums/Role.md)

### Classes

- [Client](classes/Client.md)
- [LocalStream](classes/LocalStream.md)
- [NoTracksError](classes/NoTracksError.md)
- [Signal](classes/Signal.md)
- [Transport](classes/Transport.md)

### Interfaces

- [ActiveLayer](interfaces/ActiveLayer.md)
- [Configuration](interfaces/Configuration.md)
- [Constraints](interfaces/Constraints.md)
- [DeviceInfo](interfaces/DeviceInfo.md)
- [Encoding](interfaces/Encoding.md)
- [Face](interfaces/Face.md)
- [Head](interfaces/Head.md)
- [ICECandidateInit](interfaces/ICECandidateInit.md)
- [Motion](interfaces/Motion.md)
- [Operation](interfaces/Operation.md)
- [RemoteStream](interfaces/RemoteStream.md)
- [SignalConfig](interfaces/SignalConfig.md)
- [Trickle](interfaces/Trickle.md)
- [Vector3](interfaces/Vector3.md)
- [VideoConstraints](interfaces/VideoConstraints.md)

### Type aliases

- [Layer](modules.md#layer)
- [Resolutions](modules.md#resolutions)
- [TrackKind](modules.md#trackkind)

### Variables

- [VideoConstraints](modules.md#videoconstraints)

### Functions

- [makeRemoteStream](modules.md#makeremotestream)

## Type aliases

### Layer

Ƭ **Layer**: ``"low"`` \| ``"medium"`` \| ``"high"``

Layer は、サイマルキャストのレイヤーを定義します。

___

### Resolutions

Ƭ **Resolutions**: typeof `resolutions`[`number`]

Resolutions は、解像度定義の型を表します。

___

### TrackKind

Ƭ **TrackKind**: ``"audio"`` \| ``"video"``

TrackKind は、メディアトラックの種類を表します。

## Variables

### VideoConstraints

• **VideoConstraints**: [`VideoConstraints`](modules.md#videoconstraints)

VideoConstraints は、解像度毎のビデオの制限を定義します。

## Functions

### makeRemoteStream

▸ **makeRemoteStream**(`stream`, `transport`, `deviceInfo`): [`RemoteStream`](interfaces/RemoteStream.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `stream` | `MediaStream` |
| `transport` | [`Transport`](classes/Transport.md) |
| `deviceInfo` | [`DeviceInfo`](interfaces/DeviceInfo.md) |

#### Returns

[`RemoteStream`](interfaces/RemoteStream.md)
