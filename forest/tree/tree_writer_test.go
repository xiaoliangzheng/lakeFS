package tree

import (
	"bufio"
	"log"
	"os"
	"testing"

	gr "github.com/treeverse/lakefs/graveler"

	"github.com/treeverse/lakefs/forest/tree/mocks"
)

var treeID gr.TreeID

func TestSimpleTree(t *testing.T) {
	trees := InitTreesRepository(&mocks.SstMgr{})
	b := mocks.NewBatchCloser()
	//tw := treeWriter{closeAsync: b,
	//	isSplitKeyFunc: testSplitter}
	tw := treeWriter{closeAsync: b,
		splitFactor: 4_000,
		trees:       trees}
	f, err := os.Open(`testdata/test_input.csv`)
	if err != nil {
		t.Fatal("open csv failed: ", err)
	}
	var lastKey string
	var lineCount int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		key := scanner.Text()
		if key == lastKey {
			continue
		}
		lineCount++
		if key < lastKey {
			panic(" unsorted keys:" + lastKey + "  :  " + key)
		}
		lastKey = key
		r := gr.ValueRecord{
			Key: gr.Key(key),
		}
		if err := tw.WriteEntry(r); err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	treeID, err = tw.SaveTree(TreeType{})
	if err != nil {
		panic(err)
	}
	//var lineCount, lastLineCount int
	//for inp := range input {
	//	if lastKey == inp[0] {
	//		continue
	//	}
	//	lineCount++
	//	if inp[0] < lastKey {
	//		panic(" unsorted keys:" + lastKey + inp[0])
	//	}
	//	lastKey = inp[0]
	//
	//	r := gr.ValueRecord{
	//		Key: gr.Key(inp[0]),
	//	}
	//	if err := tw.WriteEntry(r); err != nil {
	//		log.Fatal(err)
	//	}
	//	if !tw.HasOpenWriter() {
	//		diff := lineCount - lastLineCount
	//		lastLineCount = lineCount
	//		fmt.Printf("writer closed line %d diff %d\n", lineCount, diff)
	//	}
	//}
	//_, err := tw.SaveTree(nil)
	//if err != nil {
	//	panic(err)
	//}

}

//func testSplitter(path rocks.Path, rowNum int) bool {
//	return true
//}

//func readGzip(output chan []string) {
//	defer close(output)
//	f, err := os.Open("simmilar.gz")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer f.Close()
//	gr, err := gzip.NewReader(f)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer gr.Close()
//
//	cr := csv.NewReader(gr)
//	l, _ := cr.Read() // skip header
//	_ = l
//	for {
//		rec, err := cr.Read()
//		if err != nil {
//			//	log.Fatal(err)
//			break
//		}
//		out := make([]string, 2)
//		//t := rec[1] + "," + rec[2]
//		t := rec[2]
//		out[0] = t
//		//out[1]= rec[3] + "," + rec[7]
//
//		out[1] = t[len(t)-50:]
//		output <- out
//	}
//
//}

//func TestCreateTestFile(t *testing.T) {
//	f, err := os.Open("simmilar.gz")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer f.Close()
//	gr, err := gzip.NewReader(f)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer gr.Close()
//
//	cr := csv.NewReader(gr)
//	l, _ := cr.Read() // skip header
//	_ = l
//	o, err := os.Create("test_input.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer o.Close()
//	for i := 0; i < 2_500_000; i++ {
//		rec, err := cr.Read()
//		if err != nil {
//			break
//		}
//		if i%5 == 0 {
//			_, err = o.WriteString(rec[2] + "\n")
//			if err != nil {
//				t.Fatal(err)
//			}
//		}
//	}
//}