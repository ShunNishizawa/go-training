# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# boiler-plateディレクトリの作成
# RUN mkdir /go/src/github.com/boiler-plate
# ワーキングディレクトリの設定
WORKDIR /go/src/github.com/training
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/training
# パッケージのインポート
RUN go install golang.org/x/tools/cmd/goimports@latest