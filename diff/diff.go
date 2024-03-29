package diff

import (
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/mylxsw/go-utils/array"
	"github.com/pmezard/go-difflib/difflib"
)

// Diff 差异对比结果对象
type Diff struct {
	diff  string
	save  func() error
	clean func(keepOnly uint) error
}

// String 返回差异对比结果
func (d Diff) String() string {
	return d.diff
}

// Save 保存最后一次状态
func (d Diff) Save() error {
	return d.save()
}

// Clean 清理保存的状态文件，只保留 keep 个版本
// 实际上是保留 keep + 1 个版本，始终保留当前版本
func (d Diff) Clean(keep uint) error {
	return d.clean(keep)
}

// PrintAndSave 将差异对比信息输出并且保存最后一次状态
func (d Diff) PrintAndSave(out io.Writer) error {
	if d.diff == "" {
		return nil
	}

	_, _ = io.WriteString(out, d.String())
	return d.Save()
}

type FS interface {
	Exist(path string) bool
	WriteFile(path string, data []byte) error
	ReadFile(path string) ([]byte, error)
	ListFiles(dir string) ([]string, error)
	MkDir(path string) error
	Delete(path string) error
}

// Differ 文档对比工具
type Differ struct {
	fs          FS
	dataDir     string
	contextLine int
}

// NewDiffer create a new Differ
func NewDiffer(fs FS, dataDir string, contextLine int) *Differ {
	return &Differ{fs: fs, dataDir: dataDir, contextLine: contextLine}
}

// DiffLatest 将当前文档与最后一次保存的文档对比
func (ds Differ) DiffLatest(name string, targetStr string) Diff {
	var original []byte
	idx, _ := ds.fs.ReadFile(filepath.Join(ds.dataDir, name+".idx"))
	if string(idx) != "" {
		idxFilepath := filepath.Join(ds.dataDir, string(idx))
		if ds.fs.Exist(idxFilepath) {
			original, _ = ds.fs.ReadFile(idxFilepath)
		}
	}

	diffRes := ds.Diff(string(idx), string(original), name+".new", targetStr)

	fileMatchRegexp, _ := regexp.Compile(fmt.Sprintf(`^%s\.(\d+)\.stat$`, name))
	return Diff{
		diff: diffRes,
		save: func() error {
			targetName := fmt.Sprintf("%s.%s.stat", name, time.Now().Format("20060102150405"))
			_ = ds.fs.WriteFile(filepath.Join(ds.dataDir, targetName+".diff"), []byte(diffRes))
			_ = ds.fs.WriteFile(filepath.Join(ds.dataDir, targetName), []byte(targetStr))

			return ds.fs.WriteFile(filepath.Join(ds.dataDir, name+".idx"), []byte(targetName))
		},
		clean: func(keepOnly uint) error {
			files, err := ds.fs.ListFiles(ds.dataDir)
			if err != nil {
				return err
			}

			tss := array.Map(
				array.Filter(files, func(file string, _ int) bool { return fileMatchRegexp.MatchString(file) }),
				func(file string, _ int) string {
					return strings.Split(file, ".")[1]
				},
			)

			sort.Strings(tss)
			if len(tss) <= int(keepOnly)+1 {
				return nil
			}

			array.Each(
				array.Map(tss[:len(tss)-int(keepOnly)-1], func(ts string, _ int) string { return filepath.Join(ds.dataDir, fmt.Sprintf("%s.%s.stat", name, ts)) }),
				func(targetFile string, _ int) {
					_ = ds.fs.Delete(targetFile)
					_ = ds.fs.Delete(targetFile + ".diff")
				},
			)

			return nil
		},
	}
}

// Diff 文档差异对比
func (ds Differ) Diff(s1name, s1, s2name, s2 string) string {
	udiff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(s1),
		B:        difflib.SplitLines(s2),
		FromFile: s1name,
		ToFile:   s2name,
		Context:  ds.contextLine,
	}

	text, _ := difflib.GetUnifiedDiffString(udiff)
	return text
}
