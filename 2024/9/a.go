package day9

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day9A",
		Aliases: []string{"day9a"},
		Usage:   "Day 9 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day9B",
		Aliases: []string{"day9b"},
		Usage:   "Day 9 part B",
		Action:  partB,
	}
)

func partA(ctx *cli.Context) error {
	answer, err := processA(aoc.PreparePuzzle(puzzle))
	if err != nil {
		return err
	}
	log.Printf("Answer A: %v", answer)
	return nil
}

func partB(ctx *cli.Context) error {
	answer, err := processB(aoc.PreparePuzzle(puzzle))
	if err != nil {
		return err
	}
	log.Printf("Answer B: %v", answer)
	return nil
}

func processA(puzzle []string) (int, error) {
	d := newDisk(puzzle[0])
	d.dump()
	for fileId := d.maxId; fileId >= 0; fileId-- {
		d.moveFile(fileId, d.findFreeBlock)
	}
	d.dump()
	return d.checksum(), nil
}

func processB(puzzle []string) (int, error) {
	d := newDisk(puzzle[0])
	d.dump()
	for fileId := d.maxId; fileId >= 0; fileId-- {
		d.moveFile(fileId, d.findBigEnoughFreeBlock)
	}
	d.dump()
	return d.checksum(), nil
}

type disk struct {
	filemap map[int][]fileInfo
	maxId   int
}

func newDisk(puzzle string) *disk {
	filemap := make(map[int][]fileInfo)
	var freelist []fileInfo
	id := 0
	start := 0
	file := true
	for _, p := range puzzle {
		size := aoc.MustAtoi(string(p))
		if size == 0 {
			file = !file
			continue
		}
		if file {
			fileInfos, ok := filemap[id]
			if !ok {
				fileInfos = nil
			}
			filemap[id] = append(fileInfos, fileInfo{id: id, start: start, size: size})
			id++
		} else {
			freelist = append(freelist, fileInfo{id: -1, start: start, size: size})
		}
		start += size
		file = !file
	}
	filemap[-1] = freelist
	return &disk{filemap: filemap, maxId: id - 1}
}

func (d *disk) dump() {
	return
	log.Println(d.filemap)
	// Compile an array of fileInfo sorted by start.
	var orderedFiles []fileInfo
	for _, fis := range d.filemap {
		for _, fi := range fis {
			if fi.size == 0 {
				continue
			}
			orderedFiles = append(orderedFiles, fi)
		}
	}
	sort.Slice(orderedFiles, func(i, j int) bool {
		if orderedFiles[i].start == orderedFiles[j].start {
			log.Printf("Overlapping starts for %v & %v", orderedFiles[i], orderedFiles[j])
		}
		return orderedFiles[i].start < orderedFiles[j].start
	})
	//log.Printf("ordered files: %v", orderedFiles)
	s := ""
	for _, of := range orderedFiles {
		fid := fmt.Sprintf("%v", of.id)
		if of.id == -1 {
			fid = "."
		}
		s += strings.Repeat(fid, of.size)
	}
	s2 := ""
	for n, c := range s {
		if n%10 == 0 {
			s2 += " "
		}
		s2 += string(c)
	}
	log.Printf("%v", s2)
}

func (d *disk) checksum() int {
	sum := 0
	for id, fis := range d.filemap {
		if id == -1 {
			continue
		}
		for _, fi := range fis {
			for n := 0; n < fi.size; n++ {
				sum += id * (fi.start + n)
			}
		}
	}
	return sum
}

func (d *disk) moveFile(fileId int, findFreeBlock func(int) ([]fileInfo, int)) {
	log.Printf("-------------------- moving %v", fileId)
	fis, ok := d.filemap[fileId]
	if !ok {
		panic(fmt.Sprintf("Failed to find file #%v", fileId))
	}
	if len(fis) != 1 {
		panic(fmt.Sprintf("File %v is already fragmented into %v parts", len(fis)))
	}
	completed := false
	for !completed {
		//fis, _ = d.filemap[fileId]
		//log.Printf("fis=%v, len=%v", fis, len(fis))
		fileFragment := fis[len(fis)-1]
		freelist, fbi := findFreeBlock(fileFragment.size)
		if freelist == nil {
			return
		}
		if freelist[fbi].start > fileFragment.start {
			//log.Printf("not moving %v, free space %v", fileFragment, d.filemap[-1])
			return
		}
		//log.Printf("free block %v found for %v", freelist[fbi], fileFragment)
		originalFileStart := fileFragment.start
		originalFileSize := fileFragment.size
		switch {
		case fileFragment.size > freelist[fbi].size:
			// Multiple free blocks required.
			//log.Printf("partial move %v to %v", fileFragment, freelist[fbi])
			fis[len(fis)-1] = fileInfo{id: fileFragment.id, start: freelist[fbi].start, size: freelist[fbi].size}
			remainingFile := fileInfo{id: fileFragment.id, start: originalFileStart + freelist[fbi].size, size: originalFileSize - freelist[fbi].size}
			fis = append(fis, remainingFile)
			//log.Printf("appended %v to give %v", remainingFile, fis)
			freeSize := freelist[fbi].size
			freelist = append(freelist, fileInfo{id: -1, start: originalFileStart, size: freeSize})
			freelist[fbi].size = 0
		case fileFragment.size == freelist[fbi].size:
			// File fits exactly into the free block.
			//log.Printf("exact move %v to %v", fileFragment, freelist[fbi])
			fis[len(fis)-1].start = freelist[fbi].start
			freelist = append(freelist, fileInfo{id: -1, start: originalFileStart, size: originalFileSize})
			freelist[fbi].size = 0
			completed = true
		case fileFragment.size < freelist[fbi].size:
			// File fits into the free block with space to spare.
			//log.Printf("move with space %v to %v", fileFragment, freelist[fbi])
			fis[len(fis)-1].start = freelist[fbi].start
			freelist = append(freelist, fileInfo{id: -1, start: originalFileStart, size: originalFileSize})
			freelist[fbi].start += fileFragment.size
			freelist[fbi].size -= fileFragment.size
			completed = true
		default:
			panic("Unhandled case!")
		}
		sort.Slice(freelist, func(i, j int) bool {
			return freelist[i].start < freelist[j].start
		})
		d.filemap[-1] = freelist
		log.Printf("fis=%v", fis)
		d.filemap[fileId] = fis
		d.dump()
	}
}

func (d *disk) findFreeBlock(_ int) ([]fileInfo, int) {
	freelist, ok := d.filemap[-1]
	if !ok {
		panic("Failed to find freelist")
	}
	for i, fi := range freelist {
		if fi.size > 0 {
			return freelist, i
		}
	}
	panic("No free blocks!")
}

func (d *disk) findBigEnoughFreeBlock(size int) ([]fileInfo, int) {
	freelist, ok := d.filemap[-1]
	if !ok {
		panic("Failed to find freelist")
	}
	for i, fi := range freelist {
		if fi.size >= size {
			return freelist, i
		}
	}
	return nil, -1
}

type fileInfo struct {
	id, start, size int
}

func (fi fileInfo) String() string {
	if fi.id == -1 {
		return fmt.Sprintf("free @%v len %v", fi.start, fi.size)
	}
	return fmt.Sprintf("#%v @%v len %v", fi.id, fi.start, fi.size)
}
