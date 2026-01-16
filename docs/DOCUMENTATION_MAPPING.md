# ドキュメントマッピング

## 概要

このドキュメントは、既存のドキュメント内容と新規作成するドキュメントの対応関係を示します。これにより、情報の重複を避け、適切な参照関係を構築できます。

## 既存ドキュメントの内容分析

### README.md

**現在の内容:**
- プロジェクト概要
- 前提条件（Datadog Agentのインストール）
- 使用方法（IHPAマニフェストの説明）
- インストール方法
- FittingJobの説明
- 他のドキュメントへのリンク

**新規ドキュメントへのマッピング:**
- プロジェクト概要 → **保持**（エントリーポイントとして）
- 前提条件 → `operations-guide.md`で詳細化
- 使用方法 → `operations-guide.md`で詳細化
- インストール方法 → `operations-guide.md`で詳細化
- FittingJobの説明 → `system-specification.md`で詳細化

### docs/architecture.md

**現在の内容:**
- システム全体のアーキテクチャ図
- IHPAリソースの展開方法
- リソース間の関係
- HPA生成の仕組み
- メトリクス送信の流れ
- EstimatorControllerの動作フロー
- FittingJobの動作フロー

**新規ドキュメントへのマッピング:**
- システム全体のアーキテクチャ → `system-specification.md`の「アーキテクチャ詳細」セクションで参照
- リソース間の関係 → `system-specification.md`で詳細化
- HPA生成の仕組み → `api-specification.md`で仕様化
- メトリクス送信の流れ → `system-specification.md`で詳細化
- コントローラーの動作フロー → `system-specification.md`で詳細化

**保持理由:**
- 視覚的なアーキテクチャ概要として有用
- 新規ドキュメントから参照される

### docs/estimator.md

**現在の内容:**
- Estimatorの概要
- データ受信処理
  - CSV解析
  - タイムスタンプ調整（gapMinutes）
  - データマージロジック
- データ送信処理
  - 調整モード（none/adjust）
  - 調整計算式
  - yhat_upper/lowerの扱い
- Datadog固有の実装
  - Aggregator（avg, sum, max, min）
  - HPAのアルゴリズム
  - 合計値での学習の重要性

**新規ドキュメントへのマッピング:**
- Estimatorの概要 → `system-specification.md`の「Estimator Controller」セクション
- データ受信処理 → `system-specification.md`で詳細化
- データ送信処理 → `system-specification.md`で詳細化
- 調整メカニズム → `system-specification.md`で詳細化、`api-specification.md`で仕様化
- Datadog固有の実装 → `system-specification.md`の「メトリクスプロバイダ統合」セクション

**保持理由:**
- Estimatorの動作原理を深く説明
- 数式や詳細なロジックを含む
- 新規ドキュメントから参照される

### docs/fittingjob.md

**現在の内容:**
- FittingJobの概要（Pythonプログラム）
- 時系列予測
  - Prophetの使用
  - パラメータ説明（seasonality_mode, growth, interval_width）
- 変化点検知
  - 特異スペクトル解析法（SST）
  - SSTのアルゴリズム詳細
  - パラメータ説明
- 参考文献

**新規ドキュメントへのマッピング:**
- FittingJobの概要 → `system-specification.md`の「機械学習実装詳細」セクション
- Prophetモデル → `system-specification.md`で詳細化
- 変化点検知 → `system-specification.md`で詳細化
- パラメータ → `api-specification.md`で仕様化、`operations-guide.md`でチューニング方法を説明

**保持理由:**
- 機械学習アルゴリズムの詳細な説明
- 学術的な背景を含む
- 新規ドキュメントから参照される

### docs/developer.md

**現在の内容:**
- マニフェスト生成方法
- Kubernetes 1.18のValidation問題と回避方法
- FittingJobのビルド方法
- FittingJobの開発環境セットアップ
- 予測データのCSVフォーマット
- IHPA Controllerのビルド方法
- Kubebuilderのscaffold生成コマンド

**新規ドキュメントへのマッピング:**
- マニフェスト生成方法 → **保持**（開発者向け情報）
- ビルド方法 → **保持**（開発者向け情報）
- 開発環境セットアップ → **保持**（開発者向け情報）
- CSVフォーマット → `api-specification.md`で仕様化

**保持理由:**
- 開発者向けの技術情報に特化
- ビルドやセットアップの手順
- 新規ドキュメントとは異なる対象読者

## 新規ドキュメントの内容ソース

### system-specification.md

**情報ソース:**

