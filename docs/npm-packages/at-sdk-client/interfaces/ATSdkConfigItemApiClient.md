[@spiralmindjp/at-sdk-client](../README.md) / [Exports](../modules.md) / ATSdkConfigItemApiClient

# Interface: ATSdkConfigItemApiClient

Interface for configuring the API client.

## Table of contents

### Properties

- [accessTokenURL](ATSdkConfigItemApiClient.md#accesstokenurl)
- [contentAutomaticUpdateInterval](ATSdkConfigItemApiClient.md#contentautomaticupdateinterval)
- [coreBaseURL](ATSdkConfigItemApiClient.md#corebaseurl)
- [coreWebSocketBaseURL](ATSdkConfigItemApiClient.md#corewebsocketbaseurl)
- [debugApiClient](ATSdkConfigItemApiClient.md#debugapiclient)
- [sfuConfig](ATSdkConfigItemApiClient.md#sfuconfig)

## Properties

### accessTokenURL

• **accessTokenURL**: `string`

Relative URL (relative to core server) to get an access token, such as '/api/token/access'

___

### contentAutomaticUpdateInterval

• **contentAutomaticUpdateInterval**: `number`

Millisecond interval between automatic content checking and download (recommended value 600000 ms = 10 minutes)

___

### coreBaseURL

• **coreBaseURL**: `string`

URL to the core server, such as 'https://at-core.internal:50001'

___

### coreWebSocketBaseURL

• **coreWebSocketBaseURL**: `string`

URL to the core websocket server, such as 'wss://at-core.internal:50001'

___

### debugApiClient

• **debugApiClient**: `boolean`

Boolean flag to enable additional console debug messages relating to the API client

___

### sfuConfig

• **sfuConfig**: `Configuration`

Configuration object for the SFU client
