# API仕様書

## 概要

[API仕様書の目的と対象範囲]

## カスタムリソース定義 (CRD)

### [リソース名]

#### APIバージョン

- Group: `[group]`
- Version: `[version]`
- Kind: `[Kind]`

#### リソース構造

```yaml
apiVersion: [group]/[version]
kind: [Kind]
metadata:
  name: [name]
  namespace: [namespace]
spec:
  # 仕様フィールド
status:
  # ステータスフィールド
```

#### Spec フィールド

| フィールド | 型 | 必須 | デフォルト | 説明 | 制約 |
|---------|-----|------|----------|------|------|
| field1 | string | Yes | - | 説明 | 制約条件 |
| field2 | object | No | {} | 説明 | - |

##### [ネストされたオブジェクト名]

| フィールド | 型 | 必須 | デフォルト | 説明 |
|---------|-----|------|----------|------|
| subfield1 | string | Yes | - | 説明 |

#### Status フィールド

| フィールド | 型 | 説明 |
|---------|-----|------|
| conditions | []Condition | リソースの状態 |
| phase | string | 現在のフェーズ |

#### 使用例

```yaml
apiVersion: example.com/v1
kind: Example
metadata:
  name: example-instance
  namespace: default
spec:
  field1: "value"
  field2:
    subfield1: "subvalue"
```

#### バリデーション

- [バリデーションルール1]
- [バリデーションルール2]

#### デフォルト値の適用

- [デフォルト値の適用ロジック]

## データモデル

### [モデル名]

**構造:**

```go
type ModelName struct {
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
}
```

**説明:**
[モデルの説明]

## メトリクス命名規則

### 予測メトリクス

- プレフィックス: `[prefix]`
- 命名パターン: `[pattern]`

**例:**
- 元のメトリクス: `original.metric.name`
- 予測メトリクス: `prefix.original.metric.name`

### メトリクスタイプ

| タイプ | 説明 | 命名規則 |
|------|------|---------|
| Resource | Kubernetesリソースメトリクス | 規則1 |
| External | 外部メトリクス | 規則2 |

## ConfigMap データフォーマット

### [ConfigMap名]

**用途:** [用途の説明]

**データ構造:**

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: [name]
data:
  key1: |
    [データ内容]
```

**フォーマット仕様:**

[フォーマットの詳細説明]

## アノテーションとラベル

### アノテーション

| キー | 値の型 | 説明 | 例 |
|-----|--------|------|-----|
| annotation.key | string | 説明 | "value" |

### ラベル

| キー | 値の型 | 説明 | 例 |
|-----|--------|------|-----|
| label.key | string | 説明 | "value" |

## RBAC リソース

### ServiceAccount

[ServiceAccountの説明]

### Role

**権限:**

| APIGroup | Resources | Verbs |
|----------|-----------|-------|
| "" | configmaps | get, list, watch, create, update |

### RoleBinding

[RoleBindingの説明]

## バージョン互換性

### v1beta1 → v1beta2

**変更点:**
- [変更1]
- [変更2]

**移行ガイド:**
1. [手順1]
2. [手順2]

## 参考資料

- [関連ドキュメント]
