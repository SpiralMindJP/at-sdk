[@spiralmindjp/at-sfu-client](../README.md) / [Exports](../modules.md) / ICECandidateInit

# Interface: ICECandidateInit

ICE の接続候補を表します。

## Table of contents

### Properties

- [candidate](ICECandidateInit.md#candidate)
- [sdpMLineIndex](ICECandidateInit.md#sdpmlineindex)
- [sdpMid](ICECandidateInit.md#sdpmid)
- [usernameFragment](ICECandidateInit.md#usernamefragment)

## Properties

### candidate

• **candidate**: `string`

接続チェックに使用可能な候補のトランスポートアドレス。

___

### sdpMLineIndex

• **sdpMLineIndex**: ``null`` \| `number`

候補が関連付けられているSDPのメディア記述のゼロベースのインデックス番号。

___

### sdpMid

• **sdpMid**: ``null`` \| `string`

候補が関連付けられているコンポーネント内のメディアストリームを一意に識別するメディアストリーム識別タグ。

___

### usernameFragment

• **usernameFragment**: ``null`` \| `string`

ICEがメッセージの整合性に使用するランダムに生成されたユーザー名フラグメント（ice-ufrag）とランダムに生成されたパスワード（ice-pwd）を含む文字列。
