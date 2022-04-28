[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / Configuration

# Interface: Configuration

Configuration は、RTCConfiguration を拡張します。

## Hierarchy

- `RTCConfiguration`

  ↳ **`Configuration`**

## Table of contents

### Properties

- [bundlePolicy](Configuration.md#bundlepolicy)
- [certificates](Configuration.md#certificates)
- [codec](Configuration.md#codec)
- [iceCandidatePoolSize](Configuration.md#icecandidatepoolsize)
- [iceServers](Configuration.md#iceservers)
- [iceTransportPolicy](Configuration.md#icetransportpolicy)
- [rtcpMuxPolicy](Configuration.md#rtcpmuxpolicy)

## Properties

### bundlePolicy

• `Optional` **bundlePolicy**: `RTCBundlePolicy`

#### Inherited from

RTCConfiguration.bundlePolicy

___

### certificates

• `Optional` **certificates**: `RTCCertificate`[]

#### Inherited from

RTCConfiguration.certificates

___

### codec

• **codec**: ``"vp8"`` \| ``"vp9"`` \| ``"h264"``

ビデオコーデック。

___

### iceCandidatePoolSize

• `Optional` **iceCandidatePoolSize**: `number`

#### Inherited from

RTCConfiguration.iceCandidatePoolSize

___

### iceServers

• `Optional` **iceServers**: `RTCIceServer`[]

#### Inherited from

RTCConfiguration.iceServers

___

### iceTransportPolicy

• `Optional` **iceTransportPolicy**: `RTCIceTransportPolicy`

#### Inherited from

RTCConfiguration.iceTransportPolicy

___

### rtcpMuxPolicy

• `Optional` **rtcpMuxPolicy**: ``"require"``

#### Inherited from

RTCConfiguration.rtcpMuxPolicy
