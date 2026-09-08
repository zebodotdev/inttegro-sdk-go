package purchaseintent_test

import (
	"fmt"

	"github.com/zebodotdev/inttegro-sdk-go/v6/purchaseintent"
)

func ExampleStatus() {
	intent := purchaseintent.PurchaseIntent{Status: purchaseintent.StatusActive}
	fmt.Println(intent.Status)
	// Output: active
}
