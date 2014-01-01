// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestConvertM3uAndOutputPls(t *testing.T) {
	file, err := os.Create("testfile.pls")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	ConvertM3uAndOutputPls("David-Bowie-Singles.m3u", file)

	output, err := ioutil.ReadFile("testfile.pls")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := ioutil.ReadFile("David-Bowie-Singles.pls")
	if err != nil {
		t.Fatal(err)
	}
	if string(output) != string(expected) {
		t.Error("Generated pls file is not as expected")
	}
}

func TestConvertPlsAndOutputM3u(t *testing.T) {
	file, err := os.Create("testfile.m3u")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	ConvertPlsAndOutputM3u("David-Bowie-Singles.pls", file)

	output, err := ioutil.ReadFile("testfile.m3u")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := ioutil.ReadFile("David-Bowie-Singles.m3u")
	if err != nil {
		t.Fatal(err)
	}
	if string(output) != string(expected) {
		t.Error("Generated m3u file is not as expected")
	}
}
