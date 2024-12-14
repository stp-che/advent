package solution

import (
	"log"
	"os"
)

type Solution struct {
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) Run(inputPath string) int {
	diskMap := s.readData(inputPath)
	sum := 0

	latestFileID := len(diskMap) / 2
	latestFilePos := len(diskMap) - 1
	mapPos := 1
	blockPos := int(diskMap[0])
	fileID := 0
	for mapPos <= latestFilePos {
		spaces := int(diskMap[mapPos])
		for i := 0; i < spaces; i++ {
			// fmt.Printf("%d -> %d\n", blockPos, latestFileID)
			sum += latestFileID * blockPos
			blockPos++
			diskMap[latestFilePos]--
			if diskMap[latestFilePos] == 0 {
				latestFileID--
				latestFilePos -= 2
			}
		}
		if mapPos > latestFilePos {
			break
		}
		mapPos++
		fileID++
		fileBlocks := int(diskMap[mapPos])
		for i := 0; i < fileBlocks; i++ {
			// fmt.Printf("%d -> %d\n", blockPos, fileID)
			sum += fileID * blockPos
			blockPos++
		}
		mapPos++
	}

	return sum
}

func (s *Solution) Run1(inputPath string) int {
	diskMap := s.readData(inputPath)

	spacesOffsets := make([]int, len(diskMap)/2)
	filesOffsets := make([]int, len(diskMap)/2+1)
	offset := 0
	for i := 0; i < len(diskMap)-1; i += 2 {
		filesOffsets[i/2] = offset
		offset += int(diskMap[i])
		spacesOffsets[i/2] = offset
		offset += int(diskMap[i+1])
	}
	filesOffsets[len(filesOffsets)-1] = offset

	fileId := len(filesOffsets) - 1
	for fileId > 0 {
		i := 0
		fileLen := diskMap[fileId*2]
		for i < fileId {
			spaceLen := diskMap[i*2+1]
			if spaceLen >= fileLen {
				break
			}
			i++
		}
		if i < fileId {
			filesOffsets[fileId] = spacesOffsets[i]
			spacesOffsets[i] += int(fileLen)
			diskMap[i*2+1] -= fileLen
		}
		fileId--
	}

	// fmt.Println(filesOffsets)

	sum := 0
	for fileId, offset := range filesOffsets {
		fileSize := int(diskMap[fileId*2])
		sum += fileId * ((offset+fileSize-1)*(offset+fileSize) - (offset-1)*offset) / 2
	}

	return sum
}

func (s *Solution) readData(inputPath string) []byte {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	zero := byte('0')
	for i := 0; i < len(data); i++ {
		data[i] -= zero
	}

	return data
}
