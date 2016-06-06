# `go generate`

 - ここでの内容
     - go generateとはどういったものか
     - 自作のgo generateってどんな感じなのか
     - go generateでさくさくコーディング

## go generateとは？

 - 2014/12 の Go1.4で追加されたコードを生成する機能
 -  `//go:generate [go generateのコマンド] [そのコマンドに対応したパラメータ]`

```main.go
//go:generate gen -force

package food

// +gen *
type Food struct {
    Name string
    Price int
}
```

 - 内部的にはコンパイルされたGoのバイナリを動かしているだけ。上記の場合は `gen` コマンドにパラメータを与えて実行しているだけ。
 - `gen`とは何？`gen`とは、`go generate`の外部ライブラリ。上記は実行前にこちらhttps://github.com/clipperhouse/genを `go get`してます。
 - コマンドは `$GOPATH/bin` にあるものが実行される。
 - これを実行すると `food_gen.go` というファイルが同階層に作成される `https://gist.github.com/mattn/6daf8425916fc9a9f516`

## stringerでgo generate してみる

 - cloneしてください。

```
git clone git@github.com:sgswtky/mokumoku4.git
cd tutorial1
go run main.go
```

 - このまま実行すると、 `const` の中身がそのままpirntされます。
 - go getします

```
go get golang.org/x/tools/cmd/stringer
```

 - ここで、8行目に `//go:generate stringer -type=Water` を追加して `go generate` してください。
 - `water_string.go` というファイルが生成されました。

```この時点でのディレクトリ
 - tutorial1
  |- main.go
  |- water_string.go
```

 - これらのファイルを実行します `go run main.go water_string.go`
 - 先ほどはconstの数字がprintされたのに対して、今度は対応付けられた constの名前がprintされました。
 - `water_string.go`をみてもらうとわかりますが、 generateした時点でのconstを元にして、対応付けられた数字からconstの名前を返却する関数を作成しています。

## 自作のプログラムで go generateしてみる

### 準備

 - 必要な物を`go get`して `cd` します

```
go get github.com/labstack/echo
go get github.com/labstack/echo/engine/standard
go get github.com/sgswtky/echo-generate-router
cd mokumoku4/tutorial2
```

 -  `cd` したディレクトリにあるプログラムはechoフレームワークを使用したwebアプリケーションです。
 - `main.go` に `// @router`という所がありますが、これを元にルーティングさせるコードを自動生成するのが、↑こちらで go getした `echo-generate-router` です。

### go generate

 - コードを生成してください。

```
go generate
```

### 成果物の確認

 - 生成されたファイルを確認すると、ルーティング処理が記述されたファイルが生成されます。
 - これを含めて実行して、curlで叩いてみましょう

```
go run *
# 別のウィンドウで↓実行
curl http://localhost:7070/mobile/smartphone/ios/version/changelog
```

 - 自作のルーティング生成プログラムで `go generate` してみました。

### この go generateでできること

     - ルーティングの自動生成
     - 関数名をルーティング内容から自動生成するのでコーディング規約のボリュームが減る。
         - `@router GET sample` → `e.GET("/sample", getSample)`
     - echoでは対応していないルーティングのグループ化が可能

### 自作のgo generatenの注意点

 - 乱用するとただの黒魔術

## genで go generateしてみる

 - go get して cd して go generate

```
github.com/clipperhouse/gen
cd mokumoku4/tutorial3
```

### 生成されたものを見てみる

 - genというgo generateのプログラムを使ってみます。
 - genは宣言しているtypeの上に `// +gen` と書くことで色々なコードを生成します。
 - 例えばこのチュートリアルでは、Water構造体のスライスを構造体として定義し、その構造体をレシーバーとしたメソッドの生成を行っています。
 - 生成された `water_slice.go` の中身を確認してみると、以下のWaterスライスのレシーバーのメソッドが生成されていることが確認できます。

```
func (rcv WaterSlice) Where(fn func(Water) bool) (result WaterSlice)
```

### レシーバーを追加してみる

 - main.goの `// +gen slice:"Count" `に Whereを追加してください。　
 - `go generate` して `water_slice.go` を確認すると、レシーバーにCountが追加されていることがわかります。

```
func (rcv WaterSlice) Count(fn func(Water) bool) (result int)
```

### 構造体のレシーバーを呼び出す

 - `gen`のマニュアル http://clipperhouse.github.io/gen/
 - それぞれmainの最後に追加して`go run *`で実行してみてください。

```count
// 全てのWaterを数える
	countLogic := func(w Water) bool {
		return true
	}
	fmt.Println(waters.Count(countLogic))
```

```where
// priceが50より大きいWaterを取得
	whereLogic := func(w Water) bool {
		return w.price > 50
	}
	fmt.Println(fmt.Sprintf("%+v", waters.Where(whereLogic)))
```

## もくもくたいむ

 - `gen` のマニュアルを見て、必要に応じて `go generate`でシーバー追加していく感じで進めてください。　
 - cdしてください

```
cd mokumoku4/moku2
```

### 課題①

 - あなたは1日に水を10万本を激安で売る水の商人です。
 - 明日納品分のCSVを取り込んで必要な情報だけを表示する必要があります。
 - 色々と確認したところ以下の情報をCSVから確認したいことがわかりました。
     - 金額の合計（商品別、全商品）
     - 容量の合計（商品別、全商品）
     - 商品別の利益、全商品の利益、全商品の利率

```出力フォーマット。とりあえずこのフォーマット埋める感じで出力してもらえれば大丈夫です。

------------------------------
金額の合計
------------------------------
Irohasu : int
Volvic : int
CrystalGeyser : int
AlcaliIonWater : int
Perrier :  int
Contrex : int
Wilkinson :  int
AmebaWater :  int
Onsensui :  int
全ての商品の合計 : int

------------------------------
容量の合計
------------------------------
Irohasu : int
Volvic : int
CrystalGeyser : int
AlcaliIonWater : int
Perrier :  int
Contrex : int
Wilkinson :  int
AmebaWater :  int
Onsensui :  int
全ての商品の合計 : int

------------------------------
全商品の利益
------------------------------
Irohasu : int
Volvic : int
CrystalGeyser : int
AlcaliIonWater : int
Perrier :  int
Contrex : int
Wilkinson :  int
AmebaWater :  int
Onsensui :  int
全ての商品の利益 : int

------------------------------
全商品の利率
------------------------------
Irohasu : float32
Volvic : float32
CrystalGeyser : float32
AlcaliIonWater : float32
Perrier :  float32
Contrex : float32
Wilkinson :  float32
AmebaWater :  float32
Onsensui :  float32
```

```商品の利益率
Irohasu : 0.12
Volvic : 0.2
CrystalGeyser : 0.05
AlcaliIonWater : 0.1
Perrier :  0.1
Contrex : 0.2
Wilkinson :  0.13
AmebaWater : 0
Onsensui : 0.4
```

### 課題②

 - 水を売り続けていると、お客さんから当日の水の情報が即座にわかるようにしてほしいという要望がありました。
 - あなたはお客さんのために水の情報に関する全てを見える化しようときめました。
     - 水で容量あたり一番安い水はどれか
     - 入荷している水のリストが欲しい（商品名、価格、在庫数）


### 課題③

 - もし時間余ったら処理見なおしたり並列化して早くしてください。