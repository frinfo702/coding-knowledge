# Programming Knowledge DeepWiki Seed

このレポジトリは **「プログラミングそのもの」** に関する巨大ナレッジを Markdown と実践コードで蓄積し、  
[deepwiki](https://deepwiki.example/) にインポートして AI が構造化 Wiki を生成できるようにするためのスターターです。

## ディレクトリ構成

```
.
├── docs/            # 各言語・トピックの解説 Markdown
│   ├── go/
│   │   ├── hooks.md
│   │   └── errors.md
│   └── python/
│       ├── hooks.md
│       └── errors.md
└── examples/        # 実践コード（.go, .py など）
    ├── go/
    │   ├── hook_server/
    │   │   ├── main.go
    │   │   ├── server.go
    │   │   └── server_test.go
    │   └── error_examples/
    │       └── error_wrap.go
    └── python/
        ├── hooks/
        │   └── plugin_system.py
        └── errors/
            └── custom_error.py
```

deepwiki は `docs/` と `examples/` の両方を解析します。Markdown には **目的 / 仕組み / サンプル** を、  
code には **最小実装 / テスト / ベンチ** を入れることで、AI が依存関係と役割を推定しやすくなります。

---
