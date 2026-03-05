# GoNest Tools

Development and release management tools for GoNest projects.

## 🛠️ Tools

### 📊 badge - Coverage Badge Generator

Generates beautiful coverage badges from Go test coverage profiles.

**Features:**
- Automatic color coding based on coverage percentage
- Integration with shields.io
- Configurable labels and output paths

**Usage:**
```bash
gonest-badge -in coverage.out -out .public/coverage.svg -label coverage
```

**Flags:**
- `-in` - Input coverage profile (default: `coverage.out`)
- `-out` - Output SVG file (default: `.public/coverage.svg`)
- `-label` - Badge label (default: `coverage`)

**Colors:**
| Coverage | Color        |
| -------- | ------------ |
| 90%+     | Bright Green |
| 80-89%   | Green        |
| 70-79%   | Yellow Green |
| 60-69%   | Yellow       |
| 50-59%   | Orange       |
| <50%     | Red          |

---

### 🧹 clean - File Cleanup Tool

Removes files matching glob patterns.

**Usage:**
```bash
gonest-clean "*.out" "*.coverage.out" ".public/*.svg"
```

**Examples:**
```bash
# Clean all coverage files
gonest-clean coverage.out "*.coverage.out"

# Clean badges
gonest-clean ".public/*.svg"

# Clean test artifacts
gonest-clean "*.test" "*.out"

# Multiple patterns at once
gonest-clean "*.out" "*.test" "tmp/*"
```

---

### 🏷️ tag - Git Tag Management

Manages Git tags for multi-module repositories. Automatically discovers modules and creates/pushes/deletes tags for each.

**Features:**
- Auto-discovery of modules (directories with `go.mod`)
- Parallel tag operations for speed
- Version bumping (minor/patch)
- Tag purging (cleanup old versions)
- Safe deletion (local + remote)

**Usage:**

**Create and push tags:**
```bash
gonest-tag --create --push v0.1.0
```

**Bump patch version (0.1.0 → 0.1.1):**
```bash
gonest-tag --bump patch
```

**Bump minor version (0.1.0 → 0.2.0):**
```bash
gonest-tag --bump minor
```

**Delete tags:**
```bash
gonest-tag --delete v0.1.0
```

**Purge all tags except v0.1.0:**
```bash
gonest-tag --purge v0.1.0
```

**Flags:**
- `--create` - Create tags locally
- `--push` - Push tags to remote
- `--delete` - Delete tags locally and remotely
- `--bump <minor|patch>` - Auto-bump version and replace old tags
- `--purge <version>` - Delete all tags except specified version

**How it works:**
1. Discovers all modules by finding `go.mod` files
2. Excludes: root, `examples/*`, `_tools/*`
3. Creates tags: `v0.1.0` (root) and `core/v0.1.0`, `validator/v0.1.0`, etc
4. Executes git commands in parallel

**Example output:**
```
Bumping version from v0.1.0 to v0.1.1 (level: patch)
Creating tag v0.1.1 for root...
Creating tag core/v0.1.1...
Creating tag validator/v0.1.1...
Creating tag controller/v0.1.1...
...
Pushing tag v0.1.1...
Pushing tag core/v0.1.1...
...
Deleting tag v0.1.0...
Deleting tag core/v0.1.0...
...
```

---

## 📦 Installation

### Install all tools at once:
```bash
make install
```

Tools will be installed to `$GOPATH/bin` (usually `~/go/bin`).

### Install individually:
```bash
cd badge && go install
cd clean && go install
cd tag && go install
```

### Build binaries (without installing):
```bash
make build
```

Binaries will be in `./bin/`:
- `./bin/gonest-badge`
- `./bin/gonest-clean`
- `./bin/gonest-tag`

---

## 🚀 Usage in Projects

### In Makefile:
```makefile
# Badge generation
badge:
	gonest-badge -in coverage.out -out .public/coverage.svg

# Cleanup
clean:
	gonest-clean coverage.out "*.coverage.out" ".public/*.svg"

# Release
tag:
	gonest-tag --create --push v0.1.0

tag-bump:
	gonest-tag --bump patch
```

### In CI/CD:
```yaml
- name: Generate coverage badge
  run: gonest-badge -in coverage.out -out .public/coverage.svg

- name: Clean artifacts
  run: gonest-clean "*.out" "*.test"

- name: Create release tags
  run: gonest-tag --create --push ${{ github.ref_name }}
```

---

## 🔧 Development

### Structure
```
gonest-tools/
├── badge/
│   ├── go.mod
│   └── main.go
├── clean/
│   ├── go.mod
│   └── main.go
├── tag/
│   ├── go.mod
│   ├── main.go
│   └── version_test.go
├── go.work          # Workspace for developing together
├── Makefile
└── README.md
```

### Commands
```bash
# Build all tools
make build

# Test
make test

# Install
make install

# Clean
make clean

# Tidy modules
make mod-tidy
```

### Module Structure
Each tool is an independent Go module:
```go
// badge/go.mod
module github.com/gonest-dev/gonest-tools/badge

// clean/go.mod
module github.com/gonest-dev/gonest-tools/clean

// tag/go.mod
module github.com/gonest-dev/gonest-tools/tag
```

This isolation prevents tool dependencies from polluting consumer projects.

---

## 📝 License

MIT

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

---

## 🔗 Links

- [GoNest Framework](https://github.com/gonest-dev/gonest)
- [Documentation](https://gonest.dev)
- [Examples](https://github.com/gonest-dev/gonest-examples)