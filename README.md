# 概要

## 背景

テレビには「全録」という全ての番組を録画するというものがあります。

それをラジオでも実現したいと思い、このようなものを作りました。

## 概要

取得したラジオ番組表からcron形式の番組予約リストを出力します。

番組予約リストを```crontab -e```でcronに手動登録します。

# 前提

ラジコを録音できる環境が整っていることを前提にしています。

# 使い方

## 録音パスの変更

録音したファイルを格納するディレクトリを任意のものに変更します。

```
recorded_path := "/home/chinachu/chinachu/recorded/radio/"
```

## 起動方法

```
go run radiko_all_rec.go
```
