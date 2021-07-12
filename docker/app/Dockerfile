# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work \
&& apt-get update -y \
&& apt-get install -y mariadb-client \
&& go get github.com/labstack/echo/middleware \
&& go get github.com/ipfans/echo-session \
&& go get github.com/go-mysql-org/go-mysql

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work
