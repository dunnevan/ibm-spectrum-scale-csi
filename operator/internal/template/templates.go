package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"

	scalev1 "github.com/IBM/ibm-spectrum-scale-csi/operator/apis/scale/v1"
)

var (
	yamls *template.Template

	//TODO: use the scheme we setup in main.go
	decode runtime.Decoder = scheme.Codecs.UniversalDeserializer()
)

func init() {
	cwd, _ := os.Getwd()

	yamls = template.New("csi")

	yamls = template.Must(parseGlob(
		yamls,
		filepath.Join(cwd, "./templates/attacher", "*.yaml")))

	yamls = template.Must(parseGlob(
		yamls,
		filepath.Join(cwd, "./templates/driver", "*.yaml")))

	yamls = template.Must(parseGlob(
		yamls,
		filepath.Join(cwd, "./templates/provisioner", "*.yaml")))

	yamls = template.Must(parseGlob(
		yamls,
		filepath.Join(cwd, "./templates/snapshotter", "*.yaml")))
}

func ExecuteTemplate(cr *scalev1.CSI, templateName string, into runtime.Object) error {
	var b bytes.Buffer

	err := yamls.ExecuteTemplate(&b, templateName, cr)
	if err != nil {
		return err
	}

	_, _, err = decode.Decode(b.Bytes(), nil, into)
	if err != nil {
		return err
	}

	return nil
}

func Execute(cr *scalev1.CSI, group string, into runtime.Object) error {

	intoType := reflect.TypeOf(into).Elem()
	templateName := strings.ToLower(path.Join(group, intoType.Name()))

	return ExecuteTemplate(cr, templateName, into)
}

func parseFiles(t *template.Template, filenames ...string) (*template.Template, error) {
	if len(filenames) == 0 {
		return nil, fmt.Errorf("template: no files named in call to parseFiles")
	}
	for _, f := range filenames {
		b, err := ioutil.ReadFile(filepath.Clean(f))
		if err != nil {
			return nil, err
		}
		s := string(b)
		filename := filepath.Base(f)
		fileext := filepath.Ext(f)
		group := filepath.Base(filepath.Dir(f))
		name := path.Join(group, strings.TrimSuffix(filename, fileext))

		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func parseGlob(t *template.Template, pattern string) (*template.Template, error) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(filenames) == 0 {
		return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
	}
	return parseFiles(t, filenames...)
}
