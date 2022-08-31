package maps

/*
 * This file is part of the ObjectVault Project.
 * Copyright (C) 2020-2022 Paulo Ferreira <vault at sourcenotes.org>
 *
 * This work is published under the GNU AGPLv3.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

type MapWrapper struct {
	inner    map[string]interface{}
	modified bool
}

func NewMapWrapper(m map[string]interface{}) *MapWrapper {
	v := &MapWrapper{}

	// Is Incoming Map Set?
	if m != nil { // YES
		// TODO: Should we clone the incoming map (deep, shallow)
		v.inner = m
	}

	return v
}

func (o *MapWrapper) IsEmpty() bool {
	return len(o.inner) == 0
}

func (o *MapWrapper) IsModified() bool {
	return o.modified
}

func (o *MapWrapper) Map() map[string]interface{} {
	// Is Map Empty?
	if len(o.inner) == 0 {
		return nil
	}

	return o.inner
}

func (o *MapWrapper) Has(path string) bool {
	return Has(o.inner, path)
}

func (o *MapWrapper) Get(path interface{}) (interface{}, error) {
	return Get(o.inner, path)
}

func (o *MapWrapper) GetDefault(path string, d interface{}) (interface{}, error) {
	return GetDefault(o.inner, path, d)
}

func (o *MapWrapper) Set(path string, v interface{}, force bool) error {
	var m map[string]interface{}
	var e error

	// Is VALUE to be set?
	if v == nil { // NO: Clear it
		m, e = Clear(o.inner, path)
	} else { // YES
		m, e = Set(o.inner, path, v, force)
	}

	// Error Occurred?
	if e == nil { // NO
		o.inner = m
		o.modified = true
	}

	return e
}

func (o *MapWrapper) Clear(path interface{}) error {
	m, e := Clear(o.inner, path)

	// Error Occurred?
	if e == nil { // NO
		o.inner = m
		o.modified = true
	}

	return e
}

func (o *MapWrapper) ClearModified(path interface{}) bool {
	c := o.modified
	o.modified = false
	return c
}

func (o *MapWrapper) Import(s string) error {
	m, e := FromJSONString(s)
	if e == nil {
		o.inner = m
		o.modified = true
	}
	return e
}

func (o *MapWrapper) Export() string {
	json, _ := ToJSONString(o.inner)
	return json
}

func (o *MapWrapper) Reset() *MapWrapper {
	o.inner = nil
	o.modified = false
	return o
}
