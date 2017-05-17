package ricebox

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/mattes/migrate/source"
)

type Ricebox struct {
	box        *rice.Box
	migrations *source.Migrations
}

func WithInstance(box *rice.Box) (source.Driver, error) {
	migrations := source.NewMigrations()
	err := box.Walk("", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		m, err := source.DefaultParse(info.Name())
		if err != nil {
			return nil
		}
		if !migrations.Append(m) {
			return fmt.Errorf("unable to parse file %v", info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Ricebox{
		box:        box,
		migrations: migrations,
	}, nil
}

func (r *Ricebox) Open(sourceURL string) (source.Driver, error) {
	return nil, errors.New("ricebox: creating source with Open() is not supported, please use WithInstance() to pass rice-box that was opened in application package")
}

func (r *Ricebox) Close() error {
	return nil
}

func (r *Ricebox) First() (version uint, err error) {
	if v, ok := r.migrations.First(); !ok {
		return 0, os.ErrNotExist
	} else {
		return v, nil
	}
}

func (r *Ricebox) Prev(version uint) (prevVersion uint, err error) {
	if v, ok := r.migrations.Prev(version); !ok {
		return 0, os.ErrNotExist
	} else {
		return v, nil
	}
}

func (r *Ricebox) Next(version uint) (nextVersion uint, err error) {
	if v, ok := r.migrations.Next(version); !ok {
		return 0, os.ErrNotExist
	} else {
		return v, nil
	}
}

func (r *Ricebox) ReadUp(version uint) (io.ReadCloser, string, error) {
	if m, ok := r.migrations.Up(version); ok {
		bs, err := r.box.Bytes(m.Raw)
		if err != nil {
			return nil, "", err
		}
		return ioutil.NopCloser(bytes.NewBuffer(bs)), m.Identifier, nil
	}
	return nil, "", os.ErrNotExist
}

func (r *Ricebox) ReadDown(version uint) (io.ReadCloser, string, error) {
	if m, ok := r.migrations.Down(version); ok {
		bs, err := r.box.Bytes(m.Raw)
		if err != nil {
			return nil, "", err
		}
		return ioutil.NopCloser(bytes.NewBuffer(bs)), m.Identifier, nil
	}
	return nil, "", os.ErrNotExist
}
