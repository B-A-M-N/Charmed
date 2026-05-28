---
name: tui-audit
description: "Evidence-based architecture linting — checks TUI projects against canonical runtime constraints (bubbletea/bubbles/lipgloss) and empirical constraints derived from ecosystem patterns. Reports violations with confidence scores and shows evidence from similar repos. Trigger when: user says 'audit', 'check my TUI', 'find anti-patterns', 'lint', 'review architecture', or asks about code quality."
argument-hint: "[path-to-project] [--canonical-only] [--empirical-only] [--min-confidence 0.75] [--format text|json] [--with-evidence]"
allowed-tools: [Read, Write, Bash, Grep, Glob, ATree, mcp__charmed__scan_project, mcp__charmed__list_constraints, mcp__charmed__list_patterns, mcp__charmed__explain_concept]
---

# tui-audit — Evidence-Based Constraint Enforcement

Architecture linting using **two-tier knowledge system**:
- **Canonical constraints** (confidence 1.0): Runtime semantics of Bubble Tea core
- **Empirical constraints** (confidence <1.0): Patterns derived from ecosystem observations

Every finding includes: rule ID, severity, confidence, location, violation description, fix pattern, and optionally evidence from other repos.

## How It Works

### Canonical Mode (default)
Checks against hand-authored constraints based on runtime semantics:
- Bubble Tea lifecycle rules
- Bubble component contracts
- Lip Gloss rendering constraints

**Confidence: 1.0** - these are runtime facts, not observations.

### Empirical Mode (optional, enabled by default)
Checks against pattern-derived constraints:
- Storage must be async (14 repos, 11 fixes)
- Viewport needs diffing (9 repos, 4 fixes)
- Tickers need cleanup (17 repos, 5 fixes)

**Confidence: 0.75-0.95** - based on ecosystem evidence.

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Project path | Yes | User argument (default: current directory) |
| Mode | No | `--canonical-only` or `--empirical-only` (default: both) |
| Confidence threshold | No | `--min-confidence 0.75` (default: 0.75) |
| Output format | No | `--format text\|json` (default: text) |
| Evidence display | No | `--with-evidence` (show supporting repos/fixes) |

## Execution Pipeline

### Phase 1: Extract Evidence
Call `evidence_extract` (or equivalent MCP tool) to get architectural signals:

```
mcp__charmed__extract_evidence(
  project_path: <path>
)
→ Returns evidence graph with signals, topologies, primitives
```

### Phase 2: Load Constraints
Load **both** canonical and empirical constraints:

```
mcp__charmed__list_constraints(
  type: "all"  // or "canonical" / "empirical"
)
→ Returns constraints with confidence scores and evidence references
```

Example constraint:
```yaml
id: storage_must_be_async
type: empirical
confidence: 0.94
derived_from_pattern: storage_cmd_boundary
evidence:
  observations: 14
  violations: 12
  fixes: 11
```

### Phase 3: Constraint Evaluation

For each constraint, check project evidence against constraint requirements.

#### Canonical Constraint Example: no_blocking_update
```python
for signal in evidence.signals:
    if signal.primitive == "update_phase":
        for child_signal in signal.children:
            if child_signal.primitive == "storage_boundary":
                report(ERROR, confidence=1.0,
                       "Storage I/O in Update() blocks event loop",
                       fix="Wrap in tea.Cmd")
```

#### Empirical Constraint Example: viewport_churn  
```python
for signal in evidence.signals:
    if signal.primitive == "viewport":
        setcontent_calls = count_setcontent_per_frame(signal)
        if setcontent_calls > 1:
            report(ERROR, confidence=0.94,
                   "SetContent called multiple times per frame",
                   fix="Add content diffing (observed in 9 repos)",
                   evidence={
                       "repos_using_pattern": ["soft-serve", "glow", ...],
                       "violation_issues": ["issue#12", "issue#34"],
                       "fix_commits": ["abc123", "def456"]
                   })
```

#### viewport_recreation
```python
for binding in ir.component_bindings:
    if binding.type == "viewport.Model":
        for site in binding.construction_sites:
            if site.context in ("Update", "View"):
                report(CRITICAL, site.file, site.line,
                       "viewport.Model 在 Update/View 中被重建，导致滚动位置重置",
                       "在 Init() 中创建一次，仅调用 SetContent() 更新内容")
```

