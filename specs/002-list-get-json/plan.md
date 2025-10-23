# Implementation Plan: /list エンドポイントの追加

**Branch**: `002-list-get-json` | **Date**: 2025-10-23 | **Spec**: [link](./spec.md)
**Input**: Feature specification from `/specs/002-list-get-json/spec.md`

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

**IMPORTANT**: The /plan command STOPS at step 7. Phases 2-4 are executed by other commands:
- Phase 2: /tasks command creates tasks.md
- Phase 3-4: Implementation execution (manual or via tools)

## Summary
APIクライアントが北海道の難読地名とその読み仮名の一覧をリスト形式のJSONで一度に取得できるように、`GET /list` エンドポイントを実装する。データは既存のCSVファイルから読み込み、Goの標準ライブラリを用いてAPIを提供する。

## Technical Context
**Language/Version**: Go 1.21
**Primary Dependencies**: net/http, encoding/csv
**Storage**: `source/data/nandoku_chimei.csv` (in-memory)
**Testing**: Go standard testing library
**Target Platform**: Linux server (Docker container)
**Project Type**: single
**Performance Goals**: N/A
**Constraints**: N/A
**Scale/Scope**: ~100 records

## Constitution Check
*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **TDD mandatory**: PASS (tests will be written first)
- **Simplicity**: PASS (no unnecessary complexity)

## Project Structure

### Documentation (this feature)
```
specs/002-list-get-json/
├── plan.md              # This file (/plan command output)
├── research.md          # Phase 0 output (/plan command)
├── data-model.md        # Phase 1 output (/plan command)
├── quickstart.md        # Phase 1 output (/plan command)
├── contracts/           # Phase 1 output (/plan command)
│   └── openapi.yaml
└── tasks.md             # Phase 2 output (/tasks command - NOT created by /plan)
```

### Source Code (repository root)
```
# Option 1: Single project (DEFAULT)
source/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── model/
│   └── repository/
└── pkg/
```
**Structure Decision**: Option 1: Single project

## Phase 0: Outline & Research
1. **Extract unknowns from Technical Context** above:
   - All clear.

2. **Generate and dispatch research agents**:
   - No research needed.

3. **Consolidate findings** in `research.md`:
   - Done.

**Output**: research.md

## Phase 1: Design & Contracts
*Prerequisites: research.md complete*

1. **Extract entities from feature spec** → `data-model.md`:
   - Done.

2. **Generate API contracts** from functional requirements:
   - Done: `contracts/openapi.yaml`

3. **Generate contract tests** from contracts:
   - To be done in implementation phase.

4. **Extract test scenarios** from user stories:
   - Done: `quickstart.md`

5. **Update agent file incrementally**:
   - Not applicable for this feature.

**Output**: data-model.md, /contracts/*, quickstart.md

## Phase 2: Task Planning Approach
*This section describes what the /tasks command will do - DO NOT execute during /plan*

**Task Generation Strategy**:
- Load `.specify/templates/tasks-template.md` as base
- Generate tasks from Phase 1 design docs (contracts, data model, quickstart)
- Each contract → contract test task [P]
- Each entity → model creation task [P]
- Each user story → integration test task
- Implementation tasks to make tests pass

**Ordering Strategy**:
- TDD order: Tests before implementation
- Dependency order: Models before services before UI
- Mark [P] for parallel execution (independent files)

**Estimated Output**: 5-7 numbered, ordered tasks in tasks.md

**IMPORTANT**: This phase is executed by the /tasks command, NOT by /plan

## Phase 3+: Future Implementation
*These phases are beyond the scope of the /plan command*

**Phase 3**: Task execution (/tasks command creates tasks.md)
**Phase 4**: Implementation (execute tasks.md following constitutional principles)
**Phase 5**: Validation (run tests, execute quickstart.md, performance validation)

## Complexity Tracking
*Fill ONLY if Constitution Check has violations that must be justified*

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| N/A       | N/A        | N/A                                 |


## Progress Tracking
*This checklist is updated during execution flow*

**Phase Status**:
- [X] Phase 0: Research complete (/plan command)
- [X] Phase 1: Design complete (/plan command)
- [X] Phase 2: Task planning complete (/plan command - describe approach only)
- [ ] Phase 3: Tasks generated (/tasks command)
- [ ] Phase 4: Implementation complete
- [ ] Phase 5: Validation passed

**Gate Status**:
- [X] Initial Constitution Check: PASS
- [X] Post-Design Constitution Check: PASS
- [X] All NEEDS CLARIFICATION resolved
- [X] Complexity deviations documented

---
*Based on Constitution v2.1.1 - See `/memory/constitution.md`*
