# Tree: directory statistics and snapshot comparison

## Goal

Extend the existing directory-tree CLI with size summaries and snapshot comparison.

Example: scan a generated build directory before and after a build. See which directories grew, which file types account for the growth, and which paths were added or removed. No file contents need to be read.

## Current code

The Go CLI traverses directories and prints a tree, optionally including files and sizes. The repository includes fixtures and output tests. The work below adds statistics, filtering, and comparison to that implementation.

## Work plan: ~ 56 hours

| Task | Hours | Done when |
|---|---:|---|
| 1. Update build and package setup | 4 | Go modules support a documented clean build and test run. Keep the existing command usable. |
| 2. Add depth and path filters | 8 | Maximum depth, hidden files, and include/exclude patterns work together predictably. The report states which paths were excluded. |
| 3. Define filesystem error behavior | 6 | Symlinks are listed but not followed by default. Unreadable or disappearing entries produce structured warnings and an incomplete-scan indicator. Traversal cannot recurse through symlink cycles. |
| 4. Add directory statistics | 8 | Reports show logical bytes, file/directory counts, extension totals, largest files, and largest subdirectories. Small fixtures have exact expected totals. |
| 5. Save versioned JSON snapshots | 8 | Snapshots contain relative paths, entry types, file sizes, and modification times. Output is sorted and includes schema version, filter settings, root information, and scan completeness. |
| 6. Compare snapshots | 8 | The CLI lists added, removed, and changed paths, ranks directory growth, and summarizes extension changes. Incompatible filters or incomplete scans are not treated as a clean comparison. |
| 7. Bound traversal and output memory | 6 | A depth-first mode can emit entries incrementally. Largest-file lists use a bounded top-N structure. Users can stop after a maximum entry count and see that the scan was partial. |
| 8. Finish integration tests and documentation | 8 | Tests cover errors, Unicode paths, filters, snapshots, and partial scans. Docs include one reproducible build-directory comparison and explain the size/count semantics. |
| **Total** | **56** | |

## Statistics semantics

- Report logical file size, not actual allocated disk blocks. Do not label the result as disk usage identical to `du`.
- Count discovered paths. Hard-linked paths may count the same underlying data more than once; document this instead of claiming deduplicated storage totals.
- Keep directory totals and visible output depth separate. Clearly state whether a depth limit stops scanning or only hides printed rows; use one documented policy in the first release.
- Treat filename changes as removal plus addition. Do not claim rename detection without content or file-identity evidence.
- A size or modification-time change is a metadata change, not proof that file content changed. Content hashing is outside this phase.
- Snapshot comparison loads the selected snapshots. Do not claim that this mode has the same memory bound as streaming a live traversal.

## Simple implementation

Keep filesystem traversal, aggregation, and text/JSON rendering separate. Use Go's standard library for walking, JSON, sorting, and bounded top-N selection. Avoid a database, daemon, GUI, or analytics service.

## Development data

Build temporary directory trees with generated filenames and known sizes. Do not scan home directories or publish snapshots of real personal or client projects during CC-CLI work. Snapshots contain path metadata, so their contents are not automatically free of personal data.
