# Intelligent HPA ドキュメンテーションスタイルガイド

## 目的

このスタイルガイドは、Intelligent HPA プロジェクトの全ドキュメントにおいて一貫性のある、読みやすく、保守しやすい文書を作成するための基準を定めます。

## 言語

- **主要言語**: 日本語
- **コード例**: 英語（変数名、コメント等）
- **技術用語**: 英語の専門用語はそのまま使用（例: Controller, ConfigMap, CronJob）

## ファイル構成

### ディレクトリ構造

```
docs/
├── templates/              # ドキュメントテンプレート
│   ├── specification-template.md
│   └── api-specification-template.md
├── architecture.md         # 既存: アーキテクチャ概要
├── developer.md           # 既存: 開発者向けガイド
├── estimator.md           # 既存: Estimator詳細
├── fittingjob.md          # 既存: FittingJob詳細
├── system-specification.md # 新規: システム全体仕様書
├── api-specification.md    # 新規: API詳細仕様書
├── operations-guide.md     # 新規: 運用ガイド
└── DOCUMENTATION_STYLE_GUIDE.md # このファイル
```

### ファイル命名規則

- 小文字とハイフンを使用: `system-specification.md`
- 日本語ファイル名は避ける
- 内容を明確に表す名前を使用

## ドキュメント構造

### 必須セクション

すべての仕様書には以下のセクションを含める：

1. **概要** - ドキュメントの目的と対象範囲
2. **アーキテクチャ/構造** - システムの構成
3. **実装詳細** - 具体的な実装内容
4. **設定/パラメータ** - 設定可能な項目
5. **運用ガイド** - 使用方法とトラブルシューティング
6. **参考資料** - 関連ドキュメントへのリンク

### セクション階層

```markdown
# タイトル (H1) - ドキュメントタイトルのみ

## メインセクション (H2)

### サブセクション (H3)

#### 詳細項目 (H4)
```

- H1は1つのみ（ドキュメントタイトル）
- H2以降で構造化
- 5階層以上のネストは避ける

## 書式規則

### コードブロック

言語を明示的に指定：

```markdown
\```yaml
apiVersion: v1
kind: ConfigMap
\```

\```go
func main() {
    // code
}
\```

\```python
def train_model():
    pass
\```
```

### テーブル

- ヘッダー行を必ず含める
- 列の整列は左揃え（`:--`）または中央揃え（`:-:`）
- 複雑な内容は箇条書きで記述

```markdown
| フィールド名 | 型 | 必須 | デフォルト値 | 説明 |
|------------|-----|------|------------|------|
| name | string | Yes | - | リソース名 |
| replicas | int | No | 1 | レプリカ数 |
```

### リスト

**箇条書き:**
```markdown
- 項目1
- 項目2
  - サブ項目2.1
  - サブ項目2.2
```

**番号付きリスト:**
```markdown
1. 手順1
2. 手順2
3. 手順3
```

### 強調

- **太字**: 重要な用語、キーワード
- *イタリック*: 軽い強調（日本語では使用頻度低）
- `コード`: コマンド、変数名、ファイル名、技術用語

### リンク

```markdown
[リンクテキスト](./relative-path.md)
[外部リンク](https://example.com)
```

## 図表

### Mermaid図

アーキテクチャやフローチャートにはMermaidを使用：

```markdown
\```mermaid
graph TB
    A[開始] --> B[処理]
    B --> C[終了]
\```
```

### 画像

```markdown
![代替テキスト](../misc/image-name.svg)
```

- SVG形式を優先
- 画像は`misc/`ディレクトリに配置
- 代替テキストを必ず記述

## 用語の統一

### カスタムリソース

- IntelligentHorizontalPodAutoscaler (IHPA)
- FittingJob
- Estimator

### Kubernetesリソース

- HorizontalPodAutoscaler (HPA)
- ConfigMap
- CronJob
- ServiceAccount
- Role
- RoleBinding

### コンポーネント

- IHPA Controller
- FittingJob Controller
- Estimator Controller
- MetricProvider

### 技術用語

- 時系列予測 (Time Series Prediction)
- 変化点検知 (Change Point Detection)
- Prophet モデル
- SST (Singular Spectrum Transformation)

## コード例

### YAML マニフェスト

- インデント: 2スペース
- コメントで重要な設定を説明
- 実際に動作する完全な例を提供

```yaml
apiVersion: ihpa.ake.cyberagent.co.jp/v1beta2
kind: IntelligentHorizontalPodAutoscaler
metadata:
  name: example
  namespace: default
spec:
  # メトリクスプロバイダの設定
  metricProvider:
    name: datadog
    datadog:
      apikey: xxx
      appkey: yyy
```

### Go コード

- Go標準のフォーマット規則に従う
- 重要な部分にコメントを追加

```go
// IHPAReconciler は IHPA リソースを調整します
type IHPAReconciler struct {
    client.Client
    Scheme *runtime.Scheme
}
```

### Python コード

- PEP 8に従う
- 関数とクラスにdocstringを追加

```python
def train_model(data: pd.DataFrame) -> Prophet:
    """
    Prophetモデルを訓練します。
    
    Args:
        data: 訓練データ
        
    Returns:
        訓練済みモデル
    """
    model = Prophet()
    model.fit(data)
    return model
```

## API仕様の記述

### フィールド定義

テーブル形式で記述：

| フィールド | 型 | 必須 | デフォルト | 説明 | 制約 |
|---------|-----|------|----------|------|------|
| name | string | Yes | - | リソース名 | 1-253文字 |
| replicas | int | No | 1 | レプリカ数 | 1以上 |

### 使用例

必ず動作する完全な例を提供：

```yaml
# 完全なマニフェスト例
apiVersion: ihpa.ake.cyberagent.co.jp/v1beta2
kind: IntelligentHorizontalPodAutoscaler
metadata:
  name: nginx
  namespace: default
spec:
  # ... 完全な設定
```

## トラブルシューティング

問題解決セクションは以下の形式で記述：

```markdown
#### 問題: [問題の簡潔な説明]

**症状:**
- 症状1
- 症状2

**原因:**
[原因の説明]

**解決方法:**
1. 手順1
2. 手順2
3. 確認方法

**関連ログ:**
\```
エラーログの例
\```
```

## バージョン管理

### ドキュメントバージョン

- ドキュメントの最終更新日を記載
- 大きな変更時はバージョン番号を付与

```markdown
---
version: 1.0.0
last_updated: 2024-01-15
---
```

### API バージョン

- 対応するAPIバージョンを明記
- v1beta1とv1beta2の違いを明確に

## レビューチェックリスト

ドキュメント作成後、以下を確認：

- [ ] タイトルと概要が明確
- [ ] セクション構造が適切
- [ ] コード例が動作する
- [ ] 用語が統一されている
- [ ] リンクが正しい
- [ ] 図表が適切に配置されている
- [ ] 誤字脱字がない
- [ ] 技術的に正確
- [ ] 対象読者に適した内容

## 参考資料

- [Markdown Guide](https://www.markdownguide.org/)
- [Mermaid Documentation](https://mermaid-js.github.io/)
- [Kubernetes Documentation Style Guide](https://kubernetes.io/docs/contribute/style/style-guide/)
