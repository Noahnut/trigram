package trigram

import (
	"os"
	"testing"
)

func TestCreateContextIndex(t *testing.T) {
	tri := NewTrigram()
	testFileMap := make(map[string]string)

	testFileMap["testFileOne"] = "abc"

	tri.createContextIndex(testFileMap["testFileOne"], "testFileOne")

	if tri.trigramMap[6382179][0] != "testFileOne" {
		t.Error("Index is wrong")
	}

	testFileMap["testFileTwo"] = "abcd"

	tri.createContextIndex(testFileMap["testFileTwo"], "testFileTwo")

	if len(tri.trigramMap) != 2 {
		t.Error("Index generator is wrong")
	}

	if tri.trigramMap[6382179][0] != "testFileOne" && tri.trigramMap[6382179][1] != "testFileTwo" {
		t.Error("Index is wrong")
	}

	if tri.trigramMap[6447972][0] != "testFileTwo" {
		t.Error("Index is wrong")
	}
}

func createFakeFile(fileName string, context string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(context)

	file.Sync()

	return nil
}

func TestSimpleTrig(t *testing.T) {
	tri := NewTrigram()
	createFakeFile("testOne", "Code Search")
	createFakeFile("testTwo", "Code Search")
	createFakeFile("testThree", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris sed rhoncus turpis. Nullam gravida dictum justo, sit amet luctus orci lacinia eu. Sed convallis lacus ac eros vestibulum, sit amet lacinia nulla luctus. Aenean facilisis velit vitae elit pellentesque, at vulputate tellus suscipit. Integer aliquet euismod tincidunt. Etiam dignissim fermentum turpis a ultricies. Vestibulum turpis lacus, cursus ut dictum vitae, viverra vel ante. Curabitur et tortor tellus. Aliquam posuere in ipsum ut vehicula.")

	err := tri.Add("testOne")

	if err != nil {
		t.Fatal(err)
	}

	err = tri.Add("testTwo")

	if err != nil {
		t.Fatal(err)
	}

	err = tri.Add("testThree")

	if err != nil {
		t.Fatal(err)
	}

	result := tri.Find("Search")

	if result[0] != "testOne" && result[1] != "testTwo" {
		t.Error("Wrong to find the fileName")
	}

	result = tri.Find("Cod")

	if result[0] != "testOne" && result[1] != "testTwo" {
		t.Error("Wrong to find the fileName")
	}

	result = tri.Find("consectetur adipiscing elit")

	if result[0] != "testThree" {
		t.Error("Wrong to find the fileName")
	}

	result = tri.Find("sit amet, ")

	if result[0] != "testThree" {
		t.Error("Wrong to find the fileName")
	}

	result = tri.Find("ullam gr")

	if result[0] != "testThree" {
		t.Error("Wrong to find the fileName")
	}

	os.Remove("testOne")
	os.Remove("testTwo")
	os.Remove("testThree")
}
