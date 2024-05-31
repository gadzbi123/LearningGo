package main

import (
	"testing"

	_ "github.com/google/pprof/profile"
)

func BenchmarkPerformance(b *testing.B) {
	b.ReportAllocs()
	main()
}

// func TestProfile(t *testing.T) {
// 	// Create an empty new Profile we can start adding stuff to.
// 	p := profile.Profile{}

// 	// SampleType specifies what type the samples are of.
// 	// Usually shown in the top of the profile visualization.
// 	p.SampleType = []*profile.ValueType{{
// 		Type: "alloc_space", // We want to have a profile that describes the allocations
// 		Unit: "bytes",       // in bytes.
// 	}, {Type: "cpu", Unit: "%"}}

// 	// Create a new file to write the profile to.
// 	f, err := os.Create("profile.pb.gz")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// Write the profile to the file.
// 	if err := p.Write(f); err != nil {
// 		t.Fatal(err)
// 	}
// }
