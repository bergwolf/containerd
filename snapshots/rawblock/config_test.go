// +build linux,!no_rawblock

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package rawblock

import (
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestConfigSetDefaults(t *testing.T) {
	c := &SnapshotterConfig{}
	c.setDefaults("foo")

	assert.Equal(t, c.RootPath, "foo")
	assert.Equal(t, c.SizeMB, uint32(defaultImageSizeMB))
	assert.Equal(t, c.FsType, defaultFsType)

	c.FsType = "xfs"
	c.RootPath = "bar"

	c.setDefaults("foo")
	assert.Equal(t, c.RootPath, "bar")
	assert.Equal(t, c.FsType, "xfs")
	assert.Assert(t, strings.Contains(strings.Join(c.Options, ","), "nouuid"))
}

func TestConfigValidate(t *testing.T) {
	c := &SnapshotterConfig{}

	err := c.validate()
	assert.Assert(t, err != nil, "empty snapshotter config should fail validation")

	c.setDefaults("")
	err = c.validate()
	assert.NilError(t, err)
}
