package masking

import (
	"fmt"
	"testing"
)

type Card struct {
	Pin        string
	Password   string
	Cvv        string
	CardNumber string
	Signature  string
	Name       string
	IsPin      bool
	hidden     float64
	NewPin     string
	Parent     *Crd
	Children   []Crd
}
type Crd struct {
	Pin        string
	Password   string
	Cvv        string
	CardNumber string
	Signature  string
	Name       string
	IsPin      bool
	hidden     float64
	NewPin     string
}

func TestMaskSensitive(t *testing.T) {

	conf := &Config{
		MaskingKeys:       DefaultMaskingKey,
		MaskingPercentage: 100,
		MaskingPosition:   0,
		MaskingCharacter:  "*",
	}
	MaskConfig = conf
	//p1 := Card{
	//	Pin:        "123456",
	//	Password:   "123456",
	//	Cvv:        "123456",
	//	CardNumber: "123456",
	//	Signature:  "123456",
	//	Name:       "Card holder name",
	//	IsPin:      true,
	//	hidden:     50,
	//	NewPin:     "3432423",
	//	Parent: &Crd{
	//		Pin:        "123456",
	//		Password:   "123456",
	//		Cvv:        "123456",
	//		CardNumber: "123456",
	//		Signature:  "123456",
	//		Name:       "Parent card name",
	//		IsPin:      true,
	//		hidden:     50,
	//		NewPin:     "3432423",
	//	},
	//	Children: []Crd{
	//		{
	//			Pin:        "123456",
	//			Password:   "123456",
	//			Cvv:        "123456",
	//			CardNumber: "1212",
	//			Signature:  "Sign",
	//			Name:       "Nama child 1",
	//			IsPin:      false,
	//			hidden:     0,
	//			NewPin:     "23123",
	//		},
	//		{
	//			Pin:        "123456",
	//			Password:   "123456",
	//			Cvv:        "123456",
	//			CardNumber: "1212",
	//			Signature:  "Sign",
	//			Name:       "Nama child 2",
	//			IsPin:      false,
	//			hidden:     0,
	//			NewPin:     "23123",
	//		},
	//	},
	//}
	//result := MaskSensitive(p1)
	//if c, ok := result.(Card); ok {
	//	fmt.Println(c)
	//	fmt.Println(c.Parent.NewPin)
	//	if !strings.Contains(c.Cvv, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Cvv)
	//	}
	//	if !strings.Contains(c.Pin, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Pin)
	//	}
	//	if !strings.Contains(c.Password, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Password)
	//	}
	//	if !strings.Contains(c.CardNumber, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.CardNumber)
	//	}
	//	if !strings.Contains(c.Signature, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Signature)
	//	}
	//	if c.Name != "Card holder name" {
	//		t.Fatalf("expected value: %v, got: %v", "Card holder name", c.Name)
	//	}
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "Card type", result)
	//}
	//
	//result = MaskSensitive(Card{
	//	Pin:        "123456",
	//	Password:   "123456",
	//	Cvv:        "123456",
	//	CardNumber: "123456",
	//	Signature:  "123456",
	//	Name:       "Card holder name",
	//}, &Config{
	//	MaskingKeys: []string{"pin"},
	//})
	//if c, ok := result.(Card); ok {
	//	fmt.Println(c)
	//	if strings.Contains(c.Cvv, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Cvv)
	//	}
	//	if !strings.Contains(c.Pin, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Pin)
	//	}
	//	if strings.Contains(c.Password, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Password)
	//	}
	//	if strings.Contains(c.CardNumber, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.CardNumber)
	//	}
	//	if strings.Contains(c.Signature, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Signature)
	//	}
	//	if c.Name != "Card holder name" {
	//		t.Fatalf("expected value: %v, got: %v", "Card holder name", c.Name)
	//	}
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "Card type", result)
	//}
	//
	//result = MaskSensitive(&Card{
	//	Pin:        "123456",
	//	Password:   "123456",
	//	Cvv:        "123456",
	//	CardNumber: "123456",
	//	Signature:  "123456",
	//	Name:       "Card holder name",
	//}, &Config{
	//	MaskingKeys: []string{"pin"},
	//})
	//if c, ok := result.(Card); ok {
	//	fmt.Println("result pointer:", c)
	//
	//	if strings.Contains(c.Cvv, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Cvv)
	//	}
	//	if !strings.Contains(c.Pin, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Pin)
	//	}
	//	if strings.Contains(c.Password, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Password)
	//	}
	//	if strings.Contains(c.CardNumber, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.CardNumber)
	//	}
	//	if strings.Contains(c.Signature, "*") {
	//		t.Fatalf("expected mask: %v, got: %v", "**", c.Signature)
	//	}
	//	if c.Name != "Card holder name" {
	//		t.Fatalf("expected value: %v, got: %v", "Card holder name", c.Name)
	//	}
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "Card type", result)
	//}
	//
	//result = MaskSensitive(`{
	//	"pin": "123",
	//	"retypePin": "123",
	//	"name": "Full Name"
	//}`)
	//if c, ok := result.(string); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}
	//
	//result = MaskSensitive(`{
	//	"pin": "123456",
	//	"retypePin": "123456",
	//	"name": "Full Name",
	//	"parent": {
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 1"
	//		},
	//	"children": [
	//		{
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 1"
	//		},
	//		{
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 2"
	//		}
	//	]
	//}`)
	//if c, ok := result.(string); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}
	//
	//result = MaskSensitive(map[string]interface{}{
	//	"pin":       "123456",
	//	"retypePin": "123456",
	//	"name":      "Full Name",
	//})
	//
	//if c, ok := result.(map[string]interface{}); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}
	//
	//result = MaskSensitive(map[string]interface{}{
	//	"pin":       "123456",
	//	"retypePin": "123456",
	//	"name":      "Full Name",
	//	"children": []Crd{
	//		{
	//			Pin:        "123456",
	//			Password:   "123456",
	//			Cvv:        "123456",
	//			CardNumber: "123456",
	//			Name:       "Child Name 1",
	//			IsPin:      true,
	//		},
	//		{
	//			Pin:        "123456",
	//			Password:   "123456",
	//			Cvv:        "123456",
	//			CardNumber: "123456",
	//			Name:       "Child Name 2",
	//			IsPin:      true,
	//		},
	//	},
	//	"parent": Crd{
	//		Pin:        "123456",
	//		Password:   "123456",
	//		Cvv:        "123456",
	//		CardNumber: "123456",
	//		Name:       "Parent Name",
	//		IsPin:      true,
	//	},
	//})
	//
	//if c, ok := result.(map[string]interface{}); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}
	//
	//result = MaskSensitive(map[string]interface{}{
	//	"pin":       "123456",
	//	"retypePin": "123456",
	//	"name":      "Full Name",
	//	"children": []Crd{
	//		{
	//			Pin:        "123456",
	//			Password:   "123456",
	//			Cvv:        "123456",
	//			CardNumber: "123456",
	//			Name:       "Child Name 1",
	//			IsPin:      true,
	//		},
	//		{
	//			Pin:        "123456",
	//			Password:   "123456",
	//			Cvv:        "123456",
	//			CardNumber: "123456",
	//			Name:       "Child Name 2",
	//			IsPin:      true,
	//		},
	//	},
	//	"parent": &Crd{
	//		Pin:        "123456",
	//		Password:   "123456",
	//		Cvv:        "123456",
	//		CardNumber: "123456",
	//		Name:       "Parent Name",
	//		IsPin:      true,
	//	},
	//})
	//
	//if c, ok := result.(map[string]interface{}); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}
	//
	//result = MaskSensitive(map[string]interface{}{
	//	"pin":       "123456",
	//	"retypePin": "123456",
	//	"name":      "Full Name",
	//	"children": []Crd{
	//		{
	//			Pin:        "123456",
	//			Password:   "123456",
	//			Cvv:        "123456",
	//			CardNumber: "123456",
	//			Name:       "Child Name 1",
	//			IsPin:      true,
	//		},
	//		{
	//			Pin:        "123456",
	//			Password:   "123456",
	//			Cvv:        "123456",
	//			CardNumber: "123456",
	//			Name:       "Child Name 2",
	//			IsPin:      true,
	//		},
	//	},
	//	"parent": map[string]interface{}{
	//		"pin":       "123456",
	//		"retypePin": "123456",
	//		"name":      "Parent Name",
	//	},
	//})
	//
	//if c, ok := result.(map[string]interface{}); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}
	//
	//result = MaskSensitive(map[string]interface{}{
	//	"pin":       "123456",
	//	"retypePin": "123456",
	//	"name":      "Full Name",
	//	"children": []map[string]interface{}{
	//		{
	//			"pin":       "123456",
	//			"retypePin": "123456",
	//			"name":      "Child Name 1",
	//		},
	//		{
	//			"pin":       "123456",
	//			"retypePin": "123456",
	//			"name":      "Child Name 2",
	//		},
	//	},
	//	"parent": map[string]interface{}{
	//		"pin":       "123456",
	//		"retypePin": "123456",
	//		"name":      "Parent Name",
	//	},
	//})
	//
	//if c, ok := result.(map[string]interface{}); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}
	//
	//result = MaskSensitive(`{
	//	"pin": "123456",
	//	"retypePin": "123456",
	//	"name": "Full Name",
	//	"parent": {
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 1",
	//			"data": null
	//		},
	//	"children": [
	//		{
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 1"
	//		},
	//		{
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 2"
	//		}
	//	]
	//}`)
	//if c, ok := result.(string); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}
	//
	//result = MaskSensitive(nil)
	//if result != nil {
	//	t.Fatalf("expected: %v, got: %v", "nil", result)
	//}
	//
	//om := orderedmap.New()
	//if err := json.Unmarshal([]byte(`{
	//	"pin": "123456",
	//	"retypePin": "123456",
	//	"name": "Full Name",
	//	"parent": {
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 1",
	//			"data": null
	//		},
	//	"children": [
	//		{
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 1"
	//		},
	//		{
	//			"pin": "123456",
	//			"retypePin": "123456",
	//			"name": "Anak 2"
	//		}
	//	]
	//}`), &om); err == nil {
	//	result = MaskSensitive(om.String())
	//	fmt.Println(result)
	//}
	//
	//result = MaskSensitive(map[string][]interface{}{
	//	"Accept": []interface{}{
	//		"application/json",
	//	},
	//	"Accept-Encoding": []interface{}{
	//		"application/json",
	//	},
	//	"Authorization": []interface{}{
	//		"Bearer 123456",
	//	},
	//	"X-Signature": []interface{}{
	//		"efdsacsdcsd",
	//	},
	//})
	//if c, ok := result.(map[string][]interface{}); ok {
	//	fmt.Println(c)
	//} else {
	//	t.Fatalf("expected: %v, got: %v", "string", result)
	//}

	result := MaskSensitive(`{"Accept":["application/json"],"Accept-Encoding":["gzip"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiYXBwSWRcIjpcInNuYXBfbWVyY2hhbnRfZ3dfMTQ1MjVcIixcInNlc3Npb25JZFwiOlwiNWM1M2E5ZTY5NGMzNGY5MWE5M2FmZDNiNmE3YTU1ZjZcIn0iLCJleHAiOjE3NTM2Nzc2MzJ9.bIXKPltEW7RmAfLNC-apBeHJJPWV5cIieuIzoTxnys8"],"Channel-Id":["5098"],"Content-Length":["883"],"Content-Type":["application/json"],"Host":["snap-mandiriva-fp-service"],"User-Agent":["go-resty/2.15.3 (https://github.com/go-resty/resty)"],"X-External-Id":["1753676732"],"X-Partner-Id":["CONSUMANDIRIVAFP02HRUOZ"],"X-Signature":["kJOnBET6QyKx5NLW2dsTLqW3LRpEDWnuHw46R1/2PP4zAuKvcuJxLO+hdzD9Eblm+3swCxYCWkkHF+C2rsuzO3y5y6bTFpl3TNujjaWA6hfogmohWq8Ow2KLSzCYwr9zkIAb1PSR4kpWpMXRsaJTXlhK3L8D6kzDt0Hl7PNRlnoTegFkgoHsZB6/LaN0CP/jgFxqgqtzRZKIQWSAjFzX3i/v+9f/bJtRP0mw8w54b7cSCK8LeIjzOpr9Yc6dfUaH8R97pn9ZMguC0l1U9HWUttRWBYG+6QTYLSb73GGsuLYZqgdjmUIvwI4dmu1/tYcx02dllHYDNdugM4q0pImG1XJ62yXnRfdNc40KTN8hBFv9Hh0JwX/5QcMb9Vbng8Boa4shG2qog9PG82sU+1+bKzgNETMIl/3jXy37oq4Qc3gA05VobEbh3BgouDbkOS149MWDijluXclWKmdFEcy1MCYH9C316WxxA/NeTBpXnDshm6oW46QJszsI7WqSNe4spCrNbJ7g4fTGgLMClu6rqGKvH+03Jw8ZYgmogcBBtKwqiopoyF/X8wCtQxE3VKTOSFq3opn4qY3sXEwLCH5TLw+zP+1DTr0nlgmDGiIQJ3fNDXnnnVNOQoLEn+odM0I4JxP631s+aAXsrdmf5llYp2u7M4AaZC1/O7SlA67fV2s="],"X-Timestamp":["2025-07-28T04:25:32+00:00"]}`)
	if c, ok := result.(string); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

}
