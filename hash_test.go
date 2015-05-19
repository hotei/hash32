// Copyright 2012 Apcera Inc. All rights reserved.
// run with: go test -test.bench=".*" 


// mdr added crc[32/64/256] to the bench mix for comparison

package hash32

import (
	"hash/crc32"
	"hash/crc64"
	"testing"
	//
	"github.com/hotei/mdr"
)

var crc64Table = crc64.MakeTable(crc64.ECMA)

// Representative of small and medium subjects.
var smlKey = []byte("foo")
var medKey = []byte("apcera.continuum.router.foo.bar")
var lrgKey = []byte("alksjf;lajsdfl;jasdlfjsldjfasdjf;alskdjf;laskjdf;lkajsdf;ljasd;fljasd;lfjas;ldfja;lsdjkflakjsdf" +
	"alksjf;lajsdfl;jasdlfjsldjfasdjf;alskdjf;laskjdf;lkajsdf;ljasd;fljasd;lfjas;ldfja;lsdjkflakjsdf" +
	"alksjf;lajsdfl;jasdlfjsldjfasdjf;alskdjf;laskjdf;lkajsdf;ljasd;fljasd;lfjas;ldfja;lsdjkflakjsdf" +
	"alksjf;lajsdfl;jasdlfjsldjfasdjf;alskdjf;laskjdf;lkajsdf;ljasd;fljasd;lfjas;ldfja;lsdjkflakjsdf" +
	"alksjf;lajsdfl;jasdlfjsldjfasdjf;alskdjf;laskjdf;lkajsdf;ljasd;fljasd;lfjas;ldfja;lsdjkflakjsdf")

func Benchmark_CRC32_SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		crc32.Checksum(smlKey, crc32.IEEETable)
	}
}

func Benchmark_CRC64_SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		crc64.Checksum(smlKey, crc64Table)
	}
}

func Benchmark_CRC256_SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		mdr.BufSHA256(smlKey)
	}
}

func Benchmark_Bernstein_SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Bernstein(smlKey)
	}
}

func Benchmark_Murmur3___SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Murmur3(smlKey, M3Seed)
	}
}

func Benchmark_FNV1A_____SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		FNV1A(smlKey)
	}
}

func Benchmark_Meiyan____SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Meiyan(smlKey)
	}
}

func Benchmark_Jesteress_SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Jesteress(smlKey)
	}
}

func Benchmark_Yorikke___SmallKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Yorikke(smlKey)
	}
}

func Benchmark_CRC32_MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		crc32.Checksum(medKey, crc32.IEEETable)
	}
}

func Benchmark_CRC64_MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		crc64.Checksum(medKey, crc64Table)
	}
}

func Benchmark_CRC256_MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		mdr.BufSHA256(medKey)
	}
}

func Benchmark_Bernstein___MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Bernstein(medKey)
	}
}

func Benchmark_Murmur3_____MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Murmur3(medKey, M3Seed)
	}
}

func Benchmark_FNV1A_______MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		FNV1A(medKey)
	}
}

func Benchmark_Meiyan______MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Meiyan(medKey)
	}
}

func Benchmark_Jesteress___MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Jesteress(medKey)
	}
}

func Benchmark_Yorikke_____MedKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Yorikke(medKey)
	}
}

func Benchmark_CRC32_LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		crc32.Checksum(lrgKey, crc32.IEEETable)
	}
}

func Benchmark_CRC64_LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		crc64.Checksum(lrgKey, crc64Table)
	}
}

func Benchmark_CRC256_LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		mdr.BufSHA256(lrgKey)
	}
}

func Benchmark_Bernstein___LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Bernstein(lrgKey)
	}
}

func Benchmark_Murmur3_____LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Murmur3(lrgKey, M3Seed)
	}
}

func Benchmark_FNV1A_______LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		FNV1A(lrgKey)
	}
}

func Benchmark_Meiyan______LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Meiyan(lrgKey)
	}
}

func Benchmark_Jesteress___LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Jesteress(lrgKey)
	}
}

func Benchmark_Yorikke_____LrgKey(b *testing.B) {
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		Yorikke(lrgKey)
	}
}
