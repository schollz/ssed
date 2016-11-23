package ssed

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/schollz/bol/utils"
)

func TestCreateDirs(t *testing.T) {
	dir, _ := homedir.Dir()
	os.RemoveAll(path.Join(dir, ".config", "ssed"))
	os.RemoveAll(path.Join(dir, ".config", "cache"))
	createDirs()
	if !utils.Exists(path.Join(dir, ".config", "ssed")) || !utils.Exists(path.Join(dir, ".cache", "ssed")) {
		t.Errorf("Problem creating dirs")
	}
}

func TestConfig(t *testing.T) {
	EraseConfig()
	dir, _ := homedir.Dir()
	configFile := path.Join(dir, ".config", "ssed", "config.json")
	Open("zack", "test", "ssh://server1")
	if !utils.Exists(configFile) {
		t.Errorf("Problem creating configuation file")
	}
	fs, _ := Open("zack", "test", "")
	if fs.ReturnMethod() != "ssh://server1" {
		t.Errorf("Problem reloading method")
	}
	_, err := Open("zack", "wrongpassword", "")
	if err == nil {
		t.Errorf("Should have an error")
	}

	// Test the setting and getting of methods in memory
	firstMethod := fs.ReturnMethod()
	fs.SetMethod("http://someothermethod")
	secondMethod := fs.ReturnMethod()
	if firstMethod != "ssh://server1" || secondMethod != "http://someothermethod" {
		t.Errorf("Problem using pointers in structs")
	}
	// Test setting bad method
	err = fs.SetMethod("badmethod")
	if err == nil {
		t.Errorf("Error should be thrown for bad method")
	}
	// Test getting changed method from disk
	fs2, _ := Open("zack", "test", "")
	if fs2.ReturnMethod() != "http://someothermethod" {
		t.Errorf("Problem with persistence of method")
	}

	// Test loading the default user with corret password
	fs2, _ = Open("", "test", "")
	if fs2.username != "zack" {
		t.Errorf("Could not load default user")
	}
	// Test loading the default user with incorrect password
	_, err = Open("", "tesjkljlt", "")
	if err == nil {
		t.Errorf("Problem with password")
	}

	// Test listing configs
	Open("zack2", "test2", "ssh://server2")
	configs := ListConfigs()
	if configs[0].Username != "zack2" || configs[1].Username != "zack" {
		t.Errorf("Error setting configs: %+v", configs) // last name should be listed first
	}

}

func TestEntries(t *testing.T) {
	// Test adding a entry
	DebugMode()
	fs, _ := Open("zack", "test", "")
	fs.Update("some text", "notes", "", "2014-11-20T13:00:00-05:00")
	fs.Update("some other test", "journal", "", "2014-11-20T13:00:00-05:00")
	fs.Update("some text2", "notes", "", "2015-11-23T13:00:00-05:00")
	fs.Update("some text3", "notes", "entry1", "2016-11-20T13:00:00-05:00")
	fs.Update("some text3, edited", "notes", "entry1", "2016-11-23T13:00:00-05:00")
	for i := 0; i < 1000; i++ {
		text := strconv.Itoa(i)
		fs.Update(text, "test", text, "")
	}

	// check if ordering is correct
	for _, entry := range fs.GetDocument("notes") {
		fmt.Println(entry.Document, entry.Timestamp, entry.Text)
	}

	fs.Close()

	// fs2, _ := Open("zack2", "test2", "http://something")
	// fs2.Update("blah", "texts", "", "2014-11-21T13:00:00-05:00")
	// fs2.Update("ghjgjgj", "texts", "", "2014-11-20T13:00:00-05:00")
	// // check if ordering is correct
	// for _, entry := range fs2.GetDocument("texts") {
	// 	fmt.Println(entry.Document, entry.Timestamp, entry.Text)
	// }
	// fs2.Close()

}

func processFile(srcFile string, num int) {
	f, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tarReader := tar.NewReader(gzf)

	i := 0
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		name := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			fmt.Println("(", i, ")", "Name: ", name)
			if i == num {
				fmt.Println(" --- ")
				io.Copy(os.Stdout, tarReader)
				fmt.Println(" --- ")
				os.Exit(0)
			}
		default:
			fmt.Printf("%s : %c %s %s\n",
				"Yikes! Unable to figure out type",
				header.Typeflag,
				"in file",
				name,
			)
		}

		i++
	}
}