# Intelligent HPA ドキュメント構成計画

## 概要

このドキュメントは、Intelligent HPA プロジェクトの包括的なドキュメント体系の構成計画を示します。既存のドキュメントを保持しつつ、より詳細な仕様書を追加します。

## 現在のドキュメント構造

### 既存ドキュメント

```
docs/
├── architecture.md    # アーキテクチャ概要（保持）
├── developer.md       # 開発者向けガイド（保持）
├── estimator.md       # Estimator詳細（保持）
└── fittingjob.md      # FittingJob詳細（保持）

README.md              # プロジェクト概要（保持）
```

**既存ドキュメントの役割:**
- `architecture.md`: システム全体のアーキテクチャ概要、リソース間の関係
- `developer.md`: ビルド方法、開発環境のセットアップ
- `estimator.md`: Estimatorの動作原理、データ受信・送信の詳細
- `fittingjob.md`: FittingJobの機械学習実装、Prophet、変化点検知
- `README.md`: プロジェクト紹介、クイックスタート

## 新規ドキュメント構造

### 追加するドキュメント

```
docs/
├── templates/                          # テンプレート
│   ├── specification-template.md       # 仕様書テンプレート
│   └── api-specification-template.md   # API仕様書テンプレート
├── system-specification.md             # システム全体仕様書（新規）
├── api-specification.md                # API詳細仕様書（新規）
├── operations-guide.md                 # 運用ガイド（新規）
└── DOCUMENTATION_STYLE_GUIDE.md        # スタイルガイド（新規）
```

## 各ドキュメントの役割と内容

### 1. system-specification.md（システム全体仕様書）

**目的:** システム全体の包括的な仕様を提供

**対象読者:** システムアーキテクト、上級開発者、運用エンジニア

**主要セクション:**
1. システム概要
   - Intelligent HPAの目的と価値提案
   - 従来のHPAとの違い
   - ユースケース

2. アーキテクチャ詳細
   - システム全体構成
   - コンポーネント間の相互作用
   - データフロー
   - 実行フロー（IHPA作成から予測実行まで）

3. コントローラー実装詳細
   - IHPA Controller
     - Reconcilerの動作フロー
     - リソース生成ロジック
     - ライフサイクル管理
   - FittingJob Controller
     - CronJob管理
     - ConfigMap管理
     - スケジューリングロジック
   - Estimator Controller
     - ConfigMap監視
     - メトリクス送信ロジック
     - 調整メカニズム

4. 機械学習実装詳細
   - Prophet モデル
     - パラメータ設定
     - 訓練プロセス
     - 予測生成
   - 変化点検知
     - SSTアルゴリズム
     - パラメータチューニング
     - データフィルタリング

5. メトリクスプロバイダ統合
   - Datadog統合
     - API認証
     - メトリクス取得
     - メトリクス送信
     - Aggregator処理
   - Prometheus統合（将来）

6. データ管理
   - ConfigMapデータ構造
   - 予測データフォーマット
   - データマージロジック
   - データクリーンアップ

7. セキュリティとRBAC
   - RBAC実装
   - 権限分離
   - シークレット管理

8. エラーハンドリングと監視
   - エラー処理戦略
   - ログ出力
   - メトリクス露出
   - 可観測性

9. パフォーマンス考慮事項
   - スケーラビリティ
   - リソース使用量
   - 最適化手法

**既存ドキュメントとの関係:**
- `architecture.md`の内容を包含し、より詳細化
- `estimator.md`と`fittingjob.md`の内容を統合
- コントローラー実装の詳細を追加

### 2. api-specification.md（API詳細仕様書）

**目的:** すべてのカスタムリソースとAPIの完全な仕様を提供

**対象読者:** 開発者、API利用者、統合開発者

**主要セクション:**
1. API概要
   - APIバージョン（v1beta1, v1beta2）
   - バージョン間の違い
   - 互換性ポリシー

