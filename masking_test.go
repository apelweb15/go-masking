package masking

import (
	"encoding/json"
	"fmt"
	"github.com/haifenghuang/orderedmap"
	"strings"
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
	p1 := Card{
		Pin:        "123456",
		Password:   "123456",
		Cvv:        "123456",
		CardNumber: "123456",
		Signature:  "123456",
		Name:       "Card holder name",
		IsPin:      true,
		hidden:     50,
		NewPin:     "3432423",
		Parent: &Crd{
			Pin:        "123456",
			Password:   "123456",
			Cvv:        "123456",
			CardNumber: "123456",
			Signature:  "123456",
			Name:       "Parent card name",
			IsPin:      true,
			hidden:     50,
			NewPin:     "3432423",
		},
		Children: []Crd{
			{
				Pin:        "123456",
				Password:   "123456",
				Cvv:        "123456",
				CardNumber: "1212",
				Signature:  "Sign",
				Name:       "Nama child 1",
				IsPin:      false,
				hidden:     0,
				NewPin:     "23123",
			},
			{
				Pin:        "123456",
				Password:   "123456",
				Cvv:        "123456",
				CardNumber: "1212",
				Signature:  "Sign",
				Name:       "Nama child 2",
				IsPin:      false,
				hidden:     0,
				NewPin:     "23123",
			},
		},
	}
	result := MaskSensitive(p1)
	if c, ok := result.(Card); ok {
		fmt.Println(c)
		fmt.Println(c.Parent.NewPin)
		if !strings.Contains(c.Cvv, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Cvv)
		}
		if !strings.Contains(c.Pin, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Pin)
		}
		if !strings.Contains(c.Password, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Password)
		}
		if !strings.Contains(c.CardNumber, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.CardNumber)
		}
		if !strings.Contains(c.Signature, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Signature)
		}
		if c.Name != "Card holder name" {
			t.Fatalf("expected value: %v, got: %v", "Card holder name", c.Name)
		}
	} else {
		t.Fatalf("expected: %v, got: %v", "Card type", result)
	}

	result = MaskSensitive(Card{
		Pin:        "123456",
		Password:   "123456",
		Cvv:        "123456",
		CardNumber: "123456",
		Signature:  "123456",
		Name:       "Card holder name",
	}, &Config{
		MaskingKeys: []string{"pin"},
	})
	if c, ok := result.(Card); ok {
		fmt.Println(c)
		if strings.Contains(c.Cvv, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Cvv)
		}
		if !strings.Contains(c.Pin, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Pin)
		}
		if strings.Contains(c.Password, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Password)
		}
		if strings.Contains(c.CardNumber, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.CardNumber)
		}
		if strings.Contains(c.Signature, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Signature)
		}
		if c.Name != "Card holder name" {
			t.Fatalf("expected value: %v, got: %v", "Card holder name", c.Name)
		}
	} else {
		t.Fatalf("expected: %v, got: %v", "Card type", result)
	}

	result = MaskSensitive(&Card{
		Pin:        "123456",
		Password:   "123456",
		Cvv:        "123456",
		CardNumber: "123456",
		Signature:  "123456",
		Name:       "Card holder name",
	}, &Config{
		MaskingKeys: []string{"pin"},
	})
	if c, ok := result.(Card); ok {
		fmt.Println("result pointer:", c)

		if strings.Contains(c.Cvv, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Cvv)
		}
		if !strings.Contains(c.Pin, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Pin)
		}
		if strings.Contains(c.Password, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Password)
		}
		if strings.Contains(c.CardNumber, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.CardNumber)
		}
		if strings.Contains(c.Signature, "*") {
			t.Fatalf("expected mask: %v, got: %v", "**", c.Signature)
		}
		if c.Name != "Card holder name" {
			t.Fatalf("expected value: %v, got: %v", "Card holder name", c.Name)
		}
	} else {
		t.Fatalf("expected: %v, got: %v", "Card type", result)
	}

	result = MaskSensitive(`{
		"pin": "123",
		"retypePin": "123",
		"name": "Full Name"
	}`)
	if c, ok := result.(string); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

	result = MaskSensitive(`{
		"pin": "123456",
		"retypePin": "123456",
		"name": "Full Name",
		"parent": {
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 1"
			},
		"children": [
			{
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 1"
			},
			{
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 2"
			}
		]
	}`)
	if c, ok := result.(string); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

	result = MaskSensitive(map[string]interface{}{
		"pin":       "123456",
		"retypePin": "123456",
		"name":      "Full Name",
	})

	if c, ok := result.(map[string]interface{}); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

	result = MaskSensitive(map[string]interface{}{
		"pin":       "123456",
		"retypePin": "123456",
		"name":      "Full Name",
		"children": []Crd{
			{
				Pin:        "123456",
				Password:   "123456",
				Cvv:        "123456",
				CardNumber: "123456",
				Name:       "Child Name 1",
				IsPin:      true,
			},
			{
				Pin:        "123456",
				Password:   "123456",
				Cvv:        "123456",
				CardNumber: "123456",
				Name:       "Child Name 2",
				IsPin:      true,
			},
		},
		"parent": Crd{
			Pin:        "123456",
			Password:   "123456",
			Cvv:        "123456",
			CardNumber: "123456",
			Name:       "Parent Name",
			IsPin:      true,
		},
	})

	if c, ok := result.(map[string]interface{}); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

	result = MaskSensitive(map[string]interface{}{
		"pin":       "123456",
		"retypePin": "123456",
		"name":      "Full Name",
		"children": []Crd{
			{
				Pin:        "123456",
				Password:   "123456",
				Cvv:        "123456",
				CardNumber: "123456",
				Name:       "Child Name 1",
				IsPin:      true,
			},
			{
				Pin:        "123456",
				Password:   "123456",
				Cvv:        "123456",
				CardNumber: "123456",
				Name:       "Child Name 2",
				IsPin:      true,
			},
		},
		"parent": &Crd{
			Pin:        "123456",
			Password:   "123456",
			Cvv:        "123456",
			CardNumber: "123456",
			Name:       "Parent Name",
			IsPin:      true,
		},
	})

	if c, ok := result.(map[string]interface{}); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

	result = MaskSensitive(map[string]interface{}{
		"pin":       "123456",
		"retypePin": "123456",
		"name":      "Full Name",
		"children": []Crd{
			{
				Pin:        "123456",
				Password:   "123456",
				Cvv:        "123456",
				CardNumber: "123456",
				Name:       "Child Name 1",
				IsPin:      true,
			},
			{
				Pin:        "123456",
				Password:   "123456",
				Cvv:        "123456",
				CardNumber: "123456",
				Name:       "Child Name 2",
				IsPin:      true,
			},
		},
		"parent": map[string]interface{}{
			"pin":       "123456",
			"retypePin": "123456",
			"name":      "Parent Name",
		},
	})

	if c, ok := result.(map[string]interface{}); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

	result = MaskSensitive(map[string]interface{}{
		"pin":       "123456",
		"retypePin": "123456",
		"name":      "Full Name",
		"children": []map[string]interface{}{
			{
				"pin":       "123456",
				"retypePin": "123456",
				"name":      "Child Name 1",
			},
			{
				"pin":       "123456",
				"retypePin": "123456",
				"name":      "Child Name 2",
			},
		},
		"parent": map[string]interface{}{
			"pin":       "123456",
			"retypePin": "123456",
			"name":      "Parent Name",
		},
	})

	if c, ok := result.(map[string]interface{}); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

	result = MaskSensitive(`{
		"pin": "123456",
		"retypePin": "123456",
		"name": "Full Name",
		"parent": {
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 1",
				"data": null
			},
		"children": [
			{
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 1"
			},
			{
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 2"
			}
		]
	}`)
	if c, ok := result.(string); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

	result = MaskSensitive(nil)
	if result != nil {
		t.Fatalf("expected: %v, got: %v", "nil", result)
	}

	om := orderedmap.New()
	if err := json.Unmarshal([]byte(`{
		"pin": "123456",
		"retypePin": "123456",
		"name": "Full Name",
		"parent": {
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 1",
				"data": null
			},
		"children": [
			{
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 1"
			},
			{
				"pin": "123456",
				"retypePin": "123456",
				"name": "Anak 2"
			}
		]
	}`), &om); err == nil {
		result = MaskSensitive(om.String())
		fmt.Println(result)
	}

	result = MaskSensitive(map[string][]interface{}{
		"Accept": []interface{}{
			"application/json",
		},
		"Accept-Encoding": []interface{}{
			"application/json",
		},
		"Authorization": []interface{}{
			"Bearer 123456",
		},
		"X-Signature": []interface{}{
			"efdsacsdcsd",
		},
	})
	if c, ok := result.(map[string][]interface{}); ok {
		fmt.Println(c)
	} else {
		t.Fatalf("expected: %v, got: %v", "string", result)
	}

}
