# Programming Knowledge DeepWiki Seed

このレポジトリはdeepwikiにプログラミングの設計や特定のプログラミング言語技術を超えた抽象的な問題を解決する知識や技術をまとめさせて、
プログラミングで困ったときに見れば問題が解決するようにするためのものです。

# Instruction to　AI(LLM)
wikiを作成する際、そのwikiは日本語で作成するようにしてください。
文章は全て日本語で作成してください
多言語の使用は強く禁止します。(**Never use foreign languages other than Japanese!!**)

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
