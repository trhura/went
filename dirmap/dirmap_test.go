package dirmap

import "testing"
import "github.com/stretchr/testify/assert"

func TestHas(t *testing.T) {
	d := NewDirMap()

	d.Add("lib", "/usr/lib")
	assert.Equal(t, true, d.Has("lib"))
	assert.Equal(t, false, d.Has("libi"))
}

func TestAdd (t *testing.T) {
	d := NewDirMap()

	d.Add("lib", "/usr/lib")
	assert.Equal(t, "/usr/lib", d.Get("lib"))
	assert.Equal(t, "", d.Get("libi"))

	d.Add("lib", "/usr/share/lib")
	assert.Equal(t, "/usr/share/lib", d.Get("lib"))
	set1 := []string {"/usr/share/lib", "/usr/lib"}
	assert.Equal(t, set1, d.GetAll("lib"))

	d.Add("lib", "/usr/local/share/lib")
	assert.Equal(t, "/usr/local/share/lib", d.Get("lib"))
	set2 := []string {"/usr/local/share/lib", "/usr/share/lib", "/usr/lib"}
	assert.Equal(t, set2, d.GetAll("lib"))

	d.Add ("lib", "/usr/lib")
	assert.Equal(t, "/usr/lib", d.Get("lib"))
	set3 := []string {"/usr/lib", "/usr/local/share/lib", "/usr/share/lib"}
	assert.Equal(t, set3, d.GetAll("lib"))

	d.Add ("lib", "/opt/lib")
	assert.Equal(t, "/opt/lib", d.Get("lib"))
	set4 := []string {"/opt/lib", "/usr/lib", "/usr/local/share/lib"}
	assert.Equal(t, set4, d.GetAll("lib"))
}

func TestLoadSave (t *testing.T) {
	d := NewDirMap()

	d.Add("lib", "/usr/lib")
	d.Add("lib", "/usr/share/lib")
	d.Add("lib", "/usr/local/share/lib")
	d.Add ("lib", "/usr/lib")
	d.Add("bin", "/usr/bin")
	d.Add("bin", "/usr/share/bin")
	d.Add("bin", "/usr/local/share/bin")
	d.Add ("bin", "/usr/bin")

	filename :="dirmaptest.csv"
	d.Save(filename)

	e := LoadDirMap(filename)
	assert.Equal(t, d.Len(), e.Len())
	assert.Equal(t, d.GetAll("lib"), e.GetAll("lib"))
	assert.Equal(t, d.GetAll("bin"), e.GetAll("bin"))
}
