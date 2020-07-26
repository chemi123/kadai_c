# kadai_c
## 課題C
必要なくなれば人知れずこのレポジトリは消え去ります。

### 実行環境
macOSX(goでビルドできて、python3が実行できればどこでも動きます)

### 確認方法
```
// 単体テストの実行
$ go test ./

// ビルド
$ go build c.go

// データセットの作成。以下の例では3つのデータセットを作成
$ python3 create_dataset.py 3 > dataset

// データセットを使って確認
$ ./c < dataset
```

### テスト実行結果例
```
$ python3 create_dataset.py 3 > dataset

$ cat dataset
//　フォーマットに則った数字列(大量にあるため、ここでは表示しない)

$ ./c < dataset
-1
11
19
```
