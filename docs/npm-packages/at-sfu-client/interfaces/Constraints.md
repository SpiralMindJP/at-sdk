[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / Constraints

# Interface: Constraints

Constraints は、MediaStreamConstraints を拡張します。

## Hierarchy

- `MediaStreamConstraints`

  ↳ **`Constraints`**

## Table of contents

### Properties

- [audio](Constraints.md#audio)
- [codec](Constraints.md#codec)
- [peerIdentity](Constraints.md#peeridentity)
- [preferCurrentTab](Constraints.md#prefercurrenttab)
- [preferredCodecProfile](Constraints.md#preferredcodecprofile)
- [resolution](Constraints.md#resolution)
- [sendEmptyOnMute](Constraints.md#sendemptyonmute)
- [simulcast](Constraints.md#simulcast)
- [video](Constraints.md#video)

## Properties

### audio

• `Optional` **audio**: `boolean` \| `MediaTrackConstraints`

#### Inherited from

MediaStreamConstraints.audio

___

### codec

• **codec**: `string`

コーデック。

___

### peerIdentity

• `Optional` **peerIdentity**: `string`

#### Inherited from

MediaStreamConstraints.peerIdentity

___

### preferCurrentTab

• `Optional` **preferCurrentTab**: `boolean`

#### Inherited from

MediaStreamConstraints.preferCurrentTab

___

### preferredCodecProfile

• `Optional` **preferredCodecProfile**: `string`

最適なコーデックプロファイルを使用するかどうか。

___

### resolution

• **resolution**: ``"qvga"`` \| ``"vga"`` \| ``"shd"`` \| ``"hd"`` \| ``"fhd"`` \| ``"qhd"``

解像度。

___

### sendEmptyOnMute

• `Optional` **sendEmptyOnMute**: `boolean`

ミュート時に空のストリームを送信するかどうか。

___

### simulcast

• `Optional` **simulcast**: `boolean`

サイマルキャストを使用するかどうか。

___

### video

• `Optional` **video**: `boolean` \| `MediaTrackConstraints`

#### Inherited from

MediaStreamConstraints.video
