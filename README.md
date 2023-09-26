# go-sudoku
数独の問題を回答するプログラム

[Goライブコーディング(技育祭2020)](https://www.youtube.com/watch?v=IyPujQ381YY)の写経リポジトリ

# Usage
オプションを見る
```
$ go-sudoku -h
```
数独の問題を指定して実行する
- `9×9`の問題にすること(`.`は`0`を表す)
- 指定しない場合は全て空欄の状態から回答を出力する
```
$ go-sudoku -question=.5..83.17...1..4..3.4..56.8....3...9.9.8245....6....7...9....5...729..861.36.72.4
```
回答の探索状況を出力しながら実行する
- 1マスずつ回答を試行して出力する
- 指定しない場合は回答のみを表示する
```
$ go-sudoku -showProgress
```
# Installation
```
$ go install github.com/cr4ne89/go-sudoku@latest
```