2. IntelligentHorizontalPodAutoscaler (IHPA)
   - APIバージョン: v1beta1, v1beta2
   - リソース構造
   - Specフィールド詳細
   - Statusフィールド詳細
   - バリデーションルール
   - デフォルト値
   - 使用例
   - アノテーションとラベル

3. FittingJob
   - APIバージョン: v1beta1, v1beta2
   - リソース構造
   - Specフィールド詳細
   - 設定パラメータ
   - CronJob生成仕様
   - ConfigMap構造
   - 使用例

4. Estimator
   - APIバージョン: v1beta2
   - リソース構造
   - Specフィールド詳細
   - 調整モード（raw/adjust）
   - ConfigMapデータフォーマット
   - 使用例

5. データモデル
   - 予測データCSVフォーマット
   - ConfigMapデータ構造
   - メトリクス命名規則
   - タグとラベル規則

6. RBAC リソース
   - ServiceAccount仕様
   - Role権限定義
   - RoleBinding構造

7. 生成されるリソース
   - HPA生成仕様
   - メトリクス注入ロジック
   - ラベルとセレクタ

8. バージョン移行ガイド
   - v1beta1からv1beta2への移行
   - 非推奨フィールド
   - 移行手順

**既存ドキュメントとの関係:**
- 既存ドキュメントに散在するAPI情報を集約
- より構造化された形式で提供

### 3. operations-guide.md（運用ガイド）

**目的:** 実運用に必要な情報を提供

**対象読者:** 運用エンジニア、SRE、システム管理者

**主要セクション:**
1. インストールガイド
   - 前提条件
   - Datadog Agent/Cluster Agentのセットアップ
   - IHPA Controllerのインストール
   - マニフェスト生成方法
   - バージョン確認

2. 設定ガイド
   - メトリクスプロバイダ設定
     - Datadog設定
     - API/APPキー管理
   - IHPA リソース設定
     - 基本設定
     - メトリクス設定
     - FittingJob設定
     - Estimator設定
   - パラメータチューニング
     - gapMinutes調整
     - 季節性設定
     - 変化点検知パラメータ

3. 運用手順
   - IHPAリソースの作成
   - FittingJobの手動実行
   - 予測データの確認
   - メトリクスの監視
   - リソースの更新
   - リソースの削除

4. モニタリング
   - ログの確認方法
   - メトリクスの確認
   - HPA describeコマンド
   - 予測精度の確認
   - 異常検知

5. トラブルシューティング
   - 一般的な問題と解決方法
     - FittingJobが実行されない
     - 予測メトリクスが送信されない
     - HPAがスケールしない
     - メトリクスプロバイダ接続エラー
     - ConfigMapサイズ超過
     - 権限エラー
   - ログ分析
   - デバッグ手順

6. パフォーマンスチューニング
   - 変化点検知パラメータの調整
   - 予測精度向上のベストプラクティス
   - リソース使用量の最適化

7. アップグレードガイド
   - バージョンアップ手順
   - ダウンタイム最小化
   - ロールバック手順

8. ベストプラクティス
   - メトリクス選択
   - しきい値設定
   - スケーリング戦略
   - 本番環境での推奨設定

**既存ドキュメントとの関係:**
- `README.md`のUsageセクションを拡張
- `developer.md`の運用関連部分を抽出・拡張
- より実践的な内容に焦点

### 4. DOCUMENTATION_STYLE_GUIDE.md（スタイルガイド）

**目的:** ドキュメント作成の一貫性を保つ

**対象読者:** ドキュメント作成者、コントリビューター

**内容:**
- 言語とフォーマット規則
- ファイル構成
- セクション構造
- コードブロック
- 図表の使用方法
- 用語の統一
- レビューチェックリスト

### 5. templates/（テンプレート）

**目的:** 新規ドキュメント作成時の雛形を提供

**内容:**
- `specification-template.md`: 仕様書テンプレート
- `api-specification-template.md`: API仕様書テンプレート

## ドキュメント間の関係

