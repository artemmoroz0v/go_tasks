package random
import "math/rand"
import "github.com/e154/vydumschik"
import "github.com/essentialkaos/translit"
import "log"
import "time"
import "encoding/json"
import "L0/mainstruct"

func Random_variable(information []byte) []byte {
	var variable mainstruct.MainStruct
	json.Unmarshal(information, &variable)
	size := 19
	var order_uid string
	rand.Seed(time.Now().UnixNano())
	bigchars := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	smallchars := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	numbers := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < size; i++ {
		ok := rand.Intn(10)
		if ok < 3 {
			order_uid += bigchars[rand.Intn(26)]
		} else if ok >= 3 && ok <= 8 {
			order_uid += smallchars[rand.Intn(26)]
		} else {
			order_uid += numbers[ok]
		}
	}
	variable.Order_uid = order_uid
	variable.Payment.Transaction = order_uid
	var name vydumschik.Name
	variable.Delivery.Name = translit.EncodeToPCGN(name.Full_name(""))
	var address vydumschik.Address
	variable.Delivery.Address = translit.EncodeToPCGN(address.Street_address())
	towns_array := [10]string{"Tula", "Serpuhov", "Kaliningrad", "Chelyabinsk", "Mahachkala", "Ekaterinburg", "Kazan", "Saint-Petersbourg", "Krasnoyarsk", "Nizhnekamsk"}
	variable.Delivery.City = towns_array[rand.Intn(10)]
	var email string
	email_size := rand.Intn(20)
	for i := 0; i < email_size; i++ {
		ok := rand.Intn(10)
		if ok < 3 {
			email += bigchars[rand.Intn(26)]
		} else if ok >= 3 && ok <= 8 {
			email += smallchars[rand.Intn(26)]
		} else {
			email += numbers[ok]
		}
	}
	variable.Payment.Amount = rand.Intn(2000)
	variable.Payment.Payment_dt = rand.Intn(3000000000)
	variable.Payment.Delivery_cost = rand.Intn(500)
	variable.Date_created = time.Now()
	slice_of_bytes, err := json.MarshalIndent(variable, "", "\t")
	if err != nil {
		log.Print(err)
		return []byte{}
	}
	return slice_of_bytes
}
