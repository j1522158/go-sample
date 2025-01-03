#!/bin/bash

# サーバーにアップロード
scp -r ./release/* user@server:/path/to/deploy

# リモートサーバーでデプロイコマンドを実行
ssh user@server << EOF
cd /path/to/deploy
./bin/app &
EOF
