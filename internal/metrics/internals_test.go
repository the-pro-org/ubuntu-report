package metrics

import (
	"flag"
	"path"
	"testing"

	"github.com/ubuntu/ubuntu-report/internal/helper"
)

/*
 * Tests here some private functions to gather metrics
 * Collect() public API is calling out a lot of functions,
 * that's why we add some unit tests on direct Collect() callees here
 * for finer-graind results in case of failure.
 */

var update = flag.Bool("update", false, "update golden files")

func TestInstallerInfo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		root string
	}{
		{"regular", "testdata/good"},
		{"empty file", "testdata/empty"},
		{"doesn't exist", "testdata/none"},
		{"garbage content", "testdata/garbage"},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable for parallel execution
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := helper.Asserter{T: t}

			m := newTestMetrics(t, WithRootAt(tc.root))
			got := []byte(*m.installerInfo())
			want := helper.LoadOrUpdateGolden(path.Join(m.root, "gold", "intallerInfo"), got, *update, t)

			a.Equal(string(got), string(want))
		})
	}
}

func TestUpgradeInfo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		root string
	}{
		{"regular", "testdata/good"},
		{"empty file", "testdata/empty"},
		{"doesn't exist", "testdata/none"},
		{"garbage content", "testdata/garbage"},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable for parallel execution
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := helper.Asserter{T: t}

			m := newTestMetrics(t, WithRootAt(tc.root))
			got := []byte(*m.upgradeInfo())
			want := helper.LoadOrUpdateGolden(path.Join(m.root, "gold", "upgradeInfo"), got, *update, t)

			a.Equal(string(got), string(want))
		})
	}
}

func TestGetVersion(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		root string

		want string
	}{
		{"regular", "testdata/good", "18.04"},
		{"empty file", "testdata/empty", ""},
		{"doesn't exist", "testdata/none", ""},
		{"garbage content", "testdata/garbage", ""},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable for parallel execution
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := helper.Asserter{T: t}

			m := newTestMetrics(t, WithRootAt(tc.root))
			got := m.getVersion()

			a.Equal(got, tc.want)
		})
	}
}

func TestGetRAM(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		root string

		want string
	}{
		{"regular", "testdata/good", "8048100"},
		{"empty file", "testdata/empty", ""},
		{"doesn't exist", "testdata/none", ""},
		{"garbage content", "testdata/garbage", ""},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable for parallel execution
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := helper.Asserter{T: t}

			m := newTestMetrics(t, WithRootAt(tc.root))
			got := m.getRAM()

			a.Equal(got, tc.want)
		})
	}
}

func TestGetTimeZone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		root string

		want string
	}{
		{"regular", "testdata/good", "Europe/Paris"},
		{"empty file", "testdata/empty", ""},
		{"doesn't exist", "testdata/none", ""},
		{"garbage content", "testdata/garbage", ""},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable for parallel execution
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := helper.Asserter{T: t}

			m := newTestMetrics(t, WithRootAt(tc.root))
			got := m.getTimeZone()

			a.Equal(got, tc.want)
		})
	}
}

func newTestMetrics(t *testing.T, fixtures ...func(m *Metrics) error) Metrics {
	t.Helper()
	m, err := New(fixtures...)
	if err != nil {
		t.Fatal("can't create metrics object", err)
	}
	return m
}