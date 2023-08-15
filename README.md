<p align="center">
 <img src="https://github.com/jeronimo3875br/rikka/blob/master/assets/rikka_main.gif" alt="rikka_main" width="300"/>
</p>

# 言語ガイド

お好きな言語を選択してください：

- [English](./README.en.md)
- [Português (Portuguese)](./README.pt.md)

# 導入
Rikkaは、<a href="https://github.com/remy/nodemon">nodemon</a>のようなソリューションと同様に、アプリケーションのモニタリングと自動リロードのために開発されたCLIです。しかしながら、Rikkaは任意のプログラミング言語をサポートするように設計され、アプリケーションを手動で再起動する必要や、さまざまなテクノロジー用に異なるパッケージをインストールする必要がありません。

## インストール
Rikkaをインストールするには、お使いのマシンにGOがインストールされている必要があります。最新バージョンのインストール手順は以下で確認できます： https://go.dev/dl/

GOをインストールした後、<a href="https://git-scm.com/">GIT<a/>を使用してRikkaをマシンにダウンロードすることができます。ターミナルで以下のコマンドを実行してください：

```sh 
git clone jeronimo3875br/rikka
```

お使いのマシンの希望する場所に合わせて、パスを調整してください。

## コンパイル 
Rikkaをマシンにインストールした後、次のコマンドを使用してそれを実行可能ファイルにコンパイルできます：

```sh
go build -o rikka main.go
```

コンパイルが完了すると、実行可能なファイルが生成されます。このファイルをお好きなフォルダに移動することができます。

システムの環境変数PATHに実行可能ファイルへのパスを追加し、アクセスを簡便にすることをおすすめします。

生成されたファイルの場所に合わせて、実行可能ファイルのパスを適切に調整してください。

## 使用方法
すべてのコマンドとフラグを確認するには、実行時に**"--help"または"-h"**フラグを渡してください。このフラグはグローバルに機能し、CLIのどこでも使用できます。例えば、以下のように使用します：

```sh
./rikka --help
```

もしくは、環境変数に設定されている場合は、単に：
 
 ```sh
 rikka --help
```

## ファイルの監視と自動リロード
プロジェクトの変更を監視し、実行の自動再読み込みを行うためには、サブコマンドwatchを使用します。このコマンドはいくつかの設定フラグを受け入れます：

- **"--path"** または **"-p"**: プロジェクトへのパスを指定します。デフォルト値は**"./"**で、現在のディレクトリに対応します。
- **"--reflect"** または **"-r"**: プロジェクトの変更が検出された後に実行されるアクション（コマンド）を指定します （必須）。
- **"--ignore"** または **"-i"**: 変更があってもアクションが起こらないようにするために、無視されるフォルダやファイルを設定することができます。

このコマンドを使用するための構文は以下の通りです： `rikka watch --path /caminho/do/seu/programa --reflect "comando de recarregamento"`

以下はいくつかの使用例です：

```sh
# Python
rikka watch --path C:\Users\Jeronimo\projects\web_scraping --reflect "python main.py"

# JavaScript（フォルダ「node_modules」を無視）
rikka watch --path C:\Users\Jeronimo\projects\api --reflect "node bin/server.js" --ignore node_modules

# Golang（「go.mod」と「README.md」を無視）
rikka watch --path C:\Users\Jeronimo\projects\cli --reflect "go run cmd/main.go" --ignore go.mod,README.md
```

# ライセンス
RikkaはMITライセンスの下でリリースされています。詳細はこちらをご覧ください：<a href="https://github.com/jeronimo3875br/rikka/blob/master/LICENSE">ライセンス</a>
