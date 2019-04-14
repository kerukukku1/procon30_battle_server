# procon30_battle_server
第30回全国高専プログラミングコンテストの競技部門対戦サーバ
# 開発環境
- go version go1.11.1 darwin/amd64

# 開発規約
## Goフォルダ構成
[Standard Go Project Layout](https://github.com/golang-standards/project-layout)
## Gitブランチ構成
[A successful Git branching model](https://nvie.com/posts/a-successful-git-branching-model/)

## ブランチ詳細
- master
    
    完成したプログラムのみ。versionをコミットメッセージにつけてマージ。ある程度機能がまとまるまではマージしない。

- hotfix

    masterブランチにバグを発見した場合の緊急修正用ブランチ。

- dev

    メインの開発ブランチ。ここから機能の実装のためのブランチを生やす。

    - featureブランチ

        細かい機能をブランチに分けて実装する。実装番号をブランチに記載する。
        
        例: socket.ioの通信部分の実装
        
        feature00001/socket.io_transmission