# 通信方法

![EndPoint](./images/endpoint.png)

## 参加者

1. /api/regist/{user_name}にGETしてアクセストークンを取得
2. /api/game/{room_id}にroom_idを自分で決めてアクセス。パラメータにはトークンを含める。
3. ゲーム開始