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
}