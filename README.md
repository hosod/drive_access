# drive_access
Google Drive のファイルを操作するためのツールです。
企業用アカウントなら容量に制限がないので、学習用の画像データなどをGoogle Drive上で管理することでコストを削減できるらしいです。
ただGUIだと大量のファイルの管理が手間なので、楽をするためのツールが必要になります。


## 注意
Google Driveのファイルシステムは少し癖があります。
Google Drive APIではマイドライブは`root`という名前で扱われます。パスの指定をするときは`/root/hoge/...`という形で指定するようにしてください。

ファイルやフォルダには全てIDが振られています。Google Drive APIでは基本的にこのIDでファイルの管理を行なっています。このおかげでGoogle Driveでは同じパスのファイルやフォルダが共存できます。(`/root/hoge/fuga`という同じパスのファイルやフォルダが複数存在できる。)
このような重複するパスを指定した時、実際にどのファイルやフォルダを指定するかはファイルのDescriptionを元にユーザーが決定します。事前に入力しておいてください。

また、一度消去してゴミ箱に入ったファイルやフォルダも`root`フォルダ直下のファイルとして扱われます。
不要なファイルは`完全に削除`するようにしてください



## 使い方 
PATHを通して`drive_access/cmd/drive`を実行してください。
初回は認証が必要になります。
token情報が`drive_access/configs`以下にキャッシュされるので2回目以降は認証の必要はありません。

- `ls`: Google Drive上の指定したフォルダの下のファイル一覧を出力します。
  - `-p, --path`: フォルダを指定します。指定しなかった場合`/root`フォルダ(マイドライブ)の直下のファイル一覧を出力します。
    - `drive ls -p=/root/hoge`:マイドライブ直下の`hoge`の下にあるファイル一覧が出力されます。
    - `drive ls`: 何も指定しなければ`/root`フォルダ直下のファイル一覧が出力されます。
- `create`: Google Drive上に新しくフォルダを作成します。
  - `-d, --drive`: 新しいフォルダを作成する場所を指定します。指定しなかった場合`root`フォルダ（マイ ドライブ）の直下に作成します。
    - `drive create -d=/root/hoge/fuga`:マイドライブの直下の`hoge`の下に新しく`fuga`フォルダを作成します。`/root/hoge`が存在していない場合失敗します。
    - `drive create -d=hoge`: `/root`フォルダ(マイ ドライブ)の直下に`hoge`フォルダを作成します。
    
- `download`: Google Drive上のファイルをダウンロードします。フォルダを指定した場合はそのフォルダの下のファイルについても全て再帰的にダウンロードされます。  
  - `-l`: ダウンロード先になるローカルのパス
  
  
  