#### viewport_churn
```python
for boundary in ir.async_boundaries:
    if boundary.type == "SetContent" and boundary.call_frequency > 1:
        report(ERROR, boundary.file, boundary.line,
               f"SetContent 每帧调用 {boundary.call_frequency} 次，内容被重置而非追加",
               "增量构建内容，仅在内容实际变化时调用 SetContent")
```

#### no_blocking_update
```python
for model in ir.models:
    if has_blocking_io(model.update_body):
        report(ERROR, model.file, model.line,
               "Update() 中检测到阻塞 I/O（http.Get、文件读取、time.Sleep）",
               "包装为 tea.Cmd: return func() tea.Msg { ... }")
```

#### unbounded_cmd_spawning
```python
for boundary in ir.async_boundaries:
    if boundary.type == "CmdSpawn" and boundary.in_loop and boundary.rate == "every_iteration":
        report(ERROR, boundary.file, boundary.line,
               "每次迭代产生命令，创建无界 goroutine",
               "使用 time.Tick 限流，或合并为单个批处理命令")
```

#### sync_io_in_init
```python
for model in ir.models:
    if has_blocking_io(model.init_body):
        report(ERROR, model.file, model.line,
               "Init() 包含阻塞操作，延迟 TUI 启动",
               "将 I/O 移至 Init() 返回的 tea.Cmd 中")
```

#### viewport_without_resize_handling
```python
for binding in ir.component_bindings:
    if binding.type == "viewport.Model":
        if not has_handler(model, "WindowSizeMsg"):
            report(ERROR, model.file, model.line,
                   "viewport 未处理 WindowSizeMsg，尺寸将为 0",
                   "在 Update() 中处理 tea.WindowSizeMsg，传播尺寸到 viewport")
```

#### style_recomputation_hot_path
```python
for model in ir.models:
    if model.view_has("lipgloss.NewStyle") or model.view_has_render_call():
        report(WARN, model.file, model.line,
               "View() 中重新计算样式，每帧浪费 CPU",
               "缓存来自模型状态的样式，仅在输入数据变化时重新计算")
```

#### string_allocation_pressure
```python
for model in ir.models:
    sprintf_count = count_sprintf(model.view_body)
    if sprintf_count > 5:
        report(INFO, model.file, model.line,
               f"View() 中有 {sprintf_count} 次 fmt.Sprintf，造成 GC 压力",
               "使用 lipgloss JoinHorizontal/JoinVertical，预计算静态字符串")
```

#### missing_key_bindings_help
```python
if not has_import(ir.files, "charmbracelet/bubbles/help"):
    report(WARN, "-", "-",
           "未检测到 key.Binding 导入，用户没有可见的帮助",
           "添加 bubbles/help，使用 key bindings 显示可用按键")
```

#### tea_batch_misuse
```python
for boundary in ir.async_boundaries:
    if boundary.type == "Batch" and boundary.has_state_dependency:
        report(INFO, boundary.file, boundary.line,
               "tea.Batch 并发运行命令，若顺序重要应改用 tea.Sequence",
               "状态相关操作使用 tea.Sequence，独立操作使用 tea.Batch")
```

#### missing_alt_screen_cleanup
```python
if has_option(ir.program, "EnterAltScreen") and not has_cleanup(ir.program, "ExitAltScreen"):
    report(WARN, "-", "-",
           "使用了 Alt Screen 但未保证清理，崩溃时终端可能损坏",
           "在 defer 或清理处理程序中调用 ExitAltScreen")
```

### Phase 4: 报告

文本格式输出：

```
╔══════════════════════════════════════════════════════════════╗
║  tui-audit — <project-name>                                 ║
╠══════════════════════════════════════════════════════════════╣
║  扫描: <N> 个文件, <M> 个模型, <K> 个消息类型               ║
║  检查: <total> 条规则 · <passed> 通过 · <failed> 失败       ║
╠══════════════════════════════════════════════════════════════╣
║                                                            ║
║  🔴 CRITICAL (N)                                           ║
║  ────────────────────────────────────────────────────────  ║
║  [viewport_recreation] cmd/app.go:42                       ║
║    viewport.Model 在 Update() 中被重建                     ║
║    修复: 在 Init() 中创建，仅使用 SetContent()             ║
║                                                            ║
║  🛑 ERROR (N)                                              ║
║  ────────────────────────────────────────────────────────  ║
║  [god_model] cmd/app.go:87                                 ║
║    Update() 为 340 行，拥有太多状态和消息类型               ║
║    修复: 分解为子模型，父模型协调消息传递                   ║
║                                                            ║
║  ⚠️  WARNING (N)                                           ║
║  ────────────────────────────────────────────────────────  ║
║  [missing_key_bindings_help]                               ║
║    未检测到 key.Binding 导入                               ║
║    修复: 添加 bubbles/help，显示可用按键                   ║
║                                                            ║
║  ℹ️  INFO (N) [strict 模式]                                ║
║  ────────────────────────────────────────────────────────  ║
║  [string_allocation_pressure] cmd/app.go:120               ║
║    View() 中有大量字符串格式化                              ║
║    修复: 使用 strings.Builder，缓存静态字符串               ║
║                                                            ║
╚══════════════════════════════════════════════════════════════╝
```

