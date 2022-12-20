package shorturl

import (
	"testing"
)

func TestStoreFind(t *testing.T) {
	r := NewMemRepo()
	_ = r.StoreUrl(ShortUrl{
		Url:     "_URL_01_",
		ShortId: "_SHORT_ID_01_",
	})
	_ = r.StoreUrl(ShortUrl{
		Url:     "_URL_02_",
		ShortId: "_SHORT_ID_02_",
	})

	u01 := r.Find("_SHORT_ID_01_")
	if u01 != "_URL_01_" {
		t.Errorf("Cannot find URL for ID _SHORT_ID_01_")
	}
	u02 := r.Find("_SHORT_ID_02_")
	if u02 != "_URL_02_" {
		t.Errorf("Cannot find URL for ID _SHORT_ID_02_")
	}
}
