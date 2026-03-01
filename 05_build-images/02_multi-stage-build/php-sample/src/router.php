<?php

use Pecee\SimpleRouter\SimpleRouter;

// ビュー関数を使ってテンプレートを読み込む
SimpleRouter::get('/', function() {
    // テンプレートに渡すデータ
    $data = [
        'message' => 'Hello! Docker'
    ];

    // ビュー関数を使用してテンプレートを表示
    view('index', $data);
});

// ルーターの起動
SimpleRouter::start();

