package gowurfl

import (
	"testing"
)

func TestNew(t *testing.T) {
	w, err := New()
	defer w.Close()

	if err != nil {
		t.Errorf("failed to create wurfl handle: %s", err)
	}
}

var rootFile = "/usr/share/wurfl/wurfl.xml"

var uas = []string{
	"Mozilla/5.0 (PLAYSTATION 3; 3.55)",
	"Opera/9.80 (J2ME/MIDP; Opera Mini/9.80 (S60; SymbOS; Opera Mobi/23.348; U; en) Presto/2.5.25 Version/10.54",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2715.0 Safari/537.36",
	"Dillo/2.0",
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X; en-US) AppleWebKit/528.16 (KHTML, like Gecko, Safari/528.16) OmniWeb/v622.8.0.112941",
	"NetSurf/2.0 (RISC OS; armv5l)",
	"Mozilla/5.0 (Linux; Android 4.4.2; Lenovo S860 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.105 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; U; Android 2.3.6; es-es; LG-E400 Build/GRK39F) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"Mozilla/5.0 (Linux; Android 4.1.2; C1905 Build/15.1.C.2.8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.128 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; U; Android 4.1.2; it-it; GT-I8190 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	"Mozilla/5.0 (Linux; Android 4.1.2; Nokia_XL Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.72 Mobile Safari/537.36 OPR/19.0.1340.71596",
	"Mozilla/5.0 (Linux; U; Android 4.0.4; es-es; SAMSUNG GT-S7560/S7560XXAME9 Build/IMM76I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	"Mozilla/5.0 (Linux; Android 4.1.2; Nokia_XL Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.72 Mobile Safari/537.36 OPR/19.0.1340.71596",
	"Mozilla/5.0 (Linux; U; Android 4.3; es-es; SAMSUNG GT-I9300/I9300XXUGNA5 Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	"Mozilla/5.0 (Linux; Android 4.4.2; Lenovo S860 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.105 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 4.3; C1905 Build/15.4.A.1.9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.128 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 4.2.2; C6603 Build/10.3.1.A.2.67) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 4.2.2; GT-I9060 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 4.4.2; Lenovo S860 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.105 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 6.0; Infinix X510 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Mobile Safari/537.36 OPR/36.1.2126.102083",
	"Mozilla/5.0 (Linux; Android 5.1.1; Infinix X510 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Mobile Safari/537.36 OPR/36.1.2126.102083",
	"Mozilla/5.0 (Linux; Android 6.0; Infinix X510 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Mobile Safari/537.36 OPR/36.1.2126.102083",
	"Mozilla/5.0 (Linux; U; Android 4.3; es-es; Orange Gova Build/OrangeGova) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	"Mozilla/5.0 (Linux; Android 6.0; Infinix X510 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Mobile Safari/537.36 OPR/36.1.2126.102083",
	"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; NOKIA; Lumia 520) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.10586",
	"Mozilla/5.0 (Linux; U; Android 2.3.5; hu-hu; SAMSUNG GT-S5830/S5830BUKS2 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"Mozilla/5.0 (Linux; Android 4.4.4; SM-G316HU Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.64 Mobile Safari/537.36 OPR/36.0.2126.101126",
	"Mozilla/5.0 (Linux; U; Android 4.1.2; es-es; Elektra L Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	"Mozilla/5.0 (Linux; Android 4.2.2; C2105 Build/15.3.A.1.14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.128 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; U; Android 4.4.4; en-gb; SM-G360H Build/KTU84P) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
}

func testNewEngine(t testing.TB) *WURFL {
	w, err := New()

	if err != nil {
		t.Fatal(err)
	}

	return w
}

func testLoadRepository(p string, w *WURFL, t testing.TB) {
	if err := w.SetRoot(p); err != nil {
		t.Fatal(err)
	}

	if err := w.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestGetEngineTarget(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()

	et := w.GetEngineTarget()

	if et != EngineTargetHighPerformance {
		t.Errorf("expected %v as default engine target but got %v", EngineTargetHighPerformance, et)
	}
}

func TestSetEngineTarget(t *testing.T) {
	tcs := []struct {
		in   EngineTarget
		out  EngineTarget
		fail bool
	}{
		{EngineTargetHighAccuracy, EngineTargetHighAccuracy, false},
		{EngineTargetHighPerformance, EngineTargetHighPerformance, false},
		{EngineTargetInvalid, EngineTargetHighPerformance, true},
	}

	for _, tc := range tcs {
		w := testNewEngine(t)
		defer w.Close()

		err := w.SetEngineTarget(tc.in)
		if (err == nil) == tc.fail {
			t.Errorf("SetEngineTarget(%v) expected to fail but did not", tc.in)
		}

		et := w.GetEngineTarget()
		if et != tc.out {
			t.Errorf("GetEngineTarget() expected %v but got %v", tc.out, et)
		}
	}
}

func TestSetCacheProvider(t *testing.T) {
	tcs := []struct {
		in    CacheProvider
		sizes []int
		fail  bool
	}{
		{CacheProviderNone, []int{}, false},
		{CacheProviderNone, []int{0}, false},
		{CacheProviderNone, []int{0, 0}, false},
		{CacheProviderNone, []int{10000, 3000}, false},
		{CacheProviderLRU, []int{}, false},
		{CacheProviderLRU, []int{0}, true},
		{CacheProviderLRU, []int{10000}, false},
		{CacheProviderLRU, []int{10000, 3000}, false},
		{CacheProviderLRU, []int{0, 3000}, true},
		{CacheProviderDoubleLRU, []int{}, false},
		{CacheProviderDoubleLRU, []int{0}, false},
		{CacheProviderDoubleLRU, []int{10000, 3000}, false},
		{CacheProviderDoubleLRU, []int{10000, 0}, true},
		{CacheProviderDoubleLRU, []int{0, 3000}, true},
		{CacheProviderDoubleLRU, []int{0, 0}, true},
	}

	for _, tc := range tcs {
		w := testNewEngine(t)
		defer w.Close()

		var err error
		switch len(tc.sizes) {
		default:
			t.Fatal("invalid sizes array")
		case 0:
			err = w.SetCacheProvider(tc.in)
		case 1:
			err = w.SetCacheProvider(tc.in, tc.sizes[0])
		case 2:
			err = w.SetCacheProvider(tc.in, tc.sizes[0], tc.sizes[1])
		}

		if (err == nil) == tc.fail {
			t.Errorf("SetCacheProvider(%v, %v) expected to fail but did not", tc.in, tc.sizes)
		}
	}
}

func diff(a, b []string) []string {
	d := make(map[string]bool)

	for _, e := range b {
		d[e] = true
	}

	var ret []string
	for _, e := range a {
		if !d[e] {
			ret = append(ret, e)
		}
	}

	return ret
}

func TestGetMandatoryCapabilities(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()

	caps, err := w.GetMandatoryCapabilities()

	if err != nil {
		t.Errorf("GetMandatoryCapabilities() failed with: %s", err)
	}

	d := diff(MandatoryCapabilities, caps)
	if len(d) > 0 {
		t.Errorf("MandatoryCapabilities differ:\nwant: %v \nhave: %v\ndiff: %v", caps, MandatoryCapabilities, d)
	}
}

func TestLookupUserAgent(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()
	testLoadRepository(rootFile, w, t)

	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2715.0 Safari/537.36"
	_, err := w.LookupUserAgent(ua)

	if err != nil {
		t.Errorf("LookupUserAgent(%q) failed with: %s", ua, err)
	}
}

func TestDeviceGetVirtualCapability(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()
	testLoadRepository(rootFile, w, t)

	for _, ua := range uas {
		d, err := w.LookupUserAgent(ua)
		defer d.Close()
		if err != nil {
			t.Errorf("LookupUserAgent(%q) failed with: %s", ua, err)
		}

		_, err = d.GetVirtualCapability("advertised_device_os")
		if err != nil {
			t.Errorf("GetVirtualCapability() failed with: %s", err)
		}
	}
}

func TestGetCapabilities(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()
	testLoadRepository(rootFile, w, t)

	caps, err := w.GetCapabilities()

	if err != nil {
		t.Errorf("GetCapabilities() failed with: %s", err)
	}

	if len(caps) <= len(MandatoryCapabilities) {
		t.Errorf("default capability loading behaviour changed")
	}
}

func TestAddRequestedCapabilities(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()

	if err := w.AddRequestedCapabilities(MandatoryCapabilities); err != nil {
		t.Errorf("AddRequestedCapabilities() failed with: %s", err)
	}

	testLoadRepository(rootFile, w, t)

	for _, cap := range MandatoryCapabilities {
		if !w.HasCapability(cap) {
			t.Errorf("should have capability %q", cap)
		}
	}
}

func TestDeviceGetVirtualCapabilities(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()
	testLoadRepository(rootFile, w, t)

	for _, ua := range uas {
		d, err := w.LookupUserAgent(ua)
		defer d.Close()
		if err != nil {
			t.Errorf("LookupUserAgent(%q) failed with: %s", ua, err)
		}

		_, err = d.GetVirtualCapabilities()
		if err != nil {
			t.Errorf("GetVirtualCapability() failed with: %s", err)
		}

		//fmt.Printf("%s has os: %s\n", ua, caps)
	}
}

func TestDeviceGetID(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()
	testLoadRepository(rootFile, w, t)

	for _, ua := range uas {
		d, err := w.LookupUserAgent(ua)
		defer d.Close()
		if err != nil {
			t.Errorf("LookupUserAgent(%q) failed with: %s", ua, err)
		}

		id, err := d.GetID()
		if err != nil {
			t.Errorf("GetID(%q) failed with: %s", id, err)
		}
	}
}

func TestLoad(t *testing.T) {
	w := testNewEngine(t)
	defer w.Close()

	if err := w.Load(); err == nil {
		t.Errorf("load should fail without a prior call to SetRoot()")
	}
}

func BenchmarkLoadWithDefaultCapabilities(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping load benchmark in short mode")
	}

	for i := 0; i < b.N; i++ {
		w := testNewEngine(b)
		testLoadRepository(rootFile, w, b)
		w.Close()
	}
}

func BenchmarkLoadWithMandatoryCapabilities(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping load benchmark in short mode")
	}

	for i := 0; i < b.N; i++ {
		w := testNewEngine(b)

		if err := w.AddRequestedCapabilities(MandatoryCapabilities); err != nil {
			b.Errorf("AddRequestedCapabilities() failed with: %s", err)
		}

		testLoadRepository(rootFile, w, b)

		w.Close()
	}
}

func BenchmarkLookupWithDefaultCapabilities(b *testing.B) {
	w := testNewEngine(b)
	defer w.Close()
	testLoadRepository(rootFile, w, b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ua := range uas {
			d, err := w.LookupUserAgent(ua)
			if err != nil {
				d.Close()
				b.Errorf("LookupUserAgent(%q) failed with: %s", ua, err)
			}
			d.Close()
		}
	}
}

func BenchmarkLookupWithMandatoryCapabilities(b *testing.B) {
	w := testNewEngine(b)
	defer w.Close()

	if err := w.AddRequestedCapabilities(MandatoryCapabilities); err != nil {
		b.Errorf("AddRequestedCapabilities() failed with: %s", err)
	}

	testLoadRepository(rootFile, w, b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ua := range uas {
			d, err := w.LookupUserAgent(ua)
			if err != nil {
				d.Close()
				b.Errorf("LookupUserAgent(%q) failed with: %s", ua, err)
			}
			d.Close()
		}
	}
}
