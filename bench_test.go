package raftmemlog

import (
	"os"
	"testing"

	"github.com/hashicorp/raft/bench"
)

func BenchmarkMemLogStore_FirstIndex(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.FirstIndex(b, store)
}

func BenchmarkMemLogStore_LastIndex(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.LastIndex(b, store)
}

func BenchmarkMemLogStore_GetLog(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.GetLog(b, store)
}

func BenchmarkMemLogStore_StoreLog(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.StoreLog(b, store)
}

func BenchmarkMemLogStore_StoreLogs(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.StoreLogs(b, store)
}

func BenchmarkMemLogStore_DeleteRange(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.DeleteRange(b, store)
}

func BenchmarkMemLogStore_Set(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.Set(b, store)
}

func BenchmarkMemLogStore_Get(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.Get(b, store)
}

func BenchmarkMemLogStore_SetUint64(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.SetUint64(b, store)
}

func BenchmarkMemLogStore_GetUint64(b *testing.B) {
	store := testMemLogStore(b)
	defer store.Close()
	defer os.Remove(store.path)

	raftbench.GetUint64(b, store)
}
