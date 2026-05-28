---
name: tui-refactor
description: "Evidence-based refactoring — identifies violations via tui-audit, proposes fixes based on empirical patterns with confidence scores and evidence from similar repos. Shows 'X repos fixed this by doing Y'. Generates patches: god model decomposition, storage Cmd wrapping, viewport diffing, ticker cleanup. References pattern evidence and fix commits. Dry-run by default. Trigger when: user says 'refactor', 'fix this', 'apply pattern', 'too complex', or when audit findings exist."
argument-hint: "[path-to-project] [--apply-pattern <id>] [--min-confidence 0.80] [--dry-run] [--apply]"
allowed-tools: [Read, Write, Edit, Bash, Grep, Glob, Agent, ATree, mcp__charmed__extract_evidence, mcp__charmed__list_patterns, mcp__charmed__get_fixes]
---

# tui-refactor — Architectural Refactoring Engine

Produces safe, minimal structural patches for Bubble Tea projects. Every patch is verified against the IR and ATree impact analysis before application. Dry-run by default.

## How It Works

1. **Scan** — Calls `charmed-mcp scan_project` to produce the TUI Structural IR
2. **Evaluate** — Runs constraint enforcement (equivalent to tui-audit) against the IR
3. **Classify** — Maps each violation to a refactoring action
4. **Impact** — Uses IR data (and ATree when available) to assess cross-file impact
5. **Patch** — Generates targeted diffs that fix the violation while preserving external interface
6. **Verify** — Ensures patched code compiles and doesn't introduce new violations

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Project path | Yes | User argument (default: current directory) |
| Target symbol | No | `--target ModelName` (refactor specific model) |
| Dry run | Yes | Default on, `--apply` to write |

## Execution Pipeline

### Phase 1: IR-Enforced Audit

调用 `scan_project` MCP 工具获取 IR，然后评估所有约束规则（与 tui-audit Phase 2-3 相同）。

从 IR 中提取违规列表，包含每个违规的 rule_id、severity、file、line、message、fix。

### Phase 2: Refactor Classification

将每条违规映射到具体的结构重构动作：

| Violation | Action | Description |
|---|---|---|
| `god_model` (>300 LOC Update) | `split_model` | 提取子模型，父模型协调消息路由 |
| `giga_switch_update` (>20 branches) | `split_model` | 按消息处理模式分组字段 |
| `viewport_recreation` | `stabilize_viewport` | 移至 Init()，添加内容 diff 检查 |
| `viewport_churn` | `stabilize_viewport` | 在 SetContent 前添加 diff 检查 |
| `blocking_update` | `extract_command_boundary` | 将阻塞 I/O 包装为 tea.Cmd |
| `unbounded_cmd_spawning` | `throttle_commands` | 添加 tick/debounce/single-cmd 合并 |
| `viewport_without_resize_handling` | `add_resize_handler` | 传播 WindowSizeMsg 到 viewport |
| `style_recomputation_hot_path` | `cache_styles` | 移至结构体字段或计算属性 |
| `string_allocation_pressure` | `optimize_renders` | 使用 strings.Builder，缓存静态字符串 |

仅处理 CRITICAL 和 ERROR 级别的违规。WARNING 和 INFO 级别的建议单独列出但不自动修复。

### Phase 3: Impact Analysis (from IR)

对于每个计划的重构动作，从 IR 数据评估影响：

```
split_model impact analysis:
  affected_models: [ListModel, DetailModel, StatusModel]  (from IR models[])
  affected_messages: [12 messages routed to children]      (from IR messages[])
  affected_bindings: [viewport, list, textarea]            (from IR component_bindings[])
  state_splits: {
    AppModel → ListModel: [items, filterState, selected],
    AppModel → DetailModel: [viewport, content],
    AppModel → StatusModel: [spinner, status]
  }
  external_interface_preserved: true  (AppModel struct name, exported methods preserved)
```

如果影响范围超过 10 个符号，在继续前警告用户。

### Phase 4: Patch Generation

每个动作生成一个精确的 patch。示例：

