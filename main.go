package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	N    = 100
	M    = 100
	MaxV = 1000
)

var (
	dis       [][]int
	nghiem    [][]int
	f         []int
	minNghiem = 100000000
)

func main() {
	nghiem = make([][]int, M)
	f = make([]int, N)
	initDistance()
	initNghiem()
	for i := 0; i < 10000; i++ {
		danhGia()
		print()
		luaChon()
		laiGhep()
		dotBien()
	}
	fmt.Println("Best result: ", minNghiem)
}

func print() {
	cl := f
	sort.Ints(cl)
	fmt.Println(cl[0])
	if minNghiem > cl[0] {
		minNghiem = cl[0]
	}
	// for i := 0; i < N; i++ {
	// 	if f[i] == cl[0] {
	// 		for _, v := range nghiem[i] {
	// 			fmt.Print(v, " ")
	// 		}
	// 		fmt.Println()
	// 	}
	// }
}

func dotBien() {
	rand.Seed(time.Now().UTC().UnixNano())
	var (
		a = rand.Intn(N)
		b = rand.Intn(M)
		c = rand.Intn(M)
	)

	var tmp = nghiem[a][b]
	nghiem[a][b] = nghiem[a][c]
	nghiem[a][c] = tmp
}

func laiGhep() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 20; i++ {
		var cha = rand.Intn(N)
		var me = rand.Intn(N)
		var vitri = rand.Intn(M-1) + 1
		// fmt.Println(cha, me, vitri)
		con1 := LG(cha, me, vitri)
		con2 := LG(me, cha, vitri)
		nghiem[cha] = con1
		nghiem[me] = con2
	}
}

func LG(cha int, me int, vitri int) (con []int) {
	con = make([]int, M)
	trung := make([]int, M)
	for i := 0; i < vitri; i++ {
		con[i] = nghiem[cha][i]
		trung[con[i]] = 1
	}

	for i := vitri; i < M; i++ {
		if trung[nghiem[me][i]] == 0 {
			con[i] = nghiem[me][i]
			trung[con[i]] = 1
		} else {
			con[i] = -1
		}
	}

	j := 0
	for i := vitri; i < M; i++ {
		if con[i] == -1 {
			for trung[j] == 1 {
				j++
			}
			con[i] = j
			trung[j] = 1
		}
	}

	return
}

func luaChon() {
	var cl []int
	cl = f
	sort.Ints(cl)
	nguong := cl[N*80/100]
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < N; i++ {
		if f[i] > nguong {
			nghiem[i] = nghiem[rand.Intn(N)]
		}
	}
}

func danhGia() {
	for i := 0; i < N; i++ {
		f[i] = quangDuong(nghiem[i])
	}
}

func quangDuong(a []int) int {
	var s = 0
	for i := 0; i < len(a)-1; i++ {

		s += dis[a[i]][a[i+1]]
	}

	return s + dis[a[0]][a[len(a)-1]]
}

func initDistance() {
	rand.Seed(1)
	dis = make([][]int, M)
	for i := 0; i < M; i++ {
		dis[i] = make([]int, M)
		for j := 0; j < M; j++ {
			dis[i][j] = rand.Intn(MaxV)
		}
	}
}

func initNghiem() {
	for i := 0; i < N; i++ {
		nghiem[i] = make([]int, M)
		nghiem[i] = ChinhHop()
	}
}

func ChinhHop() []int {
	rand.Seed(time.Now().UTC().UnixNano())
	a := make([]int, M)
	for i := 0; i < M; i++ {
		a[i] = i
	}

	for i := 0; i < 100; i++ {
		x := rand.Intn(M)
		y := rand.Intn(M)
		tmp := a[x]
		a[x] = a[y]
		a[y] = tmp
	}
	return a
}
