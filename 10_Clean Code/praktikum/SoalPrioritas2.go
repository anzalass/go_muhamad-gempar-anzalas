package main

type Kendaraan struct {
	TotalRoda int

	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {

	m.TambahKecepatan(10)

}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {

	m.KecepatanPerJam = m.KecepatanPerJam + kecepatanBaru

}

func main() {

	mobilcepat := Mobil{}

	mobilcepat.Berjalan()

	mobilcepat.Berjalan()

	mobilcepat.Berjalan()

	mobillamban := Mobil{}

	mobillamban.Berjalan()

}
