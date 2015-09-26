AviUtl Script Library

## AviUtl Script Libraryとは
有志の方々が公開されているAviUtl スクリプトをできる限り集め、CSV形式でまとめたもの。
現在はbowlrollで配布中(https://bowlroll.net/file/66110)

## ソースコードについて
AviUtl Script Libraryを作るために作ったプログラムおよびデータ。(2014年に使えたAPIが使えなくなっていたため、一部書き換えて公開)

全体の流れとしてはニコニコ動画のAPIを叩き、AviUtl スクリプト・AviUtl スクリプト講座・AviUtl講座の3つのタグでそれぞれ検索される動画のvideo IDを取得し、重複分を削除した後、それぞれの動画の情報を取得し、CSV形式に変換するまでのプログラム。
もちろんこの後AviUtl Scriptとは関係ない動画情報は削除したり、ニコニコ動画では公開されていないスクリプトの追加を行ったりしたけどほとんど手動なので省略(大変だった)。

今回はAviUtl Scriptに重点を置いてまとめたけど、ニコニコ動画で配布されているMMDのデータをまとめるのにも使えるかも(と個人的に思っている、やらないけど・・・)。

## 実行環境
msys2(bash、perl)、Go 


