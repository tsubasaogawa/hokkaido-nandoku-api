# Implementation Plan: 北海道の難読地名を取得するAPI

**Branch**: `001-api` | **Date**: 2025-09-25 | **Spec**: [./spec.md](./spec.md)
**Input**: Feature specification from `/specs/001-api/spec.md`

## Execution Flow (/plan command scope)
```
1. Load feature spec from Input path
   → If not found: ERROR "No feature spec at {path}"
2. Fill Technical Context (scan for NEEDS CLARIFICATION)
   → Detect Project Type from context (web=frontend+backend, mobile=app+api)
   → Set Structure Decision based on project type
3. Fill the Constitution Check section based on the content of the constitution document.
4. Evaluate Constitution Check section below
   → If violations exist: Document in Complexity Tracking
   → If no justification possible: ERROR "Simplify approach first"
   → Update Progress Tracking: Initial Constitution Check
5. Execute Phase 0 → research.md
   → If NEEDS CLARIFICATION remain: ERROR "Resolve unknowns"
6. Execute Phase 1 → contracts, data-model.md, quickstart.md, agent-specific template file (e.g., `CLAUDE.md` for Claude Code, `.github/copilot-instructions.md` for GitHub Copilot, `GEMINI.md` for Gemini CLI, `QWEN.md` for Qwen Code or `AGENTS.md` for opencode).
7. Re-evaluate Constitution Check section
   → If new violations: Refactor design, return to Phase 1
   → Update Progress Tracking: Post-Design Constitution Check
8. Plan Phase 2 → Describe task generation approach (DO NOT create tasks.md)
9. STOP - Ready for /tasks command
```

## Summary
API利用者が北海道の難読地名をランダムに取得できるAPIを開発する。技術スタックはGo言語とAWS Lambdaを採用し、サーバーレスアーキテクチャで構築する。データはアプリケーションに同梱されたCSVファイルから提供される。

## Technical Context
**Language/Version**: Go (1.22 or later)
**Primary Dependencies**: なし (標準ライブラリのみ)
**Storage**: 同梱のCSVファイル
**Testing**: `testing` 標準パッケージ
**Target Platform**: AWS Lambda (Container Image) + Amazon API Gateway
**Infrastructure as Code**: Terraform (with `terraform-aws-modules/lambda`)
**Project Type**: single
**Performance Goals**: p99で200ms未満のレスポンスタイム
**Constraints**: グローバルで毎秒1リクエストのレートリミット
**Scale/Scope**: 単一エンドポイントのシンプルなAPI

## Constitution Check
*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

憲法の原則に違反する点はない。シンプルさを重視し、外部依存を最小限に抑えるアプローチは原則に合致している。

## Project Structure

### Documentation (this feature)
```
specs/001-api/
├── plan.md              # This file (/plan command output)
├── research.md          # Phase 0 output (/plan command)
├── data-model.md        # Phase 1 output (/plan command)
├── quickstart.md        # Phase 1 output (/plan command)
├── contracts/
│   └── openapi.yaml     # Phase 1 output (/plan command)
└── tasks.md             # Phase 2 output (/tasks command - NOT created by /plan)
```

### Source Code (repository root)
```
# Option 1: Single project (DEFAULT)
cmd/
└── api/
    └── main.go
internal/
├── handler/
├── model/
└── repository/
pkg/
└── csvloader/

tests/
├── integration/
└── unit/

data/
└── nandoku_chimei.csv

terraform/
├── main.tf
├── variables.tf
└── outputs.tf

Dockerfile
go.mod
go.sum
```

**Structure Decision**: Option 1: Single project

## Phase 0: Outline & Research
調査は完了し、結果は `research.md` にまとめられている。技術選定に関する不明点はない。

**Output**: [research.md](./research.md)

## Phase 1: Design & Contracts
*Prerequisites: research.md complete*

設計と契約定義は完了した。データモデル、APIのOpenAPI仕様、およびローカルでの実行手順が定義されている。

**Output**:
- [data-model.md](./data-model.md)
- [contracts/openapi.yaml](./contracts/openapi.yaml)
- [quickstart.md](./quickstart.md)
- `GEMINI.md` (更新済み)

## Phase 2: Task Planning Approach
*This section describes what the /tasks command will do - DO NOT execute during /plan*

**Task Generation Strategy**:
- `tasks-template.md` をベースにタスクを生成する。
- `data-model.md` からモデル（struct）定義のタスクを作成する。
- CSVローダーパッケージ `csvloader` の実装タスクを作成する。
- `internal` パッケージ（repository, handler）の実装タスクを作成する。
- `cmd/api/main.go` の実装タスク（Lambdaハンドラのエントリポイント）を作成する。
- `Dockerfile` の作成タスクを生成する。
- `terraform/` ディレクトリに `terraform-aws-modules/lambda` を利用して、ECRリポジトリ作成、コンテナイメージのビルド、Lambda関数、API Gatewayのエンドポイントを含むAWSリソース一式を定義するTerraform設定ファイルの作成タスクを生成する。
- 各コンポーネントに対応するユニットテストの作成タスクを生成する。
- `quickstart.md` に基づくインテグレーションテストのタスクを作成する。

**Ordering Strategy**:
- TDDの順序: テストを実装より先に記述する。
- 依存関係の順序: model → csvloader → repository → handler → main
- [P] で並行実行可能なタスクをマークする。

**Estimated Output**: 約15-20の順序付けられたタスクリスト。

## Phase 3+: Future Implementation
*These phases are beyond the scope of the /plan command*

**Phase 3**: Task execution (/tasks command creates tasks.md)  
**Phase 4**: Implementation (execute tasks.md following constitutional principles)  
**Phase 5**: Validation (run tests, execute quickstart.md, performance validation)

## Complexity Tracking
*Fill ONLY if Constitution Check has violations that must be justified*

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| (なし) | - | - |


## Progress Tracking
*This checklist is updated during execution flow*

**Phase Status**:
- [x] Phase 0: Research complete (/plan command)
- [x] Phase 1: Design complete (/plan command)
- [x] Phase 2: Task planning complete (/plan command - describe approach only)
- [ ] Phase 3: Tasks generated (/tasks command)
- [ ] Phase 4: Implementation complete
- [ ] Phase 5: Validation passed

**Gate Status**:
- [x] Initial Constitution Check: PASS
- [x] Post-Design Constitution Check: PASS
- [x] All NEEDS CLARIFICATION resolved
- [ ] Complexity deviations documented

---
*Based on Constitution v2.1.1 - See `/.specify/memory/constitution.md`*