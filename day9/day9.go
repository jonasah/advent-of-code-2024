package day9

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2024/lib/slicelib"
)

type File struct {
	id    int
	start int
	size  int
}

type Empty struct {
	start int
	size  int
}

func Part1(input string) int {
	diskMap := slicelib.Atoi(strings.Split(input, ""))

	fileID := 0
	var files []File
	var freeSpace []int
	for i, size := range diskMap {
		if i%2 == 0 {
			files = append(files, File{id: fileID, size: size})
			fileID++
		} else {
			freeSpace = append(freeSpace, size)
		}
	}

	checksum := 0
	position := 0
	for len(files) != 0 {
		file := files[0]
		files = files[1:]

		for range file.size {
			checksum += position * file.id
			position++
		}

		if len(freeSpace) != 0 {
			free := freeSpace[0]
			freeSpace = freeSpace[1:]

			for range free {
				if len(files) == 0 {
					break
				}

				checksum += position * files[len(files)-1].id
				position++

				files[len(files)-1].size--

				if files[len(files)-1].size == 0 {
					files = files[:len(files)-1]
				}
			}
		}
	}

	return checksum
}

func Part2(input string) int {
	diskMap := slicelib.Atoi(strings.Split(input, ""))

	position := 0
	fileID := 0
	files := map[int]*File{}
	var empty []*Empty
	for i, size := range diskMap {
		if i%2 == 0 {
			files[fileID] = &File{id: fileID, start: position, size: size}
			fileID++
		} else if size > 0 {
			empty = append(empty, &Empty{start: position, size: size})
		}

		position += size
	}

	for fileID > 0 {
		fileID--
		f := files[fileID]
		i := slices.IndexFunc(empty, func(e *Empty) bool { return e.start < f.start && e.size >= f.size })
		if i == -1 {
			continue
		}

		e := empty[i]
		f.start, e.start = e.start, f.start

		if e.size > f.size {
			avail := e.size - f.size
			e.size = f.size
			empty = append(empty, &Empty{start: f.start + f.size, size: avail})
			slices.SortFunc(empty, func(a, b *Empty) int { return a.start - b.start })
		}
	}

	checksum := 0
	for _, f := range files {
		for offset := range f.size {
			checksum += f.id * (f.start + offset)
		}
	}

	return checksum
}
