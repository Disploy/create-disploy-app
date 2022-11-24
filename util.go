package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func GetFirstHeading(markdown string) string {
	head := strings.Split(markdown, "<div align=\"center\">")[1]
	head = strings.Split(head, "</div>")[0]

	r := strings.NewReader(head)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return ""
		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "h3" {
				z.Next()
				return strings.TrimSpace(z.Token().Data)
			}
		}
	}

}

func WriteToJson(file string, property string, value string, subproperty ...string) {
	f, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bb, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc := make(map[string]interface{})

	if err := json.Unmarshal(bb, &doc); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc[property] = value

	data, err := json.MarshalIndent(doc, "", "\t")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile(file, data, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func PkgInstallCommand(manager string) string {
	if manager != "yarn" {
		return "install"
	} else {
		return "add"
	}
}

func Copy(dir, project string) {
	os.Rename(".disploy/create-disploy-app-main/assets/"+dir, project)

	WriteToJson(project+"/package.json", "name", project)
	WriteToJson(project+"/disploy.json", "name", project)

	m, ok := PackageManagerChoiceModel().(PackageManagerOptionStruct)

	if !ok {
		fmt.Println("Error: PackageManagerChoiceModel() is not PackageManagerOptionStruct")
		os.Exit(1)
	}

	cmd := exec.Command(m.choice, PkgInstallCommand(m.choice), "disploy@dev")
	cmd.Dir = project

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.RemoveAll(".disploy")
	os.Remove(".disploy.zip")
}

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func UnzipFile(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
