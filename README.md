# drive_access
Google Drive のファイルを操作するためのツールです。
企業用アカウントなら容量に制限がないので、学習用の画像データなどをGoogle Drive上で管理することでコストを削減できるらしいです。
ただGUIだと大量のファイルの管理が手間なので、楽をするためのツールが必要になります。


## はじめに
Google Driveのファイルシステムは少し癖があります。
Google Drive APIではマイドライブは`root`という名前で扱われます。パスの指定をするときは`/root/hoge/...`という形で指定するようにしてください。

ファイルやフォルダには全てIDが振られています。Google Drive APIでは基本的にこのIDでファイルの管理を行なっています。このおかげでGoogle Driveでは同じパスのファイルやフォルダが共存できます。(`/root/hoge/fuga`という同じパスのファイルやフォルダが複数存在できる。)
このような重複するパスを指定した時、実際にどのファイルやフォルダを指定するかはファイルのDescriptionを元にユーザーが決定します。事前に入力しておいてください。



## 使い方 
実行ファイル`drive_access/cmd/access`を実行してください。
初回は認証が必要になります。
token情報が`drive_access/configs`以下にキャッシュされるので2回目以降は認証の必要はありません。

`/root/your_dir`
`your_dir`

- `create`: Google Drive上に新しくフォルダを作成します。
  - `-d, --drive`: 新しいフォルダを作成する場所を指定します。指定しなかった場合`root`フォルダ（マイ ドライブ）の直下に作成します。
    - `drive create -d=/root/hoge/your_dir`:マイドライブの直下の`hoge`の下に新しく`your_dir`を作成します。`/root/hoge`が存在していない場合失敗します。
    - `drive create -d=your_dir`: `root`フォルダ(マイ ドライブ)の直下に`your_dir`を作成します。
