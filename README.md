# Go v1.16に追加されembedを試す

# embedのメモ
- ディレクティブで埋め込める変数の型は、string、[]byte、embed.FS
- 埋め込みは関数の内部など閉じたスコープで行うことができない
- hoge/*や、hoge/*.sqlのようなワイルドカードが利用できる
- ../のように親ファイルまで遡って読み込みはできない
- //の後に半角スペースを入れてしまうと本当にコメントアウトとして処理されてしまう

# サンプル実行
```
$ go run 1_read_text_file/main.go
```

# サンプルビルド
```
$ go build -o read_text_file 1_read_text_file/main.go
# 実行する
$ ./read_text_file
```
