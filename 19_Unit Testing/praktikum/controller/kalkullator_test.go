package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKalkulatorPlus(t *testing.T) {
	res := KalkulatorPlus(3, 2, 5)
	assert.Equal(t, float64(10), res)

	if KalkulatorPlus(3, 2, 5) != 10 {
		t.Error("Seharusnya 10")
	}
	if KalkulatorPlus(3, -2, 5) != 6 {
		t.Error("Seharusnya 6")
	}

}

func TestKalkulatorMinus(t *testing.T) {
	res := KalkulatorMinus(10, 5)
	assert.Equal(t, 5, res)

	if KalkulatorMinus(3, 2, 5) != -4 {
		t.Error("Seharusnya -4")
	}
	if KalkulatorMinus(3, -2, 5) != 0 {
		t.Error("Seharusnya 10")
	}
}

func TestKalkulatorKali(t *testing.T) {
	res := KalkulatorKali(10, 5)
	assert.Equal(t, 50, res)

	if KalkulatorKali(3, 2, 5) != 30 {
		t.Error("Seharusnya 30")
	}
	if KalkulatorKali(3, 3, 5) != 45 {
		t.Error("Seharusnya 45")
	}
}
func TestKalkulatorBagi(t *testing.T) {
	res := KalkulatorBagi(10, 5)
	assert.Equal(t, float64(2), res)

	if KalkulatorBagi(10, 5) != 2 {
		t.Error("Seharusnya 2")
	}
	if KalkulatorBagi(18, 6) != 3 {
		t.Error("Seharusnya 3")
	}
}
