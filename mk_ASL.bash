#!/bin/bash

# ビルド
go build getsmid.go
go build smi2csv.go

# 指定タグで検索される動画のvideo IDを取得
./getsmid.exe AviUtlスクリプト 1 2 > smid.list
./getsmid.exe AviUtlスクリプト講座 1 7 >> smid.list
./getsmid.exe AviUtl講座 1 21 >> smid.list

# 重複を取り除く
cat smid.list | sort | uniq | perl -lne 'if(/(sm\d+)/){print $1}' > smid_sort.list

# 動画の詳細情報を取得し、csv書式に変換
./smi2csv.exe smid_sort.list > aviutl_script_ori.csv
