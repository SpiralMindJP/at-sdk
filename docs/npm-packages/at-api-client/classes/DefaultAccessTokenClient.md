[@spiralmindjp/at-api-client](../README.md) / [Exports](../modules.md) / DefaultAccessTokenClient

# Class: DefaultAccessTokenClient

DefaultAccessTokenClient は、AccessTokenClient のデフォルト実装を提供します。

## Implements

- [`AccessTokenClient`](../interfaces/AccessTokenClient.md)

## Table of contents

### Constructors

- [constructor](DefaultAccessTokenClient.md#constructor)

### Methods

- [getAccessToken](DefaultAccessTokenClient.md#getaccesstoken)

## Constructors

### constructor

• **new DefaultAccessTokenClient**(`url?`)

指定された URL で、新しいインスタンスを初期化します。

#### Parameters

| Name | Type | Default value | Description |
| :------ | :------ | :------ | :------ |
| `url` | `string` | `'/token/access'` | AccessToken を取得する API のエンドポイント URL。 |

## Methods

### getAccessToken

▸ **getAccessToken**(): `PromiseResult`<[`AccessToken`](../interfaces/AccessToken.md), `Error`\>

getAccessToken は、AccessToken を取得します。

#### Returns

`PromiseResult`<[`AccessToken`](../interfaces/AccessToken.md), `Error`\>

AccessToken を待機する Promise。

#### Implementation of

[AccessTokenClient](../interfaces/AccessTokenClient.md).[getAccessToken](../interfaces/AccessTokenClient.md#getaccesstoken)
