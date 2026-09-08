package purchaseintent_test

import (
	"fmt"

	"github.com/zebodotdev/inttegro-sdk-go/v5/purchaseintent"
)

func ExampleStatus() {
	intent := purchaseintent.Resource{Status: purchaseintent.StatusActive}
	fmt.Println(intent.Status)
	// Output: active
}
