<?php

function view($name, $data = [])
{
    // htmlspecialcharsを使って配列内のすべてのデータをエスケープする
    $escapedData = array_map(function($item) {
        if (is_string($item)) {
            return htmlspecialchars($item, ENT_QUOTES, 'UTF-8');
        }
        return $item;
    }, $data);

    // 配列のキーを変数として展開（$escapedDataに変更）
    extract($escapedData);

    // テンプレートファイルを読み込む
    include sprintf('%s/resources/%s.tpl.php', __DIR__, $name);
}

