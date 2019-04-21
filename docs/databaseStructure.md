# データベースの構造

## ゲーム進行DB
- id, number, not null
- roomName, string, not null
- phase, number, not null
- playerId1, string
- playerId2, string

## ユーザDB
- id, number, not null
- token, string, not null
- name, string, not null

## ゲームヒストリDB
- id, number, not null
- gameId, number, not null
- playerId1Action, text
- playerId2Action, text