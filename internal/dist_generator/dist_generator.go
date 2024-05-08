package dist_generator

import (
	"bytes"
	"io"
	"io/fs"
	"nausea-web/internal/compress"
	"nausea-web/internal/minify"
	"os"
	"path/filepath"
)

type DistCreator struct {
	Minifier    minify.Minifier
	Compressor  compress.Compressor
	AssetsRoot  string
	currentRoot string
	dist        string
}

func NewDistCreator(root string, m minify.Minifier, c compress.Compressor) *DistCreator {
	dist := "./dist"
	return &DistCreator{
		Minifier:    m,
		Compressor:  c,
		AssetsRoot:  root,
		currentRoot: root,
		dist:        dist,
	}
}

func (w *DistCreator) prepareFile(data []byte, path string) ([]byte, error) {
	file := bytes.NewReader(data)
	var bufferToCompress io.Reader
	if w.Minifier.CanMinify(path) {
		var minified bytes.Buffer
		err := w.Minifier.Minify(&minified, file, path)
		if err != nil {
			return []byte{}, err
		}
		bufferToCompress = &minified
	} else {
		bufferToCompress = file
	}
	var compressed bytes.Buffer
	err := w.Compressor.Compress(&compressed, bufferToCompress)
	return compressed.Bytes(), err
}

func (w *DistCreator) Walker(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	dist := filepath.Join(w.dist, path)
	if info.IsDir() {
		if err := os.MkdirAll(dist, 0777); err != nil {
			return err
		}
		if path == w.currentRoot || path == w.AssetsRoot {
			return nil
		}
		w.currentRoot = path
		return w.Walk(path)
	}
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	data, err := w.prepareFile(fileBytes, path)
	err = os.WriteFile(dist, data, 0777)
	return err
}

func (w *DistCreator) Walk(path string) error {
	err := filepath.Walk(path, w.Walker)
	return err
}

func (w *DistCreator) RebuildDist() error {
	err := os.RemoveAll(w.dist)
	if err != nil {
		return err
	}
	return filepath.Walk(w.AssetsRoot, w.Walker)
}