| セクション | 既存ドキュメント | コードベース | 新規分析 |
|---------|--------------|------------|---------|
| システム概要 | README.md | - | ✓ |
| アーキテクチャ詳細 | architecture.md | ihpa-controller/ | ✓ |
| IHPA Controller | architecture.md | ihpa-controller/controllers/intelligenthorizontalpodautoscaler_* | ✓ |
| FittingJob Controller | architecture.md | ihpa-controller/controllers/fittingjob_* | ✓ |
| Estimator Controller | architecture.md, estimator.md | ihpa-controller/controllers/estimator_* | ✓ |
| Prophet モデル | fittingjob.md | fittingjob/fittingjob/model.py | ✓ |
| 変化点検知 | fittingjob.md | fittingjob/fittingjob/model.py | ✓ |
| Datadog統合 | estimator.md | fittingjob/fittingjob/datadog.py, ihpa-controller/controllers/metricprovider/datadog/ | ✓ |
| データ管理 | estimator.md, developer.md | - | ✓ |
| RBAC | - | ihpa-controller/controllers/intelligenthorizontalpodautoscaler_generator_impl.go | ✓ |
| エラーハンドリング | - | ihpa-controller/controllers/, fittingjob/ | ✓ |

### api-specification.md

**情報ソース:**

| セクション | 既存ドキュメント | コードベース | 新規分析 |
|---------|--------------|------------|---------|
| API概要 | - | ihpa-controller/api/ | ✓ |
| IHPA仕様 | README.md, architecture.md | ihpa-controller/api/v1beta2/intelligenthorizontalpodautoscaler_types.go | ✓ |
| FittingJob仕様 | README.md | ihpa-controller/api/v1beta2/fittingjob_types.go | ✓ |
| Estimator仕様 | - | ihpa-controller/api/v1beta2/estimator_types.go | ✓ |
| データモデル | developer.md, estimator.md | - | ✓ |
| メトリクス命名規則 | architecture.md, estimator.md | ihpa-controller/controllers/ | ✓ |
| RBAC仕様 | - | ihpa-controller/controllers/ | ✓ |
| バージョン移行 | - | ihpa-controller/api/v1beta1/, ihpa-controller/api/v1beta2/ | ✓ |

### operations-guide.md

**情報ソース:**

| セクション | 既存ドキュメント | コードベース | 新規分析 |
|---------|--------------|------------|---------|
| インストールガイド | README.md | manifests/ | ✓ |
| 設定ガイド | README.md | - | ✓ |
| 運用手順 | README.md | - | ✓ |
| モニタリング | - | - | ✓ |
| トラブルシューティング | - | - | ✓ |
| パフォーマンスチューニング | fittingjob.md, README.md | - | ✓ |
| ベストプラクティス | - | - | ✓ |

## 参照関係

### system-specification.md からの参照

```markdown
## Estimatorの詳細な動作原理

Estimatorの数学的な調整メカニズムや、タイムスタンプ調整の詳細については、
[Estimator詳細ドキュメント](./estimator.md)を参照してください。

ここでは、システム全体におけるEstimatorの役割と、
他のコンポーネントとの相互作用について説明します。
```

### api-specification.md からの参照

```markdown
## FittingJob パラメータの詳細

変化点検知パラメータの理論的背景とチューニング方法については、
[FittingJob詳細ドキュメント](./fittingjob.md)を参照してください。
```

### operations-guide.md からの参照

```markdown
## システムアーキテクチャの理解

運用を開始する前に、システムの全体像を理解することを推奨します。
詳細は以下のドキュメントを参照してください：

- [システム全体仕様書](./system-specification.md)
- [アーキテクチャ概要](./architecture.md)
```

## 重複回避の原則

### 原則1: 詳細度による分離

- **既存ドキュメント**: 深い技術的詳細、アルゴリズム、数式
- **新規ドキュメント**: システム全体の文脈、統合的な視点、実装詳細

### 原則2: 対象読者による分離

- **architecture.md, estimator.md, fittingjob.md**: 技術的に深く理解したい読者
- **system-specification.md**: システム全体を理解したい読者
- **api-specification.md**: API利用者、統合開発者
- **operations-guide.md**: 運用エンジニア
- **developer.md**: 開発者

### 原則3: 参照による統合

- 新規ドキュメントは既存ドキュメントを**参照**する
- 既存ドキュメントの内容を**コピー**しない
- 必要に応じて**要約**を提供し、詳細は既存ドキュメントへ誘導

## 実装時の注意事項

### 新規ドキュメント作成時

1. **既存ドキュメントを確認**: 同じ内容が既に存在しないか確認
2. **参照を追加**: 既存ドキュメントへの適切な参照を追加
3. **要約を提供**: 必要に応じて簡潔な要約を提供
4. **新しい視点を追加**: 既存ドキュメントにない視点や情報を追加

### 既存ドキュメントの更新

既存ドキュメントは**原則として変更しない**が、以下の場合は更新を検討：

1. **明らかな誤り**: 技術的な誤りの修正
2. **リンク追加**: 新規ドキュメントへのリンク追加
3. **非推奨の明示**: 古い情報の非推奨マーク

## まとめ

このマッピングにより：

1. ✅ 既存ドキュメントの価値を保持
2. ✅ 情報の重複を回避
3. ✅ 適切な参照関係を構築
4. ✅ 各ドキュメントの役割を明確化
5. ✅ 対象読者に応じた情報提供

新規ドキュメント作成時は、このマッピングを参照して、適切な内容と参照関係を構築してください。
