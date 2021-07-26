package trigram

import (
	"io/ioutil"
	"os"
	"sort"
	"sync"
)

type Trigram struct {
	syncLock   sync.Mutex
	trigramMap map[uint32][]string
}

func NewTrigram() *Trigram {
	t := Trigram{
		trigramMap: make(map[uint32][]string),
	}
	return &t
}

func (t *Trigram) getTrigram(fileContext string) []uint32 {
	trig := make([]uint32, 0)
	for i := 0; i < len(fileContext)-2; i++ {
		context := uint32(uint32(fileContext[i])<<16 | uint32(fileContext[i+1])<<8 | uint32(fileContext[i+2]))
		trig = append(trig, context)
	}

	return trig
}

func (t *Trigram) addTextToTrigramMap(context uint32, fileName string) {
	t.syncLock.Lock()
	defer t.syncLock.Unlock()
	if _, ok := t.trigramMap[uint32(context)]; !ok {
		t.trigramMap[uint32(context)] = make([]string, 0)
	}

	if pos := t.findTheString(t.trigramMap[context], fileName); pos == len(t.trigramMap[context]) {
		t.trigramMap[context] = append(t.trigramMap[context], fileName)
	}

	sort.Strings(t.trigramMap[context])
}

func (t *Trigram) createContextIndex(fileContext string, fileName string) {
	trigm := t.getTrigram(fileContext)

	for _, tr := range trigm {
		t.addTextToTrigramMap(tr, fileName)
	}
}

func (t *Trigram) Add(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	fileContext, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	t.createContextIndex(string(fileContext), fileName)

	return nil
}

func (t *Trigram) Delete(fileName string) error {
	for tri, files := range t.trigramMap {
		if pos := t.findTheString(files, fileName); pos != len(files) {
			if len(files) == 1 {
				delete(t.trigramMap, tri)
			} else {
				copy(files[pos:], files[pos+1:])
			}
		}
	}
	return nil
}

func (t *Trigram) Find(queryString string) []string {
	trg := t.getTrigram(queryString)
	resultFileMap := make(map[string]bool)

	for _, tr := range trg {
		files := t.trigramMap[tr]

		for _, file := range files {
			resultFileMap[file] = true
		}
	}

	return t.mapToSlice(resultFileMap)
}

func (t *Trigram) mapToSlice(resultMap map[string]bool) []string {
	result := make([]string, 0)
	for k := range resultMap {
		result = append(result, k)
	}

	return result
}

func (t *Trigram) findTheString(files []string, fileName string) int {
	start, end := 0, len(files)

	for start < end {
		mid := (start + end) / 2
		if files[mid] < fileName {
			start = mid + 1
		} else if files[mid] > fileName {
			end = mid
		} else {
			return mid
		}
	}
	return start
}
