// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavutil
//#include <libavutil/dict.h>
import "C"

import (
	//"unsafe"
)

func (d *Dictionary) Count() int {
	return int(C.av_dict_count((*C.struct_AVDictionary)(d)))
}

func (d *Dictionary) Get(key string) *DictionaryEntry {
	var entry *DictionaryEntry;
	entry = (*DictionaryEntry)(C.av_dict_get((*C.struct_AVDictionary)(d), C.CString(key), (*C.struct_AVDictionaryEntry)(entry), C.AV_DICT_MATCH_CASE));

	return entry
}

func (de *DictionaryEntry) Value() string {
	return C.GoString(de.value)
}