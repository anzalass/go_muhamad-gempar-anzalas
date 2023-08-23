# Data Structure (Array, Slice, Map) & Function

1. Array didalam go terdapat structur data array cara deklarasi array nya

   prime := [5]int{0, 1, 2, 3, 4}
   var y [5]int = [5]int{0, 1, 2, 3}
   var country [5]string = [5]string{"indo", "usa", "uk"}

2. Slice , slice hampir mirip dengan go tetapi dia tidak perlu menentukan panjang dari array tersebut cara mendeklarasikan nya

   var number []int
   var odd_number = []int{1, 3, 4, 6, 7, 8, 9}
   numbers := []int{1, 2, 3, 4, 5, 6, 7}
   numbers = append(numbers, 111)

3. Map di go ada tipe data map, berikut cara membuat dan mendekalarasikan map, map pasti unik dan tidak bisa diurutkan

   var harga = map[string]string{
   "a": "anzalas",
   "b": "gempar",
   "c": "muhammad",
   }

   hargaa := map[string]string{"a": "anzalas", "b": "gempar", "c": "muhammad"}

   var hargaaa = make(map[string]string)

4. Function adalah sebuah fungsi berikut adalah contoh penggunaan function dengan parameter dan return value

func calculateSquare(side int) int {
return side \* side
}
