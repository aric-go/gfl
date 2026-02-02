# gfl copy

Copy the current branch to a new branch, providing a simpler alternative to `gfl start --base=@`.

## Usage

```bash
gfl copy [new-branch-name]
gfl copy [-y|--confirm]
```

### Aliases

- `gfl cp` - Short alias for copy command

### Arguments

- `new-branch-name` (optional) - The name for the new branch
  - If not provided, automatically generates `{current-branch-name}-copyed`
  - When omitted, the `-y` or `--confirm` flag is required

### Options

This command inherits all global flags:
- `--confirm, -y` - Confirm operation without prompting (required when branch name is omitted)
- `--debug, -d` - Enable debug mode

## Description

The `gfl copy` command creates a new branch from the current branch. It's a convenience command that simplifies the workflow of copying a branch, which would otherwise require using `gfl start --base=@ [new-name]`.

### Key Features

1. **Auto-Generate Name**: When branch name is omitted with `-y` flag, automatically generates `{current-branch-name}-copyed`
2. **Branch Naming**: Automatically applies the configured `feature/` prefix and nickname (if configured)
3. **Case Formatting**: Respects the `branchCaseFormat` configuration for consistent branch naming
4. **Safety Checks**: Validates the working directory, remote branch existence, and duplicate names before creating
5. **Remote Sync**: Fetches latest remote information before creating the branch

## Examples

### Auto-Generate Branch Name (Quick Copy)

```bash
# Copy current branch with auto-generated name
gfl copy -y

# Current branch: feature/kat-10522-staging
# Creates: feature/kat-10522-staging-copyed

# Current branch: feature/user-auth
# Creates: feature/user-auth-copyed
```

### Basic Copy

```bash
# Copy current branch to a new branch with custom name
gfl copy kat-10520-copyed

# If nickname is configured as "aric":
# Creates: feature/aric/kat-10520-copyed

# If nickname is not configured:
# Creates: feature/kat-10520-copyed
```

### Using Alias

```bash
# Using the short alias with auto-generation
gfl cp -y

# Using the short alias with custom name
gfl cp test-branch-copy-2
```

### With Case Formatting

If `branchCaseFormat` is set to `lower`:

```bash
gfl copy KAT-10520-FIX
# Creates: feature/kat-10520-fix  (note: lowercase conversion)
```

## Comparison with gfl start

| Aspect | gfl start | gfl copy |
|--------|-----------|----------|
| Base branch | Configurable (dev/main via config, or --base flag) | Always current branch |
| Branch naming | feature/nickname/name (same pattern) | feature/nickname/name (same pattern) |
| Use case | Start new feature from base branch | Copy/duplicate current branch |
| Remote requirement | Base branch must exist in remote | Current branch must exist in remote |
| Prefix configuration | Uses configured featurePrefix | Always uses "feature" prefix |

### When to use gfl copy vs gfl start

**Use `gfl copy` when:**
- You want to create a variation of your current branch
- You're working on a feature and want to try a different approach
- You need to create a similar feature based on existing work
- You want to quickly duplicate the current branch state

**Use `gfl start` when:**
- Starting a completely new feature from the development branch
- You need to specify a custom base branch
- Starting work from main/develop branch

## Error Scenarios

### Working Directory Not Clean

```bash
# If you have uncommitted changes
$ gfl copy new-branch
ERROR: Working directory is not clean, please commit or stash changes first

# Solution: Commit or stash your changes first
git stash
gfl copy new-branch
```

### Current Branch Not in Remote

```bash
# If current branch hasn't been pushed
$ gfl copy local-copy
ERROR: Current branch 'local-only-branch' does not exist in remote repository, please push it first

# Solution: Push the current branch first
git push -u origin HEAD
gfl copy local-copy
```

### Branch Already Exists

```bash
# If target branch name already exists
$ gfl copy existing-branch
ERROR: Branch 'feature/existing-branch' already exists, please use a different name

# Solution: Use a different branch name
gfl copy existing-branch-v2
```

## Workflow Example

### Quick Copy with Auto-Generated Name

```bash
# Scenario: Quickly duplicate current branch without thinking of a name

# 1. Current branch
$ git branch
* feature/user-auth-v1

# 2. Quick copy with auto-generated name
$ gfl copy -y
 Syncing remote branches...
 Copying branch...
✅ Successfully copied branch 'feature/user-auth-v1' to 'feature/user-auth-v1-copyed'

# 3. Now on the new branch
$ git branch
  feature/user-auth-v1
* feature/user-auth-v1-copyed
```

### Custom Name Copy

```bash
# Scenario: Working on a feature and need to try a different approach

# 1. Current branch with some work
$ git branch
* feature/user-auth-v1

# 2. Want to try a different approach without losing current work
$ gfl copy user-auth-v2
 Syncing remote branches...
 Copying branch...
✅ Successfully copied branch 'feature/user-auth-v1' to 'feature/user-auth-v2'

# 3. Now on the new branch with all previous work
$ git branch
  feature/user-auth-v1
* feature/user-auth-v2
```

## Implementation Details

### Auto-Generated Branch Names

When no branch name is provided with the `-y` flag:

1. Extracts the current branch name
2. Removes prefix and nickname if present (e.g., from `feature/aric/user-auth` extracts `user-auth`)
3. Appends `-copyed` suffix (e.g., `user-auth` → `user-auth-copyed`)
4. Applies the full branch naming rules (prefix + nickname if configured)

**Examples**:
- `feature/kat-10522-staging` → `feature/kat-10522-staging-copyed`
- `feature/aric/user-auth` → `feature/aric/user-auth-copyed`
- `feature/bug-fix` → `feature/bug-fix-copyed`

### Branch Naming

The `gfl copy` command always uses the `feature/` prefix (respects `featurePrefix` configuration):

1. **With nickname configured**: `feature/{nickname}/{new-branch-name}`
2. **Without nickname**: `feature/{new-branch-name}`

The branch name also respects the `branchCaseFormat` configuration for automatic case conversion.

### Pre-copy Validation

Before creating the new branch, the command performs these checks:

1. ✓ Working directory is clean (no uncommitted changes)
2. ✓ Current branch exists in remote repository
3. ✓ Generated branch name doesn't already exist locally

### Creation Process

1. Fetches latest remote information (`git fetch origin`)
2. Validates current branch exists in remote
3. Validates generated branch name is unique
4. Creates new branch from remote version of current branch
5. Checks out the new branch

## See Also

- [gfl start](./start.md) - Start a new feature from base branch
- [gfl rename](./rename.md) - Rename an existing branch
- [gfl checkout](./checkout.md) - Interactive branch switching
