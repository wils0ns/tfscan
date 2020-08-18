package terraform

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"github.com/Masterminds/semver"
)

// File interface
type File interface {
	ResourceLookup(address string) ([]*Resource, error)
	ResourceTypes() ([]string, error)
}

// NewFile creates a new File object
func NewFile(f interface{}, r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, f)
	if err != nil {
		return err
	}

	return nil
}

// CheckVersionConstraint checks if a version matches the given constraint
func CheckVersionConstraint(version, constraint string) bool {
	sversion, err := semver.NewVersion(version)
	if err != nil {
		log.Println(err)
		return false
	}
	c, err := semver.NewConstraint(constraint)
	if err != nil {
		log.Println(err)
		return false
	}
	return c.Check(sversion)
}