**split_model (god_model):**

```diff
  // BEFORE: single 350-line Update handling 14 message types
  type AppModel struct {
      viewport viewport.Model
      items    []Item
      selected int
      spinner  spinner.Model
      status   string
  }

+ // AFTER: decomposed into focused child models
+ type AppModel struct {
+     list     ListModel
+     detail   DetailModel
+     status   StatusModel
+ }
+ 
+ func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
+     switch msg := msg.(type) {
+     case tea.WindowSizeMsg:
+         m.list.SetSize(msg.Width, msg.Height/2)
+         m.detail.SetSize(msg.Width, msg.Height/2)
+         return m, nil
+     case tea.KeyMsg:
+         return m, m.list.Update(msg)
+     default:
+         var cmd tea.Cmd
+         m.list, cmd = m.list.Update(msg)
+         if cmd != nil { return m, cmd }
+         m.status, cmd = m.status.Update(msg)
+         return m, cmd
+     }
+ }
```

**extract_command_boundary (no_blocking_update):**

```diff
  // BEFORE: blocking HTTP call inside Update
  func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
      case fetchMsg:
-         resp, err := http.Get(m.url)
-         if err != nil {
-             m.err = err
-             return m, nil
-         }
-         m.data = parse(resp)
+         return m, fetchCmd(m.url)
  }

+ func fetchCmd(url string) tea.Cmd {
+     return func() tea.Msg {
+         resp, err := http.Get(url)
+         if err != nil { return fetchErrMsg{err} }
+         return fetchRespMsg{parse(resp)}
+     }
+ }
```

**stabilize_viewport (viewport_recreation):**

```diff
  // BEFORE: viewport recreated every frame, scroll resets
  func (m Model) View() string {
-     vp := viewport.New(80, 24)
-     vp.SetContent(m.content)
-     return vp.View()
+     return m.viewport.View()
  }

  func (m Model) Init() (Model, tea.Cmd) {
+     m.viewport = viewport.New(80, 24)
      return m, nil
  }

  func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
      case tea.WindowSizeMsg:
+         m.viewport.Width = msg.Width
+         m.viewport.Height = msg.Height
      case contentMsg:
-         m.viewport.SetContent(msg.text)
+         if msg.text != m.lastContent {
+             m.viewport.SetContent(msg.text)
+             m.lastContent = msg.text
+         }
  }
```

### Phase 5: Output

**Dry run (default):** 在聊天中显示所有计划的动作和 patch。不修改文件。

```
Refactor plan for <project>:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  [1/3] split_model: Extract ListModel from AppModel
          Impact: 5 symbols affected, 0 dep cycles
          Files: +list_model.go, ~app.go
          Fixes: god_model

  [2/3] stabilize_viewport: Move viewport.New to Init
          Impact: 1 symbol, no side effects
          Files: ~app.go
          Fixes: viewport_recreation

  [3/3] extract_command_boundary: Wrap HTTP call in Cmd
          Impact: 2 symbols (Update, newCmd func)
          Files: ~app.go
          Fixes: no_blocking_update

  Warnings (not auto-fixed):
  - [WARN] missing_key_bindings_help: Add bubbles/help for discoverability
  - [INFO] string_allocation_pressure: Cache formatted strings in View()

Apply? [y/N] (or use --apply flag)
```

**With `--apply`:** 按顺序将 patch 写入磁盘。每次写入后验证代码仍然编译（`go build ./...`）。如果编译失败，回滚该文件并报告错误。

## Safety Checks

应用任何 patch 前：
1. 创建 `.refactor-backup/` 目录，备份原始文件
2. 验证 IR 数据显示修复后违规将被消除
3. 验证外部接口保持不变（结构体名称、导出方法）
4. 编写后运行 `go build ./...`
5. 可选：运行 `go test ./...` 验证测试未破坏

## Environment

- `ATREE_DB_PATH`: ATree database for impact analysis
- `CHARM_TUI_REFACTOR_BACKUP`: Create backups before applying (default: true)
