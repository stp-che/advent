package solution

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type dirSizeHandler interface {
	HandleSize(int)
}

type spaceManager interface {
	SetSpaceUsed(int)
}

type Solution struct {
	dirSizeHandler dirSizeHandler
	spaceManager   spaceManager
}

func New(h dirSizeHandler, s spaceManager) *Solution {
	return &Solution{h, s}
}

func (s *Solution) Run(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	s.handleInput(file)
}

func (s *Solution) handleInput(file *os.File) {
	t := newTerminal()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t.handleRow(scanner.Text())
	}

	if s.spaceManager != nil {
		s.spaceManager.SetSpaceUsed(t.rootDir.size())
	}

	t.rootDir.traverseDirs(func(d *dir) {
		s.dirSizeHandler.HandleSize(d.size())
	})
}

type terminal struct {
	rootDir *dir
	curDir  *dir
}

func newTerminal() *terminal {
	rootDir := newDir(nil)
	return &terminal{rootDir: rootDir, curDir: rootDir}
}

func (t *terminal) handleRow(row string) {
	// fmt.Println(row)
	parts := strings.Split(row, " ")
	switch parts[0] {
	case "$":
		t.exec(parts[1], parts[2:]...)
	case "dir":
		t.addDir(parts[1])
	default:
		size, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		t.addFile(parts[1], int(size))
	}
}

func (t *terminal) exec(cmd string, args ...string) {
	if cmd == "cd" {
		t.cd(args[0])
	}
}

func (t *terminal) cd(path string) {
	switch path {
	case "/":
		t.curDir = t.rootDir
	case "..":
		if t.curDir.parent != nil {
			t.curDir = t.curDir.parent
		}
	default:
		t.curDir = t.curDir.getDir(path)
	}
}

func (t *terminal) addDir(name string) {
	t.curDir.addDir(name)
}

func (t *terminal) addFile(name string, size int) {
	t.curDir.addFile(name, size)
}

type dir struct {
	parent         *dir
	dirs           map[string]*dir
	files          map[string]int
	sizeCache      int
	sizeCalculated bool
}

func newDir(parent *dir) *dir {
	return &dir{
		parent: parent,
		dirs:   make(map[string]*dir),
		files:  make(map[string]int),
	}
}

func (d *dir) getDir(name string) *dir {
	return d.dirs[name]
}

func (d *dir) addDir(name string) {
	d.dirs[name] = newDir(d)
}

func (d *dir) addFile(name string, size int) {
	d.files[name] = size
}

func (d *dir) size() int {
	if !d.sizeCalculated {
		d.sizeCache = 0
		for _, child := range d.dirs {
			d.sizeCache += child.size()
		}
		for _, fSize := range d.files {
			d.sizeCache += fSize
		}
		d.sizeCalculated = true
	}

	return d.sizeCache
}

func (d *dir) traverseDirs(f func(*dir)) {
	for _, child := range d.dirs {
		child.traverseDirs(f)
	}
	f(d)
}