```
README.md (エントリーポイント)
    ├─→ docs/system-specification.md (詳細仕様)
    │       ├─→ docs/architecture.md (既存: アーキテクチャ概要)
    │       ├─→ docs/estimator.md (既存: Estimator詳細)
    │       └─→ docs/fittingjob.md (既存: FittingJob詳細)
    │
    ├─→ docs/api-specification.md (API仕様)
    │
    ├─→ docs/operations-guide.md (運用ガイド)
    │
    └─→ docs/developer.md (既存: 開発者ガイド)
```

## 実装計画

### フェーズ1: 基盤準備（完了）
- [x] ドキュメント構造の分析
- [x] テンプレート作成
- [x] スタイルガイド作成
- [x] 構成計画の策定

### フェーズ2: コントローラー実装の文書化
- [ ] IHPA Controller実装詳細の分析と文書化
- [ ] FittingJob Controller実装詳細の分析と文書化
- [ ] Estimator Controller実装詳細の分析と文書化

### フェーズ3: 機械学習実装の文書化
- [ ] Prophet モデル実装の分析と文書化
- [ ] 変化点検知実装の分析と文書化
- [ ] メトリクスプロバイダ統合の分析と文書化

### フェーズ4: API仕様の文書化
- [ ] カスタムリソース定義の詳細分析
- [ ] データフォーマットと命名規則の文書化

### フェーズ5: セキュリティとRBACの文書化
- [ ] RBAC リソース生成ロジックの分析
- [ ] セキュリティベストプラクティスの文書化

### フェーズ6: エラーハンドリングと監視の文書化
- [ ] エラーハンドリング戦略の分析
- [ ] ログ出力とモニタリングの分析

### フェーズ7: 運用ガイドの作成
- [ ] インストールと設定ガイドの作成
- [ ] トラブルシューティングガイドの作成
- [ ] パフォーマンスチューニングガイドの作成

### フェーズ8: 統合とレビュー
- [ ] すべてのドキュメントの統合
- [ ] 一貫性の確認
- [ ] 既存ドキュメントとの整合性確認

## 既存ドキュメントの扱い

### 保持するドキュメント

以下のドキュメントは**そのまま保持**します：

1. **architecture.md**
   - 理由: アーキテクチャ概要として有用
   - 新規ドキュメントとの関係: system-specification.mdから参照

2. **developer.md**
   - 理由: 開発環境セットアップに特化した内容
   - 新規ドキュメントとの関係: 独立して保持

3. **estimator.md**
   - 理由: Estimatorの詳細な動作原理を説明
   - 新規ドキュメントとの関係: system-specification.mdから参照

4. **fittingjob.md**
   - 理由: 機械学習実装の詳細を説明
   - 新規ドキュメントとの関係: system-specification.mdから参照

5. **README.md**
   - 理由: プロジェクトのエントリーポイント
   - 新規ドキュメントとの関係: 新規ドキュメントへのリンクを追加

### 重複の回避

新規ドキュメントは既存ドキュメントの内容を**参照**する形で記述し、重複を避けます：

```markdown
## Estimatorの動作詳細

Estimatorの詳細な動作については、[Estimator詳細ドキュメント](./estimator.md)を参照してください。

ここでは、システム全体の文脈におけるEstimatorの役割を説明します...
```

## 成果物

このタスク完了時の成果物：

1. ✅ `docs/templates/specification-template.md`
2. ✅ `docs/templates/api-specification-template.md`
3. ✅ `docs/DOCUMENTATION_STYLE_GUIDE.md`
4. ✅ `docs/DOCUMENTATION_PLAN.md`（このファイル）

次のタスクで作成される成果物：

5. `docs/system-specification.md`
6. `docs/api-specification.md`
7. `docs/operations-guide.md`

## 参考資料

- 既存ドキュメント: `docs/architecture.md`, `docs/estimator.md`, `docs/fittingjob.md`, `docs/developer.md`
- 要件定義書: `.kiro/specs/intelligent-hpa-documentation/requirements.md`
- 設計書: `.kiro/specs/intelligent-hpa-documentation/design.md`