使用 `--format json` 时，输出结构化 JSON：
```json
{
  "project": "my-tui",
  "scanned_at": "2026-05-26T...",
  "mode": "full",
  "ir_version": "1",
  "summary": {
    "total_rules": 13,
    "passed": 8,
    "failed": 5,
    "by_severity": {"critical": 1, "error": 2, "warning": 1, "info": 1}
  },
  "ir": {
    "models": 3,
    "messages": 12,
    "components_bound": 4,
    "async_boundaries": 5
  },
  "findings": [
    {
      "rule_id": "viewport_recreation",
      "severity": "CRITICAL",
      "file": "cmd/app.go",
      "line": 42,
      "message": "viewport.Model recreated inside Update()",
      "fix": "Create once in Init(), use SetContent() only"
    }
  ]
}
```

## Quick Mode Fallback

当 ATree 不可用时，使用 `--quick` 参数或自动降级。在此模式下，通过 grep 而非 IR 评估规则。

快速模式覆盖范围：
- ✅ god_model（通过 grep `func.*Update` 并计算行数）
- ✅ giga_switch_update（计算 `case` 语句数量）
- ✅ viewport_recreation（grep `viewport.New(` 上下文）
- ✅ no_blocking_update（grep `http\.\|ioutil\.\|time\.Sleep` 在 Update 中）
- ✅ sync_io_in_init（grep I/O 模式在 Init 中）
- ✅ viewport_without_resize_handling（grep viewport 使用，无 WindowSizeMsg）
- ✅ style_recomputation_hot_path（grep `lipgloss.NewStyle\|Render` 在 View 中）
- ✅ string_allocation_pressure（计算 `fmt.Sprintf` 调用）
- ❌ unbounded_cmd_spawning（需要 ATree 进行循环检测）
- ❌ viewport_churn（需要 ATree 进行调用频率分析）
- ❌ tea_batch_misuse（需要 ATree 进行依赖性分析）
- ❌ missing_key_bindings_help（可通过 import grep 部分检测）
- ❌ missing_alt_screen_cleanup（可通过 grep 部分检测）

快速模式报告应明确标注：
```
⚠️  Quick mode: 8/13 rules evaluated. Run without --quick for full analysis.
```

## Constraint Rule Reference

| Rule | Severity | Category | Detectable in Quick Mode |
|------|----------|----------|--------------------------|
| `god_model` | CRITICAL | ARCHITECTURE | ✅ |
| `giga_switch_update` | CRITICAL | ARCHITECTURE | ✅ |
| `viewport_recreation` | CRITICAL | PERFORMANCE | ✅ |
| `no_blocking_update` | ERROR | CORRECTNESS | ✅ |
| `viewport_churn` | ERROR | PERFORMANCE | ❌ Needs ATree |
| `unbounded_cmd_spawning` | ERROR | CONCURRENCY | ❌ Needs ATree |
| `sync_io_in_init` | ERROR | CORRECTNESS | ✅ |
| `viewport_without_resize_handling` | ERROR | ROBUSTNESS | ✅ |
| `style_recomputation_hot_path` | WARN | PERFORMANCE | ✅ |
| `string_allocation_pressure` | INFO | PERFORMANCE | ✅ |
| `missing_key_bindings_help` | WARN | UX | ⚠️ Partial |
| `missing_alt_screen_cleanup` | WARN | ROBUSTNESS | ⚠️ Partial |
| `tea_batch_misuse` | INFO | CORRECTNESS | ❌ Needs ATree |

## Environment

- `ATREE_DB_PATH`: Override auto-detected ATree database path
- `CHARM_TUI_STRICT`: Enable strict mode by default
- `CHARM_TUI_FORMAT`: Default output format (text|json)
